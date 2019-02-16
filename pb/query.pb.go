// Code generated by protoc-gen-go. DO NOT EDIT.
// source: query.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/golang/protobuf/ptypes/any"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type QueryType int32

const (
	QueryType_THREAD_BACKUPS QueryType = 0
	QueryType_CONTACTS       QueryType = 1
)

var QueryType_name = map[int32]string{
	0: "THREAD_BACKUPS",
	1: "CONTACTS",
}
var QueryType_value = map[string]int32{
	"THREAD_BACKUPS": 0,
	"CONTACTS":       1,
}

func (x QueryType) String() string {
	return proto.EnumName(QueryType_name, int32(x))
}
func (QueryType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_query_15ca1d6cddf39022, []int{0}
}

type QueryOptions_FilterType int32

const (
	QueryOptions_NO_FILTER  QueryOptions_FilterType = 0
	QueryOptions_HIDE_OLDER QueryOptions_FilterType = 1
)

var QueryOptions_FilterType_name = map[int32]string{
	0: "NO_FILTER",
	1: "HIDE_OLDER",
}
var QueryOptions_FilterType_value = map[string]int32{
	"NO_FILTER":  0,
	"HIDE_OLDER": 1,
}

func (x QueryOptions_FilterType) String() string {
	return proto.EnumName(QueryOptions_FilterType_name, int32(x))
}
func (QueryOptions_FilterType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_query_15ca1d6cddf39022, []int{0, 0}
}

type PubSubQuery_ResponseType int32

const (
	PubSubQuery_P2P    PubSubQuery_ResponseType = 0
	PubSubQuery_PUBSUB PubSubQuery_ResponseType = 1
)

var PubSubQuery_ResponseType_name = map[int32]string{
	0: "P2P",
	1: "PUBSUB",
}
var PubSubQuery_ResponseType_value = map[string]int32{
	"P2P":    0,
	"PUBSUB": 1,
}

func (x PubSubQuery_ResponseType) String() string {
	return proto.EnumName(PubSubQuery_ResponseType_name, int32(x))
}
func (PubSubQuery_ResponseType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_query_15ca1d6cddf39022, []int{2, 0}
}

type QueryEvent_Type int32

const (
	QueryEvent_DATA QueryEvent_Type = 0
	QueryEvent_DONE QueryEvent_Type = 1
)

var QueryEvent_Type_name = map[int32]string{
	0: "DATA",
	1: "DONE",
}
var QueryEvent_Type_value = map[string]int32{
	"DATA": 0,
	"DONE": 1,
}

func (x QueryEvent_Type) String() string {
	return proto.EnumName(QueryEvent_Type_name, int32(x))
}
func (QueryEvent_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_query_15ca1d6cddf39022, []int{6, 0}
}

type QueryOptions struct {
	Local                bool                    `protobuf:"varint,1,opt,name=local,proto3" json:"local,omitempty"`
	Limit                int32                   `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Wait                 int32                   `protobuf:"varint,3,opt,name=wait,proto3" json:"wait,omitempty"`
	Filter               QueryOptions_FilterType `protobuf:"varint,4,opt,name=filter,proto3,enum=QueryOptions_FilterType" json:"filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *QueryOptions) Reset()         { *m = QueryOptions{} }
func (m *QueryOptions) String() string { return proto.CompactTextString(m) }
func (*QueryOptions) ProtoMessage()    {}
func (*QueryOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_query_15ca1d6cddf39022, []int{0}
}
func (m *QueryOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryOptions.Unmarshal(m, b)
}
func (m *QueryOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryOptions.Marshal(b, m, deterministic)
}
func (dst *QueryOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryOptions.Merge(dst, src)
}
func (m *QueryOptions) XXX_Size() int {
	return xxx_messageInfo_QueryOptions.Size(m)
}
func (m *QueryOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryOptions.DiscardUnknown(m)
}

var xxx_messageInfo_QueryOptions proto.InternalMessageInfo

func (m *QueryOptions) GetLocal() bool {
	if m != nil {
		return m.Local
	}
	return false
}

func (m *QueryOptions) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *QueryOptions) GetWait() int32 {
	if m != nil {
		return m.Wait
	}
	return 0
}

