// Copyright 2017 Capsule8, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sensor

import (
	"reflect"
	"sync"
	"unicode"

	api "github.com/capsule8/capsule8/api/v0"

	"github.com/capsule8/capsule8/pkg/expression"
	"github.com/capsule8/capsule8/pkg/sys/perf"

	"github.com/gobwas/glob"
	"github.com/golang/glog"

	"golang.org/x/sys/unix"
)

var containerEventTypes = expression.FieldTypeMap{
	"name":             int32(api.ValueType_STRING),
	"image_id":         int32(api.ValueType_STRING),
	"image_name":       int32(api.ValueType_STRING),
	"host_pid":         int32(api.ValueType_SINT32),
	"exit_code":        int32(api.ValueType_SINT32),
	"exit_status":      int32(api.ValueType_UINT32),
	"exit_signal":      int32(api.ValueType_UINT32),
	"exit_core_dumped": int32(api.ValueType_BOOL),
}

type containerCache struct {
	sync.Mutex
	cache map[string]*ContainerInfo

	sensor *Sensor

	// These are external event IDs registered with the sensor's event
	// monitor instance. The cache will enqueue these events as appropriate
	// as the cache is updated.
	containerCreatedEventID   uint64 // api.ContainerEventType_CONTAINER_EVENT_TYPE_CREATED
	containerRunningEventID   uint64 // api.ContainerEventType_CONTAINER_EVENT_TYPE_RUNNING
	containerExitedEventID    uint64 // api.ContainerEventType_CONTAINER_EVENT_TYPE_EXITED
	containerDestroyedEventID uint64 // api.ContainerEventType_CONTAINER_EVENT_TYPE_DESTROYED
	containerUpdatedEventID   uint64 // api.ContainerEventType_CONTAINER_EVENT_TYPE_UPDATED
}

// ContainerState represents the state of a container (created, running, etc.)
type ContainerState uint

const (
	// ContainerStateUnknown indicates that the container is in an unknown
	// state.
	ContainerStateUnknown ContainerState = iota

	// ContainerStateCreated indicates the container exists, but is not
	// running.
	ContainerStateCreated

	// ContainerStatePaused indicates the container is paused.
	ContainerStatePaused

	// ContainerStateRunning indicates the container is running.
	ContainerStateRunning

	// ContainerStateRestarting indicates the container is in the process
	// of restarting.
	ContainerStateRestarting

	// ContainerStateExited indicates the container has exited.
	ContainerStateExited

	// ContainerStateRemoving indicates the container is being removed.
	ContainerStateRemoving
)

// ContainerStateNames is a mapping of container states to printable names.
var ContainerStateNames = map[ContainerState]string{
	ContainerStateCreated:    "created",
	ContainerStateRestarting: "restarting",
	ContainerStateRunning:    "running",
	ContainerStateRemoving:   "removing",
	ContainerStatePaused:     "paused",
	ContainerStateExited:     "exited",
}

// ContainerRuntime represents the runtime used to manager a container
type ContainerRuntime uint

const (
	// ContainerRuntimeUnknown means the container runtime managing the
	// container is unknown. Information about the container comes from
	// runc, the kernel, or other generic sources.
	ContainerRuntimeUnknown ContainerRuntime = iota

	// ContainerRuntimeDocker means the container is managed by Docker.
	ContainerRuntimeDocker
)

// ContainerInfo records interesting information known about a container.
type ContainerInfo struct {
	cache *containerCache // the cache to which this info belongs

	ID        string
	Name      string
	ImageID   string
	ImageName string

	Pid      int
	ExitCode int

	Runtime ContainerRuntime
	State   ContainerState

	JSONConfig string
	OCIConfig  string
}

