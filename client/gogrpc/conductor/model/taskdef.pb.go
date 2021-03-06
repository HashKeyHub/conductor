// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model/taskdef.proto

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

type TaskDef_RetryLogic int32

const (
	TaskDef_FIXED               TaskDef_RetryLogic = 0
	TaskDef_EXPONENTIAL_BACKOFF TaskDef_RetryLogic = 1
)

var TaskDef_RetryLogic_name = map[int32]string{
	0: "FIXED",
	1: "EXPONENTIAL_BACKOFF",
}
var TaskDef_RetryLogic_value = map[string]int32{
	"FIXED":               0,
	"EXPONENTIAL_BACKOFF": 1,
}

func (x TaskDef_RetryLogic) String() string {
	return proto.EnumName(TaskDef_RetryLogic_name, int32(x))
}
func (TaskDef_RetryLogic) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_taskdef_a78be32c045d2b55, []int{0, 0}
}

type TaskDef_TimeoutPolicy int32

const (
	TaskDef_RETRY       TaskDef_TimeoutPolicy = 0
	TaskDef_TIME_OUT_WF TaskDef_TimeoutPolicy = 1
	TaskDef_ALERT_ONLY  TaskDef_TimeoutPolicy = 2
)

var TaskDef_TimeoutPolicy_name = map[int32]string{
	0: "RETRY",
	1: "TIME_OUT_WF",
	2: "ALERT_ONLY",
}
var TaskDef_TimeoutPolicy_value = map[string]int32{
	"RETRY":       0,
	"TIME_OUT_WF": 1,
	"ALERT_ONLY":  2,
}

func (x TaskDef_TimeoutPolicy) String() string {
	return proto.EnumName(TaskDef_TimeoutPolicy_name, int32(x))
}
func (TaskDef_TimeoutPolicy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_taskdef_a78be32c045d2b55, []int{0, 1}
}

type TaskDef struct {
	Name                        string                    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description                 string                    `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	RetryCount                  int32                     `protobuf:"varint,3,opt,name=retry_count,json=retryCount,proto3" json:"retry_count,omitempty"`
	TimeoutSeconds              int64                     `protobuf:"varint,4,opt,name=timeout_seconds,json=timeoutSeconds,proto3" json:"timeout_seconds,omitempty"`
	InputKeys                   []string                  `protobuf:"bytes,5,rep,name=input_keys,json=inputKeys,proto3" json:"input_keys,omitempty"`
	OutputKeys                  []string                  `protobuf:"bytes,6,rep,name=output_keys,json=outputKeys,proto3" json:"output_keys,omitempty"`
	TimeoutPolicy               TaskDef_TimeoutPolicy     `protobuf:"varint,7,opt,name=timeout_policy,json=timeoutPolicy,proto3,enum=conductor.proto.TaskDef_TimeoutPolicy" json:"timeout_policy,omitempty"`
	RetryLogic                  TaskDef_RetryLogic        `protobuf:"varint,8,opt,name=retry_logic,json=retryLogic,proto3,enum=conductor.proto.TaskDef_RetryLogic" json:"retry_logic,omitempty"`
	RetryDelaySeconds           int32                     `protobuf:"varint,9,opt,name=retry_delay_seconds,json=retryDelaySeconds,proto3" json:"retry_delay_seconds,omitempty"`
	ResponseTimeoutSeconds      int64                     `protobuf:"varint,10,opt,name=response_timeout_seconds,json=responseTimeoutSeconds,proto3" json:"response_timeout_seconds,omitempty"`
	ConcurrentExecLimit         int32                     `protobuf:"varint,11,opt,name=concurrent_exec_limit,json=concurrentExecLimit,proto3" json:"concurrent_exec_limit,omitempty"`
	InputTemplate               map[string]*_struct.Value `protobuf:"bytes,12,rep,name=input_template,json=inputTemplate,proto3" json:"input_template,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	RateLimitPerFrequency       int32                     `protobuf:"varint,14,opt,name=rate_limit_per_frequency,json=rateLimitPerFrequency,proto3" json:"rate_limit_per_frequency,omitempty"`
	RateLimitFrequencyInSeconds int32                     `protobuf:"varint,15,opt,name=rate_limit_frequency_in_seconds,json=rateLimitFrequencyInSeconds,proto3" json:"rate_limit_frequency_in_seconds,omitempty"`
	IsolationGroupId            string                    `protobuf:"bytes,16,opt,name=isolation_group_id,json=isolationGroupId,proto3" json:"isolation_group_id,omitempty"`
	ExecutionNameSpace          string                    `protobuf:"bytes,17,opt,name=execution_name_space,json=executionNameSpace,proto3" json:"execution_name_space,omitempty"`
	OwnerEmail                  string                    `protobuf:"bytes,18,opt,name=owner_email,json=ownerEmail,proto3" json:"owner_email,omitempty"`
	PollTimeoutSeconds          int32                     `protobuf:"varint,19,opt,name=poll_timeout_seconds,json=pollTimeoutSeconds,proto3" json:"poll_timeout_seconds,omitempty"`
	XXX_NoUnkeyedLiteral        struct{}                  `json:"-"`
	XXX_unrecognized            []byte                    `json:"-"`
	XXX_sizecache               int32                     `json:"-"`
}