func (m *QueryOptions) GetFilter() QueryOptions_FilterType {
	if m != nil {
		return m.Filter
	}
	return QueryOptions_NO_FILTER
}

type Query struct {
	Id                   string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Token                string        `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Type                 QueryType     `protobuf:"varint,3,opt,name=type,proto3,enum=QueryType" json:"type,omitempty"`
	Options              *QueryOptions `protobuf:"bytes,4,opt,name=options,proto3" json:"options,omitempty"`
	Payload              *any.Any      `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Query) Reset()         { *m = Query{} }
func (m *Query) String() string { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()    {}
func (*Query) Descriptor() ([]byte, []int) {
	return fileDescriptor_query_15ca1d6cddf39022, []int{1}
}
func (m *Query) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Query.Unmarshal(m, b)
}
func (m *Query) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Query.Marshal(b, m, deterministic)
}
func (dst *Query) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Query.Merge(dst, src)
}
func (m *Query) XXX_Size() int {
	return xxx_messageInfo_Query.Size(m)
}
func (m *Query) XXX_DiscardUnknown() {
	xxx_messageInfo_Query.DiscardUnknown(m)
}

var xxx_messageInfo_Query proto.InternalMessageInfo

func (m *Query) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Query) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Query) GetType() QueryType {
	if m != nil {
		return m.Type
	}
	return QueryType_THREAD_BACKUPS
}

func (m *Query) GetOptions() *QueryOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *Query) GetPayload() *any.Any {
	if m != nil {
		return m.Payload
	}
	return nil
}

type PubSubQuery struct {
	Id                   string                   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type                 QueryType                `protobuf:"varint,2,opt,name=type,proto3,enum=QueryType" json:"type,omitempty"`
	Payload              *any.Any                 `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	ResponseType         PubSubQuery_ResponseType `protobuf:"varint,4,opt,name=responseType,proto3,enum=PubSubQuery_ResponseType" json:"responseType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *PubSubQuery) Reset()         { *m = PubSubQuery{} }
func (m *PubSubQuery) String() string { return proto.CompactTextString(m) }
func (*PubSubQuery) ProtoMessage()    {}
func (*PubSubQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_query_15ca1d6cddf39022, []int{2}
}
func (m *PubSubQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PubSubQuery.Unmarshal(m, b)
}
func (m *PubSubQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PubSubQuery.Marshal(b, m, deterministic)
}
func (dst *PubSubQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PubSubQuery.Merge(dst, src)
}
func (m *PubSubQuery) XXX_Size() int {
	return xxx_messageInfo_PubSubQuery.Size(m)
}
func (m *PubSubQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_PubSubQuery.DiscardUnknown(m)
}

var xxx_messageInfo_PubSubQuery proto.InternalMessageInfo

func (m *PubSubQuery) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PubSubQuery) GetType() QueryType {
	if m != nil {
		return m.Type
	}
	return QueryType_THREAD_BACKUPS
}

func (m *PubSubQuery) GetPayload() *any.Any {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *PubSubQuery) GetResponseType() PubSubQuery_ResponseType {
	if m != nil {
		return m.ResponseType
	}
	return PubSubQuery_P2P
}

type QueryResult struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Date                 *timestamp.Timestamp `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
	Local                bool                 `protobuf:"varint,3,opt,name=local,proto3" json:"local,omitempty"`
	Value                *any.Any             `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *QueryResult) Reset()         { *m = QueryResult{} }
func (m *QueryResult) String() string { return proto.CompactTextString(m) }
func (*QueryResult) ProtoMessage()    {}
func (*QueryResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_query_15ca1d6cddf39022, []int{3}
}
func (m *QueryResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryResult.Unmarshal(m, b)
}
func (m *QueryResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryResult.Marshal(b, m, deterministic)
}
func (dst *QueryResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryResult.Merge(dst, src)
}
func (m *QueryResult) XXX_Size() int {
	return xxx_messageInfo_QueryResult.Size(m)
}
func (m *QueryResult) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryResult.DiscardUnknown(m)
}

