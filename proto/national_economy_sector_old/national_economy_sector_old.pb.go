// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.12
// source: national_economy_sector_old/national_economy_sector_old.proto

package proto_national_economy_sector_old

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

type ListNationalEconomySectorOldRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageNumber int32 `protobuf:"varint,1,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	PageSize   int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *ListNationalEconomySectorOldRequest) Reset() {
	*x = ListNationalEconomySectorOldRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_national_economy_sector_old_national_economy_sector_old_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNationalEconomySectorOldRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNationalEconomySectorOldRequest) ProtoMessage() {}

func (x *ListNationalEconomySectorOldRequest) ProtoReflect() protoreflect.Message {
	mi := &file_national_economy_sector_old_national_economy_sector_old_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNationalEconomySectorOldRequest.ProtoReflect.Descriptor instead.
func (*ListNationalEconomySectorOldRequest) Descriptor() ([]byte, []int) {
	return file_national_economy_sector_old_national_economy_sector_old_proto_rawDescGZIP(), []int{0}
}

func (x *ListNationalEconomySectorOldRequest) GetPageNumber() int32 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *ListNationalEconomySectorOldRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListNationalEconomySectorOldResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OldNationalEconomySectors []*NationalEconomySectorOldProfileResponse `protobuf:"bytes,1,rep,name=old_national_economy_sectors,json=oldNationalEconomySectors,proto3" json:"old_national_economy_sectors,omitempty"`
}

func (x *ListNationalEconomySectorOldResponse) Reset() {
	*x = ListNationalEconomySectorOldResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_national_economy_sector_old_national_economy_sector_old_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNationalEconomySectorOldResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNationalEconomySectorOldResponse) ProtoMessage() {}

func (x *ListNationalEconomySectorOldResponse) ProtoReflect() protoreflect.Message {
	mi := &file_national_economy_sector_old_national_economy_sector_old_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNationalEconomySectorOldResponse.ProtoReflect.Descriptor instead.
func (*ListNationalEconomySectorOldResponse) Descriptor() ([]byte, []int) {
	return file_national_economy_sector_old_national_economy_sector_old_proto_rawDescGZIP(), []int{1}
}

func (x *ListNationalEconomySectorOldResponse) GetOldNationalEconomySectors() []*NationalEconomySectorOldProfileResponse {
	if x != nil {
		return x.OldNationalEconomySectors
	}
	return nil
}

type CreateNationalEconomySectorOldRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code             string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Name             string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CbuCode          int64  `protobuf:"varint,3,opt,name=cbu_code,json=cbuCode,proto3" json:"cbu_code,omitempty"`
	CbuGroupCode     int64  `protobuf:"varint,4,opt,name=cbu_group_code,json=cbuGroupCode,proto3" json:"cbu_group_code,omitempty"`
	ActivationDate   string `protobuf:"bytes,5,opt,name=activation_date,json=activationDate,proto3" json:"activation_date,omitempty"`
	DeactivationDate string `protobuf:"bytes,6,opt,name=deactivation_date,json=deactivationDate,proto3" json:"deactivation_date,omitempty"`
	CbuReferenceKey  int32  `protobuf:"varint,7,opt,name=cbu_reference_key,json=cbuReferenceKey,proto3" json:"cbu_reference_key,omitempty"`
	FlexFinId        string `protobuf:"bytes,8,opt,name=flex_fin_id,json=flexFinId,proto3" json:"flex_fin_id,omitempty"`
}

func (x *CreateNationalEconomySectorOldRequest) Reset() {
	*x = CreateNationalEconomySectorOldRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_national_economy_sector_old_national_economy_sector_old_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNationalEconomySectorOldRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNationalEconomySectorOldRequest) ProtoMessage() {}

func (x *CreateNationalEconomySectorOldRequest) ProtoReflect() protoreflect.Message {
	mi := &file_national_economy_sector_old_national_economy_sector_old_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNationalEconomySectorOldRequest.ProtoReflect.Descriptor instead.
func (*CreateNationalEconomySectorOldRequest) Descriptor() ([]byte, []int) {
	return file_national_economy_sector_old_national_economy_sector_old_proto_rawDescGZIP(), []int{2}
}

func (x *CreateNationalEconomySectorOldRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CreateNationalEconomySectorOldRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateNationalEconomySectorOldRequest) GetCbuCode() int64 {
	if x != nil {
		return x.CbuCode
	}
	return 0
}

func (x *CreateNationalEconomySectorOldRequest) GetCbuGroupCode() int64 {
	if x != nil {
		return x.CbuGroupCode
	}
	return 0
}

func (x *CreateNationalEconomySectorOldRequest) GetActivationDate() string {
	if x != nil {
		return x.ActivationDate
	}
	return ""
}

func (x *CreateNationalEconomySectorOldRequest) GetDeactivationDate() string {
	if x != nil {
		return x.DeactivationDate
	}
	return ""
}

func (x *CreateNationalEconomySectorOldRequest) GetCbuReferenceKey() int32 {
	if x != nil {
		return x.CbuReferenceKey
	}
	return 0
}

func (x *CreateNationalEconomySectorOldRequest) GetFlexFinId() string {
	if x != nil {
		return x.FlexFinId
	}
	return ""
}

type SingleNationalEconomySectorOldRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to OldNationalEconomySectorIdentifier:
	//
	//	*SingleNationalEconomySectorOldRequest_Id
	//	*SingleNationalEconomySectorOldRequest_Code
	OldNationalEconomySectorIdentifier isSingleNationalEconomySectorOldRequest_OldNationalEconomySectorIdentifier `protobuf_oneof:"old_national_economy_sector_identifier"`
}

func (x *SingleNationalEconomySectorOldRequest) Reset() {
	*x = SingleNationalEconomySectorOldRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_national_economy_sector_old_national_economy_sector_old_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SingleNationalEconomySectorOldRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SingleNationalEconomySectorOldRequest) ProtoMessage() {}

func (x *SingleNationalEconomySectorOldRequest) ProtoReflect() protoreflect.Message {
	mi := &file_national_economy_sector_old_national_economy_sector_old_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SingleNationalEconomySectorOldRequest.ProtoReflect.Descriptor instead.
func (*SingleNationalEconomySectorOldRequest) Descriptor() ([]byte, []int) {
	return file_national_economy_sector_old_national_economy_sector_old_proto_rawDescGZIP(), []int{3}
}

func (m *SingleNationalEconomySectorOldRequest) GetOldNationalEconomySectorIdentifier() isSingleNationalEconomySectorOldRequest_OldNationalEconomySectorIdentifier {
	if m != nil {
		return m.OldNationalEconomySectorIdentifier
	}
	return nil
}

func (x *SingleNationalEconomySectorOldRequest) GetId() string {
	if x, ok := x.GetOldNationalEconomySectorIdentifier().(*SingleNationalEconomySectorOldRequest_Id); ok {
		return x.Id
	}
	return ""
}

func (x *SingleNationalEconomySectorOldRequest) GetCode() string {
	if x, ok := x.GetOldNationalEconomySectorIdentifier().(*SingleNationalEconomySectorOldRequest_Code); ok {
		return x.Code
	}
	return ""
}

type isSingleNationalEconomySectorOldRequest_OldNationalEconomySectorIdentifier interface {
	isSingleNationalEconomySectorOldRequest_OldNationalEconomySectorIdentifier()
}

type SingleNationalEconomySectorOldRequest_Id struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3,oneof"`
}

