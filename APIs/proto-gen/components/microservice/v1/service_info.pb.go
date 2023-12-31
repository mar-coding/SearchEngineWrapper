// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: microservice/v1/service_info.proto

package microservice

import (
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/golang/protobuf/ptypes/struct"
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
	Permissions_NONE Permissions = 0
)

// Enum value maps for Permissions.
var (
	Permissions_name = map[int32]string{
		0: "NONE",
	}
	Permissions_value = map[string]int32{
		"NONE": 0,
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
	return file_microservice_v1_service_info_proto_enumTypes[0].Descriptor()
}

func (Permissions) Type() protoreflect.EnumType {
	return &file_microservice_v1_service_info_proto_enumTypes[0]
}

func (x Permissions) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Permissions.Descriptor instead.
func (Permissions) EnumDescriptor() ([]byte, []int) {
	return file_microservice_v1_service_info_proto_rawDescGZIP(), []int{0}
}

type Permission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Permissions         []Permissions `protobuf:"varint,1,rep,packed,name=permissions,proto3,enum=components.microservice.v1.Permissions" json:"permissions,omitempty"`
	Optional            bool          `protobuf:"varint,2,opt,name=optional,proto3" json:"optional,omitempty"`
	ValidatePermissions bool          `protobuf:"varint,3,opt,name=validate_permissions,json=validatePermissions,proto3" json:"validate_permissions,omitempty"`
}

func (x *Permission) Reset() {
	*x = Permission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservice_v1_service_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Permission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Permission) ProtoMessage() {}

func (x *Permission) ProtoReflect() protoreflect.Message {
	mi := &file_microservice_v1_service_info_proto_msgTypes[0]
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
	return file_microservice_v1_service_info_proto_rawDescGZIP(), []int{0}
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

type GetServicePermissionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceCode int32  `protobuf:"varint,1,opt,name=service_code,json=serviceCode,proto3" json:"service_code,omitempty"`
	Version     string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *GetServicePermissionsRequest) Reset() {
	*x = GetServicePermissionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservice_v1_service_info_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServicePermissionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServicePermissionsRequest) ProtoMessage() {}

func (x *GetServicePermissionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_microservice_v1_service_info_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServicePermissionsRequest.ProtoReflect.Descriptor instead.
func (*GetServicePermissionsRequest) Descriptor() ([]byte, []int) {
	return file_microservice_v1_service_info_proto_rawDescGZIP(), []int{1}
}

func (x *GetServicePermissionsRequest) GetServiceCode() int32 {
	if x != nil {
		return x.ServiceCode
	}
	return 0
}

func (x *GetServicePermissionsRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type ServiceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string               `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Code          int32                `protobuf:"varint,4,opt,name=code,proto3" json:"code,omitempty"`
	Version       string               `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`
	Permissions   []*ServicePermission `protobuf:"bytes,6,rep,name=permissions,proto3" json:"permissions,omitempty"`
	Config        *ServiceConfig       `protobuf:"bytes,7,opt,name=config,proto3" json:"config,omitempty"`
	ApplicationId string               `protobuf:"bytes,8,opt,name=application_id,json=applicationId,proto3" json:"application_id,omitempty"`
}

func (x *ServiceInfo) Reset() {
	*x = ServiceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservice_v1_service_info_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceInfo) ProtoMessage() {}

func (x *ServiceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_microservice_v1_service_info_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceInfo.ProtoReflect.Descriptor instead.
func (*ServiceInfo) Descriptor() ([]byte, []int) {
	return file_microservice_v1_service_info_proto_rawDescGZIP(), []int{2}
}

func (x *ServiceInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ServiceInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ServiceInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ServiceInfo) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ServiceInfo) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ServiceInfo) GetPermissions() []*ServicePermission {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *ServiceInfo) GetConfig() *ServiceConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *ServiceInfo) GetApplicationId() string {
	if x != nil {
		return x.ApplicationId
	}
	return ""
}

type ServicePermission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Code        int32  `protobuf:"varint,3,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *ServicePermission) Reset() {
	*x = ServicePermission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservice_v1_service_info_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServicePermission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServicePermission) ProtoMessage() {}

func (x *ServicePermission) ProtoReflect() protoreflect.Message {
	mi := &file_microservice_v1_service_info_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServicePermission.ProtoReflect.Descriptor instead.
func (*ServicePermission) Descriptor() ([]byte, []int) {
	return file_microservice_v1_service_info_proto_rawDescGZIP(), []int{3}
}

func (x *ServicePermission) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ServicePermission) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ServicePermission) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type ServiceConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmailTemplates []*Template `protobuf:"bytes,1,rep,name=email_templates,json=emailTemplates,proto3" json:"email_templates,omitempty"`
	SmsTemplates   []*Template `protobuf:"bytes,2,rep,name=sms_templates,json=smsTemplates,proto3" json:"sms_templates,omitempty"`
}

