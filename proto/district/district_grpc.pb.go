// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: district.proto

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

const (
	DistrictService_List_FullMethodName   = "/DistrictService/List"
	DistrictService_Create_FullMethodName = "/DistrictService/Create"
	DistrictService_Get_FullMethodName    = "/DistrictService/Get"
	DistrictService_Update_FullMethodName = "/DistrictService/Update"
	DistrictService_Delete_FullMethodName = "/DistrictService/Delete"
)

// DistrictServiceClient is the client API for DistrictService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DistrictServiceClient interface {
	List(ctx context.Context, in *ListDistrictRequest, opts ...grpc.CallOption) (*ListDistrictResponse, error)
	Create(ctx context.Context, in *CreateDistrictRequest, opts ...grpc.CallOption) (*DistrictProfileResponse, error)
	Get(ctx context.Context, in *SingleDistrictRequest, opts ...grpc.CallOption) (*DistrictProfileResponse, error)
	Update(ctx context.Context, in *UpdateDistrictRequest, opts ...grpc.CallOption) (*SuccessResponse, error)
	Delete(ctx context.Context, in *SingleDistrictRequest, opts ...grpc.CallOption) (*SuccessResponse, error)
}

type districtServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDistrictServiceClient(cc grpc.ClientConnInterface) DistrictServiceClient {
	return &districtServiceClient{cc}
}

func (c *districtServiceClient) List(ctx context.Context, in *ListDistrictRequest, opts ...grpc.CallOption) (*ListDistrictResponse, error) {
	out := new(ListDistrictResponse)
	err := c.cc.Invoke(ctx, DistrictService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *districtServiceClient) Create(ctx context.Context, in *CreateDistrictRequest, opts ...grpc.CallOption) (*DistrictProfileResponse, error) {
	out := new(DistrictProfileResponse)
	err := c.cc.Invoke(ctx, DistrictService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *districtServiceClient) Get(ctx context.Context, in *SingleDistrictRequest, opts ...grpc.CallOption) (*DistrictProfileResponse, error) {
	out := new(DistrictProfileResponse)
	err := c.cc.Invoke(ctx, DistrictService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *districtServiceClient) Update(ctx context.Context, in *UpdateDistrictRequest, opts ...grpc.CallOption) (*SuccessResponse, error) {
	out := new(SuccessResponse)
	err := c.cc.Invoke(ctx, DistrictService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *districtServiceClient) Delete(ctx context.Context, in *SingleDistrictRequest, opts ...grpc.CallOption) (*SuccessResponse, error) {
	out := new(SuccessResponse)
	err := c.cc.Invoke(ctx, DistrictService_Delete_FullMethodName, in, out, opts...)
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
	Update(context.Context, *UpdateDistrictRequest) (*SuccessResponse, error)
	Delete(context.Context, *SingleDistrictRequest) (*SuccessResponse, error)
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
func (UnimplementedDistrictServiceServer) Update(context.Context, *UpdateDistrictRequest) (*SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedDistrictServiceServer) Delete(context.Context, *SingleDistrictRequest) (*SuccessResponse, error) {
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
		FullMethod: DistrictService_List_FullMethodName,
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
		FullMethod: DistrictService_Create_FullMethodName,
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
		FullMethod: DistrictService_Get_FullMethodName,
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
		FullMethod: DistrictService_Update_FullMethodName,
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
		FullMethod: DistrictService_Delete_FullMethodName,
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
	Metadata: "district.proto",
}