func (m *TaskDef) Reset()         { *m = TaskDef{} }
func (m *TaskDef) String() string { return proto.CompactTextString(m) }
func (*TaskDef) ProtoMessage()    {}
func (*TaskDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_taskdef_a78be32c045d2b55, []int{0}
}
func (m *TaskDef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskDef.Unmarshal(m, b)
}
func (m *TaskDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskDef.Marshal(b, m, deterministic)
}
func (dst *TaskDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskDef.Merge(dst, src)
}
func (m *TaskDef) XXX_Size() int {
	return xxx_messageInfo_TaskDef.Size(m)
}
func (m *TaskDef) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskDef.DiscardUnknown(m)
}

var xxx_messageInfo_TaskDef proto.InternalMessageInfo

func (m *TaskDef) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TaskDef) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *TaskDef) GetRetryCount() int32 {
	if m != nil {
		return m.RetryCount
	}
	return 0
}

func (m *TaskDef) GetTimeoutSeconds() int64 {
	if m != nil {
		return m.TimeoutSeconds
	}
	return 0
}

func (m *TaskDef) GetInputKeys() []string {
	if m != nil {
		return m.InputKeys
	}
	return nil
}

func (m *TaskDef) GetOutputKeys() []string {
	if m != nil {
		return m.OutputKeys
	}
	return nil
}

func (m *TaskDef) GetTimeoutPolicy() TaskDef_TimeoutPolicy {
	if m != nil {
		return m.TimeoutPolicy
	}
	return TaskDef_RETRY
}

func (m *TaskDef) GetRetryLogic() TaskDef_RetryLogic {
	if m != nil {
		return m.RetryLogic
	}
	return TaskDef_FIXED
}

func (m *TaskDef) GetRetryDelaySeconds() int32 {
	if m != nil {
		return m.RetryDelaySeconds
	}
	return 0
}

func (m *TaskDef) GetResponseTimeoutSeconds() int64 {
	if m != nil {
		return m.ResponseTimeoutSeconds
	}
	return 0
}

func (m *TaskDef) GetConcurrentExecLimit() int32 {
	if m != nil {
		return m.ConcurrentExecLimit
	}
	return 0
}

func (m *TaskDef) GetInputTemplate() map[string]*_struct.Value {
	if m != nil {
		return m.InputTemplate
	}
	return nil
}

func (m *TaskDef) GetRateLimitPerFrequency() int32 {
	if m != nil {
		return m.RateLimitPerFrequency
	}
	return 0
}

func (m *TaskDef) GetRateLimitFrequencyInSeconds() int32 {
	if m != nil {
		return m.RateLimitFrequencyInSeconds
	}
	return 0
}

func (m *TaskDef) GetIsolationGroupId() string {
	if m != nil {
		return m.IsolationGroupId
	}
	return ""
}

func (m *TaskDef) GetExecutionNameSpace() string {
	if m != nil {
		return m.ExecutionNameSpace
	}
	return ""
}

func (m *TaskDef) GetOwnerEmail() string {
	if m != nil {
		return m.OwnerEmail
	}
	return ""
}

func (m *TaskDef) GetPollTimeoutSeconds() int32 {
	if m != nil {
		return m.PollTimeoutSeconds
	}
	return 0
}

func init() {
	proto.RegisterType((*TaskDef)(nil), "conductor.proto.TaskDef")
	proto.RegisterMapType((map[string]*_struct.Value)(nil), "conductor.proto.TaskDef.InputTemplateEntry")
	proto.RegisterEnum("conductor.proto.TaskDef_RetryLogic", TaskDef_RetryLogic_name, TaskDef_RetryLogic_value)
	proto.RegisterEnum("conductor.proto.TaskDef_TimeoutPolicy", TaskDef_TimeoutPolicy_name, TaskDef_TimeoutPolicy_value)
}

func init() { proto.RegisterFile("model/taskdef.proto", fileDescriptor_taskdef_a78be32c045d2b55) }

