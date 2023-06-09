// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: trade_core.proto

package trade_core

import (
	base "github.com/jlu-cow-studio/common/dal/rpc/base"
	product_core "github.com/jlu-cow-studio/common/dal/rpc/product_core"
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

type OrderForList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId    int64                  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ItemId    int64                  `protobuf:"varint,3,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	Quantity  int32                  `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
	CreatedAt int64                  `protobuf:"varint,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt int64                  `protobuf:"varint,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	ItemInfo  *product_core.ItemInfo `protobuf:"bytes,7,opt,name=item_info,json=itemInfo,proto3" json:"item_info,omitempty"`
}

func (x *OrderForList) Reset() {
	*x = OrderForList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trade_core_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderForList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderForList) ProtoMessage() {}

func (x *OrderForList) ProtoReflect() protoreflect.Message {
	mi := &file_trade_core_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderForList.ProtoReflect.Descriptor instead.
func (*OrderForList) Descriptor() ([]byte, []int) {
	return file_trade_core_proto_rawDescGZIP(), []int{0}
}

func (x *OrderForList) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OrderForList) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *OrderForList) GetItemId() int64 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *OrderForList) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *OrderForList) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *OrderForList) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *OrderForList) GetItemInfo() *product_core.ItemInfo {
	if x != nil {
		return x.ItemInfo
	}
	return nil
}

type RechargeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base  *base.BaseReq `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Money float64       `protobuf:"fixed64,2,opt,name=money,proto3" json:"money,omitempty"`
}

func (x *RechargeRequest) Reset() {
	*x = RechargeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trade_core_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RechargeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RechargeRequest) ProtoMessage() {}

func (x *RechargeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trade_core_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RechargeRequest.ProtoReflect.Descriptor instead.
func (*RechargeRequest) Descriptor() ([]byte, []int) {
	return file_trade_core_proto_rawDescGZIP(), []int{1}
}

func (x *RechargeRequest) GetBase() *base.BaseReq {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *RechargeRequest) GetMoney() float64 {
	if x != nil {
		return x.Money
	}
	return 0
}

type RechargeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *base.BaseRes `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
}

func (x *RechargeResponse) Reset() {
	*x = RechargeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trade_core_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RechargeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RechargeResponse) ProtoMessage() {}

func (x *RechargeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_trade_core_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RechargeResponse.ProtoReflect.Descriptor instead.
func (*RechargeResponse) Descriptor() ([]byte, []int) {
	return file_trade_core_proto_rawDescGZIP(), []int{2}
}

func (x *RechargeResponse) GetBase() *base.BaseRes {
	if x != nil {
		return x.Base
	}
	return nil
}

type OrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base   *base.BaseReq `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	ItemId int64         `protobuf:"varint,2,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	Count  int32         `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *OrderRequest) Reset() {
	*x = OrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trade_core_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderRequest) ProtoMessage() {}

func (x *OrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trade_core_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderRequest.ProtoReflect.Descriptor instead.
func (*OrderRequest) Descriptor() ([]byte, []int) {
	return file_trade_core_proto_rawDescGZIP(), []int{3}
}

func (x *OrderRequest) GetBase() *base.BaseReq {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *OrderRequest) GetItemId() int64 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *OrderRequest) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type OrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base *base.BaseRes `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
}

func (x *OrderResponse) Reset() {
	*x = OrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trade_core_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderResponse) ProtoMessage() {}

func (x *OrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_trade_core_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderResponse.ProtoReflect.Descriptor instead.
func (*OrderResponse) Descriptor() ([]byte, []int) {
	return file_trade_core_proto_rawDescGZIP(), []int{4}
}

func (x *OrderResponse) GetBase() *base.BaseRes {
	if x != nil {
		return x.Base
	}
	return nil
}

type OrderListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base     *base.BaseReq `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	UserId   int64         `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Page     int32         `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int32         `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *OrderListRequest) Reset() {
	*x = OrderListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trade_core_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderListRequest) ProtoMessage() {}

func (x *OrderListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trade_core_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderListRequest.ProtoReflect.Descriptor instead.
func (*OrderListRequest) Descriptor() ([]byte, []int) {
	return file_trade_core_proto_rawDescGZIP(), []int{5}
}

func (x *OrderListRequest) GetBase() *base.BaseReq {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *OrderListRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *OrderListRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *OrderListRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type OrderListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base      *base.BaseRes   `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	OrderList []*OrderForList `protobuf:"bytes,2,rep,name=order_list,json=orderList,proto3" json:"order_list,omitempty"`
}

