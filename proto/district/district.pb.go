// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: district/district.proto

package proto_district

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

type ListDistrictRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageNumber int32 `protobuf:"varint,1,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	PageSize   int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *ListDistrictRequest) Reset() {
	*x = ListDistrictRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_district_district_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDistrictRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDistrictRequest) ProtoMessage() {}

func (x *ListDistrictRequest) ProtoReflect() protoreflect.Message {
	mi := &file_district_district_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDistrictRequest.ProtoReflect.Descriptor instead.
func (*ListDistrictRequest) Descriptor() ([]byte, []int) {
	return file_district_district_proto_rawDescGZIP(), []int{0}
}

func (x *ListDistrictRequest) GetPageNumber() int32 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *ListDistrictRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListDistrictResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Districts []*DistrictProfileResponse `protobuf:"bytes,1,rep,name=districts,proto3" json:"districts,omitempty"`
}

func (x *ListDistrictResponse) Reset() {
	*x = ListDistrictResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_district_district_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDistrictResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDistrictResponse) ProtoMessage() {}

func (x *ListDistrictResponse) ProtoReflect() protoreflect.Message {
	mi := &file_district_district_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDistrictResponse.ProtoReflect.Descriptor instead.
func (*ListDistrictResponse) Descriptor() ([]byte, []int) {
	return file_district_district_proto_rawDescGZIP(), []int{1}
}

func (x *ListDistrictResponse) GetDistricts() []*DistrictProfileResponse {
	if x != nil {
		return x.Districts
	}
	return nil
}

type CreateDistrictRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code             string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Name             string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	RegionId         int64  `protobuf:"varint,3,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	OldCode          int64  `protobuf:"varint,4,opt,name=old_code,json=oldCode,proto3" json:"old_code,omitempty"`
	OldName          string `protobuf:"bytes,5,opt,name=old_name,json=oldName,proto3" json:"old_name,omitempty"`
	ActivationDate   string `protobuf:"bytes,6,opt,name=activation_date,json=activationDate,proto3" json:"activation_date,omitempty"`
	DeactivationDate string `protobuf:"bytes,7,opt,name=deactivation_date,json=deactivationDate,proto3" json:"deactivation_date,omitempty"`
	FlexFinId        string `protobuf:"bytes,8,opt,name=flex_fin_id,json=flexFinId,proto3" json:"flex_fin_id,omitempty"`
}

func (x *CreateDistrictRequest) Reset() {
	*x = CreateDistrictRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_district_district_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDistrictRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDistrictRequest) ProtoMessage() {}

func (x *CreateDistrictRequest) ProtoReflect() protoreflect.Message {
	mi := &file_district_district_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDistrictRequest.ProtoReflect.Descriptor instead.
func (*CreateDistrictRequest) Descriptor() ([]byte, []int) {
	return file_district_district_proto_rawDescGZIP(), []int{2}
}

func (x *CreateDistrictRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CreateDistrictRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateDistrictRequest) GetRegionId() int64 {
	if x != nil {
		return x.RegionId
	}
	return 0
}

func (x *CreateDistrictRequest) GetOldCode() int64 {
	if x != nil {
		return x.OldCode
	}
	return 0
}

func (x *CreateDistrictRequest) GetOldName() string {
	if x != nil {
		return x.OldName
	}
	return ""
}

func (x *CreateDistrictRequest) GetActivationDate() string {
	if x != nil {
		return x.ActivationDate
	}
	return ""
}

func (x *CreateDistrictRequest) GetDeactivationDate() string {
	if x != nil {
		return x.DeactivationDate
	}
	return ""
}

func (x *CreateDistrictRequest) GetFlexFinId() string {
	if x != nil {
		return x.FlexFinId
	}
	return ""
}

type SingleDistrictRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to DistrictIdentifier:
	//
	//	*SingleDistrictRequest_Id
	//	*SingleDistrictRequest_Code
	DistrictIdentifier isSingleDistrictRequest_DistrictIdentifier `protobuf_oneof:"district_identifier"`
}

