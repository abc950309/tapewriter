// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.10
// source: job.proto

package entity

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type JobStatus int32

const (
	JobStatus_Draft      JobStatus = 0
	JobStatus_NotReady   JobStatus = 1 // dependencies not satisfied
	JobStatus_Pending    JobStatus = 2 // waiting in queue
	JobStatus_Processing JobStatus = 3
	JobStatus_Completed  JobStatus = 4
	JobStatus_Failed     JobStatus = 255
)

// Enum value maps for JobStatus.
var (
	JobStatus_name = map[int32]string{
		0:   "Draft",
		1:   "NotReady",
		2:   "Pending",
		3:   "Processing",
		4:   "Completed",
		255: "Failed",
	}
	JobStatus_value = map[string]int32{
		"Draft":      0,
		"NotReady":   1,
		"Pending":    2,
		"Processing": 3,
		"Completed":  4,
		"Failed":     255,
	}
)

func (x JobStatus) Enum() *JobStatus {
	p := new(JobStatus)
	*p = x
	return p
}

func (x JobStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JobStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_job_proto_enumTypes[0].Descriptor()
}

func (JobStatus) Type() protoreflect.EnumType {
	return &file_job_proto_enumTypes[0]
}

func (x JobStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JobStatus.Descriptor instead.
func (JobStatus) EnumDescriptor() ([]byte, []int) {
	return file_job_proto_rawDescGZIP(), []int{0}
}

type Job struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Status     JobStatus `protobuf:"varint,2,opt,name=status,proto3,enum=job.JobStatus" json:"status,omitempty"`
	Priority   int64     `protobuf:"varint,3,opt,name=priority,proto3" json:"priority,omitempty"`
	CreateTime int64     `protobuf:"varint,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime int64     `protobuf:"varint,5,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	State      *JobState `protobuf:"bytes,17,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *Job) Reset() {
	*x = Job{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Job) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Job) ProtoMessage() {}

func (x *Job) ProtoReflect() protoreflect.Message {
	mi := &file_job_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Job.ProtoReflect.Descriptor instead.
func (*Job) Descriptor() ([]byte, []int) {
	return file_job_proto_rawDescGZIP(), []int{0}
}

func (x *Job) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Job) GetStatus() JobStatus {
	if x != nil {
		return x.Status
	}
	return JobStatus_Draft
}

func (x *Job) GetPriority() int64 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *Job) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *Job) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

func (x *Job) GetState() *JobState {
	if x != nil {
		return x.State
	}
	return nil
}

type JobParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Param:
	//	*JobParam_Archive
	Param isJobParam_Param `protobuf_oneof:"param"`
}

func (x *JobParam) Reset() {
	*x = JobParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobParam) ProtoMessage() {}

func (x *JobParam) ProtoReflect() protoreflect.Message {
	mi := &file_job_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobParam.ProtoReflect.Descriptor instead.
func (*JobParam) Descriptor() ([]byte, []int) {
	return file_job_proto_rawDescGZIP(), []int{1}
}

func (m *JobParam) GetParam() isJobParam_Param {
	if m != nil {
		return m.Param
	}
	return nil
}

func (x *JobParam) GetArchive() *JobParamArchive {
	if x, ok := x.GetParam().(*JobParam_Archive); ok {
		return x.Archive
	}
	return nil
}

type isJobParam_Param interface {
	isJobParam_Param()
}

type JobParam_Archive struct {
	Archive *JobParamArchive `protobuf:"bytes,1,opt,name=Archive,proto3,oneof"`
}

func (*JobParam_Archive) isJobParam_Param() {}

type JobState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to State:
	//	*JobState_Archive
	State isJobState_State `protobuf_oneof:"state"`
}

func (x *JobState) Reset() {
	*x = JobState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobState) ProtoMessage() {}

func (x *JobState) ProtoReflect() protoreflect.Message {
	mi := &file_job_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobState.ProtoReflect.Descriptor instead.
func (*JobState) Descriptor() ([]byte, []int) {
	return file_job_proto_rawDescGZIP(), []int{2}
}

func (m *JobState) GetState() isJobState_State {
	if m != nil {
		return m.State
	}
	return nil
}

func (x *JobState) GetArchive() *JobStateArchive {
	if x, ok := x.GetState().(*JobState_Archive); ok {
		return x.Archive
	}
	return nil
}

type isJobState_State interface {
	isJobState_State()
}

type JobState_Archive struct {
	Archive *JobStateArchive `protobuf:"bytes,1,opt,name=Archive,proto3,oneof"`
}

func (*JobState_Archive) isJobState_State() {}

type JobNextParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Param:
	//	*JobNextParam_Archive
	Param isJobNextParam_Param `protobuf_oneof:"param"`
}

func (x *JobNextParam) Reset() {
	*x = JobNextParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobNextParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobNextParam) ProtoMessage() {}

