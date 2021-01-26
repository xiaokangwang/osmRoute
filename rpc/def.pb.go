// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: def.proto

package rpc

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ObjectResolveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FeatureID       string `protobuf:"bytes,1,opt,name=FeatureID,proto3" json:"FeatureID,omitempty"`
	SkipIfNotCached bool   `protobuf:"varint,2,opt,name=SkipIfNotCached,proto3" json:"SkipIfNotCached,omitempty"`
}

func (x *ObjectResolveRequest) Reset() {
	*x = ObjectResolveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_def_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectResolveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectResolveRequest) ProtoMessage() {}

func (x *ObjectResolveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_def_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectResolveRequest.ProtoReflect.Descriptor instead.
func (*ObjectResolveRequest) Descriptor() ([]byte, []int) {
	return file_def_proto_rawDescGZIP(), []int{0}
}

func (x *ObjectResolveRequest) GetFeatureID() string {
	if x != nil {
		return x.FeatureID
	}
	return ""
}

func (x *ObjectResolveRequest) GetSkipIfNotCached() bool {
	if x != nil {
		return x.SkipIfNotCached
	}
	return false
}

type ReturnedObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FeatureID     string `protobuf:"bytes,1,opt,name=FeatureID,proto3" json:"FeatureID,omitempty"`
	ObjectContent []byte `protobuf:"bytes,2,opt,name=ObjectContent,proto3" json:"ObjectContent,omitempty"`
	Found         bool   `protobuf:"varint,3,opt,name=Found,proto3" json:"Found,omitempty"`
}

func (x *ReturnedObject) Reset() {
	*x = ReturnedObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_def_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReturnedObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReturnedObject) ProtoMessage() {}

func (x *ReturnedObject) ProtoReflect() protoreflect.Message {
	mi := &file_def_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReturnedObject.ProtoReflect.Descriptor instead.
func (*ReturnedObject) Descriptor() ([]byte, []int) {
	return file_def_proto_rawDescGZIP(), []int{1}
}

func (x *ReturnedObject) GetFeatureID() string {
	if x != nil {
		return x.FeatureID
	}
	return ""
}

func (x *ReturnedObject) GetObjectContent() []byte {
	if x != nil {
		return x.ObjectContent
	}
	return nil
}

func (x *ReturnedObject) GetFound() bool {
	if x != nil {
		return x.Found
	}
	return false
}

type ScanRegionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lat float64 `protobuf:"fixed64,1,opt,name=Lat,proto3" json:"Lat,omitempty"`
	Lon float64 `protobuf:"fixed64,2,opt,name=Lon,proto3" json:"Lon,omitempty"`
}

func (x *ScanRegionRequest) Reset() {
	*x = ScanRegionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_def_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScanRegionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScanRegionRequest) ProtoMessage() {}

func (x *ScanRegionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_def_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScanRegionRequest.ProtoReflect.Descriptor instead.
func (*ScanRegionRequest) Descriptor() ([]byte, []int) {
	return file_def_proto_rawDescGZIP(), []int{2}
}

func (x *ScanRegionRequest) GetLat() float64 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *ScanRegionRequest) GetLon() float64 {
	if x != nil {
		return x.Lon
	}
	return 0
}

type GetAssociatedObjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FeatureID string `protobuf:"bytes,1,opt,name=FeatureID,proto3" json:"FeatureID,omitempty"`
}

func (x *GetAssociatedObjectRequest) Reset() {
	*x = GetAssociatedObjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_def_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAssociatedObjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAssociatedObjectRequest) ProtoMessage() {}

func (x *GetAssociatedObjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_def_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAssociatedObjectRequest.ProtoReflect.Descriptor instead.
func (*GetAssociatedObjectRequest) Descriptor() ([]byte, []int) {
	return file_def_proto_rawDescGZIP(), []int{3}
}

func (x *GetAssociatedObjectRequest) GetFeatureID() string {
	if x != nil {
		return x.FeatureID
	}
	return ""
}

type ObjectList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FeatureID []string `protobuf:"bytes,1,rep,name=FeatureID,proto3" json:"FeatureID,omitempty"`
}

func (x *ObjectList) Reset() {
	*x = ObjectList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_def_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectList) ProtoMessage() {}

