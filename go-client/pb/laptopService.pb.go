// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: proto/laptopService.proto

package ____laptop

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type CreateLaptopRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Laptop *Laptop `protobuf:"bytes,1,opt,name=laptop,proto3" json:"laptop,omitempty"`
}

func (x *CreateLaptopRequest) Reset() {
	*x = CreateLaptopRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_laptopService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLaptopRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLaptopRequest) ProtoMessage() {}

func (x *CreateLaptopRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_laptopService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLaptopRequest.ProtoReflect.Descriptor instead.
func (*CreateLaptopRequest) Descriptor() ([]byte, []int) {
	return file_proto_laptopService_proto_rawDescGZIP(), []int{0}
}

func (x *CreateLaptopRequest) GetLaptop() *Laptop {
	if x != nil {
		return x.Laptop
	}
	return nil
}

type CreateLaptopResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Laptop *Laptop `protobuf:"bytes,1,opt,name=laptop,proto3" json:"laptop,omitempty"`
}

func (x *CreateLaptopResponse) Reset() {
	*x = CreateLaptopResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_laptopService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLaptopResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLaptopResponse) ProtoMessage() {}

func (x *CreateLaptopResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_laptopService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLaptopResponse.ProtoReflect.Descriptor instead.
func (*CreateLaptopResponse) Descriptor() ([]byte, []int) {
	return file_proto_laptopService_proto_rawDescGZIP(), []int{1}
}

func (x *CreateLaptopResponse) GetLaptop() *Laptop {
	if x != nil {
		return x.Laptop
	}
	return nil
}

type FetchLaptopResposne struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Laptop *Laptop `protobuf:"bytes,1,opt,name=laptop,proto3" json:"laptop,omitempty"`
}

func (x *FetchLaptopResposne) Reset() {
	*x = FetchLaptopResposne{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_laptopService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchLaptopResposne) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchLaptopResposne) ProtoMessage() {}

func (x *FetchLaptopResposne) ProtoReflect() protoreflect.Message {
	mi := &file_proto_laptopService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchLaptopResposne.ProtoReflect.Descriptor instead.
func (*FetchLaptopResposne) Descriptor() ([]byte, []int) {
	return file_proto_laptopService_proto_rawDescGZIP(), []int{2}
}

func (x *FetchLaptopResposne) GetLaptop() *Laptop {
	if x != nil {
		return x.Laptop
	}
	return nil
}

type Hello struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hello string `protobuf:"bytes,1,opt,name=Hello,proto3" json:"Hello,omitempty"`
}

func (x *Hello) Reset() {
	*x = Hello{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_laptopService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hello) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hello) ProtoMessage() {}

func (x *Hello) ProtoReflect() protoreflect.Message {
	mi := &file_proto_laptopService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hello.ProtoReflect.Descriptor instead.
func (*Hello) Descriptor() ([]byte, []int) {
	return file_proto_laptopService_proto_rawDescGZIP(), []int{3}
}

func (x *Hello) GetHello() string {
	if x != nil {
		return x.Hello
	}
	return ""
}

type Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Void) Reset() {
	*x = Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_laptopService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_proto_laptopService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Void.ProtoReflect.Descriptor instead.
func (*Void) Descriptor() ([]byte, []int) {
	return file_proto_laptopService_proto_rawDescGZIP(), []int{4}
}

type UploadImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//	*UploadImageRequest_Info
	//	*UploadImageRequest_ChunkData
	Data isUploadImageRequest_Data `protobuf_oneof:"data"`
}

func (x *UploadImageRequest) Reset() {
	*x = UploadImageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_laptopService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadImageRequest) ProtoMessage() {}

