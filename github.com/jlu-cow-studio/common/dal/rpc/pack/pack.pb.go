// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: pack.proto

package pack

import (
	base "github.com/jlu-cow-studio/common/dal/rpc/base"
	product_core "github.com/jlu-cow-studio/common/dal/rpc/product-core"
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

type PackItemsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base       *base.BaseReq `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	ItemIdList []int32       `protobuf:"varint,2,rep,packed,name=item_id_list,json=itemIdList,proto3" json:"item_id_list,omitempty"`
}

func (x *PackItemsReq) Reset() {
	*x = PackItemsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pack_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PackItemsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PackItemsReq) ProtoMessage() {}

func (x *PackItemsReq) ProtoReflect() protoreflect.Message {
	mi := &file_pack_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PackItemsReq.ProtoReflect.Descriptor instead.
func (*PackItemsReq) Descriptor() ([]byte, []int) {
	return file_pack_proto_rawDescGZIP(), []int{0}
}

func (x *PackItemsReq) GetBase() *base.BaseReq {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *PackItemsReq) GetItemIdList() []int32 {
	if x != nil {
		return x.ItemIdList
	}
	return nil
}

type PackItemsRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Base     *base.BaseRes            `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	ItemList []*product_core.ItemInfo `protobuf:"bytes,2,rep,name=item_list,json=itemList,proto3" json:"item_list,omitempty"`
}

func (x *PackItemsRes) Reset() {
	*x = PackItemsRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pack_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PackItemsRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PackItemsRes) ProtoMessage() {}

func (x *PackItemsRes) ProtoReflect() protoreflect.Message {
	mi := &file_pack_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PackItemsRes.ProtoReflect.Descriptor instead.
func (*PackItemsRes) Descriptor() ([]byte, []int) {
	return file_pack_proto_rawDescGZIP(), []int{1}
}

func (x *PackItemsRes) GetBase() *base.BaseRes {
	if x != nil {
		return x.Base
	}
	return nil
}

func (x *PackItemsRes) GetItemList() []*product_core.ItemInfo {
	if x != nil {
		return x.ItemList
	}
	return nil
}

var File_pack_proto protoreflect.FileDescriptor

var file_pack_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x70, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x6a, 0x6c,
	0x75, 0x5f, 0x63, 0x6f, 0x77, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x70, 0x61, 0x63,
	0x6b, 0x1a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x62, 0x0a, 0x0c, 0x50, 0x61, 0x63, 0x6b, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65,
	0x71, 0x12, 0x30, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x6a, 0x6c, 0x75, 0x5f, 0x63, 0x6f, 0x77, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f,
	0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x52, 0x04, 0x62,
	0x61, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0c, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0a, 0x69, 0x74, 0x65, 0x6d, 0x49,
	0x64, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x84, 0x01, 0x0a, 0x0c, 0x50, 0x61, 0x63, 0x6b, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6a, 0x6c, 0x75, 0x5f, 0x63, 0x6f, 0x77, 0x5f, 0x73,
	0x74, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52,
	0x65, 0x73, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x09, 0x69, 0x74, 0x65, 0x6d,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6a, 0x6c,
	0x75, 0x5f, 0x63, 0x6f, 0x77, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x32, 0x62, 0x0a, 0x0b,
	0x50, 0x61, 0x63, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53, 0x0a, 0x09, 0x50,
	0x61, 0x63, 0x6b, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x21, 0x2e, 0x6a, 0x6c, 0x75, 0x5f, 0x63,
	0x6f, 0x77, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x2e, 0x50,
	0x61, 0x63, 0x6b, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x21, 0x2e, 0x6a, 0x6c,
	0x75, 0x5f, 0x63, 0x6f, 0x77, 0x5f, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x70, 0x61, 0x63,
	0x6b, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x22, 0x00,
	0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a,
	0x6c, 0x75, 0x2d, 0x63, 0x6f, 0x77, 0x2d, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x64, 0x61, 0x6c, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x61, 0x63,
	0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pack_proto_rawDescOnce sync.Once
	file_pack_proto_rawDescData = file_pack_proto_rawDesc
)

func file_pack_proto_rawDescGZIP() []byte {
	file_pack_proto_rawDescOnce.Do(func() {
		file_pack_proto_rawDescData = protoimpl.X.CompressGZIP(file_pack_proto_rawDescData)
	})
	return file_pack_proto_rawDescData
}

var file_pack_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pack_proto_goTypes = []interface{}{
	(*PackItemsReq)(nil),          // 0: jlu_cow_studio.pack.PackItemsReq
	(*PackItemsRes)(nil),          // 1: jlu_cow_studio.pack.PackItemsRes
	(*base.BaseReq)(nil),          // 2: jlu_cow_studio.base.BaseReq
	(*base.BaseRes)(nil),          // 3: jlu_cow_studio.base.BaseRes
	(*product_core.ItemInfo)(nil), // 4: jlu_cow_studio.product_core.ItemInfo
}
var file_pack_proto_depIdxs = []int32{
	2, // 0: jlu_cow_studio.pack.PackItemsReq.base:type_name -> jlu_cow_studio.base.BaseReq
	3, // 1: jlu_cow_studio.pack.PackItemsRes.base:type_name -> jlu_cow_studio.base.BaseRes
	4, // 2: jlu_cow_studio.pack.PackItemsRes.item_list:type_name -> jlu_cow_studio.product_core.ItemInfo
	0, // 3: jlu_cow_studio.pack.PackService.PackItems:input_type -> jlu_cow_studio.pack.PackItemsReq
	1, // 4: jlu_cow_studio.pack.PackService.PackItems:output_type -> jlu_cow_studio.pack.PackItemsRes
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pack_proto_init() }
func file_pack_proto_init() {
	if File_pack_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pack_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PackItemsReq); i {
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
		file_pack_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PackItemsRes); i {
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
			RawDescriptor: file_pack_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pack_proto_goTypes,
		DependencyIndexes: file_pack_proto_depIdxs,
		MessageInfos:      file_pack_proto_msgTypes,
	}.Build()
	File_pack_proto = out.File
	file_pack_proto_rawDesc = nil
	file_pack_proto_goTypes = nil
	file_pack_proto_depIdxs = nil
}