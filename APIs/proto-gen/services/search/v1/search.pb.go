// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: search.proto

package search

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "github.com/mar-coding/SearchEngineWrapper/APIs/proto-gen/components/list/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Permissions int32

const (
	Permissions_NONE           Permissions = 0
	Permissions_LIST_ITEM_POST Permissions = 1
)

// Enum value maps for Permissions.
var (
	Permissions_name = map[int32]string{
		0: "NONE",
		1: "LIST_ITEM_POST",
	}
	Permissions_value = map[string]int32{
		"NONE":           0,
		"LIST_ITEM_POST": 1,
	}
)

func (x Permissions) Enum() *Permissions {
	p := new(Permissions)
	*p = x
	return p
}

func (x Permissions) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Permissions) Descriptor() protoreflect.EnumDescriptor {
	return file_search_proto_enumTypes[0].Descriptor()
}

func (Permissions) Type() protoreflect.EnumType {
	return &file_search_proto_enumTypes[0]
}

func (x Permissions) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Permissions.Descriptor instead.
func (Permissions) EnumDescriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{0}
}

type Permission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Permissions         []Permissions `protobuf:"varint,1,rep,packed,name=permissions,proto3,enum=services.search.v1.Permissions" json:"permissions,omitempty"`
	Optional            bool          `protobuf:"varint,2,opt,name=optional,proto3" json:"optional,omitempty"`
	ValidatePermissions bool          `protobuf:"varint,3,opt,name=validate_permissions,json=validatePermissions,proto3" json:"validate_permissions,omitempty"`
	Captcha             bool          `protobuf:"varint,4,opt,name=captcha,proto3" json:"captcha,omitempty"`
}

func (x *Permission) Reset() {
	*x = Permission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Permission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Permission) ProtoMessage() {}

func (x *Permission) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Permission.ProtoReflect.Descriptor instead.
func (*Permission) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{0}
}

func (x *Permission) GetPermissions() []Permissions {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *Permission) GetOptional() bool {
	if x != nil {
		return x.Optional
	}
	return false
}

func (x *Permission) GetValidatePermissions() bool {
	if x != nil {
		return x.ValidatePermissions
	}
	return false
}

func (x *Permission) GetCaptcha() bool {
	if x != nil {
		return x.Captcha
	}
	return false
}

type ListItemsSearchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items           []*Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	PageSize        int32   `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageNo          int32   `protobuf:"varint,3,opt,name=page_no,json=pageNo,proto3" json:"page_no,omitempty"`
	TotalItemsCount int32   `protobuf:"varint,4,opt,name=total_items_count,json=totalItemsCount,proto3" json:"total_items_count,omitempty"`
}

func (x *ListItemsSearchResponse) Reset() {
	*x = ListItemsSearchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListItemsSearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListItemsSearchResponse) ProtoMessage() {}

func (x *ListItemsSearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListItemsSearchResponse.ProtoReflect.Descriptor instead.
func (*ListItemsSearchResponse) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{1}
}

func (x *ListItemsSearchResponse) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListItemsSearchResponse) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListItemsSearchResponse) GetPageNo() int32 {
	if x != nil {
		return x.PageNo
	}
	return 0
}

func (x *ListItemsSearchResponse) GetTotalItemsCount() int32 {
	if x != nil {
		return x.TotalItemsCount
	}
	return 0
}

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Url         string `protobuf:"bytes,10,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{2}
}

func (x *Item) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Item) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Item) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Item) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type ListItemsSearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// required
	Q        string `protobuf:"bytes,1,opt,name=q,proto3" json:"q,omitempty"`
	PageSize int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageNo   int32  `protobuf:"varint,3,opt,name=page_no,json=pageNo,proto3" json:"page_no,omitempty"`
}

func (x *ListItemsSearchRequest) Reset() {
	*x = ListItemsSearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListItemsSearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListItemsSearchRequest) ProtoMessage() {}

func (x *ListItemsSearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListItemsSearchRequest.ProtoReflect.Descriptor instead.
func (*ListItemsSearchRequest) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{3}
}

func (x *ListItemsSearchRequest) GetQ() string {
	if x != nil {
		return x.Q
	}
	return ""
}

func (x *ListItemsSearchRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListItemsSearchRequest) GetPageNo() int32 {
	if x != nil {
		return x.PageNo
	}
	return 0
}

var file_search_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.MethodOptions)(nil),
		ExtensionType: (*Permission)(nil),
		Field:         1000001,
		Name:          "services.search.v1.permission",
		Tag:           "bytes,1000001,opt,name=permission",
		Filename:      "search.proto",
	},
}

// Extension fields to descriptor.MethodOptions.
var (
	// optional services.search.v1.Permission permission = 1000001;
	E_Permission = &file_search_proto_extTypes[0]
)

var File_search_proto protoreflect.FileDescriptor

