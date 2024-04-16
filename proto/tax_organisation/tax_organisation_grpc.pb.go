// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: tax_organisation.proto

package proto_tax_organisation

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
	TaxOrganisationService_List_FullMethodName   = "/TaxOrganisationService/List"
	TaxOrganisationService_Create_FullMethodName = "/TaxOrganisationService/Create"
	TaxOrganisationService_Get_FullMethodName    = "/TaxOrganisationService/Get"
	TaxOrganisationService_Update_FullMethodName = "/TaxOrganisationService/Update"
	TaxOrganisationService_Delete_FullMethodName = "/TaxOrganisationService/Delete"
)

// TaxOrganisationServiceClient is the client API for TaxOrganisationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaxOrganisationServiceClient interface {
	List(ctx context.Context, in *ListTaxOrganisationRequest, opts ...grpc.CallOption) (*ListTaxOrganisationResponse, error)
	Create(ctx context.Context, in *CreateTaxOrganisationRequest, opts ...grpc.CallOption) (*TaxOrganisationProfileResponse, error)
	Get(ctx context.Context, in *SingleTaxOrganisationRequest, opts ...grpc.CallOption) (*TaxOrganisationProfileResponse, error)
	Update(ctx context.Context, in *UpdateTaxOrganisationRequest, opts ...grpc.CallOption) (*SuccessResponse, error)
	Delete(ctx context.Context, in *SingleTaxOrganisationRequest, opts ...grpc.CallOption) (*SuccessResponse, error)
}

type taxOrganisationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTaxOrganisationServiceClient(cc grpc.ClientConnInterface) TaxOrganisationServiceClient {
	return &taxOrganisationServiceClient{cc}
}

func (c *taxOrganisationServiceClient) List(ctx context.Context, in *ListTaxOrganisationRequest, opts ...grpc.CallOption) (*ListTaxOrganisationResponse, error) {
	out := new(ListTaxOrganisationResponse)
	err := c.cc.Invoke(ctx, TaxOrganisationService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taxOrganisationServiceClient) Create(ctx context.Context, in *CreateTaxOrganisationRequest, opts ...grpc.CallOption) (*TaxOrganisationProfileResponse, error) {
	out := new(TaxOrganisationProfileResponse)
	err := c.cc.Invoke(ctx, TaxOrganisationService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taxOrganisationServiceClient) Get(ctx context.Context, in *SingleTaxOrganisationRequest, opts ...grpc.CallOption) (*TaxOrganisationProfileResponse, error) {
	out := new(TaxOrganisationProfileResponse)
	err := c.cc.Invoke(ctx, TaxOrganisationService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taxOrganisationServiceClient) Update(ctx context.Context, in *UpdateTaxOrganisationRequest, opts ...grpc.CallOption) (*SuccessResponse, error) {
	out := new(SuccessResponse)
	err := c.cc.Invoke(ctx, TaxOrganisationService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taxOrganisationServiceClient) Delete(ctx context.Context, in *SingleTaxOrganisationRequest, opts ...grpc.CallOption) (*SuccessResponse, error) {
	out := new(SuccessResponse)
	err := c.cc.Invoke(ctx, TaxOrganisationService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaxOrganisationServiceServer is the server API for TaxOrganisationService service.
// All implementations must embed UnimplementedTaxOrganisationServiceServer
// for forward compatibility
type TaxOrganisationServiceServer interface {
	List(context.Context, *ListTaxOrganisationRequest) (*ListTaxOrganisationResponse, error)
	Create(context.Context, *CreateTaxOrganisationRequest) (*TaxOrganisationProfileResponse, error)
	Get(context.Context, *SingleTaxOrganisationRequest) (*TaxOrganisationProfileResponse, error)
	Update(context.Context, *UpdateTaxOrganisationRequest) (*SuccessResponse, error)
	Delete(context.Context, *SingleTaxOrganisationRequest) (*SuccessResponse, error)
	mustEmbedUnimplementedTaxOrganisationServiceServer()
}

// UnimplementedTaxOrganisationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTaxOrganisationServiceServer struct {
}

func (UnimplementedTaxOrganisationServiceServer) List(context.Context, *ListTaxOrganisationRequest) (*ListTaxOrganisationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedTaxOrganisationServiceServer) Create(context.Context, *CreateTaxOrganisationRequest) (*TaxOrganisationProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedTaxOrganisationServiceServer) Get(context.Context, *SingleTaxOrganisationRequest) (*TaxOrganisationProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedTaxOrganisationServiceServer) Update(context.Context, *UpdateTaxOrganisationRequest) (*SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedTaxOrganisationServiceServer) Delete(context.Context, *SingleTaxOrganisationRequest) (*SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedTaxOrganisationServiceServer) mustEmbedUnimplementedTaxOrganisationServiceServer() {
}

// UnsafeTaxOrganisationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaxOrganisationServiceServer will
// result in compilation errors.
type UnsafeTaxOrganisationServiceServer interface {
	mustEmbedUnimplementedTaxOrganisationServiceServer()
}

func RegisterTaxOrganisationServiceServer(s grpc.ServiceRegistrar, srv TaxOrganisationServiceServer) {
	s.RegisterService(&TaxOrganisationService_ServiceDesc, srv)
}

func _TaxOrganisationService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTaxOrganisationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaxOrganisationServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaxOrganisationService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaxOrganisationServiceServer).List(ctx, req.(*ListTaxOrganisationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaxOrganisationService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaxOrganisationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaxOrganisationServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaxOrganisationService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaxOrganisationServiceServer).Create(ctx, req.(*CreateTaxOrganisationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaxOrganisationService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SingleTaxOrganisationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaxOrganisationServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaxOrganisationService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaxOrganisationServiceServer).Get(ctx, req.(*SingleTaxOrganisationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaxOrganisationService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTaxOrganisationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaxOrganisationServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaxOrganisationService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaxOrganisationServiceServer).Update(ctx, req.(*UpdateTaxOrganisationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaxOrganisationService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SingleTaxOrganisationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaxOrganisationServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaxOrganisationService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaxOrganisationServiceServer).Delete(ctx, req.(*SingleTaxOrganisationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TaxOrganisationService_ServiceDesc is the grpc.ServiceDesc for TaxOrganisationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TaxOrganisationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TaxOrganisationService",
	HandlerType: (*TaxOrganisationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _TaxOrganisationService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _TaxOrganisationService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _TaxOrganisationService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _TaxOrganisationService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _TaxOrganisationService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tax_organisation.proto",
}
