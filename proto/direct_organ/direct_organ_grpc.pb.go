// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: direct_organ.proto

package proto_bank

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

// DirectOrganServiceClient is the client API for DirectOrganService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DirectOrganServiceClient interface {
	Create(ctx context.Context, in *CreateDirectOrganRequest, opts ...grpc.CallOption) (*DirectOrganProfileResponse, error)
	Read(ctx context.Context, in *SingleDirectOrganRequest, opts ...grpc.CallOption) (*DirectOrganProfileResponse, error)
	Update(ctx context.Context, in *UpdateDirectOrganRequest, opts ...grpc.CallOption) (*SuccessResponse, error)
	Delete(ctx context.Context, in *SingleDirectOrganRequest, opts ...grpc.CallOption) (*SuccessResponse, error)
}

type directOrganServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDirectOrganServiceClient(cc grpc.ClientConnInterface) DirectOrganServiceClient {
	return &directOrganServiceClient{cc}
}

func (c *directOrganServiceClient) Create(ctx context.Context, in *CreateDirectOrganRequest, opts ...grpc.CallOption) (*DirectOrganProfileResponse, error) {
	out := new(DirectOrganProfileResponse)
	err := c.cc.Invoke(ctx, "/DirectOrganService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directOrganServiceClient) Read(ctx context.Context, in *SingleDirectOrganRequest, opts ...grpc.CallOption) (*DirectOrganProfileResponse, error) {
	out := new(DirectOrganProfileResponse)
	err := c.cc.Invoke(ctx, "/DirectOrganService/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directOrganServiceClient) Update(ctx context.Context, in *UpdateDirectOrganRequest, opts ...grpc.CallOption) (*SuccessResponse, error) {
	out := new(SuccessResponse)
	err := c.cc.Invoke(ctx, "/DirectOrganService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directOrganServiceClient) Delete(ctx context.Context, in *SingleDirectOrganRequest, opts ...grpc.CallOption) (*SuccessResponse, error) {
	out := new(SuccessResponse)
	err := c.cc.Invoke(ctx, "/DirectOrganService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DirectOrganServiceServer is the server API for DirectOrganService service.
// All implementations must embed UnimplementedDirectOrganServiceServer
// for forward compatibility
type DirectOrganServiceServer interface {
	Create(context.Context, *CreateDirectOrganRequest) (*DirectOrganProfileResponse, error)
	Read(context.Context, *SingleDirectOrganRequest) (*DirectOrganProfileResponse, error)
	Update(context.Context, *UpdateDirectOrganRequest) (*SuccessResponse, error)
	Delete(context.Context, *SingleDirectOrganRequest) (*SuccessResponse, error)
	mustEmbedUnimplementedDirectOrganServiceServer()
}

// UnimplementedDirectOrganServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDirectOrganServiceServer struct {
}

func (UnimplementedDirectOrganServiceServer) Create(context.Context, *CreateDirectOrganRequest) (*DirectOrganProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedDirectOrganServiceServer) Read(context.Context, *SingleDirectOrganRequest) (*DirectOrganProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedDirectOrganServiceServer) Update(context.Context, *UpdateDirectOrganRequest) (*SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedDirectOrganServiceServer) Delete(context.Context, *SingleDirectOrganRequest) (*SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedDirectOrganServiceServer) mustEmbedUnimplementedDirectOrganServiceServer() {}

// UnsafeDirectOrganServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DirectOrganServiceServer will
// result in compilation errors.
type UnsafeDirectOrganServiceServer interface {
	mustEmbedUnimplementedDirectOrganServiceServer()
}

func RegisterDirectOrganServiceServer(s grpc.ServiceRegistrar, srv DirectOrganServiceServer) {
	s.RegisterService(&DirectOrganService_ServiceDesc, srv)
}

func _DirectOrganService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDirectOrganRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectOrganServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DirectOrganService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectOrganServiceServer).Create(ctx, req.(*CreateDirectOrganRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirectOrganService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SingleDirectOrganRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectOrganServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DirectOrganService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectOrganServiceServer).Read(ctx, req.(*SingleDirectOrganRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirectOrganService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDirectOrganRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectOrganServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DirectOrganService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectOrganServiceServer).Update(ctx, req.(*UpdateDirectOrganRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirectOrganService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SingleDirectOrganRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectOrganServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DirectOrganService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectOrganServiceServer).Delete(ctx, req.(*SingleDirectOrganRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DirectOrganService_ServiceDesc is the grpc.ServiceDesc for DirectOrganService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DirectOrganService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "DirectOrganService",
	HandlerType: (*DirectOrganServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _DirectOrganService_Create_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _DirectOrganService_Read_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _DirectOrganService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _DirectOrganService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "direct_organ.proto",
}
