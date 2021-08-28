// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/ova-food-api/ova-food-api.proto

package ova_food_api

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FoodType int32

const (
	FoodType_Unknown FoodType = 0
	FoodType_Drinks  FoodType = 1
	FoodType_Foods   FoodType = 2
)

// Enum value maps for FoodType.
var (
	FoodType_name = map[int32]string{
		0: "Unknown",
		1: "Drinks",
		2: "Foods",
	}
	FoodType_value = map[string]int32{
		"Unknown": 0,
		"Drinks":  1,
		"Foods":   2,
	}
)

func (x FoodType) Enum() *FoodType {
	p := new(FoodType)
	*p = x
	return p
}

func (x FoodType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FoodType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_ova_food_api_ova_food_api_proto_enumTypes[0].Descriptor()
}

func (FoodType) Type() protoreflect.EnumType {
	return &file_api_ova_food_api_ova_food_api_proto_enumTypes[0]
}

func (x FoodType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FoodType.Descriptor instead.
func (FoodType) EnumDescriptor() ([]byte, []int) {
	return file_api_ova_food_api_ova_food_api_proto_rawDescGZIP(), []int{0}
}

type CreationFood struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      uint64   `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FoodT       FoodType `protobuf:"varint,3,opt,name=food_t,json=foodT,proto3,enum=ova.food.api.FoodType" json:"food_t,omitempty"`
	Name        string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	PortionSize float32  `protobuf:"fixed32,5,opt,name=portion_size,json=portionSize,proto3" json:"portion_size,omitempty"`
}

func (x *CreationFood) Reset() {
	*x = CreationFood{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ova_food_api_ova_food_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreationFood) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreationFood) ProtoMessage() {}

func (x *CreationFood) ProtoReflect() protoreflect.Message {
	mi := &file_api_ova_food_api_ova_food_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreationFood.ProtoReflect.Descriptor instead.
func (*CreationFood) Descriptor() ([]byte, []int) {
	return file_api_ova_food_api_ova_food_api_proto_rawDescGZIP(), []int{0}
}

func (x *CreationFood) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreationFood) GetFoodT() FoodType {
	if x != nil {
		return x.FoodT
	}
	return FoodType_Unknown
}

func (x *CreationFood) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreationFood) GetPortionSize() float32 {
	if x != nil {
		return x.PortionSize
	}
	return 0
}

type Food struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FoodId      uint64   `protobuf:"varint,1,opt,name=food_id,json=foodId,proto3" json:"food_id,omitempty"`
	UserId      uint64   `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FoodT       FoodType `protobuf:"varint,3,opt,name=food_t,json=foodT,proto3,enum=ova.food.api.FoodType" json:"food_t,omitempty"`
	Name        string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	PortionSize float32  `protobuf:"fixed32,5,opt,name=portion_size,json=portionSize,proto3" json:"portion_size,omitempty"`
}

func (x *Food) Reset() {
	*x = Food{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ova_food_api_ova_food_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Food) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Food) ProtoMessage() {}

func (x *Food) ProtoReflect() protoreflect.Message {
	mi := &file_api_ova_food_api_ova_food_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Food.ProtoReflect.Descriptor instead.
func (*Food) Descriptor() ([]byte, []int) {
	return file_api_ova_food_api_ova_food_api_proto_rawDescGZIP(), []int{1}
}

func (x *Food) GetFoodId() uint64 {
	if x != nil {
		return x.FoodId
	}
	return 0
}

func (x *Food) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Food) GetFoodT() FoodType {
	if x != nil {
		return x.FoodT
	}
	return FoodType_Unknown
}

func (x *Food) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Food) GetPortionSize() float32 {
	if x != nil {
		return x.PortionSize
	}
	return 0
}

type CreateFoodV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Food *CreationFood `protobuf:"bytes,1,opt,name=food,proto3" json:"food,omitempty"`
}

func (x *CreateFoodV1Request) Reset() {
	*x = CreateFoodV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ova_food_api_ova_food_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFoodV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFoodV1Request) ProtoMessage() {}