func (x *SingleDistrictRequest) Reset() {
	*x = SingleDistrictRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_district_district_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SingleDistrictRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SingleDistrictRequest) ProtoMessage() {}

func (x *SingleDistrictRequest) ProtoReflect() protoreflect.Message {
	mi := &file_district_district_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SingleDistrictRequest.ProtoReflect.Descriptor instead.
func (*SingleDistrictRequest) Descriptor() ([]byte, []int) {
	return file_district_district_proto_rawDescGZIP(), []int{3}
}

func (m *SingleDistrictRequest) GetDistrictIdentifier() isSingleDistrictRequest_DistrictIdentifier {
	if m != nil {
		return m.DistrictIdentifier
	}
	return nil
}

func (x *SingleDistrictRequest) GetId() string {
	if x, ok := x.GetDistrictIdentifier().(*SingleDistrictRequest_Id); ok {
		return x.Id
	}
	return ""
}

func (x *SingleDistrictRequest) GetCode() string {
	if x, ok := x.GetDistrictIdentifier().(*SingleDistrictRequest_Code); ok {
		return x.Code
	}
	return ""
}

type isSingleDistrictRequest_DistrictIdentifier interface {
	isSingleDistrictRequest_DistrictIdentifier()
}

type SingleDistrictRequest_Id struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3,oneof"`
}

type SingleDistrictRequest_Code struct {
	Code string `protobuf:"bytes,2,opt,name=code,proto3,oneof"`
}

func (*SingleDistrictRequest_Id) isSingleDistrictRequest_DistrictIdentifier() {}

func (*SingleDistrictRequest_Code) isSingleDistrictRequest_DistrictIdentifier() {}

type DistrictProfileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Code             string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Name             string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	RegionId         int64  `protobuf:"varint,4,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	OldCode          int64  `protobuf:"varint,5,opt,name=old_code,json=oldCode,proto3" json:"old_code,omitempty"`
	OldName          string `protobuf:"bytes,6,opt,name=old_name,json=oldName,proto3" json:"old_name,omitempty"`
	ActivationDate   string `protobuf:"bytes,7,opt,name=activation_date,json=activationDate,proto3" json:"activation_date,omitempty"`
	DeactivationDate string `protobuf:"bytes,8,opt,name=deactivation_date,json=deactivationDate,proto3" json:"deactivation_date,omitempty"`
	FlexFinId        string `protobuf:"bytes,9,opt,name=flex_fin_id,json=flexFinId,proto3" json:"flex_fin_id,omitempty"`
	DeletedAt        int64  `protobuf:"varint,10,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *DistrictProfileResponse) Reset() {
	*x = DistrictProfileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_district_district_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DistrictProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DistrictProfileResponse) ProtoMessage() {}

func (x *DistrictProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_district_district_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DistrictProfileResponse.ProtoReflect.Descriptor instead.
func (*DistrictProfileResponse) Descriptor() ([]byte, []int) {
	return file_district_district_proto_rawDescGZIP(), []int{4}
}

func (x *DistrictProfileResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DistrictProfileResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *DistrictProfileResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DistrictProfileResponse) GetRegionId() int64 {
	if x != nil {
		return x.RegionId
	}
	return 0
}

func (x *DistrictProfileResponse) GetOldCode() int64 {
	if x != nil {
		return x.OldCode
	}
	return 0
}

func (x *DistrictProfileResponse) GetOldName() string {
	if x != nil {
		return x.OldName
	}
	return ""
}

func (x *DistrictProfileResponse) GetActivationDate() string {
	if x != nil {
		return x.ActivationDate
	}
	return ""
}

func (x *DistrictProfileResponse) GetDeactivationDate() string {
	if x != nil {
		return x.DeactivationDate
	}
	return ""
}

func (x *DistrictProfileResponse) GetFlexFinId() string {
	if x != nil {
		return x.FlexFinId
	}
	return ""
}

func (x *DistrictProfileResponse) GetDeletedAt() int64 {
	if x != nil {
		return x.DeletedAt
	}
	return 0
}

type DistrictSuccessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *DistrictSuccessResponse) Reset() {
	*x = DistrictSuccessResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_district_district_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DistrictSuccessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DistrictSuccessResponse) ProtoMessage() {}

