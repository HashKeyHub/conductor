// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model/eventhandler.proto

package model // import "github.com/netflix/conductor/client/gogrpc/conductor/model"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/golang/protobuf/ptypes/any"
import _struct "github.com/golang/protobuf/ptypes/struct"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EventHandler_Action_Type int32

const (
	EventHandler_Action_START_WORKFLOW EventHandler_Action_Type = 0
	EventHandler_Action_COMPLETE_TASK  EventHandler_Action_Type = 1
	EventHandler_Action_FAIL_TASK      EventHandler_Action_Type = 2
)

var EventHandler_Action_Type_name = map[int32]string{
	0: "START_WORKFLOW",
	1: "COMPLETE_TASK",
	2: "FAIL_TASK",
}
var EventHandler_Action_Type_value = map[string]int32{
	"START_WORKFLOW": 0,
	"COMPLETE_TASK":  1,
	"FAIL_TASK":      2,
}

func (x EventHandler_Action_Type) String() string {
	return proto.EnumName(EventHandler_Action_Type_name, int32(x))
}
func (EventHandler_Action_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_eventhandler_b0d6776f4643d25e, []int{0, 2, 0}
}

type EventHandler struct {
	Name                 string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Event                string                 `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
	Condition            string                 `protobuf:"bytes,3,opt,name=condition,proto3" json:"condition,omitempty"`
	Actions              []*EventHandler_Action `protobuf:"bytes,4,rep,name=actions,proto3" json:"actions,omitempty"`
	Active               bool                   `protobuf:"varint,5,opt,name=active,proto3" json:"active,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *EventHandler) Reset()         { *m = EventHandler{} }
func (m *EventHandler) String() string { return proto.CompactTextString(m) }
func (*EventHandler) ProtoMessage()    {}
func (*EventHandler) Descriptor() ([]byte, []int) {
	return fileDescriptor_eventhandler_b0d6776f4643d25e, []int{0}
}
func (m *EventHandler) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventHandler.Unmarshal(m, b)
}
func (m *EventHandler) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventHandler.Marshal(b, m, deterministic)
}
func (dst *EventHandler) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventHandler.Merge(dst, src)
}
func (m *EventHandler) XXX_Size() int {
	return xxx_messageInfo_EventHandler.Size(m)
}
func (m *EventHandler) XXX_DiscardUnknown() {
	xxx_messageInfo_EventHandler.DiscardUnknown(m)
}

var xxx_messageInfo_EventHandler proto.InternalMessageInfo

func (m *EventHandler) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EventHandler) GetEvent() string {
	if m != nil {
		return m.Event
	}
	return ""
}

func (m *EventHandler) GetCondition() string {
	if m != nil {
		return m.Condition
	}
	return ""
}

func (m *EventHandler) GetActions() []*EventHandler_Action {
	if m != nil {
		return m.Actions
	}
	return nil
}

func (m *EventHandler) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

