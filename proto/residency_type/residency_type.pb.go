// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.12
// source: residency_type.proto

package proto_residency_type

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

type ListResidencyTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageNumber int32 `protobuf:"varint,1,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	PageSize   int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *ListResidencyTypeRequest) Reset() {
	*x = ListResidencyTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_residency_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResidencyTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResidencyTypeRequest) ProtoMessage() {}

func (x *ListResidencyTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_residency_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResidencyTypeRequest.ProtoReflect.Descriptor instead.
func (*ListResidencyTypeRequest) Descriptor() ([]byte, []int) {
	return file_residency_type_proto_rawDescGZIP(), []int{0}
}

func (x *ListResidencyTypeRequest) GetPageNumber() int32 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *ListResidencyTypeRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListResidencyTypeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResidencyTypes []*ResidencyTypeProfileResponse `protobuf:"bytes,1,rep,name=residency_types,json=residencyTypes,proto3" json:"residency_types,omitempty"`
}

func (x *ListResidencyTypeResponse) Reset() {
	*x = ListResidencyTypeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_residency_type_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResidencyTypeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResidencyTypeResponse) ProtoMessage() {}

func (x *ListResidencyTypeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_residency_type_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResidencyTypeResponse.ProtoReflect.Descriptor instead.
func (*ListResidencyTypeResponse) Descriptor() ([]byte, []int) {
	return file_residency_type_proto_rawDescGZIP(), []int{1}
}

func (x *ListResidencyTypeResponse) GetResidencyTypes() []*ResidencyTypeProfileResponse {
	if x != nil {
		return x.ResidencyTypes
	}
	return nil
}

type CreateResidencyTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code             string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Name             string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ActivationDate   string `protobuf:"bytes,3,opt,name=activation_date,json=activationDate,proto3" json:"activation_date,omitempty"`
	DeactivationDate string `protobuf:"bytes,4,opt,name=deactivation_date,json=deactivationDate,proto3" json:"deactivation_date,omitempty"`
	FlexFinId        string `protobuf:"bytes,5,opt,name=flex_fin_id,json=flexFinId,proto3" json:"flex_fin_id,omitempty"`
	NameUz           string `protobuf:"bytes,6,opt,name=name_uz,json=nameUz,proto3" json:"name_uz,omitempty"`
}

func (x *CreateResidencyTypeRequest) Reset() {
	*x = CreateResidencyTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_residency_type_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResidencyTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResidencyTypeRequest) ProtoMessage() {}

func (x *CreateResidencyTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_residency_type_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResidencyTypeRequest.ProtoReflect.Descriptor instead.
func (*CreateResidencyTypeRequest) Descriptor() ([]byte, []int) {
	return file_residency_type_proto_rawDescGZIP(), []int{2}
}

func (x *CreateResidencyTypeRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CreateResidencyTypeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateResidencyTypeRequest) GetActivationDate() string {
	if x != nil {
		return x.ActivationDate
	}
	return ""
}

func (x *CreateResidencyTypeRequest) GetDeactivationDate() string {
	if x != nil {
		return x.DeactivationDate
	}
	return ""
}

func (x *CreateResidencyTypeRequest) GetFlexFinId() string {
	if x != nil {
		return x.FlexFinId
	}
	return ""
}

func (x *CreateResidencyTypeRequest) GetNameUz() string {
	if x != nil {
		return x.NameUz
	}
	return ""
}

type SingleResidencyTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to ResidencyTypeIdentifier:
	//
	//	*SingleResidencyTypeRequest_Id
	//	*SingleResidencyTypeRequest_Code
	ResidencyTypeIdentifier isSingleResidencyTypeRequest_ResidencyTypeIdentifier `protobuf_oneof:"residency_type_identifier"`
}

func (x *SingleResidencyTypeRequest) Reset() {
	*x = SingleResidencyTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_residency_type_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SingleResidencyTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SingleResidencyTypeRequest) ProtoMessage() {}

func (x *SingleResidencyTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_residency_type_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SingleResidencyTypeRequest.ProtoReflect.Descriptor instead.
func (*SingleResidencyTypeRequest) Descriptor() ([]byte, []int) {
	return file_residency_type_proto_rawDescGZIP(), []int{3}
}

func (m *SingleResidencyTypeRequest) GetResidencyTypeIdentifier() isSingleResidencyTypeRequest_ResidencyTypeIdentifier {
	if m != nil {
		return m.ResidencyTypeIdentifier
	}
	return nil
}

func (x *SingleResidencyTypeRequest) GetId() string {
	if x, ok := x.GetResidencyTypeIdentifier().(*SingleResidencyTypeRequest_Id); ok {
		return x.Id
	}
	return ""
}

func (x *SingleResidencyTypeRequest) GetCode() string {
	if x, ok := x.GetResidencyTypeIdentifier().(*SingleResidencyTypeRequest_Code); ok {
		return x.Code
	}
	return ""
}

type isSingleResidencyTypeRequest_ResidencyTypeIdentifier interface {
	isSingleResidencyTypeRequest_ResidencyTypeIdentifier()
}

type SingleResidencyTypeRequest_Id struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3,oneof"`
}

