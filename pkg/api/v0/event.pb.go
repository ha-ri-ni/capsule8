// Code generated by protoc-gen-go. DO NOT EDIT.
// source: event.proto

package v0

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ContainerEventType int32

const (
	ContainerEventType_CONTAINER_EVENT_TYPE_UNKNOWN   ContainerEventType = 0
	ContainerEventType_CONTAINER_EVENT_TYPE_CREATED   ContainerEventType = 1
	ContainerEventType_CONTAINER_EVENT_TYPE_RUNNING   ContainerEventType = 2
	ContainerEventType_CONTAINER_EVENT_TYPE_EXITED    ContainerEventType = 3
	ContainerEventType_CONTAINER_EVENT_TYPE_DESTROYED ContainerEventType = 4
)

var ContainerEventType_name = map[int32]string{
	0: "CONTAINER_EVENT_TYPE_UNKNOWN",
	1: "CONTAINER_EVENT_TYPE_CREATED",
	2: "CONTAINER_EVENT_TYPE_RUNNING",
	3: "CONTAINER_EVENT_TYPE_EXITED",
	4: "CONTAINER_EVENT_TYPE_DESTROYED",
}
var ContainerEventType_value = map[string]int32{
	"CONTAINER_EVENT_TYPE_UNKNOWN":   0,
	"CONTAINER_EVENT_TYPE_CREATED":   1,
	"CONTAINER_EVENT_TYPE_RUNNING":   2,
	"CONTAINER_EVENT_TYPE_EXITED":    3,
	"CONTAINER_EVENT_TYPE_DESTROYED": 4,
}

func (x ContainerEventType) String() string {
	return proto.EnumName(ContainerEventType_name, int32(x))
}
func (ContainerEventType) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

type ProcessEventType int32

const (
	ProcessEventType_PROCESS_EVENT_TYPE_UNKNOWN ProcessEventType = 0
	ProcessEventType_PROCESS_EVENT_TYPE_FORK    ProcessEventType = 1
	ProcessEventType_PROCESS_EVENT_TYPE_EXEC    ProcessEventType = 2
	ProcessEventType_PROCESS_EVENT_TYPE_EXIT    ProcessEventType = 3
)

var ProcessEventType_name = map[int32]string{
	0: "PROCESS_EVENT_TYPE_UNKNOWN",
	1: "PROCESS_EVENT_TYPE_FORK",
	2: "PROCESS_EVENT_TYPE_EXEC",
	3: "PROCESS_EVENT_TYPE_EXIT",
}
var ProcessEventType_value = map[string]int32{
	"PROCESS_EVENT_TYPE_UNKNOWN": 0,
	"PROCESS_EVENT_TYPE_FORK":    1,
	"PROCESS_EVENT_TYPE_EXEC":    2,
	"PROCESS_EVENT_TYPE_EXIT":    3,
}

func (x ProcessEventType) String() string {
	return proto.EnumName(ProcessEventType_name, int32(x))
}
func (ProcessEventType) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

type SyscallEventType int32

const (
	SyscallEventType_SYSCALL_EVENT_TYPE_UNKNOWN SyscallEventType = 0
	SyscallEventType_SYSCALL_EVENT_TYPE_ENTER   SyscallEventType = 1
	SyscallEventType_SYSCALL_EVENT_TYPE_EXIT    SyscallEventType = 2
)

var SyscallEventType_name = map[int32]string{
	0: "SYSCALL_EVENT_TYPE_UNKNOWN",
	1: "SYSCALL_EVENT_TYPE_ENTER",
	2: "SYSCALL_EVENT_TYPE_EXIT",
}
var SyscallEventType_value = map[string]int32{
	"SYSCALL_EVENT_TYPE_UNKNOWN": 0,
	"SYSCALL_EVENT_TYPE_ENTER":   1,
	"SYSCALL_EVENT_TYPE_EXIT":    2,
}

func (x SyscallEventType) String() string {
	return proto.EnumName(SyscallEventType_name, int32(x))
}
func (SyscallEventType) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

type FileEventType int32

const (
	FileEventType_FILE_EVENT_TYPE_UNKNOWN FileEventType = 0
	FileEventType_FILE_EVENT_TYPE_OPEN    FileEventType = 1
)

var FileEventType_name = map[int32]string{
	0: "FILE_EVENT_TYPE_UNKNOWN",
	1: "FILE_EVENT_TYPE_OPEN",
}
var FileEventType_value = map[string]int32{
	"FILE_EVENT_TYPE_UNKNOWN": 0,
	"FILE_EVENT_TYPE_OPEN":    1,
}