func (x *OrderListResponse) Reset() {
	*x = OrderListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trade_core_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderListResponse) ProtoMessage() {}

func (x *OrderListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_trade_core_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderListResponse.ProtoReflect.Descriptor instead.
func (*OrderListResponse) Descriptor() ([]byte, []int) {
	return file_trade_core_proto_rawDescGZIP(), []int{6}
}

func (x *OrderListResponse) GetBase() *base.BaseRes {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *OrderListResponse) GetOrderList() []*OrderForList {
	if x != nil {
		return x.OrderList
	}
	return nil
}

var File_trade_core_proto protoreflect.FileDescriptor

var file_trade_core_proto_rawDesc = []byte{
	0x0a, 0x10, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x19, 0x6a, 0x6c, 0x75, 0x5f, 0x63, 0x6f, 0x77, 0x5f, 0x73, 0x74, 0x75, 0x64,
	0x69, 0x6f, 0x2e, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x1a, 0x0a, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xee, 0x01,
	0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x42, 0x0a, 0x09, 0x69, 0x74,
	0x65, 0x6d, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e,
	0x6a, 0x6c, 0x75, 0x5f, 0x63, 0x6f, 0x77, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x49, 0x74, 0x65, 0x6d,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x59,
	0x0a, 0x0f, 0x52, 0x65, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x30, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x6a, 0x6c, 0x75, 0x5f, 0x63, 0x6f, 0x77, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f,
	0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x52, 0x04, 0x62,
	0x61, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x22, 0x44, 0x0a, 0x10, 0x52, 0x65, 0x63,
	0x68, 0x61, 0x72, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a,
	0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6a, 0x6c,
	0x75, 0x5f, 0x63, 0x6f, 0x77, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x22,
	0x6f, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x30, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x6a, 0x6c, 0x75, 0x5f, 0x63, 0x6f, 0x77, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x52, 0x04, 0x62, 0x61, 0x73,
	0x65, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0x41, 0x0a, 0x0d, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x30, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x6a, 0x6c, 0x75, 0x5f, 0x63, 0x6f, 0x77, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f,
	0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x52, 0x04, 0x62,
	0x61, 0x73, 0x65, 0x22, 0x8e, 0x01, 0x0a, 0x10, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6a, 0x6c, 0x75, 0x5f, 0x63, 0x6f, 0x77,
	0x5f, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x42, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x71, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x22, 0x8d, 0x01, 0x0a, 0x11, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x62, 0x61,
	0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6a, 0x6c, 0x75, 0x5f, 0x63,
	0x6f, 0x77, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x42,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x0a,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x6a, 0x6c, 0x75, 0x5f, 0x63, 0x6f, 0x77, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x69,
	0x6f, 0x2e, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x46, 0x6f, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x32, 0xbb, 0x02, 0x0a, 0x10, 0x54, 0x72, 0x61, 0x64, 0x65, 0x43, 0x6f,
	0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x63, 0x0a, 0x08, 0x52, 0x65, 0x63,
	0x68, 0x61, 0x72, 0x67, 0x65, 0x12, 0x2a, 0x2e, 0x6a, 0x6c, 0x75, 0x5f, 0x63, 0x6f, 0x77, 0x5f,
	0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x52, 0x65, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2b, 0x2e, 0x6a, 0x6c, 0x75, 0x5f, 0x63, 0x6f, 0x77, 0x5f, 0x73, 0x74, 0x75, 0x64,
	0x69, 0x6f, 0x2e, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x65,
	0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a,
	0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x27, 0x2e, 0x6a, 0x6c, 0x75, 0x5f, 0x63, 0x6f,
	0x77, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x28, 0x2e, 0x6a, 0x6c, 0x75, 0x5f, 0x63, 0x6f, 0x77, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x69,
	0x6f, 0x2e, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x66, 0x0a, 0x09, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2b, 0x2e, 0x6a, 0x6c, 0x75, 0x5f, 0x63, 0x6f,
	0x77, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x6a, 0x6c, 0x75, 0x5f, 0x63, 0x6f, 0x77, 0x5f, 0x73,
	0x74, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6a, 0x6c, 0x75, 0x2d, 0x63, 0x6f, 0x77, 0x2d, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x64, 0x61, 0x6c, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x74,
	0x72, 0x61, 0x64, 0x65, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_trade_core_proto_rawDescOnce sync.Once
	file_trade_core_proto_rawDescData = file_trade_core_proto_rawDesc
)

func file_trade_core_proto_rawDescGZIP() []byte {
	file_trade_core_proto_rawDescOnce.Do(func() {
		file_trade_core_proto_rawDescData = protoimpl.X.CompressGZIP(file_trade_core_proto_rawDescData)
	})
	return file_trade_core_proto_rawDescData
}

var file_trade_core_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_trade_core_proto_goTypes = []interface{}{
	(*OrderForList)(nil),          // 0: jlu_cow_studio.trade_core.OrderForList
	(*RechargeRequest)(nil),       // 1: jlu_cow_studio.trade_core.RechargeRequest
	(*RechargeResponse)(nil),      // 2: jlu_cow_studio.trade_core.RechargeResponse
	(*OrderRequest)(nil),          // 3: jlu_cow_studio.trade_core.OrderRequest
	(*OrderResponse)(nil),         // 4: jlu_cow_studio.trade_core.OrderResponse
	(*OrderListRequest)(nil),      // 5: jlu_cow_studio.trade_core.OrderListRequest
	(*OrderListResponse)(nil),     // 6: jlu_cow_studio.trade_core.OrderListResponse
	(*product_core.ItemInfo)(nil), // 7: jlu_cow_studio.product_core.ItemInfo
	(*base.BaseReq)(nil),          // 8: jlu_cow_studio.base.BaseReq
	(*base.BaseRes)(nil),          // 9: jlu_cow_studio.base.BaseRes
}
var file_trade_core_proto_depIdxs = []int32{
	7,  // 0: jlu_cow_studio.trade_core.OrderForList.item_info:type_name -> jlu_cow_studio.product_core.ItemInfo
	8,  // 1: jlu_cow_studio.trade_core.RechargeRequest.base:type_name -> jlu_cow_studio.base.BaseReq
	9,  // 2: jlu_cow_studio.trade_core.RechargeResponse.base:type_name -> jlu_cow_studio.base.BaseRes
	8,  // 3: jlu_cow_studio.trade_core.OrderRequest.base:type_name -> jlu_cow_studio.base.BaseReq
	9,  // 4: jlu_cow_studio.trade_core.OrderResponse.base:type_name -> jlu_cow_studio.base.BaseRes
	8,  // 5: jlu_cow_studio.trade_core.OrderListRequest.base:type_name -> jlu_cow_studio.base.BaseReq
	9,  // 6: jlu_cow_studio.trade_core.OrderListResponse.base:type_name -> jlu_cow_studio.base.BaseRes
	0,  // 7: jlu_cow_studio.trade_core.OrderListResponse.order_list:type_name -> jlu_cow_studio.trade_core.OrderForList
	1,  // 8: jlu_cow_studio.trade_core.TradeCoreService.Recharge:input_type -> jlu_cow_studio.trade_core.RechargeRequest
	3,  // 9: jlu_cow_studio.trade_core.TradeCoreService.Order:input_type -> jlu_cow_studio.trade_core.OrderRequest
	5,  // 10: jlu_cow_studio.trade_core.TradeCoreService.OrderList:input_type -> jlu_cow_studio.trade_core.OrderListRequest
	2,  // 11: jlu_cow_studio.trade_core.TradeCoreService.Recharge:output_type -> jlu_cow_studio.trade_core.RechargeResponse
	4,  // 12: jlu_cow_studio.trade_core.TradeCoreService.Order:output_type -> jlu_cow_studio.trade_core.OrderResponse
	6,  // 13: jlu_cow_studio.trade_core.TradeCoreService.OrderList:output_type -> jlu_cow_studio.trade_core.OrderListResponse
	11, // [11:14] is the sub-list for method output_type
	8,  // [8:11] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_trade_core_proto_init() }
func file_trade_core_proto_init() {
	if File_trade_core_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_trade_core_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderForList); i {
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
		file_trade_core_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RechargeRequest); i {
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
		file_trade_core_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RechargeResponse); i {
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
		file_trade_core_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderRequest); i {
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
		file_trade_core_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderResponse); i {
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
		file_trade_core_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderListRequest); i {
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
		file_trade_core_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderListResponse); i {
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
			RawDescriptor: file_trade_core_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_trade_core_proto_goTypes,
		DependencyIndexes: file_trade_core_proto_depIdxs,
		MessageInfos:      file_trade_core_proto_msgTypes,
	}.Build()
	File_trade_core_proto = out.File
	file_trade_core_proto_rawDesc = nil
	file_trade_core_proto_goTypes = nil
	file_trade_core_proto_depIdxs = nil
}
