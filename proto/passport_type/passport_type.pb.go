// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: passport_type/passport_type.proto

package proto_passport_type

import (
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

type ListPassportTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageNumber int32 `protobuf:"varint,1,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	PageSize   int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *ListPassportTypeRequest) Reset() {
	*x = ListPassportTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_passport_type_passport_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPassportTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPassportTypeRequest) ProtoMessage() {}

func (x *ListPassportTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_passport_type_passport_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPassportTypeRequest.ProtoReflect.Descriptor instead.
func (*ListPassportTypeRequest) Descriptor() ([]byte, []int) {
	return file_passport_type_passport_type_proto_rawDescGZIP(), []int{0}
}

func (x *ListPassportTypeRequest) GetPageNumber() int32 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *ListPassportTypeRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListPassportTypeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PassportTypes []*PassportTypeProfileResponse `protobuf:"bytes,1,rep,name=passport_types,json=passportTypes,proto3" json:"passport_types,omitempty"`
}

func (x *ListPassportTypeResponse) Reset() {
	*x = ListPassportTypeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_passport_type_passport_type_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPassportTypeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPassportTypeResponse) ProtoMessage() {}

func (x *ListPassportTypeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_passport_type_passport_type_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPassportTypeResponse.ProtoReflect.Descriptor instead.
func (*ListPassportTypeResponse) Descriptor() ([]byte, []int) {
	return file_passport_type_passport_type_proto_rawDescGZIP(), []int{1}
}

func (x *ListPassportTypeResponse) GetPassportTypes() []*PassportTypeProfileResponse {
	if x != nil {
		return x.PassportTypes
	}
	return nil
}

type CreatePassportTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code             string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Name             string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	OldCode          string `protobuf:"bytes,3,opt,name=old_code,json=oldCode,proto3" json:"old_code,omitempty"`
	OldName          string `protobuf:"bytes,4,opt,name=old_name,json=oldName,proto3" json:"old_name,omitempty"`
	ActivationDate   string `protobuf:"bytes,5,opt,name=activation_date,json=activationDate,proto3" json:"activation_date,omitempty"`
	DeactivationDate string `protobuf:"bytes,6,opt,name=deactivation_date,json=deactivationDate,proto3" json:"deactivation_date,omitempty"`
	FlexFinId        string `protobuf:"bytes,7,opt,name=flex_fin_id,json=flexFinId,proto3" json:"flex_fin_id,omitempty"`
}

func (x *CreatePassportTypeRequest) Reset() {
	*x = CreatePassportTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_passport_type_passport_type_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePassportTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePassportTypeRequest) ProtoMessage() {}

func (x *CreatePassportTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_passport_type_passport_type_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePassportTypeRequest.ProtoReflect.Descriptor instead.
func (*CreatePassportTypeRequest) Descriptor() ([]byte, []int) {
	return file_passport_type_passport_type_proto_rawDescGZIP(), []int{2}
}

func (x *CreatePassportTypeRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CreatePassportTypeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreatePassportTypeRequest) GetOldCode() string {
	if x != nil {
		return x.OldCode
	}
	return ""
}

func (x *CreatePassportTypeRequest) GetOldName() string {
	if x != nil {
		return x.OldName
	}
	return ""
}

func (x *CreatePassportTypeRequest) GetActivationDate() string {
	if x != nil {
		return x.ActivationDate
	}
	return ""
}

func (x *CreatePassportTypeRequest) GetDeactivationDate() string {
	if x != nil {
		return x.DeactivationDate
	}
	return ""
}

func (x *CreatePassportTypeRequest) GetFlexFinId() string {
	if x != nil {
		return x.FlexFinId
	}
	return ""
}

type SinglePassportTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to PassportTypeIdentifier:
	//
	//	*SinglePassportTypeRequest_Id
	//	*SinglePassportTypeRequest_Code
	PassportTypeIdentifier isSinglePassportTypeRequest_PassportTypeIdentifier `protobuf_oneof:"passport_type_identifier"`
}

func (x *SinglePassportTypeRequest) Reset() {
	*x = SinglePassportTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_passport_type_passport_type_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SinglePassportTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SinglePassportTypeRequest) ProtoMessage() {}

func (x *SinglePassportTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_passport_type_passport_type_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SinglePassportTypeRequest.ProtoReflect.Descriptor instead.
func (*SinglePassportTypeRequest) Descriptor() ([]byte, []int) {
	return file_passport_type_passport_type_proto_rawDescGZIP(), []int{3}
}

func (m *SinglePassportTypeRequest) GetPassportTypeIdentifier() isSinglePassportTypeRequest_PassportTypeIdentifier {
	if m != nil {
		return m.PassportTypeIdentifier
	}
	return nil
}

func (x *SinglePassportTypeRequest) GetId() string {
	if x, ok := x.GetPassportTypeIdentifier().(*SinglePassportTypeRequest_Id); ok {
		return x.Id
	}
	return ""
}

func (x *SinglePassportTypeRequest) GetCode() string {
	if x, ok := x.GetPassportTypeIdentifier().(*SinglePassportTypeRequest_Code); ok {
		return x.Code
	}
	return ""
}

type isSinglePassportTypeRequest_PassportTypeIdentifier interface {
	isSinglePassportTypeRequest_PassportTypeIdentifier()
}

type SinglePassportTypeRequest_Id struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3,oneof"`
}

