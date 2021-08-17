// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: auth.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetTokenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PushAgent string `protobuf:"bytes,1,opt,name=push_agent,json=pushAgent,proto3" json:"push_agent,omitempty"`
	BundleId  string `protobuf:"bytes,2,opt,name=bundle_id,json=bundleId,proto3" json:"bundle_id,omitempty"`
	OldToken  string `protobuf:"bytes,3,opt,name=old_token,json=oldToken,proto3" json:"old_token,omitempty"`
}

func (x *GetTokenReq) Reset() {
	*x = GetTokenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTokenReq) ProtoMessage() {}

func (x *GetTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTokenReq.ProtoReflect.Descriptor instead.
func (*GetTokenReq) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{0}
}

func (x *GetTokenReq) GetPushAgent() string {
	if x != nil {
		return x.PushAgent
	}
	return ""
}

func (x *GetTokenReq) GetBundleId() string {
	if x != nil {
		return x.BundleId
	}
	return ""
}

func (x *GetTokenReq) GetOldToken() string {
	if x != nil {
		return x.OldToken
	}
	return ""
}

type GetTokenResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token    string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	ExpireAt int64  `protobuf:"varint,2,opt,name=expire_at,json=expireAt,proto3" json:"expire_at,omitempty"` // unix_second
}

func (x *GetTokenResp) Reset() {
	*x = GetTokenResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTokenResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTokenResp) ProtoMessage() {}

func (x *GetTokenResp) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTokenResp.ProtoReflect.Descriptor instead.
func (*GetTokenResp) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{1}
}

func (x *GetTokenResp) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *GetTokenResp) GetExpireAt() int64 {
	if x != nil {
		return x.ExpireAt
	}
	return 0
}

type SetTokenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PushAgent string `protobuf:"bytes,1,opt,name=push_agent,json=pushAgent,proto3" json:"push_agent,omitempty"`
	BundleId  string `protobuf:"bytes,2,opt,name=bundle_id,json=bundleId,proto3" json:"bundle_id,omitempty"`
	Token     string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	ExpireAt  int64  `protobuf:"varint,4,opt,name=expire_at,json=expireAt,proto3" json:"expire_at,omitempty"` // unix_second
}

func (x *SetTokenReq) Reset() {
	*x = SetTokenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetTokenReq) ProtoMessage() {}

func (x *SetTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetTokenReq.ProtoReflect.Descriptor instead.
func (*SetTokenReq) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{2}
}

func (x *SetTokenReq) GetPushAgent() string {
	if x != nil {
		return x.PushAgent
	}
	return ""
}

func (x *SetTokenReq) GetBundleId() string {
	if x != nil {
		return x.BundleId
	}
	return ""
}

func (x *SetTokenReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *SetTokenReq) GetExpireAt() int64 {
	if x != nil {
		return x.ExpireAt
	}
	return 0
}

type SetTokenResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetTokenResp) Reset() {
	*x = SetTokenResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetTokenResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetTokenResp) ProtoMessage() {}

func (x *SetTokenResp) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetTokenResp.ProtoReflect.Descriptor instead.
func (*SetTokenResp) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{3}
}

type DelTokenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PushAgent string `protobuf:"bytes,1,opt,name=push_agent,json=pushAgent,proto3" json:"push_agent,omitempty"`
	BundleId  string `protobuf:"bytes,2,opt,name=bundle_id,json=bundleId,proto3" json:"bundle_id,omitempty"`
	Token     string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *DelTokenReq) Reset() {
	*x = DelTokenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelTokenReq) ProtoMessage() {}

func (x *DelTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelTokenReq.ProtoReflect.Descriptor instead.
func (*DelTokenReq) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{4}
}

func (x *DelTokenReq) GetPushAgent() string {
	if x != nil {
		return x.PushAgent
	}
	return ""
}

func (x *DelTokenReq) GetBundleId() string {
	if x != nil {
		return x.BundleId
	}
	return ""
}

func (x *DelTokenReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type DelTokenResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DelTokenResp) Reset() {
	*x = DelTokenResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelTokenResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelTokenResp) ProtoMessage() {}

func (x *DelTokenResp) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelTokenResp.ProtoReflect.Descriptor instead.
func (*DelTokenResp) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{5}
}

var File_auth_proto protoreflect.FileDescriptor

var file_auth_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x4f, 0x66,
	0x66, 0x6c, 0x69, 0x6e, 0x65, 0x50, 0x75, 0x73, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x22, 0x66,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x75, 0x73, 0x68, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x75, 0x73, 0x68, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x6c, 0x64,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x6c,
	0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x41, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1b, 0x0a, 0x09,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x41, 0x74, 0x22, 0x7c, 0x0a, 0x0b, 0x53, 0x65, 0x74,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x73, 0x68,
	0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75,
	0x73, 0x68, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x75, 0x6e, 0x64, 0x6c,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x75, 0x6e, 0x64,
	0x6c, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x41, 0x74, 0x22, 0x0e, 0x0a, 0x0c, 0x53, 0x65, 0x74, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x5f, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x73, 0x68, 0x5f, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x73, 0x68,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x0e, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x32, 0xed, 0x01, 0x0a, 0x04, 0x41, 0x75, 0x74,
	0x68, 0x12, 0x4b, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x2e,
	0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x50, 0x75, 0x73, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x4f,
	0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x50, 0x75, 0x73, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x2e,
	0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x4b,
	0x0a, 0x08, 0x53, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x2e, 0x4f, 0x66, 0x66,
	0x6c, 0x69, 0x6e, 0x65, 0x50, 0x75, 0x73, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x2e, 0x53, 0x65,
	0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x4f, 0x66, 0x66, 0x6c,
	0x69, 0x6e, 0x65, 0x50, 0x75, 0x73, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x2e, 0x53, 0x65, 0x74,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x08, 0x44,
	0x65, 0x6c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x2e, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e,
	0x65, 0x50, 0x75, 0x73, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x2e, 0x44, 0x65, 0x6c, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65,
	0x50, 0x75, 0x73, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x2e, 0x44, 0x65, 0x6c, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_proto_rawDescOnce sync.Once
	file_auth_proto_rawDescData = file_auth_proto_rawDesc
)

