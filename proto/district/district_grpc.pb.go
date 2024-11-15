// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: district/district.proto

package proto_district

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

// DistrictServiceClient is the client API for DistrictService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DistrictServiceClient interface {
	List(ctx context.Context, in *ListDistrictRequest, opts ...grpc.CallOption) (*ListDistrictResponse, error)
	Create(ctx context.Context, in *CreateDistrictRequest, opts ...grpc.CallOption) (*DistrictProfileResponse, error)
	Get(ctx context.Context, in *SingleDistrictRequest, opts ...grpc.CallOption) (*DistrictProfileResponse, error)
	Update(ctx context.Context, in *UpdateDistrictRequest, opts ...grpc.CallOption) (*DistrictSuccessResponse, error)
	Delete(ctx context.Context, in *SingleDistrictRequest, opts ...grpc.CallOption) (*DistrictSuccessResponse, error)
}

type districtServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDistrictServiceClient(cc grpc.ClientConnInterface) DistrictServiceClient {
	return &districtServiceClient{cc}
}

func (c *districtServiceClient) List(ctx context.Context, in *ListDistrictRequest, opts ...grpc.CallOption) (*ListDistrictResponse, error) {
	out := new(ListDistrictResponse)
	err := c.cc.Invoke(ctx, "/DistrictService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *districtServiceClient) Create(ctx context.Context, in *CreateDistrictRequest, opts ...grpc.CallOption) (*DistrictProfileResponse, error) {
	out := new(DistrictProfileResponse)
	err := c.cc.Invoke(ctx, "/DistrictService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *districtServiceClient) Get(ctx context.Context, in *SingleDistrictRequest, opts ...grpc.CallOption) (*DistrictProfileResponse, error) {
	out := new(DistrictProfileResponse)
	err := c.cc.Invoke(ctx, "/DistrictService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *districtServiceClient) Update(ctx context.Context, in *UpdateDistrictRequest, opts ...grpc.CallOption) (*DistrictSuccessResponse, error) {
	out := new(DistrictSuccessResponse)
	err := c.cc.Invoke(ctx, "/DistrictService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *districtServiceClient) Delete(ctx context.Context, in *SingleDistrictRequest, opts ...grpc.CallOption) (*DistrictSuccessResponse, error) {
	out := new(DistrictSuccessResponse)
	err := c.cc.Invoke(ctx, "/DistrictService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DistrictServiceServer is the server API for DistrictService service.
// All implementations must embed UnimplementedDistrictServiceServer
// for forward compatibility
type DistrictServiceServer interface {
	List(context.Context, *ListDistrictRequest) (*ListDistrictResponse, error)
	Create(context.Context, *CreateDistrictRequest) (*DistrictProfileResponse, error)
	Get(context.Context, *SingleDistrictRequest) (*DistrictProfileResponse, error)
	Update(context.Context, *UpdateDistrictRequest) (*DistrictSuccessResponse, error)
	Delete(context.Context, *SingleDistrictRequest) (*DistrictSuccessResponse, error)
	mustEmbedUnimplementedDistrictServiceServer()
}

// UnimplementedDistrictServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDistrictServiceServer struct {
}

func (UnimplementedDistrictServiceServer) List(context.Context, *ListDistrictRequest) (*ListDistrictResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedDistrictServiceServer) Create(context.Context, *CreateDistrictRequest) (*DistrictProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedDistrictServiceServer) Get(context.Context, *SingleDistrictRequest) (*DistrictProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedDistrictServiceServer) Update(context.Context, *UpdateDistrictRequest) (*DistrictSuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedDistrictServiceServer) Delete(context.Context, *SingleDistrictRequest) (*DistrictSuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedDistrictServiceServer) mustEmbedUnimplementedDistrictServiceServer() {}

// UnsafeDistrictServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DistrictServiceServer will
// result in compilation errors.
type UnsafeDistrictServiceServer interface {
	mustEmbedUnimplementedDistrictServiceServer()
}

func RegisterDistrictServiceServer(s grpc.ServiceRegistrar, srv DistrictServiceServer) {
	s.RegisterService(&DistrictService_ServiceDesc, srv)
}

func _DistrictService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDistrictRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistrictServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DistrictService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistrictServiceServer).List(ctx, req.(*ListDistrictRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DistrictService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDistrictRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistrictServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DistrictService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistrictServiceServer).Create(ctx, req.(*CreateDistrictRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DistrictService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SingleDistrictRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistrictServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DistrictService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistrictServiceServer).Get(ctx, req.(*SingleDistrictRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DistrictService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDistrictRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistrictServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DistrictService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistrictServiceServer).Update(ctx, req.(*UpdateDistrictRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DistrictService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SingleDistrictRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistrictServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DistrictService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistrictServiceServer).Delete(ctx, req.(*SingleDistrictRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DistrictService_ServiceDesc is the grpc.ServiceDesc for DistrictService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DistrictService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "DistrictService",
	HandlerType: (*DistrictServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _DistrictService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _DistrictService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _DistrictService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _DistrictService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _DistrictService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "district/district.proto",
}
