// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: residency_type.proto

package proto_residency_type

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

const (
	ResidencyTypeService_List_FullMethodName   = "/ResidencyTypeService/List"
	ResidencyTypeService_Create_FullMethodName = "/ResidencyTypeService/Create"
	ResidencyTypeService_Get_FullMethodName    = "/ResidencyTypeService/Get"
	ResidencyTypeService_Update_FullMethodName = "/ResidencyTypeService/Update"
	ResidencyTypeService_Delete_FullMethodName = "/ResidencyTypeService/Delete"
)

// ResidencyTypeServiceClient is the client API for ResidencyTypeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResidencyTypeServiceClient interface {
	List(ctx context.Context, in *ListResidencyTypeRequest, opts ...grpc.CallOption) (*ListResidencyTypeResponse, error)
	Create(ctx context.Context, in *CreateResidencyTypeRequest, opts ...grpc.CallOption) (*ResidencyTypeProfileResponse, error)
	Get(ctx context.Context, in *SingleResidencyTypeRequest, opts ...grpc.CallOption) (*ResidencyTypeProfileResponse, error)
	Update(ctx context.Context, in *UpdateResidencyTypeRequest, opts ...grpc.CallOption) (*ResidencyTypeSuccessResponse, error)
	Delete(ctx context.Context, in *SingleResidencyTypeRequest, opts ...grpc.CallOption) (*ResidencyTypeSuccessResponse, error)
}

type residencyTypeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewResidencyTypeServiceClient(cc grpc.ClientConnInterface) ResidencyTypeServiceClient {
	return &residencyTypeServiceClient{cc}
}

func (c *residencyTypeServiceClient) List(ctx context.Context, in *ListResidencyTypeRequest, opts ...grpc.CallOption) (*ListResidencyTypeResponse, error) {
	out := new(ListResidencyTypeResponse)
	err := c.cc.Invoke(ctx, ResidencyTypeService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *residencyTypeServiceClient) Create(ctx context.Context, in *CreateResidencyTypeRequest, opts ...grpc.CallOption) (*ResidencyTypeProfileResponse, error) {
	out := new(ResidencyTypeProfileResponse)
	err := c.cc.Invoke(ctx, ResidencyTypeService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *residencyTypeServiceClient) Get(ctx context.Context, in *SingleResidencyTypeRequest, opts ...grpc.CallOption) (*ResidencyTypeProfileResponse, error) {
	out := new(ResidencyTypeProfileResponse)
	err := c.cc.Invoke(ctx, ResidencyTypeService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *residencyTypeServiceClient) Update(ctx context.Context, in *UpdateResidencyTypeRequest, opts ...grpc.CallOption) (*ResidencyTypeSuccessResponse, error) {
	out := new(ResidencyTypeSuccessResponse)
	err := c.cc.Invoke(ctx, ResidencyTypeService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *residencyTypeServiceClient) Delete(ctx context.Context, in *SingleResidencyTypeRequest, opts ...grpc.CallOption) (*ResidencyTypeSuccessResponse, error) {
	out := new(ResidencyTypeSuccessResponse)
	err := c.cc.Invoke(ctx, ResidencyTypeService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResidencyTypeServiceServer is the server API for ResidencyTypeService service.
// All implementations must embed UnimplementedResidencyTypeServiceServer
// for forward compatibility
type ResidencyTypeServiceServer interface {
	List(context.Context, *ListResidencyTypeRequest) (*ListResidencyTypeResponse, error)
	Create(context.Context, *CreateResidencyTypeRequest) (*ResidencyTypeProfileResponse, error)
	Get(context.Context, *SingleResidencyTypeRequest) (*ResidencyTypeProfileResponse, error)
	Update(context.Context, *UpdateResidencyTypeRequest) (*ResidencyTypeSuccessResponse, error)
	Delete(context.Context, *SingleResidencyTypeRequest) (*ResidencyTypeSuccessResponse, error)
	mustEmbedUnimplementedResidencyTypeServiceServer()
}

// UnimplementedResidencyTypeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedResidencyTypeServiceServer struct {
}

func (UnimplementedResidencyTypeServiceServer) List(context.Context, *ListResidencyTypeRequest) (*ListResidencyTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedResidencyTypeServiceServer) Create(context.Context, *CreateResidencyTypeRequest) (*ResidencyTypeProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedResidencyTypeServiceServer) Get(context.Context, *SingleResidencyTypeRequest) (*ResidencyTypeProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedResidencyTypeServiceServer) Update(context.Context, *UpdateResidencyTypeRequest) (*ResidencyTypeSuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedResidencyTypeServiceServer) Delete(context.Context, *SingleResidencyTypeRequest) (*ResidencyTypeSuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedResidencyTypeServiceServer) mustEmbedUnimplementedResidencyTypeServiceServer() {}

// UnsafeResidencyTypeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResidencyTypeServiceServer will
// result in compilation errors.
type UnsafeResidencyTypeServiceServer interface {
	mustEmbedUnimplementedResidencyTypeServiceServer()
}

func RegisterResidencyTypeServiceServer(s grpc.ServiceRegistrar, srv ResidencyTypeServiceServer) {
	s.RegisterService(&ResidencyTypeService_ServiceDesc, srv)
}

func _ResidencyTypeService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListResidencyTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResidencyTypeServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResidencyTypeService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResidencyTypeServiceServer).List(ctx, req.(*ListResidencyTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResidencyTypeService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateResidencyTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResidencyTypeServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResidencyTypeService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResidencyTypeServiceServer).Create(ctx, req.(*CreateResidencyTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResidencyTypeService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SingleResidencyTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResidencyTypeServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResidencyTypeService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResidencyTypeServiceServer).Get(ctx, req.(*SingleResidencyTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResidencyTypeService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateResidencyTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResidencyTypeServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResidencyTypeService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResidencyTypeServiceServer).Update(ctx, req.(*UpdateResidencyTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResidencyTypeService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SingleResidencyTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResidencyTypeServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResidencyTypeService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResidencyTypeServiceServer).Delete(ctx, req.(*SingleResidencyTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ResidencyTypeService_ServiceDesc is the grpc.ServiceDesc for ResidencyTypeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ResidencyTypeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ResidencyTypeService",
	HandlerType: (*ResidencyTypeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _ResidencyTypeService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _ResidencyTypeService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ResidencyTypeService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ResidencyTypeService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ResidencyTypeService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "residency_type.proto",
}