type EventHandler_StartWorkflow struct {
	Name                 string                    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version              int32                     `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	CorrelationId        string                    `protobuf:"bytes,3,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	Input                map[string]*_struct.Value `protobuf:"bytes,4,rep,name=input,proto3" json:"input,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	InputMessage         *any.Any                  `protobuf:"bytes,5,opt,name=input_message,json=inputMessage,proto3" json:"input_message,omitempty"`
	TaskToDomain         map[string]string         `protobuf:"bytes,6,rep,name=task_to_domain,json=taskToDomain,proto3" json:"task_to_domain,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *EventHandler_StartWorkflow) Reset()         { *m = EventHandler_StartWorkflow{} }
func (m *EventHandler_StartWorkflow) String() string { return proto.CompactTextString(m) }
func (*EventHandler_StartWorkflow) ProtoMessage()    {}
func (*EventHandler_StartWorkflow) Descriptor() ([]byte, []int) {
	return fileDescriptor_eventhandler_b0d6776f4643d25e, []int{0, 0}
}
func (m *EventHandler_StartWorkflow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventHandler_StartWorkflow.Unmarshal(m, b)
}
func (m *EventHandler_StartWorkflow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventHandler_StartWorkflow.Marshal(b, m, deterministic)
}
func (dst *EventHandler_StartWorkflow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventHandler_StartWorkflow.Merge(dst, src)
}
func (m *EventHandler_StartWorkflow) XXX_Size() int {
	return xxx_messageInfo_EventHandler_StartWorkflow.Size(m)
}
func (m *EventHandler_StartWorkflow) XXX_DiscardUnknown() {
	xxx_messageInfo_EventHandler_StartWorkflow.DiscardUnknown(m)
}

var xxx_messageInfo_EventHandler_StartWorkflow proto.InternalMessageInfo

func (m *EventHandler_StartWorkflow) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EventHandler_StartWorkflow) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *EventHandler_StartWorkflow) GetCorrelationId() string {
	if m != nil {
		return m.CorrelationId
	}
	return ""
}

func (m *EventHandler_StartWorkflow) GetInput() map[string]*_struct.Value {
	if m != nil {
		return m.Input
	}
	return nil
}

func (m *EventHandler_StartWorkflow) GetInputMessage() *any.Any {
	if m != nil {
		return m.InputMessage
	}
	return nil
}

func (m *EventHandler_StartWorkflow) GetTaskToDomain() map[string]string {
	if m != nil {
		return m.TaskToDomain
	}
	return nil
}

type EventHandler_TaskDetails struct {
	WorkflowId           string                    `protobuf:"bytes,1,opt,name=workflow_id,json=workflowId,proto3" json:"workflow_id,omitempty"`
	TaskRefName          string                    `protobuf:"bytes,2,opt,name=task_ref_name,json=taskRefName,proto3" json:"task_ref_name,omitempty"`
	Output               map[string]*_struct.Value `protobuf:"bytes,3,rep,name=output,proto3" json:"output,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	OutputMessage        *any.Any                  `protobuf:"bytes,4,opt,name=output_message,json=outputMessage,proto3" json:"output_message,omitempty"`
	TaskId               string                    `protobuf:"bytes,5,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *EventHandler_TaskDetails) Reset()         { *m = EventHandler_TaskDetails{} }
func (m *EventHandler_TaskDetails) String() string { return proto.CompactTextString(m) }
func (*EventHandler_TaskDetails) ProtoMessage()    {}
func (*EventHandler_TaskDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_eventhandler_b0d6776f4643d25e, []int{0, 1}
}
func (m *EventHandler_TaskDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventHandler_TaskDetails.Unmarshal(m, b)
}
func (m *EventHandler_TaskDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventHandler_TaskDetails.Marshal(b, m, deterministic)
}
func (dst *EventHandler_TaskDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventHandler_TaskDetails.Merge(dst, src)
}
func (m *EventHandler_TaskDetails) XXX_Size() int {
	return xxx_messageInfo_EventHandler_TaskDetails.Size(m)
}
func (m *EventHandler_TaskDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_EventHandler_TaskDetails.DiscardUnknown(m)
}

var xxx_messageInfo_EventHandler_TaskDetails proto.InternalMessageInfo

func (m *EventHandler_TaskDetails) GetWorkflowId() string {
	if m != nil {
		return m.WorkflowId
	}
	return ""
}

func (m *EventHandler_TaskDetails) GetTaskRefName() string {
	if m != nil {
		return m.TaskRefName
	}
	return ""
}

func (m *EventHandler_TaskDetails) GetOutput() map[string]*_struct.Value {
	if m != nil {
		return m.Output
	}
	return nil
}

func (m *EventHandler_TaskDetails) GetOutputMessage() *any.Any {
	if m != nil {
		return m.OutputMessage
	}
	return nil
}

func (m *EventHandler_TaskDetails) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

type EventHandler_Action struct {
	Action               EventHandler_Action_Type    `protobuf:"varint,1,opt,name=action,proto3,enum=conductor.proto.EventHandler_Action_Type" json:"action,omitempty"`
	StartWorkflow        *EventHandler_StartWorkflow `protobuf:"bytes,2,opt,name=start_workflow,json=startWorkflow,proto3" json:"start_workflow,omitempty"`
	CompleteTask         *EventHandler_TaskDetails   `protobuf:"bytes,3,opt,name=complete_task,json=completeTask,proto3" json:"complete_task,omitempty"`
	FailTask             *EventHandler_TaskDetails   `protobuf:"bytes,4,opt,name=fail_task,json=failTask,proto3" json:"fail_task,omitempty"`
	ExpandInlineJson     bool                        `protobuf:"varint,5,opt,name=expand_inline_json,json=expandInlineJson,proto3" json:"expand_inline_json,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *EventHandler_Action) Reset()         { *m = EventHandler_Action{} }
func (m *EventHandler_Action) String() string { return proto.CompactTextString(m) }
func (*EventHandler_Action) ProtoMessage()    {}
func (*EventHandler_Action) Descriptor() ([]byte, []int) {
	return fileDescriptor_eventhandler_b0d6776f4643d25e, []int{0, 2}
}
func (m *EventHandler_Action) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventHandler_Action.Unmarshal(m, b)
}
func (m *EventHandler_Action) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventHandler_Action.Marshal(b, m, deterministic)
}
func (dst *EventHandler_Action) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventHandler_Action.Merge(dst, src)
}
func (m *EventHandler_Action) XXX_Size() int {
	return xxx_messageInfo_EventHandler_Action.Size(m)
}
func (m *EventHandler_Action) XXX_DiscardUnknown() {
	xxx_messageInfo_EventHandler_Action.DiscardUnknown(m)
}