func (x *UploadImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_laptopService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadImageRequest.ProtoReflect.Descriptor instead.
func (*UploadImageRequest) Descriptor() ([]byte, []int) {
	return file_proto_laptopService_proto_rawDescGZIP(), []int{5}
}

func (m *UploadImageRequest) GetData() isUploadImageRequest_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *UploadImageRequest) GetInfo() *ImageInfo {
	if x, ok := x.GetData().(*UploadImageRequest_Info); ok {
		return x.Info
	}
	return nil
}

func (x *UploadImageRequest) GetChunkData() []byte {
	if x, ok := x.GetData().(*UploadImageRequest_ChunkData); ok {
		return x.ChunkData
	}
	return nil
}

type isUploadImageRequest_Data interface {
	isUploadImageRequest_Data()
}

type UploadImageRequest_Info struct {
	Info *ImageInfo `protobuf:"bytes,1,opt,name=info,proto3,oneof"`
}

type UploadImageRequest_ChunkData struct {
	ChunkData []byte `protobuf:"bytes,2,opt,name=chunk_data,json=chunkData,proto3,oneof"`
}

func (*UploadImageRequest_Info) isUploadImageRequest_Data() {}

func (*UploadImageRequest_ChunkData) isUploadImageRequest_Data() {}

type ImageInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LaptopId  string `protobuf:"bytes,1,opt,name=laptop_id,json=laptopId,proto3" json:"laptop_id,omitempty"`
	ImageType string `protobuf:"bytes,2,opt,name=image_type,json=imageType,proto3" json:"image_type,omitempty"`
}

func (x *ImageInfo) Reset() {
	*x = ImageInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_laptopService_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageInfo) ProtoMessage() {}

func (x *ImageInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_laptopService_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageInfo.ProtoReflect.Descriptor instead.
func (*ImageInfo) Descriptor() ([]byte, []int) {
	return file_proto_laptopService_proto_rawDescGZIP(), []int{6}
}

func (x *ImageInfo) GetLaptopId() string {
	if x != nil {
		return x.LaptopId
	}
	return ""
}

func (x *ImageInfo) GetImageType() string {
	if x != nil {
		return x.ImageType
	}
	return ""
}

type UploadImageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LaptopId  string `protobuf:"bytes,1,opt,name=laptop_id,json=laptopId,proto3" json:"laptop_id,omitempty"`
	ImageId   string `protobuf:"bytes,2,opt,name=image_id,json=imageId,proto3" json:"image_id,omitempty"`
	ImageSize uint32 `protobuf:"varint,3,opt,name=image_size,json=imageSize,proto3" json:"image_size,omitempty"`
}

func (x *UploadImageResponse) Reset() {
	*x = UploadImageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_laptopService_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadImageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadImageResponse) ProtoMessage() {}