func (x *ServiceConfig) Reset() {
	*x = ServiceConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservice_v1_service_info_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceConfig) ProtoMessage() {}

func (x *ServiceConfig) ProtoReflect() protoreflect.Message {
	mi := &file_microservice_v1_service_info_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceConfig.ProtoReflect.Descriptor instead.
func (*ServiceConfig) Descriptor() ([]byte, []int) {
	return file_microservice_v1_service_info_proto_rawDescGZIP(), []int{4}
}

func (x *ServiceConfig) GetEmailTemplates() []*Template {
	if x != nil {
		return x.EmailTemplates
	}
	return nil
}

func (x *ServiceConfig) GetSmsTemplates() []*Template {
	if x != nil {
		return x.SmsTemplates
	}
	return nil
}

type Template struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code     int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Template string `protobuf:"bytes,2,opt,name=template,proto3" json:"template,omitempty"`
}

func (x *Template) Reset() {
	*x = Template{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microservice_v1_service_info_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Template) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Template) ProtoMessage() {}

func (x *Template) ProtoReflect() protoreflect.Message {
	mi := &file_microservice_v1_service_info_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Template.ProtoReflect.Descriptor instead.
func (*Template) Descriptor() ([]byte, []int) {
	return file_microservice_v1_service_info_proto_rawDescGZIP(), []int{5}
}

func (x *Template) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Template) GetTemplate() string {
	if x != nil {
		return x.Template
	}
	return ""
}

var file_microservice_v1_service_info_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.MethodOptions)(nil),
		ExtensionType: (*Permission)(nil),
		Field:         1000000,
		Name:          "components.microservice.v1.permission",
		Tag:           "bytes,1000000,opt,name=permission",
		Filename:      "microservice/v1/service_info.proto",
	},
}

// Extension fields to descriptor.MethodOptions.
var (
	// optional components.microservice.v1.Permission permission = 1000000;
	E_Permission = &file_microservice_v1_service_info_proto_extTypes[0]
)

var File_microservice_v1_service_info_proto protoreflect.FileDescriptor

var file_microservice_v1_service_info_proto_rawDesc = []byte{
	0x0a, 0x22, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x01,
	0x0a, 0x0a, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x49, 0x0a, 0x0b,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0e, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x12, 0x31, 0x0a, 0x14, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x13, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x5b, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x22, 0xbc, 0x02, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x4f, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x41, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x25, 0x0a, 0x0e, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x22, 0x5d, 0x0a, 0x11, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x22, 0xa9, 0x01, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x4d, 0x0a, 0x0f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x52, 0x0e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x73, 0x12, 0x49, 0x0a, 0x0d, 0x73, 0x6d, 0x73, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52,
	0x0c, 0x73, 0x6d, 0x73, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x22, 0x3a, 0x0a,
	0x08, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2a, 0x17, 0x0a, 0x0b, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45,
	0x10, 0x00, 0x32, 0xf9, 0x01, 0x0a, 0x12, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5c, 0x0a, 0x0f, 0x53, 0x79, 0x6e,
	0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x27, 0x2e, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x08, 0x82,
	0xa4, 0xe8, 0x03, 0x03, 0x0a, 0x01, 0x00, 0x12, 0x84, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x38, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x63, 0x6f,
	0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x22, 0x08, 0x82, 0xa4, 0xe8, 0x03, 0x03, 0x0a, 0x01, 0x00, 0x3a, 0x68,
	0x0a, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xc0, 0x84, 0x3d,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x65, 0x5a, 0x63, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x72, 0x2d, 0x63, 0x6f, 0x64, 0x69, 0x6e,
	0x67, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74,
	0x65, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x41, 0x50, 0x49, 0x73, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x73, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x76, 0x31, 0x3b, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_microservice_v1_service_info_proto_rawDescOnce sync.Once
	file_microservice_v1_service_info_proto_rawDescData = file_microservice_v1_service_info_proto_rawDesc
)