func (x FileEventType) String() string {
	return proto.EnumName(FileEventType_name, int32(x))
}
func (FileEventType) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{3} }

type Event struct {
	// Unique identifier for the event
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// Unique process identifier associated with the event to differentiate
	// reused values of the pid below.
	ProcessId string `protobuf:"bytes,2,opt,name=process_id,json=processId" json:"process_id,omitempty"`
	// Unix pid of the process associated with the event
	ProcessPid int32 `protobuf:"zigzag32,3,opt,name=process_pid,json=processPid" json:"process_pid,omitempty"`
	// Container identifier associated with the event
	ContainerId string `protobuf:"bytes,4,opt,name=container_id,json=containerId" json:"container_id,omitempty"`
	// Sensor identifier of the sensor instance that observed the event
	SensorId string `protobuf:"bytes,5,opt,name=sensor_id,json=sensorId" json:"sensor_id,omitempty"`
	// Sequence number from some unspecified starting point unique
	// to the Sensor. Provides a strict linear ordering of events with
	// the same sensor_id where no two events can have the same sequence
	// number. If it is present, it must be greater than zero. A zero
	// value indicates that there is no sequence number associated with
	// the event.
	SensorSequenceNumber uint64 `protobuf:"varint,6,opt,name=sensor_sequence_number,json=sensorSequenceNumber" json:"sensor_sequence_number,omitempty"`
	// Monotonic nanosecond timestamp from some unspecified starting
	// point unique to the Sensor. Can only be used to calculate time
	// intervals between events with the same sensor_id.
	SensorMonotimeNanos int64 `protobuf:"varint,7,opt,name=sensor_monotime_nanos,json=sensorMonotimeNanos" json:"sensor_monotime_nanos,omitempty"`
	// Types that are valid to be assigned to Event:
	//	*Event_Syscall
	//	*Event_Process
	//	*Event_File
	//	*Event_Container
	//	*Event_Chargen
	//	*Event_Ticker
	Event isEvent_Event `protobuf_oneof:"event"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

type isEvent_Event interface {
	isEvent_Event()
}

type Event_Syscall struct {
	Syscall *SyscallEvent `protobuf:"bytes,10,opt,name=syscall,oneof"`
}
type Event_Process struct {
	Process *ProcessEvent `protobuf:"bytes,11,opt,name=process,oneof"`
}
type Event_File struct {
	File *FileEvent `protobuf:"bytes,12,opt,name=file,oneof"`
}
type Event_Container struct {
	Container *ContainerEvent `protobuf:"bytes,20,opt,name=container,oneof"`
}
type Event_Chargen struct {
	Chargen *ChargenEvent `protobuf:"bytes,100,opt,name=chargen,oneof"`
}
type Event_Ticker struct {
	Ticker *TickerEvent `protobuf:"bytes,101,opt,name=ticker,oneof"`
}

func (*Event_Syscall) isEvent_Event()   {}
func (*Event_Process) isEvent_Event()   {}
func (*Event_File) isEvent_Event()      {}
func (*Event_Container) isEvent_Event() {}
func (*Event_Chargen) isEvent_Event()   {}
func (*Event_Ticker) isEvent_Event()    {}

func (m *Event) GetEvent() isEvent_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *Event) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Event) GetProcessId() string {
	if m != nil {
		return m.ProcessId
	}
	return ""
}

func (m *Event) GetProcessPid() int32 {
	if m != nil {
		return m.ProcessPid
	}
	return 0
}

func (m *Event) GetContainerId() string {
	if m != nil {
		return m.ContainerId
	}
	return ""
}

func (m *Event) GetSensorId() string {
	if m != nil {
		return m.SensorId
	}
	return ""
}

func (m *Event) GetSensorSequenceNumber() uint64 {
	if m != nil {
		return m.SensorSequenceNumber
	}
	return 0
}

func (m *Event) GetSensorMonotimeNanos() int64 {
	if m != nil {
		return m.SensorMonotimeNanos
	}
	return 0
}

func (m *Event) GetSyscall() *SyscallEvent {
	if x, ok := m.GetEvent().(*Event_Syscall); ok {
		return x.Syscall
	}
	return nil
}

func (m *Event) GetProcess() *ProcessEvent {
	if x, ok := m.GetEvent().(*Event_Process); ok {
		return x.Process
	}
	return nil
}

func (m *Event) GetFile() *FileEvent {
	if x, ok := m.GetEvent().(*Event_File); ok {
		return x.File
	}
	return nil
}

func (m *Event) GetContainer() *ContainerEvent {
	if x, ok := m.GetEvent().(*Event_Container); ok {
		return x.Container
	}
	return nil
}

func (m *Event) GetChargen() *ChargenEvent {
	if x, ok := m.GetEvent().(*Event_Chargen); ok {
		return x.Chargen
	}
	return nil
}

func (m *Event) GetTicker() *TickerEvent {
	if x, ok := m.GetEvent().(*Event_Ticker); ok {
		return x.Ticker
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Event) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Event_OneofMarshaler, _Event_OneofUnmarshaler, _Event_OneofSizer, []interface{}{
		(*Event_Syscall)(nil),
		(*Event_Process)(nil),
		(*Event_File)(nil),
		(*Event_Container)(nil),
		(*Event_Chargen)(nil),
		(*Event_Ticker)(nil),
	}
}

func _Event_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Event)
	// event
	switch x := m.Event.(type) {
	case *Event_Syscall:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Syscall); err != nil {
			return err
		}
	case *Event_Process:
		b.EncodeVarint(11<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Process); err != nil {
			return err
		}
	case *Event_File:
		b.EncodeVarint(12<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.File); err != nil {
			return err
		}
	case *Event_Container:
		b.EncodeVarint(20<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Container); err != nil {
			return err
		}
	case *Event_Chargen:
		b.EncodeVarint(100<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Chargen); err != nil {
			return err
		}
	case *Event_Ticker:
		b.EncodeVarint(101<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Ticker); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Event.Event has unexpected type %T", x)
	}
	return nil
}

func _Event_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Event)
	switch tag {
	case 10: // event.syscall
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SyscallEvent)
		err := b.DecodeMessage(msg)
		m.Event = &Event_Syscall{msg}
		return true, err
	case 11: // event.process
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ProcessEvent)
		err := b.DecodeMessage(msg)
		m.Event = &Event_Process{msg}
		return true, err
	case 12: // event.file
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(FileEvent)
		err := b.DecodeMessage(msg)
		m.Event = &Event_File{msg}
		return true, err
	case 20: // event.container
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ContainerEvent)
		err := b.DecodeMessage(msg)
		m.Event = &Event_Container{msg}
		return true, err
	case 100: // event.chargen
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ChargenEvent)
		err := b.DecodeMessage(msg)
		m.Event = &Event_Chargen{msg}
		return true, err
	case 101: // event.ticker
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TickerEvent)
		err := b.DecodeMessage(msg)
		m.Event = &Event_Ticker{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Event_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Event)
	// event
	switch x := m.Event.(type) {
	case *Event_Syscall:
		s := proto.Size(x.Syscall)
		n += proto.SizeVarint(10<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_Process:
		s := proto.Size(x.Process)
		n += proto.SizeVarint(11<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_File:
		s := proto.Size(x.File)
		n += proto.SizeVarint(12<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_Container:
		s := proto.Size(x.Container)
		n += proto.SizeVarint(20<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_Chargen:
		s := proto.Size(x.Chargen)
		n += proto.SizeVarint(100<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Event_Ticker:
		s := proto.Size(x.Ticker)
		n += proto.SizeVarint(101<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type ChargenEvent struct {
	// Index of the first character in this Event in relation to all of
	// the characters that have been generated in this stream.
	Index uint64 `protobuf:"varint,1,opt,name=index" json:"index,omitempty"`
	// The next one or more characters in the autogenerated stream
	Characters string `protobuf:"bytes,2,opt,name=characters" json:"characters,omitempty"`
}

func (m *ChargenEvent) Reset()                    { *m = ChargenEvent{} }
func (m *ChargenEvent) String() string            { return proto.CompactTextString(m) }
func (*ChargenEvent) ProtoMessage()               {}
func (*ChargenEvent) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *ChargenEvent) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *ChargenEvent) GetCharacters() string {
	if m != nil {
		return m.Characters
	}
	return ""
}

type TickerEvent struct {
	// The number of seconds elapsed since January 1, 1970 UTC.
	//
	// https://golang.org/pkg/time/#Time.Unix
	Seconds int64 `protobuf:"varint,1,opt,name=seconds" json:"seconds,omitempty"`
	// The number of nanoseconds elapsed since January 1, 1970 UTC
	//
	// https://golang.org/pkg/time/#Time.UnixNano
	Nanoseconds int64 `protobuf:"varint,2,opt,name=nanoseconds" json:"nanoseconds,omitempty"`
}

func (m *TickerEvent) Reset()                    { *m = TickerEvent{} }
func (m *TickerEvent) String() string            { return proto.CompactTextString(m) }
func (*TickerEvent) ProtoMessage()               {}
func (*TickerEvent) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

func (m *TickerEvent) GetSeconds() int64 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

func (m *TickerEvent) GetNanoseconds() int64 {
	if m != nil {
		return m.Nanoseconds
	}
	return 0
}

// ContainerEvent describes a Docker container or Rkt App lifecycle event
type ContainerEvent struct {
	Type ContainerEventType `protobuf:"varint,1,opt,name=type,enum=capsule8.v0.ContainerEventType" json:"type,omitempty"`
	Name string             `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// Unique identifier of the container image
	ImageId string `protobuf:"bytes,10,opt,name=image_id,json=imageId" json:"image_id,omitempty"`
	//
	// Name of the container image (i.e. "busybox" or
	// "gcr.io/google_containers/nginx-ingress-controller")
	//
	ImageName string `protobuf:"bytes,11,opt,name=image_name,json=imageName" json:"image_name,omitempty"`
	// Host process identifier of the container's init process.
	HostPid int32 `protobuf:"zigzag32,20,opt,name=host_pid,json=hostPid" json:"host_pid,omitempty"`
	// The exit code of the container if it has exited
	ExitCode int32 `protobuf:"zigzag32,30,opt,name=exit_code,json=exitCode" json:"exit_code,omitempty"`
}

