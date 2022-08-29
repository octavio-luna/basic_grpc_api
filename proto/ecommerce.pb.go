// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: proto/ecommerce.proto

package proto

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

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Label        string   `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	Type         int32    `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	Downloadlink *string  `protobuf:"bytes,4,opt,name=downloadlink,proto3,oneof" json:"downloadlink,omitempty"`
	Weight       *float32 `protobuf:"fixed32,5,opt,name=weight,proto3,oneof" json:"weight,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ecommerce_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ecommerce_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_proto_ecommerce_proto_rawDescGZIP(), []int{0}
}

func (x *Product) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Product) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *Product) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *Product) GetDownloadlink() string {
	if x != nil && x.Downloadlink != nil {
		return *x.Downloadlink
	}
	return ""
}

func (x *Product) GetWeight() float32 {
	if x != nil && x.Weight != nil {
		return *x.Weight
	}
	return 0
}

type AddProduct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductID int32 `protobuf:"varint,1,opt,name=productID,proto3" json:"productID,omitempty"`
	Quantity  int32 `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	UserID    int32 `protobuf:"varint,3,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *AddProduct) Reset() {
	*x = AddProduct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ecommerce_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddProduct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddProduct) ProtoMessage() {}

func (x *AddProduct) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ecommerce_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddProduct.ProtoReflect.Descriptor instead.
func (*AddProduct) Descriptor() ([]byte, []int) {
	return file_proto_ecommerce_proto_rawDescGZIP(), []int{1}
}

func (x *AddProduct) GetProductID() int32 {
	if x != nil {
		return x.ProductID
	}
	return 0
}

func (x *AddProduct) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *AddProduct) GetUserID() int32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

type OrderedProduct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Quantity int32 `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *OrderedProduct) Reset() {
	*x = OrderedProduct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ecommerce_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderedProduct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderedProduct) ProtoMessage() {}

func (x *OrderedProduct) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ecommerce_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderedProduct.ProtoReflect.Descriptor instead.
func (*OrderedProduct) Descriptor() ([]byte, []int) {
	return file_proto_ecommerce_proto_rawDescGZIP(), []int{2}
}

func (x *OrderedProduct) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OrderedProduct) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type OrderedProductList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Products []*OrderedProduct `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *OrderedProductList) Reset() {
	*x = OrderedProductList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ecommerce_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderedProductList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderedProductList) ProtoMessage() {}

func (x *OrderedProductList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ecommerce_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderedProductList.ProtoReflect.Descriptor instead.
func (*OrderedProductList) Descriptor() ([]byte, []int) {
	return file_proto_ecommerce_proto_rawDescGZIP(), []int{3}
}

func (x *OrderedProductList) GetProducts() []*OrderedProduct {
	if x != nil {
		return x.Products
	}
	return nil
}

type Category struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Label string `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *Category) Reset() {
	*x = Category{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ecommerce_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Category) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Category) ProtoMessage() {}

func (x *Category) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ecommerce_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Category.ProtoReflect.Descriptor instead.
func (*Category) Descriptor() ([]byte, []int) {
	return file_proto_ecommerce_proto_rawDescGZIP(), []int{4}
}

func (x *Category) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Category) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

type OrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Customerid int32             `protobuf:"varint,1,opt,name=customerid,proto3" json:"customerid,omitempty"`
	Products   []*OrderedProduct `protobuf:"bytes,2,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *OrderRequest) Reset() {
	*x = OrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ecommerce_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderRequest) ProtoMessage() {}

func (x *OrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ecommerce_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderRequest.ProtoReflect.Descriptor instead.
func (*OrderRequest) Descriptor() ([]byte, []int) {
	return file_proto_ecommerce_proto_rawDescGZIP(), []int{5}
}

func (x *OrderRequest) GetCustomerid() int32 {
	if x != nil {
		return x.Customerid
	}
	return 0
}

func (x *OrderRequest) GetProducts() []*OrderedProduct {
	if x != nil {
		return x.Products
	}
	return nil
}

type OrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32             `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Customerid int32             `protobuf:"varint,2,opt,name=customerid,proto3" json:"customerid,omitempty"`
	Products   []*OrderedProduct `protobuf:"bytes,3,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *OrderResponse) Reset() {
	*x = OrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ecommerce_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderResponse) ProtoMessage() {}

func (x *OrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ecommerce_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderResponse.ProtoReflect.Descriptor instead.
func (*OrderResponse) Descriptor() ([]byte, []int) {
	return file_proto_ecommerce_proto_rawDescGZIP(), []int{6}
}

func (x *OrderResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OrderResponse) GetCustomerid() int32 {
	if x != nil {
		return x.Customerid
	}
	return 0
}

func (x *OrderResponse) GetProducts() []*OrderedProduct {
	if x != nil {
		return x.Products
	}
	return nil
}

type OrderID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *OrderID) Reset() {
	*x = OrderID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ecommerce_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderID) ProtoMessage() {}

func (x *OrderID) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ecommerce_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderID.ProtoReflect.Descriptor instead.
func (*OrderID) Descriptor() ([]byte, []int) {
	return file_proto_ecommerce_proto_rawDescGZIP(), []int{7}
}

func (x *OrderID) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool    `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message *string `protobuf:"bytes,2,opt,name=message,proto3,oneof" json:"message,omitempty"`
}

func (x *Reply) Reset() {
	*x = Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ecommerce_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reply) ProtoMessage() {}

func (x *Reply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ecommerce_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reply.ProtoReflect.Descriptor instead.
func (*Reply) Descriptor() ([]byte, []int) {
	return file_proto_ecommerce_proto_rawDescGZIP(), []int{8}
}

func (x *Reply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *Reply) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

type UserId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UserId) Reset() {
	*x = UserId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ecommerce_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserId) ProtoMessage() {}

func (x *UserId) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ecommerce_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserId.ProtoReflect.Descriptor instead.
func (*UserId) Descriptor() ([]byte, []int) {
	return file_proto_ecommerce_proto_rawDescGZIP(), []int{9}
}

