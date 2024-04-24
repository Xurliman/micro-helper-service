// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.12
// source: borrower_type/borrower_type.proto

package proto_borrower_type

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

type ListBorrowerTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageNumber int32 `protobuf:"varint,1,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	PageSize   int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *ListBorrowerTypeRequest) Reset() {
	*x = ListBorrowerTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_borrower_type_borrower_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBorrowerTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBorrowerTypeRequest) ProtoMessage() {}

func (x *ListBorrowerTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_borrower_type_borrower_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBorrowerTypeRequest.ProtoReflect.Descriptor instead.
func (*ListBorrowerTypeRequest) Descriptor() ([]byte, []int) {
	return file_borrower_type_borrower_type_proto_rawDescGZIP(), []int{0}
}

func (x *ListBorrowerTypeRequest) GetPageNumber() int32 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *ListBorrowerTypeRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListBorrowerTypeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BorrowerTypes []*BorrowerTypeProfileResponse `protobuf:"bytes,1,rep,name=borrower_types,json=borrowerTypes,proto3" json:"borrower_types,omitempty"`
}

func (x *ListBorrowerTypeResponse) Reset() {
	*x = ListBorrowerTypeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_borrower_type_borrower_type_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBorrowerTypeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBorrowerTypeResponse) ProtoMessage() {}

func (x *ListBorrowerTypeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_borrower_type_borrower_type_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBorrowerTypeResponse.ProtoReflect.Descriptor instead.
func (*ListBorrowerTypeResponse) Descriptor() ([]byte, []int) {
	return file_borrower_type_borrower_type_proto_rawDescGZIP(), []int{1}
}

func (x *ListBorrowerTypeResponse) GetBorrowerTypes() []*BorrowerTypeProfileResponse {
	if x != nil {
		return x.BorrowerTypes
	}
	return nil
}

type CreateBorrowerTypeRequest struct {
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
	NameUz           string `protobuf:"bytes,8,opt,name=name_uz,json=nameUz,proto3" json:"name_uz,omitempty"`
}

func (x *CreateBorrowerTypeRequest) Reset() {
	*x = CreateBorrowerTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_borrower_type_borrower_type_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBorrowerTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBorrowerTypeRequest) ProtoMessage() {}

func (x *CreateBorrowerTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_borrower_type_borrower_type_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBorrowerTypeRequest.ProtoReflect.Descriptor instead.
func (*CreateBorrowerTypeRequest) Descriptor() ([]byte, []int) {
	return file_borrower_type_borrower_type_proto_rawDescGZIP(), []int{2}
}

func (x *CreateBorrowerTypeRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CreateBorrowerTypeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateBorrowerTypeRequest) GetOldCode() string {
	if x != nil {
		return x.OldCode
	}
	return ""
}

func (x *CreateBorrowerTypeRequest) GetOldName() string {
	if x != nil {
		return x.OldName
	}
	return ""
}

func (x *CreateBorrowerTypeRequest) GetActivationDate() string {
	if x != nil {
		return x.ActivationDate
	}
	return ""
}

func (x *CreateBorrowerTypeRequest) GetDeactivationDate() string {
	if x != nil {
		return x.DeactivationDate
	}
	return ""
}

func (x *CreateBorrowerTypeRequest) GetFlexFinId() string {
	if x != nil {
		return x.FlexFinId
	}
	return ""
}

func (x *CreateBorrowerTypeRequest) GetNameUz() string {
	if x != nil {
		return x.NameUz
	}
	return ""
}

type SingleBorrowerTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to BorrowerTypeIdentifier:
	//
	//	*SingleBorrowerTypeRequest_Id
	//	*SingleBorrowerTypeRequest_Code
	BorrowerTypeIdentifier isSingleBorrowerTypeRequest_BorrowerTypeIdentifier `protobuf_oneof:"borrower_type_identifier"`
}

func (x *SingleBorrowerTypeRequest) Reset() {
	*x = SingleBorrowerTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_borrower_type_borrower_type_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SingleBorrowerTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SingleBorrowerTypeRequest) ProtoMessage() {}

func (x *SingleBorrowerTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_borrower_type_borrower_type_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SingleBorrowerTypeRequest.ProtoReflect.Descriptor instead.
func (*SingleBorrowerTypeRequest) Descriptor() ([]byte, []int) {
	return file_borrower_type_borrower_type_proto_rawDescGZIP(), []int{3}
}

func (m *SingleBorrowerTypeRequest) GetBorrowerTypeIdentifier() isSingleBorrowerTypeRequest_BorrowerTypeIdentifier {
	if m != nil {
		return m.BorrowerTypeIdentifier
	}
	return nil
}

func (x *SingleBorrowerTypeRequest) GetId() string {
	if x, ok := x.GetBorrowerTypeIdentifier().(*SingleBorrowerTypeRequest_Id); ok {
		return x.Id
	}
	return ""
}

func (x *SingleBorrowerTypeRequest) GetCode() string {
	if x, ok := x.GetBorrowerTypeIdentifier().(*SingleBorrowerTypeRequest_Code); ok {
		return x.Code
	}
	return ""
}

type isSingleBorrowerTypeRequest_BorrowerTypeIdentifier interface {
	isSingleBorrowerTypeRequest_BorrowerTypeIdentifier()
}

type SingleBorrowerTypeRequest_Id struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3,oneof"`
}