type SinglePassportTypeRequest_Code struct {
	Code string `protobuf:"bytes,2,opt,name=code,proto3,oneof"`
}

func (*SinglePassportTypeRequest_Id) isSinglePassportTypeRequest_PassportTypeIdentifier() {}

func (*SinglePassportTypeRequest_Code) isSinglePassportTypeRequest_PassportTypeIdentifier() {}

type PassportTypeProfileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Code             string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Name             string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	OldCode          string `protobuf:"bytes,4,opt,name=old_code,json=oldCode,proto3" json:"old_code,omitempty"`
	OldName          string `protobuf:"bytes,5,opt,name=old_name,json=oldName,proto3" json:"old_name,omitempty"`
	ActivationDate   string `protobuf:"bytes,6,opt,name=activation_date,json=activationDate,proto3" json:"activation_date,omitempty"`
	DeactivationDate string `protobuf:"bytes,7,opt,name=deactivation_date,json=deactivationDate,proto3" json:"deactivation_date,omitempty"`
	FlexFinId        string `protobuf:"bytes,8,opt,name=flex_fin_id,json=flexFinId,proto3" json:"flex_fin_id,omitempty"`
	DeletedAt        string `protobuf:"bytes,9,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *PassportTypeProfileResponse) Reset() {
	*x = PassportTypeProfileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_passport_type_passport_type_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PassportTypeProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PassportTypeProfileResponse) ProtoMessage() {}

func (x *PassportTypeProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_passport_type_passport_type_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PassportTypeProfileResponse.ProtoReflect.Descriptor instead.
func (*PassportTypeProfileResponse) Descriptor() ([]byte, []int) {
	return file_passport_type_passport_type_proto_rawDescGZIP(), []int{4}
}

func (x *PassportTypeProfileResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PassportTypeProfileResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *PassportTypeProfileResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PassportTypeProfileResponse) GetOldCode() string {
	if x != nil {
		return x.OldCode
	}
	return ""
}

func (x *PassportTypeProfileResponse) GetOldName() string {
	if x != nil {
		return x.OldName
	}
	return ""
}

func (x *PassportTypeProfileResponse) GetActivationDate() string {
	if x != nil {
		return x.ActivationDate
	}
	return ""
}

func (x *PassportTypeProfileResponse) GetDeactivationDate() string {
	if x != nil {
		return x.DeactivationDate
	}
	return ""
}

func (x *PassportTypeProfileResponse) GetFlexFinId() string {
	if x != nil {
		return x.FlexFinId
	}
	return ""
}

func (x *PassportTypeProfileResponse) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

type PassportTypeSuccessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *PassportTypeSuccessResponse) Reset() {
	*x = PassportTypeSuccessResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_passport_type_passport_type_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PassportTypeSuccessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PassportTypeSuccessResponse) ProtoMessage() {}

func (x *PassportTypeSuccessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_passport_type_passport_type_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PassportTypeSuccessResponse.ProtoReflect.Descriptor instead.
func (*PassportTypeSuccessResponse) Descriptor() ([]byte, []int) {
	return file_passport_type_passport_type_proto_rawDescGZIP(), []int{5}
}

func (x *PassportTypeSuccessResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type UpdatePassportTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Code             string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Name             string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	OldCode          string `protobuf:"bytes,4,opt,name=old_code,json=oldCode,proto3" json:"old_code,omitempty"`
	OldName          string `protobuf:"bytes,5,opt,name=old_name,json=oldName,proto3" json:"old_name,omitempty"`
	ActivationDate   string `protobuf:"bytes,6,opt,name=activation_date,json=activationDate,proto3" json:"activation_date,omitempty"`
	DeactivationDate string `protobuf:"bytes,7,opt,name=deactivation_date,json=deactivationDate,proto3" json:"deactivation_date,omitempty"`
	FlexFinId        string `protobuf:"bytes,8,opt,name=flex_fin_id,json=flexFinId,proto3" json:"flex_fin_id,omitempty"`
}

func (x *UpdatePassportTypeRequest) Reset() {
	*x = UpdatePassportTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_passport_type_passport_type_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePassportTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePassportTypeRequest) ProtoMessage() {}

func (x *UpdatePassportTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_passport_type_passport_type_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePassportTypeRequest.ProtoReflect.Descriptor instead.
func (*UpdatePassportTypeRequest) Descriptor() ([]byte, []int) {
	return file_passport_type_passport_type_proto_rawDescGZIP(), []int{6}
}

func (x *UpdatePassportTypeRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdatePassportTypeRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *UpdatePassportTypeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdatePassportTypeRequest) GetOldCode() string {
	if x != nil {
		return x.OldCode
	}
	return ""
}

func (x *UpdatePassportTypeRequest) GetOldName() string {
	if x != nil {
		return x.OldName
	}
	return ""
}

func (x *UpdatePassportTypeRequest) GetActivationDate() string {
	if x != nil {
		return x.ActivationDate
	}
	return ""
}

func (x *UpdatePassportTypeRequest) GetDeactivationDate() string {
	if x != nil {
		return x.DeactivationDate
	}
	return ""
}

func (x *UpdatePassportTypeRequest) GetFlexFinId() string {
	if x != nil {
		return x.FlexFinId
	}
	return ""
}

var File_passport_type_passport_type_proto protoreflect.FileDescriptor

var file_passport_type_passport_type_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2f,
	0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x57, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x73, 0x73, 0x70,
	0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x5f, 0x0a, 0x18,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0e, 0x70, 0x61, 0x73, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x50, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0d,
	0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x22, 0xef, 0x01,
	0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x6c, 0x64, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x6c, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x6f, 0x6c, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6f, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61,
	0x74, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x64, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x64,
	0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12,
	0x1e, 0x0a, 0x0b, 0x66, 0x6c, 0x65, 0x78, 0x5f, 0x66, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x6c, 0x65, 0x78, 0x46, 0x69, 0x6e, 0x49, 0x64, 0x22,
	0x5f, 0x0a, 0x19, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x50, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x42, 0x1a, 0x0a, 0x18, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x22, 0xa0, 0x02, 0x0a, 0x1b, 0x50, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x6c, 0x64, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x6c, 0x64, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x6c, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x27,
	0x0a, 0x0f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x64, 0x65, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x10, 0x64, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x0a, 0x0b, 0x66, 0x6c, 0x65, 0x78, 0x5f, 0x66, 0x69, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x6c, 0x65, 0x78, 0x46,
	0x69, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0x39, 0x0a, 0x1b, 0x50, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xff,
	0x01, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x6c, 0x64, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x6c, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x6f, 0x6c, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6f, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x64, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10,
	0x64, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65,
	0x12, 0x1e, 0x0a, 0x0b, 0x66, 0x6c, 0x65, 0x78, 0x5f, 0x66, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x6c, 0x65, 0x78, 0x46, 0x69, 0x6e, 0x49, 0x64,
	0x32, 0xdf, 0x02, 0x0a, 0x13, 0x50, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x18, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x1a, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x50, 0x61,
	0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x03, 0x47, 0x65, 0x74,
	0x12, 0x1a, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x50, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x50,
	0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x06, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x73,
	0x73, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x50, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x53,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42,
	0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c,
	0x65, 0x50, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x50, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x15, 0x5a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x61, 0x73, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_passport_type_passport_type_proto_rawDescOnce sync.Once
	file_passport_type_passport_type_proto_rawDescData = file_passport_type_passport_type_proto_rawDesc
)

func file_passport_type_passport_type_proto_rawDescGZIP() []byte {
	file_passport_type_passport_type_proto_rawDescOnce.Do(func() {
		file_passport_type_passport_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_passport_type_passport_type_proto_rawDescData)
	})
	return file_passport_type_passport_type_proto_rawDescData
}

var file_passport_type_passport_type_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_passport_type_passport_type_proto_goTypes = []interface{}{
	(*ListPassportTypeRequest)(nil),     // 0: ListPassportTypeRequest
	(*ListPassportTypeResponse)(nil),    // 1: ListPassportTypeResponse
	(*CreatePassportTypeRequest)(nil),   // 2: CreatePassportTypeRequest
	(*SinglePassportTypeRequest)(nil),   // 3: SinglePassportTypeRequest
	(*PassportTypeProfileResponse)(nil), // 4: PassportTypeProfileResponse
	(*PassportTypeSuccessResponse)(nil), // 5: PassportTypeSuccessResponse
	(*UpdatePassportTypeRequest)(nil),   // 6: UpdatePassportTypeRequest
}
var file_passport_type_passport_type_proto_depIdxs = []int32{
	4, // 0: ListPassportTypeResponse.passport_types:type_name -> PassportTypeProfileResponse
	0, // 1: PassportTypeService.List:input_type -> ListPassportTypeRequest
	2, // 2: PassportTypeService.Create:input_type -> CreatePassportTypeRequest
	3, // 3: PassportTypeService.Get:input_type -> SinglePassportTypeRequest
	6, // 4: PassportTypeService.Update:input_type -> UpdatePassportTypeRequest
	3, // 5: PassportTypeService.Delete:input_type -> SinglePassportTypeRequest
	1, // 6: PassportTypeService.List:output_type -> ListPassportTypeResponse
	4, // 7: PassportTypeService.Create:output_type -> PassportTypeProfileResponse
	4, // 8: PassportTypeService.Get:output_type -> PassportTypeProfileResponse
	5, // 9: PassportTypeService.Update:output_type -> PassportTypeSuccessResponse
	5, // 10: PassportTypeService.Delete:output_type -> PassportTypeSuccessResponse
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_passport_type_passport_type_proto_init() }
func file_passport_type_passport_type_proto_init() {
	if File_passport_type_passport_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_passport_type_passport_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPassportTypeRequest); i {
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
		file_passport_type_passport_type_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPassportTypeResponse); i {
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
		file_passport_type_passport_type_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePassportTypeRequest); i {
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
		file_passport_type_passport_type_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SinglePassportTypeRequest); i {
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
		file_passport_type_passport_type_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PassportTypeProfileResponse); i {
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
		file_passport_type_passport_type_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PassportTypeSuccessResponse); i {
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
		file_passport_type_passport_type_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePassportTypeRequest); i {
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
	file_passport_type_passport_type_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*SinglePassportTypeRequest_Id)(nil),
		(*SinglePassportTypeRequest_Code)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_passport_type_passport_type_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_passport_type_passport_type_proto_goTypes,
		DependencyIndexes: file_passport_type_passport_type_proto_depIdxs,
		MessageInfos:      file_passport_type_passport_type_proto_msgTypes,
	}.Build()
	File_passport_type_passport_type_proto = out.File
	file_passport_type_passport_type_proto_rawDesc = nil
	file_passport_type_passport_type_proto_goTypes = nil
	file_passport_type_passport_type_proto_depIdxs = nil
}