func (m *ContainerEvent) Reset()                    { *m = ContainerEvent{} }
func (m *ContainerEvent) String() string            { return proto.CompactTextString(m) }
func (*ContainerEvent) ProtoMessage()               {}
func (*ContainerEvent) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{3} }

func (m *ContainerEvent) GetType() ContainerEventType {
	if m != nil {
		return m.Type
	}
	return ContainerEventType_CONTAINER_EVENT_TYPE_UNKNOWN
}

func (m *ContainerEvent) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ContainerEvent) GetImageId() string {
	if m != nil {
		return m.ImageId
	}
	return ""
}

func (m *ContainerEvent) GetImageName() string {
	if m != nil {
		return m.ImageName
	}
	return ""
}

func (m *ContainerEvent) GetHostPid() int32 {
	if m != nil {
		return m.HostPid
	}
	return 0
}

func (m *ContainerEvent) GetExitCode() int32 {
	if m != nil {
		return m.ExitCode
	}
	return 0
}

type ProcessEvent struct {
	Type ProcessEventType `protobuf:"varint,1,opt,name=type,enum=capsule8.v0.ProcessEventType" json:"type,omitempty"`
	// Optional
	ForkChildPid int32 `protobuf:"zigzag32,10,opt,name=fork_child_pid,json=forkChildPid" json:"fork_child_pid,omitempty"`
	// Optional
	ExecFilename    string   `protobuf:"bytes,20,opt,name=exec_filename,json=execFilename" json:"exec_filename,omitempty"`
	ExecCommandLine []string `protobuf:"bytes,21,rep,name=exec_command_line,json=execCommandLine" json:"exec_command_line,omitempty"`
	// Optional
	ExitCode int32 `protobuf:"zigzag32,30,opt,name=exit_code,json=exitCode" json:"exit_code,omitempty"`
}

