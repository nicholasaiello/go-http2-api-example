// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: proto/holdings.proto

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

type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	EncodedId   string `protobuf:"bytes,2,opt,name=encoded_id,json=encodedId,proto3" json:"encoded_id,omitempty"`
	Name        string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	ShortName   string `protobuf:"bytes,4,opt,name=short_name,json=shortName,proto3" json:"short_name,omitempty"`
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Mode        string `protobuf:"bytes,6,opt,name=mode,proto3" json:"mode,omitempty"`
	Status      string `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
	Type        string `protobuf:"bytes,8,opt,name=type,proto3" json:"type,omitempty"`
	Created     string `protobuf:"bytes,9,opt,name=created,proto3" json:"created,omitempty"`
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_holdings_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_proto_holdings_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_proto_holdings_proto_rawDescGZIP(), []int{0}
}

func (x *Account) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Account) GetEncodedId() string {
	if x != nil {
		return x.EncodedId
	}
	return ""
}

func (x *Account) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Account) GetShortName() string {
	if x != nil {
		return x.ShortName
	}
	return ""
}

func (x *Account) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Account) GetMode() string {
	if x != nil {
		return x.Mode
	}
	return ""
}

func (x *Account) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Account) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Account) GetCreated() string {
	if x != nil {
		return x.Created
	}
	return ""
}

type Lot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Quantity     int32   `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	PricePaid    float32 `protobuf:"fixed32,3,opt,name=price_paid,json=pricePaid,proto3" json:"price_paid,omitempty"`
	TotalCost    float32 `protobuf:"fixed32,4,opt,name=total_cost,json=totalCost,proto3" json:"total_cost,omitempty"`
	AcquiredDate string  `protobuf:"bytes,5,opt,name=acquired_date,json=acquiredDate,proto3" json:"acquired_date,omitempty"`
}

func (x *Lot) Reset() {
	*x = Lot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_holdings_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Lot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Lot) ProtoMessage() {}

func (x *Lot) ProtoReflect() protoreflect.Message {
	mi := &file_proto_holdings_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Lot.ProtoReflect.Descriptor instead.
func (*Lot) Descriptor() ([]byte, []int) {
	return file_proto_holdings_proto_rawDescGZIP(), []int{1}
}

func (x *Lot) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Lot) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *Lot) GetPricePaid() float32 {
	if x != nil {
		return x.PricePaid
	}
	return 0
}

func (x *Lot) GetTotalCost() float32 {
	if x != nil {
		return x.TotalCost
	}
	return 0
}

func (x *Lot) GetAcquiredDate() string {
	if x != nil {
		return x.AcquiredDate
	}
	return ""
}

type Holding struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Symbol        string  `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Lots          []*Lot  `protobuf:"bytes,3,rep,name=lots,proto3" json:"lots,omitempty"`
	Quantity      int32   `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
	PricePaid     float32 `protobuf:"fixed32,5,opt,name=price_paid,json=pricePaid,proto3" json:"price_paid,omitempty"`
	TotalCost     float32 `protobuf:"fixed32,6,opt,name=total_cost,json=totalCost,proto3" json:"total_cost,omitempty"`
	LastTradeDate string  `protobuf:"bytes,7,opt,name=last_trade_date,json=lastTradeDate,proto3" json:"last_trade_date,omitempty"`
	AcquiredDate  string  `protobuf:"bytes,8,opt,name=acquired_date,json=acquiredDate,proto3" json:"acquired_date,omitempty"`
}

func (x *Holding) Reset() {
	*x = Holding{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_holdings_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Holding) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Holding) ProtoMessage() {}

func (x *Holding) ProtoReflect() protoreflect.Message {
	mi := &file_proto_holdings_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Holding.ProtoReflect.Descriptor instead.
func (*Holding) Descriptor() ([]byte, []int) {
	return file_proto_holdings_proto_rawDescGZIP(), []int{2}
}

func (x *Holding) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Holding) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *Holding) GetLots() []*Lot {
	if x != nil {
		return x.Lots
	}
	return nil
}

func (x *Holding) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *Holding) GetPricePaid() float32 {
	if x != nil {
		return x.PricePaid
	}
	return 0
}

func (x *Holding) GetTotalCost() float32 {
	if x != nil {
		return x.TotalCost
	}
	return 0
}

func (x *Holding) GetLastTradeDate() string {
	if x != nil {
		return x.LastTradeDate
	}
	return ""
}

func (x *Holding) GetAcquiredDate() string {
	if x != nil {
		return x.AcquiredDate
	}
	return ""
}

