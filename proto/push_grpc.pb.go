// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// PushClient is the client API for Push service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PushClient interface {
	PushToSingle(ctx context.Context, in *PushToSingleReq, opts ...grpc.CallOption) (*PushToSingleResp, error)
	CreateTask(ctx context.Context, in *CreateTaskReq, opts ...grpc.CallOption) (*CreateTaskResp, error)
	PushToList(ctx context.Context, in *PushToListRep, opts ...grpc.CallOption) (*PushToListResp, error)
	PushToApp(ctx context.Context, in *PushToAppReq, opts ...grpc.CallOption) (*PushToAppResp, error)
	StopTask(ctx context.Context, in *StopTaskReq, opts ...grpc.CallOption) (*StopTaskResp, error)
	CheckTask(ctx context.Context, in *CheckTaskReq, opts ...grpc.CallOption) (*CheckTaskResp, error)
	RemoveTask(ctx context.Context, in *RemoveTaskReq, opts ...grpc.CallOption) (*RemoveTaskResp, error)
	ViewDetail(ctx context.Context, in *ViewDetailReq, opts ...grpc.CallOption) (*ViewDetailResp, error)
}

type pushClient struct {
	cc grpc.ClientConnInterface
}

func NewPushClient(cc grpc.ClientConnInterface) PushClient {
	return &pushClient{cc}
}

