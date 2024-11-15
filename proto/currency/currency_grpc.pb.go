// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: currency/currency.proto

package proto_currency

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

// CurrencyServiceClient is the client API for CurrencyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CurrencyServiceClient interface {
	List(ctx context.Context, in *ListCurrencyRequest, opts ...grpc.CallOption) (*ListCurrencyResponse, error)
	Create(ctx context.Context, in *CreateCurrencyRequest, opts ...grpc.CallOption) (*CurrencyProfileResponse, error)
	Get(ctx context.Context, in *SingleCurrencyRequest, opts ...grpc.CallOption) (*CurrencyProfileResponse, error)
	Update(ctx context.Context, in *UpdateCurrencyRequest, opts ...grpc.CallOption) (*CurrencySuccessResponse, error)
	Delete(ctx context.Context, in *SingleCurrencyRequest, opts ...grpc.CallOption) (*CurrencySuccessResponse, error)
}

type currencyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCurrencyServiceClient(cc grpc.ClientConnInterface) CurrencyServiceClient {
	return &currencyServiceClient{cc}
}

func (c *currencyServiceClient) List(ctx context.Context, in *ListCurrencyRequest, opts ...grpc.CallOption) (*ListCurrencyResponse, error) {
	out := new(ListCurrencyResponse)
	err := c.cc.Invoke(ctx, "/CurrencyService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currencyServiceClient) Create(ctx context.Context, in *CreateCurrencyRequest, opts ...grpc.CallOption) (*CurrencyProfileResponse, error) {
	out := new(CurrencyProfileResponse)
	err := c.cc.Invoke(ctx, "/CurrencyService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currencyServiceClient) Get(ctx context.Context, in *SingleCurrencyRequest, opts ...grpc.CallOption) (*CurrencyProfileResponse, error) {
	out := new(CurrencyProfileResponse)
	err := c.cc.Invoke(ctx, "/CurrencyService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currencyServiceClient) Update(ctx context.Context, in *UpdateCurrencyRequest, opts ...grpc.CallOption) (*CurrencySuccessResponse, error) {
	out := new(CurrencySuccessResponse)
	err := c.cc.Invoke(ctx, "/CurrencyService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currencyServiceClient) Delete(ctx context.Context, in *SingleCurrencyRequest, opts ...grpc.CallOption) (*CurrencySuccessResponse, error) {
	out := new(CurrencySuccessResponse)
	err := c.cc.Invoke(ctx, "/CurrencyService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CurrencyServiceServer is the server API for CurrencyService service.
// All implementations must embed UnimplementedCurrencyServiceServer
// for forward compatibility
type CurrencyServiceServer interface {
	List(context.Context, *ListCurrencyRequest) (*ListCurrencyResponse, error)
	Create(context.Context, *CreateCurrencyRequest) (*CurrencyProfileResponse, error)
	Get(context.Context, *SingleCurrencyRequest) (*CurrencyProfileResponse, error)
	Update(context.Context, *UpdateCurrencyRequest) (*CurrencySuccessResponse, error)
	Delete(context.Context, *SingleCurrencyRequest) (*CurrencySuccessResponse, error)
	mustEmbedUnimplementedCurrencyServiceServer()
}

// UnimplementedCurrencyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCurrencyServiceServer struct {
}

func (UnimplementedCurrencyServiceServer) List(context.Context, *ListCurrencyRequest) (*ListCurrencyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedCurrencyServiceServer) Create(context.Context, *CreateCurrencyRequest) (*CurrencyProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCurrencyServiceServer) Get(context.Context, *SingleCurrencyRequest) (*CurrencyProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedCurrencyServiceServer) Update(context.Context, *UpdateCurrencyRequest) (*CurrencySuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedCurrencyServiceServer) Delete(context.Context, *SingleCurrencyRequest) (*CurrencySuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCurrencyServiceServer) mustEmbedUnimplementedCurrencyServiceServer() {}

// UnsafeCurrencyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CurrencyServiceServer will
// result in compilation errors.
type UnsafeCurrencyServiceServer interface {
	mustEmbedUnimplementedCurrencyServiceServer()
}

func RegisterCurrencyServiceServer(s grpc.ServiceRegistrar, srv CurrencyServiceServer) {
	s.RegisterService(&CurrencyService_ServiceDesc, srv)
}

func _CurrencyService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCurrencyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrencyServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CurrencyService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrencyServiceServer).List(ctx, req.(*ListCurrencyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CurrencyService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCurrencyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrencyServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CurrencyService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrencyServiceServer).Create(ctx, req.(*CreateCurrencyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CurrencyService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SingleCurrencyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrencyServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CurrencyService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrencyServiceServer).Get(ctx, req.(*SingleCurrencyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CurrencyService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCurrencyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrencyServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CurrencyService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrencyServiceServer).Update(ctx, req.(*UpdateCurrencyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CurrencyService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SingleCurrencyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrencyServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CurrencyService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrencyServiceServer).Delete(ctx, req.(*SingleCurrencyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CurrencyService_ServiceDesc is the grpc.ServiceDesc for CurrencyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CurrencyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CurrencyService",
	HandlerType: (*CurrencyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _CurrencyService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _CurrencyService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _CurrencyService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _CurrencyService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CurrencyService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "currency/currency.proto",
}