func (x *CreateFoodV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_ova_food_api_ova_food_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFoodV1Request.ProtoReflect.Descriptor instead.
func (*CreateFoodV1Request) Descriptor() ([]byte, []int) {
	return file_api_ova_food_api_ova_food_api_proto_rawDescGZIP(), []int{2}
}

func (x *CreateFoodV1Request) GetFood() *CreationFood {
	if x != nil {
		return x.Food
	}
	return nil
}

type DescribeFoodV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FoodId uint64 `protobuf:"varint,1,opt,name=food_id,json=foodId,proto3" json:"food_id,omitempty"`
}

func (x *DescribeFoodV1Request) Reset() {
	*x = DescribeFoodV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ova_food_api_ova_food_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeFoodV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeFoodV1Request) ProtoMessage() {}

func (x *DescribeFoodV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_ova_food_api_ova_food_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeFoodV1Request.ProtoReflect.Descriptor instead.
func (*DescribeFoodV1Request) Descriptor() ([]byte, []int) {
	return file_api_ova_food_api_ova_food_api_proto_rawDescGZIP(), []int{3}
}

func (x *DescribeFoodV1Request) GetFoodId() uint64 {
	if x != nil {
		return x.FoodId
	}
	return 0
}

type DescribeFoodV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Food *Food `protobuf:"bytes,1,opt,name=food,proto3" json:"food,omitempty"`
}

func (x *DescribeFoodV1Response) Reset() {
	*x = DescribeFoodV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ova_food_api_ova_food_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeFoodV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeFoodV1Response) ProtoMessage() {}

func (x *DescribeFoodV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_ova_food_api_ova_food_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeFoodV1Response.ProtoReflect.Descriptor instead.
func (*DescribeFoodV1Response) Descriptor() ([]byte, []int) {
	return file_api_ova_food_api_ova_food_api_proto_rawDescGZIP(), []int{4}
}

func (x *DescribeFoodV1Response) GetFood() *Food {
	if x != nil {
		return x.Food
	}
	return nil
}

type ListFoodsV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []uint64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *ListFoodsV1Request) Reset() {
	*x = ListFoodsV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ova_food_api_ova_food_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFoodsV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFoodsV1Request) ProtoMessage() {}

