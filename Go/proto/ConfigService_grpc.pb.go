// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: proto/ConfigService.proto

package ConfigService

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ConfigServiceClient is the client API for ConfigService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConfigServiceClient interface {
	AddConfig(ctx context.Context, in *Config, opts ...grpc.CallOption) (*Config, error)
	GetConfig(ctx context.Context, in *ConfigNameRequest, opts ...grpc.CallOption) (*Config, error)
	GetAllVersionsOfConfig(ctx context.Context, in *ConfigNameRequest, opts ...grpc.CallOption) (*Configs, error)
	GetAllConfigs(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Configs, error)
	UpdateConfig(ctx context.Context, in *Config, opts ...grpc.CallOption) (*Config, error)
	DeleteConfig(ctx context.Context, in *ConfigNameRequest, opts ...grpc.CallOption) (*Config, error)
	UseConfig(ctx context.Context, in *ConfigNameRequest, opts ...grpc.CallOption) (ConfigService_UseConfigClient, error)
	StopConfigUseForAll(ctx context.Context, in *ConfigNameRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type configServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConfigServiceClient(cc grpc.ClientConnInterface) ConfigServiceClient {
	return &configServiceClient{cc}
}

func (c *configServiceClient) AddConfig(ctx context.Context, in *Config, opts ...grpc.CallOption) (*Config, error) {
	out := new(Config)
	err := c.cc.Invoke(ctx, "/ConfigService.ConfigService/addConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) GetConfig(ctx context.Context, in *ConfigNameRequest, opts ...grpc.CallOption) (*Config, error) {
	out := new(Config)
	err := c.cc.Invoke(ctx, "/ConfigService.ConfigService/getConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) GetAllVersionsOfConfig(ctx context.Context, in *ConfigNameRequest, opts ...grpc.CallOption) (*Configs, error) {
	out := new(Configs)
	err := c.cc.Invoke(ctx, "/ConfigService.ConfigService/getAllVersionsOfConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) GetAllConfigs(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Configs, error) {
	out := new(Configs)
	err := c.cc.Invoke(ctx, "/ConfigService.ConfigService/getAllConfigs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) UpdateConfig(ctx context.Context, in *Config, opts ...grpc.CallOption) (*Config, error) {
	out := new(Config)
	err := c.cc.Invoke(ctx, "/ConfigService.ConfigService/updateConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) DeleteConfig(ctx context.Context, in *ConfigNameRequest, opts ...grpc.CallOption) (*Config, error) {
	out := new(Config)
	err := c.cc.Invoke(ctx, "/ConfigService.ConfigService/deleteConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) UseConfig(ctx context.Context, in *ConfigNameRequest, opts ...grpc.CallOption) (ConfigService_UseConfigClient, error) {
	stream, err := c.cc.NewStream(ctx, &ConfigService_ServiceDesc.Streams[0], "/ConfigService.ConfigService/useConfig", opts...)
	if err != nil {
		return nil, err
	}
	x := &configServiceUseConfigClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ConfigService_UseConfigClient interface {
	Recv() (*Config, error)
	grpc.ClientStream
}

type configServiceUseConfigClient struct {
	grpc.ClientStream
}

func (x *configServiceUseConfigClient) Recv() (*Config, error) {
	m := new(Config)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *configServiceClient) StopConfigUseForAll(ctx context.Context, in *ConfigNameRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ConfigService.ConfigService/stopConfigUseForAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfigServiceServer is the server API for ConfigService service.
// All implementations must embed UnimplementedConfigServiceServer
// for forward compatibility
type ConfigServiceServer interface {
	AddConfig(context.Context, *Config) (*Config, error)
	GetConfig(context.Context, *ConfigNameRequest) (*Config, error)
	GetAllVersionsOfConfig(context.Context, *ConfigNameRequest) (*Configs, error)
	GetAllConfigs(context.Context, *emptypb.Empty) (*Configs, error)
	UpdateConfig(context.Context, *Config) (*Config, error)
	DeleteConfig(context.Context, *ConfigNameRequest) (*Config, error)
	UseConfig(*ConfigNameRequest, ConfigService_UseConfigServer) error
	StopConfigUseForAll(context.Context, *ConfigNameRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedConfigServiceServer()
}

// UnimplementedConfigServiceServer must be embedded to have forward compatible implementations.
type UnimplementedConfigServiceServer struct {
}

func (UnimplementedConfigServiceServer) AddConfig(context.Context, *Config) (*Config, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddConfig not implemented")
}
func (UnimplementedConfigServiceServer) GetConfig(context.Context, *ConfigNameRequest) (*Config, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfig not implemented")
}
func (UnimplementedConfigServiceServer) GetAllVersionsOfConfig(context.Context, *ConfigNameRequest) (*Configs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllVersionsOfConfig not implemented")
}
func (UnimplementedConfigServiceServer) GetAllConfigs(context.Context, *emptypb.Empty) (*Configs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllConfigs not implemented")
}
func (UnimplementedConfigServiceServer) UpdateConfig(context.Context, *Config) (*Config, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateConfig not implemented")
}
func (UnimplementedConfigServiceServer) DeleteConfig(context.Context, *ConfigNameRequest) (*Config, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteConfig not implemented")
}
func (UnimplementedConfigServiceServer) UseConfig(*ConfigNameRequest, ConfigService_UseConfigServer) error {
	return status.Errorf(codes.Unimplemented, "method UseConfig not implemented")
}
func (UnimplementedConfigServiceServer) StopConfigUseForAll(context.Context, *ConfigNameRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopConfigUseForAll not implemented")
}
func (UnimplementedConfigServiceServer) mustEmbedUnimplementedConfigServiceServer() {}

// UnsafeConfigServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConfigServiceServer will
// result in compilation errors.
type UnsafeConfigServiceServer interface {
	mustEmbedUnimplementedConfigServiceServer()
}

func RegisterConfigServiceServer(s grpc.ServiceRegistrar, srv ConfigServiceServer) {
	s.RegisterService(&ConfigService_ServiceDesc, srv)
}

func _ConfigService_AddConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Config)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).AddConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ConfigService.ConfigService/addConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).AddConfig(ctx, req.(*Config))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_GetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).GetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ConfigService.ConfigService/getConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).GetConfig(ctx, req.(*ConfigNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_GetAllVersionsOfConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).GetAllVersionsOfConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ConfigService.ConfigService/getAllVersionsOfConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).GetAllVersionsOfConfig(ctx, req.(*ConfigNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_GetAllConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).GetAllConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ConfigService.ConfigService/getAllConfigs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).GetAllConfigs(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_UpdateConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Config)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).UpdateConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ConfigService.ConfigService/updateConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).UpdateConfig(ctx, req.(*Config))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_DeleteConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).DeleteConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ConfigService.ConfigService/deleteConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).DeleteConfig(ctx, req.(*ConfigNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_UseConfig_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ConfigNameRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ConfigServiceServer).UseConfig(m, &configServiceUseConfigServer{stream})
}

type ConfigService_UseConfigServer interface {
	Send(*Config) error
	grpc.ServerStream
}

type configServiceUseConfigServer struct {
	grpc.ServerStream
}

func (x *configServiceUseConfigServer) Send(m *Config) error {
	return x.ServerStream.SendMsg(m)
}

func _ConfigService_StopConfigUseForAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).StopConfigUseForAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ConfigService.ConfigService/stopConfigUseForAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).StopConfigUseForAll(ctx, req.(*ConfigNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ConfigService_ServiceDesc is the grpc.ServiceDesc for ConfigService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConfigService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ConfigService.ConfigService",
	HandlerType: (*ConfigServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "addConfig",
			Handler:    _ConfigService_AddConfig_Handler,
		},
		{
			MethodName: "getConfig",
			Handler:    _ConfigService_GetConfig_Handler,
		},
		{
			MethodName: "getAllVersionsOfConfig",
			Handler:    _ConfigService_GetAllVersionsOfConfig_Handler,
		},
		{
			MethodName: "getAllConfigs",
			Handler:    _ConfigService_GetAllConfigs_Handler,
		},
		{
			MethodName: "updateConfig",
			Handler:    _ConfigService_UpdateConfig_Handler,
		},
		{
			MethodName: "deleteConfig",
			Handler:    _ConfigService_DeleteConfig_Handler,
		},
		{
			MethodName: "stopConfigUseForAll",
			Handler:    _ConfigService_StopConfigUseForAll_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "useConfig",
			Handler:       _ConfigService_UseConfig_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/ConfigService.proto",
}
