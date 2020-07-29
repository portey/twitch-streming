// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package service

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetSignUrlRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSignUrlRequest) Reset()         { *m = GetSignUrlRequest{} }
func (m *GetSignUrlRequest) String() string { return proto.CompactTextString(m) }
func (*GetSignUrlRequest) ProtoMessage()    {}
func (*GetSignUrlRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *GetSignUrlRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSignUrlRequest.Unmarshal(m, b)
}
func (m *GetSignUrlRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSignUrlRequest.Marshal(b, m, deterministic)
}
func (m *GetSignUrlRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSignUrlRequest.Merge(m, src)
}
func (m *GetSignUrlRequest) XXX_Size() int {
	return xxx_messageInfo_GetSignUrlRequest.Size(m)
}
func (m *GetSignUrlRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSignUrlRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetSignUrlRequest proto.InternalMessageInfo

type GetSignUrlResponse struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSignUrlResponse) Reset()         { *m = GetSignUrlResponse{} }
func (m *GetSignUrlResponse) String() string { return proto.CompactTextString(m) }
func (*GetSignUrlResponse) ProtoMessage()    {}
func (*GetSignUrlResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *GetSignUrlResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSignUrlResponse.Unmarshal(m, b)
}
func (m *GetSignUrlResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSignUrlResponse.Marshal(b, m, deterministic)
}
func (m *GetSignUrlResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSignUrlResponse.Merge(m, src)
}
func (m *GetSignUrlResponse) XXX_Size() int {
	return xxx_messageInfo_GetSignUrlResponse.Size(m)
}
func (m *GetSignUrlResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSignUrlResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetSignUrlResponse proto.InternalMessageInfo

func (m *GetSignUrlResponse) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type ExchangeTokenRequest struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExchangeTokenRequest) Reset()         { *m = ExchangeTokenRequest{} }
func (m *ExchangeTokenRequest) String() string { return proto.CompactTextString(m) }
func (*ExchangeTokenRequest) ProtoMessage()    {}
func (*ExchangeTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}

func (m *ExchangeTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExchangeTokenRequest.Unmarshal(m, b)
}
func (m *ExchangeTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExchangeTokenRequest.Marshal(b, m, deterministic)
}
func (m *ExchangeTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExchangeTokenRequest.Merge(m, src)
}
func (m *ExchangeTokenRequest) XXX_Size() int {
	return xxx_messageInfo_ExchangeTokenRequest.Size(m)
}
func (m *ExchangeTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExchangeTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExchangeTokenRequest proto.InternalMessageInfo

func (m *ExchangeTokenRequest) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

type ExchangeTokenResponse struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExchangeTokenResponse) Reset()         { *m = ExchangeTokenResponse{} }
func (m *ExchangeTokenResponse) String() string { return proto.CompactTextString(m) }
func (*ExchangeTokenResponse) ProtoMessage()    {}
func (*ExchangeTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{3}
}

