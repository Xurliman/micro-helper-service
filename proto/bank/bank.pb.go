// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: bank.proto

package bank

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

type CreateBankRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code             int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Name             string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ShortName        string `protobuf:"bytes,3,opt,name=short_name,json=shortName,proto3" json:"short_name,omitempty"`
	CountryId        int64  `protobuf:"varint,4,opt,name=country_id,json=countryId,proto3" json:"country_id,omitempty"`
	OpenDate         string `protobuf:"bytes,5,opt,name=open_date,json=openDate,proto3" json:"open_date,omitempty"`
	CloseDate        string `protobuf:"bytes,6,opt,name=close_date,json=closeDate,proto3" json:"close_date,omitempty"`
	ActivationDate   string `protobuf:"bytes,7,opt,name=activation_date,json=activationDate,proto3" json:"activation_date,omitempty"`
	DeactivationDate string `protobuf:"bytes,8,opt,name=deactivation_date,json=deactivationDate,proto3" json:"deactivation_date,omitempty"`
}

func (x *CreateBankRequest) Reset() {
	*x = CreateBankRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBankRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBankRequest) ProtoMessage() {}

func (x *CreateBankRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bank_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBankRequest.ProtoReflect.Descriptor instead.
func (*CreateBankRequest) Descriptor() ([]byte, []int) {
	return file_bank_proto_rawDescGZIP(), []int{0}
}

func (x *CreateBankRequest) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CreateBankRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateBankRequest) GetShortName() string {
	if x != nil {
		return x.ShortName
	}
	return ""
}

func (x *CreateBankRequest) GetCountryId() int64 {
	if x != nil {
		return x.CountryId
	}
	return 0
}

func (x *CreateBankRequest) GetOpenDate() string {
	if x != nil {
		return x.OpenDate
	}
	return ""
}

func (x *CreateBankRequest) GetCloseDate() string {
	if x != nil {
		return x.CloseDate
	}
	return ""
}

func (x *CreateBankRequest) GetActivationDate() string {
	if x != nil {
		return x.ActivationDate
	}
	return ""
}

func (x *CreateBankRequest) GetDeactivationDate() string {
	if x != nil {
		return x.DeactivationDate
	}
	return ""
}

type SingleBankRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SingleBankRequest) Reset() {
	*x = SingleBankRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SingleBankRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SingleBankRequest) ProtoMessage() {}

