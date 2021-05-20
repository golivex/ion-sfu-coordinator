// Copyright 2020 Google LLC
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: google/appengine/v1/instance.proto

package appengine

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Availability of the instance.
type Instance_Availability int32

const (
	Instance_UNSPECIFIED Instance_Availability = 0
	Instance_RESIDENT    Instance_Availability = 1
	Instance_DYNAMIC     Instance_Availability = 2
)

// Enum value maps for Instance_Availability.
var (
	Instance_Availability_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "RESIDENT",
		2: "DYNAMIC",
	}
	Instance_Availability_value = map[string]int32{
		"UNSPECIFIED": 0,
		"RESIDENT":    1,
		"DYNAMIC":     2,
	}
)

func (x Instance_Availability) Enum() *Instance_Availability {
	p := new(Instance_Availability)
	*p = x
	return p
}

func (x Instance_Availability) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Instance_Availability) Descriptor() protoreflect.EnumDescriptor {
	return file_google_appengine_v1_instance_proto_enumTypes[0].Descriptor()
}

func (Instance_Availability) Type() protoreflect.EnumType {
	return &file_google_appengine_v1_instance_proto_enumTypes[0]
}

func (x Instance_Availability) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Instance_Availability.Descriptor instead.
func (Instance_Availability) EnumDescriptor() ([]byte, []int) {
	return file_google_appengine_v1_instance_proto_rawDescGZIP(), []int{0, 0}
}

// An Instance resource is the computing unit that App Engine uses to
// automatically scale an application.
type Instance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Full path to the Instance resource in the API.
	// Example: `apps/myapp/services/default/versions/v1/instances/instance-1`.
	//
	// @OutputOnly
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Relative name of the instance within the version.
	// Example: `instance-1`.
	//
	// @OutputOnly
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// App Engine release this instance is running on.
	//
	// @OutputOnly
	AppEngineRelease string `protobuf:"bytes,3,opt,name=app_engine_release,json=appEngineRelease,proto3" json:"app_engine_release,omitempty"`
	// Availability of the instance.
	//
	// @OutputOnly
	Availability Instance_Availability `protobuf:"varint,4,opt,name=availability,proto3,enum=google.appengine.v1.Instance_Availability" json:"availability,omitempty"`
	// Name of the virtual machine where this instance lives. Only applicable
	// for instances in App Engine flexible environment.
	//
	// @OutputOnly
	VmName string `protobuf:"bytes,5,opt,name=vm_name,json=vmName,proto3" json:"vm_name,omitempty"`
	// Zone where the virtual machine is located. Only applicable for instances
	// in App Engine flexible environment.
	//
	// @OutputOnly
	VmZoneName string `protobuf:"bytes,6,opt,name=vm_zone_name,json=vmZoneName,proto3" json:"vm_zone_name,omitempty"`
	// Virtual machine ID of this instance. Only applicable for instances in
	// App Engine flexible environment.
	//
	// @OutputOnly
	VmId string `protobuf:"bytes,7,opt,name=vm_id,json=vmId,proto3" json:"vm_id,omitempty"`
	// Time that this instance was started.
	//
	// @OutputOnly
	StartTime *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// Number of requests since this instance was started.
	//
	// @OutputOnly
	Requests int32 `protobuf:"varint,9,opt,name=requests,proto3" json:"requests,omitempty"`
	// Number of errors since this instance was started.
	//
	// @OutputOnly
	Errors int32 `protobuf:"varint,10,opt,name=errors,proto3" json:"errors,omitempty"`
	// Average queries per second (QPS) over the last minute.
	//
	// @OutputOnly
	Qps float32 `protobuf:"fixed32,11,opt,name=qps,proto3" json:"qps,omitempty"`
	// Average latency (ms) over the last minute.
	//
	// @OutputOnly
	AverageLatency int32 `protobuf:"varint,12,opt,name=average_latency,json=averageLatency,proto3" json:"average_latency,omitempty"`
	// Total memory in use (bytes).
	//
	// @OutputOnly
	MemoryUsage int64 `protobuf:"varint,13,opt,name=memory_usage,json=memoryUsage,proto3" json:"memory_usage,omitempty"`
	// Status of the virtual machine where this instance lives. Only applicable
	// for instances in App Engine flexible environment.
	//
	// @OutputOnly
	VmStatus string `protobuf:"bytes,14,opt,name=vm_status,json=vmStatus,proto3" json:"vm_status,omitempty"`
	// Whether this instance is in debug mode. Only applicable for instances in
	// App Engine flexible environment.
	//
	// @OutputOnly
	VmDebugEnabled bool `protobuf:"varint,15,opt,name=vm_debug_enabled,json=vmDebugEnabled,proto3" json:"vm_debug_enabled,omitempty"`
	// The IP address of this instance. Only applicable for instances in App
	// Engine flexible environment.
	//
	// @OutputOnly
	VmIp string `protobuf:"bytes,16,opt,name=vm_ip,json=vmIp,proto3" json:"vm_ip,omitempty"`
}

func (x *Instance) Reset() {
	*x = Instance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_appengine_v1_instance_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Instance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Instance) ProtoMessage() {}

func (x *Instance) ProtoReflect() protoreflect.Message {
	mi := &file_google_appengine_v1_instance_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Instance.ProtoReflect.Descriptor instead.
func (*Instance) Descriptor() ([]byte, []int) {
	return file_google_appengine_v1_instance_proto_rawDescGZIP(), []int{0}
}

func (x *Instance) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Instance) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Instance) GetAppEngineRelease() string {
	if x != nil {
		return x.AppEngineRelease
	}
	return ""
}