func (x *UserId) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ProductList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Products []*Product `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *ProductList) Reset() {
	*x = ProductList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ecommerce_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductList) ProtoMessage() {}

func (x *ProductList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ecommerce_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductList.ProtoReflect.Descriptor instead.
func (*ProductList) Descriptor() ([]byte, []int) {
	return file_proto_ecommerce_proto_rawDescGZIP(), []int{10}
}

func (x *ProductList) GetProducts() []*Product {
	if x != nil {
		return x.Products
	}
	return nil
}

var File_proto_ecommerce_proto protoreflect.FileDescriptor

var file_proto_ecommerce_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x22, 0xa5, 0x01,
	0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x27, 0x0a, 0x0c, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x6c,
	0x69, 0x6e, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0c, 0x64, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x6b, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06,
	0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x48, 0x01, 0x52, 0x06,
	0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x88, 0x01, 0x01, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x64, 0x6f,
	0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x6b, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x77,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x5e, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49,
	0x44, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x3c, 0x0a, 0x0e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x64,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x22, 0x46, 0x0a, 0x12, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x64, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x08, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x22, 0x30, 0x0a, 0x08, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0x60, 0x0a,
	0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a,
	0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x69, 0x64, 0x12, 0x30, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x64, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x22,
	0x71, 0x0a, 0x0d, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x69, 0x64,
	0x12, 0x30, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x65,
	0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x73, 0x22, 0x19, 0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4c, 0x0a,
	0x05, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x12, 0x1d, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x42,
	0x0a, 0x0a, 0x08, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x18, 0x0a, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x38, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x32,
	0xf4, 0x02, 0x0a, 0x10, 0x45, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x12, 0x10, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x41, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x1a, 0x0b, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3d, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x61, 0x72, 0x74, 0x12, 0x0c,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x18, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x61, 0x72, 0x74, 0x12, 0x0c, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x13, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x46, 0x72, 0x6f, 0x6d,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x44, 0x1a, 0x18, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x33,
	0x0a, 0x08, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x12, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x73, 0x46, 0x72, 0x6f, 0x6d, 0x49, 0x44, 0x73, 0x12, 0x18, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x63, 0x74, 0x61, 0x76, 0x69, 0x6f, 0x2d, 0x6c, 0x75, 0x6e,
	0x61, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x61, 0x70, 0x69,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_ecommerce_proto_rawDescOnce sync.Once
	file_proto_ecommerce_proto_rawDescData = file_proto_ecommerce_proto_rawDesc
)

func file_proto_ecommerce_proto_rawDescGZIP() []byte {
	file_proto_ecommerce_proto_rawDescOnce.Do(func() {
		file_proto_ecommerce_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_ecommerce_proto_rawDescData)
	})
	return file_proto_ecommerce_proto_rawDescData
}

var file_proto_ecommerce_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_ecommerce_proto_goTypes = []interface{}{
	(*Product)(nil),            // 0: grpc.Product
	(*AddProduct)(nil),         // 1: grpc.AddProduct
	(*OrderedProduct)(nil),     // 2: grpc.OrderedProduct
	(*OrderedProductList)(nil), // 3: grpc.OrderedProductList
	(*Category)(nil),           // 4: grpc.Category
	(*OrderRequest)(nil),       // 5: grpc.OrderRequest
	(*OrderResponse)(nil),      // 6: grpc.OrderResponse
	(*OrderID)(nil),            // 7: grpc.OrderID
	(*Reply)(nil),              // 8: grpc.Reply
	(*UserId)(nil),             // 9: grpc.UserId
	(*ProductList)(nil),        // 10: grpc.ProductList
}
var file_proto_ecommerce_proto_depIdxs = []int32{
	2,  // 0: grpc.OrderedProductList.products:type_name -> grpc.OrderedProduct
	2,  // 1: grpc.OrderRequest.products:type_name -> grpc.OrderedProduct
	2,  // 2: grpc.OrderResponse.products:type_name -> grpc.OrderedProduct
	0,  // 3: grpc.ProductList.products:type_name -> grpc.Product
	1,  // 4: grpc.ECommerceService.AddProductToCart:input_type -> grpc.AddProduct
	9,  // 5: grpc.ECommerceService.GetProductsFromCart:input_type -> grpc.UserId
	9,  // 6: grpc.ECommerceService.AddOrderFromCart:input_type -> grpc.UserId
	7,  // 7: grpc.ECommerceService.GetProductsFromOrder:input_type -> grpc.OrderID
	5,  // 8: grpc.ECommerceService.AddOrder:input_type -> grpc.OrderRequest
	3,  // 9: grpc.ECommerceService.GetProductsFromIDs:input_type -> grpc.OrderedProductList
	8,  // 10: grpc.ECommerceService.AddProductToCart:output_type -> grpc.Reply
	3,  // 11: grpc.ECommerceService.GetProductsFromCart:output_type -> grpc.OrderedProductList
	6,  // 12: grpc.ECommerceService.AddOrderFromCart:output_type -> grpc.OrderResponse
	3,  // 13: grpc.ECommerceService.GetProductsFromOrder:output_type -> grpc.OrderedProductList
	6,  // 14: grpc.ECommerceService.AddOrder:output_type -> grpc.OrderResponse
	10, // 15: grpc.ECommerceService.GetProductsFromIDs:output_type -> grpc.ProductList
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_proto_ecommerce_proto_init() }
func file_proto_ecommerce_proto_init() {
	if File_proto_ecommerce_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_ecommerce_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
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
		file_proto_ecommerce_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddProduct); i {
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
		file_proto_ecommerce_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderedProduct); i {
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
		file_proto_ecommerce_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderedProductList); i {
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
		file_proto_ecommerce_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Category); i {
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
		file_proto_ecommerce_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderRequest); i {
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
		file_proto_ecommerce_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderResponse); i {
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
		file_proto_ecommerce_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderID); i {
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
		file_proto_ecommerce_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reply); i {
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
		file_proto_ecommerce_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserId); i {
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
		file_proto_ecommerce_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductList); i {
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
	file_proto_ecommerce_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_proto_ecommerce_proto_msgTypes[8].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_ecommerce_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_ecommerce_proto_goTypes,
		DependencyIndexes: file_proto_ecommerce_proto_depIdxs,
		MessageInfos:      file_proto_ecommerce_proto_msgTypes,
	}.Build()
	File_proto_ecommerce_proto = out.File
	file_proto_ecommerce_proto_rawDesc = nil
	file_proto_ecommerce_proto_goTypes = nil
	file_proto_ecommerce_proto_depIdxs = nil
}