type Holdings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                  string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Positions           []*Holding `protobuf:"bytes,2,rep,name=positions,proto3" json:"positions,omitempty"`
	TotalPositionsCount int32      `protobuf:"varint,3,opt,name=total_positions_count,json=totalPositionsCount,proto3" json:"total_positions_count,omitempty"`
	ViewName            string     `protobuf:"bytes,4,opt,name=view_name,json=viewName,proto3" json:"view_name,omitempty"`
	LastUpdated         string     `protobuf:"bytes,5,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
}

func (x *Holdings) Reset() {
	*x = Holdings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_holdings_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Holdings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Holdings) ProtoMessage() {}

func (x *Holdings) ProtoReflect() protoreflect.Message {
	mi := &file_proto_holdings_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Holdings.ProtoReflect.Descriptor instead.
func (*Holdings) Descriptor() ([]byte, []int) {
	return file_proto_holdings_proto_rawDescGZIP(), []int{3}
}

func (x *Holdings) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Holdings) GetPositions() []*Holding {
	if x != nil {
		return x.Positions
	}
	return nil
}

func (x *Holdings) GetTotalPositionsCount() int32 {
	if x != nil {
		return x.TotalPositionsCount
	}
	return 0
}

func (x *Holdings) GetViewName() string {
	if x != nil {
		return x.ViewName
	}
	return ""
}

func (x *Holdings) GetLastUpdated() string {
	if x != nil {
		return x.LastUpdated
	}
	return ""
}

var File_proto_holdings_proto protoreflect.FileDescriptor

var file_proto_holdings_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x6f, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe7, 0x01,
	0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x6e, 0x63,
	0x6f, 0x64, 0x65, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65,
	0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x6f, 0x64,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x22, 0x94, 0x01, 0x0a, 0x03, 0x4c, 0x6f, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x61, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x09, 0x70, 0x72, 0x69, 0x63, 0x65, 0x50, 0x61, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x63, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x61, 0x63, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x22, 0xf8,
	0x01, 0x0a, 0x07, 0x48, 0x6f, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79,
	0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62,
	0x6f, 0x6c, 0x12, 0x1e, 0x0a, 0x04, 0x6c, 0x6f, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x74, 0x52, 0x04, 0x6c, 0x6f,
	0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x61, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x09, 0x70, 0x72, 0x69, 0x63, 0x65, 0x50, 0x61, 0x69, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0f,
	0x6c, 0x61, 0x73, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x63, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x63, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x22, 0xbc, 0x01, 0x0a, 0x08, 0x48, 0x6f,
	0x6c, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2c, 0x0a, 0x09, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x48, 0x6f, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x09, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x32, 0x0a, 0x15, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x13, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x76, 0x69, 0x65, 0x77,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x76, 0x69, 0x65,
	0x77, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6c, 0x61, 0x73,
	0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x32, 0xa7, 0x01, 0x0a, 0x0f, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x48, 0x6f, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x30, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x48, 0x6f, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x0e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x6f, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x00, 0x12, 0x2e,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x48, 0x6f, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x6f, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x1a, 0x0e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x6f, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x22, 0x00, 0x12, 0x32,
	0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x6f, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x0e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x0e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x6f, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x22, 0x00,
	0x30, 0x01, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6e, 0x69, 0x63, 0x68, 0x6f, 0x6c, 0x61, 0x73, 0x61, 0x69, 0x65, 0x6c, 0x6c, 0x6f, 0x2f,
	0x67, 0x6f, 0x2d, 0x68, 0x74, 0x74, 0x70, 0x32, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_holdings_proto_rawDescOnce sync.Once
	file_proto_holdings_proto_rawDescData = file_proto_holdings_proto_rawDesc
)

func file_proto_holdings_proto_rawDescGZIP() []byte {
	file_proto_holdings_proto_rawDescOnce.Do(func() {
		file_proto_holdings_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_holdings_proto_rawDescData)
	})
	return file_proto_holdings_proto_rawDescData
}

var file_proto_holdings_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_holdings_proto_goTypes = []interface{}{
	(*Account)(nil),  // 0: proto.Account
	(*Lot)(nil),      // 1: proto.Lot
	(*Holding)(nil),  // 2: proto.Holding
	(*Holdings)(nil), // 3: proto.Holdings
}
var file_proto_holdings_proto_depIdxs = []int32{
	1, // 0: proto.Holding.lots:type_name -> proto.Lot
	2, // 1: proto.Holdings.positions:type_name -> proto.Holding
	0, // 2: proto.AccountHoldings.GetHoldings:input_type -> proto.Account
	2, // 3: proto.AccountHoldings.GetHolding:input_type -> proto.Holding
	0, // 4: proto.AccountHoldings.ListHoldings:input_type -> proto.Account
	3, // 5: proto.AccountHoldings.GetHoldings:output_type -> proto.Holdings
	2, // 6: proto.AccountHoldings.GetHolding:output_type -> proto.Holding
	2, // 7: proto.AccountHoldings.ListHoldings:output_type -> proto.Holding
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_holdings_proto_init() }
func file_proto_holdings_proto_init() {
	if File_proto_holdings_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_holdings_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account); i {
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
		file_proto_holdings_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Lot); i {
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
		file_proto_holdings_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Holding); i {
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
		file_proto_holdings_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Holdings); i {
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
			RawDescriptor: file_proto_holdings_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_holdings_proto_goTypes,
		DependencyIndexes: file_proto_holdings_proto_depIdxs,
		MessageInfos:      file_proto_holdings_proto_msgTypes,
	}.Build()
	File_proto_holdings_proto = out.File
	file_proto_holdings_proto_rawDesc = nil
	file_proto_holdings_proto_goTypes = nil
	file_proto_holdings_proto_depIdxs = nil
}