type SingleNationalEconomySectorOldRequest_Code struct {
	Code string `protobuf:"bytes,2,opt,name=code,proto3,oneof"`
}

func (*SingleNationalEconomySectorOldRequest_Id) isSingleNationalEconomySectorOldRequest_OldNationalEconomySectorIdentifier() {
}

func (*SingleNationalEconomySectorOldRequest_Code) isSingleNationalEconomySectorOldRequest_OldNationalEconomySectorIdentifier() {
}

type NationalEconomySectorOldProfileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Code             string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Name             string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	CbuCode          int64  `protobuf:"varint,4,opt,name=cbu_code,json=cbuCode,proto3" json:"cbu_code,omitempty"`
	CbuGroupCode     int64  `protobuf:"varint,5,opt,name=cbu_group_code,json=cbuGroupCode,proto3" json:"cbu_group_code,omitempty"`
	ActivationDate   string `protobuf:"bytes,6,opt,name=activation_date,json=activationDate,proto3" json:"activation_date,omitempty"`
	DeactivationDate string `protobuf:"bytes,7,opt,name=deactivation_date,json=deactivationDate,proto3" json:"deactivation_date,omitempty"`
	CbuReferenceKey  int32  `protobuf:"varint,8,opt,name=cbu_reference_key,json=cbuReferenceKey,proto3" json:"cbu_reference_key,omitempty"`
	FlexFinId        string `protobuf:"bytes,9,opt,name=flex_fin_id,json=flexFinId,proto3" json:"flex_fin_id,omitempty"`
	DeletedAt        string `protobuf:"bytes,10,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *NationalEconomySectorOldProfileResponse) Reset() {
	*x = NationalEconomySectorOldProfileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_national_economy_sector_old_national_economy_sector_old_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NationalEconomySectorOldProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NationalEconomySectorOldProfileResponse) ProtoMessage() {}

func (x *NationalEconomySectorOldProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_national_economy_sector_old_national_economy_sector_old_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NationalEconomySectorOldProfileResponse.ProtoReflect.Descriptor instead.
func (*NationalEconomySectorOldProfileResponse) Descriptor() ([]byte, []int) {
	return file_national_economy_sector_old_national_economy_sector_old_proto_rawDescGZIP(), []int{4}
}

func (x *NationalEconomySectorOldProfileResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NationalEconomySectorOldProfileResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *NationalEconomySectorOldProfileResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NationalEconomySectorOldProfileResponse) GetCbuCode() int64 {
	if x != nil {
		return x.CbuCode
	}
	return 0
}

func (x *NationalEconomySectorOldProfileResponse) GetCbuGroupCode() int64 {
	if x != nil {
		return x.CbuGroupCode
	}
	return 0
}

func (x *NationalEconomySectorOldProfileResponse) GetActivationDate() string {
	if x != nil {
		return x.ActivationDate
	}
	return ""
}

func (x *NationalEconomySectorOldProfileResponse) GetDeactivationDate() string {
	if x != nil {
		return x.DeactivationDate
	}
	return ""
}

func (x *NationalEconomySectorOldProfileResponse) GetCbuReferenceKey() int32 {
	if x != nil {
		return x.CbuReferenceKey
	}
	return 0
}

func (x *NationalEconomySectorOldProfileResponse) GetFlexFinId() string {
	if x != nil {
		return x.FlexFinId
	}
	return ""
}

func (x *NationalEconomySectorOldProfileResponse) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

type NationalEconomySectorOldSuccessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *NationalEconomySectorOldSuccessResponse) Reset() {
	*x = NationalEconomySectorOldSuccessResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_national_economy_sector_old_national_economy_sector_old_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NationalEconomySectorOldSuccessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NationalEconomySectorOldSuccessResponse) ProtoMessage() {}

func (x *NationalEconomySectorOldSuccessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_national_economy_sector_old_national_economy_sector_old_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NationalEconomySectorOldSuccessResponse.ProtoReflect.Descriptor instead.
func (*NationalEconomySectorOldSuccessResponse) Descriptor() ([]byte, []int) {
	return file_national_economy_sector_old_national_economy_sector_old_proto_rawDescGZIP(), []int{5}
}

func (x *NationalEconomySectorOldSuccessResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type UpdateNationalEconomySectorOldRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Code             string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Name             string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	CbuCode          int64  `protobuf:"varint,4,opt,name=cbu_code,json=cbuCode,proto3" json:"cbu_code,omitempty"`
	CbuGroupCode     int64  `protobuf:"varint,5,opt,name=cbu_group_code,json=cbuGroupCode,proto3" json:"cbu_group_code,omitempty"`
	ActivationDate   string `protobuf:"bytes,6,opt,name=activation_date,json=activationDate,proto3" json:"activation_date,omitempty"`
	DeactivationDate string `protobuf:"bytes,7,opt,name=deactivation_date,json=deactivationDate,proto3" json:"deactivation_date,omitempty"`
	CbuReferenceKey  int32  `protobuf:"varint,8,opt,name=cbu_reference_key,json=cbuReferenceKey,proto3" json:"cbu_reference_key,omitempty"`
	FlexFinId        string `protobuf:"bytes,9,opt,name=flex_fin_id,json=flexFinId,proto3" json:"flex_fin_id,omitempty"`
}

func (x *UpdateNationalEconomySectorOldRequest) Reset() {
	*x = UpdateNationalEconomySectorOldRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_national_economy_sector_old_national_economy_sector_old_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNationalEconomySectorOldRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNationalEconomySectorOldRequest) ProtoMessage() {}

func (x *UpdateNationalEconomySectorOldRequest) ProtoReflect() protoreflect.Message {
	mi := &file_national_economy_sector_old_national_economy_sector_old_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNationalEconomySectorOldRequest.ProtoReflect.Descriptor instead.
func (*UpdateNationalEconomySectorOldRequest) Descriptor() ([]byte, []int) {
	return file_national_economy_sector_old_national_economy_sector_old_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateNationalEconomySectorOldRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateNationalEconomySectorOldRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *UpdateNationalEconomySectorOldRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateNationalEconomySectorOldRequest) GetCbuCode() int64 {
	if x != nil {
		return x.CbuCode
	}
	return 0
}

func (x *UpdateNationalEconomySectorOldRequest) GetCbuGroupCode() int64 {
	if x != nil {
		return x.CbuGroupCode
	}
	return 0
}

func (x *UpdateNationalEconomySectorOldRequest) GetActivationDate() string {
	if x != nil {
		return x.ActivationDate
	}
	return ""
}

func (x *UpdateNationalEconomySectorOldRequest) GetDeactivationDate() string {
	if x != nil {
		return x.DeactivationDate
	}
	return ""
}

func (x *UpdateNationalEconomySectorOldRequest) GetCbuReferenceKey() int32 {
	if x != nil {
		return x.CbuReferenceKey
	}
	return 0
}

func (x *UpdateNationalEconomySectorOldRequest) GetFlexFinId() string {
	if x != nil {
		return x.FlexFinId
	}
	return ""
}

var File_national_economy_sector_old_national_economy_sector_old_proto protoreflect.FileDescriptor

var file_national_economy_sector_old_national_economy_sector_old_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x65, 0x63, 0x6f, 0x6e, 0x6f,
	0x6d, 0x79, 0x5f, 0x73, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x6f, 0x6c, 0x64, 0x2f, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x65, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x5f, 0x73,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x6f, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x63, 0x0a, 0x23, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x45,
	0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x53, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x4f, 0x6c, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x61, 0x67,
	0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x22, 0x91, 0x01, 0x0a, 0x24, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x45, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x53, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x4f, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x69, 0x0a,
	0x1c, 0x6f, 0x6c, 0x64, 0x5f, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x65, 0x63,
	0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x5f, 0x73, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x4e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x45, 0x63,
	0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x53, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x4f, 0x6c, 0x64, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x19, 0x6f,
	0x6c, 0x64, 0x4e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x45, 0x63, 0x6f, 0x6e, 0x6f, 0x6d,
	0x79, 0x53, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x22, 0xb2, 0x02, 0x0a, 0x25, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x45, 0x63, 0x6f, 0x6e, 0x6f,
	0x6d, 0x79, 0x53, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x4f, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x62,
	0x75, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x62,
	0x75, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x63, 0x62, 0x75, 0x5f, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63,
	0x62, 0x75, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x64, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x64, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x2a, 0x0a, 0x11, 0x63, 0x62, 0x75, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x63, 0x62,
	0x75, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x1e, 0x0a,
	0x0b, 0x66, 0x6c, 0x65, 0x78, 0x5f, 0x66, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x66, 0x6c, 0x65, 0x78, 0x46, 0x69, 0x6e, 0x49, 0x64, 0x22, 0x79, 0x0a,
	0x25, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x4e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x45,
	0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x53, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x4f, 0x6c, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x28,
	0x0a, 0x26, 0x6f, 0x6c, 0x64, 0x5f, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x65,
	0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x5f, 0x73, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x22, 0xe3, 0x02, 0x0a, 0x27, 0x4e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x45, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x53, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x4f, 0x6c, 0x64, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x63, 0x62, 0x75, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x63, 0x62, 0x75, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x63, 0x62, 0x75, 0x5f, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0c, 0x63, 0x62, 0x75, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x27, 0x0a,
	0x0f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x64, 0x65, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x64, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x63, 0x62, 0x75, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f,
	0x63, 0x62, 0x75, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x4b, 0x65, 0x79, 0x12,
	0x1e, 0x0a, 0x0b, 0x66, 0x6c, 0x65, 0x78, 0x5f, 0x66, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x6c, 0x65, 0x78, 0x46, 0x69, 0x6e, 0x49, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x45,
	0x0a, 0x27, 0x4e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x45, 0x63, 0x6f, 0x6e, 0x6f, 0x6d,
	0x79, 0x53, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x4f, 0x6c, 0x64, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xc2, 0x02, 0x0a, 0x25, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x45, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x53,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x4f, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x62, 0x75, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x62, 0x75, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x63, 0x62, 0x75, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x62, 0x75, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x2b, 0x0a, 0x11, 0x64, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x64, 0x65,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x2a,
	0x0a, 0x11, 0x63, 0x62, 0x75, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x63, 0x62, 0x75, 0x52, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x1e, 0x0a, 0x0b, 0x66, 0x6c,
	0x65, 0x78, 0x5f, 0x66, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x66, 0x6c, 0x65, 0x78, 0x46, 0x69, 0x6e, 0x49, 0x64, 0x32, 0xe3, 0x03, 0x0a, 0x1f, 0x4e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x45, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x53, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x4f, 0x6c, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53,
	0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x24, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x45, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x53, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x4f, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x4e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x45, 0x63, 0x6f, 0x6e, 0x6f,
	0x6d, 0x79, 0x53, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x4f, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x26, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x45, 0x63,
	0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x53, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x4f, 0x6c, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x4e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x45, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x53, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x4f, 0x6c, 0x64,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x57, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x26, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x4e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x45, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x53, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x4f, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28,
	0x2e, 0x4e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x45, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79,
	0x53, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x4f, 0x6c, 0x64, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x26, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x45, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x53, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x4f, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x4e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x45, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x53, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x4f, 0x6c, 0x64, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x26,
	0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x4e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x45,
	0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x53, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x4f, 0x6c, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x4e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x45, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x53, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x4f, 0x6c,
	0x64, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x23, 0x5a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x2e, 0x65, 0x63, 0x6f, 0x6e, 0x6f, 0x6d, 0x79, 0x2e, 0x73, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x2e, 0x6f, 0x6c, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_national_economy_sector_old_national_economy_sector_old_proto_rawDescOnce sync.Once
	file_national_economy_sector_old_national_economy_sector_old_proto_rawDescData = file_national_economy_sector_old_national_economy_sector_old_proto_rawDesc
)

func file_national_economy_sector_old_national_economy_sector_old_proto_rawDescGZIP() []byte {
	file_national_economy_sector_old_national_economy_sector_old_proto_rawDescOnce.Do(func() {
		file_national_economy_sector_old_national_economy_sector_old_proto_rawDescData = protoimpl.X.CompressGZIP(file_national_economy_sector_old_national_economy_sector_old_proto_rawDescData)
	})
	return file_national_economy_sector_old_national_economy_sector_old_proto_rawDescData
}

var file_national_economy_sector_old_national_economy_sector_old_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_national_economy_sector_old_national_economy_sector_old_proto_goTypes = []interface{}{
	(*ListNationalEconomySectorOldRequest)(nil),     // 0: ListNationalEconomySectorOldRequest
	(*ListNationalEconomySectorOldResponse)(nil),    // 1: ListNationalEconomySectorOldResponse
	(*CreateNationalEconomySectorOldRequest)(nil),   // 2: CreateNationalEconomySectorOldRequest
	(*SingleNationalEconomySectorOldRequest)(nil),   // 3: SingleNationalEconomySectorOldRequest
	(*NationalEconomySectorOldProfileResponse)(nil), // 4: NationalEconomySectorOldProfileResponse
	(*NationalEconomySectorOldSuccessResponse)(nil), // 5: NationalEconomySectorOldSuccessResponse
	(*UpdateNationalEconomySectorOldRequest)(nil),   // 6: UpdateNationalEconomySectorOldRequest
}
var file_national_economy_sector_old_national_economy_sector_old_proto_depIdxs = []int32{
	4, // 0: ListNationalEconomySectorOldResponse.old_national_economy_sectors:type_name -> NationalEconomySectorOldProfileResponse
	0, // 1: NationalEconomySectorOldService.List:input_type -> ListNationalEconomySectorOldRequest
	2, // 2: NationalEconomySectorOldService.Create:input_type -> CreateNationalEconomySectorOldRequest
	3, // 3: NationalEconomySectorOldService.Get:input_type -> SingleNationalEconomySectorOldRequest
	6, // 4: NationalEconomySectorOldService.Update:input_type -> UpdateNationalEconomySectorOldRequest
	3, // 5: NationalEconomySectorOldService.Delete:input_type -> SingleNationalEconomySectorOldRequest
	1, // 6: NationalEconomySectorOldService.List:output_type -> ListNationalEconomySectorOldResponse
	4, // 7: NationalEconomySectorOldService.Create:output_type -> NationalEconomySectorOldProfileResponse
	4, // 8: NationalEconomySectorOldService.Get:output_type -> NationalEconomySectorOldProfileResponse
	5, // 9: NationalEconomySectorOldService.Update:output_type -> NationalEconomySectorOldSuccessResponse
	5, // 10: NationalEconomySectorOldService.Delete:output_type -> NationalEconomySectorOldSuccessResponse
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_national_economy_sector_old_national_economy_sector_old_proto_init() }
func file_national_economy_sector_old_national_economy_sector_old_proto_init() {
	if File_national_economy_sector_old_national_economy_sector_old_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_national_economy_sector_old_national_economy_sector_old_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNationalEconomySectorOldRequest); i {
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
		file_national_economy_sector_old_national_economy_sector_old_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNationalEconomySectorOldResponse); i {
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
		file_national_economy_sector_old_national_economy_sector_old_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNationalEconomySectorOldRequest); i {
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
		file_national_economy_sector_old_national_economy_sector_old_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SingleNationalEconomySectorOldRequest); i {
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
		file_national_economy_sector_old_national_economy_sector_old_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NationalEconomySectorOldProfileResponse); i {
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
		file_national_economy_sector_old_national_economy_sector_old_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NationalEconomySectorOldSuccessResponse); i {
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
		file_national_economy_sector_old_national_economy_sector_old_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNationalEconomySectorOldRequest); i {
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
	file_national_economy_sector_old_national_economy_sector_old_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*SingleNationalEconomySectorOldRequest_Id)(nil),
		(*SingleNationalEconomySectorOldRequest_Code)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_national_economy_sector_old_national_economy_sector_old_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_national_economy_sector_old_national_economy_sector_old_proto_goTypes,
		DependencyIndexes: file_national_economy_sector_old_national_economy_sector_old_proto_depIdxs,
		MessageInfos:      file_national_economy_sector_old_national_economy_sector_old_proto_msgTypes,
	}.Build()
	File_national_economy_sector_old_national_economy_sector_old_proto = out.File
	file_national_economy_sector_old_national_economy_sector_old_proto_rawDesc = nil
	file_national_economy_sector_old_national_economy_sector_old_proto_goTypes = nil
	file_national_economy_sector_old_national_economy_sector_old_proto_depIdxs = nil
}
