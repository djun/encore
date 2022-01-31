// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: encore/daemon/daemon.proto

package daemon

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

// DaemonClient is the client API for Daemon service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DaemonClient interface {
	// Run runs the application.
	Run(ctx context.Context, in *RunRequest, opts ...grpc.CallOption) (Daemon_RunClient, error)
	// Test runs tests.
	Test(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (Daemon_TestClient, error)
	// Check checks the app for compilation errors.
	Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (Daemon_CheckClient, error)
	// Logs streams logs from the encore.dev platform.
	Logs(ctx context.Context, in *LogsRequest, opts ...grpc.CallOption) (Daemon_LogsClient, error)
	// DBConnect starts the database and returns the DSN for connecting to it.
	DBConnect(ctx context.Context, in *DBConnectRequest, opts ...grpc.CallOption) (*DBConnectResponse, error)
	// DBProxy starts a local database proxy for connecting to remote databases
	// on the encore.dev platform.
	DBProxy(ctx context.Context, in *DBProxyRequest, opts ...grpc.CallOption) (Daemon_DBProxyClient, error)
	// DBReset resets the given databases, recreating them from scratch.
	DBReset(ctx context.Context, in *DBResetRequest, opts ...grpc.CallOption) (Daemon_DBResetClient, error)
	// GenClient generates a client based on the app's API.
	GenClient(ctx context.Context, in *GenClientRequest, opts ...grpc.CallOption) (*GenClientResponse, error)
	// SetSecret sets a secret key on the encore.dev platform.
	SetSecret(ctx context.Context, in *SetSecretRequest, opts ...grpc.CallOption) (*SetSecretResponse, error)
	// Version reports the daemon version.
	Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*VersionResponse, error)
}

type daemonClient struct {
	cc grpc.ClientConnInterface
}

func NewDaemonClient(cc grpc.ClientConnInterface) DaemonClient {
	return &daemonClient{cc}
}

