// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: user_core.proto

package user_core

import (
	base "github.com/jlu-cow-studio/common/dao/rpc/base"
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

type UserLoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base     *base.BaseReq `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Username string        `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password string        `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *UserLoginReq) Reset() {
	*x = UserLoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_core_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLoginReq) ProtoMessage() {}

func (x *UserLoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_core_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLoginReq.ProtoReflect.Descriptor instead.
func (*UserLoginReq) Descriptor() ([]byte, []int) {
	return file_user_core_proto_rawDescGZIP(), []int{0}
}

func (x *UserLoginReq) GetBase() *base.BaseReq {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *UserLoginReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserLoginReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type UserLoginRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base  *base.BaseRes `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Token string        `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *UserLoginRes) Reset() {
	*x = UserLoginRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_core_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLoginRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLoginRes) ProtoMessage() {}

func (x *UserLoginRes) ProtoReflect() protoreflect.Message {
	mi := &file_user_core_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLoginRes.ProtoReflect.Descriptor instead.
func (*UserLoginRes) Descriptor() ([]byte, []int) {
	return file_user_core_proto_rawDescGZIP(), []int{1}
}

func (x *UserLoginRes) GetBase() *base.BaseRes {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *UserLoginRes) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid      string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Province string `protobuf:"bytes,4,opt,name=province,proto3" json:"province,omitempty"`
	Role     string `protobuf:"bytes,5,opt,name=role,proto3" json:"role,omitempty"`
	City     string `protobuf:"bytes,6,opt,name=city,proto3" json:"city,omitempty"`
	District string `protobuf:"bytes,7,opt,name=district,proto3" json:"district,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_core_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_user_core_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_user_core_proto_rawDescGZIP(), []int{2}
}

func (x *UserInfo) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *UserInfo) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserInfo) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserInfo) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *UserInfo) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *UserInfo) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *UserInfo) GetDistrict() string {
	if x != nil {
		return x.District
	}
	return ""
}

type UserRegisterReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base     *base.BaseReq `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	UserInfo *UserInfo     `protobuf:"bytes,2,opt,name=userInfo,proto3" json:"userInfo,omitempty"`
}

func (x *UserRegisterReq) Reset() {
	*x = UserRegisterReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_core_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRegisterReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRegisterReq) ProtoMessage() {}

func (x *UserRegisterReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_core_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRegisterReq.ProtoReflect.Descriptor instead.
func (*UserRegisterReq) Descriptor() ([]byte, []int) {
	return file_user_core_proto_rawDescGZIP(), []int{3}
}

func (x *UserRegisterReq) GetBase() *base.BaseReq {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *UserRegisterReq) GetUserInfo() *UserInfo {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

type UserRegisterRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *base.BaseRes `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
}

func (x *UserRegisterRes) Reset() {
	*x = UserRegisterRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_core_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRegisterRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRegisterRes) ProtoMessage() {}

func (x *UserRegisterRes) ProtoReflect() protoreflect.Message {
	mi := &file_user_core_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRegisterRes.ProtoReflect.Descriptor instead.
func (*UserRegisterRes) Descriptor() ([]byte, []int) {
	return file_user_core_proto_rawDescGZIP(), []int{4}
}

func (x *UserRegisterRes) GetBase() *base.BaseRes {
	if x != nil {
		return x.Base
	}
	return nil
}

type UserInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *base.BaseReq `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
}

func (x *UserInfoReq) Reset() {
	*x = UserInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_core_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoReq) ProtoMessage() {}

func (x *UserInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_core_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoReq.ProtoReflect.Descriptor instead.
func (*UserInfoReq) Descriptor() ([]byte, []int) {
	return file_user_core_proto_rawDescGZIP(), []int{5}
}

func (x *UserInfoReq) GetBase() *base.BaseReq {
	if x != nil {
		return x.Base
	}
	return nil
}

type UserInfoRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base     *base.BaseRes `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	UserInfo *UserInfo     `protobuf:"bytes,2,opt,name=userInfo,proto3" json:"userInfo,omitempty"`
}

func (x *UserInfoRes) Reset() {
	*x = UserInfoRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_core_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfoRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoRes) ProtoMessage() {}

