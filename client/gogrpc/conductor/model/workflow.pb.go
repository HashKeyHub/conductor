// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model/workflow.proto

package model // import "github.com/netflix/conductor/client/gogrpc/conductor/model"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
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

type Workflow_WorkflowStatus int32

const (
	Workflow_RUNNING    Workflow_WorkflowStatus = 0
	Workflow_COMPLETED  Workflow_WorkflowStatus = 1
	Workflow_FAILED     Workflow_WorkflowStatus = 2
	Workflow_TIMED_OUT  Workflow_WorkflowStatus = 3
	Workflow_TERMINATED Workflow_WorkflowStatus = 4
	Workflow_PAUSED     Workflow_WorkflowStatus = 5
)

var Workflow_WorkflowStatus_name = map[int32]string{
	0: "RUNNING",
	1: "COMPLETED",
	2: "FAILED",
	3: "TIMED_OUT",
	4: "TERMINATED",
	5: "PAUSED",
}
var Workflow_WorkflowStatus_value = map[string]int32{
	"RUNNING":    0,
	"COMPLETED":  1,
	"FAILED":     2,
	"TIMED_OUT":  3,
	"TERMINATED": 4,
	"PAUSED":     5,
}

func (x Workflow_WorkflowStatus) String() string {
	return proto.EnumName(Workflow_WorkflowStatus_name, int32(x))
}
func (Workflow_WorkflowStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_workflow_4ef5cdb8c99357a4, []int{0, 0}
}

type Workflow struct {
	Status                           Workflow_WorkflowStatus   `protobuf:"varint,1,opt,name=status,proto3,enum=conductor.proto.Workflow_WorkflowStatus" json:"status,omitempty"`
	EndTime                          int64                     `protobuf:"varint,2,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	WorkflowId                       string                    `protobuf:"bytes,3,opt,name=workflow_id,json=workflowId,proto3" json:"workflow_id,omitempty"`
	ParentWorkflowId                 string                    `protobuf:"bytes,4,opt,name=parent_workflow_id,json=parentWorkflowId,proto3" json:"parent_workflow_id,omitempty"`
	ParentWorkflowTaskId             string                    `protobuf:"bytes,5,opt,name=parent_workflow_task_id,json=parentWorkflowTaskId,proto3" json:"parent_workflow_task_id,omitempty"`
	Tasks                            []*Task                   `protobuf:"bytes,6,rep,name=tasks,proto3" json:"tasks,omitempty"`
	Input                            map[string]*_struct.Value `protobuf:"bytes,8,rep,name=input,proto3" json:"input,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Output                           map[string]*_struct.Value `protobuf:"bytes,9,rep,name=output,proto3" json:"output,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	WorkflowType                     string                    `protobuf:"bytes,10,opt,name=workflow_type,json=workflowType,proto3" json:"workflow_type,omitempty"`
	Version                          int32                     `protobuf:"varint,11,opt,name=version,proto3" json:"version,omitempty"`
	CorrelationId                    string                    `protobuf:"bytes,12,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	ReRunFromWorkflowId              string                    `protobuf:"bytes,13,opt,name=re_run_from_workflow_id,json=reRunFromWorkflowId,proto3" json:"re_run_from_workflow_id,omitempty"`
	ReasonForIncompletion            string                    `protobuf:"bytes,14,opt,name=reason_for_incompletion,json=reasonForIncompletion,proto3" json:"reason_for_incompletion,omitempty"`
	SchemaVersion                    int32                     `protobuf:"varint,15,opt,name=schema_version,json=schemaVersion,proto3" json:"schema_version,omitempty"`
	Event                            string                    `protobuf:"bytes,16,opt,name=event,proto3" json:"event,omitempty"`
	TaskToDomain                     map[string]string         `protobuf:"bytes,17,rep,name=task_to_domain,json=taskToDomain,proto3" json:"task_to_domain,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	FailedReferenceTaskNames         []string                  `protobuf:"bytes,18,rep,name=failed_reference_task_names,json=failedReferenceTaskNames,proto3" json:"failed_reference_task_names,omitempty"`
	WorkflowDefinition               *WorkflowDef              `protobuf:"bytes,19,opt,name=workflow_definition,json=workflowDefinition,proto3" json:"workflow_definition,omitempty"`
	ExternalInputPayloadStoragePath  string                    `protobuf:"bytes,20,opt,name=external_input_payload_storage_path,json=externalInputPayloadStoragePath,proto3" json:"external_input_payload_storage_path,omitempty"`
	ExternalOutputPayloadStoragePath string                    `protobuf:"bytes,21,opt,name=external_output_payload_storage_path,json=externalOutputPayloadStoragePath,proto3" json:"external_output_payload_storage_path,omitempty"`
	Priority                         int32                     `protobuf:"varint,22,opt,name=priority,proto3" json:"priority,omitempty"`
	XXX_NoUnkeyedLiteral             struct{}                  `json:"-"`
	XXX_unrecognized                 []byte                    `json:"-"`
	XXX_sizecache                    int32                     `json:"-"`
}

