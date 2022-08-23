// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: orderingpb/api.proto

package orderingpb

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

type CreateOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items      []*Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	CustomerId string  `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	PaymentId  string  `protobuf:"bytes,3,opt,name=payment_id,json=paymentId,proto3" json:"payment_id,omitempty"`
}

func (x *CreateOrderRequest) Reset() {
	*x = CreateOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderingpb_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderRequest) ProtoMessage() {}

func (x *CreateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orderingpb_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderRequest.ProtoReflect.Descriptor instead.
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return file_orderingpb_api_proto_rawDescGZIP(), []int{0}
}

func (x *CreateOrderRequest) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *CreateOrderRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *CreateOrderRequest) GetPaymentId() string {
	if x != nil {
		return x.PaymentId
	}
	return ""
}

type CreateOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateOrderResponse) Reset() {
	*x = CreateOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderingpb_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderResponse) ProtoMessage() {}

func (x *CreateOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orderingpb_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderResponse.ProtoReflect.Descriptor instead.
func (*CreateOrderResponse) Descriptor() ([]byte, []int) {
	return file_orderingpb_api_proto_rawDescGZIP(), []int{1}
}

func (x *CreateOrderResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetOrderRequest) Reset() {
	*x = GetOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderingpb_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderRequest) ProtoMessage() {}

func (x *GetOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orderingpb_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderRequest.ProtoReflect.Descriptor instead.
func (*GetOrderRequest) Descriptor() ([]byte, []int) {
	return file_orderingpb_api_proto_rawDescGZIP(), []int{2}
}

func (x *GetOrderRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Order *Order `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *GetOrderResponse) Reset() {
	*x = GetOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderingpb_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderResponse) ProtoMessage() {}

func (x *GetOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orderingpb_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderResponse.ProtoReflect.Descriptor instead.
func (*GetOrderResponse) Descriptor() ([]byte, []int) {
	return file_orderingpb_api_proto_rawDescGZIP(), []int{3}
}

func (x *GetOrderResponse) GetOrder() *Order {
	if x != nil {
		return x.Order
	}
	return nil
}

type CancelOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CancelOrderRequest) Reset() {
	*x = CancelOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderingpb_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelOrderRequest) ProtoMessage() {}