func (x *SingleBankRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bank_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SingleBankRequest.ProtoReflect.Descriptor instead.
func (*SingleBankRequest) Descriptor() ([]byte, []int) {
	return file_bank_proto_rawDescGZIP(), []int{1}
}

func (x *SingleBankRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type BankProfileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name             string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Code             int64  `protobuf:"varint,3,opt,name=code,proto3" json:"code,omitempty"`
	ShortName        string `protobuf:"bytes,4,opt,name=short_name,json=shortName,proto3" json:"short_name,omitempty"`
	CountryId        int64  `protobuf:"varint,5,opt,name=country_id,json=countryId,proto3" json:"country_id,omitempty"`
	OpenDate         string `protobuf:"bytes,6,opt,name=open_date,json=openDate,proto3" json:"open_date,omitempty"`
	CloseDate        string `protobuf:"bytes,7,opt,name=close_date,json=closeDate,proto3" json:"close_date,omitempty"`
	ActivationDate   string `protobuf:"bytes,8,opt,name=activation_date,json=activationDate,proto3" json:"activation_date,omitempty"`
	DeactivationDate string `protobuf:"bytes,9,opt,name=deactivation_date,json=deactivationDate,proto3" json:"deactivation_date,omitempty"`
}

func (x *BankProfileResponse) Reset() {
	*x = BankProfileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BankProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BankProfileResponse) ProtoMessage() {}

func (x *BankProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bank_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BankProfileResponse.ProtoReflect.Descriptor instead.
func (*BankProfileResponse) Descriptor() ([]byte, []int) {
	return file_bank_proto_rawDescGZIP(), []int{2}
}

func (x *BankProfileResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BankProfileResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BankProfileResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *BankProfileResponse) GetShortName() string {
	if x != nil {
		return x.ShortName
	}
	return ""
}

func (x *BankProfileResponse) GetCountryId() int64 {
	if x != nil {
		return x.CountryId
	}
	return 0
}

func (x *BankProfileResponse) GetOpenDate() string {
	if x != nil {
		return x.OpenDate
	}
	return ""
}

func (x *BankProfileResponse) GetCloseDate() string {
	if x != nil {
		return x.CloseDate
	}
	return ""
}

func (x *BankProfileResponse) GetActivationDate() string {
	if x != nil {
		return x.ActivationDate
	}
	return ""
}

func (x *BankProfileResponse) GetDeactivationDate() string {
	if x != nil {
		return x.DeactivationDate
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
		mi := &file_bank_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuccessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuccessResponse) ProtoMessage() {}

func (x *SuccessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bank_proto_msgTypes[3]
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
	return file_bank_proto_rawDescGZIP(), []int{3}
}

func (x *SuccessResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type UpdateBankRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name             string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Code             int64  `protobuf:"varint,3,opt,name=code,proto3" json:"code,omitempty"`
	ShortName        string `protobuf:"bytes,4,opt,name=short_name,json=shortName,proto3" json:"short_name,omitempty"`
	CountryId        int64  `protobuf:"varint,5,opt,name=country_id,json=countryId,proto3" json:"country_id,omitempty"`
	OpenDate         string `protobuf:"bytes,6,opt,name=open_date,json=openDate,proto3" json:"open_date,omitempty"`
	CloseDate        string `protobuf:"bytes,7,opt,name=close_date,json=closeDate,proto3" json:"close_date,omitempty"`
	ActivationDate   string `protobuf:"bytes,8,opt,name=activation_date,json=activationDate,proto3" json:"activation_date,omitempty"`
	DeactivationDate string `protobuf:"bytes,9,opt,name=deactivation_date,json=deactivationDate,proto3" json:"deactivation_date,omitempty"`
}

func (x *UpdateBankRequest) Reset() {
	*x = UpdateBankRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bank_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBankRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBankRequest) ProtoMessage() {}

func (x *UpdateBankRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bank_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBankRequest.ProtoReflect.Descriptor instead.
func (*UpdateBankRequest) Descriptor() ([]byte, []int) {
	return file_bank_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateBankRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateBankRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateBankRequest) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *UpdateBankRequest) GetShortName() string {
	if x != nil {
		return x.ShortName
	}
	return ""
}

func (x *UpdateBankRequest) GetCountryId() int64 {
	if x != nil {
		return x.CountryId
	}
	return 0
}

func (x *UpdateBankRequest) GetOpenDate() string {
	if x != nil {
		return x.OpenDate
	}
	return ""
}

func (x *UpdateBankRequest) GetCloseDate() string {
	if x != nil {
		return x.CloseDate
	}
	return ""
}

func (x *UpdateBankRequest) GetActivationDate() string {
	if x != nil {
		return x.ActivationDate
	}
	return ""
}

func (x *UpdateBankRequest) GetDeactivationDate() string {
	if x != nil {
		return x.DeactivationDate
	}
	return ""
}

var File_bank_proto protoreflect.FileDescriptor

var file_bank_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x02, 0x0a,
	0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x6e,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x70, 0x65,
	0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x6f, 0x73, 0x65,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x0a,
	0x11, 0x64, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x64, 0x65, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x22, 0x23, 0x0a, 0x11, 0x53, 0x69,
	0x6e, 0x67, 0x6c, 0x65, 0x42, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x9d, 0x02, 0x0a, 0x13, 0x42, 0x61, 0x6e, 0x6b, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6f, 0x70, 0x65, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c,
	0x6f, 0x73, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x6c, 0x6f, 0x73, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61,
	0x74, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x64, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x64,
	0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x22,
	0x2d, 0x0a, 0x0f, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x9b,
	0x02, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x70,
	0x65, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f,
	0x70, 0x65, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x6f, 0x73, 0x65,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x6f,
	0x73, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12,
	0x2b, 0x0a, 0x11, 0x64, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x64, 0x65, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x32, 0xd3, 0x01, 0x0a,
	0x0b, 0x42, 0x61, 0x6e, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x06,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x61, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x42, 0x61, 0x6e,
	0x6b, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x30, 0x0a, 0x04, 0x52, 0x65, 0x61, 0x64, 0x12, 0x12, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c,
	0x65, 0x42, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x42,
	0x61, 0x6e, 0x6b, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x10, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x53,
	0x69, 0x6e, 0x67, 0x6c, 0x65, 0x42, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x10, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x61,
	0x6e, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bank_proto_rawDescOnce sync.Once
	file_bank_proto_rawDescData = file_bank_proto_rawDesc
)

func file_bank_proto_rawDescGZIP() []byte {
	file_bank_proto_rawDescOnce.Do(func() {
		file_bank_proto_rawDescData = protoimpl.X.CompressGZIP(file_bank_proto_rawDescData)
	})
	return file_bank_proto_rawDescData
}

var file_bank_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_bank_proto_goTypes = []interface{}{
	(*CreateBankRequest)(nil),   // 0: CreateBankRequest
	(*SingleBankRequest)(nil),   // 1: SingleBankRequest
	(*BankProfileResponse)(nil), // 2: BankProfileResponse
	(*SuccessResponse)(nil),     // 3: SuccessResponse
	(*UpdateBankRequest)(nil),   // 4: UpdateBankRequest
}
var file_bank_proto_depIdxs = []int32{
	0, // 0: BankService.Create:input_type -> CreateBankRequest
	1, // 1: BankService.Read:input_type -> SingleBankRequest
	4, // 2: BankService.Update:input_type -> UpdateBankRequest
	1, // 3: BankService.Delete:input_type -> SingleBankRequest
	2, // 4: BankService.Create:output_type -> BankProfileResponse
	2, // 5: BankService.Read:output_type -> BankProfileResponse
	3, // 6: BankService.Update:output_type -> SuccessResponse
	3, // 7: BankService.Delete:output_type -> SuccessResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bank_proto_init() }
func file_bank_proto_init() {
	if File_bank_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bank_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBankRequest); i {
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
		file_bank_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SingleBankRequest); i {
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
		file_bank_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BankProfileResponse); i {
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
		file_bank_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_bank_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBankRequest); i {
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
			RawDescriptor: file_bank_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bank_proto_goTypes,
		DependencyIndexes: file_bank_proto_depIdxs,
		MessageInfos:      file_bank_proto_msgTypes,
	}.Build()
	File_bank_proto = out.File
	file_bank_proto_rawDesc = nil
	file_bank_proto_goTypes = nil
	file_bank_proto_depIdxs = nil
}