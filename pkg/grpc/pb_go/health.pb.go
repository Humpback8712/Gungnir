// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: pkg/grpc/health.proto

package pb_go

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

type HealthCheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterName string `protobuf:"bytes,1,opt,name=clusterName,proto3" json:"clusterName,omitempty"`
}

func (x *HealthCheckRequest) Reset() {
	*x = HealthCheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_health_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckRequest) ProtoMessage() {}

func (x *HealthCheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_health_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckRequest.ProtoReflect.Descriptor instead.
func (*HealthCheckRequest) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_health_proto_rawDescGZIP(), []int{0}
}

func (x *HealthCheckRequest) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

type Resource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CpuTotal  float32 `protobuf:"fixed32,1,opt,name=cpuTotal,proto3" json:"cpuTotal,omitempty"`
	CpuRest   float32 `protobuf:"fixed32,2,opt,name=cpuRest,proto3" json:"cpuRest,omitempty"`
	DiskTotal float32 `protobuf:"fixed32,3,opt,name=diskTotal,proto3" json:"diskTotal,omitempty"`
	DiskRest  float32 `protobuf:"fixed32,4,opt,name=diskRest,proto3" json:"diskRest,omitempty"`
}

func (x *Resource) Reset() {
	*x = Resource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_health_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resource) ProtoMessage() {}

func (x *Resource) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_health_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resource.ProtoReflect.Descriptor instead.
func (*Resource) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_health_proto_rawDescGZIP(), []int{1}
}

func (x *Resource) GetCpuTotal() float32 {
	if x != nil {
		return x.CpuTotal
	}
	return 0
}

func (x *Resource) GetCpuRest() float32 {
	if x != nil {
		return x.CpuRest
	}
	return 0
}

func (x *Resource) GetDiskTotal() float32 {
	if x != nil {
		return x.DiskTotal
	}
	return 0
}

func (x *Resource) GetDiskRest() float32 {
	if x != nil {
		return x.DiskRest
	}
	return 0
}

type HealthPackage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterName string    `protobuf:"bytes,1,opt,name=clusterName,proto3" json:"clusterName,omitempty"`
	Status      string    `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Info        string    `protobuf:"bytes,3,opt,name=info,proto3" json:"info,omitempty"`
	Resource    *Resource `protobuf:"bytes,4,opt,name=resource,proto3" json:"resource,omitempty"`
}

func (x *HealthPackage) Reset() {
	*x = HealthPackage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_health_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthPackage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthPackage) ProtoMessage() {}

func (x *HealthPackage) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_health_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthPackage.ProtoReflect.Descriptor instead.
func (*HealthPackage) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_health_proto_rawDescGZIP(), []int{2}
}

func (x *HealthPackage) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *HealthPackage) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *HealthPackage) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

func (x *HealthPackage) GetResource() *Resource {
	if x != nil {
		return x.Resource
	}
	return nil
}

var File_pkg_grpc_health_proto protoreflect.FileDescriptor

var file_pkg_grpc_health_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x36, 0x0a, 0x12, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a,
	0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0x7a, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x70, 0x75, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x63,
	0x70, 0x75, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x70, 0x75, 0x52, 0x65,
	0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x63, 0x70, 0x75, 0x52, 0x65, 0x73,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69, 0x73, 0x6b, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x64, 0x69, 0x73, 0x6b, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12,
	0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x08, 0x64, 0x69, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x74, 0x22, 0x84, 0x01, 0x0a, 0x0d,
	0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x25, 0x0a, 0x08, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x32, 0x39, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x2f, 0x0a, 0x04,
	0x4b, 0x65, 0x65, 0x70, 0x12, 0x13, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x09, 0x5a,
	0x07, 0x2e, 0x2f, 0x70, 0x62, 0x5f, 0x67, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_grpc_health_proto_rawDescOnce sync.Once
	file_pkg_grpc_health_proto_rawDescData = file_pkg_grpc_health_proto_rawDesc
)

func file_pkg_grpc_health_proto_rawDescGZIP() []byte {
	file_pkg_grpc_health_proto_rawDescOnce.Do(func() {
		file_pkg_grpc_health_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_grpc_health_proto_rawDescData)
	})
	return file_pkg_grpc_health_proto_rawDescData
}

var file_pkg_grpc_health_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pkg_grpc_health_proto_goTypes = []interface{}{
	(*HealthCheckRequest)(nil), // 0: HealthCheckRequest
	(*Resource)(nil),           // 1: Resource
	(*HealthPackage)(nil),      // 2: HealthPackage
}
var file_pkg_grpc_health_proto_depIdxs = []int32{
	1, // 0: HealthPackage.resource:type_name -> Resource
	0, // 1: Health.Keep:input_type -> HealthCheckRequest
	2, // 2: Health.Keep:output_type -> HealthPackage
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_grpc_health_proto_init() }
func file_pkg_grpc_health_proto_init() {
	if File_pkg_grpc_health_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_grpc_health_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckRequest); i {
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
		file_pkg_grpc_health_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resource); i {
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
		file_pkg_grpc_health_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthPackage); i {
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
			RawDescriptor: file_pkg_grpc_health_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_grpc_health_proto_goTypes,
		DependencyIndexes: file_pkg_grpc_health_proto_depIdxs,
		MessageInfos:      file_pkg_grpc_health_proto_msgTypes,
	}.Build()
	File_pkg_grpc_health_proto = out.File
	file_pkg_grpc_health_proto_rawDesc = nil
	file_pkg_grpc_health_proto_goTypes = nil
	file_pkg_grpc_health_proto_depIdxs = nil
}