func (m *Workflow) Reset()         { *m = Workflow{} }
func (m *Workflow) String() string { return proto.CompactTextString(m) }
func (*Workflow) ProtoMessage()    {}
func (*Workflow) Descriptor() ([]byte, []int) {
	return fileDescriptor_workflow_4ef5cdb8c99357a4, []int{0}
}
func (m *Workflow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Workflow.Unmarshal(m, b)
}
func (m *Workflow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Workflow.Marshal(b, m, deterministic)
}
func (dst *Workflow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Workflow.Merge(dst, src)
}
func (m *Workflow) XXX_Size() int {
	return xxx_messageInfo_Workflow.Size(m)
}
func (m *Workflow) XXX_DiscardUnknown() {
	xxx_messageInfo_Workflow.DiscardUnknown(m)
}

var xxx_messageInfo_Workflow proto.InternalMessageInfo

func (m *Workflow) GetStatus() Workflow_WorkflowStatus {
	if m != nil {
		return m.Status
	}
	return Workflow_RUNNING
}

func (m *Workflow) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func (m *Workflow) GetWorkflowId() string {
	if m != nil {
		return m.WorkflowId
	}
	return ""
}

func (m *Workflow) GetParentWorkflowId() string {
	if m != nil {
		return m.ParentWorkflowId
	}
	return ""
}

func (m *Workflow) GetParentWorkflowTaskId() string {
	if m != nil {
		return m.ParentWorkflowTaskId
	}
	return ""
}

func (m *Workflow) GetTasks() []*Task {
	if m != nil {
		return m.Tasks
	}
	return nil
}

func (m *Workflow) GetInput() map[string]*_struct.Value {
	if m != nil {
		return m.Input
	}
	return nil
}

func (m *Workflow) GetOutput() map[string]*_struct.Value {
	if m != nil {
		return m.Output
	}
	return nil
}

func (m *Workflow) GetWorkflowType() string {
	if m != nil {
		return m.WorkflowType
	}
	return ""
}

func (m *Workflow) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Workflow) GetCorrelationId() string {
	if m != nil {
		return m.CorrelationId
	}
	return ""
}

func (m *Workflow) GetReRunFromWorkflowId() string {
	if m != nil {
		return m.ReRunFromWorkflowId
	}
	return ""
}

func (m *Workflow) GetReasonForIncompletion() string {
	if m != nil {
		return m.ReasonForIncompletion
	}
	return ""
}

func (m *Workflow) GetSchemaVersion() int32 {
	if m != nil {
		return m.SchemaVersion
	}
	return 0
}

func (m *Workflow) GetEvent() string {
	if m != nil {
		return m.Event
	}
	return ""
}

func (m *Workflow) GetTaskToDomain() map[string]string {
	if m != nil {
		return m.TaskToDomain
	}
	return nil
}

func (m *Workflow) GetFailedReferenceTaskNames() []string {
	if m != nil {
		return m.FailedReferenceTaskNames
	}
	return nil
}

func (m *Workflow) GetWorkflowDefinition() *WorkflowDef {
	if m != nil {
		return m.WorkflowDefinition
	}
	return nil
}