func (x *UploadImageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_laptopService_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadImageResponse.ProtoReflect.Descriptor instead.
func (*UploadImageResponse) Descriptor() ([]byte, []int) {
	return file_proto_laptopService_proto_rawDescGZIP(), []int{7}
}

func (x *UploadImageResponse) GetLaptopId() string {
	if x != nil {
		return x.LaptopId
	}
	return ""
}

func (x *UploadImageResponse) GetImageId() string {
	if x != nil {
		return x.ImageId
	}
	return ""
}

func (x *UploadImageResponse) GetImageSize() uint32 {
	if x != nil {
		return x.ImageSize
	}
	return 0
}

var File_proto_laptopService_proto protoreflect.FileDescriptor

var file_proto_laptopService_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6c, 0x61, 0x70,
	0x74, 0x6f, 0x70, 0x1a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x61, 0x70, 0x74, 0x6f,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3d, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26,
	0x0a, 0x06, 0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x2e, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x52, 0x06,
	0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x22, 0x3e, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26,
	0x0a, 0x06, 0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x2e, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x52, 0x06,
	0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x22, 0x3d, 0x0a, 0x13, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4c,
	0x61, 0x70, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x73, 0x6e, 0x65, 0x12, 0x26, 0x0a,
	0x06, 0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x2e, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x52, 0x06, 0x6c,
	0x61, 0x70, 0x74, 0x6f, 0x70, 0x22, 0x1d, 0x0a, 0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x14,
	0x0a, 0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x22, 0x06, 0x0a, 0x04, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x66, 0x0a, 0x12,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x27, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0a, 0x63,
	0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48,
	0x00, 0x52, 0x09, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x42, 0x06, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x47, 0x0a, 0x09, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x6c, 0x0a,
	0x13, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x49,
	0x64, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x32, 0xab, 0x02, 0x0a, 0x0d,
	0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a,
	0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x12, 0x1b, 0x2e,
	0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x70,
	0x74, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6c, 0x61, 0x70,
	0x74, 0x6f, 0x70, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0f, 0x46, 0x65,
	0x74, 0x63, 0x68, 0x41, 0x6c, 0x6c, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x73, 0x12, 0x0c, 0x2e,
	0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x1a, 0x1b, 0x2e, 0x6c, 0x61,
	0x70, 0x74, 0x6f, 0x70, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x73, 0x6e, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x50, 0x0a, 0x11,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x12, 0x1a, 0x2e, 0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x39,
	0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x0d, 0x2e, 0x6c, 0x61, 0x70,
	0x74, 0x6f, 0x70, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x1a, 0x1c, 0x2e, 0x6c, 0x61, 0x70, 0x74,
	0x6f, 0x70, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x3b, 0x5f,
	0x5f, 0x5f, 0x6c, 0x61, 0x70, 0x74, 0x6f, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_laptopService_proto_rawDescOnce sync.Once
	file_proto_laptopService_proto_rawDescData = file_proto_laptopService_proto_rawDesc
)

func file_proto_laptopService_proto_rawDescGZIP() []byte {
	file_proto_laptopService_proto_rawDescOnce.Do(func() {
		file_proto_laptopService_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_laptopService_proto_rawDescData)
	})
	return file_proto_laptopService_proto_rawDescData
}

var file_proto_laptopService_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_laptopService_proto_goTypes = []interface{}{
	(*CreateLaptopRequest)(nil),  // 0: laptop.CreateLaptopRequest
	(*CreateLaptopResponse)(nil), // 1: laptop.CreateLaptopResponse
	(*FetchLaptopResposne)(nil),  // 2: laptop.FetchLaptopResposne
	(*Hello)(nil),                // 3: laptop.Hello
	(*Void)(nil),                 // 4: laptop.Void
	(*UploadImageRequest)(nil),   // 5: laptop.UploadImageRequest
	(*ImageInfo)(nil),            // 6: laptop.ImageInfo
	(*UploadImageResponse)(nil),  // 7: laptop.UploadImageResponse
	(*Laptop)(nil),               // 8: laptop.Laptop
}
var file_proto_laptopService_proto_depIdxs = []int32{
	8, // 0: laptop.CreateLaptopRequest.laptop:type_name -> laptop.Laptop
	8, // 1: laptop.CreateLaptopResponse.laptop:type_name -> laptop.Laptop
	8, // 2: laptop.FetchLaptopResposne.laptop:type_name -> laptop.Laptop
	6, // 3: laptop.UploadImageRequest.info:type_name -> laptop.ImageInfo
	0, // 4: laptop.LaptopService.CreateLaptop:input_type -> laptop.CreateLaptopRequest
	4, // 5: laptop.LaptopService.FetchAllLaptops:input_type -> laptop.Void
	5, // 6: laptop.LaptopService.UploadLaptopImage:input_type -> laptop.UploadImageRequest
	3, // 7: laptop.LaptopService.SayHello:input_type -> laptop.Hello
	1, // 8: laptop.LaptopService.CreateLaptop:output_type -> laptop.CreateLaptopResponse
	2, // 9: laptop.LaptopService.FetchAllLaptops:output_type -> laptop.FetchLaptopResposne
	7, // 10: laptop.LaptopService.UploadLaptopImage:output_type -> laptop.UploadImageResponse
	1, // 11: laptop.LaptopService.SayHello:output_type -> laptop.CreateLaptopResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_laptopService_proto_init() }
func file_proto_laptopService_proto_init() {
	if File_proto_laptopService_proto != nil {
		return
	}
	file_proto_laptop_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_laptopService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLaptopRequest); i {
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
		file_proto_laptopService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLaptopResponse); i {
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
		file_proto_laptopService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchLaptopResposne); i {
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
		file_proto_laptopService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hello); i {
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
		file_proto_laptopService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Void); i {
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
		file_proto_laptopService_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadImageRequest); i {
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
		file_proto_laptopService_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageInfo); i {
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
		file_proto_laptopService_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadImageResponse); i {
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
	file_proto_laptopService_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*UploadImageRequest_Info)(nil),
		(*UploadImageRequest_ChunkData)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_laptopService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_laptopService_proto_goTypes,
		DependencyIndexes: file_proto_laptopService_proto_depIdxs,
		MessageInfos:      file_proto_laptopService_proto_msgTypes,
	}.Build()
	File_proto_laptopService_proto = out.File
	file_proto_laptopService_proto_rawDesc = nil
	file_proto_laptopService_proto_goTypes = nil
	file_proto_laptopService_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LaptopServiceClient is the client API for LaptopService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LaptopServiceClient interface {
	CreateLaptop(ctx context.Context, in *CreateLaptopRequest, opts ...grpc.CallOption) (*CreateLaptopResponse, error)
	FetchAllLaptops(ctx context.Context, in *Void, opts ...grpc.CallOption) (LaptopService_FetchAllLaptopsClient, error)
	UploadLaptopImage(ctx context.Context, opts ...grpc.CallOption) (LaptopService_UploadLaptopImageClient, error)
	SayHello(ctx context.Context, in *Hello, opts ...grpc.CallOption) (*CreateLaptopResponse, error)
}

type laptopServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLaptopServiceClient(cc grpc.ClientConnInterface) LaptopServiceClient {
	return &laptopServiceClient{cc}
}

func (c *laptopServiceClient) CreateLaptop(ctx context.Context, in *CreateLaptopRequest, opts ...grpc.CallOption) (*CreateLaptopResponse, error) {
	out := new(CreateLaptopResponse)
	err := c.cc.Invoke(ctx, "/laptop.LaptopService/CreateLaptop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *laptopServiceClient) FetchAllLaptops(ctx context.Context, in *Void, opts ...grpc.CallOption) (LaptopService_FetchAllLaptopsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LaptopService_serviceDesc.Streams[0], "/laptop.LaptopService/FetchAllLaptops", opts...)
	if err != nil {
		return nil, err
	}
	x := &laptopServiceFetchAllLaptopsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LaptopService_FetchAllLaptopsClient interface {
	Recv() (*FetchLaptopResposne, error)
	grpc.ClientStream
}

type laptopServiceFetchAllLaptopsClient struct {
	grpc.ClientStream
}

func (x *laptopServiceFetchAllLaptopsClient) Recv() (*FetchLaptopResposne, error) {
	m := new(FetchLaptopResposne)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *laptopServiceClient) UploadLaptopImage(ctx context.Context, opts ...grpc.CallOption) (LaptopService_UploadLaptopImageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LaptopService_serviceDesc.Streams[1], "/laptop.LaptopService/UploadLaptopImage", opts...)
	if err != nil {
		return nil, err
	}
	x := &laptopServiceUploadLaptopImageClient{stream}
	return x, nil
}

type LaptopService_UploadLaptopImageClient interface {
	Send(*UploadImageRequest) error
	CloseAndRecv() (*UploadImageResponse, error)
	grpc.ClientStream
}

type laptopServiceUploadLaptopImageClient struct {
	grpc.ClientStream
}

func (x *laptopServiceUploadLaptopImageClient) Send(m *UploadImageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *laptopServiceUploadLaptopImageClient) CloseAndRecv() (*UploadImageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadImageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *laptopServiceClient) SayHello(ctx context.Context, in *Hello, opts ...grpc.CallOption) (*CreateLaptopResponse, error) {
	out := new(CreateLaptopResponse)
	err := c.cc.Invoke(ctx, "/laptop.LaptopService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LaptopServiceServer is the server API for LaptopService service.
type LaptopServiceServer interface {
	CreateLaptop(context.Context, *CreateLaptopRequest) (*CreateLaptopResponse, error)
	FetchAllLaptops(*Void, LaptopService_FetchAllLaptopsServer) error
	UploadLaptopImage(LaptopService_UploadLaptopImageServer) error
	SayHello(context.Context, *Hello) (*CreateLaptopResponse, error)
}

// UnimplementedLaptopServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLaptopServiceServer struct {
}

func (*UnimplementedLaptopServiceServer) CreateLaptop(context.Context, *CreateLaptopRequest) (*CreateLaptopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLaptop not implemented")
}
func (*UnimplementedLaptopServiceServer) FetchAllLaptops(*Void, LaptopService_FetchAllLaptopsServer) error {
	return status.Errorf(codes.Unimplemented, "method FetchAllLaptops not implemented")
}
func (*UnimplementedLaptopServiceServer) UploadLaptopImage(LaptopService_UploadLaptopImageServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadLaptopImage not implemented")
}
func (*UnimplementedLaptopServiceServer) SayHello(context.Context, *Hello) (*CreateLaptopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}

func RegisterLaptopServiceServer(s *grpc.Server, srv LaptopServiceServer) {
	s.RegisterService(&_LaptopService_serviceDesc, srv)
}

func _LaptopService_CreateLaptop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLaptopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LaptopServiceServer).CreateLaptop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/laptop.LaptopService/CreateLaptop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LaptopServiceServer).CreateLaptop(ctx, req.(*CreateLaptopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LaptopService_FetchAllLaptops_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Void)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LaptopServiceServer).FetchAllLaptops(m, &laptopServiceFetchAllLaptopsServer{stream})
}

