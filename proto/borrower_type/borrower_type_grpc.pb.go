// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: borrower_type.proto

package proto_borrower_type

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

// BorrowerTypeServiceClient is the client API for BorrowerTypeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BorrowerTypeServiceClient interface {
	List(ctx context.Context, in *ListBorrowerTypeRequest, opts ...grpc.CallOption) (*ListBorrowerTypeResponse, error)
	Create(ctx context.Context, in *CreateBorrowerTypeRequest, opts ...grpc.CallOption) (*BorrowerTypeResponse, error)
	Get(ctx context.Context, in *SingleBorrowerTypeRequest, opts ...grpc.CallOption) (*BorrowerTypeResponse, error)
	Update(ctx context.Context, in *SingleBorrowerTypeRequest, opts ...grpc.CallOption) (*SuccessResponse, error)
	Delete(ctx context.Context, in *SingleBorrowerTypeRequest, opts ...grpc.CallOption) (*SuccessResponse, error)
}

type borrowerTypeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBorrowerTypeServiceClient(cc grpc.ClientConnInterface) BorrowerTypeServiceClient {
	return &borrowerTypeServiceClient{cc}
}

func (c *borrowerTypeServiceClient) List(ctx context.Context, in *ListBorrowerTypeRequest, opts ...grpc.CallOption) (*ListBorrowerTypeResponse, error) {
	out := new(ListBorrowerTypeResponse)
	err := c.cc.Invoke(ctx, "/BorrowerTypeService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *borrowerTypeServiceClient) Create(ctx context.Context, in *CreateBorrowerTypeRequest, opts ...grpc.CallOption) (*BorrowerTypeResponse, error) {
	out := new(BorrowerTypeResponse)
	err := c.cc.Invoke(ctx, "/BorrowerTypeService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *borrowerTypeServiceClient) Get(ctx context.Context, in *SingleBorrowerTypeRequest, opts ...grpc.CallOption) (*BorrowerTypeResponse, error) {
	out := new(BorrowerTypeResponse)
	err := c.cc.Invoke(ctx, "/BorrowerTypeService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *borrowerTypeServiceClient) Update(ctx context.Context, in *SingleBorrowerTypeRequest, opts ...grpc.CallOption) (*SuccessResponse, error) {
	out := new(SuccessResponse)
	err := c.cc.Invoke(ctx, "/BorrowerTypeService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *borrowerTypeServiceClient) Delete(ctx context.Context, in *SingleBorrowerTypeRequest, opts ...grpc.CallOption) (*SuccessResponse, error) {
	out := new(SuccessResponse)
	err := c.cc.Invoke(ctx, "/BorrowerTypeService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BorrowerTypeServiceServer is the server API for BorrowerTypeService service.
// All implementations must embed UnimplementedBorrowerTypeServiceServer
// for forward compatibility
type BorrowerTypeServiceServer interface {
	List(context.Context, *ListBorrowerTypeRequest) (*ListBorrowerTypeResponse, error)
	Create(context.Context, *CreateBorrowerTypeRequest) (*BorrowerTypeResponse, error)
	Get(context.Context, *SingleBorrowerTypeRequest) (*BorrowerTypeResponse, error)
	Update(context.Context, *SingleBorrowerTypeRequest) (*SuccessResponse, error)
	Delete(context.Context, *SingleBorrowerTypeRequest) (*SuccessResponse, error)
	mustEmbedUnimplementedBorrowerTypeServiceServer()
}

// UnimplementedBorrowerTypeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBorrowerTypeServiceServer struct {
}

func (UnimplementedBorrowerTypeServiceServer) List(context.Context, *ListBorrowerTypeRequest) (*ListBorrowerTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedBorrowerTypeServiceServer) Create(context.Context, *CreateBorrowerTypeRequest) (*BorrowerTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedBorrowerTypeServiceServer) Get(context.Context, *SingleBorrowerTypeRequest) (*BorrowerTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedBorrowerTypeServiceServer) Update(context.Context, *SingleBorrowerTypeRequest) (*SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedBorrowerTypeServiceServer) Delete(context.Context, *SingleBorrowerTypeRequest) (*SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedBorrowerTypeServiceServer) mustEmbedUnimplementedBorrowerTypeServiceServer() {}

// UnsafeBorrowerTypeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BorrowerTypeServiceServer will
// result in compilation errors.
type UnsafeBorrowerTypeServiceServer interface {
	mustEmbedUnimplementedBorrowerTypeServiceServer()
}

func RegisterBorrowerTypeServiceServer(s grpc.ServiceRegistrar, srv BorrowerTypeServiceServer) {
	s.RegisterService(&BorrowerTypeService_ServiceDesc, srv)
}

func _BorrowerTypeService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBorrowerTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BorrowerTypeServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BorrowerTypeService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BorrowerTypeServiceServer).List(ctx, req.(*ListBorrowerTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BorrowerTypeService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBorrowerTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BorrowerTypeServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BorrowerTypeService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BorrowerTypeServiceServer).Create(ctx, req.(*CreateBorrowerTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BorrowerTypeService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SingleBorrowerTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BorrowerTypeServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BorrowerTypeService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BorrowerTypeServiceServer).Get(ctx, req.(*SingleBorrowerTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BorrowerTypeService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SingleBorrowerTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BorrowerTypeServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BorrowerTypeService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BorrowerTypeServiceServer).Update(ctx, req.(*SingleBorrowerTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BorrowerTypeService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SingleBorrowerTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BorrowerTypeServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BorrowerTypeService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BorrowerTypeServiceServer).Delete(ctx, req.(*SingleBorrowerTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BorrowerTypeService_ServiceDesc is the grpc.ServiceDesc for BorrowerTypeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BorrowerTypeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "BorrowerTypeService",
	HandlerType: (*BorrowerTypeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _BorrowerTypeService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _BorrowerTypeService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _BorrowerTypeService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _BorrowerTypeService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _BorrowerTypeService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "borrower_type.proto",
}