func (x *UserInfoRes) ProtoReflect() protoreflect.Message {
	mi := &file_user_core_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoRes.ProtoReflect.Descriptor instead.
func (*UserInfoRes) Descriptor() ([]byte, []int) {
	return file_user_core_proto_rawDescGZIP(), []int{6}
}

func (x *UserInfoRes) GetBase() *base.BaseRes {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *UserInfoRes) GetUserInfo() *UserInfo {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

type UserAuthReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *base.BaseReq `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Role int32         `protobuf:"varint,2,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *UserAuthReq) Reset() {
	*x = UserAuthReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_core_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserAuthReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAuthReq) ProtoMessage() {}

func (x *UserAuthReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_core_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAuthReq.ProtoReflect.Descriptor instead.
func (*UserAuthReq) Descriptor() ([]byte, []int) {
	return file_user_core_proto_rawDescGZIP(), []int{7}
}

func (x *UserAuthReq) GetBase() *base.BaseReq {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *UserAuthReq) GetRole() int32 {
	if x != nil {
		return x.Role
	}
	return 0
}

type UserAuthRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *base.BaseRes `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
}

func (x *UserAuthRes) Reset() {
	*x = UserAuthRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_core_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserAuthRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAuthRes) ProtoMessage() {}

func (x *UserAuthRes) ProtoReflect() protoreflect.Message {
	mi := &file_user_core_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAuthRes.ProtoReflect.Descriptor instead.
func (*UserAuthRes) Descriptor() ([]byte, []int) {
	return file_user_core_proto_rawDescGZIP(), []int{8}
}

func (x *UserAuthRes) GetBase() *base.BaseRes {
	if x != nil {
		return x.Base
	}
	return nil
}

var File_user_core_proto protoreflect.FileDescriptor

var file_user_core_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x18, 0x6a, 0x6c, 0x75, 0x5f, 0x63, 0x6f, 0x77, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x69,
	0x6f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x1a, 0x0a, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x78, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x30, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6a, 0x6c, 0x75, 0x5f, 0x63, 0x6f, 0x77, 0x5f,
	0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x71, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x22, 0x56, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x73, 0x12, 0x30, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x6a, 0x6c, 0x75, 0x5f, 0x63, 0x6f, 0x77, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f,
	0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x52, 0x04, 0x62,
	0x61, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xb4, 0x01, 0x0a, 0x08, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x72, 0x6f, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x63, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74,
	0x22, 0x83, 0x01, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x12, 0x30, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6a, 0x6c, 0x75, 0x5f, 0x63, 0x6f, 0x77, 0x5f, 0x73, 0x74, 0x75,
	0x64, 0x69, 0x6f, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71,
	0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6a, 0x6c, 0x75, 0x5f, 0x63,
	0x6f, 0x77, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x43, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x04, 0x62, 0x61, 0x73,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6a, 0x6c, 0x75, 0x5f, 0x63, 0x6f,
	0x77, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x42, 0x61,
	0x73, 0x65, 0x52, 0x65, 0x73, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x22, 0x3f, 0x0a, 0x0b, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x30, 0x0a, 0x04, 0x62, 0x61,
	0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6a, 0x6c, 0x75, 0x5f, 0x63,
	0x6f, 0x77, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x42,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x22, 0x7f, 0x0a, 0x0b,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x04, 0x62,
	0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6a, 0x6c, 0x75, 0x5f,
	0x63, 0x6f, 0x77, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x3e, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x6a, 0x6c, 0x75, 0x5f, 0x63, 0x6f, 0x77, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x53, 0x0a,
	0x0b, 0x55, 0x73, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x12, 0x30, 0x0a, 0x04,
	0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6a, 0x6c, 0x75,
	0x5f, 0x63, 0x6f, 0x77, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x72, 0x6f,
	0x6c, 0x65, 0x22, 0x3f, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65,
	0x73, 0x12, 0x30, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x6a, 0x6c, 0x75, 0x5f, 0x63, 0x6f, 0x77, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f,
	0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x52, 0x04, 0x62,
	0x61, 0x73, 0x65, 0x32, 0x90, 0x03, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x72, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5d, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x12, 0x26, 0x2e, 0x6a, 0x6c, 0x75, 0x5f, 0x63, 0x6f, 0x77, 0x5f, 0x73,
	0x74, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x26, 0x2e, 0x6a,
	0x6c, 0x75, 0x5f, 0x63, 0x6f, 0x77, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x66, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x29, 0x2e, 0x6a, 0x6c, 0x75, 0x5f, 0x63, 0x6f, 0x77,
	0x5f, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x1a, 0x29, 0x2e, 0x6a, 0x6c, 0x75, 0x5f, 0x63, 0x6f, 0x77, 0x5f, 0x73, 0x74, 0x75, 0x64,
	0x69, 0x6f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x5a,
	0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x25, 0x2e, 0x6a, 0x6c, 0x75,
	0x5f, 0x63, 0x6f, 0x77, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x71, 0x1a, 0x25, 0x2e, 0x6a, 0x6c, 0x75, 0x5f, 0x63, 0x6f, 0x77, 0x5f, 0x73, 0x74, 0x75, 0x64,
	0x69, 0x6f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x5a, 0x0a, 0x08, 0x55, 0x73,
	0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x12, 0x25, 0x2e, 0x6a, 0x6c, 0x75, 0x5f, 0x63, 0x6f, 0x77,
	0x5f, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x1a, 0x25, 0x2e,
	0x6a, 0x6c, 0x75, 0x5f, 0x63, 0x6f, 0x77, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x75, 0x74,
	0x68, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x6c, 0x75, 0x2d, 0x63, 0x6f, 0x77, 0x2d, 0x73, 0x74, 0x75,
	0x64, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_core_proto_rawDescOnce sync.Once
	file_user_core_proto_rawDescData = file_user_core_proto_rawDesc
)

func file_user_core_proto_rawDescGZIP() []byte {
	file_user_core_proto_rawDescOnce.Do(func() {
		file_user_core_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_core_proto_rawDescData)
	})
	return file_user_core_proto_rawDescData
}

var file_user_core_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_user_core_proto_goTypes = []interface{}{
	(*UserLoginReq)(nil),    // 0: jlu_cow_studio.user_core.UserLoginReq
	(*UserLoginRes)(nil),    // 1: jlu_cow_studio.user_core.UserLoginRes
	(*UserInfo)(nil),        // 2: jlu_cow_studio.user_core.UserInfo
	(*UserRegisterReq)(nil), // 3: jlu_cow_studio.user_core.UserRegisterReq
	(*UserRegisterRes)(nil), // 4: jlu_cow_studio.user_core.UserRegisterRes
	(*UserInfoReq)(nil),     // 5: jlu_cow_studio.user_core.UserInfoReq
	(*UserInfoRes)(nil),     // 6: jlu_cow_studio.user_core.UserInfoRes
	(*UserAuthReq)(nil),     // 7: jlu_cow_studio.user_core.UserAuthReq
	(*UserAuthRes)(nil),     // 8: jlu_cow_studio.user_core.UserAuthRes
	(*base.BaseReq)(nil),    // 9: jlu_cow_studio.base.BaseReq
	(*base.BaseRes)(nil),    // 10: jlu_cow_studio.base.BaseRes
}
var file_user_core_proto_depIdxs = []int32{
	9,  // 0: jlu_cow_studio.user_core.UserLoginReq.base:type_name -> jlu_cow_studio.base.BaseReq
	10, // 1: jlu_cow_studio.user_core.UserLoginRes.base:type_name -> jlu_cow_studio.base.BaseRes
	9,  // 2: jlu_cow_studio.user_core.UserRegisterReq.base:type_name -> jlu_cow_studio.base.BaseReq
	2,  // 3: jlu_cow_studio.user_core.UserRegisterReq.userInfo:type_name -> jlu_cow_studio.user_core.UserInfo
	10, // 4: jlu_cow_studio.user_core.UserRegisterRes.base:type_name -> jlu_cow_studio.base.BaseRes
	9,  // 5: jlu_cow_studio.user_core.UserInfoReq.base:type_name -> jlu_cow_studio.base.BaseReq
	10, // 6: jlu_cow_studio.user_core.UserInfoRes.base:type_name -> jlu_cow_studio.base.BaseRes
	2,  // 7: jlu_cow_studio.user_core.UserInfoRes.userInfo:type_name -> jlu_cow_studio.user_core.UserInfo
	9,  // 8: jlu_cow_studio.user_core.UserAuthReq.base:type_name -> jlu_cow_studio.base.BaseReq
	10, // 9: jlu_cow_studio.user_core.UserAuthRes.base:type_name -> jlu_cow_studio.base.BaseRes
	0,  // 10: jlu_cow_studio.user_core.UserCoreService.UserLogin:input_type -> jlu_cow_studio.user_core.UserLoginReq
	3,  // 11: jlu_cow_studio.user_core.UserCoreService.UserRegister:input_type -> jlu_cow_studio.user_core.UserRegisterReq
	5,  // 12: jlu_cow_studio.user_core.UserCoreService.UserInfo:input_type -> jlu_cow_studio.user_core.UserInfoReq
	7,  // 13: jlu_cow_studio.user_core.UserCoreService.UserAuth:input_type -> jlu_cow_studio.user_core.UserAuthReq
	1,  // 14: jlu_cow_studio.user_core.UserCoreService.UserLogin:output_type -> jlu_cow_studio.user_core.UserLoginRes
	4,  // 15: jlu_cow_studio.user_core.UserCoreService.UserRegister:output_type -> jlu_cow_studio.user_core.UserRegisterRes
	6,  // 16: jlu_cow_studio.user_core.UserCoreService.UserInfo:output_type -> jlu_cow_studio.user_core.UserInfoRes
	8,  // 17: jlu_cow_studio.user_core.UserCoreService.UserAuth:output_type -> jlu_cow_studio.user_core.UserAuthRes
	14, // [14:18] is the sub-list for method output_type
	10, // [10:14] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_user_core_proto_init() }
func file_user_core_proto_init() {
	if File_user_core_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_core_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserLoginReq); i {
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
		file_user_core_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserLoginRes); i {
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
		file_user_core_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
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
		file_user_core_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRegisterReq); i {
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
		file_user_core_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRegisterRes); i {
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
		file_user_core_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfoReq); i {
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
		file_user_core_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfoRes); i {
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
		file_user_core_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserAuthReq); i {
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
		file_user_core_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserAuthRes); i {
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
			RawDescriptor: file_user_core_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_core_proto_goTypes,
		DependencyIndexes: file_user_core_proto_depIdxs,
		MessageInfos:      file_user_core_proto_msgTypes,
	}.Build()
	File_user_core_proto = out.File
	file_user_core_proto_rawDesc = nil
	file_user_core_proto_goTypes = nil
	file_user_core_proto_depIdxs = nil
}
