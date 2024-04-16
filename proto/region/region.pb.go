// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.12
// source: region.proto

package proto_region

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

type ListRegionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageNumber int32 `protobuf:"varint,1,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	PageSize   int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *ListRegionRequest) Reset() {
	*x = ListRegionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_region_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRegionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRegionRequest) ProtoMessage() {}

func (x *ListRegionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_region_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRegionRequest.ProtoReflect.Descriptor instead.
func (*ListRegionRequest) Descriptor() ([]byte, []int) {
	return file_region_proto_rawDescGZIP(), []int{0}
}

func (x *ListRegionRequest) GetPageNumber() int32 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *ListRegionRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListRegionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PaymentTypes []*RegionProfileResponse `protobuf:"bytes,1,rep,name=payment_types,json=paymentTypes,proto3" json:"payment_types,omitempty"`
}

func (x *ListRegionResponse) Reset() {
	*x = ListRegionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_region_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRegionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRegionResponse) ProtoMessage() {}

func (x *ListRegionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_region_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRegionResponse.ProtoReflect.Descriptor instead.
func (*ListRegionResponse) Descriptor() ([]byte, []int) {
	return file_region_proto_rawDescGZIP(), []int{1}
}

func (x *ListRegionResponse) GetPaymentTypes() []*RegionProfileResponse {
	if x != nil {
		return x.PaymentTypes
	}
	return nil
}

type CreateRegionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code             string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Name             string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CountryId        int64  `protobuf:"varint,3,opt,name=country_id,json=countryId,proto3" json:"country_id,omitempty"`
	ActivationDate   string `protobuf:"bytes,4,opt,name=activation_date,json=activationDate,proto3" json:"activation_date,omitempty"`
	DeactivationDate string `protobuf:"bytes,5,opt,name=deactivation_date,json=deactivationDate,proto3" json:"deactivation_date,omitempty"`
	OldCode          string `protobuf:"bytes,6,opt,name=old_code,json=oldCode,proto3" json:"old_code,omitempty"`
	OldName          string `protobuf:"bytes,7,opt,name=old_name,json=oldName,proto3" json:"old_name,omitempty"`
	FlexFinId        string `protobuf:"bytes,8,opt,name=flex_fin_id,json=flexFinId,proto3" json:"flex_fin_id,omitempty"`
}

func (x *CreateRegionRequest) Reset() {
	*x = CreateRegionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_region_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRegionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRegionRequest) ProtoMessage() {}

func (x *CreateRegionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_region_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRegionRequest.ProtoReflect.Descriptor instead.
func (*CreateRegionRequest) Descriptor() ([]byte, []int) {
	return file_region_proto_rawDescGZIP(), []int{2}
}

func (x *CreateRegionRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CreateRegionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateRegionRequest) GetCountryId() int64 {
	if x != nil {
		return x.CountryId
	}
	return 0
}

func (x *CreateRegionRequest) GetActivationDate() string {
	if x != nil {
		return x.ActivationDate
	}
	return ""
}

func (x *CreateRegionRequest) GetDeactivationDate() string {
	if x != nil {
		return x.DeactivationDate
	}
	return ""
}

func (x *CreateRegionRequest) GetOldCode() string {
	if x != nil {
		return x.OldCode
	}
	return ""
}

func (x *CreateRegionRequest) GetOldName() string {
	if x != nil {
		return x.OldName
	}
	return ""
}

func (x *CreateRegionRequest) GetFlexFinId() string {
	if x != nil {
		return x.FlexFinId
	}
	return ""
}

type SingleRegionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to RegionIdentifier:
	//
	//	*SingleRegionRequest_Id
	//	*SingleRegionRequest_Code
	RegionIdentifier isSingleRegionRequest_RegionIdentifier `protobuf_oneof:"region_identifier"`
}

func (x *SingleRegionRequest) Reset() {
	*x = SingleRegionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_region_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SingleRegionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SingleRegionRequest) ProtoMessage() {}

func (x *SingleRegionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_region_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SingleRegionRequest.ProtoReflect.Descriptor instead.
func (*SingleRegionRequest) Descriptor() ([]byte, []int) {
	return file_region_proto_rawDescGZIP(), []int{3}
}

func (m *SingleRegionRequest) GetRegionIdentifier() isSingleRegionRequest_RegionIdentifier {
	if m != nil {
		return m.RegionIdentifier
	}
	return nil
}

func (x *SingleRegionRequest) GetId() string {
	if x, ok := x.GetRegionIdentifier().(*SingleRegionRequest_Id); ok {
		return x.Id
	}
	return ""
}

func (x *SingleRegionRequest) GetCode() string {
	if x, ok := x.GetRegionIdentifier().(*SingleRegionRequest_Code); ok {
		return x.Code
	}
	return ""
}

type isSingleRegionRequest_RegionIdentifier interface {
	isSingleRegionRequest_RegionIdentifier()
}

type SingleRegionRequest_Id struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3,oneof"`
}