type LaptopService_FetchAllLaptopsServer interface {
	Send(*FetchLaptopResposne) error
	grpc.ServerStream
}

type laptopServiceFetchAllLaptopsServer struct {
	grpc.ServerStream
}

func (x *laptopServiceFetchAllLaptopsServer) Send(m *FetchLaptopResposne) error {
	return x.ServerStream.SendMsg(m)
}

func _LaptopService_UploadLaptopImage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LaptopServiceServer).UploadLaptopImage(&laptopServiceUploadLaptopImageServer{stream})
}

type LaptopService_UploadLaptopImageServer interface {
	SendAndClose(*UploadImageResponse) error
	Recv() (*UploadImageRequest, error)
	grpc.ServerStream
}

type laptopServiceUploadLaptopImageServer struct {
	grpc.ServerStream
}

func (x *laptopServiceUploadLaptopImageServer) SendAndClose(m *UploadImageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *laptopServiceUploadLaptopImageServer) Recv() (*UploadImageRequest, error) {
	m := new(UploadImageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _LaptopService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Hello)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LaptopServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/laptop.LaptopService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LaptopServiceServer).SayHello(ctx, req.(*Hello))
	}
	return interceptor(ctx, in, info, handler)
}

var _LaptopService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "laptop.LaptopService",
	HandlerType: (*LaptopServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLaptop",
			Handler:    _LaptopService_CreateLaptop_Handler,
		},
		{
			MethodName: "SayHello",
			Handler:    _LaptopService_SayHello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FetchAllLaptops",
			Handler:       _LaptopService_FetchAllLaptops_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "UploadLaptopImage",
			Handler:       _LaptopService_UploadLaptopImage_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "proto/laptopService.proto",
}