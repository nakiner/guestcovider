// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.15.2
// source: guestcovider-services.proto

package guestcoviderpb

import (
	context "context"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_guestcovider_services_proto protoreflect.FileDescriptor

var file_guestcovider_services_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x67, 0x69, 0x6d, 0x61, 0x2d, 0x67, 0x75, 0x65, 0x73, 0x74, 0x63, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x13, 0x61, 0x67, 0x69, 0x6d, 0x61, 0x67, 0x75, 0x65, 0x73, 0x74, 0x63,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x70, 0x62, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67,
	0x65, 0x6e, 0x2d, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x61, 0x67, 0x69, 0x6d, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x61, 0x67, 0x69, 0x6d, 0x61,
	0x2d, 0x67, 0x75, 0x65, 0x73, 0x74, 0x63, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2d, 0x68, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x61, 0x67, 0x69, 0x6d,
	0x61, 0x2d, 0x67, 0x75, 0x65, 0x73, 0x74, 0x63, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2d, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xa6, 0x03, 0x0a, 0x0d, 0x48, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7a, 0x0a, 0x08, 0x4c,
	0x69, 0x76, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x24, 0x2e, 0x61, 0x67, 0x69, 0x6d, 0x61, 0x67,
	0x75, 0x65, 0x73, 0x74, 0x63, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x4c, 0x69,
	0x76, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e,
	0x61, 0x67, 0x69, 0x6d, 0x61, 0x67, 0x75, 0x65, 0x73, 0x74, 0x63, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x76, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f, 0x6c,
	0x69, 0x76, 0x65, 0x6e, 0x65, 0x73, 0x73, 0x92, 0x41, 0x0d, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x7e, 0x0a, 0x09, 0x52, 0x65, 0x61, 0x64, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x12, 0x25, 0x2e, 0x61, 0x67, 0x69, 0x6d, 0x61, 0x67, 0x75, 0x65, 0x73,
	0x74, 0x63, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x61, 0x67,
	0x69, 0x6d, 0x61, 0x67, 0x75, 0x65, 0x73, 0x74, 0x63, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x70,
	0x62, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x72, 0x65,
	0x61, 0x64, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x92, 0x41, 0x0d, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x76, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x23, 0x2e, 0x61, 0x67, 0x69, 0x6d, 0x61, 0x67, 0x75, 0x65, 0x73, 0x74, 0x63,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x67, 0x69, 0x6d, 0x61, 0x67,
	0x75, 0x65, 0x73, 0x74, 0x63, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x92,
	0x41, 0x0d, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x1a,
	0x21, 0xa2, 0xc5, 0xb6, 0x03, 0x1c, 0x0a, 0x02, 0x10, 0x01, 0x12, 0x02, 0x10, 0x01, 0x1a, 0x02,
	0x10, 0x01, 0x22, 0x02, 0x10, 0x01, 0x2a, 0x02, 0x10, 0x01, 0x32, 0x02, 0x10, 0x01, 0x3a, 0x02,
	0x10, 0x00, 0x32, 0xa5, 0x02, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x7c, 0x0a, 0x0a, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x26, 0x2e, 0x61, 0x67, 0x69, 0x6d, 0x61, 0x67, 0x75, 0x65, 0x73, 0x74, 0x63, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x61, 0x67, 0x69, 0x6d, 0x61,
	0x67, 0x75, 0x65, 0x73, 0x74, 0x63, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x92, 0x41, 0x06, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x12, 0x75, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x26,
	0x2e, 0x61, 0x67, 0x69, 0x6d, 0x61, 0x67, 0x75, 0x65, 0x73, 0x74, 0x63, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x61, 0x67, 0x69, 0x6d, 0x61, 0x67, 0x75,
	0x65, 0x73, 0x74, 0x63, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x07, 0x1a, 0x05, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x92, 0x41,
	0x06, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x21, 0xa2, 0xc5, 0xb6, 0x03, 0x1c, 0x0a, 0x02,
	0x10, 0x01, 0x12, 0x02, 0x10, 0x01, 0x1a, 0x02, 0x10, 0x01, 0x22, 0x02, 0x10, 0x01, 0x2a, 0x02,
	0x10, 0x01, 0x32, 0x02, 0x10, 0x01, 0x3a, 0x02, 0x10, 0x00, 0x42, 0xa4, 0x01, 0x5a, 0x1c, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x67, 0x69, 0x6d, 0x61, 0x67, 0x75, 0x65,
	0x73, 0x74, 0x63, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x70, 0x62, 0x92, 0x41, 0x82, 0x01, 0x12,
	0x1c, 0x0a, 0x15, 0x43, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x20, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x2a, 0x01, 0x01,
	0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73,
	0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x6a, 0x73, 0x6f, 0x6e, 0x52, 0x3b, 0x0a, 0x03, 0x34, 0x30, 0x34, 0x12, 0x34, 0x0a, 0x2a, 0x52,
	0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x74, 0x68, 0x65,
	0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x64, 0x6f, 0x65, 0x73, 0x20, 0x6e,
	0x6f, 0x74, 0x20, 0x65, 0x78, 0x69, 0x73, 0x74, 0x2e, 0x12, 0x06, 0x0a, 0x04, 0x9a, 0x02, 0x01,
	0x07, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_guestcovider_services_proto_goTypes = []interface{}{
	(*LivenessRequest)(nil),    // 0: guestcoviderpb.LivenessRequest
	(*ReadinessRequest)(nil),   // 1: guestcoviderpb.ReadinessRequest
	(*VersionRequest)(nil),     // 2: guestcoviderpb.VersionRequest
	(*SearchUserRequest)(nil),  // 3: guestcoviderpb.SearchUserRequest
	(*UpdateUserRequest)(nil),  // 4: guestcoviderpb.UpdateUserRequest
	(*LivenessResponse)(nil),   // 5: guestcoviderpb.LivenessResponse
	(*ReadinessResponse)(nil),  // 6: guestcoviderpb.ReadinessResponse
	(*VersionResponse)(nil),    // 7: guestcoviderpb.VersionResponse
	(*SearchUserResponse)(nil), // 8: guestcoviderpb.SearchUserResponse
	(*UpdateUserResponse)(nil), // 9: guestcoviderpb.UpdateUserResponse
}
var file_guestcovider_services_proto_depIdxs = []int32{
	0, // 0: guestcoviderpb.HealthService.Liveness:input_type -> guestcoviderpb.LivenessRequest
	1, // 1: guestcoviderpb.HealthService.Readiness:input_type -> guestcoviderpb.ReadinessRequest
	2, // 2: guestcoviderpb.HealthService.Version:input_type -> guestcoviderpb.VersionRequest
	3, // 3: guestcoviderpb.UserService.SearchUser:input_type -> guestcoviderpb.SearchUserRequest
	4, // 4: guestcoviderpb.UserService.UpdateUser:input_type -> guestcoviderpb.UpdateUserRequest
	5, // 5: guestcoviderpb.HealthService.Liveness:output_type -> guestcoviderpb.LivenessResponse
	6, // 6: guestcoviderpb.HealthService.Readiness:output_type -> guestcoviderpb.ReadinessResponse
	7, // 7: guestcoviderpb.HealthService.Version:output_type -> guestcoviderpb.VersionResponse
	8, // 8: guestcoviderpb.UserService.SearchUser:output_type -> guestcoviderpb.SearchUserResponse
	9, // 9: guestcoviderpb.UserService.UpdateUser:output_type -> guestcoviderpb.UpdateUserResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_guestcovider_services_proto_init() }
func file_guestcovider_services_proto_init() {
	if File_guestcovider_services_proto != nil {
		return
	}
	file_guestcovider_health_proto_init()
	file_guestcovider_user_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_guestcovider_services_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_guestcovider_services_proto_goTypes,
		DependencyIndexes: file_guestcovider_services_proto_depIdxs,
	}.Build()
	File_guestcovider_services_proto = out.File
	file_guestcovider_services_proto_rawDesc = nil
	file_guestcovider_services_proto_goTypes = nil
	file_guestcovider_services_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HealthServiceClient is the client API for HealthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HealthServiceClient interface {
	// returns a error if service doesn`t live.
	Liveness(ctx context.Context, in *LivenessRequest, opts ...grpc.CallOption) (*LivenessResponse, error)
	// returns a error if service doesn`t ready.
	Readiness(ctx context.Context, in *ReadinessRequest, opts ...grpc.CallOption) (*ReadinessResponse, error)
	// returns build time, last commit and version app
	Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error)
}

type healthServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHealthServiceClient(cc grpc.ClientConnInterface) HealthServiceClient {
	return &healthServiceClient{cc}
}

func (c *healthServiceClient) Liveness(ctx context.Context, in *LivenessRequest, opts ...grpc.CallOption) (*LivenessResponse, error) {
	out := new(LivenessResponse)
	err := c.cc.Invoke(ctx, "/guestcoviderpb.HealthService/Liveness", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthServiceClient) Readiness(ctx context.Context, in *ReadinessRequest, opts ...grpc.CallOption) (*ReadinessResponse, error) {
	out := new(ReadinessResponse)
	err := c.cc.Invoke(ctx, "/guestcoviderpb.HealthService/Readiness", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthServiceClient) Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/guestcoviderpb.HealthService/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HealthServiceServer is the server API for HealthService service.
type HealthServiceServer interface {
	// returns a error if service doesn`t live.
	Liveness(context.Context, *LivenessRequest) (*LivenessResponse, error)
	// returns a error if service doesn`t ready.
	Readiness(context.Context, *ReadinessRequest) (*ReadinessResponse, error)
	// returns build time, last commit and version app
	Version(context.Context, *VersionRequest) (*VersionResponse, error)
}

// UnimplementedHealthServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHealthServiceServer struct {
}

func (*UnimplementedHealthServiceServer) Liveness(context.Context, *LivenessRequest) (*LivenessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Liveness not implemented")
}
func (*UnimplementedHealthServiceServer) Readiness(context.Context, *ReadinessRequest) (*ReadinessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Readiness not implemented")
}
func (*UnimplementedHealthServiceServer) Version(context.Context, *VersionRequest) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}

func RegisterHealthServiceServer(s *grpc.Server, srv HealthServiceServer) {
	s.RegisterService(&_HealthService_serviceDesc, srv)
}

func _HealthService_Liveness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LivenessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServiceServer).Liveness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/guestcoviderpb.HealthService/Liveness",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServiceServer).Liveness(ctx, req.(*LivenessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthService_Readiness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadinessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServiceServer).Readiness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/guestcoviderpb.HealthService/Readiness",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServiceServer).Readiness(ctx, req.(*ReadinessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthService_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServiceServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/guestcoviderpb.HealthService/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServiceServer).Version(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HealthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "guestcoviderpb.HealthService",
	HandlerType: (*HealthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Liveness",
			Handler:    _HealthService_Liveness_Handler,
		},
		{
			MethodName: "Readiness",
			Handler:    _HealthService_Readiness_Handler,
		},
		{
			MethodName: "Version",
			Handler:    _HealthService_Version_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "guestcovider-services.proto",
}

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	SearchUser(ctx context.Context, in *SearchUserRequest, opts ...grpc.CallOption) (*SearchUserResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) SearchUser(ctx context.Context, in *SearchUserRequest, opts ...grpc.CallOption) (*SearchUserResponse, error) {
	out := new(SearchUserResponse)
	err := c.cc.Invoke(ctx, "/guestcoviderpb.UserService/SearchUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	out := new(UpdateUserResponse)
	err := c.cc.Invoke(ctx, "/guestcoviderpb.UserService/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	SearchUser(context.Context, *SearchUserRequest) (*SearchUserResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
}

// UnimplementedUserServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (*UnimplementedUserServiceServer) SearchUser(context.Context, *SearchUserRequest) (*SearchUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchUser not implemented")
}
func (*UnimplementedUserServiceServer) UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_SearchUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SearchUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/guestcoviderpb.UserService/SearchUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SearchUser(ctx, req.(*SearchUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/guestcoviderpb.UserService/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "guestcoviderpb.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchUser",
			Handler:    _UserService_SearchUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserService_UpdateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "guestcovider-services.proto",
}
