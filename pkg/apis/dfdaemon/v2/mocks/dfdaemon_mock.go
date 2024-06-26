// Code generated by MockGen. DO NOT EDIT.
// Source: ../dfdaemon_grpc.pb.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	common "d7y.io/api/v2/pkg/apis/common/v2"
	dfdaemon "d7y.io/api/v2/pkg/apis/dfdaemon/v2"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// MockDfdaemonClient is a mock of DfdaemonClient interface.
type MockDfdaemonClient struct {
	ctrl     *gomock.Controller
	recorder *MockDfdaemonClientMockRecorder
}

// MockDfdaemonClientMockRecorder is the mock recorder for MockDfdaemonClient.
type MockDfdaemonClientMockRecorder struct {
	mock *MockDfdaemonClient
}

// NewMockDfdaemonClient creates a new mock instance.
func NewMockDfdaemonClient(ctrl *gomock.Controller) *MockDfdaemonClient {
	mock := &MockDfdaemonClient{ctrl: ctrl}
	mock.recorder = &MockDfdaemonClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDfdaemonClient) EXPECT() *MockDfdaemonClientMockRecorder {
	return m.recorder
}

// DeleteTask mocks base method.
func (m *MockDfdaemonClient) DeleteTask(ctx context.Context, in *dfdaemon.DeleteTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteTask", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteTask indicates an expected call of DeleteTask.
func (mr *MockDfdaemonClientMockRecorder) DeleteTask(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTask", reflect.TypeOf((*MockDfdaemonClient)(nil).DeleteTask), varargs...)
}

// DownloadTask mocks base method.
func (m *MockDfdaemonClient) DownloadTask(ctx context.Context, in *dfdaemon.DownloadTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DownloadTask", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DownloadTask indicates an expected call of DownloadTask.
func (mr *MockDfdaemonClientMockRecorder) DownloadTask(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadTask", reflect.TypeOf((*MockDfdaemonClient)(nil).DownloadTask), varargs...)
}

// StatTask mocks base method.
func (m *MockDfdaemonClient) StatTask(ctx context.Context, in *dfdaemon.StatTaskRequest, opts ...grpc.CallOption) (*common.Task, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StatTask", varargs...)
	ret0, _ := ret[0].(*common.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StatTask indicates an expected call of StatTask.
func (mr *MockDfdaemonClientMockRecorder) StatTask(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatTask", reflect.TypeOf((*MockDfdaemonClient)(nil).StatTask), varargs...)
}

// SyncPieces mocks base method.
func (m *MockDfdaemonClient) SyncPieces(ctx context.Context, opts ...grpc.CallOption) (dfdaemon.Dfdaemon_SyncPiecesClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SyncPieces", varargs...)
	ret0, _ := ret[0].(dfdaemon.Dfdaemon_SyncPiecesClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SyncPieces indicates an expected call of SyncPieces.
func (mr *MockDfdaemonClientMockRecorder) SyncPieces(ctx interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncPieces", reflect.TypeOf((*MockDfdaemonClient)(nil).SyncPieces), varargs...)
}

// UploadTask mocks base method.
func (m *MockDfdaemonClient) UploadTask(ctx context.Context, in *dfdaemon.UploadTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UploadTask", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadTask indicates an expected call of UploadTask.
func (mr *MockDfdaemonClientMockRecorder) UploadTask(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadTask", reflect.TypeOf((*MockDfdaemonClient)(nil).UploadTask), varargs...)
}

// MockDfdaemon_SyncPiecesClient is a mock of Dfdaemon_SyncPiecesClient interface.
type MockDfdaemon_SyncPiecesClient struct {
	ctrl     *gomock.Controller
	recorder *MockDfdaemon_SyncPiecesClientMockRecorder
}

// MockDfdaemon_SyncPiecesClientMockRecorder is the mock recorder for MockDfdaemon_SyncPiecesClient.
type MockDfdaemon_SyncPiecesClientMockRecorder struct {
	mock *MockDfdaemon_SyncPiecesClient
}

// NewMockDfdaemon_SyncPiecesClient creates a new mock instance.
func NewMockDfdaemon_SyncPiecesClient(ctrl *gomock.Controller) *MockDfdaemon_SyncPiecesClient {
	mock := &MockDfdaemon_SyncPiecesClient{ctrl: ctrl}
	mock.recorder = &MockDfdaemon_SyncPiecesClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDfdaemon_SyncPiecesClient) EXPECT() *MockDfdaemon_SyncPiecesClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method.
func (m *MockDfdaemon_SyncPiecesClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockDfdaemon_SyncPiecesClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockDfdaemon_SyncPiecesClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockDfdaemon_SyncPiecesClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockDfdaemon_SyncPiecesClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockDfdaemon_SyncPiecesClient)(nil).Context))
}

// Header mocks base method.
func (m *MockDfdaemon_SyncPiecesClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockDfdaemon_SyncPiecesClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockDfdaemon_SyncPiecesClient)(nil).Header))
}

// Recv mocks base method.
func (m *MockDfdaemon_SyncPiecesClient) Recv() (*dfdaemon.SyncPiecesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*dfdaemon.SyncPiecesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockDfdaemon_SyncPiecesClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockDfdaemon_SyncPiecesClient)(nil).Recv))
}