func file_auth_proto_rawDescGZIP() []byte {
	file_auth_proto_rawDescOnce.Do(func() {
		file_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_proto_rawDescData)
	})
	return file_auth_proto_rawDescData
}

var file_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_auth_proto_goTypes = []interface{}{
	(*GetTokenReq)(nil),  // 0: OfflinePush.Auth.GetTokenReq
	(*GetTokenResp)(nil), // 1: OfflinePush.Auth.GetTokenResp
	(*SetTokenReq)(nil),  // 2: OfflinePush.Auth.SetTokenReq
	(*SetTokenResp)(nil), // 3: OfflinePush.Auth.SetTokenResp
	(*DelTokenReq)(nil),  // 4: OfflinePush.Auth.DelTokenReq
	(*DelTokenResp)(nil), // 5: OfflinePush.Auth.DelTokenResp
}
var file_auth_proto_depIdxs = []int32{
	0, // 0: OfflinePush.Auth.Auth.GetToken:input_type -> OfflinePush.Auth.GetTokenReq
	2, // 1: OfflinePush.Auth.Auth.SetToken:input_type -> OfflinePush.Auth.SetTokenReq
	4, // 2: OfflinePush.Auth.Auth.DelToken:input_type -> OfflinePush.Auth.DelTokenReq
	1, // 3: OfflinePush.Auth.Auth.GetToken:output_type -> OfflinePush.Auth.GetTokenResp
	3, // 4: OfflinePush.Auth.Auth.SetToken:output_type -> OfflinePush.Auth.SetTokenResp
	5, // 5: OfflinePush.Auth.Auth.DelToken:output_type -> OfflinePush.Auth.DelTokenResp
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_auth_proto_init() }
func file_auth_proto_init() {
	if File_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTokenReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTokenResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetTokenReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetTokenResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelTokenReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelTokenResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_proto_goTypes,
		DependencyIndexes: file_auth_proto_depIdxs,
		MessageInfos:      file_auth_proto_msgTypes,
	}.Build()
	File_auth_proto = out.File
	file_auth_proto_rawDesc = nil
	file_auth_proto_goTypes = nil
	file_auth_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthClient interface {
	GetToken(ctx context.Context, in *GetTokenReq, opts ...grpc.CallOption) (*GetTokenResp, error)
	SetToken(ctx context.Context, in *SetTokenReq, opts ...grpc.CallOption) (*SetTokenResp, error)
	DelToken(ctx context.Context, in *DelTokenReq, opts ...grpc.CallOption) (*DelTokenResp, error)
}

type authClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthClient(cc grpc.ClientConnInterface) AuthClient {
	return &authClient{cc}
}

func (c *authClient) GetToken(ctx context.Context, in *GetTokenReq, opts ...grpc.CallOption) (*GetTokenResp, error) {
	out := new(GetTokenResp)
	err := c.cc.Invoke(ctx, "/OfflinePush.Auth.Auth/GetToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) SetToken(ctx context.Context, in *SetTokenReq, opts ...grpc.CallOption) (*SetTokenResp, error) {
	out := new(SetTokenResp)
	err := c.cc.Invoke(ctx, "/OfflinePush.Auth.Auth/SetToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) DelToken(ctx context.Context, in *DelTokenReq, opts ...grpc.CallOption) (*DelTokenResp, error) {
	out := new(DelTokenResp)
	err := c.cc.Invoke(ctx, "/OfflinePush.Auth.Auth/DelToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
type AuthServer interface {
	GetToken(context.Context, *GetTokenReq) (*GetTokenResp, error)
	SetToken(context.Context, *SetTokenReq) (*SetTokenResp, error)
	DelToken(context.Context, *DelTokenReq) (*DelTokenResp, error)
}

// UnimplementedAuthServer can be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (*UnimplementedAuthServer) GetToken(context.Context, *GetTokenReq) (*GetTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetToken not implemented")
}
func (*UnimplementedAuthServer) SetToken(context.Context, *SetTokenReq) (*SetTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetToken not implemented")
}
func (*UnimplementedAuthServer) DelToken(context.Context, *DelTokenReq) (*DelTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelToken not implemented")
}

func RegisterAuthServer(s *grpc.Server, srv AuthServer) {
	s.RegisterService(&_Auth_serviceDesc, srv)
}

func _Auth_GetToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OfflinePush.Auth.Auth/GetToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetToken(ctx, req.(*GetTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_SetToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).SetToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OfflinePush.Auth.Auth/SetToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).SetToken(ctx, req.(*SetTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_DelToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).DelToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OfflinePush.Auth.Auth/DelToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).DelToken(ctx, req.(*DelTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Auth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "OfflinePush.Auth.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetToken",
			Handler:    _Auth_GetToken_Handler,
		},
		{
			MethodName: "SetToken",
			Handler:    _Auth_SetToken_Handler,
		},
		{
			MethodName: "DelToken",
			Handler:    _Auth_DelToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
