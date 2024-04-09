// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: national_economy_sector_new.proto

package proto_national_economy_sector_new

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// NationalEconomySectorNewServiceClient is the client API for NationalEconomySectorNewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NationalEconomySectorNewServiceClient interface {
	Create(ctx context.Context, in *CreateNationalEconomySectorNewRequest, opts ...grpc.CallOption) (*NationalEconomySectorNewProfileResponse, error)
	Read(ctx context.Context, in *SingleNationalEconomySectorNewRequest, opts ...grpc.CallOption) (*NationalEconomySectorNewProfileResponse, error)
	Update(ctx context.Context, in *UpdateNationalEconomySectorNewRequest, opts ...grpc.CallOption) (*SuccessResponse, error)
	Delete(ctx context.Context, in *SingleNationalEconomySectorNewRequest, opts ...grpc.CallOption) (*SuccessResponse, error)
}

type nationalEconomySectorNewServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNationalEconomySectorNewServiceClient(cc grpc.ClientConnInterface) NationalEconomySectorNewServiceClient {
	return &nationalEconomySectorNewServiceClient{cc}
}

func (c *nationalEconomySectorNewServiceClient) Create(ctx context.Context, in *CreateNationalEconomySectorNewRequest, opts ...grpc.CallOption) (*NationalEconomySectorNewProfileResponse, error) {
	out := new(NationalEconomySectorNewProfileResponse)
	err := c.cc.Invoke(ctx, "/NationalEconomySectorNewService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nationalEconomySectorNewServiceClient) Read(ctx context.Context, in *SingleNationalEconomySectorNewRequest, opts ...grpc.CallOption) (*NationalEconomySectorNewProfileResponse, error) {
	out := new(NationalEconomySectorNewProfileResponse)
	err := c.cc.Invoke(ctx, "/NationalEconomySectorNewService/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nationalEconomySectorNewServiceClient) Update(ctx context.Context, in *UpdateNationalEconomySectorNewRequest, opts ...grpc.CallOption) (*SuccessResponse, error) {
	out := new(SuccessResponse)
	err := c.cc.Invoke(ctx, "/NationalEconomySectorNewService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nationalEconomySectorNewServiceClient) Delete(ctx context.Context, in *SingleNationalEconomySectorNewRequest, opts ...grpc.CallOption) (*SuccessResponse, error) {
	out := new(SuccessResponse)
	err := c.cc.Invoke(ctx, "/NationalEconomySectorNewService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NationalEconomySectorNewServiceServer is the server API for NationalEconomySectorNewService service.
// All implementations must embed UnimplementedNationalEconomySectorNewServiceServer
// for forward compatibility
type NationalEconomySectorNewServiceServer interface {
	Create(context.Context, *CreateNationalEconomySectorNewRequest) (*NationalEconomySectorNewProfileResponse, error)
	Read(context.Context, *SingleNationalEconomySectorNewRequest) (*NationalEconomySectorNewProfileResponse, error)
	Update(context.Context, *UpdateNationalEconomySectorNewRequest) (*SuccessResponse, error)
	Delete(context.Context, *SingleNationalEconomySectorNewRequest) (*SuccessResponse, error)
	mustEmbedUnimplementedNationalEconomySectorNewServiceServer()
}

// UnimplementedNationalEconomySectorNewServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNationalEconomySectorNewServiceServer struct {
}

func (UnimplementedNationalEconomySectorNewServiceServer) Create(context.Context, *CreateNationalEconomySectorNewRequest) (*NationalEconomySectorNewProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedNationalEconomySectorNewServiceServer) Read(context.Context, *SingleNationalEconomySectorNewRequest) (*NationalEconomySectorNewProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedNationalEconomySectorNewServiceServer) Update(context.Context, *UpdateNationalEconomySectorNewRequest) (*SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedNationalEconomySectorNewServiceServer) Delete(context.Context, *SingleNationalEconomySectorNewRequest) (*SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedNationalEconomySectorNewServiceServer) mustEmbedUnimplementedNationalEconomySectorNewServiceServer() {
}

// UnsafeNationalEconomySectorNewServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NationalEconomySectorNewServiceServer will
// result in compilation errors.
type UnsafeNationalEconomySectorNewServiceServer interface {
	mustEmbedUnimplementedNationalEconomySectorNewServiceServer()
}

func RegisterNationalEconomySectorNewServiceServer(s grpc.ServiceRegistrar, srv NationalEconomySectorNewServiceServer) {
	s.RegisterService(&NationalEconomySectorNewService_ServiceDesc, srv)
}

func _NationalEconomySectorNewService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNationalEconomySectorNewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NationalEconomySectorNewServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NationalEconomySectorNewService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NationalEconomySectorNewServiceServer).Create(ctx, req.(*CreateNationalEconomySectorNewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NationalEconomySectorNewService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SingleNationalEconomySectorNewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NationalEconomySectorNewServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NationalEconomySectorNewService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NationalEconomySectorNewServiceServer).Read(ctx, req.(*SingleNationalEconomySectorNewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NationalEconomySectorNewService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNationalEconomySectorNewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NationalEconomySectorNewServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NationalEconomySectorNewService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NationalEconomySectorNewServiceServer).Update(ctx, req.(*UpdateNationalEconomySectorNewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NationalEconomySectorNewService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SingleNationalEconomySectorNewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NationalEconomySectorNewServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NationalEconomySectorNewService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NationalEconomySectorNewServiceServer).Delete(ctx, req.(*SingleNationalEconomySectorNewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NationalEconomySectorNewService_ServiceDesc is the grpc.ServiceDesc for NationalEconomySectorNewService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NationalEconomySectorNewService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "NationalEconomySectorNewService",
	HandlerType: (*NationalEconomySectorNewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _NationalEconomySectorNewService_Create_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _NationalEconomySectorNewService_Read_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _NationalEconomySectorNewService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _NationalEconomySectorNewService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "national_economy_sector_new.proto",
}