var fileDescriptor_taskdef_a78be32c045d2b55 = []byte{
	// 707 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0x51, 0x6f, 0xdb, 0x36,
	0x10, 0xae, 0xe2, 0xba, 0xad, 0xcf, 0x8b, 0xad, 0xd2, 0x6b, 0x27, 0xb4, 0x1b, 0x6a, 0x64, 0xc0,
	0x66, 0x60, 0x85, 0x5c, 0x78, 0x0f, 0x2b, 0xba, 0xa7, 0xa4, 0x96, 0x07, 0x23, 0x8e, 0x6d, 0x28,
	0xda, 0x96, 0xec, 0x85, 0x90, 0xa9, 0xb3, 0x47, 0x58, 0x12, 0x35, 0x8a, 0xda, 0xa2, 0xdf, 0xb8,
	0x3f, 0x55, 0x90, 0xb2, 0x15, 0x27, 0x41, 0xde, 0xa8, 0xef, 0xfb, 0xee, 0x78, 0xf7, 0xf1, 0x4e,
	0xd0, 0x4b, 0x44, 0x84, 0xf1, 0x50, 0x85, 0xf9, 0x36, 0xc2, 0xb5, 0x9b, 0x49, 0xa1, 0x04, 0xe9,
	0x32, 0x91, 0x46, 0x05, 0x53, 0x42, 0x56, 0xc0, 0x9b, 0x6f, 0x37, 0x42, 0x6c, 0x62, 0x1c, 0x9a,
	0xaf, 0x55, 0xb1, 0x1e, 0xe6, 0x4a, 0x16, 0x4c, 0x55, 0xec, 0xc9, 0xff, 0x2f, 0xe0, 0x79, 0x10,
	0xe6, 0xdb, 0x31, 0xae, 0x09, 0x81, 0xa7, 0x69, 0x98, 0xa0, 0x63, 0xf5, 0xad, 0x41, 0xcb, 0x37,
	0x67, 0xd2, 0x87, 0x76, 0x84, 0x39, 0x93, 0x3c, 0x53, 0x5c, 0xa4, 0xce, 0x91, 0xa1, 0x0e, 0x21,
	0xf2, 0x0e, 0xda, 0x12, 0x95, 0x2c, 0x29, 0x13, 0x45, 0xaa, 0x9c, 0x46, 0xdf, 0x1a, 0x34, 0x7d,
	0x30, 0xd0, 0x67, 0x8d, 0x90, 0x1f, 0xa1, 0xab, 0x78, 0x82, 0xa2, 0x50, 0x34, 0x47, 0x5d, 0x5d,
	0xee, 0x3c, 0xed, 0x5b, 0x83, 0x86, 0xdf, 0xd9, 0xc1, 0x97, 0x15, 0x4a, 0xbe, 0x03, 0xe0, 0x69,
	0x56, 0x28, 0xba, 0xc5, 0x32, 0x77, 0x9a, 0xfd, 0xc6, 0xa0, 0xe5, 0xb7, 0x0c, 0x72, 0x8e, 0x65,
	0xae, 0x2f, 0x12, 0x85, 0xaa, 0xf9, 0x67, 0x86, 0x87, 0x0a, 0x32, 0x82, 0x0b, 0xd8, 0x67, 0xa4,
	0x99, 0x88, 0x39, 0x2b, 0x9d, 0xe7, 0x7d, 0x6b, 0xd0, 0x19, 0xfd, 0xe0, 0xde, 0xf3, 0xc4, 0xdd,
	0x75, 0xec, 0x06, 0x95, 0x7c, 0x69, 0xd4, 0xfe, 0xb1, 0x3a, 0xfc, 0x24, 0xe3, 0x7d, 0x63, 0xb1,
	0xd8, 0x70, 0xe6, 0xbc, 0x30, 0xb9, 0xbe, 0x7f, 0x34, 0x97, 0xaf, 0xb5, 0x33, 0x2d, 0xdd, 0x75,
	0x6f, 0xce, 0xc4, 0x85, 0x5e, 0x95, 0x25, 0xc2, 0x38, 0x2c, 0x6b, 0x07, 0x5a, 0xc6, 0xa6, 0x97,
	0x86, 0x1a, 0x6b, 0x66, 0x6f, 0xc2, 0x47, 0x70, 0x24, 0xe6, 0x99, 0x48, 0x73, 0xa4, 0xf7, 0x6d,
	0x03, 0x63, 0xdb, 0xeb, 0x3d, 0x1f, 0xdc, 0xb5, 0x6f, 0x04, 0xaf, 0x98, 0x48, 0x59, 0x21, 0x25,
	0xa6, 0x8a, 0xe2, 0x0d, 0x32, 0x1a, 0xf3, 0x84, 0x2b, 0xa7, 0x6d, 0xee, 0xea, 0xdd, 0x92, 0xde,
	0x0d, 0xb2, 0x99, 0xa6, 0x88, 0x0f, 0x9d, 0xca, 0x72, 0x85, 0x49, 0x16, 0x87, 0x0a, 0x9d, 0xaf,
	0xfa, 0x8d, 0x41, 0x7b, 0xf4, 0xd3, 0xa3, 0x6d, 0x4e, 0xb5, 0x3c, 0xd8, 0xa9, 0xbd, 0x54, 0xc9,
	0xd2, 0x3f, 0xe6, 0x87, 0x18, 0xf9, 0x05, 0x1c, 0x19, 0x2a, 0xac, 0x2e, 0xa7, 0x19, 0x4a, 0xba,
	0x96, 0xf8, 0x4f, 0x81, 0x29, 0x2b, 0x9d, 0x8e, 0x29, 0xe5, 0x95, 0xe6, 0x4d, 0x01, 0x4b, 0x94,
	0x93, 0x3d, 0x49, 0xc6, 0xf0, 0xee, 0x20, 0xb0, 0x0e, 0xa2, 0x3c, 0xad, 0x1d, 0xe8, 0x9a, 0xf8,
	0xb7, 0x75, 0x7c, 0x1d, 0x3c, 0x4d, 0xf7, 0x36, 0xbc, 0x07, 0xc2, 0x73, 0x11, 0x87, 0x7a, 0x38,
	0xe9, 0x46, 0x8a, 0x22, 0xa3, 0x3c, 0x72, 0x6c, 0x33, 0xb8, 0x76, 0xcd, 0xfc, 0xa6, 0x89, 0x69,
	0x44, 0x3e, 0xc0, 0xd7, 0xda, 0xa9, 0xc2, 0xa8, 0xf5, 0xc4, 0xd3, 0x3c, 0x0b, 0x19, 0x3a, 0x2f,
	0x8d, 0x9e, 0xd4, 0xdc, 0x3c, 0x4c, 0xf0, 0x52, 0x33, 0x66, 0x0c, 0xff, 0x4b, 0x51, 0x52, 0x4c,
	0x42, 0x1e, 0x3b, 0xc4, 0x08, 0xc1, 0x40, 0x9e, 0x46, 0x74, 0xca, 0x4c, 0xc4, 0xf1, 0x83, 0xd7,
	0xeb, 0x99, 0xda, 0x89, 0xe6, 0xee, 0xbe, 0xdc, 0x9b, 0x2b, 0x20, 0x0f, 0x6d, 0x25, 0x36, 0x34,
	0xb6, 0x58, 0xee, 0xb6, 0x51, 0x1f, 0xc9, 0x7b, 0x68, 0xfe, 0x1b, 0xc6, 0x05, 0x9a, 0x35, 0x6c,
	0x8f, 0x5e, 0xbb, 0xd5, 0x6a, 0xbb, 0xfb, 0xd5, 0x76, 0xff, 0xd0, 0xac, 0x5f, 0x89, 0x3e, 0x1d,
	0x7d, 0xb4, 0x4e, 0x3e, 0x00, 0xdc, 0xce, 0x25, 0x69, 0x41, 0x73, 0x32, 0xbd, 0xf2, 0xc6, 0xf6,
	0x13, 0xf2, 0x0d, 0xf4, 0xbc, 0xab, 0xe5, 0x62, 0xee, 0xcd, 0x83, 0xe9, 0xe9, 0x8c, 0x9e, 0x9d,
	0x7e, 0x3e, 0x5f, 0x4c, 0x26, 0xb6, 0x75, 0xf2, 0x2b, 0x1c, 0xdf, 0xd9, 0x0a, 0x1d, 0xe4, 0x7b,
	0x81, 0x7f, 0x6d, 0x3f, 0x21, 0x5d, 0x68, 0x07, 0xd3, 0x0b, 0x8f, 0x2e, 0x7e, 0x0f, 0xe8, 0x9f,
	0x13, 0xdb, 0x22, 0x1d, 0x80, 0xd3, 0x99, 0xe7, 0x07, 0x74, 0x31, 0x9f, 0x5d, 0xdb, 0x47, 0x67,
	0x11, 0xbc, 0x65, 0x22, 0x71, 0x53, 0x54, 0xeb, 0x98, 0xdf, 0xdc, 0x9f, 0xa1, 0xb3, 0xd6, 0x6e,
	0x88, 0x96, 0xab, 0xbf, 0x3e, 0x6d, 0xb8, 0xfa, 0xbb, 0x58, 0xb9, 0x4c, 0x24, 0xc3, 0x9d, 0x7c,
	0x58, 0xcb, 0x87, 0x2c, 0xe6, 0x98, 0xaa, 0xe1, 0x46, 0x6c, 0x64, 0xc6, 0x0e, 0x70, 0xf3, 0xc3,
	0x5b, 0x3d, 0x33, 0xd9, 0x7e, 0xfe, 0x12, 0x00, 0x00, 0xff, 0xff, 0xdb, 0x68, 0x64, 0x27, 0x00,
	0x05, 0x00, 0x00,
}
