// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.12
// source: client_type_classifier.proto

package proto_client_type_classifier

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

type ListClientTypeClassifierRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageNumber int32 `protobuf:"varint,1,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	PageSize   int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *ListClientTypeClassifierRequest) Reset() {
	*x = ListClientTypeClassifierRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_type_classifier_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListClientTypeClassifierRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListClientTypeClassifierRequest) ProtoMessage() {}

func (x *ListClientTypeClassifierRequest) ProtoReflect() protoreflect.Message {
	mi := &file_client_type_classifier_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListClientTypeClassifierRequest.ProtoReflect.Descriptor instead.
func (*ListClientTypeClassifierRequest) Descriptor() ([]byte, []int) {
	return file_client_type_classifier_proto_rawDescGZIP(), []int{0}
}

func (x *ListClientTypeClassifierRequest) GetPageNumber() int32 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *ListClientTypeClassifierRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListClientTypeClassifierResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientTypeClassifier []*ClientTypeClassifierProfileResponse `protobuf:"bytes,1,rep,name=client_type_classifier,json=clientTypeClassifier,proto3" json:"client_type_classifier,omitempty"`
}

func (x *ListClientTypeClassifierResponse) Reset() {
	*x = ListClientTypeClassifierResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_type_classifier_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListClientTypeClassifierResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListClientTypeClassifierResponse) ProtoMessage() {}

func (x *ListClientTypeClassifierResponse) ProtoReflect() protoreflect.Message {
	mi := &file_client_type_classifier_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListClientTypeClassifierResponse.ProtoReflect.Descriptor instead.
func (*ListClientTypeClassifierResponse) Descriptor() ([]byte, []int) {
	return file_client_type_classifier_proto_rawDescGZIP(), []int{1}
}

func (x *ListClientTypeClassifierResponse) GetClientTypeClassifier() []*ClientTypeClassifierProfileResponse {
	if x != nil {
		return x.ClientTypeClassifier
	}
	return nil
}

type CreateClientTypeClassifierRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code             string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Name             string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ShortName        string `protobuf:"bytes,3,opt,name=short_name,json=shortName,proto3" json:"short_name,omitempty"`
	ClientType       int32  `protobuf:"varint,4,opt,name=client_type,json=clientType,proto3" json:"client_type,omitempty"`
	ActivationDate   string `protobuf:"bytes,5,opt,name=activation_date,json=activationDate,proto3" json:"activation_date,omitempty"`
	DeactivationDate string `protobuf:"bytes,6,opt,name=deactivation_date,json=deactivationDate,proto3" json:"deactivation_date,omitempty"`
	CbuReferenceKey  int32  `protobuf:"varint,7,opt,name=cbu_reference_key,json=cbuReferenceKey,proto3" json:"cbu_reference_key,omitempty"`
	OldCode          int32  `protobuf:"varint,8,opt,name=old_code,json=oldCode,proto3" json:"old_code,omitempty"`
	OldName          string `protobuf:"bytes,9,opt,name=old_name,json=oldName,proto3" json:"old_name,omitempty"`
}

func (x *CreateClientTypeClassifierRequest) Reset() {
	*x = CreateClientTypeClassifierRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_type_classifier_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateClientTypeClassifierRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateClientTypeClassifierRequest) ProtoMessage() {}

func (x *CreateClientTypeClassifierRequest) ProtoReflect() protoreflect.Message {
	mi := &file_client_type_classifier_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateClientTypeClassifierRequest.ProtoReflect.Descriptor instead.
func (*CreateClientTypeClassifierRequest) Descriptor() ([]byte, []int) {
	return file_client_type_classifier_proto_rawDescGZIP(), []int{2}
}

func (x *CreateClientTypeClassifierRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CreateClientTypeClassifierRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateClientTypeClassifierRequest) GetShortName() string {
	if x != nil {
		return x.ShortName
	}
	return ""
}

func (x *CreateClientTypeClassifierRequest) GetClientType() int32 {
	if x != nil {
		return x.ClientType
	}
	return 0
}

func (x *CreateClientTypeClassifierRequest) GetActivationDate() string {
	if x != nil {
		return x.ActivationDate
	}
	return ""
}

func (x *CreateClientTypeClassifierRequest) GetDeactivationDate() string {
	if x != nil {
		return x.DeactivationDate
	}
	return ""
}

func (x *CreateClientTypeClassifierRequest) GetCbuReferenceKey() int32 {
	if x != nil {
		return x.CbuReferenceKey
	}
	return 0
}

func (x *CreateClientTypeClassifierRequest) GetOldCode() int32 {
	if x != nil {
		return x.OldCode
	}
	return 0
}

func (x *CreateClientTypeClassifierRequest) GetOldName() string {
	if x != nil {
		return x.OldName
	}
	return ""
}

type SingleClientTypeClassifierRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to ClientTypeClassifierIdentifier:
	//
	//	*SingleClientTypeClassifierRequest_Id
	//	*SingleClientTypeClassifierRequest_Code
	ClientTypeClassifierIdentifier isSingleClientTypeClassifierRequest_ClientTypeClassifierIdentifier `protobuf_oneof:"client_type_classifier_identifier"`
}

func (x *SingleClientTypeClassifierRequest) Reset() {
	*x = SingleClientTypeClassifierRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_type_classifier_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SingleClientTypeClassifierRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SingleClientTypeClassifierRequest) ProtoMessage() {}

func (x *SingleClientTypeClassifierRequest) ProtoReflect() protoreflect.Message {
	mi := &file_client_type_classifier_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SingleClientTypeClassifierRequest.ProtoReflect.Descriptor instead.
func (*SingleClientTypeClassifierRequest) Descriptor() ([]byte, []int) {
	return file_client_type_classifier_proto_rawDescGZIP(), []int{3}
}

func (m *SingleClientTypeClassifierRequest) GetClientTypeClassifierIdentifier() isSingleClientTypeClassifierRequest_ClientTypeClassifierIdentifier {
	if m != nil {
		return m.ClientTypeClassifierIdentifier
	}
	return nil
}

func (x *SingleClientTypeClassifierRequest) GetId() int64 {
	if x, ok := x.GetClientTypeClassifierIdentifier().(*SingleClientTypeClassifierRequest_Id); ok {
		return x.Id
	}
	return 0
}

func (x *SingleClientTypeClassifierRequest) GetCode() string {
	if x, ok := x.GetClientTypeClassifierIdentifier().(*SingleClientTypeClassifierRequest_Code); ok {
		return x.Code
	}
	return ""
}

type isSingleClientTypeClassifierRequest_ClientTypeClassifierIdentifier interface {
	isSingleClientTypeClassifierRequest_ClientTypeClassifierIdentifier()
}

type SingleClientTypeClassifierRequest_Id struct {
	Id int64 `protobuf:"varint,1,opt,name=id,proto3,oneof"`
}

type SingleClientTypeClassifierRequest_Code struct {
	Code string `protobuf:"bytes,2,opt,name=code,proto3,oneof"`
}

func (*SingleClientTypeClassifierRequest_Id) isSingleClientTypeClassifierRequest_ClientTypeClassifierIdentifier() {
}

func (*SingleClientTypeClassifierRequest_Code) isSingleClientTypeClassifierRequest_ClientTypeClassifierIdentifier() {
}

type ClientTypeClassifierProfileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Code             string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Name             string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	ShortName        string `protobuf:"bytes,4,opt,name=short_name,json=shortName,proto3" json:"short_name,omitempty"`
	ClientType       int32  `protobuf:"varint,5,opt,name=client_type,json=clientType,proto3" json:"client_type,omitempty"`
	ActivationDate   string `protobuf:"bytes,6,opt,name=activation_date,json=activationDate,proto3" json:"activation_date,omitempty"`
	DeactivationDate string `protobuf:"bytes,7,opt,name=deactivation_date,json=deactivationDate,proto3" json:"deactivation_date,omitempty"`
	CbuReferenceKey  int32  `protobuf:"varint,8,opt,name=cbu_reference_key,json=cbuReferenceKey,proto3" json:"cbu_reference_key,omitempty"`
	OldCode          int32  `protobuf:"varint,9,opt,name=old_code,json=oldCode,proto3" json:"old_code,omitempty"`
	OldName          string `protobuf:"bytes,10,opt,name=old_name,json=oldName,proto3" json:"old_name,omitempty"`
	DeletedAt        string `protobuf:"bytes,11,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *ClientTypeClassifierProfileResponse) Reset() {
	*x = ClientTypeClassifierProfileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_type_classifier_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientTypeClassifierProfileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientTypeClassifierProfileResponse) ProtoMessage() {}

func (x *ClientTypeClassifierProfileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_client_type_classifier_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientTypeClassifierProfileResponse.ProtoReflect.Descriptor instead.
func (*ClientTypeClassifierProfileResponse) Descriptor() ([]byte, []int) {
	return file_client_type_classifier_proto_rawDescGZIP(), []int{4}
}

func (x *ClientTypeClassifierProfileResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ClientTypeClassifierProfileResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *ClientTypeClassifierProfileResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ClientTypeClassifierProfileResponse) GetShortName() string {
	if x != nil {
		return x.ShortName
	}
	return ""
}

func (x *ClientTypeClassifierProfileResponse) GetClientType() int32 {
	if x != nil {
		return x.ClientType
	}
	return 0
}

func (x *ClientTypeClassifierProfileResponse) GetActivationDate() string {
	if x != nil {
		return x.ActivationDate
	}
	return ""
}

func (x *ClientTypeClassifierProfileResponse) GetDeactivationDate() string {
	if x != nil {
		return x.DeactivationDate
	}
	return ""
}

func (x *ClientTypeClassifierProfileResponse) GetCbuReferenceKey() int32 {
	if x != nil {
		return x.CbuReferenceKey
	}
	return 0
}

func (x *ClientTypeClassifierProfileResponse) GetOldCode() int32 {
	if x != nil {
		return x.OldCode
	}
	return 0
}

func (x *ClientTypeClassifierProfileResponse) GetOldName() string {
	if x != nil {
		return x.OldName
	}
	return ""
}

func (x *ClientTypeClassifierProfileResponse) GetDeletedAt() string {
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
		mi := &file_client_type_classifier_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuccessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuccessResponse) ProtoMessage() {}

func (x *SuccessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_client_type_classifier_proto_msgTypes[5]
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
	return file_client_type_classifier_proto_rawDescGZIP(), []int{5}
}

func (x *SuccessResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type UpdateClientTypeClassifierRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Code             string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Name             string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	ShortName        string `protobuf:"bytes,4,opt,name=short_name,json=shortName,proto3" json:"short_name,omitempty"`
	ClientType       int32  `protobuf:"varint,5,opt,name=client_type,json=clientType,proto3" json:"client_type,omitempty"`
	ActivationDate   string `protobuf:"bytes,6,opt,name=activation_date,json=activationDate,proto3" json:"activation_date,omitempty"`
	DeactivationDate string `protobuf:"bytes,7,opt,name=deactivation_date,json=deactivationDate,proto3" json:"deactivation_date,omitempty"`
	CbuReferenceKey  int32  `protobuf:"varint,8,opt,name=cbu_reference_key,json=cbuReferenceKey,proto3" json:"cbu_reference_key,omitempty"`
	OldCode          int32  `protobuf:"varint,9,opt,name=old_code,json=oldCode,proto3" json:"old_code,omitempty"`
	OldName          string `protobuf:"bytes,10,opt,name=old_name,json=oldName,proto3" json:"old_name,omitempty"`
}

func (x *UpdateClientTypeClassifierRequest) Reset() {
	*x = UpdateClientTypeClassifierRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_type_classifier_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateClientTypeClassifierRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateClientTypeClassifierRequest) ProtoMessage() {}

func (x *UpdateClientTypeClassifierRequest) ProtoReflect() protoreflect.Message {
	mi := &file_client_type_classifier_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateClientTypeClassifierRequest.ProtoReflect.Descriptor instead.
func (*UpdateClientTypeClassifierRequest) Descriptor() ([]byte, []int) {
	return file_client_type_classifier_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateClientTypeClassifierRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateClientTypeClassifierRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *UpdateClientTypeClassifierRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateClientTypeClassifierRequest) GetShortName() string {
	if x != nil {
		return x.ShortName
	}
	return ""
}

func (x *UpdateClientTypeClassifierRequest) GetClientType() int32 {
	if x != nil {
		return x.ClientType
	}
	return 0
}

func (x *UpdateClientTypeClassifierRequest) GetActivationDate() string {
	if x != nil {
		return x.ActivationDate
	}
	return ""
}

func (x *UpdateClientTypeClassifierRequest) GetDeactivationDate() string {
	if x != nil {
		return x.DeactivationDate
	}
	return ""
}

func (x *UpdateClientTypeClassifierRequest) GetCbuReferenceKey() int32 {
	if x != nil {
		return x.CbuReferenceKey
	}
	return 0
}

func (x *UpdateClientTypeClassifierRequest) GetOldCode() int32 {
	if x != nil {
		return x.OldCode
	}
	return 0
}

func (x *UpdateClientTypeClassifierRequest) GetOldName() string {
	if x != nil {
		return x.OldName
	}
	return ""
}

var File_client_type_classifier_proto protoreflect.FileDescriptor

var file_client_type_classifier_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5f,
	0x0a, 0x1f, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22,
	0x7e, 0x0a, 0x20, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x16, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x14, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x65, 0x72, 0x22,
	0xc3, 0x02, 0x0a, 0x21, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x27, 0x0a,
	0x0f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x64, 0x65, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x64, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x63, 0x62, 0x75, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f,
	0x63, 0x62, 0x75, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x4b, 0x65, 0x79, 0x12,
	0x19, 0x0a, 0x08, 0x6f, 0x6c, 0x64, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x6f, 0x6c, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x6c,
	0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x6c,
	0x64, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x70, 0x0a, 0x21, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x42, 0x23, 0x0a, 0x21, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x22, 0xf4, 0x02, 0x0a, 0x23, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65,
	0x12, 0x2b, 0x0a, 0x11, 0x64, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x64, 0x65, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x2a, 0x0a,
	0x11, 0x63, 0x62, 0x75, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x63, 0x62, 0x75, 0x52, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x6c, 0x64,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6f, 0x6c, 0x64,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x6c, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x2d,
	0x0a, 0x0f, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xd3, 0x02,
	0x0a, 0x21, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x64, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x64, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x2a, 0x0a, 0x11, 0x63, 0x62, 0x75, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x63, 0x62,
	0x75, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x19, 0x0a,
	0x08, 0x6f, 0x6c, 0x64, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x6f, 0x6c, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x6c, 0x64, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x6c, 0x64, 0x4e,
	0x61, 0x6d, 0x65, 0x32, 0x90, 0x03, 0x0a, 0x1b, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6c, 0x61, 0x73,
	0x73, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x52, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x22, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24,
	0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x04, 0x52, 0x65, 0x61, 0x64, 0x12, 0x22, 0x2e, 0x53,
	0x69, 0x6e, 0x67, 0x6c, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x24, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x69, 0x66, 0x69, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x22, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x22, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1e, 0x5a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x69, 0x66, 0x69, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_client_type_classifier_proto_rawDescOnce sync.Once
	file_client_type_classifier_proto_rawDescData = file_client_type_classifier_proto_rawDesc
)

func file_client_type_classifier_proto_rawDescGZIP() []byte {
	file_client_type_classifier_proto_rawDescOnce.Do(func() {
		file_client_type_classifier_proto_rawDescData = protoimpl.X.CompressGZIP(file_client_type_classifier_proto_rawDescData)
	})
	return file_client_type_classifier_proto_rawDescData
}

var file_client_type_classifier_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_client_type_classifier_proto_goTypes = []interface{}{
	(*ListClientTypeClassifierRequest)(nil),     // 0: ListClientTypeClassifierRequest
	(*ListClientTypeClassifierResponse)(nil),    // 1: ListClientTypeClassifierResponse
	(*CreateClientTypeClassifierRequest)(nil),   // 2: CreateClientTypeClassifierRequest
	(*SingleClientTypeClassifierRequest)(nil),   // 3: SingleClientTypeClassifierRequest
	(*ClientTypeClassifierProfileResponse)(nil), // 4: ClientTypeClassifierProfileResponse
	(*SuccessResponse)(nil),                     // 5: SuccessResponse
	(*UpdateClientTypeClassifierRequest)(nil),   // 6: UpdateClientTypeClassifierRequest
}
var file_client_type_classifier_proto_depIdxs = []int32{
	4, // 0: ListClientTypeClassifierResponse.client_type_classifier:type_name -> ClientTypeClassifierProfileResponse
	0, // 1: ClientTypeClassifierService.List:input_type -> ListClientTypeClassifierRequest
	2, // 2: ClientTypeClassifierService.Create:input_type -> CreateClientTypeClassifierRequest
	3, // 3: ClientTypeClassifierService.Read:input_type -> SingleClientTypeClassifierRequest
	6, // 4: ClientTypeClassifierService.Update:input_type -> UpdateClientTypeClassifierRequest
	3, // 5: ClientTypeClassifierService.Delete:input_type -> SingleClientTypeClassifierRequest
	1, // 6: ClientTypeClassifierService.List:output_type -> ListClientTypeClassifierResponse
	4, // 7: ClientTypeClassifierService.Create:output_type -> ClientTypeClassifierProfileResponse
	4, // 8: ClientTypeClassifierService.Read:output_type -> ClientTypeClassifierProfileResponse
	5, // 9: ClientTypeClassifierService.Update:output_type -> SuccessResponse
	5, // 10: ClientTypeClassifierService.Delete:output_type -> SuccessResponse
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_client_type_classifier_proto_init() }
func file_client_type_classifier_proto_init() {
	if File_client_type_classifier_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_client_type_classifier_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListClientTypeClassifierRequest); i {
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
		file_client_type_classifier_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListClientTypeClassifierResponse); i {
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
		file_client_type_classifier_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateClientTypeClassifierRequest); i {
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
		file_client_type_classifier_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SingleClientTypeClassifierRequest); i {
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
		file_client_type_classifier_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientTypeClassifierProfileResponse); i {
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
		file_client_type_classifier_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_client_type_classifier_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateClientTypeClassifierRequest); i {
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
	file_client_type_classifier_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*SingleClientTypeClassifierRequest_Id)(nil),
		(*SingleClientTypeClassifierRequest_Code)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_client_type_classifier_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_client_type_classifier_proto_goTypes,
		DependencyIndexes: file_client_type_classifier_proto_depIdxs,
		MessageInfos:      file_client_type_classifier_proto_msgTypes,
	}.Build()
	File_client_type_classifier_proto = out.File
	file_client_type_classifier_proto_rawDesc = nil
	file_client_type_classifier_proto_goTypes = nil
	file_client_type_classifier_proto_depIdxs = nil
}
