// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package api

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

type Record struct {
	Timestamp            int64    `protobuf:"varint,1,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	Topic                []byte   `protobuf:"bytes,2,opt,name=Topic,proto3" json:"Topic,omitempty"`
	Payload              []byte   `protobuf:"bytes,3,opt,name=Payload,proto3" json:"Payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Record) Reset()         { *m = Record{} }
func (m *Record) String() string { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()    {}
func (*Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *Record) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Record.Unmarshal(m, b)
}
func (m *Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Record.Marshal(b, m, deterministic)
}
func (m *Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Record.Merge(m, src)
}
func (m *Record) XXX_Size() int {
	return xxx_messageInfo_Record.Size(m)
}
func (m *Record) XXX_DiscardUnknown() {
	xxx_messageInfo_Record.DiscardUnknown(m)
}

var xxx_messageInfo_Record proto.InternalMessageInfo

func (m *Record) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Record) GetTopic() []byte {
	if m != nil {
		return m.Topic
	}
	return nil
}

func (m *Record) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type PutRecordsRequest struct {
	Records              []*Record `protobuf:"bytes,1,rep,name=Records,proto3" json:"Records,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PutRecordsRequest) Reset()         { *m = PutRecordsRequest{} }
func (m *PutRecordsRequest) String() string { return proto.CompactTextString(m) }
func (*PutRecordsRequest) ProtoMessage()    {}
func (*PutRecordsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *PutRecordsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutRecordsRequest.Unmarshal(m, b)
}
func (m *PutRecordsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutRecordsRequest.Marshal(b, m, deterministic)
}
func (m *PutRecordsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutRecordsRequest.Merge(m, src)
}
func (m *PutRecordsRequest) XXX_Size() int {
	return xxx_messageInfo_PutRecordsRequest.Size(m)
}
func (m *PutRecordsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PutRecordsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PutRecordsRequest proto.InternalMessageInfo

func (m *PutRecordsRequest) GetRecords() []*Record {
	if m != nil {
		return m.Records
	}
	return nil
}

type PutRecordsResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutRecordsResponse) Reset()         { *m = PutRecordsResponse{} }
func (m *PutRecordsResponse) String() string { return proto.CompactTextString(m) }
func (*PutRecordsResponse) ProtoMessage()    {}
func (*PutRecordsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2}
}

func (m *PutRecordsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutRecordsResponse.Unmarshal(m, b)
}
func (m *PutRecordsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutRecordsResponse.Marshal(b, m, deterministic)
}
func (m *PutRecordsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutRecordsResponse.Merge(m, src)
}
func (m *PutRecordsResponse) XXX_Size() int {
	return xxx_messageInfo_PutRecordsResponse.Size(m)
}
func (m *PutRecordsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PutRecordsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PutRecordsResponse proto.InternalMessageInfo

type GetRecordsRequest struct {
	FromTimestamp        int64    `protobuf:"varint,1,opt,name=FromTimestamp,proto3" json:"FromTimestamp,omitempty"`
	Patterns             [][]byte `protobuf:"bytes,2,rep,name=Patterns,proto3" json:"Patterns,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRecordsRequest) Reset()         { *m = GetRecordsRequest{} }
func (m *GetRecordsRequest) String() string { return proto.CompactTextString(m) }
func (*GetRecordsRequest) ProtoMessage()    {}
func (*GetRecordsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{3}
}

func (m *GetRecordsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRecordsRequest.Unmarshal(m, b)
}
func (m *GetRecordsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRecordsRequest.Marshal(b, m, deterministic)
}
func (m *GetRecordsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRecordsRequest.Merge(m, src)
}
func (m *GetRecordsRequest) XXX_Size() int {
	return xxx_messageInfo_GetRecordsRequest.Size(m)
}
func (m *GetRecordsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRecordsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRecordsRequest proto.InternalMessageInfo

func (m *GetRecordsRequest) GetFromTimestamp() int64 {
	if m != nil {
		return m.FromTimestamp
	}
	return 0
}

func (m *GetRecordsRequest) GetPatterns() [][]byte {
	if m != nil {
		return m.Patterns
	}
	return nil
}

type GetRecordsResponse struct {
	Records              []*Record `protobuf:"bytes,1,rep,name=Records,proto3" json:"Records,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetRecordsResponse) Reset()         { *m = GetRecordsResponse{} }
func (m *GetRecordsResponse) String() string { return proto.CompactTextString(m) }
func (*GetRecordsResponse) ProtoMessage()    {}
func (*GetRecordsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{4}
}

func (m *GetRecordsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRecordsResponse.Unmarshal(m, b)
}
func (m *GetRecordsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRecordsResponse.Marshal(b, m, deterministic)
}
func (m *GetRecordsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRecordsResponse.Merge(m, src)
}
func (m *GetRecordsResponse) XXX_Size() int {
	return xxx_messageInfo_GetRecordsResponse.Size(m)
}
func (m *GetRecordsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRecordsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetRecordsResponse proto.InternalMessageInfo

func (m *GetRecordsResponse) GetRecords() []*Record {
	if m != nil {
		return m.Records
	}
	return nil
}

type DumpRequest struct {
	DestinationURL       string   `protobuf:"bytes,1,opt,name=DestinationURL,proto3" json:"DestinationURL,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DumpRequest) Reset()         { *m = DumpRequest{} }
func (m *DumpRequest) String() string { return proto.CompactTextString(m) }
func (*DumpRequest) ProtoMessage()    {}
func (*DumpRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{5}
}

func (m *DumpRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DumpRequest.Unmarshal(m, b)
}
func (m *DumpRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DumpRequest.Marshal(b, m, deterministic)
}
func (m *DumpRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DumpRequest.Merge(m, src)
}
func (m *DumpRequest) XXX_Size() int {
	return xxx_messageInfo_DumpRequest.Size(m)
}
func (m *DumpRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DumpRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DumpRequest proto.InternalMessageInfo

func (m *DumpRequest) GetDestinationURL() string {
	if m != nil {
		return m.DestinationURL
	}
	return ""
}

type DumpResponse struct {
	StartedAt            int64    `protobuf:"varint,1,opt,name=StartedAt,proto3" json:"StartedAt,omitempty"`
	ProgressBytes        int64    `protobuf:"varint,2,opt,name=ProgressBytes,proto3" json:"ProgressBytes,omitempty"`
	TotalBytes           int64    `protobuf:"varint,3,opt,name=TotalBytes,proto3" json:"TotalBytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DumpResponse) Reset()         { *m = DumpResponse{} }
func (m *DumpResponse) String() string { return proto.CompactTextString(m) }
func (*DumpResponse) ProtoMessage()    {}
func (*DumpResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{6}
}

func (m *DumpResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DumpResponse.Unmarshal(m, b)
}
func (m *DumpResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DumpResponse.Marshal(b, m, deterministic)
}
func (m *DumpResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DumpResponse.Merge(m, src)
}
func (m *DumpResponse) XXX_Size() int {
	return xxx_messageInfo_DumpResponse.Size(m)
}
func (m *DumpResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DumpResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DumpResponse proto.InternalMessageInfo

func (m *DumpResponse) GetStartedAt() int64 {
	if m != nil {
		return m.StartedAt
	}
	return 0
}

func (m *DumpResponse) GetProgressBytes() int64 {
	if m != nil {
		return m.ProgressBytes
	}
	return 0
}

func (m *DumpResponse) GetTotalBytes() int64 {
	if m != nil {
		return m.TotalBytes
	}
	return 0
}

type SSTRequest struct {
	ToOffset             uint64   `protobuf:"varint,1,opt,name=ToOffset,proto3" json:"ToOffset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SSTRequest) Reset()         { *m = SSTRequest{} }
func (m *SSTRequest) String() string { return proto.CompactTextString(m) }
func (*SSTRequest) ProtoMessage()    {}
func (*SSTRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{7}
}

func (m *SSTRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SSTRequest.Unmarshal(m, b)
}
func (m *SSTRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SSTRequest.Marshal(b, m, deterministic)
}
func (m *SSTRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SSTRequest.Merge(m, src)
}
func (m *SSTRequest) XXX_Size() int {
	return xxx_messageInfo_SSTRequest.Size(m)
}
func (m *SSTRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SSTRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SSTRequest proto.InternalMessageInfo

func (m *SSTRequest) GetToOffset() uint64 {
	if m != nil {
		return m.ToOffset
	}
	return 0
}

type SSTResponseChunk struct {
	Chunk                []byte   `protobuf:"bytes,1,opt,name=Chunk,proto3" json:"Chunk,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SSTResponseChunk) Reset()         { *m = SSTResponseChunk{} }
func (m *SSTResponseChunk) String() string { return proto.CompactTextString(m) }
func (*SSTResponseChunk) ProtoMessage()    {}
func (*SSTResponseChunk) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{8}
}

func (m *SSTResponseChunk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SSTResponseChunk.Unmarshal(m, b)
}
func (m *SSTResponseChunk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SSTResponseChunk.Marshal(b, m, deterministic)
}
func (m *SSTResponseChunk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SSTResponseChunk.Merge(m, src)
}
func (m *SSTResponseChunk) XXX_Size() int {
	return xxx_messageInfo_SSTResponseChunk.Size(m)
}
func (m *SSTResponseChunk) XXX_DiscardUnknown() {
	xxx_messageInfo_SSTResponseChunk.DiscardUnknown(m)
}

var xxx_messageInfo_SSTResponseChunk proto.InternalMessageInfo

func (m *SSTResponseChunk) GetChunk() []byte {
	if m != nil {
		return m.Chunk
	}
	return nil
}

type LoadRequest struct {
	SourceURL            string   `protobuf:"bytes,1,opt,name=SourceURL,proto3" json:"SourceURL,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoadRequest) Reset()         { *m = LoadRequest{} }
func (m *LoadRequest) String() string { return proto.CompactTextString(m) }
func (*LoadRequest) ProtoMessage()    {}
func (*LoadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{9}
}

func (m *LoadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadRequest.Unmarshal(m, b)
}
func (m *LoadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadRequest.Marshal(b, m, deterministic)
}
func (m *LoadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadRequest.Merge(m, src)
}
func (m *LoadRequest) XXX_Size() int {
	return xxx_messageInfo_LoadRequest.Size(m)
}
func (m *LoadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoadRequest proto.InternalMessageInfo

func (m *LoadRequest) GetSourceURL() string {
	if m != nil {
		return m.SourceURL
	}
	return ""
}

type LoadResponse struct {
	StartedAt            int64    `protobuf:"varint,1,opt,name=StartedAt,proto3" json:"StartedAt,omitempty"`
	ProgressBytes        int64    `protobuf:"varint,2,opt,name=ProgressBytes,proto3" json:"ProgressBytes,omitempty"`
	TotalBytes           int64    `protobuf:"varint,3,opt,name=TotalBytes,proto3" json:"TotalBytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoadResponse) Reset()         { *m = LoadResponse{} }
func (m *LoadResponse) String() string { return proto.CompactTextString(m) }
func (*LoadResponse) ProtoMessage()    {}
func (*LoadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{10}
}

func (m *LoadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadResponse.Unmarshal(m, b)
}
func (m *LoadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadResponse.Marshal(b, m, deterministic)
}
func (m *LoadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadResponse.Merge(m, src)
}
func (m *LoadResponse) XXX_Size() int {
	return xxx_messageInfo_LoadResponse.Size(m)
}
func (m *LoadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoadResponse proto.InternalMessageInfo

func (m *LoadResponse) GetStartedAt() int64 {
	if m != nil {
		return m.StartedAt
	}
	return 0
}

func (m *LoadResponse) GetProgressBytes() int64 {
	if m != nil {
		return m.ProgressBytes
	}
	return 0
}

func (m *LoadResponse) GetTotalBytes() int64 {
	if m != nil {
		return m.TotalBytes
	}
	return 0
}

type TopicMetadata struct {
	Name                 []byte   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	MessageCount         uint64   `protobuf:"varint,2,opt,name=MessageCount,proto3" json:"MessageCount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TopicMetadata) Reset()         { *m = TopicMetadata{} }
func (m *TopicMetadata) String() string { return proto.CompactTextString(m) }
func (*TopicMetadata) ProtoMessage()    {}
func (*TopicMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{11}
}

func (m *TopicMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopicMetadata.Unmarshal(m, b)
}
func (m *TopicMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopicMetadata.Marshal(b, m, deterministic)
}
func (m *TopicMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopicMetadata.Merge(m, src)
}
func (m *TopicMetadata) XXX_Size() int {
	return xxx_messageInfo_TopicMetadata.Size(m)
}
func (m *TopicMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_TopicMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_TopicMetadata proto.InternalMessageInfo

func (m *TopicMetadata) GetName() []byte {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *TopicMetadata) GetMessageCount() uint64 {
	if m != nil {
		return m.MessageCount
	}
	return 0
}

type ListTopicsRequest struct {
	Pattern              []byte   `protobuf:"bytes,1,opt,name=Pattern,proto3" json:"Pattern,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTopicsRequest) Reset()         { *m = ListTopicsRequest{} }
func (m *ListTopicsRequest) String() string { return proto.CompactTextString(m) }
func (*ListTopicsRequest) ProtoMessage()    {}
func (*ListTopicsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{12}
}

func (m *ListTopicsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTopicsRequest.Unmarshal(m, b)
}
func (m *ListTopicsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTopicsRequest.Marshal(b, m, deterministic)
}
func (m *ListTopicsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTopicsRequest.Merge(m, src)
}
func (m *ListTopicsRequest) XXX_Size() int {
	return xxx_messageInfo_ListTopicsRequest.Size(m)
}
func (m *ListTopicsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTopicsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListTopicsRequest proto.InternalMessageInfo

func (m *ListTopicsRequest) GetPattern() []byte {
	if m != nil {
		return m.Pattern
	}
	return nil
}

type ListTopicsResponse struct {
	TopicMetadatas       []*TopicMetadata `protobuf:"bytes,1,rep,name=TopicMetadatas,proto3" json:"TopicMetadatas,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ListTopicsResponse) Reset()         { *m = ListTopicsResponse{} }
func (m *ListTopicsResponse) String() string { return proto.CompactTextString(m) }
func (*ListTopicsResponse) ProtoMessage()    {}
func (*ListTopicsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{13}
}

func (m *ListTopicsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTopicsResponse.Unmarshal(m, b)
}
func (m *ListTopicsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTopicsResponse.Marshal(b, m, deterministic)
}
func (m *ListTopicsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTopicsResponse.Merge(m, src)
}
func (m *ListTopicsResponse) XXX_Size() int {
	return xxx_messageInfo_ListTopicsResponse.Size(m)
}
func (m *ListTopicsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTopicsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListTopicsResponse proto.InternalMessageInfo

func (m *ListTopicsResponse) GetTopicMetadatas() []*TopicMetadata {
	if m != nil {
		return m.TopicMetadatas
	}
	return nil
}

func init() {
	proto.RegisterType((*Record)(nil), "api.Record")
	proto.RegisterType((*PutRecordsRequest)(nil), "api.PutRecordsRequest")
	proto.RegisterType((*PutRecordsResponse)(nil), "api.PutRecordsResponse")
	proto.RegisterType((*GetRecordsRequest)(nil), "api.GetRecordsRequest")
	proto.RegisterType((*GetRecordsResponse)(nil), "api.GetRecordsResponse")
	proto.RegisterType((*DumpRequest)(nil), "api.DumpRequest")
	proto.RegisterType((*DumpResponse)(nil), "api.DumpResponse")
	proto.RegisterType((*SSTRequest)(nil), "api.SSTRequest")
	proto.RegisterType((*SSTResponseChunk)(nil), "api.SSTResponseChunk")
	proto.RegisterType((*LoadRequest)(nil), "api.LoadRequest")
	proto.RegisterType((*LoadResponse)(nil), "api.LoadResponse")
	proto.RegisterType((*TopicMetadata)(nil), "api.TopicMetadata")
	proto.RegisterType((*ListTopicsRequest)(nil), "api.ListTopicsRequest")
	proto.RegisterType((*ListTopicsResponse)(nil), "api.ListTopicsResponse")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 541 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0x4f, 0x6f, 0xd3, 0x4e,
	0x10, 0x6d, 0xe2, 0xfc, 0xda, 0x64, 0x92, 0xf6, 0xd7, 0x8c, 0x0a, 0x58, 0x51, 0x85, 0xa2, 0x15,
	0x20, 0x4b, 0x88, 0x02, 0x45, 0x5c, 0xca, 0x01, 0x95, 0x56, 0xf4, 0x92, 0x82, 0xb5, 0x76, 0xb9,
	0x2f, 0xf1, 0xb6, 0x58, 0xd4, 0x5e, 0xe3, 0x5d, 0x1f, 0xfa, 0x7d, 0xf9, 0x20, 0xc8, 0xfb, 0x27,
	0xb6, 0xe3, 0x0b, 0x27, 0x6e, 0x9e, 0xe7, 0x37, 0x6f, 0xdf, 0xf8, 0xcd, 0x1a, 0x26, 0xac, 0x48,
	0x4f, 0x8a, 0x52, 0x28, 0x81, 0x1e, 0x2b, 0x52, 0xf2, 0x0d, 0x76, 0x29, 0x5f, 0x8b, 0x32, 0xc1,
	0x63, 0x98, 0xc4, 0x69, 0xc6, 0xa5, 0x62, 0x59, 0xe1, 0x0f, 0x96, 0x83, 0xc0, 0xa3, 0x0d, 0x80,
	0x47, 0xf0, 0x5f, 0x2c, 0x8a, 0x74, 0xed, 0x0f, 0x97, 0x83, 0x60, 0x46, 0x4d, 0x81, 0x3e, 0xec,
	0x85, 0xec, 0xe1, 0x5e, 0xb0, 0xc4, 0xf7, 0x34, 0xee, 0x4a, 0x72, 0x06, 0xf3, 0xb0, 0x52, 0x46,
	0x5a, 0x52, 0xfe, 0xab, 0xe2, 0x52, 0xe1, 0x73, 0xd8, 0xb3, 0x88, 0x3f, 0x58, 0x7a, 0xc1, 0xf4,
	0x74, 0x7a, 0x52, 0xdb, 0x31, 0x18, 0x75, 0xef, 0xc8, 0x11, 0x60, 0xbb, 0x57, 0x16, 0x22, 0x97,
	0x9c, 0xdc, 0xc0, 0xfc, 0x8a, 0x6f, 0x2b, 0x3e, 0x83, 0xfd, 0xcf, 0xa5, 0xc8, 0xb6, 0x8d, 0x77,
	0x41, 0x5c, 0xc0, 0x38, 0x64, 0x4a, 0xf1, 0x32, 0x97, 0xfe, 0x70, 0xe9, 0x05, 0x33, 0xba, 0xa9,
	0xc9, 0x07, 0xc0, 0xb6, 0xac, 0x39, 0xec, 0x6f, 0x9d, 0xbe, 0x87, 0xe9, 0x65, 0x95, 0x15, 0xce,
	0xcd, 0x0b, 0x38, 0xb8, 0xe4, 0x52, 0xa5, 0x39, 0x53, 0xa9, 0xc8, 0x6f, 0xe8, 0x4a, 0xdb, 0x99,
	0xd0, 0x2d, 0x94, 0x94, 0x30, 0x33, 0x6d, 0xf6, 0xb4, 0x63, 0x98, 0x44, 0x8a, 0x95, 0x8a, 0x27,
	0xe7, 0xca, 0x7d, 0xfa, 0x0d, 0x50, 0xcf, 0x18, 0x96, 0xe2, 0xae, 0xe4, 0x52, 0x7e, 0x7a, 0x50,
	0x5c, 0xea, 0x08, 0x3c, 0xda, 0x05, 0xf1, 0x29, 0x40, 0x2c, 0x14, 0xbb, 0x37, 0x14, 0x4f, 0x53,
	0x5a, 0x08, 0x09, 0x00, 0xa2, 0x28, 0x76, 0x4e, 0x17, 0x30, 0x8e, 0xc5, 0xd7, 0xdb, 0x5b, 0xc9,
	0xcd, 0x81, 0x23, 0xba, 0xa9, 0x49, 0x00, 0x87, 0x9a, 0x69, 0xcc, 0x5d, 0xfc, 0xa8, 0xf2, 0x9f,
	0x75, 0xfc, 0xfa, 0x41, 0x93, 0x67, 0xd4, 0x14, 0xe4, 0x25, 0x4c, 0x57, 0x82, 0x25, 0x4e, 0xb4,
	0x1e, 0x43, 0x54, 0xe5, 0x9a, 0x37, 0x93, 0x37, 0x40, 0x3d, 0xb4, 0x21, 0xff, 0xc3, 0xa1, 0xaf,
	0x60, 0x5f, 0x2f, 0xea, 0x35, 0x57, 0x2c, 0x61, 0x8a, 0x21, 0xc2, 0xe8, 0x0b, 0xcb, 0xb8, 0x1d,
	0x43, 0x3f, 0x23, 0x81, 0xd9, 0x35, 0x97, 0x92, 0xdd, 0xf1, 0x0b, 0x51, 0xe5, 0x4a, 0x9f, 0x34,
	0xa2, 0x1d, 0x8c, 0xbc, 0x82, 0xf9, 0x2a, 0x95, 0x4a, 0x8b, 0x6d, 0x96, 0x4f, 0x6f, 0xbf, 0x5e,
	0x23, 0xab, 0xe7, 0x4a, 0x12, 0x02, 0xb6, 0xe9, 0x76, 0xe2, 0x33, 0x38, 0xe8, 0xb8, 0x71, 0xbb,
	0x85, 0x7a, 0xb7, 0x3a, 0xaf, 0xe8, 0x16, 0xf3, 0xf4, 0xf7, 0x10, 0xc6, 0xd6, 0x91, 0xc4, 0xb7,
	0xe0, 0x45, 0x51, 0x8c, 0xff, 0xeb, 0xbe, 0x26, 0xd5, 0xc5, 0xa3, 0x06, 0x68, 0x85, 0x47, 0x76,
	0xde, 0x0c, 0xf0, 0x35, 0x8c, 0xea, 0x95, 0xc3, 0x43, 0x4d, 0x69, 0x2d, 0xed, 0x62, 0xde, 0x42,
	0xec, 0x55, 0xb3, 0x0d, 0x75, 0x5c, 0xb6, 0xa1, 0x15, 0xb3, 0x6d, 0x68, 0x67, 0xa9, 0x1b, 0x3e,
	0x02, 0x34, 0xb7, 0x16, 0x1f, 0x6b, 0x52, 0xef, 0x17, 0xb0, 0x78, 0xd2, 0xc3, 0x9d, 0x04, 0x9e,
	0x03, 0x34, 0x37, 0xd1, 0x0a, 0xf4, 0x6e, 0xbc, 0x15, 0xe8, 0x5f, 0x59, 0xe7, 0xa1, 0xf9, 0xee,
	0x56, 0xa2, 0x97, 0x9b, 0x95, 0xe8, 0x07, 0x44, 0x76, 0xbe, 0xef, 0xea, 0x5f, 0xe3, 0xbb, 0x3f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x11, 0xdd, 0x69, 0x61, 0x27, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MessagesClient is the client API for Messages service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MessagesClient interface {
	SST(ctx context.Context, in *SSTRequest, opts ...grpc.CallOption) (Messages_SSTClient, error)
	Dump(ctx context.Context, in *DumpRequest, opts ...grpc.CallOption) (Messages_DumpClient, error)
	Load(ctx context.Context, in *LoadRequest, opts ...grpc.CallOption) (Messages_LoadClient, error)
	PutRecords(ctx context.Context, in *PutRecordsRequest, opts ...grpc.CallOption) (*PutRecordsResponse, error)
	GetRecords(ctx context.Context, in *GetRecordsRequest, opts ...grpc.CallOption) (Messages_GetRecordsClient, error)
	ListTopics(ctx context.Context, in *ListTopicsRequest, opts ...grpc.CallOption) (*ListTopicsResponse, error)
}

type messagesClient struct {
	cc *grpc.ClientConn
}

func NewMessagesClient(cc *grpc.ClientConn) MessagesClient {
	return &messagesClient{cc}
}

func (c *messagesClient) SST(ctx context.Context, in *SSTRequest, opts ...grpc.CallOption) (Messages_SSTClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Messages_serviceDesc.Streams[0], "/api.Messages/SST", opts...)
	if err != nil {
		return nil, err
	}
	x := &messagesSSTClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Messages_SSTClient interface {
	Recv() (*SSTResponseChunk, error)
	grpc.ClientStream
}

type messagesSSTClient struct {
	grpc.ClientStream
}

func (x *messagesSSTClient) Recv() (*SSTResponseChunk, error) {
	m := new(SSTResponseChunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *messagesClient) Dump(ctx context.Context, in *DumpRequest, opts ...grpc.CallOption) (Messages_DumpClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Messages_serviceDesc.Streams[1], "/api.Messages/Dump", opts...)
	if err != nil {
		return nil, err
	}
	x := &messagesDumpClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Messages_DumpClient interface {
	Recv() (*DumpResponse, error)
	grpc.ClientStream
}

type messagesDumpClient struct {
	grpc.ClientStream
}

func (x *messagesDumpClient) Recv() (*DumpResponse, error) {
	m := new(DumpResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *messagesClient) Load(ctx context.Context, in *LoadRequest, opts ...grpc.CallOption) (Messages_LoadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Messages_serviceDesc.Streams[2], "/api.Messages/Load", opts...)
	if err != nil {
		return nil, err
	}
	x := &messagesLoadClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Messages_LoadClient interface {
	Recv() (*LoadResponse, error)
	grpc.ClientStream
}

type messagesLoadClient struct {
	grpc.ClientStream
}

func (x *messagesLoadClient) Recv() (*LoadResponse, error) {
	m := new(LoadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *messagesClient) PutRecords(ctx context.Context, in *PutRecordsRequest, opts ...grpc.CallOption) (*PutRecordsResponse, error) {
	out := new(PutRecordsResponse)
	err := c.cc.Invoke(ctx, "/api.Messages/PutRecords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagesClient) GetRecords(ctx context.Context, in *GetRecordsRequest, opts ...grpc.CallOption) (Messages_GetRecordsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Messages_serviceDesc.Streams[3], "/api.Messages/GetRecords", opts...)
	if err != nil {
		return nil, err
	}
	x := &messagesGetRecordsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Messages_GetRecordsClient interface {
	Recv() (*GetRecordsResponse, error)
	grpc.ClientStream
}

type messagesGetRecordsClient struct {
	grpc.ClientStream
}

func (x *messagesGetRecordsClient) Recv() (*GetRecordsResponse, error) {
	m := new(GetRecordsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *messagesClient) ListTopics(ctx context.Context, in *ListTopicsRequest, opts ...grpc.CallOption) (*ListTopicsResponse, error) {
	out := new(ListTopicsResponse)
	err := c.cc.Invoke(ctx, "/api.Messages/ListTopics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessagesServer is the server API for Messages service.
type MessagesServer interface {
	SST(*SSTRequest, Messages_SSTServer) error
	Dump(*DumpRequest, Messages_DumpServer) error
	Load(*LoadRequest, Messages_LoadServer) error
	PutRecords(context.Context, *PutRecordsRequest) (*PutRecordsResponse, error)
	GetRecords(*GetRecordsRequest, Messages_GetRecordsServer) error
	ListTopics(context.Context, *ListTopicsRequest) (*ListTopicsResponse, error)
}

// UnimplementedMessagesServer can be embedded to have forward compatible implementations.
type UnimplementedMessagesServer struct {
}

func (*UnimplementedMessagesServer) SST(req *SSTRequest, srv Messages_SSTServer) error {
	return status.Errorf(codes.Unimplemented, "method SST not implemented")
}
func (*UnimplementedMessagesServer) Dump(req *DumpRequest, srv Messages_DumpServer) error {
	return status.Errorf(codes.Unimplemented, "method Dump not implemented")
}
func (*UnimplementedMessagesServer) Load(req *LoadRequest, srv Messages_LoadServer) error {
	return status.Errorf(codes.Unimplemented, "method Load not implemented")
}
func (*UnimplementedMessagesServer) PutRecords(ctx context.Context, req *PutRecordsRequest) (*PutRecordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutRecords not implemented")
}
func (*UnimplementedMessagesServer) GetRecords(req *GetRecordsRequest, srv Messages_GetRecordsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetRecords not implemented")
}
func (*UnimplementedMessagesServer) ListTopics(ctx context.Context, req *ListTopicsRequest) (*ListTopicsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTopics not implemented")
}

func RegisterMessagesServer(s *grpc.Server, srv MessagesServer) {
	s.RegisterService(&_Messages_serviceDesc, srv)
}

func _Messages_SST_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SSTRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MessagesServer).SST(m, &messagesSSTServer{stream})
}

type Messages_SSTServer interface {
	Send(*SSTResponseChunk) error
	grpc.ServerStream
}

type messagesSSTServer struct {
	grpc.ServerStream
}

func (x *messagesSSTServer) Send(m *SSTResponseChunk) error {
	return x.ServerStream.SendMsg(m)
}

func _Messages_Dump_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DumpRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MessagesServer).Dump(m, &messagesDumpServer{stream})
}

type Messages_DumpServer interface {
	Send(*DumpResponse) error
	grpc.ServerStream
}

type messagesDumpServer struct {
	grpc.ServerStream
}

func (x *messagesDumpServer) Send(m *DumpResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Messages_Load_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(LoadRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MessagesServer).Load(m, &messagesLoadServer{stream})
}

type Messages_LoadServer interface {
	Send(*LoadResponse) error
	grpc.ServerStream
}

type messagesLoadServer struct {
	grpc.ServerStream
}

func (x *messagesLoadServer) Send(m *LoadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Messages_PutRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutRecordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagesServer).PutRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Messages/PutRecords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagesServer).PutRecords(ctx, req.(*PutRecordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Messages_GetRecords_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetRecordsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MessagesServer).GetRecords(m, &messagesGetRecordsServer{stream})
}

type Messages_GetRecordsServer interface {
	Send(*GetRecordsResponse) error
	grpc.ServerStream
}

type messagesGetRecordsServer struct {
	grpc.ServerStream
}

func (x *messagesGetRecordsServer) Send(m *GetRecordsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Messages_ListTopics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTopicsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagesServer).ListTopics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Messages/ListTopics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagesServer).ListTopics(ctx, req.(*ListTopicsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Messages_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Messages",
	HandlerType: (*MessagesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PutRecords",
			Handler:    _Messages_PutRecords_Handler,
		},
		{
			MethodName: "ListTopics",
			Handler:    _Messages_ListTopics_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SST",
			Handler:       _Messages_SST_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Dump",
			Handler:       _Messages_Dump_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Load",
			Handler:       _Messages_Load_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetRecords",
			Handler:       _Messages_GetRecords_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api.proto",
}