func (c *pushClient) PushToSingle(ctx context.Context, in *PushToSingleReq, opts ...grpc.CallOption) (*PushToSingleResp, error) {
	out := new(PushToSingleResp)
	err := c.cc.Invoke(ctx, "/OfflinePush.Push.Push/PushToSingle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pushClient) CreateTask(ctx context.Context, in *CreateTaskReq, opts ...grpc.CallOption) (*CreateTaskResp, error) {
	out := new(CreateTaskResp)
	err := c.cc.Invoke(ctx, "/OfflinePush.Push.Push/CreateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pushClient) PushToList(ctx context.Context, in *PushToListRep, opts ...grpc.CallOption) (*PushToListResp, error) {
	out := new(PushToListResp)
	err := c.cc.Invoke(ctx, "/OfflinePush.Push.Push/PushToList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pushClient) PushToApp(ctx context.Context, in *PushToAppReq, opts ...grpc.CallOption) (*PushToAppResp, error) {
	out := new(PushToAppResp)
	err := c.cc.Invoke(ctx, "/OfflinePush.Push.Push/PushToApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pushClient) StopTask(ctx context.Context, in *StopTaskReq, opts ...grpc.CallOption) (*StopTaskResp, error) {
	out := new(StopTaskResp)
	err := c.cc.Invoke(ctx, "/OfflinePush.Push.Push/StopTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pushClient) CheckTask(ctx context.Context, in *CheckTaskReq, opts ...grpc.CallOption) (*CheckTaskResp, error) {
	out := new(CheckTaskResp)
	err := c.cc.Invoke(ctx, "/OfflinePush.Push.Push/CheckTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pushClient) RemoveTask(ctx context.Context, in *RemoveTaskReq, opts ...grpc.CallOption) (*RemoveTaskResp, error) {
	out := new(RemoveTaskResp)
	err := c.cc.Invoke(ctx, "/OfflinePush.Push.Push/RemoveTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pushClient) ViewDetail(ctx context.Context, in *ViewDetailReq, opts ...grpc.CallOption) (*ViewDetailResp, error) {
	out := new(ViewDetailResp)
	err := c.cc.Invoke(ctx, "/OfflinePush.Push.Push/ViewDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PushServer is the server API for Push service.
// All implementations must embed UnimplementedPushServer
// for forward compatibility
type PushServer interface {
	PushToSingle(context.Context, *PushToSingleReq) (*PushToSingleResp, error)
	CreateTask(context.Context, *CreateTaskReq) (*CreateTaskResp, error)
	PushToList(context.Context, *PushToListRep) (*PushToListResp, error)
	PushToApp(context.Context, *PushToAppReq) (*PushToAppResp, error)
	StopTask(context.Context, *StopTaskReq) (*StopTaskResp, error)
	CheckTask(context.Context, *CheckTaskReq) (*CheckTaskResp, error)
	RemoveTask(context.Context, *RemoveTaskReq) (*RemoveTaskResp, error)
	ViewDetail(context.Context, *ViewDetailReq) (*ViewDetailResp, error)
	mustEmbedUnimplementedPushServer()
}

// UnimplementedPushServer must be embedded to have forward compatible implementations.
type UnimplementedPushServer struct {
}

func (UnimplementedPushServer) PushToSingle(context.Context, *PushToSingleReq) (*PushToSingleResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushToSingle not implemented")
}
func (UnimplementedPushServer) CreateTask(context.Context, *CreateTaskReq) (*CreateTaskResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTask not implemented")
}
func (UnimplementedPushServer) PushToList(context.Context, *PushToListRep) (*PushToListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushToList not implemented")
}
func (UnimplementedPushServer) PushToApp(context.Context, *PushToAppReq) (*PushToAppResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushToApp not implemented")
}
func (UnimplementedPushServer) StopTask(context.Context, *StopTaskReq) (*StopTaskResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopTask not implemented")
}
func (UnimplementedPushServer) CheckTask(context.Context, *CheckTaskReq) (*CheckTaskResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckTask not implemented")
}
func (UnimplementedPushServer) RemoveTask(context.Context, *RemoveTaskReq) (*RemoveTaskResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveTask not implemented")
}
func (UnimplementedPushServer) ViewDetail(context.Context, *ViewDetailReq) (*ViewDetailResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewDetail not implemented")
}
func (UnimplementedPushServer) mustEmbedUnimplementedPushServer() {}

// UnsafePushServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PushServer will
// result in compilation errors.
type UnsafePushServer interface {
	mustEmbedUnimplementedPushServer()
}

func RegisterPushServer(s grpc.ServiceRegistrar, srv PushServer) {
	s.RegisterService(&Push_ServiceDesc, srv)
}

func _Push_PushToSingle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushToSingleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PushServer).PushToSingle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OfflinePush.Push.Push/PushToSingle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PushServer).PushToSingle(ctx, req.(*PushToSingleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Push_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaskReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PushServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OfflinePush.Push.Push/CreateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PushServer).CreateTask(ctx, req.(*CreateTaskReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Push_PushToList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushToListRep)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PushServer).PushToList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OfflinePush.Push.Push/PushToList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PushServer).PushToList(ctx, req.(*PushToListRep))
	}
	return interceptor(ctx, in, info, handler)
}

func _Push_PushToApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushToAppReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PushServer).PushToApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OfflinePush.Push.Push/PushToApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PushServer).PushToApp(ctx, req.(*PushToAppReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Push_StopTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopTaskReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PushServer).StopTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OfflinePush.Push.Push/StopTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PushServer).StopTask(ctx, req.(*StopTaskReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Push_CheckTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckTaskReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PushServer).CheckTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OfflinePush.Push.Push/CheckTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PushServer).CheckTask(ctx, req.(*CheckTaskReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Push_RemoveTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveTaskReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PushServer).RemoveTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OfflinePush.Push.Push/RemoveTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PushServer).RemoveTask(ctx, req.(*RemoveTaskReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Push_ViewDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewDetailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PushServer).ViewDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OfflinePush.Push.Push/ViewDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PushServer).ViewDetail(ctx, req.(*ViewDetailReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Push_ServiceDesc is the grpc.ServiceDesc for Push service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Push_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "OfflinePush.Push.Push",
	HandlerType: (*PushServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PushToSingle",
			Handler:    _Push_PushToSingle_Handler,
		},
		{
			MethodName: "CreateTask",
			Handler:    _Push_CreateTask_Handler,
		},
		{
			MethodName: "PushToList",
			Handler:    _Push_PushToList_Handler,
		},
		{
			MethodName: "PushToApp",
			Handler:    _Push_PushToApp_Handler,
		},
		{
			MethodName: "StopTask",
			Handler:    _Push_StopTask_Handler,
		},
		{
			MethodName: "CheckTask",
			Handler:    _Push_CheckTask_Handler,
		},
		{
			MethodName: "RemoveTask",
			Handler:    _Push_RemoveTask_Handler,
		},
		{
			MethodName: "ViewDetail",
			Handler:    _Push_ViewDetail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "push.proto",
}