type SingleRegionRequest_Code struct {
	Code string `protobuf:"bytes,2,opt,name=code,proto3,oneof"`
}

func (*SingleRegionRequest_Id) isSingleRegionRequest_RegionIdentifier() {}

func (*SingleRegionRequest_Code) isSingleRegionRequest_RegionIdentifier() {}

type RegionProfileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Code             string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Name             string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	CountryId        int64  `protobuf:"varint,4,opt,name=country_id,json=countryId,proto3" json:"country_id,omitempty"`
	ActivationDate   string `protobuf:"bytes,5,opt,name=activation_date,json=activationDate,proto3" json:"activation_date,omitempty"`
	DeactivationDate string `protobuf:"bytes,6,opt,name=deactivation_date,json=deactivationDate,proto3" json:"deactivation_date,omitempty"`
	OldCode          string `protobuf:"bytes,7,opt,name=old_code,json=oldCode,proto3" json:"old_code,omitempty"`
	OldName          string `protobuf:"bytes,8,opt,name=old_name,json=oldName,proto3" json:"old_name,omitempty"`
	FlexFinId        string `protobuf:"bytes,9,opt,name=flex_fin_id,json=flexFinId,proto3" json:"flex_fin_id,omitempty"`
	DeletedAt        string `protobuf:"bytes,10,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *RegionProfileResponse) Reset() {
	*x = RegionProfileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_region_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegionProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegionProfileResponse) ProtoMessage() {}

func (x *RegionProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_region_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegionProfileResponse.ProtoReflect.Descriptor instead.
func (*RegionProfileResponse) Descriptor() ([]byte, []int) {
	return file_region_proto_rawDescGZIP(), []int{4}
}

func (x *RegionProfileResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RegionProfileResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *RegionProfileResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RegionProfileResponse) GetCountryId() int64 {
	if x != nil {
		return x.CountryId
	}
	return 0
}

func (x *RegionProfileResponse) GetActivationDate() string {
	if x != nil {
		return x.ActivationDate
	}
	return ""
}

func (x *RegionProfileResponse) GetDeactivationDate() string {
	if x != nil {
		return x.DeactivationDate
	}
	return ""
}

func (x *RegionProfileResponse) GetOldCode() string {
	if x != nil {
		return x.OldCode
	}
	return ""
}

func (x *RegionProfileResponse) GetOldName() string {
	if x != nil {
		return x.OldName
	}
	return ""
}

func (x *RegionProfileResponse) GetFlexFinId() string {
	if x != nil {
		return x.FlexFinId
	}
	return ""
}

func (x *RegionProfileResponse) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

type SuccessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *SuccessResponse) Reset() {
	*x = SuccessResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_region_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuccessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuccessResponse) ProtoMessage() {}

func (x *SuccessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_region_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuccessResponse.ProtoReflect.Descriptor instead.
func (*SuccessResponse) Descriptor() ([]byte, []int) {
	return file_region_proto_rawDescGZIP(), []int{5}
}

func (x *SuccessResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type UpdateRegionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Code             string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Name             string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	CountryId        int64  `protobuf:"varint,4,opt,name=country_id,json=countryId,proto3" json:"country_id,omitempty"`
	ActivationDate   string `protobuf:"bytes,5,opt,name=activation_date,json=activationDate,proto3" json:"activation_date,omitempty"`
	DeactivationDate string `protobuf:"bytes,6,opt,name=deactivation_date,json=deactivationDate,proto3" json:"deactivation_date,omitempty"`
	OldCode          string `protobuf:"bytes,7,opt,name=old_code,json=oldCode,proto3" json:"old_code,omitempty"`
	OldName          string `protobuf:"bytes,8,opt,name=old_name,json=oldName,proto3" json:"old_name,omitempty"`
	FlexFinId        string `protobuf:"bytes,9,opt,name=flex_fin_id,json=flexFinId,proto3" json:"flex_fin_id,omitempty"`
}

func (x *UpdateRegionRequest) Reset() {
	*x = UpdateRegionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_region_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRegionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRegionRequest) ProtoMessage() {}

func (x *UpdateRegionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_region_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRegionRequest.ProtoReflect.Descriptor instead.
func (*UpdateRegionRequest) Descriptor() ([]byte, []int) {
	return file_region_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateRegionRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateRegionRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *UpdateRegionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateRegionRequest) GetCountryId() int64 {
	if x != nil {
		return x.CountryId
	}
	return 0
}

func (x *UpdateRegionRequest) GetActivationDate() string {
	if x != nil {
		return x.ActivationDate
	}
	return ""
}

func (x *UpdateRegionRequest) GetDeactivationDate() string {
	if x != nil {
		return x.DeactivationDate
	}
	return ""
}

func (x *UpdateRegionRequest) GetOldCode() string {
	if x != nil {
		return x.OldCode
	}
	return ""
}

func (x *UpdateRegionRequest) GetOldName() string {
	if x != nil {
		return x.OldName
	}
	return ""
}

func (x *UpdateRegionRequest) GetFlexFinId() string {
	if x != nil {
		return x.FlexFinId
	}
	return ""
}

var File_region_proto protoreflect.FileDescriptor

var file_region_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x51,
	0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x22, 0x51, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0c, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x73, 0x22, 0x88, 0x02, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x0a, 0x11,
	0x64, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x64, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x6c, 0x64,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x6c, 0x64,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x6c, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1e, 0x0a, 0x0b, 0x66, 0x6c, 0x65, 0x78, 0x5f, 0x66, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x6c, 0x65, 0x78, 0x46, 0x69, 0x6e, 0x49, 0x64, 0x22,
	0x52, 0x0a, 0x13, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x13,
	0x0a, 0x11, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x22, 0xb9, 0x02, 0x0a, 0x15, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x0a,
	0x11, 0x64, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x64, 0x65, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x6c,
	0x64, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x6c,
	0x64, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x6c, 0x64, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1e, 0x0a, 0x0b, 0x66, 0x6c, 0x65, 0x78, 0x5f, 0x66, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x6c, 0x65, 0x78, 0x46, 0x69, 0x6e, 0x49, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22,
	0x2d, 0x0a, 0x0f, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x98,
	0x02, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x12, 0x27, 0x0a,
	0x0f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x64, 0x65, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x64, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x6c, 0x64, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x6c, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x6f, 0x6c, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6f, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0b, 0x66, 0x6c, 0x65,
	0x78, 0x5f, 0x66, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x66, 0x6c, 0x65, 0x78, 0x46, 0x69, 0x6e, 0x49, 0x64, 0x32, 0x91, 0x02, 0x0a, 0x0d, 0x52, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x12, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x06,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x52,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x14, 0x2e, 0x53, 0x69,
	0x6e, 0x67, 0x6c, 0x65, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x53, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x52, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0e, 0x5a,
	0x0c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_region_proto_rawDescOnce sync.Once
	file_region_proto_rawDescData = file_region_proto_rawDesc
)

func file_region_proto_rawDescGZIP() []byte {
	file_region_proto_rawDescOnce.Do(func() {
		file_region_proto_rawDescData = protoimpl.X.CompressGZIP(file_region_proto_rawDescData)
	})
	return file_region_proto_rawDescData
}

var file_region_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_region_proto_goTypes = []interface{}{
	(*ListRegionRequest)(nil),     // 0: ListRegionRequest
	(*ListRegionResponse)(nil),    // 1: ListRegionResponse
	(*CreateRegionRequest)(nil),   // 2: CreateRegionRequest
	(*SingleRegionRequest)(nil),   // 3: SingleRegionRequest
	(*RegionProfileResponse)(nil), // 4: RegionProfileResponse
	(*SuccessResponse)(nil),       // 5: SuccessResponse
	(*UpdateRegionRequest)(nil),   // 6: UpdateRegionRequest
}
var file_region_proto_depIdxs = []int32{
	4, // 0: ListRegionResponse.payment_types:type_name -> RegionProfileResponse
	0, // 1: RegionService.List:input_type -> ListRegionRequest
	2, // 2: RegionService.Create:input_type -> CreateRegionRequest
	3, // 3: RegionService.Get:input_type -> SingleRegionRequest
	6, // 4: RegionService.Update:input_type -> UpdateRegionRequest
	3, // 5: RegionService.Delete:input_type -> SingleRegionRequest
	1, // 6: RegionService.List:output_type -> ListRegionResponse
	4, // 7: RegionService.Create:output_type -> RegionProfileResponse
	4, // 8: RegionService.Get:output_type -> RegionProfileResponse
	5, // 9: RegionService.Update:output_type -> SuccessResponse
	5, // 10: RegionService.Delete:output_type -> SuccessResponse
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_region_proto_init() }
func file_region_proto_init() {
	if File_region_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_region_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRegionRequest); i {
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
		file_region_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRegionResponse); i {
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
		file_region_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRegionRequest); i {
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
		file_region_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SingleRegionRequest); i {
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
		file_region_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegionProfileResponse); i {
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
		file_region_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuccessResponse); i {
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
		file_region_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRegionRequest); i {
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
	file_region_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*SingleRegionRequest_Id)(nil),
		(*SingleRegionRequest_Code)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_region_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_region_proto_goTypes,
		DependencyIndexes: file_region_proto_depIdxs,
		MessageInfos:      file_region_proto_msgTypes,
	}.Build()
	File_region_proto = out.File
	file_region_proto_rawDesc = nil
	file_region_proto_goTypes = nil
	file_region_proto_depIdxs = nil
}
