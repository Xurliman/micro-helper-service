// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: national_economy_sector_old.proto

package proto_national_economy_sector_old

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

// NationalEconomySectorOldServiceClient is the client API for NationalEconomySectorOldService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NationalEconomySectorOldServiceClient interface {
	Create(ctx context.Context, in *CreateNationalEconomySectorOldRequest, opts ...grpc.CallOption) (*NationalEconomySectorOldProfileResponse, error)
	Read(ctx context.Context, in *SingleNationalEconomySectorOldRequest, opts ...grpc.CallOption) (*NationalEconomySectorOldProfileResponse, error)
	Update(ctx context.Context, in *UpdateNationalEconomySectorOldRequest, opts ...grpc.CallOption) (*SuccessResponse, error)
	Delete(ctx context.Context, in *SingleNationalEconomySectorOldRequest, opts ...grpc.CallOption) (*SuccessResponse, error)
}

type nationalEconomySectorOldServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNationalEconomySectorOldServiceClient(cc grpc.ClientConnInterface) NationalEconomySectorOldServiceClient {
	return &nationalEconomySectorOldServiceClient{cc}
}

func (c *nationalEconomySectorOldServiceClient) Create(ctx context.Context, in *CreateNationalEconomySectorOldRequest, opts ...grpc.CallOption) (*NationalEconomySectorOldProfileResponse, error) {
	out := new(NationalEconomySectorOldProfileResponse)
	err := c.cc.Invoke(ctx, "/NationalEconomySectorOldService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nationalEconomySectorOldServiceClient) Read(ctx context.Context, in *SingleNationalEconomySectorOldRequest, opts ...grpc.CallOption) (*NationalEconomySectorOldProfileResponse, error) {
	out := new(NationalEconomySectorOldProfileResponse)
	err := c.cc.Invoke(ctx, "/NationalEconomySectorOldService/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nationalEconomySectorOldServiceClient) Update(ctx context.Context, in *UpdateNationalEconomySectorOldRequest, opts ...grpc.CallOption) (*SuccessResponse, error) {
	out := new(SuccessResponse)
	err := c.cc.Invoke(ctx, "/NationalEconomySectorOldService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nationalEconomySectorOldServiceClient) Delete(ctx context.Context, in *SingleNationalEconomySectorOldRequest, opts ...grpc.CallOption) (*SuccessResponse, error) {
	out := new(SuccessResponse)
	err := c.cc.Invoke(ctx, "/NationalEconomySectorOldService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NationalEconomySectorOldServiceServer is the server API for NationalEconomySectorOldService service.
// All implementations must embed UnimplementedNationalEconomySectorOldServiceServer
// for forward compatibility
type NationalEconomySectorOldServiceServer interface {
	Create(context.Context, *CreateNationalEconomySectorOldRequest) (*NationalEconomySectorOldProfileResponse, error)
	Read(context.Context, *SingleNationalEconomySectorOldRequest) (*NationalEconomySectorOldProfileResponse, error)
	Update(context.Context, *UpdateNationalEconomySectorOldRequest) (*SuccessResponse, error)
	Delete(context.Context, *SingleNationalEconomySectorOldRequest) (*SuccessResponse, error)
	mustEmbedUnimplementedNationalEconomySectorOldServiceServer()
}

// UnimplementedNationalEconomySectorOldServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNationalEconomySectorOldServiceServer struct {
}

func (UnimplementedNationalEconomySectorOldServiceServer) Create(context.Context, *CreateNationalEconomySectorOldRequest) (*NationalEconomySectorOldProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedNationalEconomySectorOldServiceServer) Read(context.Context, *SingleNationalEconomySectorOldRequest) (*NationalEconomySectorOldProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedNationalEconomySectorOldServiceServer) Update(context.Context, *UpdateNationalEconomySectorOldRequest) (*SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedNationalEconomySectorOldServiceServer) Delete(context.Context, *SingleNationalEconomySectorOldRequest) (*SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedNationalEconomySectorOldServiceServer) mustEmbedUnimplementedNationalEconomySectorOldServiceServer() {
}

// UnsafeNationalEconomySectorOldServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NationalEconomySectorOldServiceServer will
// result in compilation errors.
type UnsafeNationalEconomySectorOldServiceServer interface {
	mustEmbedUnimplementedNationalEconomySectorOldServiceServer()
}

func RegisterNationalEconomySectorOldServiceServer(s grpc.ServiceRegistrar, srv NationalEconomySectorOldServiceServer) {
	s.RegisterService(&NationalEconomySectorOldService_ServiceDesc, srv)
}

func _NationalEconomySectorOldService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNationalEconomySectorOldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NationalEconomySectorOldServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NationalEconomySectorOldService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NationalEconomySectorOldServiceServer).Create(ctx, req.(*CreateNationalEconomySectorOldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NationalEconomySectorOldService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SingleNationalEconomySectorOldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NationalEconomySectorOldServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NationalEconomySectorOldService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NationalEconomySectorOldServiceServer).Read(ctx, req.(*SingleNationalEconomySectorOldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NationalEconomySectorOldService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNationalEconomySectorOldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NationalEconomySectorOldServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NationalEconomySectorOldService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NationalEconomySectorOldServiceServer).Update(ctx, req.(*UpdateNationalEconomySectorOldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NationalEconomySectorOldService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SingleNationalEconomySectorOldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NationalEconomySectorOldServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NationalEconomySectorOldService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NationalEconomySectorOldServiceServer).Delete(ctx, req.(*SingleNationalEconomySectorOldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NationalEconomySectorOldService_ServiceDesc is the grpc.ServiceDesc for NationalEconomySectorOldService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NationalEconomySectorOldService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "NationalEconomySectorOldService",
	HandlerType: (*NationalEconomySectorOldServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _NationalEconomySectorOldService_Create_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _NationalEconomySectorOldService_Read_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _NationalEconomySectorOldService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _NationalEconomySectorOldService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "national_economy_sector_old.proto",
}