// RecvMsg mocks base method.
func (m_2 *MockDfdaemon_SyncPiecesClient) RecvMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockDfdaemon_SyncPiecesClientMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockDfdaemon_SyncPiecesClient)(nil).RecvMsg), m)
}

// Send mocks base method.
func (m *MockDfdaemon_SyncPiecesClient) Send(arg0 *dfdaemon.SyncPiecesRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockDfdaemon_SyncPiecesClientMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockDfdaemon_SyncPiecesClient)(nil).Send), arg0)
}

// SendMsg mocks base method.
func (m_2 *MockDfdaemon_SyncPiecesClient) SendMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockDfdaemon_SyncPiecesClientMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockDfdaemon_SyncPiecesClient)(nil).SendMsg), m)
}

// Trailer mocks base method.
func (m *MockDfdaemon_SyncPiecesClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockDfdaemon_SyncPiecesClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockDfdaemon_SyncPiecesClient)(nil).Trailer))
}

// MockDfdaemonServer is a mock of DfdaemonServer interface.
type MockDfdaemonServer struct {
	ctrl     *gomock.Controller
	recorder *MockDfdaemonServerMockRecorder
}

// MockDfdaemonServerMockRecorder is the mock recorder for MockDfdaemonServer.
type MockDfdaemonServerMockRecorder struct {
	mock *MockDfdaemonServer
}

// NewMockDfdaemonServer creates a new mock instance.
func NewMockDfdaemonServer(ctrl *gomock.Controller) *MockDfdaemonServer {
	mock := &MockDfdaemonServer{ctrl: ctrl}
	mock.recorder = &MockDfdaemonServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDfdaemonServer) EXPECT() *MockDfdaemonServerMockRecorder {
	return m.recorder
}

// DeleteTask mocks base method.
func (m *MockDfdaemonServer) DeleteTask(arg0 context.Context, arg1 *dfdaemon.DeleteTaskRequest) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTask", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteTask indicates an expected call of DeleteTask.
func (mr *MockDfdaemonServerMockRecorder) DeleteTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTask", reflect.TypeOf((*MockDfdaemonServer)(nil).DeleteTask), arg0, arg1)
}

// DownloadTask mocks base method.
func (m *MockDfdaemonServer) DownloadTask(arg0 context.Context, arg1 *dfdaemon.DownloadTaskRequest) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadTask", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DownloadTask indicates an expected call of DownloadTask.
func (mr *MockDfdaemonServerMockRecorder) DownloadTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadTask", reflect.TypeOf((*MockDfdaemonServer)(nil).DownloadTask), arg0, arg1)
}

// StatTask mocks base method.
func (m *MockDfdaemonServer) StatTask(arg0 context.Context, arg1 *dfdaemon.StatTaskRequest) (*common.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StatTask", arg0, arg1)
	ret0, _ := ret[0].(*common.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StatTask indicates an expected call of StatTask.
func (mr *MockDfdaemonServerMockRecorder) StatTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatTask", reflect.TypeOf((*MockDfdaemonServer)(nil).StatTask), arg0, arg1)
}

// SyncPieces mocks base method.
func (m *MockDfdaemonServer) SyncPieces(arg0 dfdaemon.Dfdaemon_SyncPiecesServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncPieces", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SyncPieces indicates an expected call of SyncPieces.
func (mr *MockDfdaemonServerMockRecorder) SyncPieces(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncPieces", reflect.TypeOf((*MockDfdaemonServer)(nil).SyncPieces), arg0)
}

// UploadTask mocks base method.
func (m *MockDfdaemonServer) UploadTask(arg0 context.Context, arg1 *dfdaemon.UploadTaskRequest) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadTask", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadTask indicates an expected call of UploadTask.
func (mr *MockDfdaemonServerMockRecorder) UploadTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadTask", reflect.TypeOf((*MockDfdaemonServer)(nil).UploadTask), arg0, arg1)
}

// MockUnsafeDfdaemonServer is a mock of UnsafeDfdaemonServer interface.
type MockUnsafeDfdaemonServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeDfdaemonServerMockRecorder
}