var file_search_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e,
	0x76, 0x31, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f,
	0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x76,
	0x31, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb8, 0x01, 0x0a, 0x0a, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x41, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x12, 0x31, 0x0a, 0x14, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x13, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68,
	0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61,
	0x22, 0xab, 0x01, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x76, 0x31,
	0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x6e, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x61, 0x67, 0x65,
	0x4e, 0x6f, 0x12, 0x2a, 0x0a, 0x11, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x84,
	0x01, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x4a, 0x04, 0x08, 0x04, 0x10, 0x05, 0x4a, 0x04, 0x08, 0x05, 0x10, 0x06, 0x4a, 0x04, 0x08,
	0x06, 0x10, 0x07, 0x4a, 0x04, 0x08, 0x07, 0x10, 0x08, 0x4a, 0x04, 0x08, 0x08, 0x10, 0x09, 0x4a,
	0x04, 0x08, 0x09, 0x10, 0x0a, 0x22, 0x65, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x15, 0x0a, 0x01, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72,
	0x02, 0x5a, 0x00, 0x52, 0x01, 0x71, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x6f, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x6f, 0x2a, 0x2b, 0x0a, 0x0b,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x4e,
	0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x4c, 0x49, 0x53, 0x54, 0x5f, 0x49, 0x54,
	0x45, 0x4d, 0x5f, 0x50, 0x4f, 0x53, 0x54, 0x10, 0x01, 0x32, 0xbf, 0x01, 0x0a, 0x0d, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xad, 0x01, 0x0a, 0x0f,
	0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12,
	0x2a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x41, 0x92, 0x41, 0x27, 0x12, 0x25, 0x67,
	0x65, 0x74, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x20, 0x66, 0x6f, 0x72, 0x20, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x20, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x8a, 0xa4, 0xe8, 0x03, 0x03, 0x0a, 0x01, 0x00, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x09, 0x12, 0x07, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x3a, 0x60, 0x0a, 0x0a, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xc1, 0x84, 0x3d, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0xaa, 0x02,
	0x92, 0x41, 0xd2, 0x01, 0x12, 0x4e, 0x0a, 0x0e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x20, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x37, 0x0a, 0x13, 0x4d, 0x6f, 0x68, 0x61, 0x6d, 0x6d,
	0x61, 0x64, 0x41, 0x6d, 0x69, 0x6e, 0x20, 0x52, 0x61, 0x68, 0x69, 0x6d, 0x69, 0x12, 0x09, 0x31,
	0x32, 0x37, 0x2e, 0x30, 0x2e, 0x30, 0x2e, 0x31, 0x1a, 0x15, 0x6d, 0x61, 0x72, 0x63, 0x6f, 0x64,
	0x69, 0x6e, 0x67, 0x37, 0x38, 0x40, 0x67, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x32,
	0x03, 0x31, 0x2e, 0x30, 0x2a, 0x02, 0x02, 0x01, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x5a, 0x4a, 0x0a, 0x48,
	0x0a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x3e, 0x08, 0x02, 0x12, 0x29, 0x41, 0x75,
	0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x65, 0x64,
	0x20, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x1a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x02, 0x62, 0x0c, 0x0a, 0x0a, 0x0a, 0x06, 0x62, 0x65,
	0x61, 0x72, 0x65, 0x72, 0x12, 0x00, 0x5a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x72, 0x2d, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x2f, 0x41, 0x50, 0x49, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x65, 0x6e,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x2f, 0x76, 0x31, 0x3b, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_search_proto_rawDescOnce sync.Once
	file_search_proto_rawDescData = file_search_proto_rawDesc
)

func file_search_proto_rawDescGZIP() []byte {
	file_search_proto_rawDescOnce.Do(func() {
		file_search_proto_rawDescData = protoimpl.X.CompressGZIP(file_search_proto_rawDescData)
	})
	return file_search_proto_rawDescData
}

var file_search_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_search_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_search_proto_goTypes = []interface{}{
	(Permissions)(0),                 // 0: services.search.v1.Permissions
	(*Permission)(nil),               // 1: services.search.v1.Permission
	(*ListItemsSearchResponse)(nil),  // 2: services.search.v1.ListItemsSearchResponse
	(*Item)(nil),                     // 3: services.search.v1.Item
	(*ListItemsSearchRequest)(nil),   // 4: services.search.v1.ListItemsSearchRequest
	(*descriptor.MethodOptions)(nil), // 5: google.protobuf.MethodOptions
}
var file_search_proto_depIdxs = []int32{
	0, // 0: services.search.v1.Permission.permissions:type_name -> services.search.v1.Permissions
	3, // 1: services.search.v1.ListItemsSearchResponse.items:type_name -> services.search.v1.Item
	5, // 2: services.search.v1.permission:extendee -> google.protobuf.MethodOptions
	1, // 3: services.search.v1.permission:type_name -> services.search.v1.Permission
	4, // 4: services.search.v1.SearchService.ListItemsSearch:input_type -> services.search.v1.ListItemsSearchRequest
	2, // 5: services.search.v1.SearchService.ListItemsSearch:output_type -> services.search.v1.ListItemsSearchResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	3, // [3:4] is the sub-list for extension type_name
	2, // [2:3] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_search_proto_init() }
func file_search_proto_init() {
	if File_search_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_search_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Permission); i {
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
		file_search_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListItemsSearchResponse); i {
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
		file_search_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
		file_search_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListItemsSearchRequest); i {
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
			RawDescriptor: file_search_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 1,
			NumServices:   1,
		},
		GoTypes:           file_search_proto_goTypes,
		DependencyIndexes: file_search_proto_depIdxs,
		EnumInfos:         file_search_proto_enumTypes,
		MessageInfos:      file_search_proto_msgTypes,
		ExtensionInfos:    file_search_proto_extTypes,
	}.Build()
	File_search_proto = out.File
	file_search_proto_rawDesc = nil
	file_search_proto_goTypes = nil
	file_search_proto_depIdxs = nil
}