func (x *CancelOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orderingpb_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelOrderRequest.ProtoReflect.Descriptor instead.
func (*CancelOrderRequest) Descriptor() ([]byte, []int) {
	return file_orderingpb_api_proto_rawDescGZIP(), []int{4}
}

func (x *CancelOrderRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CancelOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CancelOrderResponse) Reset() {
	*x = CancelOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderingpb_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelOrderResponse) ProtoMessage() {}

func (x *CancelOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orderingpb_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelOrderResponse.ProtoReflect.Descriptor instead.
func (*CancelOrderResponse) Descriptor() ([]byte, []int) {
	return file_orderingpb_api_proto_rawDescGZIP(), []int{5}
}

type ReadyOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ReadyOrderRequest) Reset() {
	*x = ReadyOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderingpb_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadyOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadyOrderRequest) ProtoMessage() {}

func (x *ReadyOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orderingpb_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadyOrderRequest.ProtoReflect.Descriptor instead.
func (*ReadyOrderRequest) Descriptor() ([]byte, []int) {
	return file_orderingpb_api_proto_rawDescGZIP(), []int{6}
}

func (x *ReadyOrderRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ReadyOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReadyOrderResponse) Reset() {
	*x = ReadyOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderingpb_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadyOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadyOrderResponse) ProtoMessage() {}

func (x *ReadyOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orderingpb_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadyOrderResponse.ProtoReflect.Descriptor instead.
func (*ReadyOrderResponse) Descriptor() ([]byte, []int) {
	return file_orderingpb_api_proto_rawDescGZIP(), []int{7}
}

type CompleteOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	InvoiceId string `protobuf:"bytes,2,opt,name=invoice_id,json=invoiceId,proto3" json:"invoice_id,omitempty"`
}

func (x *CompleteOrderRequest) Reset() {
	*x = CompleteOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderingpb_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteOrderRequest) ProtoMessage() {}

func (x *CompleteOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orderingpb_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteOrderRequest.ProtoReflect.Descriptor instead.
func (*CompleteOrderRequest) Descriptor() ([]byte, []int) {
	return file_orderingpb_api_proto_rawDescGZIP(), []int{8}
}

func (x *CompleteOrderRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CompleteOrderRequest) GetInvoiceId() string {
	if x != nil {
		return x.InvoiceId
	}
	return ""
}

type CompleteOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CompleteOrderResponse) Reset() {
	*x = CompleteOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderingpb_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteOrderResponse) ProtoMessage() {}

func (x *CompleteOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orderingpb_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteOrderResponse.ProtoReflect.Descriptor instead.
func (*CompleteOrderResponse) Descriptor() ([]byte, []int) {
	return file_orderingpb_api_proto_rawDescGZIP(), []int{9}
}

var File_orderingpb_api_proto protoreflect.FileDescriptor

var file_orderingpb_api_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x2f, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67,
	0x70, 0x62, 0x1a, 0x19, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7c, 0x0a,
	0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x2e,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x25, 0x0a, 0x13, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x21, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3b, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x05, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x69, 0x6e, 0x67, 0x70, 0x62, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x22, 0x24, 0x0a, 0x12, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x43, 0x61, 0x6e, 0x63,
	0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x23, 0x0a, 0x11, 0x52, 0x65, 0x61, 0x64, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x52, 0x65, 0x61, 0x64, 0x79, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x45, 0x0a, 0x14, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x49,
	0x64, 0x22, 0x17, 0x0a, 0x15, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xa5, 0x03, 0x0a, 0x0f, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50,
	0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1e, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x47, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x0b, 0x43, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1e, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x69, 0x6e, 0x67, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x69, 0x6e, 0x67, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0a, 0x52,
	0x65, 0x61, 0x64, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x79, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x69, 0x6e, 0x67, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x0d, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x20, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x94, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x69, 0x6e, 0x67, 0x70, 0x62, 0x42, 0x08, 0x41, 0x70, 0x69, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x30, 0x65, 0x64, 0x61, 0x2d, 0x69, 0x6e, 0x2d, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x2f, 0x63, 0x68, 0x34, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e,
	0x67, 0x70, 0x62, 0xa2, 0x02, 0x03, 0x4f, 0x58, 0x58, 0xaa, 0x02, 0x0a, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x69, 0x6e, 0x67, 0x70, 0x62, 0xca, 0x02, 0x0a, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e,
	0x67, 0x70, 0x62, 0xe2, 0x02, 0x16, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x70, 0x62,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_orderingpb_api_proto_rawDescOnce sync.Once
	file_orderingpb_api_proto_rawDescData = file_orderingpb_api_proto_rawDesc
)

func file_orderingpb_api_proto_rawDescGZIP() []byte {
	file_orderingpb_api_proto_rawDescOnce.Do(func() {
		file_orderingpb_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_orderingpb_api_proto_rawDescData)
	})
	return file_orderingpb_api_proto_rawDescData
}

var file_orderingpb_api_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_orderingpb_api_proto_goTypes = []interface{}{
	(*CreateOrderRequest)(nil),    // 0: orderingpb.CreateOrderRequest
	(*CreateOrderResponse)(nil),   // 1: orderingpb.CreateOrderResponse
	(*GetOrderRequest)(nil),       // 2: orderingpb.GetOrderRequest
	(*GetOrderResponse)(nil),      // 3: orderingpb.GetOrderResponse
	(*CancelOrderRequest)(nil),    // 4: orderingpb.CancelOrderRequest
	(*CancelOrderResponse)(nil),   // 5: orderingpb.CancelOrderResponse
	(*ReadyOrderRequest)(nil),     // 6: orderingpb.ReadyOrderRequest
	(*ReadyOrderResponse)(nil),    // 7: orderingpb.ReadyOrderResponse
	(*CompleteOrderRequest)(nil),  // 8: orderingpb.CompleteOrderRequest
	(*CompleteOrderResponse)(nil), // 9: orderingpb.CompleteOrderResponse
	(*Item)(nil),                  // 10: orderingpb.Item
	(*Order)(nil),                 // 11: orderingpb.Order
}
var file_orderingpb_api_proto_depIdxs = []int32{
	10, // 0: orderingpb.CreateOrderRequest.items:type_name -> orderingpb.Item
	11, // 1: orderingpb.GetOrderResponse.order:type_name -> orderingpb.Order
	0,  // 2: orderingpb.OrderingService.CreateOrder:input_type -> orderingpb.CreateOrderRequest
	2,  // 3: orderingpb.OrderingService.GetOrder:input_type -> orderingpb.GetOrderRequest
	4,  // 4: orderingpb.OrderingService.CancelOrder:input_type -> orderingpb.CancelOrderRequest
	6,  // 5: orderingpb.OrderingService.ReadyOrder:input_type -> orderingpb.ReadyOrderRequest
	8,  // 6: orderingpb.OrderingService.CompleteOrder:input_type -> orderingpb.CompleteOrderRequest
	1,  // 7: orderingpb.OrderingService.CreateOrder:output_type -> orderingpb.CreateOrderResponse
	3,  // 8: orderingpb.OrderingService.GetOrder:output_type -> orderingpb.GetOrderResponse
	5,  // 9: orderingpb.OrderingService.CancelOrder:output_type -> orderingpb.CancelOrderResponse
	7,  // 10: orderingpb.OrderingService.ReadyOrder:output_type -> orderingpb.ReadyOrderResponse
	9,  // 11: orderingpb.OrderingService.CompleteOrder:output_type -> orderingpb.CompleteOrderResponse
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_orderingpb_api_proto_init() }
func file_orderingpb_api_proto_init() {
	if File_orderingpb_api_proto != nil {
		return
	}
	file_orderingpb_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_orderingpb_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrderRequest); i {
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
		file_orderingpb_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrderResponse); i {
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
		file_orderingpb_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderRequest); i {
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
		file_orderingpb_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderResponse); i {
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
		file_orderingpb_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelOrderRequest); i {
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
		file_orderingpb_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelOrderResponse); i {
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
		file_orderingpb_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadyOrderRequest); i {
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
		file_orderingpb_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadyOrderResponse); i {
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
		file_orderingpb_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompleteOrderRequest); i {
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
		file_orderingpb_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompleteOrderResponse); i {
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
			RawDescriptor: file_orderingpb_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_orderingpb_api_proto_goTypes,
		DependencyIndexes: file_orderingpb_api_proto_depIdxs,
		MessageInfos:      file_orderingpb_api_proto_msgTypes,
	}.Build()
	File_orderingpb_api_proto = out.File
	file_orderingpb_api_proto_rawDesc = nil
	file_orderingpb_api_proto_goTypes = nil
	file_orderingpb_api_proto_depIdxs = nil
}