var xxx_messageInfo_QueryResult proto.InternalMessageInfo

func (m *QueryResult) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *QueryResult) GetDate() *timestamp.Timestamp {
	if m != nil {
		return m.Date
	}
	return nil
}

func (m *QueryResult) GetLocal() bool {
	if m != nil {
		return m.Local
	}
	return false
}

func (m *QueryResult) GetValue() *any.Any {
	if m != nil {
		return m.Value
	}
	return nil
}

type QueryResults struct {
	Type                 QueryType      `protobuf:"varint,1,opt,name=type,proto3,enum=QueryType" json:"type,omitempty"`
	Items                []*QueryResult `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *QueryResults) Reset()         { *m = QueryResults{} }
func (m *QueryResults) String() string { return proto.CompactTextString(m) }
func (*QueryResults) ProtoMessage()    {}
func (*QueryResults) Descriptor() ([]byte, []int) {
	return fileDescriptor_query_15ca1d6cddf39022, []int{4}
}
func (m *QueryResults) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryResults.Unmarshal(m, b)
}
func (m *QueryResults) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryResults.Marshal(b, m, deterministic)
}
func (dst *QueryResults) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryResults.Merge(dst, src)
}
func (m *QueryResults) XXX_Size() int {
	return xxx_messageInfo_QueryResults.Size(m)
}
func (m *QueryResults) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryResults.DiscardUnknown(m)
}

var xxx_messageInfo_QueryResults proto.InternalMessageInfo

func (m *QueryResults) GetType() QueryType {
	if m != nil {
		return m.Type
	}
	return QueryType_THREAD_BACKUPS
}

func (m *QueryResults) GetItems() []*QueryResult {
	if m != nil {
		return m.Items
	}
	return nil
}

type PubSubQueryResults struct {
	Id                   string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Results              *QueryResults `protobuf:"bytes,2,opt,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *PubSubQueryResults) Reset()         { *m = PubSubQueryResults{} }
func (m *PubSubQueryResults) String() string { return proto.CompactTextString(m) }
func (*PubSubQueryResults) ProtoMessage()    {}
func (*PubSubQueryResults) Descriptor() ([]byte, []int) {
	return fileDescriptor_query_15ca1d6cddf39022, []int{5}
}
func (m *PubSubQueryResults) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PubSubQueryResults.Unmarshal(m, b)
}
func (m *PubSubQueryResults) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PubSubQueryResults.Marshal(b, m, deterministic)
}
func (dst *PubSubQueryResults) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PubSubQueryResults.Merge(dst, src)
}
func (m *PubSubQueryResults) XXX_Size() int {
	return xxx_messageInfo_PubSubQueryResults.Size(m)
}
func (m *PubSubQueryResults) XXX_DiscardUnknown() {
	xxx_messageInfo_PubSubQueryResults.DiscardUnknown(m)
}

var xxx_messageInfo_PubSubQueryResults proto.InternalMessageInfo

func (m *PubSubQueryResults) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PubSubQueryResults) GetResults() *QueryResults {
	if m != nil {
		return m.Results
	}
	return nil
}

