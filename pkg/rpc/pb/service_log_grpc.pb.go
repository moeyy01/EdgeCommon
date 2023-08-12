// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: service_log.proto

package pb

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
	LogService_CreateLog_FullMethodName             = "/pb.LogService/createLog"
	LogService_CountLogs_FullMethodName             = "/pb.LogService/countLogs"
	LogService_ListLogs_FullMethodName              = "/pb.LogService/listLogs"
	LogService_DeleteLogPermanently_FullMethodName  = "/pb.LogService/deleteLogPermanently"
	LogService_DeleteLogsPermanently_FullMethodName = "/pb.LogService/deleteLogsPermanently"
	LogService_CleanLogsPermanently_FullMethodName  = "/pb.LogService/cleanLogsPermanently"
	LogService_SumLogsSize_FullMethodName           = "/pb.LogService/sumLogsSize"
)

// LogServiceClient is the client API for LogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LogServiceClient interface {
	// 创建日志
	CreateLog(ctx context.Context, in *CreateLogRequest, opts ...grpc.CallOption) (*CreateLogResponse, error)
	// 计算日志数量
	CountLogs(ctx context.Context, in *CountLogRequest, opts ...grpc.CallOption) (*RPCCountResponse, error)
	// 列出单页日志
	ListLogs(ctx context.Context, in *ListLogsRequest, opts ...grpc.CallOption) (*ListLogsResponse, error)
	// 删除单条
	DeleteLogPermanently(ctx context.Context, in *DeleteLogPermanentlyRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 批量删除
	DeleteLogsPermanently(ctx context.Context, in *DeleteLogsPermanentlyRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 清理
	CleanLogsPermanently(ctx context.Context, in *CleanLogsPermanentlyRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 计算日志容量大小
	SumLogsSize(ctx context.Context, in *SumLogsSizeRequest, opts ...grpc.CallOption) (*SumLogsResponse, error)
}

type logServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLogServiceClient(cc grpc.ClientConnInterface) LogServiceClient {
	return &logServiceClient{cc}
}

func (c *logServiceClient) CreateLog(ctx context.Context, in *CreateLogRequest, opts ...grpc.CallOption) (*CreateLogResponse, error) {
	out := new(CreateLogResponse)
	err := c.cc.Invoke(ctx, LogService_CreateLog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logServiceClient) CountLogs(ctx context.Context, in *CountLogRequest, opts ...grpc.CallOption) (*RPCCountResponse, error) {
	out := new(RPCCountResponse)
	err := c.cc.Invoke(ctx, LogService_CountLogs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logServiceClient) ListLogs(ctx context.Context, in *ListLogsRequest, opts ...grpc.CallOption) (*ListLogsResponse, error) {
	out := new(ListLogsResponse)
	err := c.cc.Invoke(ctx, LogService_ListLogs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logServiceClient) DeleteLogPermanently(ctx context.Context, in *DeleteLogPermanentlyRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, LogService_DeleteLogPermanently_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logServiceClient) DeleteLogsPermanently(ctx context.Context, in *DeleteLogsPermanentlyRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, LogService_DeleteLogsPermanently_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logServiceClient) CleanLogsPermanently(ctx context.Context, in *CleanLogsPermanentlyRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, LogService_CleanLogsPermanently_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logServiceClient) SumLogsSize(ctx context.Context, in *SumLogsSizeRequest, opts ...grpc.CallOption) (*SumLogsResponse, error) {
	out := new(SumLogsResponse)
	err := c.cc.Invoke(ctx, LogService_SumLogsSize_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogServiceServer is the server API for LogService service.
// All implementations should embed UnimplementedLogServiceServer
// for forward compatibility
type LogServiceServer interface {
	// 创建日志
	CreateLog(context.Context, *CreateLogRequest) (*CreateLogResponse, error)
	// 计算日志数量
	CountLogs(context.Context, *CountLogRequest) (*RPCCountResponse, error)
	// 列出单页日志
	ListLogs(context.Context, *ListLogsRequest) (*ListLogsResponse, error)
	// 删除单条
	DeleteLogPermanently(context.Context, *DeleteLogPermanentlyRequest) (*RPCSuccess, error)
	// 批量删除
	DeleteLogsPermanently(context.Context, *DeleteLogsPermanentlyRequest) (*RPCSuccess, error)
	// 清理
	CleanLogsPermanently(context.Context, *CleanLogsPermanentlyRequest) (*RPCSuccess, error)
	// 计算日志容量大小
	SumLogsSize(context.Context, *SumLogsSizeRequest) (*SumLogsResponse, error)
}

// UnimplementedLogServiceServer should be embedded to have forward compatible implementations.
type UnimplementedLogServiceServer struct {
}

func (UnimplementedLogServiceServer) CreateLog(context.Context, *CreateLogRequest) (*CreateLogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLog not implemented")
}
func (UnimplementedLogServiceServer) CountLogs(context.Context, *CountLogRequest) (*RPCCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountLogs not implemented")
}
func (UnimplementedLogServiceServer) ListLogs(context.Context, *ListLogsRequest) (*ListLogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLogs not implemented")
}
func (UnimplementedLogServiceServer) DeleteLogPermanently(context.Context, *DeleteLogPermanentlyRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLogPermanently not implemented")
}
func (UnimplementedLogServiceServer) DeleteLogsPermanently(context.Context, *DeleteLogsPermanentlyRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLogsPermanently not implemented")
}
func (UnimplementedLogServiceServer) CleanLogsPermanently(context.Context, *CleanLogsPermanentlyRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CleanLogsPermanently not implemented")
}
func (UnimplementedLogServiceServer) SumLogsSize(context.Context, *SumLogsSizeRequest) (*SumLogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SumLogsSize not implemented")
}

// UnsafeLogServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogServiceServer will
// result in compilation errors.
type UnsafeLogServiceServer interface {
	mustEmbedUnimplementedLogServiceServer()
}

func RegisterLogServiceServer(s grpc.ServiceRegistrar, srv LogServiceServer) {
	s.RegisterService(&LogService_ServiceDesc, srv)
}

func _LogService_CreateLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServiceServer).CreateLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LogService_CreateLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServiceServer).CreateLog(ctx, req.(*CreateLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogService_CountLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServiceServer).CountLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LogService_CountLogs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServiceServer).CountLogs(ctx, req.(*CountLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogService_ListLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServiceServer).ListLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LogService_ListLogs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServiceServer).ListLogs(ctx, req.(*ListLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogService_DeleteLogPermanently_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLogPermanentlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServiceServer).DeleteLogPermanently(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LogService_DeleteLogPermanently_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServiceServer).DeleteLogPermanently(ctx, req.(*DeleteLogPermanentlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogService_DeleteLogsPermanently_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLogsPermanentlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServiceServer).DeleteLogsPermanently(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LogService_DeleteLogsPermanently_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServiceServer).DeleteLogsPermanently(ctx, req.(*DeleteLogsPermanentlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogService_CleanLogsPermanently_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CleanLogsPermanentlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServiceServer).CleanLogsPermanently(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LogService_CleanLogsPermanently_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServiceServer).CleanLogsPermanently(ctx, req.(*CleanLogsPermanentlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogService_SumLogsSize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumLogsSizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServiceServer).SumLogsSize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LogService_SumLogsSize_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServiceServer).SumLogsSize(ctx, req.(*SumLogsSizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LogService_ServiceDesc is the grpc.ServiceDesc for LogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.LogService",
	HandlerType: (*LogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createLog",
			Handler:    _LogService_CreateLog_Handler,
		},
		{
			MethodName: "countLogs",
			Handler:    _LogService_CountLogs_Handler,
		},
		{
			MethodName: "listLogs",
			Handler:    _LogService_ListLogs_Handler,
		},
		{
			MethodName: "deleteLogPermanently",
			Handler:    _LogService_DeleteLogPermanently_Handler,
		},
		{
			MethodName: "deleteLogsPermanently",
			Handler:    _LogService_DeleteLogsPermanently_Handler,
		},
		{
			MethodName: "cleanLogsPermanently",
			Handler:    _LogService_CleanLogsPermanently_Handler,
		},
		{
			MethodName: "sumLogsSize",
			Handler:    _LogService_SumLogsSize_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_log.proto",
}