func file_microservice_v1_service_info_proto_rawDescGZIP() []byte {
	file_microservice_v1_service_info_proto_rawDescOnce.Do(func() {
		file_microservice_v1_service_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_microservice_v1_service_info_proto_rawDescData)
	})
	return file_microservice_v1_service_info_proto_rawDescData
}

var file_microservice_v1_service_info_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_microservice_v1_service_info_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_microservice_v1_service_info_proto_goTypes = []interface{}{
	(Permissions)(0),                     // 0: components.microservice.v1.Permissions
	(*Permission)(nil),                   // 1: components.microservice.v1.Permission
	(*GetServicePermissionsRequest)(nil), // 2: components.microservice.v1.GetServicePermissionsRequest
	(*ServiceInfo)(nil),                  // 3: components.microservice.v1.ServiceInfo
	(*ServicePermission)(nil),            // 4: components.microservice.v1.ServicePermission
	(*ServiceConfig)(nil),                // 5: components.microservice.v1.ServiceConfig
	(*Template)(nil),                     // 6: components.microservice.v1.Template
	(*descriptor.MethodOptions)(nil),     // 7: google.protobuf.MethodOptions
	(*empty.Empty)(nil),                  // 8: google.protobuf.Empty
}
var file_microservice_v1_service_info_proto_depIdxs = []int32{
	0, // 0: components.microservice.v1.Permission.permissions:type_name -> components.microservice.v1.Permissions
	4, // 1: components.microservice.v1.ServiceInfo.permissions:type_name -> components.microservice.v1.ServicePermission
	5, // 2: components.microservice.v1.ServiceInfo.config:type_name -> components.microservice.v1.ServiceConfig
	6, // 3: components.microservice.v1.ServiceConfig.email_templates:type_name -> components.microservice.v1.Template
	6, // 4: components.microservice.v1.ServiceConfig.sms_templates:type_name -> components.microservice.v1.Template
	7, // 5: components.microservice.v1.permission:extendee -> google.protobuf.MethodOptions
	1, // 6: components.microservice.v1.permission:type_name -> components.microservice.v1.Permission
	3, // 7: components.microservice.v1.ServiceInfoService.SyncServiceInfo:input_type -> components.microservice.v1.ServiceInfo
	2, // 8: components.microservice.v1.ServiceInfoService.GetServicePermissions:input_type -> components.microservice.v1.GetServicePermissionsRequest
	8, // 9: components.microservice.v1.ServiceInfoService.SyncServiceInfo:output_type -> google.protobuf.Empty
	3, // 10: components.microservice.v1.ServiceInfoService.GetServicePermissions:output_type -> components.microservice.v1.ServiceInfo
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	6, // [6:7] is the sub-list for extension type_name
	5, // [5:6] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_microservice_v1_service_info_proto_init() }
func file_microservice_v1_service_info_proto_init() {
	if File_microservice_v1_service_info_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_microservice_v1_service_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_microservice_v1_service_info_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetServicePermissionsRequest); i {
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
		file_microservice_v1_service_info_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceInfo); i {
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
		file_microservice_v1_service_info_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServicePermission); i {
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
		file_microservice_v1_service_info_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceConfig); i {
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
		file_microservice_v1_service_info_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Template); i {
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
			RawDescriptor: file_microservice_v1_service_info_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 1,
			NumServices:   1,
		},
		GoTypes:           file_microservice_v1_service_info_proto_goTypes,
		DependencyIndexes: file_microservice_v1_service_info_proto_depIdxs,
		EnumInfos:         file_microservice_v1_service_info_proto_enumTypes,
		MessageInfos:      file_microservice_v1_service_info_proto_msgTypes,
		ExtensionInfos:    file_microservice_v1_service_info_proto_extTypes,
	}.Build()
	File_microservice_v1_service_info_proto = out.File
	file_microservice_v1_service_info_proto_rawDesc = nil
	file_microservice_v1_service_info_proto_goTypes = nil
	file_microservice_v1_service_info_proto_depIdxs = nil
}