type QueryEvent struct {
	Type                 QueryEvent_Type `protobuf:"varint,1,opt,name=type,proto3,enum=QueryEvent_Type" json:"type,omitempty"`
	Data                 *QueryResult    `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *QueryEvent) Reset()         { *m = QueryEvent{} }
func (m *QueryEvent) String() string { return proto.CompactTextString(m) }
func (*QueryEvent) ProtoMessage()    {}
func (*QueryEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_query_15ca1d6cddf39022, []int{6}
}
func (m *QueryEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryEvent.Unmarshal(m, b)
}
func (m *QueryEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryEvent.Marshal(b, m, deterministic)
}
func (dst *QueryEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryEvent.Merge(dst, src)
}
func (m *QueryEvent) XXX_Size() int {
	return xxx_messageInfo_QueryEvent.Size(m)
}
func (m *QueryEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryEvent.DiscardUnknown(m)
}

var xxx_messageInfo_QueryEvent proto.InternalMessageInfo

func (m *QueryEvent) GetType() QueryEvent_Type {
	if m != nil {
		return m.Type
	}
	return QueryEvent_DATA
}

func (m *QueryEvent) GetData() *QueryResult {
	if m != nil {
		return m.Data
	}
	return nil
}

type ContactQuery struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Username             string   `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContactQuery) Reset()         { *m = ContactQuery{} }
func (m *ContactQuery) String() string { return proto.CompactTextString(m) }
func (*ContactQuery) ProtoMessage()    {}
func (*ContactQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_query_15ca1d6cddf39022, []int{7}
}
func (m *ContactQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContactQuery.Unmarshal(m, b)
}
func (m *ContactQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContactQuery.Marshal(b, m, deterministic)
}
func (dst *ContactQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContactQuery.Merge(dst, src)
}
func (m *ContactQuery) XXX_Size() int {
	return xxx_messageInfo_ContactQuery.Size(m)
}
func (m *ContactQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_ContactQuery.DiscardUnknown(m)
}

var xxx_messageInfo_ContactQuery proto.InternalMessageInfo

func (m *ContactQuery) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ContactQuery) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ContactQuery) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type ThreadBackupQuery struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ThreadBackupQuery) Reset()         { *m = ThreadBackupQuery{} }
func (m *ThreadBackupQuery) String() string { return proto.CompactTextString(m) }
func (*ThreadBackupQuery) ProtoMessage()    {}
func (*ThreadBackupQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_query_15ca1d6cddf39022, []int{8}
}
func (m *ThreadBackupQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThreadBackupQuery.Unmarshal(m, b)
}
func (m *ThreadBackupQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThreadBackupQuery.Marshal(b, m, deterministic)
}
func (dst *ThreadBackupQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThreadBackupQuery.Merge(dst, src)
}
func (m *ThreadBackupQuery) XXX_Size() int {
	return xxx_messageInfo_ThreadBackupQuery.Size(m)
}
func (m *ThreadBackupQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_ThreadBackupQuery.DiscardUnknown(m)
}

var xxx_messageInfo_ThreadBackupQuery proto.InternalMessageInfo

func (m *ThreadBackupQuery) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*QueryOptions)(nil), "QueryOptions")
	proto.RegisterType((*Query)(nil), "Query")
	proto.RegisterType((*PubSubQuery)(nil), "PubSubQuery")
	proto.RegisterType((*QueryResult)(nil), "QueryResult")
	proto.RegisterType((*QueryResults)(nil), "QueryResults")
	proto.RegisterType((*PubSubQueryResults)(nil), "PubSubQueryResults")
	proto.RegisterType((*QueryEvent)(nil), "QueryEvent")
	proto.RegisterType((*ContactQuery)(nil), "ContactQuery")
	proto.RegisterType((*ThreadBackupQuery)(nil), "ThreadBackupQuery")
	proto.RegisterEnum("QueryType", QueryType_name, QueryType_value)
	proto.RegisterEnum("QueryOptions_FilterType", QueryOptions_FilterType_name, QueryOptions_FilterType_value)
	proto.RegisterEnum("PubSubQuery_ResponseType", PubSubQuery_ResponseType_name, PubSubQuery_ResponseType_value)
	proto.RegisterEnum("QueryEvent_Type", QueryEvent_Type_name, QueryEvent_Type_value)
}

func init() { proto.RegisterFile("query.proto", fileDescriptor_query_15ca1d6cddf39022) }

