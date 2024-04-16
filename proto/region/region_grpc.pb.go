// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: region.proto

package proto_region

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
	RegionService_Create_FullMethodName = "/RegionService/Create"
	RegionService_Read_FullMethodName   = "/RegionService/Read"
	RegionService_Update_FullMethodName = "/RegionService/Update"
	RegionService_Delete_FullMethodName = "/RegionService/Delete"
)

// RegionServiceClient is the client API for RegionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RegionServiceClient interface {
	Create(ctx context.Context, in *CreateRegionRequest, opts ...grpc.CallOption) (*RegionProfileResponse, error)
	Read(ctx context.Context, in *SingleRegionRequest, opts ...grpc.CallOption) (*RegionProfileResponse, error)
	Update(ctx context.Context, in *UpdateRegionRequest, opts ...grpc.CallOption) (*SuccessResponse, error)
	Delete(ctx context.Context, in *SingleRegionRequest, opts ...grpc.CallOption) (*SuccessResponse, error)
}

type regionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRegionServiceClient(cc grpc.ClientConnInterface) RegionServiceClient {
	return &regionServiceClient{cc}
}

func (c *regionServiceClient) Create(ctx context.Context, in *CreateRegionRequest, opts ...grpc.CallOption) (*RegionProfileResponse, error) {
	out := new(RegionProfileResponse)
	err := c.cc.Invoke(ctx, RegionService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *regionServiceClient) Read(ctx context.Context, in *SingleRegionRequest, opts ...grpc.CallOption) (*RegionProfileResponse, error) {
	out := new(RegionProfileResponse)
	err := c.cc.Invoke(ctx, RegionService_Read_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *regionServiceClient) Update(ctx context.Context, in *UpdateRegionRequest, opts ...grpc.CallOption) (*SuccessResponse, error) {
	out := new(SuccessResponse)
	err := c.cc.Invoke(ctx, RegionService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *regionServiceClient) Delete(ctx context.Context, in *SingleRegionRequest, opts ...grpc.CallOption) (*SuccessResponse, error) {
	out := new(SuccessResponse)
	err := c.cc.Invoke(ctx, RegionService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegionServiceServer is the server API for RegionService service.
// All implementations must embed UnimplementedRegionServiceServer
// for forward compatibility
type RegionServiceServer interface {
	Create(context.Context, *CreateRegionRequest) (*RegionProfileResponse, error)
	Read(context.Context, *SingleRegionRequest) (*RegionProfileResponse, error)
	Update(context.Context, *UpdateRegionRequest) (*SuccessResponse, error)
	Delete(context.Context, *SingleRegionRequest) (*SuccessResponse, error)
	mustEmbedUnimplementedRegionServiceServer()
}

// UnimplementedRegionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRegionServiceServer struct {
}

func (UnimplementedRegionServiceServer) Create(context.Context, *CreateRegionRequest) (*RegionProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedRegionServiceServer) Read(context.Context, *SingleRegionRequest) (*RegionProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedRegionServiceServer) Update(context.Context, *UpdateRegionRequest) (*SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedRegionServiceServer) Delete(context.Context, *SingleRegionRequest) (*SuccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedRegionServiceServer) mustEmbedUnimplementedRegionServiceServer() {}

// UnsafeRegionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RegionServiceServer will
// result in compilation errors.
type UnsafeRegionServiceServer interface {
	mustEmbedUnimplementedRegionServiceServer()
}

func RegisterRegionServiceServer(s grpc.ServiceRegistrar, srv RegionServiceServer) {
	s.RegisterService(&RegionService_ServiceDesc, srv)
}

func _RegionService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRegionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegionServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RegionService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegionServiceServer).Create(ctx, req.(*CreateRegionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegionService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SingleRegionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegionServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RegionService_Read_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegionServiceServer).Read(ctx, req.(*SingleRegionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegionService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRegionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegionServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RegionService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegionServiceServer).Update(ctx, req.(*UpdateRegionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegionService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SingleRegionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegionServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RegionService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegionServiceServer).Delete(ctx, req.(*SingleRegionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RegionService_ServiceDesc is the grpc.ServiceDesc for RegionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RegionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "RegionService",
	HandlerType: (*RegionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _RegionService_Create_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _RegionService_Read_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _RegionService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _RegionService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "region.proto",
}