func newContainerCache(sensor *Sensor) *containerCache {
	cache := &containerCache{
		cache:  make(map[string]*ContainerInfo),
		sensor: sensor,
	}

	var err error
	cache.containerCreatedEventID, err = sensor.monitor.RegisterExternalEvent(
		"CONTAINER_CREATED",
		cache.decodeContainerCreatedEvent,
		containerEventTypes,
	)
	if err != nil {
		glog.Fatalf("Failed to register external event: %s", err)
	}

	cache.containerRunningEventID, err = sensor.monitor.RegisterExternalEvent(
		"CONTAINER_RUNNING",
		cache.decodeContainerRunningEvent,
		containerEventTypes,
	)
	if err != nil {
		glog.Fatalf("Failed to register external event: %s", err)
	}

	cache.containerExitedEventID, err = sensor.monitor.RegisterExternalEvent(
		"CONTAINER_EXITED",
		cache.decodeContainerExitedEvent,
		containerEventTypes,
	)
	if err != nil {
		glog.Fatalf("Failed to register external event: %s", err)
	}

	cache.containerDestroyedEventID, err = sensor.monitor.RegisterExternalEvent(
		"CONTAINER_DESTROYED",
		cache.decodeContainerDestroyedEvent,
		containerEventTypes,
	)
	if err != nil {
		glog.Fatalf("Failed to register external event: %s", err)
	}

	cache.containerUpdatedEventID, err = sensor.monitor.RegisterExternalEvent(
		"CONTAINER_UPDATED",
		cache.decodeContainerUpdatedEvent,
		containerEventTypes,
	)
	if err != nil {
		glog.Fatalf("Failed to register extern event: %s", err)
	}

	return cache
}

func (cc *containerCache) deleteContainer(
	containerID string,
	runtime ContainerRuntime,
	sampleID perf.SampleID,
) {
	cc.Lock()
	info, ok := cc.cache[containerID]
	if ok {
		if runtime == info.Runtime {
			delete(cc.cache, containerID)
		} else {
			ok = false
		}
	}
	cc.Unlock()

	if ok {
		glog.V(2).Infof("Sending CONTAINER_DESTROYED for %s", info.ID)
		cc.enqueueContainerEvent(cc.containerDestroyedEventID,
			sampleID, info)
	}
}

func (cc *containerCache) newContainerInfo(containerID string) *ContainerInfo {
	return &ContainerInfo{
		cache:   cc,
		ID:      containerID,
		Runtime: ContainerRuntimeUnknown,
	}
}

func (cc *containerCache) lookupContainer(containerID string, create bool) *ContainerInfo {
	cc.Lock()
	defer cc.Unlock()

	info := cc.cache[containerID]
	if info == nil && create {
		info = cc.newContainerInfo(containerID)
		cc.cache[containerID] = info
	}

	return info
}

func (cc *containerCache) enqueueContainerEvent(
	eventID uint64,
	sampleID perf.SampleID,
	info *ContainerInfo,
) error {
	ws := unix.WaitStatus(info.ExitCode)
	data := map[string]interface{}{
		"container_id":     info.ID,
		"name":             info.Name,
		"image_id":         info.ImageID,
		"image_name":       info.ImageName,
		"host_pid":         int32(info.Pid),
		"exit_code":        int32(info.ExitCode),
		"exit_status":      uint32(0),
		"exit_signal":      uint32(0),
		"exit_core_dumped": ws.CoreDump(),
	}

	if ws.Exited() {
		data["exit_status"] = uint32(ws.ExitStatus())
	}
	if ws.Signaled() {
		data["exit_signal"] = uint32(ws.Signal())
	}

	return cc.sensor.monitor.EnqueueExternalSample(eventID, sampleID, data)
}

func (cc *containerCache) newContainerEvent(
	sample *perf.SampleRecord,
	data perf.TraceEventSampleData,
	eventType api.ContainerEventType,
) (*api.TelemetryEvent, error) {
	cev := &api.ContainerEvent{
		Type:           eventType,
		ImageId:        data["image_id"].(string),
		ImageName:      data["image_name"].(string),
		HostPid:        data["host_pid"].(int32),
		ExitCode:       data["exit_code"].(int32),
		ExitStatus:     data["exit_status"].(uint32),
		ExitSignal:     data["exit_signal"].(uint32),
		ExitCoreDumped: data["exit_core_dumped"].(bool),
	}

	if s, ok := data["docker_config"].(string); ok && len(s) > 0 {
		cev.DockerConfigJson = s
	}
	if s, ok := data["oci_config"].(string); ok && len(s) > 0 {
		cev.OciConfigJson = s
	}

	event := cc.sensor.NewEventFromSample(sample, data)
	event.ContainerId = data["container_id"].(string)
	event.Event = &api.TelemetryEvent_Container{
		Container: cev,
	}
	return event, nil
}

func (cc *containerCache) decodeContainerCreatedEvent(
	sample *perf.SampleRecord,
	data perf.TraceEventSampleData,
) (interface{}, error) {
	return cc.newContainerEvent(sample, data,
		api.ContainerEventType_CONTAINER_EVENT_TYPE_CREATED)
}