type SingleResidencyTypeRequest_Code struct {
	Code string `protobuf:"bytes,2,opt,name=code,proto3,oneof"`
}

func (*SingleResidencyTypeRequest_Id) isSingleResidencyTypeRequest_ResidencyTypeIdentifier() {}

func (*SingleResidencyTypeRequest_Code) isSingleResidencyTypeRequest_ResidencyTypeIdentifier() {}

type ResidencyTypeProfileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Code             string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Name             string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	ActivationDate   string `protobuf:"bytes,4,opt,name=activation_date,json=activationDate,proto3" json:"activation_date,omitempty"`
	DeactivationDate string `protobuf:"bytes,5,opt,name=deactivation_date,json=deactivationDate,proto3" json:"deactivation_date,omitempty"`
	FlexFinId        string `protobuf:"bytes,6,opt,name=flex_fin_id,json=flexFinId,proto3" json:"flex_fin_id,omitempty"`
	NameUz           string `protobuf:"bytes,7,opt,name=name_uz,json=nameUz,proto3" json:"name_uz,omitempty"`
	DeletedAt        string `protobuf:"bytes,8,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *ResidencyTypeProfileResponse) Reset() {
	*x = ResidencyTypeProfileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_residency_type_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResidencyTypeProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResidencyTypeProfileResponse) ProtoMessage() {}

func (x *ResidencyTypeProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_residency_type_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResidencyTypeProfileResponse.ProtoReflect.Descriptor instead.
func (*ResidencyTypeProfileResponse) Descriptor() ([]byte, []int) {
	return file_residency_type_proto_rawDescGZIP(), []int{4}
}

func (x *ResidencyTypeProfileResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ResidencyTypeProfileResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *ResidencyTypeProfileResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ResidencyTypeProfileResponse) GetActivationDate() string {
	if x != nil {
		return x.ActivationDate
	}
	return ""
}

func (x *ResidencyTypeProfileResponse) GetDeactivationDate() string {
	if x != nil {
		return x.DeactivationDate
	}
	return ""
}

func (x *ResidencyTypeProfileResponse) GetFlexFinId() string {
	if x != nil {
		return x.FlexFinId
	}
	return ""
}

func (x *ResidencyTypeProfileResponse) GetNameUz() string {
	if x != nil {
		return x.NameUz
	}
	return ""
}

func (x *ResidencyTypeProfileResponse) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

type ResidencyTypeSuccessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *ResidencyTypeSuccessResponse) Reset() {
	*x = ResidencyTypeSuccessResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_residency_type_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResidencyTypeSuccessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResidencyTypeSuccessResponse) ProtoMessage() {}

func (x *ResidencyTypeSuccessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_residency_type_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResidencyTypeSuccessResponse.ProtoReflect.Descriptor instead.
func (*ResidencyTypeSuccessResponse) Descriptor() ([]byte, []int) {
	return file_residency_type_proto_rawDescGZIP(), []int{5}
}

func (x *ResidencyTypeSuccessResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type UpdateResidencyTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Code             string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Name             string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	ActivationDate   string `protobuf:"bytes,4,opt,name=activation_date,json=activationDate,proto3" json:"activation_date,omitempty"`
	DeactivationDate string `protobuf:"bytes,5,opt,name=deactivation_date,json=deactivationDate,proto3" json:"deactivation_date,omitempty"`
	FlexFinId        string `protobuf:"bytes,6,opt,name=flex_fin_id,json=flexFinId,proto3" json:"flex_fin_id,omitempty"`
	NameUz           string `protobuf:"bytes,7,opt,name=name_uz,json=nameUz,proto3" json:"name_uz,omitempty"`
}

func (x *UpdateResidencyTypeRequest) Reset() {
	*x = UpdateResidencyTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_residency_type_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResidencyTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResidencyTypeRequest) ProtoMessage() {}

func (x *UpdateResidencyTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_residency_type_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResidencyTypeRequest.ProtoReflect.Descriptor instead.
func (*UpdateResidencyTypeRequest) Descriptor() ([]byte, []int) {
	return file_residency_type_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateResidencyTypeRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateResidencyTypeRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *UpdateResidencyTypeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateResidencyTypeRequest) GetActivationDate() string {
	if x != nil {
		return x.ActivationDate
	}
	return ""
}

func (x *UpdateResidencyTypeRequest) GetDeactivationDate() string {
	if x != nil {
		return x.DeactivationDate
	}
	return ""
}

func (x *UpdateResidencyTypeRequest) GetFlexFinId() string {
	if x != nil {
		return x.FlexFinId
	}
	return ""
}

func (x *UpdateResidencyTypeRequest) GetNameUz() string {
	if x != nil {
		return x.NameUz
	}
	return ""
}

var File_residency_type_proto protoreflect.FileDescriptor

var file_residency_type_proto_rawDesc = []byte{
	0x0a, 0x14, 0x72, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x58, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x22, 0x63, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63,
	0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a,
	0x0f, 0x72, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x52, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e,
	0x63, 0x79, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0e, 0x72, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x79,
	0x54, 0x79, 0x70, 0x65, 0x73, 0x22, 0xd3, 0x01, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x64, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x64, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61,
	0x74, 0x65, 0x12, 0x1e, 0x0a, 0x0b, 0x66, 0x6c, 0x65, 0x78, 0x5f, 0x66, 0x69, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x6c, 0x65, 0x78, 0x46, 0x69, 0x6e,
	0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x75, 0x7a, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x61, 0x6d, 0x65, 0x55, 0x7a, 0x22, 0x61, 0x0a, 0x1a, 0x53,
	0x69, 0x6e, 0x67, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x42, 0x1b, 0x0a, 0x19, 0x72, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x22, 0x84,
	0x02, 0x0a, 0x1c, 0x52, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x54, 0x79, 0x70, 0x65,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65,
	0x12, 0x2b, 0x0a, 0x11, 0x64, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x64, 0x65, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x0a,
	0x0b, 0x66, 0x6c, 0x65, 0x78, 0x5f, 0x66, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x66, 0x6c, 0x65, 0x78, 0x46, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x75, 0x7a, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6e, 0x61, 0x6d, 0x65, 0x55, 0x7a, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x3a, 0x0a, 0x1c, 0x52, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e,
	0x63, 0x79, 0x54, 0x79, 0x70, 0x65, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0xe3, 0x01, 0x0a, 0x1a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x69,
	0x64, 0x65, 0x6e, 0x63, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x2b, 0x0a, 0x11, 0x64, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x64, 0x65,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1e,
	0x0a, 0x0b, 0x66, 0x6c, 0x65, 0x78, 0x5f, 0x66, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x6c, 0x65, 0x78, 0x46, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x75, 0x7a, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6e, 0x61, 0x6d, 0x65, 0x55, 0x7a, 0x32, 0xea, 0x02, 0x0a, 0x14, 0x52, 0x65, 0x73, 0x69,
	0x64, 0x65, 0x6e, 0x63, 0x79, 0x54, 0x79, 0x70, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x3d, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x19, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x69, 0x64, 0x65,
	0x6e, 0x63, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x44, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x52, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e,
	0x63, 0x79, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1b, 0x2e, 0x53,
	0x69, 0x6e, 0x67, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x52, 0x65, 0x73, 0x69,
	0x64, 0x65, 0x6e, 0x63, 0x79, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x1b, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x69, 0x64,
	0x65, 0x6e, 0x63, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1d, 0x2e, 0x52, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x54, 0x79, 0x70, 0x65, 0x53,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44,
	0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x52, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x63,
	0x79, 0x54, 0x79, 0x70, 0x65, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x16, 0x5a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65,
	0x73, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_residency_type_proto_rawDescOnce sync.Once
	file_residency_type_proto_rawDescData = file_residency_type_proto_rawDesc
)

func file_residency_type_proto_rawDescGZIP() []byte {
	file_residency_type_proto_rawDescOnce.Do(func() {
		file_residency_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_residency_type_proto_rawDescData)
	})
	return file_residency_type_proto_rawDescData
}

var file_residency_type_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_residency_type_proto_goTypes = []interface{}{
	(*ListResidencyTypeRequest)(nil),     // 0: ListResidencyTypeRequest
	(*ListResidencyTypeResponse)(nil),    // 1: ListResidencyTypeResponse
	(*CreateResidencyTypeRequest)(nil),   // 2: CreateResidencyTypeRequest
	(*SingleResidencyTypeRequest)(nil),   // 3: SingleResidencyTypeRequest
	(*ResidencyTypeProfileResponse)(nil), // 4: ResidencyTypeProfileResponse
	(*ResidencyTypeSuccessResponse)(nil), // 5: ResidencyTypeSuccessResponse
	(*UpdateResidencyTypeRequest)(nil),   // 6: UpdateResidencyTypeRequest
}
var file_residency_type_proto_depIdxs = []int32{
	4, // 0: ListResidencyTypeResponse.residency_types:type_name -> ResidencyTypeProfileResponse
	0, // 1: ResidencyTypeService.List:input_type -> ListResidencyTypeRequest
	2, // 2: ResidencyTypeService.Create:input_type -> CreateResidencyTypeRequest
	3, // 3: ResidencyTypeService.Get:input_type -> SingleResidencyTypeRequest
	6, // 4: ResidencyTypeService.Update:input_type -> UpdateResidencyTypeRequest
	3, // 5: ResidencyTypeService.Delete:input_type -> SingleResidencyTypeRequest
	1, // 6: ResidencyTypeService.List:output_type -> ListResidencyTypeResponse
	4, // 7: ResidencyTypeService.Create:output_type -> ResidencyTypeProfileResponse
	4, // 8: ResidencyTypeService.Get:output_type -> ResidencyTypeProfileResponse
	5, // 9: ResidencyTypeService.Update:output_type -> ResidencyTypeSuccessResponse
	5, // 10: ResidencyTypeService.Delete:output_type -> ResidencyTypeSuccessResponse
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_residency_type_proto_init() }
func file_residency_type_proto_init() {
	if File_residency_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_residency_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResidencyTypeRequest); i {
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
		file_residency_type_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResidencyTypeResponse); i {
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
		file_residency_type_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResidencyTypeRequest); i {
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
		file_residency_type_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SingleResidencyTypeRequest); i {
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
		file_residency_type_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResidencyTypeProfileResponse); i {
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
		file_residency_type_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResidencyTypeSuccessResponse); i {
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
		file_residency_type_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateResidencyTypeRequest); i {
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
	file_residency_type_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*SingleResidencyTypeRequest_Id)(nil),
		(*SingleResidencyTypeRequest_Code)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_residency_type_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_residency_type_proto_goTypes,
		DependencyIndexes: file_residency_type_proto_depIdxs,
		MessageInfos:      file_residency_type_proto_msgTypes,
	}.Build()
	File_residency_type_proto = out.File
	file_residency_type_proto_rawDesc = nil
	file_residency_type_proto_goTypes = nil
	file_residency_type_proto_depIdxs = nil
}
