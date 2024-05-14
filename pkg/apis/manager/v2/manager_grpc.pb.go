// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: pkg/apis/manager/v2/manager.proto

package manager

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

// ManagerClient is the client API for Manager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ManagerClient interface {
	// Get SeedPeer and SeedPeer cluster configuration.
	GetSeedPeer(ctx context.Context, in *GetSeedPeerRequest, opts ...grpc.CallOption) (*SeedPeer, error)
	// Update SeedPeer configuration.
	UpdateSeedPeer(ctx context.Context, in *UpdateSeedPeerRequest, opts ...grpc.CallOption) (*SeedPeer, error)
	// Delete SeedPeer configuration.
	DeleteSeedPeer(ctx context.Context, in *DeleteSeedPeerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Get Scheduler and Scheduler cluster configuration.
	GetScheduler(ctx context.Context, in *GetSchedulerRequest, opts ...grpc.CallOption) (*Scheduler, error)
	// Update scheduler configuration.
	UpdateScheduler(ctx context.Context, in *UpdateSchedulerRequest, opts ...grpc.CallOption) (*Scheduler, error)
	// List acitve schedulers configuration.
	ListSchedulers(ctx context.Context, in *ListSchedulersRequest, opts ...grpc.CallOption) (*ListSchedulersResponse, error)
	// Get ObjectStorage configuration.
	GetObjectStorage(ctx context.Context, in *GetObjectStorageRequest, opts ...grpc.CallOption) (*ObjectStorage, error)
	// List buckets configuration.
	ListBuckets(ctx context.Context, in *ListBucketsRequest, opts ...grpc.CallOption) (*ListBucketsResponse, error)
	// List applications configuration.
	ListApplications(ctx context.Context, in *ListApplicationsRequest, opts ...grpc.CallOption) (*ListApplicationsResponse, error)
	// Create model and update data of model to object storage.
	CreateModel(ctx context.Context, in *CreateModelRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// KeepAlive with manager.
	KeepAlive(ctx context.Context, opts ...grpc.CallOption) (Manager_KeepAliveClient, error)
}

type managerClient struct {
	cc grpc.ClientConnInterface
}

func NewManagerClient(cc grpc.ClientConnInterface) ManagerClient {
	return &managerClient{cc}
}

func (c *managerClient) GetSeedPeer(ctx context.Context, in *GetSeedPeerRequest, opts ...grpc.CallOption) (*SeedPeer, error) {
	out := new(SeedPeer)
	err := c.cc.Invoke(ctx, "/manager.v2.Manager/GetSeedPeer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) UpdateSeedPeer(ctx context.Context, in *UpdateSeedPeerRequest, opts ...grpc.CallOption) (*SeedPeer, error) {
	out := new(SeedPeer)
	err := c.cc.Invoke(ctx, "/manager.v2.Manager/UpdateSeedPeer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) DeleteSeedPeer(ctx context.Context, in *DeleteSeedPeerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/manager.v2.Manager/DeleteSeedPeer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetScheduler(ctx context.Context, in *GetSchedulerRequest, opts ...grpc.CallOption) (*Scheduler, error) {
	out := new(Scheduler)
	err := c.cc.Invoke(ctx, "/manager.v2.Manager/GetScheduler", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) UpdateScheduler(ctx context.Context, in *UpdateSchedulerRequest, opts ...grpc.CallOption) (*Scheduler, error) {
	out := new(Scheduler)
	err := c.cc.Invoke(ctx, "/manager.v2.Manager/UpdateScheduler", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) ListSchedulers(ctx context.Context, in *ListSchedulersRequest, opts ...grpc.CallOption) (*ListSchedulersResponse, error) {
	out := new(ListSchedulersResponse)
	err := c.cc.Invoke(ctx, "/manager.v2.Manager/ListSchedulers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) GetObjectStorage(ctx context.Context, in *GetObjectStorageRequest, opts ...grpc.CallOption) (*ObjectStorage, error) {
	out := new(ObjectStorage)
	err := c.cc.Invoke(ctx, "/manager.v2.Manager/GetObjectStorage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) ListBuckets(ctx context.Context, in *ListBucketsRequest, opts ...grpc.CallOption) (*ListBucketsResponse, error) {
	out := new(ListBucketsResponse)
	err := c.cc.Invoke(ctx, "/manager.v2.Manager/ListBuckets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) ListApplications(ctx context.Context, in *ListApplicationsRequest, opts ...grpc.CallOption) (*ListApplicationsResponse, error) {
	out := new(ListApplicationsResponse)
	err := c.cc.Invoke(ctx, "/manager.v2.Manager/ListApplications", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) CreateModel(ctx context.Context, in *CreateModelRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/manager.v2.Manager/CreateModel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) KeepAlive(ctx context.Context, opts ...grpc.CallOption) (Manager_KeepAliveClient, error) {
	stream, err := c.cc.NewStream(ctx, &Manager_ServiceDesc.Streams[0], "/manager.v2.Manager/KeepAlive", opts...)
	if err != nil {
		return nil, err
	}
	x := &managerKeepAliveClient{stream}
	return x, nil
}

type Manager_KeepAliveClient interface {
	Send(*KeepAliveRequest) error
	CloseAndRecv() (*emptypb.Empty, error)
	grpc.ClientStream
}

type managerKeepAliveClient struct {
	grpc.ClientStream
}

func (x *managerKeepAliveClient) Send(m *KeepAliveRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *managerKeepAliveClient) CloseAndRecv() (*emptypb.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(emptypb.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ManagerServer is the server API for Manager service.
// All implementations should embed UnimplementedManagerServer
// for forward compatibility
type ManagerServer interface {
	// Get SeedPeer and SeedPeer cluster configuration.
	GetSeedPeer(context.Context, *GetSeedPeerRequest) (*SeedPeer, error)
	// Update SeedPeer configuration.
	UpdateSeedPeer(context.Context, *UpdateSeedPeerRequest) (*SeedPeer, error)
	// Delete SeedPeer configuration.
	DeleteSeedPeer(context.Context, *DeleteSeedPeerRequest) (*emptypb.Empty, error)
	// Get Scheduler and Scheduler cluster configuration.
	GetScheduler(context.Context, *GetSchedulerRequest) (*Scheduler, error)
	// Update scheduler configuration.
	UpdateScheduler(context.Context, *UpdateSchedulerRequest) (*Scheduler, error)
	// List acitve schedulers configuration.
	ListSchedulers(context.Context, *ListSchedulersRequest) (*ListSchedulersResponse, error)
	// Get ObjectStorage configuration.
	GetObjectStorage(context.Context, *GetObjectStorageRequest) (*ObjectStorage, error)
	// List buckets configuration.
	ListBuckets(context.Context, *ListBucketsRequest) (*ListBucketsResponse, error)
	// List applications configuration.
	ListApplications(context.Context, *ListApplicationsRequest) (*ListApplicationsResponse, error)
	// Create model and update data of model to object storage.
	CreateModel(context.Context, *CreateModelRequest) (*emptypb.Empty, error)
	// KeepAlive with manager.
	KeepAlive(Manager_KeepAliveServer) error
}

// UnimplementedManagerServer should be embedded to have forward compatible implementations.
type UnimplementedManagerServer struct {
}

func (UnimplementedManagerServer) GetSeedPeer(context.Context, *GetSeedPeerRequest) (*SeedPeer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSeedPeer not implemented")
}
func (UnimplementedManagerServer) UpdateSeedPeer(context.Context, *UpdateSeedPeerRequest) (*SeedPeer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSeedPeer not implemented")
}
func (UnimplementedManagerServer) DeleteSeedPeer(context.Context, *DeleteSeedPeerRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSeedPeer not implemented")
}
func (UnimplementedManagerServer) GetScheduler(context.Context, *GetSchedulerRequest) (*Scheduler, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetScheduler not implemented")
}
func (UnimplementedManagerServer) UpdateScheduler(context.Context, *UpdateSchedulerRequest) (*Scheduler, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateScheduler not implemented")
}
func (UnimplementedManagerServer) ListSchedulers(context.Context, *ListSchedulersRequest) (*ListSchedulersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSchedulers not implemented")
}
func (UnimplementedManagerServer) GetObjectStorage(context.Context, *GetObjectStorageRequest) (*ObjectStorage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObjectStorage not implemented")
}
func (UnimplementedManagerServer) ListBuckets(context.Context, *ListBucketsRequest) (*ListBucketsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBuckets not implemented")
}
func (UnimplementedManagerServer) ListApplications(context.Context, *ListApplicationsRequest) (*ListApplicationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListApplications not implemented")
}
func (UnimplementedManagerServer) CreateModel(context.Context, *CreateModelRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateModel not implemented")
}
func (UnimplementedManagerServer) KeepAlive(Manager_KeepAliveServer) error {
	return status.Errorf(codes.Unimplemented, "method KeepAlive not implemented")
}

// UnsafeManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ManagerServer will
// result in compilation errors.
type UnsafeManagerServer interface {
	mustEmbedUnimplementedManagerServer()
}

func RegisterManagerServer(s grpc.ServiceRegistrar, srv ManagerServer) {
	s.RegisterService(&Manager_ServiceDesc, srv)
}

func _Manager_GetSeedPeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSeedPeerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetSeedPeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manager.v2.Manager/GetSeedPeer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetSeedPeer(ctx, req.(*GetSeedPeerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_UpdateSeedPeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSeedPeerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).UpdateSeedPeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manager.v2.Manager/UpdateSeedPeer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).UpdateSeedPeer(ctx, req.(*UpdateSeedPeerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_DeleteSeedPeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSeedPeerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).DeleteSeedPeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manager.v2.Manager/DeleteSeedPeer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).DeleteSeedPeer(ctx, req.(*DeleteSeedPeerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetScheduler_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSchedulerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetScheduler(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manager.v2.Manager/GetScheduler",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetScheduler(ctx, req.(*GetSchedulerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_UpdateScheduler_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSchedulerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).UpdateScheduler(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manager.v2.Manager/UpdateScheduler",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).UpdateScheduler(ctx, req.(*UpdateSchedulerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_ListSchedulers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSchedulersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).ListSchedulers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manager.v2.Manager/ListSchedulers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).ListSchedulers(ctx, req.(*ListSchedulersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_GetObjectStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetObjectStorageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).GetObjectStorage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manager.v2.Manager/GetObjectStorage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).GetObjectStorage(ctx, req.(*GetObjectStorageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_ListBuckets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBucketsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).ListBuckets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manager.v2.Manager/ListBuckets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).ListBuckets(ctx, req.(*ListBucketsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_ListApplications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListApplicationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).ListApplications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manager.v2.Manager/ListApplications",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).ListApplications(ctx, req.(*ListApplicationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_CreateModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateModelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).CreateModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manager.v2.Manager/CreateModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).CreateModel(ctx, req.(*CreateModelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_KeepAlive_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ManagerServer).KeepAlive(&managerKeepAliveServer{stream})
}

type Manager_KeepAliveServer interface {
	SendAndClose(*emptypb.Empty) error
	Recv() (*KeepAliveRequest, error)
	grpc.ServerStream
}

type managerKeepAliveServer struct {
	grpc.ServerStream
}

func (x *managerKeepAliveServer) SendAndClose(m *emptypb.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *managerKeepAliveServer) Recv() (*KeepAliveRequest, error) {
	m := new(KeepAliveRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Manager_ServiceDesc is the grpc.ServiceDesc for Manager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Manager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "manager.v2.Manager",
	HandlerType: (*ManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSeedPeer",
			Handler:    _Manager_GetSeedPeer_Handler,
		},
		{
			MethodName: "UpdateSeedPeer",
			Handler:    _Manager_UpdateSeedPeer_Handler,
		},
		{
			MethodName: "DeleteSeedPeer",
			Handler:    _Manager_DeleteSeedPeer_Handler,
		},
		{
			MethodName: "GetScheduler",
			Handler:    _Manager_GetScheduler_Handler,
		},
		{
			MethodName: "UpdateScheduler",
			Handler:    _Manager_UpdateScheduler_Handler,
		},
		{
			MethodName: "ListSchedulers",
			Handler:    _Manager_ListSchedulers_Handler,
		},
		{
			MethodName: "GetObjectStorage",
			Handler:    _Manager_GetObjectStorage_Handler,
		},
		{
			MethodName: "ListBuckets",
			Handler:    _Manager_ListBuckets_Handler,
		},
		{
			MethodName: "ListApplications",
			Handler:    _Manager_ListApplications_Handler,
		},
		{
			MethodName: "CreateModel",
			Handler:    _Manager_CreateModel_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "KeepAlive",
			Handler:       _Manager_KeepAlive_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "pkg/apis/manager/v2/manager.proto",
}