func (c *daemonClient) Run(ctx context.Context, in *RunRequest, opts ...grpc.CallOption) (Daemon_RunClient, error) {
	stream, err := c.cc.NewStream(ctx, &Daemon_ServiceDesc.Streams[0], "/encore.daemon.Daemon/Run", opts...)
	if err != nil {
		return nil, err
	}
	x := &daemonRunClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Daemon_RunClient interface {
	Recv() (*CommandMessage, error)
	grpc.ClientStream
}

type daemonRunClient struct {
	grpc.ClientStream
}

func (x *daemonRunClient) Recv() (*CommandMessage, error) {
	m := new(CommandMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *daemonClient) Test(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (Daemon_TestClient, error) {
	stream, err := c.cc.NewStream(ctx, &Daemon_ServiceDesc.Streams[1], "/encore.daemon.Daemon/Test", opts...)
	if err != nil {
		return nil, err
	}
	x := &daemonTestClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Daemon_TestClient interface {
	Recv() (*CommandMessage, error)
	grpc.ClientStream
}

type daemonTestClient struct {
	grpc.ClientStream
}

func (x *daemonTestClient) Recv() (*CommandMessage, error) {
	m := new(CommandMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *daemonClient) Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (Daemon_CheckClient, error) {
	stream, err := c.cc.NewStream(ctx, &Daemon_ServiceDesc.Streams[2], "/encore.daemon.Daemon/Check", opts...)
	if err != nil {
		return nil, err
	}
	x := &daemonCheckClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Daemon_CheckClient interface {
	Recv() (*CommandMessage, error)
	grpc.ClientStream
}

type daemonCheckClient struct {
	grpc.ClientStream
}

func (x *daemonCheckClient) Recv() (*CommandMessage, error) {
	m := new(CommandMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *daemonClient) Logs(ctx context.Context, in *LogsRequest, opts ...grpc.CallOption) (Daemon_LogsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Daemon_ServiceDesc.Streams[3], "/encore.daemon.Daemon/Logs", opts...)
	if err != nil {
		return nil, err
	}
	x := &daemonLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Daemon_LogsClient interface {
	Recv() (*LogsMessage, error)
	grpc.ClientStream
}

type daemonLogsClient struct {
	grpc.ClientStream
}

func (x *daemonLogsClient) Recv() (*LogsMessage, error) {
	m := new(LogsMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *daemonClient) DBConnect(ctx context.Context, in *DBConnectRequest, opts ...grpc.CallOption) (*DBConnectResponse, error) {
	out := new(DBConnectResponse)
	err := c.cc.Invoke(ctx, "/encore.daemon.Daemon/DBConnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonClient) DBProxy(ctx context.Context, in *DBProxyRequest, opts ...grpc.CallOption) (Daemon_DBProxyClient, error) {
	stream, err := c.cc.NewStream(ctx, &Daemon_ServiceDesc.Streams[4], "/encore.daemon.Daemon/DBProxy", opts...)
	if err != nil {
		return nil, err
	}
	x := &daemonDBProxyClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Daemon_DBProxyClient interface {
	Recv() (*CommandMessage, error)
	grpc.ClientStream
}

type daemonDBProxyClient struct {
	grpc.ClientStream
}

func (x *daemonDBProxyClient) Recv() (*CommandMessage, error) {
	m := new(CommandMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *daemonClient) DBReset(ctx context.Context, in *DBResetRequest, opts ...grpc.CallOption) (Daemon_DBResetClient, error) {
	stream, err := c.cc.NewStream(ctx, &Daemon_ServiceDesc.Streams[5], "/encore.daemon.Daemon/DBReset", opts...)
	if err != nil {
		return nil, err
	}
	x := &daemonDBResetClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Daemon_DBResetClient interface {
	Recv() (*CommandMessage, error)
	grpc.ClientStream
}

type daemonDBResetClient struct {
	grpc.ClientStream
}

func (x *daemonDBResetClient) Recv() (*CommandMessage, error) {
	m := new(CommandMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *daemonClient) GenClient(ctx context.Context, in *GenClientRequest, opts ...grpc.CallOption) (*GenClientResponse, error) {
	out := new(GenClientResponse)
	err := c.cc.Invoke(ctx, "/encore.daemon.Daemon/GenClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonClient) SetSecret(ctx context.Context, in *SetSecretRequest, opts ...grpc.CallOption) (*SetSecretResponse, error) {
	out := new(SetSecretResponse)
	err := c.cc.Invoke(ctx, "/encore.daemon.Daemon/SetSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonClient) Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/encore.daemon.Daemon/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DaemonServer is the server API for Daemon service.
// All implementations must embed UnimplementedDaemonServer
// for forward compatibility
type DaemonServer interface {
	// Run runs the application.
	Run(*RunRequest, Daemon_RunServer) error
	// Test runs tests.
	Test(*TestRequest, Daemon_TestServer) error
	// Check checks the app for compilation errors.
	Check(*CheckRequest, Daemon_CheckServer) error
	// Logs streams logs from the encore.dev platform.
	Logs(*LogsRequest, Daemon_LogsServer) error
	// DBConnect starts the database and returns the DSN for connecting to it.
	DBConnect(context.Context, *DBConnectRequest) (*DBConnectResponse, error)
	// DBProxy starts a local database proxy for connecting to remote databases
	// on the encore.dev platform.
	DBProxy(*DBProxyRequest, Daemon_DBProxyServer) error
	// DBReset resets the given databases, recreating them from scratch.
	DBReset(*DBResetRequest, Daemon_DBResetServer) error
	// GenClient generates a client based on the app's API.
	GenClient(context.Context, *GenClientRequest) (*GenClientResponse, error)
	// SetSecret sets a secret key on the encore.dev platform.
	SetSecret(context.Context, *SetSecretRequest) (*SetSecretResponse, error)
	// Version reports the daemon version.
	Version(context.Context, *emptypb.Empty) (*VersionResponse, error)
	mustEmbedUnimplementedDaemonServer()
}

// UnimplementedDaemonServer must be embedded to have forward compatible implementations.
type UnimplementedDaemonServer struct {
}

func (UnimplementedDaemonServer) Run(*RunRequest, Daemon_RunServer) error {
	return status.Errorf(codes.Unimplemented, "method Run not implemented")
}
func (UnimplementedDaemonServer) Test(*TestRequest, Daemon_TestServer) error {
	return status.Errorf(codes.Unimplemented, "method Test not implemented")
}
func (UnimplementedDaemonServer) Check(*CheckRequest, Daemon_CheckServer) error {
	return status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (UnimplementedDaemonServer) Logs(*LogsRequest, Daemon_LogsServer) error {
	return status.Errorf(codes.Unimplemented, "method Logs not implemented")
}
func (UnimplementedDaemonServer) DBConnect(context.Context, *DBConnectRequest) (*DBConnectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DBConnect not implemented")
}
func (UnimplementedDaemonServer) DBProxy(*DBProxyRequest, Daemon_DBProxyServer) error {
	return status.Errorf(codes.Unimplemented, "method DBProxy not implemented")
}
func (UnimplementedDaemonServer) DBReset(*DBResetRequest, Daemon_DBResetServer) error {
	return status.Errorf(codes.Unimplemented, "method DBReset not implemented")
}
func (UnimplementedDaemonServer) GenClient(context.Context, *GenClientRequest) (*GenClientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenClient not implemented")
}
func (UnimplementedDaemonServer) SetSecret(context.Context, *SetSecretRequest) (*SetSecretResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetSecret not implemented")
}
func (UnimplementedDaemonServer) Version(context.Context, *emptypb.Empty) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedDaemonServer) mustEmbedUnimplementedDaemonServer() {}

// UnsafeDaemonServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DaemonServer will
// result in compilation errors.
type UnsafeDaemonServer interface {
	mustEmbedUnimplementedDaemonServer()
}

func RegisterDaemonServer(s grpc.ServiceRegistrar, srv DaemonServer) {
	s.RegisterService(&Daemon_ServiceDesc, srv)
}

func _Daemon_Run_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RunRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DaemonServer).Run(m, &daemonRunServer{stream})
}

type Daemon_RunServer interface {
	Send(*CommandMessage) error
	grpc.ServerStream
}

type daemonRunServer struct {
	grpc.ServerStream
}

func (x *daemonRunServer) Send(m *CommandMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _Daemon_Test_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TestRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DaemonServer).Test(m, &daemonTestServer{stream})
}

type Daemon_TestServer interface {
	Send(*CommandMessage) error
	grpc.ServerStream
}

type daemonTestServer struct {
	grpc.ServerStream
}

func (x *daemonTestServer) Send(m *CommandMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _Daemon_Check_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CheckRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DaemonServer).Check(m, &daemonCheckServer{stream})
}

type Daemon_CheckServer interface {
	Send(*CommandMessage) error
	grpc.ServerStream
}

type daemonCheckServer struct {
	grpc.ServerStream
}

func (x *daemonCheckServer) Send(m *CommandMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _Daemon_Logs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(LogsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DaemonServer).Logs(m, &daemonLogsServer{stream})
}

type Daemon_LogsServer interface {
	Send(*LogsMessage) error
	grpc.ServerStream
}

type daemonLogsServer struct {
	grpc.ServerStream
}

func (x *daemonLogsServer) Send(m *LogsMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _Daemon_DBConnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DBConnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).DBConnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/encore.daemon.Daemon/DBConnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).DBConnect(ctx, req.(*DBConnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Daemon_DBProxy_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DBProxyRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DaemonServer).DBProxy(m, &daemonDBProxyServer{stream})
}

type Daemon_DBProxyServer interface {
	Send(*CommandMessage) error
	grpc.ServerStream
}

type daemonDBProxyServer struct {
	grpc.ServerStream
}

func (x *daemonDBProxyServer) Send(m *CommandMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _Daemon_DBReset_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DBResetRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DaemonServer).DBReset(m, &daemonDBResetServer{stream})
}

type Daemon_DBResetServer interface {
	Send(*CommandMessage) error
	grpc.ServerStream
}

type daemonDBResetServer struct {
	grpc.ServerStream
}

func (x *daemonDBResetServer) Send(m *CommandMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _Daemon_GenClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).GenClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/encore.daemon.Daemon/GenClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).GenClient(ctx, req.(*GenClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Daemon_SetSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).SetSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/encore.daemon.Daemon/SetSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).SetSecret(ctx, req.(*SetSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Daemon_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/encore.daemon.Daemon/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).Version(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Daemon_ServiceDesc is the grpc.ServiceDesc for Daemon service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Daemon_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "encore.daemon.Daemon",
	HandlerType: (*DaemonServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DBConnect",
			Handler:    _Daemon_DBConnect_Handler,
		},
		{
			MethodName: "GenClient",
			Handler:    _Daemon_GenClient_Handler,
		},
		{
			MethodName: "SetSecret",
			Handler:    _Daemon_SetSecret_Handler,
		},
		{
			MethodName: "Version",
			Handler:    _Daemon_Version_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Run",
			Handler:       _Daemon_Run_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Test",
			Handler:       _Daemon_Test_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Check",
			Handler:       _Daemon_Check_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Logs",
			Handler:       _Daemon_Logs_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "DBProxy",
			Handler:       _Daemon_DBProxy_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "DBReset",
			Handler:       _Daemon_DBReset_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "encore/daemon/daemon.proto",
}