func (m *ExchangeTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExchangeTokenResponse.Unmarshal(m, b)
}
func (m *ExchangeTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExchangeTokenResponse.Marshal(b, m, deterministic)
}
func (m *ExchangeTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExchangeTokenResponse.Merge(m, src)
}
func (m *ExchangeTokenResponse) XXX_Size() int {
	return xxx_messageInfo_ExchangeTokenResponse.Size(m)
}
func (m *ExchangeTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExchangeTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExchangeTokenResponse proto.InternalMessageInfo

func (m *ExchangeTokenResponse) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

type SubscribeToStreamerRequest struct {
	StreamerName         string   `protobuf:"bytes,1,opt,name=streamer_name,json=streamerName,proto3" json:"streamer_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscribeToStreamerRequest) Reset()         { *m = SubscribeToStreamerRequest{} }
func (m *SubscribeToStreamerRequest) String() string { return proto.CompactTextString(m) }
func (*SubscribeToStreamerRequest) ProtoMessage()    {}
func (*SubscribeToStreamerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{4}
}

func (m *SubscribeToStreamerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeToStreamerRequest.Unmarshal(m, b)
}
func (m *SubscribeToStreamerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeToStreamerRequest.Marshal(b, m, deterministic)
}
func (m *SubscribeToStreamerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeToStreamerRequest.Merge(m, src)
}
func (m *SubscribeToStreamerRequest) XXX_Size() int {
	return xxx_messageInfo_SubscribeToStreamerRequest.Size(m)
}
func (m *SubscribeToStreamerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeToStreamerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeToStreamerRequest proto.InternalMessageInfo

func (m *SubscribeToStreamerRequest) GetStreamerName() string {
	if m != nil {
		return m.StreamerName
	}
	return ""
}

type SubscribeToStreamerResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscribeToStreamerResponse) Reset()         { *m = SubscribeToStreamerResponse{} }
func (m *SubscribeToStreamerResponse) String() string { return proto.CompactTextString(m) }
func (*SubscribeToStreamerResponse) ProtoMessage()    {}
func (*SubscribeToStreamerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{5}
}

func (m *SubscribeToStreamerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeToStreamerResponse.Unmarshal(m, b)
}
func (m *SubscribeToStreamerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeToStreamerResponse.Marshal(b, m, deterministic)
}
func (m *SubscribeToStreamerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeToStreamerResponse.Merge(m, src)
}
func (m *SubscribeToStreamerResponse) XXX_Size() int {
	return xxx_messageInfo_SubscribeToStreamerResponse.Size(m)
}
func (m *SubscribeToStreamerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeToStreamerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeToStreamerResponse proto.InternalMessageInfo

type GetEventCountsAggregatedByStreamerRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetEventCountsAggregatedByStreamerRequest) Reset() {
	*m = GetEventCountsAggregatedByStreamerRequest{}
}
func (m *GetEventCountsAggregatedByStreamerRequest) String() string {
	return proto.CompactTextString(m)
}
func (*GetEventCountsAggregatedByStreamerRequest) ProtoMessage() {}
func (*GetEventCountsAggregatedByStreamerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{6}
}

func (m *GetEventCountsAggregatedByStreamerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEventCountsAggregatedByStreamerRequest.Unmarshal(m, b)
}
func (m *GetEventCountsAggregatedByStreamerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEventCountsAggregatedByStreamerRequest.Marshal(b, m, deterministic)
}
func (m *GetEventCountsAggregatedByStreamerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEventCountsAggregatedByStreamerRequest.Merge(m, src)
}
func (m *GetEventCountsAggregatedByStreamerRequest) XXX_Size() int {
	return xxx_messageInfo_GetEventCountsAggregatedByStreamerRequest.Size(m)
}
func (m *GetEventCountsAggregatedByStreamerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEventCountsAggregatedByStreamerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetEventCountsAggregatedByStreamerRequest proto.InternalMessageInfo

type GetEventCountsAggregatedByStreamerResponse struct {
	Items                []*GetEventCountsAggregatedByStreamerResponse_Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                           `json:"-"`
	XXX_unrecognized     []byte                                             `json:"-"`
	XXX_sizecache        int32                                              `json:"-"`
}

func (m *GetEventCountsAggregatedByStreamerResponse) Reset() {
	*m = GetEventCountsAggregatedByStreamerResponse{}
}
func (m *GetEventCountsAggregatedByStreamerResponse) String() string {
	return proto.CompactTextString(m)
}
func (*GetEventCountsAggregatedByStreamerResponse) ProtoMessage() {}
func (*GetEventCountsAggregatedByStreamerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{7}
}

func (m *GetEventCountsAggregatedByStreamerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEventCountsAggregatedByStreamerResponse.Unmarshal(m, b)
}
func (m *GetEventCountsAggregatedByStreamerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEventCountsAggregatedByStreamerResponse.Marshal(b, m, deterministic)
}
func (m *GetEventCountsAggregatedByStreamerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEventCountsAggregatedByStreamerResponse.Merge(m, src)
}
func (m *GetEventCountsAggregatedByStreamerResponse) XXX_Size() int {
	return xxx_messageInfo_GetEventCountsAggregatedByStreamerResponse.Size(m)
}
func (m *GetEventCountsAggregatedByStreamerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEventCountsAggregatedByStreamerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetEventCountsAggregatedByStreamerResponse proto.InternalMessageInfo

func (m *GetEventCountsAggregatedByStreamerResponse) GetItems() []*GetEventCountsAggregatedByStreamerResponse_Item {
	if m != nil {
		return m.Items
	}
	return nil
}

type GetEventCountsAggregatedByStreamerResponse_Item struct {
	StreamerName         string   `protobuf:"bytes,1,opt,name=streamer_name,json=streamerName,proto3" json:"streamer_name,omitempty"`
	Count                int32    `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetEventCountsAggregatedByStreamerResponse_Item) Reset() {
	*m = GetEventCountsAggregatedByStreamerResponse_Item{}
}
func (m *GetEventCountsAggregatedByStreamerResponse_Item) String() string {
	return proto.CompactTextString(m)
}
func (*GetEventCountsAggregatedByStreamerResponse_Item) ProtoMessage() {}
func (*GetEventCountsAggregatedByStreamerResponse_Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{7, 0}
}

func (m *GetEventCountsAggregatedByStreamerResponse_Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEventCountsAggregatedByStreamerResponse_Item.Unmarshal(m, b)
}
func (m *GetEventCountsAggregatedByStreamerResponse_Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEventCountsAggregatedByStreamerResponse_Item.Marshal(b, m, deterministic)
}
func (m *GetEventCountsAggregatedByStreamerResponse_Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEventCountsAggregatedByStreamerResponse_Item.Merge(m, src)
}
func (m *GetEventCountsAggregatedByStreamerResponse_Item) XXX_Size() int {
	return xxx_messageInfo_GetEventCountsAggregatedByStreamerResponse_Item.Size(m)
}
func (m *GetEventCountsAggregatedByStreamerResponse_Item) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEventCountsAggregatedByStreamerResponse_Item.DiscardUnknown(m)
}