func (x *DistrictSuccessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_district_district_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DistrictSuccessResponse.ProtoReflect.Descriptor instead.
func (*DistrictSuccessResponse) Descriptor() ([]byte, []int) {
	return file_district_district_proto_rawDescGZIP(), []int{5}
}

func (x *DistrictSuccessResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type UpdateDistrictRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Code             string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Name             string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	RegionId         int64  `protobuf:"varint,4,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	OldCode          int64  `protobuf:"varint,5,opt,name=old_code,json=oldCode,proto3" json:"old_code,omitempty"`
	OldName          string `protobuf:"bytes,6,opt,name=old_name,json=oldName,proto3" json:"old_name,omitempty"`
	ActivationDate   string `protobuf:"bytes,7,opt,name=activation_date,json=activationDate,proto3" json:"activation_date,omitempty"`
	DeactivationDate string `protobuf:"bytes,8,opt,name=deactivation_date,json=deactivationDate,proto3" json:"deactivation_date,omitempty"`
	FlexFinId        string `protobuf:"bytes,9,opt,name=flex_fin_id,json=flexFinId,proto3" json:"flex_fin_id,omitempty"`
}

func (x *UpdateDistrictRequest) Reset() {
	*x = UpdateDistrictRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_district_district_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDistrictRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDistrictRequest) ProtoMessage() {}

func (x *UpdateDistrictRequest) ProtoReflect() protoreflect.Message {
	mi := &file_district_district_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDistrictRequest.ProtoReflect.Descriptor instead.
func (*UpdateDistrictRequest) Descriptor() ([]byte, []int) {
	return file_district_district_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateDistrictRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateDistrictRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *UpdateDistrictRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateDistrictRequest) GetRegionId() int64 {
	if x != nil {
		return x.RegionId
	}
	return 0
}

func (x *UpdateDistrictRequest) GetOldCode() int64 {
	if x != nil {
		return x.OldCode
	}
	return 0
}

func (x *UpdateDistrictRequest) GetOldName() string {
	if x != nil {
		return x.OldName
	}
	return ""
}

func (x *UpdateDistrictRequest) GetActivationDate() string {
	if x != nil {
		return x.ActivationDate
	}
	return ""
}

func (x *UpdateDistrictRequest) GetDeactivationDate() string {
	if x != nil {
		return x.DeactivationDate
	}
	return ""
}

func (x *UpdateDistrictRequest) GetFlexFinId() string {
	if x != nil {
		return x.FlexFinId
	}
	return ""
}

var File_district_district_proto protoreflect.FileDescriptor

var file_district_district_proto_rawDesc = []byte{
	0x0a, 0x17, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x72,
	0x69, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x53, 0x0a, 0x13, 0x4c, 0x69, 0x73,
	0x74, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x4e,
	0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x09, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69,
	0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x44, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x09, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x73, 0x22, 0x88,
	0x02, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x6f, 0x6c, 0x64, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x6f, 0x6c, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x6c, 0x64, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x6c, 0x64, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x0a, 0x11,
	0x64, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x64, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x0a, 0x0b, 0x66, 0x6c, 0x65,
	0x78, 0x5f, 0x66, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x66, 0x6c, 0x65, 0x78, 0x46, 0x69, 0x6e, 0x49, 0x64, 0x22, 0x56, 0x0a, 0x15, 0x53, 0x69, 0x6e,
	0x67, 0x6c, 0x65, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x15, 0x0a, 0x13, 0x64, 0x69,
	0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x22, 0xb9, 0x02, 0x0a, 0x17, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x6c, 0x64, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x6c, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x6f, 0x6c, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6f, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x2b, 0x0a, 0x11, 0x64, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x64, 0x65,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1e,
	0x0a, 0x0b, 0x66, 0x6c, 0x65, 0x78, 0x5f, 0x66, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x6c, 0x65, 0x78, 0x46, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x35, 0x0a,
	0x17, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x98, 0x02, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44,
	0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x6c, 0x64, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x6c, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x6f, 0x6c, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6f, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61,
	0x74, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x64, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x64,
	0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12,
	0x1e, 0x0a, 0x0b, 0x66, 0x6c, 0x65, 0x78, 0x5f, 0x66, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x6c, 0x65, 0x78, 0x46, 0x69, 0x6e, 0x49, 0x64, 0x32,
	0xb3, 0x02, 0x0a, 0x0f, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x15, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x16, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x69, 0x73, 0x74, 0x72,
	0x69, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x44, 0x69, 0x73,
	0x74, 0x72, 0x69, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x16, 0x2e, 0x53, 0x69,
	0x6e, 0x67, 0x6c, 0x65, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a,
	0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x44, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x44, 0x69,
	0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x10, 0x5a, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x64,
	0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_district_district_proto_rawDescOnce sync.Once
	file_district_district_proto_rawDescData = file_district_district_proto_rawDesc
)

func file_district_district_proto_rawDescGZIP() []byte {
	file_district_district_proto_rawDescOnce.Do(func() {
		file_district_district_proto_rawDescData = protoimpl.X.CompressGZIP(file_district_district_proto_rawDescData)
	})
	return file_district_district_proto_rawDescData
}

var file_district_district_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_district_district_proto_goTypes = []interface{}{
	(*ListDistrictRequest)(nil),     // 0: ListDistrictRequest
	(*ListDistrictResponse)(nil),    // 1: ListDistrictResponse
	(*CreateDistrictRequest)(nil),   // 2: CreateDistrictRequest
	(*SingleDistrictRequest)(nil),   // 3: SingleDistrictRequest
	(*DistrictProfileResponse)(nil), // 4: DistrictProfileResponse
	(*DistrictSuccessResponse)(nil), // 5: DistrictSuccessResponse
	(*UpdateDistrictRequest)(nil),   // 6: UpdateDistrictRequest
}
var file_district_district_proto_depIdxs = []int32{
	4, // 0: ListDistrictResponse.districts:type_name -> DistrictProfileResponse
	0, // 1: DistrictService.List:input_type -> ListDistrictRequest
	2, // 2: DistrictService.Create:input_type -> CreateDistrictRequest
	3, // 3: DistrictService.Get:input_type -> SingleDistrictRequest
	6, // 4: DistrictService.Update:input_type -> UpdateDistrictRequest
	3, // 5: DistrictService.Delete:input_type -> SingleDistrictRequest
	1, // 6: DistrictService.List:output_type -> ListDistrictResponse
	4, // 7: DistrictService.Create:output_type -> DistrictProfileResponse
	4, // 8: DistrictService.Get:output_type -> DistrictProfileResponse
	5, // 9: DistrictService.Update:output_type -> DistrictSuccessResponse
	5, // 10: DistrictService.Delete:output_type -> DistrictSuccessResponse
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_district_district_proto_init() }
func file_district_district_proto_init() {
	if File_district_district_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_district_district_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDistrictRequest); i {
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
		file_district_district_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDistrictResponse); i {
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
		file_district_district_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDistrictRequest); i {
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
		file_district_district_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SingleDistrictRequest); i {
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
		file_district_district_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DistrictProfileResponse); i {
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
		file_district_district_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DistrictSuccessResponse); i {
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
		file_district_district_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDistrictRequest); i {
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
	file_district_district_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*SingleDistrictRequest_Id)(nil),
		(*SingleDistrictRequest_Code)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_district_district_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_district_district_proto_goTypes,
		DependencyIndexes: file_district_district_proto_depIdxs,
		MessageInfos:      file_district_district_proto_msgTypes,
	}.Build()
	File_district_district_proto = out.File
	file_district_district_proto_rawDesc = nil
	file_district_district_proto_goTypes = nil
	file_district_district_proto_depIdxs = nil
}