func (m *ProcessEvent) Reset()                    { *m = ProcessEvent{} }
func (m *ProcessEvent) String() string            { return proto.CompactTextString(m) }
func (*ProcessEvent) ProtoMessage()               {}
func (*ProcessEvent) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{4} }

func (m *ProcessEvent) GetType() ProcessEventType {
	if m != nil {
		return m.Type
	}
	return ProcessEventType_PROCESS_EVENT_TYPE_UNKNOWN
}

func (m *ProcessEvent) GetForkChildPid() int32 {
	if m != nil {
		return m.ForkChildPid
	}
	return 0
}

func (m *ProcessEvent) GetExecFilename() string {
	if m != nil {
		return m.ExecFilename
	}
	return ""
}

func (m *ProcessEvent) GetExecCommandLine() []string {
	if m != nil {
		return m.ExecCommandLine
	}
	return nil
}

func (m *ProcessEvent) GetExitCode() int32 {
	if m != nil {
		return m.ExitCode
	}
	return 0
}

type SyscallEvent struct {
	Type SyscallEventType `protobuf:"varint,1,opt,name=type,enum=capsule8.v0.SyscallEventType" json:"type,omitempty"`
	Id   int64            `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	Arg0 uint64           `protobuf:"varint,10,opt,name=arg0" json:"arg0,omitempty"`
	Arg1 uint64           `protobuf:"varint,11,opt,name=arg1" json:"arg1,omitempty"`
	Arg2 uint64           `protobuf:"varint,12,opt,name=arg2" json:"arg2,omitempty"`
	Arg3 uint64           `protobuf:"varint,13,opt,name=arg3" json:"arg3,omitempty"`
	Arg4 uint64           `protobuf:"varint,14,opt,name=arg4" json:"arg4,omitempty"`
	Arg5 uint64           `protobuf:"varint,15,opt,name=arg5" json:"arg5,omitempty"`
	Ret  int64            `protobuf:"varint,20,opt,name=ret" json:"ret,omitempty"`
}

func (m *SyscallEvent) Reset()                    { *m = SyscallEvent{} }
func (m *SyscallEvent) String() string            { return proto.CompactTextString(m) }
func (*SyscallEvent) ProtoMessage()               {}
func (*SyscallEvent) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{5} }

func (m *SyscallEvent) GetType() SyscallEventType {
	if m != nil {
		return m.Type
	}
	return SyscallEventType_SYSCALL_EVENT_TYPE_UNKNOWN
}

func (m *SyscallEvent) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SyscallEvent) GetArg0() uint64 {
	if m != nil {
		return m.Arg0
	}
	return 0
}

func (m *SyscallEvent) GetArg1() uint64 {
	if m != nil {
		return m.Arg1
	}
	return 0
}

func (m *SyscallEvent) GetArg2() uint64 {
	if m != nil {
		return m.Arg2
	}
	return 0
}

func (m *SyscallEvent) GetArg3() uint64 {
	if m != nil {
		return m.Arg3
	}
	return 0
}

func (m *SyscallEvent) GetArg4() uint64 {
	if m != nil {
		return m.Arg4
	}
	return 0
}

func (m *SyscallEvent) GetArg5() uint64 {
	if m != nil {
		return m.Arg5
	}
	return 0
}

func (m *SyscallEvent) GetRet() int64 {
	if m != nil {
		return m.Ret
	}
	return 0
}

type FileEvent struct {
	Type      FileEventType `protobuf:"varint,1,opt,name=type,enum=capsule8.v0.FileEventType" json:"type,omitempty"`
	Filename  string        `protobuf:"bytes,10,opt,name=filename" json:"filename,omitempty"`
	OpenFlags int32         `protobuf:"zigzag32,11,opt,name=open_flags,json=openFlags" json:"open_flags,omitempty"`
	OpenMode  int32         `protobuf:"zigzag32,12,opt,name=open_mode,json=openMode" json:"open_mode,omitempty"`
}

func (m *FileEvent) Reset()                    { *m = FileEvent{} }
func (m *FileEvent) String() string            { return proto.CompactTextString(m) }
func (*FileEvent) ProtoMessage()               {}
func (*FileEvent) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{6} }

func (m *FileEvent) GetType() FileEventType {
	if m != nil {
		return m.Type
	}
	return FileEventType_FILE_EVENT_TYPE_UNKNOWN
}

func (m *FileEvent) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *FileEvent) GetOpenFlags() int32 {
	if m != nil {
		return m.OpenFlags
	}
	return 0
}

func (m *FileEvent) GetOpenMode() int32 {
	if m != nil {
		return m.OpenMode
	}
	return 0
}

func init() {
	proto.RegisterType((*Event)(nil), "capsule8.v0.Event")
	proto.RegisterType((*ChargenEvent)(nil), "capsule8.v0.ChargenEvent")
	proto.RegisterType((*TickerEvent)(nil), "capsule8.v0.TickerEvent")
	proto.RegisterType((*ContainerEvent)(nil), "capsule8.v0.ContainerEvent")
	proto.RegisterType((*ProcessEvent)(nil), "capsule8.v0.ProcessEvent")
	proto.RegisterType((*SyscallEvent)(nil), "capsule8.v0.SyscallEvent")
	proto.RegisterType((*FileEvent)(nil), "capsule8.v0.FileEvent")
	proto.RegisterEnum("capsule8.v0.ContainerEventType", ContainerEventType_name, ContainerEventType_value)
	proto.RegisterEnum("capsule8.v0.ProcessEventType", ProcessEventType_name, ProcessEventType_value)
	proto.RegisterEnum("capsule8.v0.SyscallEventType", SyscallEventType_name, SyscallEventType_value)
	proto.RegisterEnum("capsule8.v0.FileEventType", FileEventType_name, FileEventType_value)
}

func init() { proto.RegisterFile("event.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 926 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x56, 0xc1, 0x6e, 0xdb, 0x46,
	0x10, 0x0d, 0x25, 0xda, 0xb2, 0x46, 0xb2, 0x43, 0x6f, 0x95, 0x74, 0x63, 0x25, 0xb6, 0xaa, 0xf6,
	0x20, 0x08, 0x85, 0x60, 0xcb, 0x36, 0x50, 0xa0, 0x27, 0x87, 0xa6, 0x5a, 0x22, 0x0e, 0x25, 0xac,
	0x94, 0x36, 0xee, 0x85, 0x60, 0xc8, 0xb5, 0x4d, 0x44, 0x24, 0x55, 0x92, 0x36, 0xec, 0x4f, 0xe8,
	0xb9, 0x1f, 0xd4, 0x0f, 0x28, 0xd0, 0x5b, 0x4f, 0xfd, 0x99, 0x62, 0x87, 0x2b, 0x96, 0x72, 0x29,
	0xdf, 0xb8, 0xef, 0xcd, 0xdb, 0x9d, 0x99, 0x7d, 0xb3, 0x20, 0x34, 0xf8, 0x1d, 0x0f, 0xd3, 0xc1,
	0x22, 0x8e, 0xd2, 0x88, 0x34, 0x5c, 0x67, 0x91, 0xdc, 0xce, 0xf9, 0x77, 0x83, 0xbb, 0xc3, 0xee,
	0x5f, 0x2a, 0x6c, 0x18, 0x82, 0x24, 0x3b, 0x50, 0xf1, 0x3d, 0xaa, 0x74, 0x94, 0x5e, 0x9d, 0x55,
	0x7c, 0x8f, 0xbc, 0x01, 0x58, 0xc4, 0x91, 0xcb, 0x93, 0xc4, 0xf6, 0x3d, 0x5a, 0x41, 0xbc, 0x2e,
	0x11, 0xd3, 0x23, 0x07, 0xd0, 0x58, 0xd2, 0x0b, 0xdf, 0xa3, 0xd5, 0x8e, 0xd2, 0xdb, 0x65, 0x4b,
	0xc5, 0xc4, 0xf7, 0xc8, 0x57, 0xd0, 0x74, 0xa3, 0x30, 0x75, 0xfc, 0x90, 0xc7, 0x62, 0x07, 0x15,
	0x77, 0x68, 0xe4, 0x98, 0xe9, 0x91, 0x36, 0xd4, 0x13, 0x1e, 0x26, 0x11, 0xf2, 0x1b, 0xc8, 0x6f,
	0x65, 0x80, 0xe9, 0x91, 0x13, 0x78, 0x29, 0xc9, 0x84, 0xff, 0x7a, 0xcb, 0x43, 0x97, 0xdb, 0xe1,
	0x6d, 0xf0, 0x89, 0xc7, 0x74, 0xb3, 0xa3, 0xf4, 0x54, 0xd6, 0xca, 0xd8, 0xa9, 0x24, 0x2d, 0xe4,
	0xc8, 0x10, 0x5e, 0x48, 0x55, 0x10, 0x85, 0x51, 0xea, 0x07, 0xdc, 0x0e, 0x9d, 0x30, 0x4a, 0x68,
	0xad, 0xa3, 0xf4, 0xaa, 0xec, 0x8b, 0x8c, 0x7c, 0x2f, 0x39, 0x4b, 0x50, 0xe4, 0x14, 0x6a, 0xc9,
	0x43, 0xe2, 0x3a, 0xf3, 0x39, 0x85, 0x8e, 0xd2, 0x6b, 0x0c, 0x5f, 0x0d, 0x0a, 0x2d, 0x1a, 0x4c,
	0x33, 0x0e, 0xbb, 0xf4, 0xe3, 0x33, 0xb6, 0x8c, 0x15, 0x32, 0x59, 0x2e, 0x6d, 0x94, 0xc8, 0x26,
	0x19, 0x97, 0xcb, 0x64, 0x2c, 0xf9, 0x16, 0xd4, 0x2b, 0x7f, 0xce, 0x69, 0x13, 0x35, 0x2f, 0x57,
	0x34, 0x23, 0x7f, 0xce, 0x97, 0x02, 0x8c, 0x22, 0xdf, 0x43, 0x3d, 0xef, 0x18, 0x6d, 0xa1, 0xa4,
	0xbd, 0x22, 0xd1, 0x97, 0xec, 0x52, 0xf7, 0x5f, 0xbc, 0xc8, 0xd0, 0xbd, 0x71, 0xe2, 0x6b, 0x1e,
	0x52, 0xaf, 0x24, 0x43, 0x3d, 0xe3, 0xf2, 0x0c, 0x65, 0x2c, 0x19, 0xc2, 0x66, 0xea, 0xbb, 0x9f,
	0x79, 0x4c, 0x39, 0xaa, 0xe8, 0x8a, 0x6a, 0x86, 0xd4, 0x52, 0x24, 0x23, 0xdf, 0xd6, 0x60, 0x03,
	0x3d, 0xd6, 0x3d, 0x87, 0x66, 0x71, 0x5f, 0xd2, 0x82, 0x0d, 0x3f, 0xf4, 0xf8, 0x3d, 0x3a, 0x4b,
	0x65, 0xd9, 0x82, 0xec, 0x03, 0x88, 0xd3, 0x1c, 0x37, 0xe5, 0x71, 0x22, 0xcd, 0x55, 0x40, 0xba,
	0x26, 0x34, 0x0a, 0xe7, 0x10, 0x0a, 0xb5, 0x84, 0xbb, 0x51, 0xe8, 0x25, 0xb8, 0x4d, 0x95, 0x2d,
	0x97, 0xa4, 0x03, 0x0d, 0xbc, 0x5f, 0xc9, 0x56, 0x90, 0x2d, 0x42, 0xdd, 0x3f, 0x15, 0xd8, 0x59,
	0x6d, 0x12, 0x39, 0x06, 0x35, 0x7d, 0x58, 0x70, 0xdc, 0x6b, 0x67, 0x78, 0xf0, 0x44, 0x3f, 0x67,
	0x0f, 0x0b, 0xce, 0x30, 0x98, 0x10, 0x50, 0x43, 0x27, 0xe0, 0x32, 0x59, 0xfc, 0x26, 0xaf, 0x60,
	0xcb, 0x0f, 0x9c, 0x6b, 0x2e, 0xfc, 0x0b, 0x88, 0xd7, 0x70, 0x6d, 0xe2, 0xf8, 0x64, 0x14, 0x8a,
	0x1a, 0xd9, 0xf8, 0x20, 0x62, 0x49, 0xe5, 0x4d, 0x94, 0xa4, 0x38, 0x3b, 0x2d, 0x9c, 0x9d, 0x9a,
	0x58, 0x8b, 0xc1, 0x69, 0x43, 0x9d, 0xdf, 0xfb, 0xa9, 0xed, 0x46, 0x1e, 0xa7, 0xfb, 0xc8, 0x6d,
	0x09, 0x40, 0x8f, 0x3c, 0xde, 0xfd, 0x5b, 0x81, 0x66, 0xd1, 0x59, 0xe4, 0x68, 0xa5, 0x96, 0x37,
	0x6b, 0x2d, 0x58, 0xa8, 0xe4, 0x1b, 0xd8, 0xb9, 0x8a, 0xe2, 0xcf, 0xb6, 0x7b, 0xe3, 0xcf, 0x3d,
	0xcc, 0x00, 0xf0, 0x94, 0xa6, 0x40, 0x75, 0x01, 0x8a, 0x34, 0xbe, 0x86, 0x6d, 0x7e, 0xcf, 0x5d,
	0x5b, 0xd8, 0x10, 0x6b, 0x68, 0x61, 0x0d, 0x4d, 0x01, 0x8e, 0x24, 0x46, 0xfa, 0xb0, 0x8b, 0x41,
	0x6e, 0x14, 0x04, 0x4e, 0xe8, 0xd9, 0x73, 0x3f, 0xe4, 0xf4, 0x45, 0xa7, 0xda, 0xab, 0xb3, 0xe7,
	0x82, 0xd0, 0x33, 0xfc, 0xc2, 0x0f, 0xf9, 0xd3, 0x75, 0xfd, 0xa3, 0x40, 0xb3, 0x38, 0x68, 0x4f,
	0xd6, 0x55, 0x0c, 0x2c, 0xd4, 0x95, 0xbd, 0x60, 0x99, 0x05, 0xc4, 0x0b, 0x46, 0x40, 0x75, 0xe2,
	0xeb, 0x43, 0xac, 0x4e, 0x65, 0xf8, 0x2d, 0xb1, 0x23, 0xbc, 0x90, 0x0c, 0x3b, 0x92, 0xd8, 0x10,
	0x27, 0x32, 0xc3, 0x86, 0x12, 0x3b, 0xa6, 0xdb, 0x39, 0x76, 0x2c, 0xb1, 0x13, 0xba, 0x93, 0x63,
	0x27, 0x12, 0x3b, 0xa5, 0xcf, 0x73, 0xec, 0x94, 0x68, 0x50, 0x8d, 0x79, 0x8a, 0xfd, 0xaa, 0x32,
	0xf1, 0xd9, 0xfd, 0x5d, 0x81, 0x7a, 0x3e, 0xdb, 0x64, 0xb0, 0x52, 0xda, 0x5e, 0xf9, 0x0b, 0x50,
	0xa8, 0x6b, 0x0f, 0xb6, 0xf2, 0x4b, 0xc8, 0x5c, 0x96, 0xaf, 0x85, 0xcd, 0xa2, 0x05, 0x0f, 0xed,
	0xab, 0xb9, 0x73, 0x9d, 0xbd, 0x43, 0xbb, 0xac, 0x2e, 0x90, 0x91, 0x00, 0x44, 0xcf, 0x91, 0x0e,
	0x44, 0xcf, 0x9b, 0x59, 0xcf, 0x05, 0xf0, 0x3e, 0xf2, 0x78, 0xff, 0x0f, 0x05, 0xc8, 0xff, 0xed,
	0x4e, 0x3a, 0xf0, 0x5a, 0x1f, 0x5b, 0xb3, 0x33, 0xd3, 0x32, 0x98, 0x6d, 0xfc, 0x64, 0x58, 0x33,
	0x7b, 0x76, 0x39, 0x31, 0xec, 0x0f, 0xd6, 0x3b, 0x6b, 0xfc, 0xb3, 0xa5, 0x3d, 0x5b, 0x1b, 0xa1,
	0x33, 0xe3, 0x6c, 0x66, 0x9c, 0x6b, 0xca, 0xda, 0x08, 0xf6, 0xc1, 0xb2, 0x4c, 0xeb, 0x07, 0xad,
	0x42, 0x0e, 0xa0, 0x5d, 0x1a, 0x61, 0x7c, 0x34, 0xc5, 0x16, 0x55, 0xd2, 0x85, 0xfd, 0xd2, 0x80,
	0x73, 0x63, 0x3a, 0x63, 0xe3, 0x4b, 0xe3, 0x5c, 0x53, 0xfb, 0xbf, 0x29, 0xa0, 0x3d, 0x36, 0x39,
	0xd9, 0x87, 0xbd, 0x09, 0x1b, 0xeb, 0xc6, 0x74, 0x5a, 0x9e, 0x7d, 0x1b, 0xbe, 0x2c, 0xe1, 0x47,
	0x63, 0xf6, 0x4e, 0x53, 0xd6, 0x90, 0xc6, 0x47, 0x43, 0xd7, 0x2a, 0x6b, 0x49, 0x73, 0xa6, 0x55,
	0xfb, 0x01, 0x68, 0x8f, 0x7d, 0x29, 0x52, 0x99, 0x5e, 0x4e, 0xf5, 0xb3, 0x8b, 0x8b, 0xf2, 0x54,
	0x5e, 0x03, 0x2d, 0xe1, 0x0d, 0x6b, 0x66, 0xb0, 0x2c, 0x97, 0x32, 0x56, 0x1c, 0x57, 0xe9, 0x8f,
	0x60, 0x7b, 0xc5, 0x2b, 0x22, 0x7a, 0x64, 0x5e, 0x18, 0xe5, 0x07, 0x51, 0x68, 0x3d, 0x26, 0xc7,
	0x13, 0xc3, 0xd2, 0x94, 0xb7, 0xea, 0x2f, 0x95, 0xbb, 0xc3, 0x4f, 0x9b, 0xf8, 0x6b, 0x70, 0xfc,
	0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf1, 0xa1, 0x42, 0x61, 0x29, 0x08, 0x00, 0x00,
}