var xxx_messageInfo_EventHandler_Action proto.InternalMessageInfo

func (m *EventHandler_Action) GetAction() EventHandler_Action_Type {
	if m != nil {
		return m.Action
	}
	return EventHandler_Action_START_WORKFLOW
}

func (m *EventHandler_Action) GetStartWorkflow() *EventHandler_StartWorkflow {
	if m != nil {
		return m.StartWorkflow
	}
	return nil
}

func (m *EventHandler_Action) GetCompleteTask() *EventHandler_TaskDetails {
	if m != nil {
		return m.CompleteTask
	}
	return nil
}

func (m *EventHandler_Action) GetFailTask() *EventHandler_TaskDetails {
	if m != nil {
		return m.FailTask
	}
	return nil
}

func (m *EventHandler_Action) GetExpandInlineJson() bool {
	if m != nil {
		return m.ExpandInlineJson
	}
	return false
}

func init() {
	proto.RegisterType((*EventHandler)(nil), "conductor.proto.EventHandler")
	proto.RegisterType((*EventHandler_StartWorkflow)(nil), "conductor.proto.EventHandler.StartWorkflow")
	proto.RegisterMapType((map[string]*_struct.Value)(nil), "conductor.proto.EventHandler.StartWorkflow.InputEntry")
	proto.RegisterMapType((map[string]string)(nil), "conductor.proto.EventHandler.StartWorkflow.TaskToDomainEntry")
	proto.RegisterType((*EventHandler_TaskDetails)(nil), "conductor.proto.EventHandler.TaskDetails")
	proto.RegisterMapType((map[string]*_struct.Value)(nil), "conductor.proto.EventHandler.TaskDetails.OutputEntry")
	proto.RegisterType((*EventHandler_Action)(nil), "conductor.proto.EventHandler.Action")
	proto.RegisterEnum("conductor.proto.EventHandler_Action_Type", EventHandler_Action_Type_name, EventHandler_Action_Type_value)
}

func init() {
	proto.RegisterFile("model/eventhandler.proto", fileDescriptor_eventhandler_b0d6776f4643d25e)
}