func (x *Instance) GetAvailability() Instance_Availability {
	if x != nil {
		return x.Availability
	}
	return Instance_UNSPECIFIED
}

func (x *Instance) GetVmName() string {
	if x != nil {
		return x.VmName
	}
	return ""
}

func (x *Instance) GetVmZoneName() string {
	if x != nil {
		return x.VmZoneName
	}
	return ""
}

func (x *Instance) GetVmId() string {
	if x != nil {
		return x.VmId
	}
	return ""
}

func (x *Instance) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *Instance) GetRequests() int32 {
	if x != nil {
		return x.Requests
	}
	return 0
}

func (x *Instance) GetErrors() int32 {
	if x != nil {
		return x.Errors
	}
	return 0
}

func (x *Instance) GetQps() float32 {
	if x != nil {
		return x.Qps
	}
	return 0
}

func (x *Instance) GetAverageLatency() int32 {
	if x != nil {
		return x.AverageLatency
	}
	return 0
}

func (x *Instance) GetMemoryUsage() int64 {
	if x != nil {
		return x.MemoryUsage
	}
	return 0
}

func (x *Instance) GetVmStatus() string {
	if x != nil {
		return x.VmStatus
	}
	return ""
}

func (x *Instance) GetVmDebugEnabled() bool {
	if x != nil {
		return x.VmDebugEnabled
	}
	return false
}

func (x *Instance) GetVmIp() string {
	if x != nil {
		return x.VmIp
	}
	return ""
}

var File_google_appengine_v1_instance_proto protoreflect.FileDescriptor

var file_google_appengine_v1_instance_proto_rawDesc = []byte{
	0x0a, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe1, 0x04, 0x0a, 0x08, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x61, 0x70, 0x70,
	0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x61, 0x70, 0x70, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0c, 0x61, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x41, 0x76, 0x61,
	0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x0c, 0x61, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x76, 0x6d, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x76, 0x6d, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0c, 0x76, 0x6d, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x76, 0x6d, 0x5a, 0x6f, 0x6e, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x13, 0x0a, 0x05, 0x76, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x76, 0x6d, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x71, 0x70, 0x73, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x03, 0x71, 0x70, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x76, 0x65, 0x72,
	0x61, 0x67, 0x65, 0x5f, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0e, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x4c, 0x61, 0x74, 0x65, 0x6e, 0x63,
	0x79, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x75, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x55,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x76, 0x6d, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x76, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x28, 0x0a, 0x10, 0x76, 0x6d, 0x5f, 0x64, 0x65, 0x62, 0x75, 0x67, 0x5f, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x76, 0x6d, 0x44,
	0x65, 0x62, 0x75, 0x67, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x13, 0x0a, 0x05, 0x76,
	0x6d, 0x5f, 0x69, 0x70, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x76, 0x6d, 0x49, 0x70,
	0x22, 0x3a, 0x0a, 0x0c, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x53, 0x49, 0x44, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x12,
	0x0b, 0x0a, 0x07, 0x44, 0x59, 0x4e, 0x41, 0x4d, 0x49, 0x43, 0x10, 0x02, 0x42, 0xbf, 0x01, 0x0a,
	0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x70,
	0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0xaa, 0x02, 0x19, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x19, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x5c, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5c, 0x56, 0x31, 0xea,
	0x02, 0x1c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a,
	0x3a, 0x41, 0x70, 0x70, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_appengine_v1_instance_proto_rawDescOnce sync.Once
	file_google_appengine_v1_instance_proto_rawDescData = file_google_appengine_v1_instance_proto_rawDesc
)

func file_google_appengine_v1_instance_proto_rawDescGZIP() []byte {
	file_google_appengine_v1_instance_proto_rawDescOnce.Do(func() {
		file_google_appengine_v1_instance_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_appengine_v1_instance_proto_rawDescData)
	})
	return file_google_appengine_v1_instance_proto_rawDescData
}

var file_google_appengine_v1_instance_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_appengine_v1_instance_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_appengine_v1_instance_proto_goTypes = []interface{}{
	(Instance_Availability)(0),    // 0: google.appengine.v1.Instance.Availability
	(*Instance)(nil),              // 1: google.appengine.v1.Instance
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_google_appengine_v1_instance_proto_depIdxs = []int32{
	0, // 0: google.appengine.v1.Instance.availability:type_name -> google.appengine.v1.Instance.Availability
	2, // 1: google.appengine.v1.Instance.start_time:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_appengine_v1_instance_proto_init() }
func file_google_appengine_v1_instance_proto_init() {
	if File_google_appengine_v1_instance_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_appengine_v1_instance_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Instance); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_appengine_v1_instance_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_appengine_v1_instance_proto_goTypes,
		DependencyIndexes: file_google_appengine_v1_instance_proto_depIdxs,
		EnumInfos:         file_google_appengine_v1_instance_proto_enumTypes,
		MessageInfos:      file_google_appengine_v1_instance_proto_msgTypes,
	}.Build()
	File_google_appengine_v1_instance_proto = out.File
	file_google_appengine_v1_instance_proto_rawDesc = nil
	file_google_appengine_v1_instance_proto_goTypes = nil
	file_google_appengine_v1_instance_proto_depIdxs = nil
}