var fileDescriptor_query_15ca1d6cddf39022 = []byte{
	// 612 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xcf, 0x4e, 0xdb, 0x40,
	0x10, 0xc6, 0xb3, 0x89, 0x43, 0x92, 0x49, 0x88, 0xdc, 0x15, 0x07, 0x93, 0x43, 0x1b, 0x6d, 0x2b,
	0x11, 0x51, 0xb1, 0x54, 0xee, 0xb9, 0x07, 0x27, 0x31, 0x02, 0x95, 0x92, 0x74, 0x63, 0x2e, 0xbd,
	0xa0, 0x0d, 0x5e, 0xa8, 0x85, 0x63, 0xbb, 0xf6, 0x9a, 0x2a, 0x4f, 0xd1, 0xb7, 0xe8, 0xad, 0x0f,
	0xd3, 0x37, 0xaa, 0xd8, 0xb5, 0xc1, 0xe1, 0x8f, 0xd4, 0x9b, 0x67, 0xbe, 0xf1, 0xb7, 0xbf, 0x99,
	0x9d, 0x85, 0xee, 0x8f, 0x5c, 0xa4, 0x6b, 0x9a, 0xa4, 0xb1, 0x8c, 0x07, 0xbb, 0xd7, 0x71, 0x7c,
	0x1d, 0x8a, 0x43, 0x15, 0x2d, 0xf3, 0xab, 0x43, 0x1e, 0x95, 0xd2, 0x9b, 0xc7, 0x92, 0x0c, 0x56,
	0x22, 0x93, 0x7c, 0x95, 0xe8, 0x02, 0xf2, 0x07, 0x41, 0xef, 0xeb, 0x9d, 0xd7, 0x2c, 0x91, 0x41,
	0x1c, 0x65, 0x78, 0x07, 0x9a, 0x61, 0x7c, 0xc9, 0x43, 0x0b, 0x0d, 0xd1, 0xa8, 0xcd, 0x74, 0xa0,
	0xb2, 0xc1, 0x2a, 0x90, 0x56, 0x7d, 0x88, 0x46, 0x4d, 0xa6, 0x03, 0x8c, 0xc1, 0xf8, 0xc9, 0x03,
	0x69, 0x35, 0x54, 0x52, 0x7d, 0xe3, 0x0f, 0xb0, 0x75, 0x15, 0x84, 0x52, 0xa4, 0x96, 0x31, 0x44,
	0xa3, 0xbe, 0x6d, 0xd1, 0xaa, 0x3d, 0x3d, 0x52, 0x9a, 0xb7, 0x4e, 0x04, 0x2b, 0xea, 0xc8, 0x7b,
	0x80, 0x87, 0x2c, 0xde, 0x86, 0xce, 0xd9, 0xec, 0xe2, 0xe8, 0xe4, 0xd4, 0x73, 0x99, 0x59, 0xc3,
	0x7d, 0x80, 0xe3, 0x93, 0xa9, 0x7b, 0x31, 0x3b, 0x9d, 0xba, 0xcc, 0x44, 0xe4, 0x37, 0x82, 0xa6,
	0x32, 0xc4, 0x7d, 0xa8, 0x07, 0xbe, 0xa2, 0xec, 0xb0, 0x7a, 0xe0, 0xdf, 0x21, 0xca, 0xf8, 0x46,
	0x44, 0x0a, 0xb1, 0xc3, 0x74, 0x80, 0x5f, 0x83, 0x21, 0xd7, 0x89, 0x50, 0x88, 0x7d, 0x1b, 0x34,
	0x8c, 0x3a, 0x5e, 0xe5, 0xf1, 0x1e, 0xb4, 0x62, 0x8d, 0xa6, 0x78, 0xbb, 0xf6, 0xf6, 0x06, 0x2f,
	0x2b, 0x55, 0x4c, 0xa1, 0x95, 0xf0, 0x75, 0x18, 0x73, 0xdf, 0x6a, 0xaa, 0xc2, 0x1d, 0xaa, 0x67,
	0x4b, 0xcb, 0xd9, 0x52, 0x27, 0x5a, 0xb3, 0xb2, 0x88, 0xfc, 0x45, 0xd0, 0x9d, 0xe7, 0xcb, 0x45,
	0xbe, 0x7c, 0x1e, 0xb7, 0x04, 0xab, 0xbf, 0x00, 0x56, 0x39, 0xaf, 0xf1, 0x1f, 0xe7, 0xe1, 0x4f,
	0xd0, 0x4b, 0x45, 0x96, 0xc4, 0x51, 0x26, 0xee, 0x5c, 0x8a, 0xe9, 0xef, 0xd2, 0x0a, 0x03, 0x65,
	0x95, 0x02, 0xb6, 0x51, 0x4e, 0xde, 0x42, 0xaf, 0xaa, 0xe2, 0x16, 0x34, 0xe6, 0xf6, 0xdc, 0xac,
	0x61, 0x80, 0xad, 0xf9, 0xf9, 0x78, 0x71, 0x3e, 0x36, 0x11, 0xf9, 0x85, 0xa0, 0xab, 0x9c, 0x98,
	0xc8, 0xf2, 0x50, 0x3e, 0xe9, 0x89, 0x82, 0xe1, 0x73, 0xa9, 0x7b, 0xea, 0xda, 0x83, 0x27, 0xc0,
	0x5e, 0xb9, 0x7c, 0x4c, 0xd5, 0x3d, 0xec, 0x5a, 0xa3, 0xba, 0x6b, 0xfb, 0xd0, 0xbc, 0xe5, 0x61,
	0x2e, 0x8a, 0x0b, 0x79, 0xbe, 0x6f, 0x5d, 0x42, 0x58, 0xb1, 0xbd, 0x1a, 0x28, 0xbb, 0x9f, 0x2a,
	0x7a, 0x61, 0xaa, 0x04, 0x9a, 0x81, 0x14, 0xab, 0xcc, 0xaa, 0x0f, 0x1b, 0xa3, 0xae, 0xdd, 0xa3,
	0x95, 0xbf, 0x99, 0x96, 0xc8, 0x17, 0xc0, 0x95, 0xa1, 0x95, 0xce, 0x8f, 0x7b, 0xdd, 0x83, 0x56,
	0xaa, 0xa5, 0xa2, 0xdd, 0xed, 0xaa, 0x57, 0xc6, 0x4a, 0x95, 0x24, 0x00, 0x4a, 0x70, 0x6f, 0x45,
	0x24, 0xf1, 0xbb, 0x0d, 0x40, 0x93, 0x3e, 0x48, 0xb4, 0x82, 0x39, 0x54, 0x83, 0xe4, 0x85, 0xf3,
	0x26, 0xa5, 0x52, 0xc8, 0x00, 0x0c, 0x75, 0x4f, 0x6d, 0x30, 0xa6, 0x8e, 0xe7, 0x98, 0x35, 0xf5,
	0x35, 0x3b, 0x73, 0x4d, 0x44, 0x3c, 0xe8, 0x4d, 0xe2, 0x48, 0xf2, 0x4b, 0xf9, 0xfc, 0xea, 0x59,
	0xd0, 0xe2, 0xbe, 0x9f, 0x8a, 0x2c, 0x2b, 0xde, 0x4a, 0x19, 0xe2, 0x01, 0xb4, 0xf3, 0x4c, 0xa4,
	0x11, 0x5f, 0xe9, 0x17, 0xd3, 0x61, 0xf7, 0x31, 0x39, 0x80, 0x57, 0xde, 0xf7, 0x54, 0x70, 0x7f,
	0xcc, 0x2f, 0x6f, 0xf2, 0x44, 0x5b, 0x57, 0xac, 0xd0, 0x86, 0xd5, 0xfe, 0x01, 0x74, 0xee, 0x87,
	0x8f, 0x31, 0xf4, 0xbd, 0x63, 0xe6, 0x3a, 0xd3, 0x8b, 0xb1, 0x33, 0xf9, 0x7c, 0x3e, 0x5f, 0x98,
	0x35, 0xdc, 0x83, 0xf6, 0x64, 0x76, 0xe6, 0x39, 0x13, 0x6f, 0x61, 0xa2, 0xb1, 0xf1, 0xad, 0x9e,
	0x2c, 0x97, 0x5b, 0xea, 0x8e, 0x3f, 0xfe, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x2f, 0x0c, 0x8e, 0x88,
	0xdf, 0x04, 0x00, 0x00,
}