var fileDescriptor_eventhandler_b0d6776f4643d25e = []byte{
	// 727 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0xfb, 0x4e, 0x13, 0x4f,
	0x14, 0xc7, 0x7f, 0xbd, 0x42, 0xcf, 0x76, 0xfb, 0x2b, 0x13, 0x82, 0x6b, 0x25, 0x91, 0x10, 0x4d,
	0x30, 0x92, 0x6d, 0x52, 0xa3, 0x51, 0xbc, 0xa5, 0x48, 0x89, 0x95, 0x42, 0x71, 0x68, 0x24, 0xf1,
	0x9f, 0xcd, 0x76, 0x77, 0x5a, 0xc6, 0x6e, 0x67, 0x9a, 0xdd, 0x69, 0xa1, 0x8f, 0xe0, 0x73, 0xf8,
	0x36, 0xbe, 0x87, 0xef, 0x61, 0x66, 0x66, 0x17, 0x96, 0x4b, 0x50, 0x12, 0xff, 0x3b, 0x97, 0x39,
	0xdf, 0xf3, 0xd9, 0x73, 0x66, 0x5a, 0xb0, 0xc6, 0xdc, 0x27, 0x41, 0x9d, 0xcc, 0x08, 0x13, 0x27,
	0x2e, 0xf3, 0x03, 0x12, 0xda, 0x93, 0x90, 0x0b, 0x8e, 0xfe, 0xf7, 0x38, 0xf3, 0xa7, 0x9e, 0xe0,
	0x71, 0xa0, 0xb6, 0x3a, 0xe4, 0x7c, 0x18, 0x90, 0xba, 0xf2, 0xfa, 0xd3, 0x41, 0x3d, 0x12, 0xe1,
	0xd4, 0x13, 0x71, 0xf6, 0xfe, 0xd5, 0xac, 0xcb, 0xe6, 0x3a, 0xb5, 0xfe, 0xdd, 0x80, 0x72, 0x4b,
	0x36, 0xf8, 0xa8, 0x1b, 0x20, 0x04, 0x79, 0xe6, 0x8e, 0x89, 0x95, 0x59, 0xcb, 0x6c, 0x94, 0xb0,
	0xb2, 0xd1, 0x32, 0x14, 0x14, 0x84, 0x95, 0x55, 0x41, 0xed, 0xa0, 0x55, 0x28, 0x49, 0x0c, 0x2a,
	0x28, 0x67, 0x56, 0x4e, 0x65, 0x2e, 0x02, 0xe8, 0x1d, 0x2c, 0xb8, 0x9e, 0xb4, 0x22, 0x2b, 0xbf,
	0x96, 0xdb, 0x30, 0x1a, 0x8f, 0xec, 0x2b, 0xd0, 0x76, 0xba, 0xaf, 0xdd, 0x54, 0x87, 0x71, 0x52,
	0x84, 0x56, 0xa0, 0x28, 0xcd, 0x19, 0xb1, 0x0a, 0x6b, 0x99, 0x8d, 0x45, 0x1c, 0x7b, 0xb5, 0x5f,
	0x39, 0x30, 0x8f, 0x84, 0x1b, 0x8a, 0x63, 0x1e, 0x8e, 0x06, 0x01, 0x3f, 0xbd, 0x91, 0xd8, 0x82,
	0x85, 0x19, 0x09, 0x23, 0x49, 0x26, 0x99, 0x0b, 0x38, 0x71, 0xd1, 0x63, 0xa8, 0x78, 0x3c, 0x0c,
	0x49, 0xe0, 0xca, 0x3e, 0x0e, 0xf5, 0x63, 0x74, 0x33, 0x15, 0x6d, 0xfb, 0xa8, 0x03, 0x05, 0xca,
	0x26, 0x53, 0x11, 0xc3, 0xbf, 0xb8, 0x1d, 0xfe, 0x12, 0x90, 0xdd, 0x96, 0x85, 0x2d, 0x26, 0xc2,
	0x39, 0xd6, 0x22, 0xe8, 0x15, 0x98, 0xca, 0x70, 0xc6, 0x24, 0x8a, 0xdc, 0xa1, 0xfe, 0x26, 0xa3,
	0xb1, 0x6c, 0xeb, 0xc5, 0xd8, 0xc9, 0x62, 0xec, 0x26, 0x9b, 0xe3, 0xb2, 0x3a, 0xba, 0xaf, 0x4f,
	0x22, 0x0f, 0x2a, 0xc2, 0x8d, 0x46, 0x8e, 0xe0, 0x8e, 0xcf, 0xc7, 0x2e, 0x65, 0x56, 0x51, 0x11,
	0xbd, 0xbd, 0x0b, 0x51, 0xcf, 0x8d, 0x46, 0x3d, 0xbe, 0xa3, 0xea, 0x35, 0x58, 0x59, 0xa4, 0x42,
	0xb5, 0x43, 0x80, 0x0b, 0x68, 0x54, 0x85, 0xdc, 0x88, 0xcc, 0xe3, 0x79, 0x4a, 0x13, 0x6d, 0x42,
	0x61, 0xe6, 0x06, 0x53, 0xa2, 0x86, 0x69, 0x34, 0x56, 0xae, 0x71, 0x7f, 0x91, 0x59, 0xac, 0x0f,
	0x6d, 0x65, 0x5f, 0x66, 0x6a, 0xef, 0x61, 0xe9, 0x5a, 0xd3, 0x1b, 0x84, 0x97, 0xd3, 0xc2, 0xa5,
	0xb4, 0xc0, 0xcf, 0x2c, 0x18, 0x52, 0x61, 0x87, 0x08, 0x97, 0x06, 0x11, 0x7a, 0x08, 0xc6, 0x69,
	0xfc, 0x39, 0x72, 0x69, 0x5a, 0x03, 0x92, 0x50, 0xdb, 0x47, 0xeb, 0x60, 0xaa, 0x41, 0x85, 0x64,
	0xe0, 0xa8, 0xfb, 0xa0, 0x25, 0x0d, 0x19, 0xc4, 0x64, 0x70, 0x20, 0xaf, 0xc5, 0x3e, 0x14, 0xf9,
	0x54, 0xc8, 0xb5, 0xe6, 0xd4, 0x10, 0x9f, 0xdf, 0x3e, 0xc4, 0x54, 0x7f, 0xbb, 0xab, 0xea, 0xf4,
	0xf0, 0x62, 0x11, 0xf4, 0x1a, 0x2a, 0xda, 0x3a, 0xdf, 0x6b, 0xfe, 0x96, 0xbd, 0x9a, 0xfa, 0x6c,
	0xb2, 0xd8, 0x7b, 0xb0, 0xa0, 0x78, 0xa9, 0xaf, 0x6e, 0x43, 0x09, 0x17, 0xa5, 0xdb, 0xf6, 0x6b,
	0x9f, 0xc1, 0x48, 0x35, 0xfb, 0x27, 0xdb, 0xf8, 0x91, 0x83, 0xa2, 0x7e, 0x60, 0xa8, 0xa9, 0xdf,
	0x15, 0x67, 0x4a, 0xb1, 0xd2, 0x78, 0xf2, 0x37, 0xcf, 0xd2, 0xee, 0xcd, 0x27, 0x04, 0xc7, 0x85,
	0x08, 0x43, 0x25, 0x92, 0xd7, 0xcb, 0x49, 0xa6, 0x1f, 0x83, 0x3c, 0xbd, 0xc3, 0x95, 0xc4, 0x66,
	0x74, 0xe9, 0x11, 0x1f, 0x80, 0xe9, 0xf1, 0xf1, 0x24, 0x20, 0x82, 0x38, 0x72, 0x0e, 0xea, 0x55,
	0x1a, 0x7f, 0xa2, 0x4b, 0x2d, 0x08, 0x97, 0x93, 0x7a, 0x19, 0x44, 0xbb, 0x50, 0x1a, 0xb8, 0x34,
	0xd0, 0x5a, 0xf9, 0xbb, 0x6a, 0x2d, 0xca, 0x5a, 0xa5, 0xb3, 0x09, 0x88, 0x9c, 0x4d, 0x5c, 0xe6,
	0x3b, 0x94, 0x05, 0x94, 0x11, 0xe7, 0x5b, 0xc4, 0x59, 0xfc, 0x93, 0x54, 0xd5, 0x99, 0xb6, 0x4a,
	0x7c, 0x8a, 0x38, 0x5b, 0x7f, 0x03, 0x79, 0x39, 0x29, 0x84, 0xa0, 0x72, 0xd4, 0x6b, 0xe2, 0x9e,
	0x73, 0xdc, 0xc5, 0x7b, 0xbb, 0x9d, 0xee, 0x71, 0xf5, 0x3f, 0xb4, 0x04, 0xe6, 0x87, 0xee, 0xfe,
	0x61, 0xa7, 0xd5, 0x6b, 0x39, 0xbd, 0xe6, 0xd1, 0x5e, 0x35, 0x83, 0x4c, 0x28, 0xed, 0x36, 0xdb,
	0x1d, 0xed, 0x66, 0xb7, 0x29, 0x3c, 0xf0, 0xf8, 0xd8, 0x66, 0x44, 0x0c, 0x02, 0x7a, 0x76, 0x95,
	0x76, 0xbb, 0x92, 0xc6, 0x3d, 0xec, 0x7f, 0xdd, 0x1a, 0x52, 0x71, 0x32, 0xed, 0xdb, 0x1e, 0x1f,
	0xd7, 0xe3, 0x9a, 0xfa, 0x79, 0x4d, 0xdd, 0x0b, 0x28, 0x61, 0xa2, 0x3e, 0xe4, 0xc3, 0x70, 0xe2,
	0xa5, 0xe2, 0xea, 0x3f, 0xa5, 0x5f, 0x54, 0x92, 0xcf, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0xdb,
	0x35, 0x63, 0x90, 0x63, 0x06, 0x00, 0x00,
}