// MockUnsafeDfdaemonServerMockRecorder is the mock recorder for MockUnsafeDfdaemonServer.
type MockUnsafeDfdaemonServerMockRecorder struct {
	mock *MockUnsafeDfdaemonServer
}

// NewMockUnsafeDfdaemonServer creates a new mock instance.
func NewMockUnsafeDfdaemonServer(ctrl *gomock.Controller) *MockUnsafeDfdaemonServer {
	mock := &MockUnsafeDfdaemonServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeDfdaemonServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeDfdaemonServer) EXPECT() *MockUnsafeDfdaemonServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedDfdaemonServer mocks base method.
func (m *MockUnsafeDfdaemonServer) mustEmbedUnimplementedDfdaemonServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedDfdaemonServer")
}

// mustEmbedUnimplementedDfdaemonServer indicates an expected call of mustEmbedUnimplementedDfdaemonServer.
func (mr *MockUnsafeDfdaemonServerMockRecorder) mustEmbedUnimplementedDfdaemonServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedDfdaemonServer", reflect.TypeOf((*MockUnsafeDfdaemonServer)(nil).mustEmbedUnimplementedDfdaemonServer))
}

// MockDfdaemon_SyncPiecesServer is a mock of Dfdaemon_SyncPiecesServer interface.
type MockDfdaemon_SyncPiecesServer struct {
	ctrl     *gomock.Controller
	recorder *MockDfdaemon_SyncPiecesServerMockRecorder
}

// MockDfdaemon_SyncPiecesServerMockRecorder is the mock recorder for MockDfdaemon_SyncPiecesServer.
type MockDfdaemon_SyncPiecesServerMockRecorder struct {
	mock *MockDfdaemon_SyncPiecesServer
}

// NewMockDfdaemon_SyncPiecesServer creates a new mock instance.
func NewMockDfdaemon_SyncPiecesServer(ctrl *gomock.Controller) *MockDfdaemon_SyncPiecesServer {
	mock := &MockDfdaemon_SyncPiecesServer{ctrl: ctrl}
	mock.recorder = &MockDfdaemon_SyncPiecesServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDfdaemon_SyncPiecesServer) EXPECT() *MockDfdaemon_SyncPiecesServerMockRecorder {
	return m.recorder
}

// Context mocks base method.
func (m *MockDfdaemon_SyncPiecesServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockDfdaemon_SyncPiecesServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockDfdaemon_SyncPiecesServer)(nil).Context))
}

// Recv mocks base method.
func (m *MockDfdaemon_SyncPiecesServer) Recv() (*dfdaemon.SyncPiecesRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*dfdaemon.SyncPiecesRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockDfdaemon_SyncPiecesServerMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockDfdaemon_SyncPiecesServer)(nil).Recv))
}

// RecvMsg mocks base method.
func (m_2 *MockDfdaemon_SyncPiecesServer) RecvMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "RecvMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockDfdaemon_SyncPiecesServerMockRecorder) RecvMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockDfdaemon_SyncPiecesServer)(nil).RecvMsg), m)
}

// Send mocks base method.
func (m *MockDfdaemon_SyncPiecesServer) Send(arg0 *dfdaemon.SyncPiecesResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockDfdaemon_SyncPiecesServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockDfdaemon_SyncPiecesServer)(nil).Send), arg0)
}

// SendHeader mocks base method.
func (m *MockDfdaemon_SyncPiecesServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader.
func (mr *MockDfdaemon_SyncPiecesServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockDfdaemon_SyncPiecesServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method.
func (m_2 *MockDfdaemon_SyncPiecesServer) SendMsg(m interface{}) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "SendMsg", m)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockDfdaemon_SyncPiecesServerMockRecorder) SendMsg(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockDfdaemon_SyncPiecesServer)(nil).SendMsg), m)
}

// SetHeader mocks base method.
func (m *MockDfdaemon_SyncPiecesServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockDfdaemon_SyncPiecesServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockDfdaemon_SyncPiecesServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method.
func (m *MockDfdaemon_SyncPiecesServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer.
func (mr *MockDfdaemon_SyncPiecesServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockDfdaemon_SyncPiecesServer)(nil).SetTrailer), arg0)
}