func (x *ListFoodsV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_ova_food_api_ova_food_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFoodsV1Request.ProtoReflect.Descriptor instead.
func (*ListFoodsV1Request) Descriptor() ([]byte, []int) {
	return file_api_ova_food_api_ova_food_api_proto_rawDescGZIP(), []int{5}
}

func (x *ListFoodsV1Request) GetIds() []uint64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type ListFoodsV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Foods map[uint64]*Food `protobuf:"bytes,1,rep,name=foods,proto3" json:"foods,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ListFoodsV1Response) Reset() {
	*x = ListFoodsV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ova_food_api_ova_food_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFoodsV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFoodsV1Response) ProtoMessage() {}

func (x *ListFoodsV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_ova_food_api_ova_food_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFoodsV1Response.ProtoReflect.Descriptor instead.
func (*ListFoodsV1Response) Descriptor() ([]byte, []int) {
	return file_api_ova_food_api_ova_food_api_proto_rawDescGZIP(), []int{6}
}

func (x *ListFoodsV1Response) GetFoods() map[uint64]*Food {
	if x != nil {
		return x.Foods
	}
	return nil
}

type RemoveFoodV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FoodId uint64 `protobuf:"varint,1,opt,name=food_id,json=foodId,proto3" json:"food_id,omitempty"`
}

func (x *RemoveFoodV1Request) Reset() {
	*x = RemoveFoodV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ova_food_api_ova_food_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveFoodV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveFoodV1Request) ProtoMessage() {}

func (x *RemoveFoodV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_ova_food_api_ova_food_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveFoodV1Request.ProtoReflect.Descriptor instead.
func (*RemoveFoodV1Request) Descriptor() ([]byte, []int) {
	return file_api_ova_food_api_ova_food_api_proto_rawDescGZIP(), []int{7}
}

func (x *RemoveFoodV1Request) GetFoodId() uint64 {
	if x != nil {
		return x.FoodId
	}
	return 0
}

var File_api_ova_food_api_ova_food_api_proto protoreflect.FileDescriptor

var file_api_ova_food_api_ova_food_api_proto_rawDesc = []byte{
	0x0a, 0x23, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x76, 0x61, 0x2d, 0x66, 0x6f, 0x6f, 0x64, 0x2d, 0x61,
	0x70, 0x69, 0x2f, 0x6f, 0x76, 0x61, 0x2d, 0x66, 0x6f, 0x6f, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6f, 0x76, 0x61, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x2e,
	0x61, 0x70, 0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdc, 0x01, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x6f, 0x64, 0x12, 0x20, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02,
	0x28, 0x00, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x06, 0x66, 0x6f,
	0x6f, 0x64, 0x5f, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x6f, 0x76, 0x61,
	0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x46, 0x6f, 0x6f, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x05, 0x66, 0x6f,
	0x6f, 0x64, 0x54, 0x12, 0x42, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x2e, 0xfa, 0x42, 0x2b, 0x72, 0x29, 0x28, 0x80, 0x02, 0x32, 0x24, 0x5e, 0x5b, 0x5e,
	0x5b, 0x30, 0x2d, 0x39, 0x5d, 0x41, 0x2d, 0x5a, 0x61, 0x2d, 0x7a, 0x5d, 0x2b, 0x28, 0x20, 0x5b,
	0x5e, 0x5b, 0x30, 0x2d, 0x39, 0x5d, 0x41, 0x2d, 0x5a, 0x61, 0x2d, 0x7a, 0x5d, 0x2b, 0x29, 0x2a,
	0x24, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x0c, 0x70, 0x6f, 0x72, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x42, 0x0a, 0xfa,
	0x42, 0x07, 0x0a, 0x05, 0x2d, 0x00, 0x00, 0x00, 0x00, 0x52, 0x0b, 0x70, 0x6f, 0x72, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x9e, 0x01, 0x0a, 0x04, 0x46, 0x6f, 0x6f, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x66, 0x6f, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x66, 0x6f, 0x6f, 0x64, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x2d, 0x0a, 0x06, 0x66, 0x6f, 0x6f, 0x64, 0x5f, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x16, 0x2e, 0x6f, 0x76, 0x61, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x46, 0x6f, 0x6f, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x66, 0x6f, 0x6f, 0x64, 0x54,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x70, 0x6f, 0x72, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x4f, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x46, 0x6f, 0x6f, 0x64, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38,
	0x0a, 0x04, 0x66, 0x6f, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6f,
	0x76, 0x61, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x6f, 0x64, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02,
	0x10, 0x01, 0x52, 0x04, 0x66, 0x6f, 0x6f, 0x64, 0x22, 0x39, 0x0a, 0x15, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x20, 0x0a, 0x07, 0x66, 0x6f, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x28, 0x00, 0x52, 0x06, 0x66, 0x6f, 0x6f,
	0x64, 0x49, 0x64, 0x22, 0x40, 0x0a, 0x16, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x46,
	0x6f, 0x6f, 0x64, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a,
	0x04, 0x66, 0x6f, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6f, 0x76,
	0x61, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x46, 0x6f, 0x6f, 0x64, 0x52,
	0x04, 0x66, 0x6f, 0x6f, 0x64, 0x22, 0x34, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x6f,
	0x64, 0x73, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x03, 0x69,
	0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x04, 0x42, 0x0c, 0xfa, 0x42, 0x09, 0x92, 0x01, 0x06,
	0x22, 0x04, 0x32, 0x02, 0x28, 0x00, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0xa7, 0x01, 0x0a, 0x13,
	0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x6f, 0x64, 0x73, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x05, 0x66, 0x6f, 0x6f, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6f, 0x76, 0x61, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x6f, 0x64, 0x73, 0x56, 0x31, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x46, 0x6f, 0x6f, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x05, 0x66, 0x6f, 0x6f, 0x64, 0x73, 0x1a, 0x4c, 0x0a, 0x0a, 0x46, 0x6f, 0x6f, 0x64, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6f, 0x76, 0x61, 0x2e, 0x66, 0x6f, 0x6f,
	0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x37, 0x0a, 0x13, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46,
	0x6f, 0x6f, 0x64, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x07,
	0x66, 0x6f, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x32, 0x02, 0x28, 0x00, 0x52, 0x06, 0x66, 0x6f, 0x6f, 0x64, 0x49, 0x64, 0x2a, 0x2e,
	0x0a, 0x08, 0x46, 0x6f, 0x6f, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e,
	0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x72, 0x69, 0x6e, 0x6b,
	0x73, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x46, 0x6f, 0x6f, 0x64, 0x73, 0x10, 0x02, 0x32, 0xb9,
	0x03, 0x0a, 0x0a, 0x4f, 0x76, 0x61, 0x46, 0x6f, 0x6f, 0x64, 0x41, 0x70, 0x69, 0x12, 0x62, 0x0a,
	0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x56, 0x31, 0x12, 0x21, 0x2e,
	0x6f, 0x76, 0x61, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11,
	0x22, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x6f, 0x6f, 0x64, 0x73, 0x3a, 0x04, 0x66, 0x6f, 0x6f,
	0x64, 0x12, 0x78, 0x0a, 0x0e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x46, 0x6f, 0x6f,
	0x64, 0x56, 0x31, 0x12, 0x23, 0x2e, 0x6f, 0x76, 0x61, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x56,
	0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x6f, 0x76, 0x61, 0x2e, 0x66,
	0x6f, 0x6f, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x46, 0x6f, 0x6f, 0x64, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x6f, 0x6f, 0x64,
	0x73, 0x2f, 0x7b, 0x66, 0x6f, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x65, 0x0a, 0x0b, 0x4c,
	0x69, 0x73, 0x74, 0x46, 0x6f, 0x6f, 0x64, 0x73, 0x56, 0x31, 0x12, 0x20, 0x2e, 0x6f, 0x76, 0x61,
	0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f,
	0x6f, 0x64, 0x73, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6f,
	0x76, 0x61, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x46, 0x6f, 0x6f, 0x64, 0x73, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x6f, 0x6f,
	0x64, 0x73, 0x12, 0x66, 0x0a, 0x0c, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x6f, 0x6f, 0x64,
	0x56, 0x31, 0x12, 0x21, 0x2e, 0x6f, 0x76, 0x61, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x56, 0x31, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1b, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x15, 0x2a, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x6f, 0x6f, 0x64, 0x73,
	0x2f, 0x7b, 0x66, 0x6f, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x7a, 0x6f, 0x6e, 0x76, 0x61, 0x2f,
	0x6f, 0x76, 0x61, 0x2d, 0x66, 0x6f, 0x6f, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x6f, 0x76, 0x61, 0x2d, 0x66, 0x6f, 0x6f, 0x64, 0x2d, 0x61, 0x70, 0x69, 0x3b, 0x6f, 0x76,
	0x61, 0x5f, 0x66, 0x6f, 0x6f, 0x64, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_ova_food_api_ova_food_api_proto_rawDescOnce sync.Once
	file_api_ova_food_api_ova_food_api_proto_rawDescData = file_api_ova_food_api_ova_food_api_proto_rawDesc
)

func file_api_ova_food_api_ova_food_api_proto_rawDescGZIP() []byte {
	file_api_ova_food_api_ova_food_api_proto_rawDescOnce.Do(func() {
		file_api_ova_food_api_ova_food_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_ova_food_api_ova_food_api_proto_rawDescData)
	})
	return file_api_ova_food_api_ova_food_api_proto_rawDescData
}

var file_api_ova_food_api_ova_food_api_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_ova_food_api_ova_food_api_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_ova_food_api_ova_food_api_proto_goTypes = []interface{}{
	(FoodType)(0),                  // 0: ova.food.api.FoodType
	(*CreationFood)(nil),           // 1: ova.food.api.CreationFood
	(*Food)(nil),                   // 2: ova.food.api.Food
	(*CreateFoodV1Request)(nil),    // 3: ova.food.api.CreateFoodV1Request
	(*DescribeFoodV1Request)(nil),  // 4: ova.food.api.DescribeFoodV1Request
	(*DescribeFoodV1Response)(nil), // 5: ova.food.api.DescribeFoodV1Response
	(*ListFoodsV1Request)(nil),     // 6: ova.food.api.ListFoodsV1Request
	(*ListFoodsV1Response)(nil),    // 7: ova.food.api.ListFoodsV1Response
	(*RemoveFoodV1Request)(nil),    // 8: ova.food.api.RemoveFoodV1Request
	nil,                            // 9: ova.food.api.ListFoodsV1Response.FoodsEntry
	(*emptypb.Empty)(nil),          // 10: google.protobuf.Empty
}
var file_api_ova_food_api_ova_food_api_proto_depIdxs = []int32{
	0,  // 0: ova.food.api.CreationFood.food_t:type_name -> ova.food.api.FoodType
	0,  // 1: ova.food.api.Food.food_t:type_name -> ova.food.api.FoodType
	1,  // 2: ova.food.api.CreateFoodV1Request.food:type_name -> ova.food.api.CreationFood
	2,  // 3: ova.food.api.DescribeFoodV1Response.food:type_name -> ova.food.api.Food
	9,  // 4: ova.food.api.ListFoodsV1Response.foods:type_name -> ova.food.api.ListFoodsV1Response.FoodsEntry
	2,  // 5: ova.food.api.ListFoodsV1Response.FoodsEntry.value:type_name -> ova.food.api.Food
	3,  // 6: ova.food.api.OvaFoodApi.CreateFoodV1:input_type -> ova.food.api.CreateFoodV1Request
	4,  // 7: ova.food.api.OvaFoodApi.DescribeFoodV1:input_type -> ova.food.api.DescribeFoodV1Request
	6,  // 8: ova.food.api.OvaFoodApi.ListFoodsV1:input_type -> ova.food.api.ListFoodsV1Request
	8,  // 9: ova.food.api.OvaFoodApi.RemoveFoodV1:input_type -> ova.food.api.RemoveFoodV1Request
	10, // 10: ova.food.api.OvaFoodApi.CreateFoodV1:output_type -> google.protobuf.Empty
	5,  // 11: ova.food.api.OvaFoodApi.DescribeFoodV1:output_type -> ova.food.api.DescribeFoodV1Response
	7,  // 12: ova.food.api.OvaFoodApi.ListFoodsV1:output_type -> ova.food.api.ListFoodsV1Response
	10, // 13: ova.food.api.OvaFoodApi.RemoveFoodV1:output_type -> google.protobuf.Empty
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_api_ova_food_api_ova_food_api_proto_init() }
func file_api_ova_food_api_ova_food_api_proto_init() {
	if File_api_ova_food_api_ova_food_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_ova_food_api_ova_food_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreationFood); i {
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
		file_api_ova_food_api_ova_food_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Food); i {
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
		file_api_ova_food_api_ova_food_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFoodV1Request); i {
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
		file_api_ova_food_api_ova_food_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeFoodV1Request); i {
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
		file_api_ova_food_api_ova_food_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeFoodV1Response); i {
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
		file_api_ova_food_api_ova_food_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFoodsV1Request); i {
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
		file_api_ova_food_api_ova_food_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFoodsV1Response); i {
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
		file_api_ova_food_api_ova_food_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveFoodV1Request); i {
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
			RawDescriptor: file_api_ova_food_api_ova_food_api_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_ova_food_api_ova_food_api_proto_goTypes,
		DependencyIndexes: file_api_ova_food_api_ova_food_api_proto_depIdxs,
		EnumInfos:         file_api_ova_food_api_ova_food_api_proto_enumTypes,
		MessageInfos:      file_api_ova_food_api_ova_food_api_proto_msgTypes,
	}.Build()
	File_api_ova_food_api_ova_food_api_proto = out.File
	file_api_ova_food_api_ova_food_api_proto_rawDesc = nil
	file_api_ova_food_api_ova_food_api_proto_goTypes = nil
	file_api_ova_food_api_ova_food_api_proto_depIdxs = nil
}