func (x *ObjectList) ProtoReflect() protoreflect.Message {
	mi := &file_def_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectList.ProtoReflect.Descriptor instead.
func (*ObjectList) Descriptor() ([]byte, []int) {
	return file_def_proto_rawDescGZIP(), []int{4}
}

func (x *ObjectList) GetFeatureID() []string {
	if x != nil {
		return x.FeatureID
	}
	return nil
}

type ObjectListWithAssociatedObjects struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FeatureID                     []string               `protobuf:"bytes,1,rep,name=FeatureID,proto3" json:"FeatureID,omitempty"`
	FeatureIDAndAssociatedObjects map[string]*ObjectList `protobuf:"bytes,2,rep,name=FeatureIDAndAssociatedObjects,proto3" json:"FeatureIDAndAssociatedObjects,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ObjectListWithAssociatedObjects) Reset() {
	*x = ObjectListWithAssociatedObjects{}
	if protoimpl.UnsafeEnabled {
		mi := &file_def_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectListWithAssociatedObjects) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectListWithAssociatedObjects) ProtoMessage() {}

func (x *ObjectListWithAssociatedObjects) ProtoReflect() protoreflect.Message {
	mi := &file_def_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectListWithAssociatedObjects.ProtoReflect.Descriptor instead.
func (*ObjectListWithAssociatedObjects) Descriptor() ([]byte, []int) {
	return file_def_proto_rawDescGZIP(), []int{5}
}

func (x *ObjectListWithAssociatedObjects) GetFeatureID() []string {
	if x != nil {
		return x.FeatureID
	}
	return nil
}

func (x *ObjectListWithAssociatedObjects) GetFeatureIDAndAssociatedObjects() map[string]*ObjectList {
	if x != nil {
		return x.FeatureIDAndAssociatedObjects
	}
	return nil
}

type NameList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectName []string `protobuf:"bytes,1,rep,name=ObjectName,proto3" json:"ObjectName,omitempty"`
}

func (x *NameList) Reset() {
	*x = NameList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_def_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NameList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NameList) ProtoMessage() {}

func (x *NameList) ProtoReflect() protoreflect.Message {
	mi := &file_def_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NameList.ProtoReflect.Descriptor instead.
func (*NameList) Descriptor() ([]byte, []int) {
	return file_def_proto_rawDescGZIP(), []int{6}
}

func (x *NameList) GetObjectName() []string {
	if x != nil {
		return x.ObjectName
	}
	return nil
}

type NameSearch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keyword string `protobuf:"bytes,1,opt,name=Keyword,proto3" json:"Keyword,omitempty"`
}

func (x *NameSearch) Reset() {
	*x = NameSearch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_def_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NameSearch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NameSearch) ProtoMessage() {}

func (x *NameSearch) ProtoReflect() protoreflect.Message {
	mi := &file_def_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NameSearch.ProtoReflect.Descriptor instead.
func (*NameSearch) Descriptor() ([]byte, []int) {
	return file_def_proto_rawDescGZIP(), []int{7}
}

func (x *NameSearch) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

var File_def_proto protoreflect.FileDescriptor

var file_def_proto_rawDesc = []byte{
	0x0a, 0x09, 0x64, 0x65, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72, 0x70, 0x63,
	0x22, 0x5e, 0x0a, 0x14, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x46, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x46, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x49, 0x44, 0x12, 0x28, 0x0a, 0x0f, 0x53, 0x6b, 0x69, 0x70, 0x49, 0x66,
	0x4e, 0x6f, 0x74, 0x43, 0x61, 0x63, 0x68, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0f, 0x53, 0x6b, 0x69, 0x70, 0x49, 0x66, 0x4e, 0x6f, 0x74, 0x43, 0x61, 0x63, 0x68, 0x65, 0x64,
	0x22, 0x6a, 0x0a, 0x0e, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x49, 0x44,
	0x12, 0x24, 0x0a, 0x0d, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x22, 0x37, 0x0a, 0x11,
	0x53, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x4c, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03,
	0x4c, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x4c, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x03, 0x4c, 0x6f, 0x6e, 0x22, 0x3a, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x6f,
	0x63, 0x69, 0x61, 0x74, 0x65, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x49,
	0x44, 0x22, 0x2a, 0x0a, 0x0a, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x09, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x49, 0x44, 0x22, 0xb2, 0x02,
	0x0a, 0x1f, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x69, 0x74, 0x68,
	0x41, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x65, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x12, 0x1c, 0x0a, 0x09, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x49, 0x44, 0x12,
	0x8d, 0x01, 0x0a, 0x1d, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x49, 0x44, 0x41, 0x6e, 0x64,
	0x41, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x65, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x47, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x69, 0x74, 0x68, 0x41, 0x73, 0x73, 0x6f,
	0x63, 0x69, 0x61, 0x74, 0x65, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x46, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x49, 0x44, 0x41, 0x6e, 0x64, 0x41, 0x73, 0x73, 0x6f, 0x63, 0x69,
	0x61, 0x74, 0x65, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x1d, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x49, 0x44, 0x41, 0x6e, 0x64, 0x41, 0x73,
	0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x65, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x1a,
	0x61, 0x0a, 0x22, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x49, 0x44, 0x41, 0x6e, 0x64, 0x41,
	0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x65, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x25, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x2a, 0x0a, 0x08, 0x4e, 0x61, 0x6d, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1e,
	0x0a, 0x0a, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0a, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x26,
	0x0a, 0x0a, 0x4e, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x18, 0x0a, 0x07,
	0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4b,
	0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x32, 0xd5, 0x02, 0x0a, 0x0c, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x07, 0x52, 0x65, 0x73, 0x6f, 0x6c,
	0x76, 0x65, 0x12, 0x19, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x0a, 0x53, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x69, 0x74, 0x68, 0x41,
	0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x65, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x22, 0x00, 0x12, 0x49, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61,
	0x74, 0x65, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1f, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x65, 0x64, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x36, 0x0a,
	0x12, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x50, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x12, 0x0f, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x1a, 0x0d, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x11, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x42,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x45, 0x78, 0x61, 0x63, 0x74, 0x12, 0x0f, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x1a, 0x0f, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x42, 0x26,
	0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x69, 0x61,
	0x6f, 0x6b, 0x61, 0x6e, 0x67, 0x77, 0x61, 0x6e, 0x67, 0x2f, 0x6f, 0x73, 0x6d, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_def_proto_rawDescOnce sync.Once
	file_def_proto_rawDescData = file_def_proto_rawDesc
)

func file_def_proto_rawDescGZIP() []byte {
	file_def_proto_rawDescOnce.Do(func() {
		file_def_proto_rawDescData = protoimpl.X.CompressGZIP(file_def_proto_rawDescData)
	})
	return file_def_proto_rawDescData
}

var file_def_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_def_proto_goTypes = []interface{}{
	(*ObjectResolveRequest)(nil),            // 0: rpc.ObjectResolveRequest
	(*ReturnedObject)(nil),                  // 1: rpc.ReturnedObject
	(*ScanRegionRequest)(nil),               // 2: rpc.ScanRegionRequest
	(*GetAssociatedObjectRequest)(nil),      // 3: rpc.GetAssociatedObjectRequest
	(*ObjectList)(nil),                      // 4: rpc.ObjectList
	(*ObjectListWithAssociatedObjects)(nil), // 5: rpc.ObjectListWithAssociatedObjects
	(*NameList)(nil),                        // 6: rpc.NameList
	(*NameSearch)(nil),                      // 7: rpc.NameSearch
	nil,                                     // 8: rpc.ObjectListWithAssociatedObjects.FeatureIDAndAssociatedObjectsEntry
}
var file_def_proto_depIdxs = []int32{
	8, // 0: rpc.ObjectListWithAssociatedObjects.FeatureIDAndAssociatedObjects:type_name -> rpc.ObjectListWithAssociatedObjects.FeatureIDAndAssociatedObjectsEntry
	4, // 1: rpc.ObjectListWithAssociatedObjects.FeatureIDAndAssociatedObjectsEntry.value:type_name -> rpc.ObjectList
	0, // 2: rpc.RouteService.Resolve:input_type -> rpc.ObjectResolveRequest
	2, // 3: rpc.RouteService.ScanRegion:input_type -> rpc.ScanRegionRequest
	3, // 4: rpc.RouteService.GetAssociatedObject:input_type -> rpc.GetAssociatedObjectRequest
	7, // 5: rpc.RouteService.SearchByNamePrefix:input_type -> rpc.NameSearch
	7, // 6: rpc.RouteService.SearchByNameExact:input_type -> rpc.NameSearch
	1, // 7: rpc.RouteService.Resolve:output_type -> rpc.ReturnedObject
	5, // 8: rpc.RouteService.ScanRegion:output_type -> rpc.ObjectListWithAssociatedObjects
	4, // 9: rpc.RouteService.GetAssociatedObject:output_type -> rpc.ObjectList
	6, // 10: rpc.RouteService.SearchByNamePrefix:output_type -> rpc.NameList
	4, // 11: rpc.RouteService.SearchByNameExact:output_type -> rpc.ObjectList
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_def_proto_init() }
func file_def_proto_init() {
	if File_def_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_def_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectResolveRequest); i {
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
		file_def_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReturnedObject); i {
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
		file_def_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScanRegionRequest); i {
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
		file_def_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAssociatedObjectRequest); i {
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
		file_def_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectList); i {
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
		file_def_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectListWithAssociatedObjects); i {
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
		file_def_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NameList); i {
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
		file_def_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NameSearch); i {
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
			RawDescriptor: file_def_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_def_proto_goTypes,
		DependencyIndexes: file_def_proto_depIdxs,
		MessageInfos:      file_def_proto_msgTypes,
	}.Build()
	File_def_proto = out.File
	file_def_proto_rawDesc = nil
	file_def_proto_goTypes = nil
	file_def_proto_depIdxs = nil
}