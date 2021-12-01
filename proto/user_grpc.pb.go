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

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	BindAlias(ctx context.Context, in *BindAliasReq, opts ...grpc.CallOption) (*BindAliasResp, error)
	QueryAliasByCid(ctx context.Context, in *QueryAliasReq, opts ...grpc.CallOption) (*QueryAliasResp, error)
	QueryCidByAlias(ctx context.Context, in *QueryCidReq, opts ...grpc.CallOption) (*QueryCidResp, error)
	UnbindAlias(ctx context.Context, in *UnbindAliasReq, opts ...grpc.CallOption) (*UnbindAliasResp, error)
	RevokeAlias(ctx context.Context, in *RevokeAliasReq, opts ...grpc.CallOption) (*RevokeAliasResp, error)
	BindUserWithTag(ctx context.Context, in *BindUserWithTagReq, opts ...grpc.CallOption) (*BindUserWithTagResp, error)
	BindTagWithUser(ctx context.Context, in *BindTagWithUserReq, opts ...grpc.CallOption) (*BindTagWithUserResp, error)
	UnbindTagFromUser(ctx context.Context, in *UnbindTagFromUserReq, opts ...grpc.CallOption) (*UnbindTagFromUserResp, error)
	QueryUserTag(ctx context.Context, in *QueryUserTagReq, opts ...grpc.CallOption) (*QueryUserTagResp, error)
	AddBlackList(ctx context.Context, in *AddBlackListReq, opts ...grpc.CallOption) (*AddBlackListResp, error)
	DelBlackList(ctx context.Context, in *DelBlackListReq, opts ...grpc.CallOption) (*DelBlackListResp, error)
	QueryUserStatus(ctx context.Context, in *QueryUserStatusReq, opts ...grpc.CallOption) (*QueryUserStatusResp, error)
	QueryDeviceStatus(ctx context.Context, in *QueryDeviceStatusReq, opts ...grpc.CallOption) (*QueryDeviceStatusResp, error)
	QueryUserInfo(ctx context.Context, in *QueryUserInfoReq, opts ...grpc.CallOption) (*QueryUserInfoResp, error)
	SetPushBadge(ctx context.Context, in *SetPushBadgeReq, opts ...grpc.CallOption) (*SetPushBadgeResp, error)
	QueryUserCount(ctx context.Context, in *QueryUserCountReq, opts ...grpc.CallOption) (*QueryUserCountResp, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) BindAlias(ctx context.Context, in *BindAliasReq, opts ...grpc.CallOption) (*BindAliasResp, error) {
	out := new(BindAliasResp)
	err := c.cc.Invoke(ctx, "/OfflinePush.User.User/BindAlias", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) QueryAliasByCid(ctx context.Context, in *QueryAliasReq, opts ...grpc.CallOption) (*QueryAliasResp, error) {
	out := new(QueryAliasResp)
	err := c.cc.Invoke(ctx, "/OfflinePush.User.User/QueryAliasByCid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) QueryCidByAlias(ctx context.Context, in *QueryCidReq, opts ...grpc.CallOption) (*QueryCidResp, error) {
	out := new(QueryCidResp)
	err := c.cc.Invoke(ctx, "/OfflinePush.User.User/QueryCidByAlias", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UnbindAlias(ctx context.Context, in *UnbindAliasReq, opts ...grpc.CallOption) (*UnbindAliasResp, error) {
	out := new(UnbindAliasResp)
	err := c.cc.Invoke(ctx, "/OfflinePush.User.User/UnbindAlias", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) RevokeAlias(ctx context.Context, in *RevokeAliasReq, opts ...grpc.CallOption) (*RevokeAliasResp, error) {
	out := new(RevokeAliasResp)
	err := c.cc.Invoke(ctx, "/OfflinePush.User.User/RevokeAlias", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) BindUserWithTag(ctx context.Context, in *BindUserWithTagReq, opts ...grpc.CallOption) (*BindUserWithTagResp, error) {
	out := new(BindUserWithTagResp)
	err := c.cc.Invoke(ctx, "/OfflinePush.User.User/BindUserWithTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) BindTagWithUser(ctx context.Context, in *BindTagWithUserReq, opts ...grpc.CallOption) (*BindTagWithUserResp, error) {
	out := new(BindTagWithUserResp)
	err := c.cc.Invoke(ctx, "/OfflinePush.User.User/BindTagWithUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UnbindTagFromUser(ctx context.Context, in *UnbindTagFromUserReq, opts ...grpc.CallOption) (*UnbindTagFromUserResp, error) {
	out := new(UnbindTagFromUserResp)
	err := c.cc.Invoke(ctx, "/OfflinePush.User.User/UnbindTagFromUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) QueryUserTag(ctx context.Context, in *QueryUserTagReq, opts ...grpc.CallOption) (*QueryUserTagResp, error) {
	out := new(QueryUserTagResp)
	err := c.cc.Invoke(ctx, "/OfflinePush.User.User/QueryUserTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) AddBlackList(ctx context.Context, in *AddBlackListReq, opts ...grpc.CallOption) (*AddBlackListResp, error) {
	out := new(AddBlackListResp)
	err := c.cc.Invoke(ctx, "/OfflinePush.User.User/AddBlackList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) DelBlackList(ctx context.Context, in *DelBlackListReq, opts ...grpc.CallOption) (*DelBlackListResp, error) {
	out := new(DelBlackListResp)
	err := c.cc.Invoke(ctx, "/OfflinePush.User.User/DelBlackList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) QueryUserStatus(ctx context.Context, in *QueryUserStatusReq, opts ...grpc.CallOption) (*QueryUserStatusResp, error) {
	out := new(QueryUserStatusResp)
	err := c.cc.Invoke(ctx, "/OfflinePush.User.User/QueryUserStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) QueryDeviceStatus(ctx context.Context, in *QueryDeviceStatusReq, opts ...grpc.CallOption) (*QueryDeviceStatusResp, error) {
	out := new(QueryDeviceStatusResp)
	err := c.cc.Invoke(ctx, "/OfflinePush.User.User/QueryDeviceStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) QueryUserInfo(ctx context.Context, in *QueryUserInfoReq, opts ...grpc.CallOption) (*QueryUserInfoResp, error) {
	out := new(QueryUserInfoResp)
	err := c.cc.Invoke(ctx, "/OfflinePush.User.User/QueryUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) SetPushBadge(ctx context.Context, in *SetPushBadgeReq, opts ...grpc.CallOption) (*SetPushBadgeResp, error) {
	out := new(SetPushBadgeResp)
	err := c.cc.Invoke(ctx, "/OfflinePush.User.User/SetPushBadge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) QueryUserCount(ctx context.Context, in *QueryUserCountReq, opts ...grpc.CallOption) (*QueryUserCountResp, error) {
	out := new(QueryUserCountResp)
	err := c.cc.Invoke(ctx, "/OfflinePush.User.User/QueryUserCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	BindAlias(context.Context, *BindAliasReq) (*BindAliasResp, error)
	QueryAliasByCid(context.Context, *QueryAliasReq) (*QueryAliasResp, error)
	QueryCidByAlias(context.Context, *QueryCidReq) (*QueryCidResp, error)
	UnbindAlias(context.Context, *UnbindAliasReq) (*UnbindAliasResp, error)
	RevokeAlias(context.Context, *RevokeAliasReq) (*RevokeAliasResp, error)
	BindUserWithTag(context.Context, *BindUserWithTagReq) (*BindUserWithTagResp, error)
	BindTagWithUser(context.Context, *BindTagWithUserReq) (*BindTagWithUserResp, error)
	UnbindTagFromUser(context.Context, *UnbindTagFromUserReq) (*UnbindTagFromUserResp, error)
	QueryUserTag(context.Context, *QueryUserTagReq) (*QueryUserTagResp, error)
	AddBlackList(context.Context, *AddBlackListReq) (*AddBlackListResp, error)
	DelBlackList(context.Context, *DelBlackListReq) (*DelBlackListResp, error)
	QueryUserStatus(context.Context, *QueryUserStatusReq) (*QueryUserStatusResp, error)
	QueryDeviceStatus(context.Context, *QueryDeviceStatusReq) (*QueryDeviceStatusResp, error)
	QueryUserInfo(context.Context, *QueryUserInfoReq) (*QueryUserInfoResp, error)
	SetPushBadge(context.Context, *SetPushBadgeReq) (*SetPushBadgeResp, error)
	QueryUserCount(context.Context, *QueryUserCountReq) (*QueryUserCountResp, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) BindAlias(context.Context, *BindAliasReq) (*BindAliasResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BindAlias not implemented")
}
func (UnimplementedUserServer) QueryAliasByCid(context.Context, *QueryAliasReq) (*QueryAliasResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryAliasByCid not implemented")
}
func (UnimplementedUserServer) QueryCidByAlias(context.Context, *QueryCidReq) (*QueryCidResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryCidByAlias not implemented")
}
func (UnimplementedUserServer) UnbindAlias(context.Context, *UnbindAliasReq) (*UnbindAliasResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnbindAlias not implemented")
}
func (UnimplementedUserServer) RevokeAlias(context.Context, *RevokeAliasReq) (*RevokeAliasResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeAlias not implemented")
}
func (UnimplementedUserServer) BindUserWithTag(context.Context, *BindUserWithTagReq) (*BindUserWithTagResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BindUserWithTag not implemented")
}
func (UnimplementedUserServer) BindTagWithUser(context.Context, *BindTagWithUserReq) (*BindTagWithUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BindTagWithUser not implemented")
}
func (UnimplementedUserServer) UnbindTagFromUser(context.Context, *UnbindTagFromUserReq) (*UnbindTagFromUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnbindTagFromUser not implemented")
}
func (UnimplementedUserServer) QueryUserTag(context.Context, *QueryUserTagReq) (*QueryUserTagResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryUserTag not implemented")
}
func (UnimplementedUserServer) AddBlackList(context.Context, *AddBlackListReq) (*AddBlackListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBlackList not implemented")
}
func (UnimplementedUserServer) DelBlackList(context.Context, *DelBlackListReq) (*DelBlackListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelBlackList not implemented")
}
func (UnimplementedUserServer) QueryUserStatus(context.Context, *QueryUserStatusReq) (*QueryUserStatusResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryUserStatus not implemented")
}
func (UnimplementedUserServer) QueryDeviceStatus(context.Context, *QueryDeviceStatusReq) (*QueryDeviceStatusResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryDeviceStatus not implemented")
}
func (UnimplementedUserServer) QueryUserInfo(context.Context, *QueryUserInfoReq) (*QueryUserInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryUserInfo not implemented")
}
func (UnimplementedUserServer) SetPushBadge(context.Context, *SetPushBadgeReq) (*SetPushBadgeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPushBadge not implemented")
}
func (UnimplementedUserServer) QueryUserCount(context.Context, *QueryUserCountReq) (*QueryUserCountResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryUserCount not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_BindAlias_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BindAliasReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).BindAlias(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OfflinePush.User.User/BindAlias",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).BindAlias(ctx, req.(*BindAliasReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_QueryAliasByCid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAliasReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).QueryAliasByCid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OfflinePush.User.User/QueryAliasByCid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).QueryAliasByCid(ctx, req.(*QueryAliasReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_QueryCidByAlias_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCidReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).QueryCidByAlias(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OfflinePush.User.User/QueryCidByAlias",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).QueryCidByAlias(ctx, req.(*QueryCidReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UnbindAlias_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnbindAliasReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UnbindAlias(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OfflinePush.User.User/UnbindAlias",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UnbindAlias(ctx, req.(*UnbindAliasReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_RevokeAlias_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokeAliasReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).RevokeAlias(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OfflinePush.User.User/RevokeAlias",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).RevokeAlias(ctx, req.(*RevokeAliasReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_BindUserWithTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BindUserWithTagReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).BindUserWithTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OfflinePush.User.User/BindUserWithTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).BindUserWithTag(ctx, req.(*BindUserWithTagReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_BindTagWithUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BindTagWithUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).BindTagWithUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OfflinePush.User.User/BindTagWithUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).BindTagWithUser(ctx, req.(*BindTagWithUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UnbindTagFromUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnbindTagFromUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UnbindTagFromUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OfflinePush.User.User/UnbindTagFromUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UnbindTagFromUser(ctx, req.(*UnbindTagFromUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_QueryUserTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryUserTagReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).QueryUserTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OfflinePush.User.User/QueryUserTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).QueryUserTag(ctx, req.(*QueryUserTagReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_AddBlackList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddBlackListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).AddBlackList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OfflinePush.User.User/AddBlackList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).AddBlackList(ctx, req.(*AddBlackListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_DelBlackList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelBlackListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).DelBlackList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OfflinePush.User.User/DelBlackList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).DelBlackList(ctx, req.(*DelBlackListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_QueryUserStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryUserStatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).QueryUserStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OfflinePush.User.User/QueryUserStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).QueryUserStatus(ctx, req.(*QueryUserStatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_QueryDeviceStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDeviceStatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).QueryDeviceStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OfflinePush.User.User/QueryDeviceStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).QueryDeviceStatus(ctx, req.(*QueryDeviceStatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_QueryUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryUserInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).QueryUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OfflinePush.User.User/QueryUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).QueryUserInfo(ctx, req.(*QueryUserInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_SetPushBadge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetPushBadgeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).SetPushBadge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OfflinePush.User.User/SetPushBadge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).SetPushBadge(ctx, req.(*SetPushBadgeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_QueryUserCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryUserCountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).QueryUserCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OfflinePush.User.User/QueryUserCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).QueryUserCount(ctx, req.(*QueryUserCountReq))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "OfflinePush.User.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BindAlias",
			Handler:    _User_BindAlias_Handler,
		},
		{
			MethodName: "QueryAliasByCid",
			Handler:    _User_QueryAliasByCid_Handler,
		},
		{
			MethodName: "QueryCidByAlias",
			Handler:    _User_QueryCidByAlias_Handler,
		},
		{
			MethodName: "UnbindAlias",
			Handler:    _User_UnbindAlias_Handler,
		},
		{
			MethodName: "RevokeAlias",
			Handler:    _User_RevokeAlias_Handler,
		},
		{
			MethodName: "BindUserWithTag",
			Handler:    _User_BindUserWithTag_Handler,
		},
		{
			MethodName: "BindTagWithUser",
			Handler:    _User_BindTagWithUser_Handler,
		},
		{
			MethodName: "UnbindTagFromUser",
			Handler:    _User_UnbindTagFromUser_Handler,
		},
		{
			MethodName: "QueryUserTag",
			Handler:    _User_QueryUserTag_Handler,
		},
		{
			MethodName: "AddBlackList",
			Handler:    _User_AddBlackList_Handler,
		},
		{
			MethodName: "DelBlackList",
			Handler:    _User_DelBlackList_Handler,
		},
		{
			MethodName: "QueryUserStatus",
			Handler:    _User_QueryUserStatus_Handler,
		},
		{
			MethodName: "QueryDeviceStatus",
			Handler:    _User_QueryDeviceStatus_Handler,
		},
		{
			MethodName: "QueryUserInfo",
			Handler:    _User_QueryUserInfo_Handler,
		},
		{
			MethodName: "SetPushBadge",
			Handler:    _User_SetPushBadge_Handler,
		},
		{
			MethodName: "QueryUserCount",
			Handler:    _User_QueryUserCount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
