// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package config

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

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	SaveConfig(ctx context.Context, in *Config, opts ...grpc.CallOption) (*Config, error)
	QueryConfig(ctx context.Context, in *QueryConfigRequest, opts ...grpc.CallOption) (*ConfigSet, error)
	UpdateConfig(ctx context.Context, in *UpdateConfigRequest, opts ...grpc.CallOption) (*Config, error)
	DescribeConfig(ctx context.Context, in *DescribeConfigRequest, opts ...grpc.CallOption) (*Config, error)
	DeleteConfig(ctx context.Context, in *DeleteConfigRequest, opts ...grpc.CallOption) (*Config, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) SaveConfig(ctx context.Context, in *Config, opts ...grpc.CallOption) (*Config, error) {
	out := new(Config)
	err := c.cc.Invoke(ctx, "/ericyaoxr.cmdb.config.Service/SaveConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) QueryConfig(ctx context.Context, in *QueryConfigRequest, opts ...grpc.CallOption) (*ConfigSet, error) {
	out := new(ConfigSet)
	err := c.cc.Invoke(ctx, "/ericyaoxr.cmdb.config.Service/QueryConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) UpdateConfig(ctx context.Context, in *UpdateConfigRequest, opts ...grpc.CallOption) (*Config, error) {
	out := new(Config)
	err := c.cc.Invoke(ctx, "/ericyaoxr.cmdb.config.Service/UpdateConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DescribeConfig(ctx context.Context, in *DescribeConfigRequest, opts ...grpc.CallOption) (*Config, error) {
	out := new(Config)
	err := c.cc.Invoke(ctx, "/ericyaoxr.cmdb.config.Service/DescribeConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) DeleteConfig(ctx context.Context, in *DeleteConfigRequest, opts ...grpc.CallOption) (*Config, error) {
	out := new(Config)
	err := c.cc.Invoke(ctx, "/ericyaoxr.cmdb.config.Service/DeleteConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	SaveConfig(context.Context, *Config) (*Config, error)
	QueryConfig(context.Context, *QueryConfigRequest) (*ConfigSet, error)
	UpdateConfig(context.Context, *UpdateConfigRequest) (*Config, error)
	DescribeConfig(context.Context, *DescribeConfigRequest) (*Config, error)
	DeleteConfig(context.Context, *DeleteConfigRequest) (*Config, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) SaveConfig(context.Context, *Config) (*Config, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveConfig not implemented")
}
func (UnimplementedServiceServer) QueryConfig(context.Context, *QueryConfigRequest) (*ConfigSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryConfig not implemented")
}
func (UnimplementedServiceServer) UpdateConfig(context.Context, *UpdateConfigRequest) (*Config, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateConfig not implemented")
}
func (UnimplementedServiceServer) DescribeConfig(context.Context, *DescribeConfigRequest) (*Config, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeConfig not implemented")
}
func (UnimplementedServiceServer) DeleteConfig(context.Context, *DeleteConfigRequest) (*Config, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteConfig not implemented")
}
func (UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_SaveConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Config)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).SaveConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ericyaoxr.cmdb.config.Service/SaveConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).SaveConfig(ctx, req.(*Config))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_QueryConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).QueryConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ericyaoxr.cmdb.config.Service/QueryConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).QueryConfig(ctx, req.(*QueryConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_UpdateConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).UpdateConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ericyaoxr.cmdb.config.Service/UpdateConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).UpdateConfig(ctx, req.(*UpdateConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_DescribeConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DescribeConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ericyaoxr.cmdb.config.Service/DescribeConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DescribeConfig(ctx, req.(*DescribeConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_DeleteConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).DeleteConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ericyaoxr.cmdb.config.Service/DeleteConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).DeleteConfig(ctx, req.(*DeleteConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ericyaoxr.cmdb.config.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveConfig",
			Handler:    _Service_SaveConfig_Handler,
		},
		{
			MethodName: "QueryConfig",
			Handler:    _Service_QueryConfig_Handler,
		},
		{
			MethodName: "UpdateConfig",
			Handler:    _Service_UpdateConfig_Handler,
		},
		{
			MethodName: "DescribeConfig",
			Handler:    _Service_DescribeConfig_Handler,
		},
		{
			MethodName: "DeleteConfig",
			Handler:    _Service_DeleteConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/config/pb/config.proto",
}