func (cc *containerCache) decodeContainerRunningEvent(
	sample *perf.SampleRecord,
	data perf.TraceEventSampleData,
) (interface{}, error) {
	return cc.newContainerEvent(sample, data,
		api.ContainerEventType_CONTAINER_EVENT_TYPE_RUNNING)
}

func (cc *containerCache) decodeContainerExitedEvent(
	sample *perf.SampleRecord,
	data perf.TraceEventSampleData,
) (interface{}, error) {
	return cc.newContainerEvent(sample, data,
		api.ContainerEventType_CONTAINER_EVENT_TYPE_EXITED)
}

func (cc *containerCache) decodeContainerDestroyedEvent(
	sample *perf.SampleRecord,
	data perf.TraceEventSampleData,
) (interface{}, error) {
	return cc.newContainerEvent(sample, data,
		api.ContainerEventType_CONTAINER_EVENT_TYPE_DESTROYED)
}

func (cc *containerCache) decodeContainerUpdatedEvent(
	sample *perf.SampleRecord,
	data perf.TraceEventSampleData,
) (interface{}, error) {
	return cc.newContainerEvent(sample, data,
		api.ContainerEventType_CONTAINER_EVENT_TYPE_UPDATED)
}

// Update updates the data cached for a container with new information. Some
// new information may trigger telemetry events to fire.
func (info *ContainerInfo) Update(
	runtime ContainerRuntime,
	sampleID perf.SampleID,
	data map[string]interface{},
) {
	if info.Runtime == ContainerRuntimeUnknown {
		info.Runtime = runtime
	}

	oldState := info.State
	dataChanged := false

	s := reflect.ValueOf(info).Elem()
	t := s.Type()
	for i := t.NumField() - 1; i >= 0; i-- {
		f := t.Field(i)
		if !unicode.IsUpper(rune(f.Name[0])) {
			continue
		}
		v, ok := data[f.Name]
		if !ok {
			continue
		}
		if !reflect.TypeOf(v).AssignableTo(f.Type) {
			glog.Fatalf("Cannot assign %v to %s %s",
				v, f.Name, f.Type)
		}

		if s.Field(i).Interface() != v {
			if f.Name != "State" {
				dataChanged = true
			} else if info.Runtime != runtime {
				// Only allow state changes from the runtime
				// known to be managing the container.
				continue
			}
			s.Field(i).Set(reflect.ValueOf(v))
		}
	}

	if info.State != oldState {
		if oldState < ContainerStateCreated {
			glog.V(2).Infof("Sending CONTAINER_CREATED for %s", info.ID)
			info.cache.enqueueContainerEvent(
				info.cache.containerCreatedEventID, sampleID,
				info)
		}
		if oldState < ContainerStateRunning &&
			info.State >= ContainerStateRunning {
			glog.V(2).Infof("Sending CONTAINER_RUNNING for %s", info.ID)
			info.cache.enqueueContainerEvent(
				info.cache.containerRunningEventID, sampleID,
				info)
		}
		if oldState < ContainerStateRestarting &&
			info.State >= ContainerStateRestarting {
			glog.V(2).Infof("Sending CONTAINER_EXITED for %s", info.ID)
			info.cache.enqueueContainerEvent(
				info.cache.containerExitedEventID,
				sampleID, info)
		}
	} else if dataChanged {
		glog.V(2).Infof("Sending CONTAINER_UPDATED for %s", info.ID)
		info.cache.enqueueContainerEvent(
			info.cache.containerUpdatedEventID, sampleID, info)
	}
}