func (x *JobNextParam) ProtoReflect() protoreflect.Message {
	mi := &file_job_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobNextParam.ProtoReflect.Descriptor instead.
func (*JobNextParam) Descriptor() ([]byte, []int) {
	return file_job_proto_rawDescGZIP(), []int{3}
}

func (m *JobNextParam) GetParam() isJobNextParam_Param {
	if m != nil {
		return m.Param
	}
	return nil
}

func (x *JobNextParam) GetArchive() *JobArchiveNextParam {
	if x, ok := x.GetParam().(*JobNextParam_Archive); ok {
		return x.Archive
	}
	return nil
}

type isJobNextParam_Param interface {
	isJobNextParam_Param()
}

type JobNextParam_Archive struct {
	Archive *JobArchiveNextParam `protobuf:"bytes,1,opt,name=archive,proto3,oneof"`
}

func (*JobNextParam_Archive) isJobNextParam_Param() {}

type CreatableJob struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Priority int64     `protobuf:"varint,3,opt,name=priority,proto3" json:"priority,omitempty"`
	Param    *JobParam `protobuf:"bytes,17,opt,name=param,proto3" json:"param,omitempty"`
}

func (x *CreatableJob) Reset() {
	*x = CreatableJob{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatableJob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatableJob) ProtoMessage() {}

func (x *CreatableJob) ProtoReflect() protoreflect.Message {
	mi := &file_job_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatableJob.ProtoReflect.Descriptor instead.
func (*CreatableJob) Descriptor() ([]byte, []int) {
	return file_job_proto_rawDescGZIP(), []int{4}
}

func (x *CreatableJob) GetPriority() int64 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *CreatableJob) GetParam() *JobParam {
	if x != nil {
		return x.Param
	}
	return nil
}

type JobFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *JobStatus `protobuf:"varint,1,opt,name=status,proto3,enum=job.JobStatus,oneof" json:"status,omitempty"`
	Limit  *int64     `protobuf:"varint,33,opt,name=limit,proto3,oneof" json:"limit,omitempty"`
	Offset *int64     `protobuf:"varint,34,opt,name=offset,proto3,oneof" json:"offset,omitempty"`
}

func (x *JobFilter) Reset() {
	*x = JobFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobFilter) ProtoMessage() {}

func (x *JobFilter) ProtoReflect() protoreflect.Message {
	mi := &file_job_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobFilter.ProtoReflect.Descriptor instead.
func (*JobFilter) Descriptor() ([]byte, []int) {
	return file_job_proto_rawDescGZIP(), []int{5}
}

func (x *JobFilter) GetStatus() JobStatus {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return JobStatus_Draft
}

func (x *JobFilter) GetLimit() int64 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

func (x *JobFilter) GetOffset() int64 {
	if x != nil && x.Offset != nil {
		return *x.Offset
	}
	return 0
}

type JobDisplay struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Display:
	//	*JobDisplay_Archive
	Display isJobDisplay_Display `protobuf_oneof:"display"`
}

func (x *JobDisplay) Reset() {
	*x = JobDisplay{}
	if protoimpl.UnsafeEnabled {
		mi := &file_job_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobDisplay) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobDisplay) ProtoMessage() {}