var xxx_messageInfo_GetEventCountsAggregatedByStreamerResponse_Item proto.InternalMessageInfo

func (m *GetEventCountsAggregatedByStreamerResponse_Item) GetStreamerName() string {
	if m != nil {
		return m.StreamerName
	}
	return ""
}

func (m *GetEventCountsAggregatedByStreamerResponse_Item) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type GetEventCountsAggregatedByStreamerAndTypeRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetEventCountsAggregatedByStreamerAndTypeRequest) Reset() {
	*m = GetEventCountsAggregatedByStreamerAndTypeRequest{}
}
func (m *GetEventCountsAggregatedByStreamerAndTypeRequest) String() string {
	return proto.CompactTextString(m)
}
func (*GetEventCountsAggregatedByStreamerAndTypeRequest) ProtoMessage() {}
func (*GetEventCountsAggregatedByStreamerAndTypeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{8}
}

func (m *GetEventCountsAggregatedByStreamerAndTypeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEventCountsAggregatedByStreamerAndTypeRequest.Unmarshal(m, b)
}
func (m *GetEventCountsAggregatedByStreamerAndTypeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEventCountsAggregatedByStreamerAndTypeRequest.Marshal(b, m, deterministic)
}
func (m *GetEventCountsAggregatedByStreamerAndTypeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEventCountsAggregatedByStreamerAndTypeRequest.Merge(m, src)
}
func (m *GetEventCountsAggregatedByStreamerAndTypeRequest) XXX_Size() int {
	return xxx_messageInfo_GetEventCountsAggregatedByStreamerAndTypeRequest.Size(m)
}
func (m *GetEventCountsAggregatedByStreamerAndTypeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEventCountsAggregatedByStreamerAndTypeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetEventCountsAggregatedByStreamerAndTypeRequest proto.InternalMessageInfo

type GetEventCountsAggregatedByStreamerAndTypeResponse struct {
	Items                []*GetEventCountsAggregatedByStreamerAndTypeResponse_Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                  `json:"-"`
	XXX_unrecognized     []byte                                                    `json:"-"`
	XXX_sizecache        int32                                                     `json:"-"`
}

func (m *GetEventCountsAggregatedByStreamerAndTypeResponse) Reset() {
	*m = GetEventCountsAggregatedByStreamerAndTypeResponse{}
}
func (m *GetEventCountsAggregatedByStreamerAndTypeResponse) String() string {
	return proto.CompactTextString(m)
}
func (*GetEventCountsAggregatedByStreamerAndTypeResponse) ProtoMessage() {}
func (*GetEventCountsAggregatedByStreamerAndTypeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{9}
}

func (m *GetEventCountsAggregatedByStreamerAndTypeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEventCountsAggregatedByStreamerAndTypeResponse.Unmarshal(m, b)
}
func (m *GetEventCountsAggregatedByStreamerAndTypeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEventCountsAggregatedByStreamerAndTypeResponse.Marshal(b, m, deterministic)
}
func (m *GetEventCountsAggregatedByStreamerAndTypeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEventCountsAggregatedByStreamerAndTypeResponse.Merge(m, src)
}
func (m *GetEventCountsAggregatedByStreamerAndTypeResponse) XXX_Size() int {
	return xxx_messageInfo_GetEventCountsAggregatedByStreamerAndTypeResponse.Size(m)
}
func (m *GetEventCountsAggregatedByStreamerAndTypeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEventCountsAggregatedByStreamerAndTypeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetEventCountsAggregatedByStreamerAndTypeResponse proto.InternalMessageInfo

func (m *GetEventCountsAggregatedByStreamerAndTypeResponse) GetItems() []*GetEventCountsAggregatedByStreamerAndTypeResponse_Item {
	if m != nil {
		return m.Items
	}
	return nil
}

type GetEventCountsAggregatedByStreamerAndTypeResponse_Item struct {
	StreamerName         string   `protobuf:"bytes,1,opt,name=streamer_name,json=streamerName,proto3" json:"streamer_name,omitempty"`
	EventType            string   `protobuf:"bytes,2,opt,name=event_type,json=eventType,proto3" json:"event_type,omitempty"`
	Count                int32    `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetEventCountsAggregatedByStreamerAndTypeResponse_Item) Reset() {
	*m = GetEventCountsAggregatedByStreamerAndTypeResponse_Item{}
}
func (m *GetEventCountsAggregatedByStreamerAndTypeResponse_Item) String() string {
	return proto.CompactTextString(m)
}
func (*GetEventCountsAggregatedByStreamerAndTypeResponse_Item) ProtoMessage() {}
func (*GetEventCountsAggregatedByStreamerAndTypeResponse_Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{9, 0}
}

func (m *GetEventCountsAggregatedByStreamerAndTypeResponse_Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEventCountsAggregatedByStreamerAndTypeResponse_Item.Unmarshal(m, b)
}
func (m *GetEventCountsAggregatedByStreamerAndTypeResponse_Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEventCountsAggregatedByStreamerAndTypeResponse_Item.Marshal(b, m, deterministic)
}
func (m *GetEventCountsAggregatedByStreamerAndTypeResponse_Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEventCountsAggregatedByStreamerAndTypeResponse_Item.Merge(m, src)
}
func (m *GetEventCountsAggregatedByStreamerAndTypeResponse_Item) XXX_Size() int {
	return xxx_messageInfo_GetEventCountsAggregatedByStreamerAndTypeResponse_Item.Size(m)
}
func (m *GetEventCountsAggregatedByStreamerAndTypeResponse_Item) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEventCountsAggregatedByStreamerAndTypeResponse_Item.DiscardUnknown(m)
}

var xxx_messageInfo_GetEventCountsAggregatedByStreamerAndTypeResponse_Item proto.InternalMessageInfo

func (m *GetEventCountsAggregatedByStreamerAndTypeResponse_Item) GetStreamerName() string {
	if m != nil {
		return m.StreamerName
	}
	return ""
}

func (m *GetEventCountsAggregatedByStreamerAndTypeResponse_Item) GetEventType() string {
	if m != nil {
		return m.EventType
	}
	return ""
}

func (m *GetEventCountsAggregatedByStreamerAndTypeResponse_Item) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterType((*GetSignUrlRequest)(nil), "user_service_api.GetSignUrlRequest")
	proto.RegisterType((*GetSignUrlResponse)(nil), "user_service_api.GetSignUrlResponse")
	proto.RegisterType((*ExchangeTokenRequest)(nil), "user_service_api.ExchangeTokenRequest")
	proto.RegisterType((*ExchangeTokenResponse)(nil), "user_service_api.ExchangeTokenResponse")
	proto.RegisterType((*SubscribeToStreamerRequest)(nil), "user_service_api.SubscribeToStreamerRequest")
	proto.RegisterType((*SubscribeToStreamerResponse)(nil), "user_service_api.SubscribeToStreamerResponse")
	proto.RegisterType((*GetEventCountsAggregatedByStreamerRequest)(nil), "user_service_api.GetEventCountsAggregatedByStreamerRequest")
	proto.RegisterType((*GetEventCountsAggregatedByStreamerResponse)(nil), "user_service_api.GetEventCountsAggregatedByStreamerResponse")
	proto.RegisterType((*GetEventCountsAggregatedByStreamerResponse_Item)(nil), "user_service_api.GetEventCountsAggregatedByStreamerResponse.Item")
	proto.RegisterType((*GetEventCountsAggregatedByStreamerAndTypeRequest)(nil), "user_service_api.GetEventCountsAggregatedByStreamerAndTypeRequest")
	proto.RegisterType((*GetEventCountsAggregatedByStreamerAndTypeResponse)(nil), "user_service_api.GetEventCountsAggregatedByStreamerAndTypeResponse")
	proto.RegisterType((*GetEventCountsAggregatedByStreamerAndTypeResponse_Item)(nil), "user_service_api.GetEventCountsAggregatedByStreamerAndTypeResponse.Item")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 477 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x18, 0xd4, 0x12, 0x82, 0x94, 0x2f, 0x8d, 0x54, 0x36, 0x45, 0x8a, 0x8c, 0x2a, 0x85, 0x2d, 0x2a,
	0xe1, 0xcf, 0x82, 0x70, 0x82, 0x72, 0x71, 0xaa, 0xaa, 0x70, 0xe1, 0xe0, 0x04, 0x55, 0xe2, 0x80,
	0xeb, 0xb8, 0x9f, 0x5c, 0x8b, 0x7a, 0x6d, 0x76, 0xd7, 0x85, 0x3c, 0x0c, 0xe2, 0x25, 0x78, 0x08,
	0x5e, 0x88, 0x3b, 0xb2, 0x77, 0xa3, 0xd6, 0xae, 0x21, 0xee, 0xcf, 0xcd, 0x1e, 0x7d, 0x33, 0xf3,
	0xed, 0xd8, 0xb3, 0xd0, 0x93, 0x28, 0x4e, 0xa3, 0x00, 0xed, 0x54, 0x24, 0x2a, 0xa1, 0xeb, 0x99,
	0x44, 0xe1, 0x19, 0xcc, 0xf3, 0xd3, 0x88, 0xf5, 0xe1, 0xee, 0x3e, 0xaa, 0x69, 0x14, 0xf2, 0x8f,
	0xe2, 0xc4, 0xc5, 0xaf, 0x19, 0x4a, 0xc5, 0xb6, 0x81, 0x9e, 0x07, 0x65, 0x9a, 0x70, 0x89, 0x74,
	0x1d, 0x5a, 0x99, 0x38, 0x19, 0x90, 0x21, 0x19, 0x75, 0xdc, 0xfc, 0x91, 0xbd, 0x86, 0x8d, 0xbd,
	0xef, 0xc1, 0xb1, 0xcf, 0x43, 0x9c, 0x25, 0x5f, 0x90, 0x1b, 0x3e, 0x7d, 0x00, 0x6b, 0x7e, 0x10,
	0xa0, 0x94, 0x9e, 0xca, 0x61, 0x43, 0xe9, 0x6a, 0xac, 0x98, 0x64, 0x6f, 0xe0, 0x5e, 0x85, 0x6a,
	0x5c, 0x1a, 0x70, 0x1d, 0xb0, 0xa6, 0xd9, 0x5c, 0x06, 0x22, 0x9a, 0xe3, 0x2c, 0x99, 0x2a, 0x81,
	0x7e, 0x8c, 0x62, 0x69, 0xbe, 0x05, 0x3d, 0x69, 0x20, 0x8f, 0xfb, 0x31, 0x1a, 0x85, 0xb5, 0x25,
	0xf8, 0xc1, 0x8f, 0x91, 0x6d, 0xc2, 0xfd, 0x5a, 0x09, 0xbd, 0x04, 0x7b, 0x0a, 0x8f, 0xf7, 0x51,
	0xed, 0x9d, 0x22, 0x57, 0xbb, 0x49, 0xc6, 0x95, 0x74, 0xc2, 0x50, 0x60, 0xe8, 0x2b, 0x3c, 0x9a,
	0x2c, 0x2a, 0x86, 0xec, 0x37, 0x81, 0x27, 0x4d, 0xa6, 0xcd, 0x01, 0x0f, 0xa0, 0x1d, 0x29, 0x8c,
	0xe5, 0x80, 0x0c, 0x5b, 0xa3, 0xee, 0xd8, 0xb1, 0xab, 0xdf, 0xc4, 0x6e, 0x2e, 0x66, 0xbf, 0x57,
	0x18, 0xbb, 0x5a, 0xcf, 0x72, 0xe0, 0x76, 0xfe, 0xda, 0x28, 0x00, 0xba, 0x01, 0xed, 0x20, 0x97,
	0x1f, 0xdc, 0x1a, 0x92, 0x51, 0xdb, 0xd5, 0x2f, 0x6c, 0x0c, 0x2f, 0x56, 0x9b, 0x3b, 0xfc, 0x68,
	0xb6, 0x48, 0x71, 0x79, 0xfc, 0x3f, 0x04, 0x5e, 0x5e, 0x82, 0x64, 0x52, 0xf8, 0x5c, 0x4e, 0xe1,
	0xdd, 0x55, 0x52, 0xa8, 0x68, 0x96, 0xc2, 0x38, 0xbc, 0x4c, 0x18, 0x9b, 0x00, 0x98, 0x5b, 0x79,
	0x6a, 0x91, 0x62, 0x91, 0x48, 0xc7, 0xed, 0x14, 0x48, 0xae, 0x7f, 0x96, 0x55, 0xeb, 0x5c, 0x56,
	0xe3, 0x1f, 0x6d, 0xe8, 0xcf, 0xbe, 0x45, 0x2a, 0x38, 0xd6, 0xfb, 0x4c, 0xf5, 0xee, 0xf4, 0x00,
	0xe0, 0xac, 0x3c, 0x74, 0xab, 0xf6, 0x60, 0xe5, 0xbe, 0x59, 0x0f, 0xff, 0x3f, 0x64, 0x22, 0x3b,
	0x84, 0x5e, 0xa9, 0x32, 0x74, 0xfb, 0x22, 0xad, 0xae, 0x8e, 0xd6, 0xa3, 0x95, 0x73, 0xc6, 0x41,
	0x40, 0xbf, 0xa6, 0x15, 0xf4, 0xd9, 0x45, 0xfe, 0xbf, 0xfb, 0x67, 0x3d, 0x6f, 0x38, 0x6d, 0x3c,
	0x7f, 0x12, 0x60, 0xab, 0x3f, 0x35, 0xdd, 0xb9, 0x5a, 0x4d, 0xf4, 0x4a, 0x6f, 0xaf, 0xd3, 0x31,
	0xfa, 0x8b, 0x34, 0xb9, 0x0d, 0xcc, 0xcf, 0x48, 0x27, 0xd7, 0xfa, 0x93, 0xf5, 0xbe, 0xbb, 0x37,
	0xd0, 0x86, 0x49, 0xf7, 0x53, 0xc7, 0xde, 0x31, 0x12, 0xf3, 0x3b, 0xc5, 0xfd, 0xff, 0xea, 0x6f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x7e, 0x50, 0xe1, 0x31, 0x10, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TwitchStreamServiceClient is the client API for TwitchStreamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TwitchStreamServiceClient interface {
	GetSignUrl(ctx context.Context, in *GetSignUrlRequest, opts ...grpc.CallOption) (*GetSignUrlResponse, error)
	ExchangeToken(ctx context.Context, in *ExchangeTokenRequest, opts ...grpc.CallOption) (*ExchangeTokenResponse, error)
	SubscribeToStreamer(ctx context.Context, in *SubscribeToStreamerRequest, opts ...grpc.CallOption) (*SubscribeToStreamerResponse, error)
	GetEventCountsAggregatedByStreamer(ctx context.Context, in *GetEventCountsAggregatedByStreamerRequest, opts ...grpc.CallOption) (*GetEventCountsAggregatedByStreamerResponse, error)
	GetEventCountsAggregatedByStreamerAndType(ctx context.Context, in *GetEventCountsAggregatedByStreamerAndTypeRequest, opts ...grpc.CallOption) (*GetEventCountsAggregatedByStreamerAndTypeResponse, error)
}

type twitchStreamServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTwitchStreamServiceClient(cc grpc.ClientConnInterface) TwitchStreamServiceClient {
	return &twitchStreamServiceClient{cc}
}

func (c *twitchStreamServiceClient) GetSignUrl(ctx context.Context, in *GetSignUrlRequest, opts ...grpc.CallOption) (*GetSignUrlResponse, error) {
	out := new(GetSignUrlResponse)
	err := c.cc.Invoke(ctx, "/user_service_api.TwitchStreamService/GetSignUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *twitchStreamServiceClient) ExchangeToken(ctx context.Context, in *ExchangeTokenRequest, opts ...grpc.CallOption) (*ExchangeTokenResponse, error) {
	out := new(ExchangeTokenResponse)
	err := c.cc.Invoke(ctx, "/user_service_api.TwitchStreamService/ExchangeToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *twitchStreamServiceClient) SubscribeToStreamer(ctx context.Context, in *SubscribeToStreamerRequest, opts ...grpc.CallOption) (*SubscribeToStreamerResponse, error) {
	out := new(SubscribeToStreamerResponse)
	err := c.cc.Invoke(ctx, "/user_service_api.TwitchStreamService/SubscribeToStreamer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *twitchStreamServiceClient) GetEventCountsAggregatedByStreamer(ctx context.Context, in *GetEventCountsAggregatedByStreamerRequest, opts ...grpc.CallOption) (*GetEventCountsAggregatedByStreamerResponse, error) {
	out := new(GetEventCountsAggregatedByStreamerResponse)
	err := c.cc.Invoke(ctx, "/user_service_api.TwitchStreamService/GetEventCountsAggregatedByStreamer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *twitchStreamServiceClient) GetEventCountsAggregatedByStreamerAndType(ctx context.Context, in *GetEventCountsAggregatedByStreamerAndTypeRequest, opts ...grpc.CallOption) (*GetEventCountsAggregatedByStreamerAndTypeResponse, error) {
	out := new(GetEventCountsAggregatedByStreamerAndTypeResponse)
	err := c.cc.Invoke(ctx, "/user_service_api.TwitchStreamService/GetEventCountsAggregatedByStreamerAndType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TwitchStreamServiceServer is the server API for TwitchStreamService service.
type TwitchStreamServiceServer interface {
	GetSignUrl(context.Context, *GetSignUrlRequest) (*GetSignUrlResponse, error)
	ExchangeToken(context.Context, *ExchangeTokenRequest) (*ExchangeTokenResponse, error)
	SubscribeToStreamer(context.Context, *SubscribeToStreamerRequest) (*SubscribeToStreamerResponse, error)
	GetEventCountsAggregatedByStreamer(context.Context, *GetEventCountsAggregatedByStreamerRequest) (*GetEventCountsAggregatedByStreamerResponse, error)
	GetEventCountsAggregatedByStreamerAndType(context.Context, *GetEventCountsAggregatedByStreamerAndTypeRequest) (*GetEventCountsAggregatedByStreamerAndTypeResponse, error)
}

// UnimplementedTwitchStreamServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTwitchStreamServiceServer struct {
}

func (*UnimplementedTwitchStreamServiceServer) GetSignUrl(ctx context.Context, req *GetSignUrlRequest) (*GetSignUrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSignUrl not implemented")
}
func (*UnimplementedTwitchStreamServiceServer) ExchangeToken(ctx context.Context, req *ExchangeTokenRequest) (*ExchangeTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExchangeToken not implemented")
}
func (*UnimplementedTwitchStreamServiceServer) SubscribeToStreamer(ctx context.Context, req *SubscribeToStreamerRequest) (*SubscribeToStreamerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubscribeToStreamer not implemented")
}
func (*UnimplementedTwitchStreamServiceServer) GetEventCountsAggregatedByStreamer(ctx context.Context, req *GetEventCountsAggregatedByStreamerRequest) (*GetEventCountsAggregatedByStreamerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventCountsAggregatedByStreamer not implemented")
}
func (*UnimplementedTwitchStreamServiceServer) GetEventCountsAggregatedByStreamerAndType(ctx context.Context, req *GetEventCountsAggregatedByStreamerAndTypeRequest) (*GetEventCountsAggregatedByStreamerAndTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventCountsAggregatedByStreamerAndType not implemented")
}

func RegisterTwitchStreamServiceServer(s *grpc.Server, srv TwitchStreamServiceServer) {
	s.RegisterService(&_TwitchStreamService_serviceDesc, srv)
}

func _TwitchStreamService_GetSignUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSignUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwitchStreamServiceServer).GetSignUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_service_api.TwitchStreamService/GetSignUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwitchStreamServiceServer).GetSignUrl(ctx, req.(*GetSignUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TwitchStreamService_ExchangeToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExchangeTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwitchStreamServiceServer).ExchangeToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_service_api.TwitchStreamService/ExchangeToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwitchStreamServiceServer).ExchangeToken(ctx, req.(*ExchangeTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TwitchStreamService_SubscribeToStreamer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeToStreamerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwitchStreamServiceServer).SubscribeToStreamer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_service_api.TwitchStreamService/SubscribeToStreamer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwitchStreamServiceServer).SubscribeToStreamer(ctx, req.(*SubscribeToStreamerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TwitchStreamService_GetEventCountsAggregatedByStreamer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventCountsAggregatedByStreamerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwitchStreamServiceServer).GetEventCountsAggregatedByStreamer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_service_api.TwitchStreamService/GetEventCountsAggregatedByStreamer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwitchStreamServiceServer).GetEventCountsAggregatedByStreamer(ctx, req.(*GetEventCountsAggregatedByStreamerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TwitchStreamService_GetEventCountsAggregatedByStreamerAndType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventCountsAggregatedByStreamerAndTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TwitchStreamServiceServer).GetEventCountsAggregatedByStreamerAndType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_service_api.TwitchStreamService/GetEventCountsAggregatedByStreamerAndType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TwitchStreamServiceServer).GetEventCountsAggregatedByStreamerAndType(ctx, req.(*GetEventCountsAggregatedByStreamerAndTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TwitchStreamService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user_service_api.TwitchStreamService",
	HandlerType: (*TwitchStreamServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSignUrl",
			Handler:    _TwitchStreamService_GetSignUrl_Handler,
		},
		{
			MethodName: "ExchangeToken",
			Handler:    _TwitchStreamService_ExchangeToken_Handler,
		},
		{
			MethodName: "SubscribeToStreamer",
			Handler:    _TwitchStreamService_SubscribeToStreamer_Handler,
		},
		{
			MethodName: "GetEventCountsAggregatedByStreamer",
			Handler:    _TwitchStreamService_GetEventCountsAggregatedByStreamer_Handler,
		},
		{
			MethodName: "GetEventCountsAggregatedByStreamerAndType",
			Handler:    _TwitchStreamService_GetEventCountsAggregatedByStreamerAndType_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}