func registerContainerEvents(
	sensor *Sensor,
	eventMap subscriptionMap,
	events []*api.ContainerEventFilter,
) {
	var (
		filters       [6]*api.Expression
		subscriptions [6]*subscription
	)

	for _, cef := range events {
		t := cef.Type
		if t < 1 || t > 5 {
			continue
		}
		if subscriptions[t] == nil {
			var eventID uint64
			switch t {
			case api.ContainerEventType_CONTAINER_EVENT_TYPE_CREATED:
				eventID = sensor.containerCache.containerCreatedEventID
			case api.ContainerEventType_CONTAINER_EVENT_TYPE_RUNNING:
				eventID = sensor.containerCache.containerRunningEventID
			case api.ContainerEventType_CONTAINER_EVENT_TYPE_EXITED:
				eventID = sensor.containerCache.containerExitedEventID
			case api.ContainerEventType_CONTAINER_EVENT_TYPE_DESTROYED:
				eventID = sensor.containerCache.containerDestroyedEventID
			case api.ContainerEventType_CONTAINER_EVENT_TYPE_UPDATED:
				eventID = sensor.containerCache.containerUpdatedEventID
			}
			subscriptions[t] = eventMap.subscribe(eventID)
		}
		filters[t] = expression.LogicalOr(filters[t], cef.FilterExpression)
	}

	for i, s := range subscriptions {
		if filters[i] == nil {
			// No filter, no problem
			continue
		}

		expr, err := expression.NewExpression(filters[i])
		if err != nil {
			// Bad filter. Remove subscription
			glog.V(1).Infof("Invalid container filter expression: %s", err)
			eventMap.unsubscribe(s.eventID)
			continue
		}

		err = expr.Validate(containerEventTypes)
		if err != nil {
			// Bad filter. Remove subscription
			glog.V(1).Infof("Invalid container filter expression: %s", err)
			eventMap.unsubscribe(s.eventID)
			continue
		}

		s.filter = expr
	}
}

///////////////////////////////////////////////////////////////////////////////

func newContainerFilter(ecf *api.ContainerFilter) *containerFilter {
	cf := &containerFilter{}

	for _, v := range ecf.Ids {
		cf.addContainerID(v)
	}

	for _, v := range ecf.Names {
		cf.addContainerName(v)
	}

	for _, v := range ecf.ImageIds {
		cf.addImageID(v)
	}

	for _, v := range ecf.ImageNames {
		cf.addImageName(v)
	}

	return cf
}

type containerFilter struct {
	containerIds   map[string]bool
	containerNames map[string]bool
	imageIds       map[string]bool
	imageGlobs     map[string]glob.Glob
}

func (c *containerFilter) addContainerID(cid string) {
	if len(cid) > 0 {
		if c.containerIds == nil {
			c.containerIds = make(map[string]bool)
		}
		c.containerIds[cid] = true
	}
}

func (c *containerFilter) removeContainerID(cid string) {
	delete(c.containerIds, cid)
}

func (c *containerFilter) addContainerName(cname string) {
	if len(cname) > 0 {
		if c.containerNames == nil {
			c.containerNames = make(map[string]bool)
		}
		c.containerNames[cname] = true
	}
}

func (c *containerFilter) addImageID(iid string) {
	if len(iid) > 0 {
		if c.imageIds == nil {
			c.imageIds = make(map[string]bool)
		}
		c.imageIds[iid] = true
	}
}

func (c *containerFilter) addImageName(iname string) {
	if len(iname) > 0 {
		if c.imageGlobs == nil {
			c.imageGlobs = make(map[string]glob.Glob)
		} else {
			_, ok := c.imageGlobs[iname]
			if ok {
				return
			}
		}

		g, err := glob.Compile(iname, '/')
		if err == nil {
			c.imageGlobs[iname] = g
		}
	}
}

func (c *containerFilter) FilterFunc(i interface{}) bool {
	event := i.(Event)
	e := event.Event

	//
	// Fast path: Check if containerId is in containerIds map
	//
	if c.containerIds != nil && c.containerIds[e.ContainerId] {
		return true
	}

	switch e.Event.(type) {
	case *api.TelemetryEvent_Container:
		cev := e.GetContainer()

		//
		// Slow path: Check if other identifiers are in maps. If they
		// are, add the containerId to containerIds map to take fast
		// path next time.
		//

		if c.containerNames[cev.Name] {
			c.addContainerID(e.ContainerId)
			return true
		}

		if c.imageIds[cev.ImageId] {
			c.addContainerID(e.ContainerId)
			return true
		}

		if c.imageGlobs != nil && cev.ImageName != "" {
			for _, g := range c.imageGlobs {
				if g.Match(cev.ImageName) {
					c.addContainerID(e.ContainerId)
					return true
				}
			}
		}
	}

	return false
}

func (c *containerFilter) DoFunc(i interface{}) {
	event := i.(Event)
	e := event.Event

	switch e.Event.(type) {
	case *api.TelemetryEvent_Container:
		cev := e.GetContainer()
		if cev.Type == api.ContainerEventType_CONTAINER_EVENT_TYPE_DESTROYED {
			c.removeContainerID(e.ContainerId)
		}
	}
}