func (m *Workflow) GetExternalInputPayloadStoragePath() string {
	if m != nil {
		return m.ExternalInputPayloadStoragePath
	}
	return ""
}

func (m *Workflow) GetExternalOutputPayloadStoragePath() string {
	if m != nil {
		return m.ExternalOutputPayloadStoragePath
	}
	return ""
}

func (m *Workflow) GetPriority() int32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func init() {
	proto.RegisterType((*Workflow)(nil), "conductor.proto.Workflow")
	proto.RegisterMapType((map[string]*_struct.Value)(nil), "conductor.proto.Workflow.InputEntry")
	proto.RegisterMapType((map[string]*_struct.Value)(nil), "conductor.proto.Workflow.OutputEntry")
	proto.RegisterMapType((map[string]string)(nil), "conductor.proto.Workflow.TaskToDomainEntry")
	proto.RegisterEnum("conductor.proto.Workflow_WorkflowStatus", Workflow_WorkflowStatus_name, Workflow_WorkflowStatus_value)
}

func init() { proto.RegisterFile("model/workflow.proto", fileDescriptor_workflow_4ef5cdb8c99357a4) }

var fileDescriptor_workflow_4ef5cdb8c99357a4 = []byte{
	// 805 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0x6d, 0x4f, 0xdb, 0x48,
	0x10, 0xc7, 0x2f, 0xa4, 0x09, 0x64, 0x42, 0x52, 0x77, 0x09, 0xc5, 0x47, 0x2b, 0x35, 0xea, 0x83,
	0x14, 0xa9, 0x55, 0x22, 0x71, 0x0f, 0x3a, 0x21, 0x55, 0x77, 0xf4, 0x12, 0x4e, 0x96, 0x20, 0xa4,
	0x26, 0xb4, 0xd2, 0xbd, 0xb1, 0x16, 0x7b, 0x1c, 0x2c, 0xec, 0x5d, 0x6b, 0xbd, 0x86, 0xe6, 0x3b,
	0xdc, 0x87, 0x3e, 0xed, 0xae, 0x1d, 0x0c, 0x25, 0xef, 0xfa, 0xce, 0x3b, 0xf3, 0xfb, 0xff, 0x99,
	0x9d, 0x1d, 0x26, 0xd0, 0x4b, 0x78, 0x80, 0xf1, 0xe8, 0x96, 0x8b, 0xeb, 0x30, 0xe6, 0xb7, 0xc3,
	0x54, 0x70, 0xc9, 0xc9, 0x53, 0x9f, 0xb3, 0x20, 0xf7, 0x25, 0x17, 0x26, 0xb0, 0xbf, 0x77, 0x1f,
	0x0b, 0x30, 0x2c, 0x12, 0x96, 0x49, 0x48, 0x9a, 0x5d, 0x17, 0x91, 0x97, 0x0b, 0xce, 0x17, 0x31,
	0x8e, 0xf4, 0xe9, 0x32, 0x0f, 0x47, 0x99, 0x14, 0xb9, 0x2f, 0x4d, 0xf6, 0xf5, 0x7f, 0x6d, 0xd8,
	0xfa, 0x5a, 0xb8, 0x90, 0xbf, 0xa0, 0x99, 0x49, 0x2a, 0xf3, 0xcc, 0xae, 0xf5, 0x6b, 0x83, 0xee,
	0xc1, 0x60, 0xf8, 0xe0, 0xef, 0x0e, 0x4b, 0x74, 0xf5, 0x71, 0xae, 0x79, 0xb7, 0xd0, 0x91, 0x9f,
	0x61, 0x0b, 0x59, 0xe0, 0xc9, 0x28, 0x41, 0x7b, 0xa3, 0x5f, 0x1b, 0xd4, 0xdd, 0x4d, 0x64, 0xc1,
	0x3c, 0x4a, 0x90, 0xbc, 0x82, 0x76, 0x59, 0xae, 0x17, 0x05, 0x76, 0xbd, 0x5f, 0x1b, 0xb4, 0x5c,
	0x28, 0x43, 0x4e, 0x40, 0x3e, 0x00, 0x49, 0xa9, 0x40, 0x26, 0xbd, 0x2a, 0xf7, 0x44, 0x73, 0x96,
	0xc9, 0x7c, 0xbd, 0xa3, 0x7f, 0x83, 0xbd, 0x87, 0xb4, 0xba, 0xb4, 0x92, 0x34, 0xb4, 0xa4, 0x77,
	0x5f, 0x32, 0xa7, 0xd9, 0xb5, 0x13, 0x90, 0xf7, 0xd0, 0x50, 0x58, 0x66, 0x37, 0xfb, 0xf5, 0x41,
	0xfb, 0x60, 0xf7, 0xbb, 0x1b, 0x2a, 0xce, 0x35, 0x0c, 0x39, 0x84, 0x46, 0xc4, 0xd2, 0x5c, 0xda,
	0x5b, 0x1a, 0x7e, 0xbb, 0xbe, 0x1d, 0x8e, 0xc2, 0x26, 0x4c, 0x8a, 0xa5, 0x6b, 0x24, 0xe4, 0x23,
	0x34, 0x79, 0x2e, 0x95, 0xb8, 0xa5, 0xc5, 0xef, 0xd6, 0x8b, 0xcf, 0x34, 0x67, 0xd4, 0x85, 0x88,
	0xbc, 0x81, 0xce, 0xdd, 0xbd, 0x96, 0x29, 0xda, 0xa0, 0x2f, 0xb5, 0x5d, 0x06, 0xe7, 0xcb, 0x14,
	0x89, 0x0d, 0x9b, 0x37, 0x28, 0xb2, 0x88, 0x33, 0xbb, 0xdd, 0xaf, 0x0d, 0x1a, 0x6e, 0x79, 0x24,
	0xef, 0xa0, 0xeb, 0x73, 0x21, 0x30, 0xa6, 0x32, 0xe2, 0x4c, 0x35, 0x65, 0x5b, 0xeb, 0x3b, 0x95,
	0xa8, 0x13, 0x90, 0x5f, 0x61, 0x4f, 0xa0, 0x27, 0x72, 0xe6, 0x85, 0x82, 0x27, 0xf7, 0xfa, 0xde,
	0xd1, 0xfc, 0x8e, 0x40, 0x37, 0x67, 0xc7, 0x82, 0x27, 0x95, 0xd6, 0xff, 0xae, 0x54, 0x34, 0xe3,
	0xcc, 0x0b, 0xb9, 0xf0, 0x22, 0xe6, 0xf3, 0x24, 0x8d, 0x51, 0x59, 0xda, 0x5d, 0xad, 0xda, 0x35,
	0xe9, 0x63, 0x2e, 0x9c, 0x4a, 0x52, 0x15, 0x95, 0xf9, 0x57, 0x98, 0x50, 0xaf, 0xac, 0xfa, 0xa9,
	0xae, 0xba, 0x63, 0xa2, 0x5f, 0x8a, 0xda, 0x7b, 0xd0, 0xc0, 0x1b, 0x64, 0xd2, 0xb6, 0xb4, 0x99,
	0x39, 0x90, 0xcf, 0xd0, 0xd5, 0xef, 0x2b, 0xb9, 0x17, 0xf0, 0x84, 0x46, 0xcc, 0x7e, 0xa6, 0xfb,
	0xfa, 0x7e, 0x7d, 0x5f, 0xd5, 0x53, 0xce, 0xf9, 0x58, 0xd3, 0xa6, 0xbb, 0xdb, 0xb2, 0x12, 0x22,
	0x1f, 0xe1, 0x45, 0x48, 0xa3, 0x18, 0x03, 0x4f, 0x60, 0x88, 0x02, 0x99, 0x8f, 0x66, 0x86, 0x18,
	0x4d, 0x30, 0xb3, 0x49, 0xbf, 0x3e, 0x68, 0xb9, 0xb6, 0x41, 0xdc, 0x92, 0x50, 0xa6, 0x53, 0x95,
	0x27, 0xa7, 0xb0, 0xb3, 0x6a, 0x58, 0x80, 0x61, 0xc4, 0x22, 0xdd, 0x82, 0x9d, 0x7e, 0x6d, 0xd0,
	0x3e, 0x78, 0xb9, 0xb6, 0xac, 0x31, 0x86, 0x2e, 0xb9, 0xbd, 0x3b, 0x14, 0x3a, 0x72, 0x02, 0x6f,
	0xf0, 0x9b, 0x44, 0xc1, 0x68, 0xec, 0xe9, 0x11, 0xf2, 0x52, 0xba, 0x8c, 0x39, 0x0d, 0xbc, 0x4c,
	0x72, 0x41, 0x17, 0xe8, 0xa5, 0x54, 0x5e, 0xd9, 0x3d, 0xdd, 0x94, 0x57, 0x25, 0xaa, 0x07, 0x6f,
	0x66, 0xc0, 0x73, 0xc3, 0xcd, 0xa8, 0xbc, 0x22, 0x53, 0x78, 0xbb, 0x72, 0x33, 0x23, 0xf5, 0xb8,
	0xdd, 0xae, 0xb6, 0xeb, 0x97, 0xac, 0x19, 0xc5, 0x47, 0xfc, 0xf6, 0x61, 0x2b, 0x15, 0x11, 0x17,
	0x91, 0x5c, 0xda, 0xcf, 0xf5, 0xab, 0xad, 0xce, 0xfb, 0x33, 0x80, 0xbb, 0xf9, 0x27, 0x16, 0xd4,
	0xaf, 0x71, 0xa9, 0x37, 0x48, 0xcb, 0x55, 0x9f, 0xe4, 0x03, 0x34, 0x6e, 0x68, 0x9c, 0x9b, 0x8d,
	0xd0, 0x3e, 0x78, 0x3e, 0x34, 0x1b, 0x69, 0x58, 0x6e, 0xa4, 0xe1, 0x17, 0x95, 0x75, 0x0d, 0x74,
	0xb8, 0xf1, 0x47, 0x6d, 0xff, 0x33, 0xb4, 0x2b, 0xff, 0x14, 0x3f, 0xc4, 0xf2, 0x4f, 0x78, 0xf6,
	0xdd, 0x3c, 0x3c, 0x62, 0xdc, 0xab, 0x1a, 0xb7, 0x2a, 0x06, 0xaf, 0x7d, 0xe8, 0xde, 0x5f, 0x7a,
	0xa4, 0x0d, 0x9b, 0xee, 0xc5, 0x74, 0xea, 0x4c, 0xff, 0xb1, 0x7e, 0x22, 0x1d, 0x68, 0xfd, 0x7d,
	0x76, 0x3a, 0x3b, 0x99, 0xcc, 0x27, 0x63, 0xab, 0x46, 0x00, 0x9a, 0xc7, 0x47, 0xce, 0xc9, 0x64,
	0x6c, 0x6d, 0xa8, 0xd4, 0xdc, 0x39, 0x9d, 0x8c, 0xbd, 0xb3, 0x8b, 0xb9, 0x55, 0x27, 0x5d, 0x80,
	0xf9, 0xc4, 0x3d, 0x75, 0xa6, 0x47, 0x0a, 0x7d, 0xa2, 0xd0, 0xd9, 0xd1, 0xc5, 0xf9, 0x64, 0x6c,
	0x35, 0x3e, 0x21, 0xbc, 0xf0, 0x79, 0x32, 0x64, 0x28, 0xc3, 0x38, 0xfa, 0xf6, 0x70, 0x86, 0x3e,
	0x41, 0x59, 0xc1, 0xec, 0xf2, 0xdf, 0xc3, 0x45, 0x24, 0xaf, 0xf2, 0xcb, 0xa1, 0xcf, 0x93, 0x51,
	0xc1, 0x8f, 0x56, 0xfc, 0xc8, 0x8f, 0x23, 0x64, 0x72, 0xb4, 0xe0, 0x0b, 0x91, 0xfa, 0x95, 0xb8,
	0xfe, 0x79, 0xb8, 0x6c, 0x6a, 0xbb, 0x5f, 0xfe, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x9f, 0x99, 0x18,
	0x17, 0x6e, 0x06, 0x00, 0x00,
}