type SingleBorrowerTypeRequest_Code struct {
	Code string `protobuf:"bytes,2,opt,name=code,proto3,oneof"`
}

func (*SingleBorrowerTypeRequest_Id) isSingleBorrowerTypeRequest_BorrowerTypeIdentifier() {}

func (*SingleBorrowerTypeRequest_Code) isSingleBorrowerTypeRequest_BorrowerTypeIdentifier() {}

type BorrowerTypeProfileResponse struct {
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
	NameUz           string `protobuf:"bytes,9,opt,name=name_uz,json=nameUz,proto3" json:"name_uz,omitempty"`
	DeletedAt        string `protobuf:"bytes,10,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *BorrowerTypeProfileResponse) Reset() {
	*x = BorrowerTypeProfileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_borrower_type_borrower_type_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BorrowerTypeProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BorrowerTypeProfileResponse) ProtoMessage() {}

func (x *BorrowerTypeProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_borrower_type_borrower_type_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BorrowerTypeProfileResponse.ProtoReflect.Descriptor instead.
func (*BorrowerTypeProfileResponse) Descriptor() ([]byte, []int) {
	return file_borrower_type_borrower_type_proto_rawDescGZIP(), []int{4}
}

func (x *BorrowerTypeProfileResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BorrowerTypeProfileResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *BorrowerTypeProfileResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BorrowerTypeProfileResponse) GetOldCode() string {
	if x != nil {
		return x.OldCode
	}
	return ""
}

func (x *BorrowerTypeProfileResponse) GetOldName() string {
	if x != nil {
		return x.OldName
	}
	return ""
}

func (x *BorrowerTypeProfileResponse) GetActivationDate() string {
	if x != nil {
		return x.ActivationDate
	}
	return ""
}

func (x *BorrowerTypeProfileResponse) GetDeactivationDate() string {
	if x != nil {
		return x.DeactivationDate
	}
	return ""
}

func (x *BorrowerTypeProfileResponse) GetFlexFinId() string {
	if x != nil {
		return x.FlexFinId
	}
	return ""
}

func (x *BorrowerTypeProfileResponse) GetNameUz() string {
	if x != nil {
		return x.NameUz
	}
	return ""
}

func (x *BorrowerTypeProfileResponse) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

type BorrowerTypeSuccessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *BorrowerTypeSuccessResponse) Reset() {
	*x = BorrowerTypeSuccessResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_borrower_type_borrower_type_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BorrowerTypeSuccessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BorrowerTypeSuccessResponse) ProtoMessage() {}

func (x *BorrowerTypeSuccessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_borrower_type_borrower_type_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BorrowerTypeSuccessResponse.ProtoReflect.Descriptor instead.
func (*BorrowerTypeSuccessResponse) Descriptor() ([]byte, []int) {
	return file_borrower_type_borrower_type_proto_rawDescGZIP(), []int{5}
}

func (x *BorrowerTypeSuccessResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type UpdateBorrowerTypeRequest struct {
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
	NameUz           string `protobuf:"bytes,9,opt,name=name_uz,json=nameUz,proto3" json:"name_uz,omitempty"`
}

func (x *UpdateBorrowerTypeRequest) Reset() {
	*x = UpdateBorrowerTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_borrower_type_borrower_type_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBorrowerTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBorrowerTypeRequest) ProtoMessage() {}

func (x *UpdateBorrowerTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_borrower_type_borrower_type_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBorrowerTypeRequest.ProtoReflect.Descriptor instead.
func (*UpdateBorrowerTypeRequest) Descriptor() ([]byte, []int) {
	return file_borrower_type_borrower_type_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateBorrowerTypeRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateBorrowerTypeRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *UpdateBorrowerTypeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateBorrowerTypeRequest) GetOldCode() string {
	if x != nil {
		return x.OldCode
	}
	return ""
}

func (x *UpdateBorrowerTypeRequest) GetOldName() string {
	if x != nil {
		return x.OldName
	}
	return ""
}

func (x *UpdateBorrowerTypeRequest) GetActivationDate() string {
	if x != nil {
		return x.ActivationDate
	}
	return ""
}

func (x *UpdateBorrowerTypeRequest) GetDeactivationDate() string {
	if x != nil {
		return x.DeactivationDate
	}
	return ""
}

func (x *UpdateBorrowerTypeRequest) GetFlexFinId() string {
	if x != nil {
		return x.FlexFinId
	}
	return ""
}

func (x *UpdateBorrowerTypeRequest) GetNameUz() string {
	if x != nil {
		return x.NameUz
	}
	return ""
}

var File_borrower_type_borrower_type_proto protoreflect.FileDescriptor

var file_borrower_type_borrower_type_proto_rawDesc = []byte{
	0x0a, 0x21, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2f,
	0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x57, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x72, 0x72, 0x6f,
	0x77, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x5f, 0x0a, 0x18,
	0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0e, 0x62, 0x6f, 0x72, 0x72,
	0x6f, 0x77, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0d,
	0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x73, 0x22, 0x88, 0x02,
	0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x72,
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
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x6c, 0x65, 0x78, 0x46, 0x69, 0x6e, 0x49, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x75, 0x7a, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6e, 0x61, 0x6d, 0x65, 0x55, 0x7a, 0x22, 0x5f, 0x0a, 0x19, 0x53, 0x69, 0x6e, 0x67,
	0x6c, 0x65, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x1a, 0x0a,
	0x18, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x22, 0xb9, 0x02, 0x0a, 0x1b, 0x42, 0x6f,
	0x72, 0x72, 0x6f, 0x77, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x6c, 0x64, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x6c, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x6f, 0x6c, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6f, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65,
	0x12, 0x2b, 0x0a, 0x11, 0x64, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x64, 0x65, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x0a,
	0x0b, 0x66, 0x6c, 0x65, 0x78, 0x5f, 0x66, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x66, 0x6c, 0x65, 0x78, 0x46, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x75, 0x7a, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6e, 0x61, 0x6d, 0x65, 0x55, 0x7a, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x39, 0x0a, 0x1b, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x98, 0x02, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x72, 0x72, 0x6f,
	0x77, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x6c, 0x64, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x6c, 0x64, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x6c, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x64, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x64, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61,
	0x74, 0x65, 0x12, 0x1e, 0x0a, 0x0b, 0x66, 0x6c, 0x65, 0x78, 0x5f, 0x66, 0x69, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x6c, 0x65, 0x78, 0x46, 0x69, 0x6e,
	0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x75, 0x7a, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x61, 0x6d, 0x65, 0x55, 0x7a, 0x32, 0xdf, 0x02, 0x0a, 0x13,
	0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x72, 0x72,
	0x6f, 0x77, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x42, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1a, 0x2e, 0x53, 0x69,
	0x6e, 0x67, 0x6c, 0x65, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77,
	0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x1a, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x42, 0x6f,
	0x72, 0x72, 0x6f, 0x77, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x42, 0x6f, 0x72, 0x72,
	0x6f, 0x77, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x15, 0x5a,
	0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x65, 0x72, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_borrower_type_borrower_type_proto_rawDescOnce sync.Once
	file_borrower_type_borrower_type_proto_rawDescData = file_borrower_type_borrower_type_proto_rawDesc
)

func file_borrower_type_borrower_type_proto_rawDescGZIP() []byte {
	file_borrower_type_borrower_type_proto_rawDescOnce.Do(func() {
		file_borrower_type_borrower_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_borrower_type_borrower_type_proto_rawDescData)
	})
	return file_borrower_type_borrower_type_proto_rawDescData
}

var file_borrower_type_borrower_type_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_borrower_type_borrower_type_proto_goTypes = []interface{}{
	(*ListBorrowerTypeRequest)(nil),     // 0: ListBorrowerTypeRequest
	(*ListBorrowerTypeResponse)(nil),    // 1: ListBorrowerTypeResponse
	(*CreateBorrowerTypeRequest)(nil),   // 2: CreateBorrowerTypeRequest
	(*SingleBorrowerTypeRequest)(nil),   // 3: SingleBorrowerTypeRequest
	(*BorrowerTypeProfileResponse)(nil), // 4: BorrowerTypeProfileResponse
	(*BorrowerTypeSuccessResponse)(nil), // 5: BorrowerTypeSuccessResponse
	(*UpdateBorrowerTypeRequest)(nil),   // 6: UpdateBorrowerTypeRequest
}
var file_borrower_type_borrower_type_proto_depIdxs = []int32{
	4, // 0: ListBorrowerTypeResponse.borrower_types:type_name -> BorrowerTypeProfileResponse
	0, // 1: BorrowerTypeService.List:input_type -> ListBorrowerTypeRequest
	2, // 2: BorrowerTypeService.Create:input_type -> CreateBorrowerTypeRequest
	3, // 3: BorrowerTypeService.Get:input_type -> SingleBorrowerTypeRequest
	6, // 4: BorrowerTypeService.Update:input_type -> UpdateBorrowerTypeRequest
	3, // 5: BorrowerTypeService.Delete:input_type -> SingleBorrowerTypeRequest
	1, // 6: BorrowerTypeService.List:output_type -> ListBorrowerTypeResponse
	4, // 7: BorrowerTypeService.Create:output_type -> BorrowerTypeProfileResponse
	4, // 8: BorrowerTypeService.Get:output_type -> BorrowerTypeProfileResponse
	5, // 9: BorrowerTypeService.Update:output_type -> BorrowerTypeSuccessResponse
	5, // 10: BorrowerTypeService.Delete:output_type -> BorrowerTypeSuccessResponse
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_borrower_type_borrower_type_proto_init() }
func file_borrower_type_borrower_type_proto_init() {
	if File_borrower_type_borrower_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_borrower_type_borrower_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBorrowerTypeRequest); i {
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
		file_borrower_type_borrower_type_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBorrowerTypeResponse); i {
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
		file_borrower_type_borrower_type_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBorrowerTypeRequest); i {
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
		file_borrower_type_borrower_type_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SingleBorrowerTypeRequest); i {
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
		file_borrower_type_borrower_type_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BorrowerTypeProfileResponse); i {
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
		file_borrower_type_borrower_type_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BorrowerTypeSuccessResponse); i {
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
		file_borrower_type_borrower_type_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBorrowerTypeRequest); i {
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
	file_borrower_type_borrower_type_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*SingleBorrowerTypeRequest_Id)(nil),
		(*SingleBorrowerTypeRequest_Code)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_borrower_type_borrower_type_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_borrower_type_borrower_type_proto_goTypes,
		DependencyIndexes: file_borrower_type_borrower_type_proto_depIdxs,
		MessageInfos:      file_borrower_type_borrower_type_proto_msgTypes,
	}.Build()
	File_borrower_type_borrower_type_proto = out.File
	file_borrower_type_borrower_type_proto_rawDesc = nil
	file_borrower_type_borrower_type_proto_goTypes = nil
	file_borrower_type_borrower_type_proto_depIdxs = nil
}