func (x *JobDisplay) ProtoReflect() protoreflect.Message {
	mi := &file_job_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobDisplay.ProtoReflect.Descriptor instead.
func (*JobDisplay) Descriptor() ([]byte, []int) {
	return file_job_proto_rawDescGZIP(), []int{6}
}

func (m *JobDisplay) GetDisplay() isJobDisplay_Display {
	if m != nil {
		return m.Display
	}
	return nil
}

func (x *JobDisplay) GetArchive() *JobDisplayArchive {
	if x, ok := x.GetDisplay().(*JobDisplay_Archive); ok {
		return x.Archive
	}
	return nil
}

type isJobDisplay_Display interface {
	isJobDisplay_Display()
}

type JobDisplay_Archive struct {
	Archive *JobDisplayArchive `protobuf:"bytes,1,opt,name=archive,proto3,oneof"`
}

func (*JobDisplay_Archive) isJobDisplay_Display() {}

var File_job_proto protoreflect.FileDescriptor

var file_job_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6a, 0x6f, 0x62,
	0x1a, 0x11, 0x6a, 0x6f, 0x62, 0x5f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x01, 0x0a, 0x03, 0x4a, 0x6f, 0x62, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x6a, 0x6f,
	0x62, 0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12,
	0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x23, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x4d, 0x0a, 0x08, 0x4a, 0x6f, 0x62, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x12, 0x38, 0x0a, 0x07, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76,
	0x65, 0x2e, 0x4a, 0x6f, 0x62, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76,
	0x65, 0x48, 0x00, 0x52, 0x07, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x42, 0x07, 0x0a, 0x05,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x22, 0x4d, 0x0a, 0x08, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x38, 0x0a, 0x07, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65,
	0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x65, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65,
	0x48, 0x00, 0x52, 0x07, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x22, 0x55, 0x0a, 0x0c, 0x4a, 0x6f, 0x62, 0x4e, 0x65, 0x78, 0x74, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x12, 0x3c, 0x0a, 0x07, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x61, 0x72, 0x63, 0x68,
	0x69, 0x76, 0x65, 0x2e, 0x4a, 0x6f, 0x62, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x4e, 0x65,
	0x78, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x48, 0x00, 0x52, 0x07, 0x61, 0x72, 0x63, 0x68, 0x69,
	0x76, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x22, 0x4f, 0x0a, 0x0c, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4a, 0x6f, 0x62, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70,
	0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x4a, 0x6f, 0x62,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x22, 0x90, 0x01, 0x0a,
	0x09, 0x4a, 0x6f, 0x62, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x6a, 0x6f, 0x62,
	0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48, 0x00, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x21, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x88,
	0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x22, 0x20, 0x01,
	0x28, 0x03, 0x48, 0x02, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x88, 0x01, 0x01, 0x42,
	0x09, 0x0a, 0x07, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22,
	0x53, 0x0a, 0x0a, 0x4a, 0x6f, 0x62, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x12, 0x3a, 0x0a,
	0x07, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x6a, 0x6f, 0x62, 0x5f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x2e, 0x4a, 0x6f, 0x62,
	0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x48, 0x00,
	0x52, 0x07, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x2a, 0x5d, 0x0a, 0x09, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x72, 0x61, 0x66, 0x74, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08,
	0x4e, 0x6f, 0x74, 0x52, 0x65, 0x61, 0x64, 0x79, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x65,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x69, 0x6e, 0x67, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x06, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64,
	0x10, 0xff, 0x01, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x62, 0x63, 0x39, 0x35, 0x30, 0x33, 0x30, 0x39, 0x2f, 0x74, 0x61, 0x70, 0x65,
	0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_job_proto_rawDescOnce sync.Once
	file_job_proto_rawDescData = file_job_proto_rawDesc
)

func file_job_proto_rawDescGZIP() []byte {
	file_job_proto_rawDescOnce.Do(func() {
		file_job_proto_rawDescData = protoimpl.X.CompressGZIP(file_job_proto_rawDescData)
	})
	return file_job_proto_rawDescData
}

var file_job_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_job_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_job_proto_goTypes = []interface{}{
	(JobStatus)(0),              // 0: job.JobStatus
	(*Job)(nil),                 // 1: job.Job
	(*JobParam)(nil),            // 2: job.JobParam
	(*JobState)(nil),            // 3: job.JobState
	(*JobNextParam)(nil),        // 4: job.JobNextParam
	(*CreatableJob)(nil),        // 5: job.CreatableJob
	(*JobFilter)(nil),           // 6: job.JobFilter
	(*JobDisplay)(nil),          // 7: job.JobDisplay
	(*JobParamArchive)(nil),     // 8: job_archive.JobParamArchive
	(*JobStateArchive)(nil),     // 9: job_archive.JobStateArchive
	(*JobArchiveNextParam)(nil), // 10: job_archive.JobArchiveNextParam
	(*JobDisplayArchive)(nil),   // 11: job_archive.JobDisplayArchive
}
var file_job_proto_depIdxs = []int32{
	0,  // 0: job.Job.status:type_name -> job.JobStatus
	3,  // 1: job.Job.state:type_name -> job.JobState
	8,  // 2: job.JobParam.Archive:type_name -> job_archive.JobParamArchive
	9,  // 3: job.JobState.Archive:type_name -> job_archive.JobStateArchive
	10, // 4: job.JobNextParam.archive:type_name -> job_archive.JobArchiveNextParam
	2,  // 5: job.CreatableJob.param:type_name -> job.JobParam
	0,  // 6: job.JobFilter.status:type_name -> job.JobStatus
	11, // 7: job.JobDisplay.archive:type_name -> job_archive.JobDisplayArchive
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_job_proto_init() }
func file_job_proto_init() {
	if File_job_proto != nil {
		return
	}
	file_job_archive_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_job_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Job); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_job_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobParam); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_job_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobState); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_job_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobNextParam); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_job_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatableJob); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_job_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobFilter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_job_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobDisplay); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_job_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*JobParam_Archive)(nil),
	}
	file_job_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*JobState_Archive)(nil),
	}
	file_job_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*JobNextParam_Archive)(nil),
	}
	file_job_proto_msgTypes[5].OneofWrappers = []interface{}{}
	file_job_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*JobDisplay_Archive)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_job_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_job_proto_goTypes,
		DependencyIndexes: file_job_proto_depIdxs,
		EnumInfos:         file_job_proto_enumTypes,
		MessageInfos:      file_job_proto_msgTypes,
	}.Build()
	File_job_proto = out.File
	file_job_proto_rawDesc = nil
	file_job_proto_goTypes = nil
	file_job_proto_depIdxs = nil
}
