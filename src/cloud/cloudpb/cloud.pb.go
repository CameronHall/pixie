// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: src/cloud/cloudpb/cloud.proto

package cloudpb

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	proto1 "pixielabs.ai/pixielabs/src/common/uuid/proto"
	reflect "reflect"
	strconv "strconv"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type RegisterVizierAck_RegistrationStatus int32

const (
	ST_UNKNOWN          RegisterVizierAck_RegistrationStatus = 0
	ST_OK               RegisterVizierAck_RegistrationStatus = 1
	ST_FAILED_NOT_FOUND RegisterVizierAck_RegistrationStatus = 2
)

var RegisterVizierAck_RegistrationStatus_name = map[int32]string{
	0: "ST_UNKNOWN",
	1: "ST_OK",
	2: "ST_FAILED_NOT_FOUND",
}

var RegisterVizierAck_RegistrationStatus_value = map[string]int32{
	"ST_UNKNOWN":          0,
	"ST_OK":               1,
	"ST_FAILED_NOT_FOUND": 2,
}

func (RegisterVizierAck_RegistrationStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1fc79a51051c56a2, []int{1, 0}
}

type VizierHeartbeatAck_HeartbeatStatus int32

const (
	HB_UNKNOWN VizierHeartbeatAck_HeartbeatStatus = 0
	HB_OK      VizierHeartbeatAck_HeartbeatStatus = 1
	HB_ERROR   VizierHeartbeatAck_HeartbeatStatus = 2
)

var VizierHeartbeatAck_HeartbeatStatus_name = map[int32]string{
	0: "HB_UNKNOWN",
	1: "HB_OK",
	2: "HB_ERROR",
}

var VizierHeartbeatAck_HeartbeatStatus_value = map[string]int32{
	"HB_UNKNOWN": 0,
	"HB_OK":      1,
	"HB_ERROR":   2,
}

func (VizierHeartbeatAck_HeartbeatStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1fc79a51051c56a2, []int{3, 0}
}

type RegisterVizierRequest struct {
	VizierID *proto1.UUID `protobuf:"bytes,1,opt,name=vizier_id,json=vizierId,proto3" json:"vizier_id,omitempty"`
	JwtKey   string       `protobuf:"bytes,2,opt,name=jwt_key,json=jwtKey,proto3" json:"jwt_key,omitempty"`
}

func (m *RegisterVizierRequest) Reset()      { *m = RegisterVizierRequest{} }
func (*RegisterVizierRequest) ProtoMessage() {}
func (*RegisterVizierRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fc79a51051c56a2, []int{0}
}
func (m *RegisterVizierRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegisterVizierRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegisterVizierRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegisterVizierRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterVizierRequest.Merge(m, src)
}
func (m *RegisterVizierRequest) XXX_Size() int {
	return m.Size()
}
func (m *RegisterVizierRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterVizierRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterVizierRequest proto.InternalMessageInfo

func (m *RegisterVizierRequest) GetVizierID() *proto1.UUID {
	if m != nil {
		return m.VizierID
	}
	return nil
}

func (m *RegisterVizierRequest) GetJwtKey() string {
	if m != nil {
		return m.JwtKey
	}
	return ""
}

type RegisterVizierAck struct {
	Status RegisterVizierAck_RegistrationStatus `protobuf:"varint,1,opt,name=status,proto3,enum=pl.services.RegisterVizierAck_RegistrationStatus" json:"status,omitempty"`
}

func (m *RegisterVizierAck) Reset()      { *m = RegisterVizierAck{} }
func (*RegisterVizierAck) ProtoMessage() {}
func (*RegisterVizierAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fc79a51051c56a2, []int{1}
}
func (m *RegisterVizierAck) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegisterVizierAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegisterVizierAck.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegisterVizierAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterVizierAck.Merge(m, src)
}
func (m *RegisterVizierAck) XXX_Size() int {
	return m.Size()
}
func (m *RegisterVizierAck) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterVizierAck.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterVizierAck proto.InternalMessageInfo

func (m *RegisterVizierAck) GetStatus() RegisterVizierAck_RegistrationStatus {
	if m != nil {
		return m.Status
	}
	return ST_UNKNOWN
}

type VizierHeartbeat struct {
	Time           int64 `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	SequenceNumber int64 `protobuf:"varint,2,opt,name=sequence_number,json=sequenceNumber,proto3" json:"sequence_number,omitempty"`
}

func (m *VizierHeartbeat) Reset()      { *m = VizierHeartbeat{} }
func (*VizierHeartbeat) ProtoMessage() {}
func (*VizierHeartbeat) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fc79a51051c56a2, []int{2}
}
func (m *VizierHeartbeat) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VizierHeartbeat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VizierHeartbeat.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VizierHeartbeat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VizierHeartbeat.Merge(m, src)
}
func (m *VizierHeartbeat) XXX_Size() int {
	return m.Size()
}
func (m *VizierHeartbeat) XXX_DiscardUnknown() {
	xxx_messageInfo_VizierHeartbeat.DiscardUnknown(m)
}

var xxx_messageInfo_VizierHeartbeat proto.InternalMessageInfo

func (m *VizierHeartbeat) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *VizierHeartbeat) GetSequenceNumber() int64 {
	if m != nil {
		return m.SequenceNumber
	}
	return 0
}

type VizierHeartbeatAck struct {
	Status         VizierHeartbeatAck_HeartbeatStatus `protobuf:"varint,1,opt,name=status,proto3,enum=pl.services.VizierHeartbeatAck_HeartbeatStatus" json:"status,omitempty"`
	Time           int64                              `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
	SequenceNumber int64                              `protobuf:"varint,3,opt,name=sequence_number,json=sequenceNumber,proto3" json:"sequence_number,omitempty"`
	ErrorMessage   string                             `protobuf:"bytes,4,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
}

func (m *VizierHeartbeatAck) Reset()      { *m = VizierHeartbeatAck{} }
func (*VizierHeartbeatAck) ProtoMessage() {}
func (*VizierHeartbeatAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fc79a51051c56a2, []int{3}
}
func (m *VizierHeartbeatAck) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VizierHeartbeatAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VizierHeartbeatAck.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VizierHeartbeatAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VizierHeartbeatAck.Merge(m, src)
}
func (m *VizierHeartbeatAck) XXX_Size() int {
	return m.Size()
}
func (m *VizierHeartbeatAck) XXX_DiscardUnknown() {
	xxx_messageInfo_VizierHeartbeatAck.DiscardUnknown(m)
}

var xxx_messageInfo_VizierHeartbeatAck proto.InternalMessageInfo

func (m *VizierHeartbeatAck) GetStatus() VizierHeartbeatAck_HeartbeatStatus {
	if m != nil {
		return m.Status
	}
	return HB_UNKNOWN
}

func (m *VizierHeartbeatAck) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *VizierHeartbeatAck) GetSequenceNumber() int64 {
	if m != nil {
		return m.SequenceNumber
	}
	return 0
}

func (m *VizierHeartbeatAck) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func init() {
	proto.RegisterEnum("pl.services.RegisterVizierAck_RegistrationStatus", RegisterVizierAck_RegistrationStatus_name, RegisterVizierAck_RegistrationStatus_value)
	proto.RegisterEnum("pl.services.VizierHeartbeatAck_HeartbeatStatus", VizierHeartbeatAck_HeartbeatStatus_name, VizierHeartbeatAck_HeartbeatStatus_value)
	proto.RegisterType((*RegisterVizierRequest)(nil), "pl.services.RegisterVizierRequest")
	proto.RegisterType((*RegisterVizierAck)(nil), "pl.services.RegisterVizierAck")
	proto.RegisterType((*VizierHeartbeat)(nil), "pl.services.VizierHeartbeat")
	proto.RegisterType((*VizierHeartbeatAck)(nil), "pl.services.VizierHeartbeatAck")
}

func init() { proto.RegisterFile("src/cloud/cloudpb/cloud.proto", fileDescriptor_1fc79a51051c56a2) }

var fileDescriptor_1fc79a51051c56a2 = []byte{
	// 520 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0x41, 0x73, 0xd2, 0x40,
	0x14, 0xce, 0x42, 0x45, 0xd8, 0x22, 0xe0, 0x3a, 0x4e, 0x99, 0xce, 0xb8, 0x32, 0xf1, 0x60, 0x2f,
	0x26, 0x5a, 0x6f, 0xf5, 0x54, 0x86, 0x56, 0x18, 0x34, 0xcc, 0x04, 0xd0, 0x19, 0x2f, 0x99, 0x24,
	0x6c, 0x71, 0x5b, 0xc2, 0xc6, 0xdd, 0x4d, 0x6b, 0x3d, 0xf9, 0x13, 0xfc, 0x19, 0xfd, 0x29, 0x1e,
	0x39, 0xf6, 0xe4, 0x48, 0xb8, 0x78, 0xec, 0xdd, 0x8b, 0xc3, 0x2e, 0x32, 0x02, 0x4e, 0x2f, 0xd9,
	0xf7, 0xbe, 0x7c, 0xdf, 0x7b, 0xef, 0x7b, 0xc9, 0xc2, 0x47, 0x82, 0x87, 0x76, 0x38, 0x62, 0xc9,
	0x40, 0x3f, 0xe3, 0x40, 0x9f, 0x56, 0xcc, 0x99, 0x64, 0x68, 0x3b, 0x1e, 0x59, 0x82, 0xf0, 0x73,
	0x1a, 0x12, 0xb1, 0xfb, 0x6c, 0x48, 0xe5, 0xc7, 0x24, 0xb0, 0x42, 0x16, 0xd9, 0x43, 0x36, 0x64,
	0xb6, 0xe2, 0x04, 0xc9, 0x89, 0xca, 0x54, 0xa2, 0x22, 0xad, 0xdd, 0xad, 0xa9, 0xd2, 0x2c, 0x8a,
	0xd8, 0xd8, 0x4e, 0x12, 0x3a, 0xd0, 0x74, 0x15, 0x6a, 0x86, 0x39, 0x82, 0x0f, 0x5d, 0x32, 0xa4,
	0x42, 0x12, 0xfe, 0x8e, 0x7e, 0xa1, 0x84, 0xbb, 0xe4, 0x53, 0x42, 0x84, 0x44, 0x07, 0xb0, 0x70,
	0xae, 0x00, 0x8f, 0x0e, 0xaa, 0xa0, 0x06, 0xf6, 0xb6, 0xf7, 0xcb, 0x56, 0x3c, 0xb2, 0xe6, 0xda,
	0x38, 0xb0, 0xfa, 0xfd, 0x56, 0xa3, 0x5e, 0x4c, 0x7f, 0x3c, 0xce, 0x6b, 0x59, 0xab, 0xe1, 0xe6,
	0x35, 0xbf, 0x35, 0x40, 0x3b, 0xf0, 0xee, 0xe9, 0x85, 0xf4, 0xce, 0xc8, 0x65, 0x35, 0x53, 0x03,
	0x7b, 0x05, 0x37, 0x77, 0x7a, 0x21, 0xdb, 0xe4, 0xd2, 0xbc, 0x02, 0xf0, 0xfe, 0x6a, 0xbb, 0xc3,
	0xf0, 0x0c, 0xb5, 0x60, 0x4e, 0x48, 0x5f, 0x26, 0x42, 0xf5, 0x29, 0xed, 0xbf, 0xb0, 0xfe, 0xb1,
	0x6c, 0x6d, 0xf0, 0x17, 0x08, 0xf7, 0x25, 0x65, 0xe3, 0xae, 0x12, 0xba, 0x8b, 0x02, 0x66, 0x13,
	0xa2, 0xcd, 0xb7, 0xa8, 0x04, 0x61, 0xb7, 0xe7, 0xf5, 0x9d, 0xb6, 0xd3, 0x79, 0xef, 0x54, 0x0c,
	0x54, 0x80, 0x77, 0xba, 0x3d, 0xaf, 0xd3, 0xae, 0x00, 0xb4, 0x03, 0x1f, 0x74, 0x7b, 0xde, 0xf1,
	0x61, 0xeb, 0xcd, 0x51, 0xc3, 0x73, 0x3a, 0x3d, 0xef, 0xb8, 0xd3, 0x77, 0x1a, 0x95, 0x8c, 0xe9,
	0xc0, 0xb2, 0xee, 0xd8, 0x24, 0x3e, 0x97, 0x01, 0xf1, 0x25, 0x42, 0x70, 0x4b, 0xd2, 0x88, 0xa8,
	0x29, 0xb3, 0xae, 0x8a, 0xd1, 0x53, 0x58, 0x16, 0xf3, 0x8d, 0x8d, 0x43, 0xe2, 0x8d, 0x93, 0x28,
	0x20, 0x5c, 0x59, 0xce, 0xba, 0xa5, 0xbf, 0xb0, 0xa3, 0x50, 0xf3, 0x37, 0x80, 0x68, 0xad, 0xe0,
	0xdc, 0xfb, 0xeb, 0x35, 0xef, 0xf6, 0x8a, 0xf7, 0x4d, 0x81, 0xb5, 0x4c, 0x56, 0x9d, 0x2f, 0x87,
	0xcb, 0xdc, 0x3e, 0x5c, 0xf6, 0x7f, 0xc3, 0xa1, 0x27, 0xf0, 0x1e, 0xe1, 0x9c, 0x71, 0x2f, 0x22,
	0x42, 0xf8, 0x43, 0x52, 0xdd, 0x52, 0x9f, 0xad, 0xa8, 0xc0, 0xb7, 0x1a, 0x33, 0x0f, 0x60, 0x79,
	0xad, 0xf9, 0x7c, 0xb1, 0xcd, 0xfa, 0xea, 0x62, 0x9b, 0x75, 0xbd, 0xd8, 0x22, 0xcc, 0x37, 0xeb,
	0xde, 0x91, 0xeb, 0x76, 0xdc, 0x4a, 0xa6, 0x7e, 0x32, 0x99, 0x62, 0xe3, 0x7a, 0x8a, 0x8d, 0x9b,
	0x29, 0x06, 0x5f, 0x53, 0x0c, 0xae, 0x52, 0x0c, 0xbe, 0xa7, 0x18, 0x4c, 0x52, 0x0c, 0x7e, 0xa6,
	0x18, 0xfc, 0x4a, 0xb1, 0x71, 0x93, 0x62, 0xf0, 0x6d, 0x86, 0x8d, 0xc9, 0x0c, 0x1b, 0xd7, 0x33,
	0x6c, 0x7c, 0x78, 0x1e, 0xd3, 0xcf, 0x94, 0x8c, 0xfc, 0x40, 0x58, 0x3e, 0xb5, 0x97, 0x89, 0xbd,
	0x71, 0x61, 0x5e, 0x2d, 0xce, 0x20, 0xa7, 0xfe, 0xea, 0x97, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff,
	0xc6, 0xf1, 0x7b, 0x04, 0x54, 0x03, 0x00, 0x00,
}

func (x RegisterVizierAck_RegistrationStatus) String() string {
	s, ok := RegisterVizierAck_RegistrationStatus_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x VizierHeartbeatAck_HeartbeatStatus) String() string {
	s, ok := VizierHeartbeatAck_HeartbeatStatus_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *RegisterVizierRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RegisterVizierRequest)
	if !ok {
		that2, ok := that.(RegisterVizierRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.VizierID.Equal(that1.VizierID) {
		return false
	}
	if this.JwtKey != that1.JwtKey {
		return false
	}
	return true
}
func (this *RegisterVizierAck) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RegisterVizierAck)
	if !ok {
		that2, ok := that.(RegisterVizierAck)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Status != that1.Status {
		return false
	}
	return true
}
func (this *VizierHeartbeat) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VizierHeartbeat)
	if !ok {
		that2, ok := that.(VizierHeartbeat)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Time != that1.Time {
		return false
	}
	if this.SequenceNumber != that1.SequenceNumber {
		return false
	}
	return true
}
func (this *VizierHeartbeatAck) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VizierHeartbeatAck)
	if !ok {
		that2, ok := that.(VizierHeartbeatAck)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Status != that1.Status {
		return false
	}
	if this.Time != that1.Time {
		return false
	}
	if this.SequenceNumber != that1.SequenceNumber {
		return false
	}
	if this.ErrorMessage != that1.ErrorMessage {
		return false
	}
	return true
}
func (this *RegisterVizierRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&cloudpb.RegisterVizierRequest{")
	if this.VizierID != nil {
		s = append(s, "VizierID: "+fmt.Sprintf("%#v", this.VizierID)+",\n")
	}
	s = append(s, "JwtKey: "+fmt.Sprintf("%#v", this.JwtKey)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *RegisterVizierAck) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&cloudpb.RegisterVizierAck{")
	s = append(s, "Status: "+fmt.Sprintf("%#v", this.Status)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *VizierHeartbeat) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&cloudpb.VizierHeartbeat{")
	s = append(s, "Time: "+fmt.Sprintf("%#v", this.Time)+",\n")
	s = append(s, "SequenceNumber: "+fmt.Sprintf("%#v", this.SequenceNumber)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *VizierHeartbeatAck) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&cloudpb.VizierHeartbeatAck{")
	s = append(s, "Status: "+fmt.Sprintf("%#v", this.Status)+",\n")
	s = append(s, "Time: "+fmt.Sprintf("%#v", this.Time)+",\n")
	s = append(s, "SequenceNumber: "+fmt.Sprintf("%#v", this.SequenceNumber)+",\n")
	s = append(s, "ErrorMessage: "+fmt.Sprintf("%#v", this.ErrorMessage)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringCloud(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *RegisterVizierRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisterVizierRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.VizierID != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCloud(dAtA, i, uint64(m.VizierID.Size()))
		n1, err := m.VizierID.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.JwtKey) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCloud(dAtA, i, uint64(len(m.JwtKey)))
		i += copy(dAtA[i:], m.JwtKey)
	}
	return i, nil
}

func (m *RegisterVizierAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisterVizierAck) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintCloud(dAtA, i, uint64(m.Status))
	}
	return i, nil
}

func (m *VizierHeartbeat) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VizierHeartbeat) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Time != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintCloud(dAtA, i, uint64(m.Time))
	}
	if m.SequenceNumber != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintCloud(dAtA, i, uint64(m.SequenceNumber))
	}
	return i, nil
}

func (m *VizierHeartbeatAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VizierHeartbeatAck) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintCloud(dAtA, i, uint64(m.Status))
	}
	if m.Time != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintCloud(dAtA, i, uint64(m.Time))
	}
	if m.SequenceNumber != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintCloud(dAtA, i, uint64(m.SequenceNumber))
	}
	if len(m.ErrorMessage) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintCloud(dAtA, i, uint64(len(m.ErrorMessage)))
		i += copy(dAtA[i:], m.ErrorMessage)
	}
	return i, nil
}

func encodeVarintCloud(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *RegisterVizierRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.VizierID != nil {
		l = m.VizierID.Size()
		n += 1 + l + sovCloud(uint64(l))
	}
	l = len(m.JwtKey)
	if l > 0 {
		n += 1 + l + sovCloud(uint64(l))
	}
	return n
}

func (m *RegisterVizierAck) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Status != 0 {
		n += 1 + sovCloud(uint64(m.Status))
	}
	return n
}

func (m *VizierHeartbeat) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Time != 0 {
		n += 1 + sovCloud(uint64(m.Time))
	}
	if m.SequenceNumber != 0 {
		n += 1 + sovCloud(uint64(m.SequenceNumber))
	}
	return n
}

func (m *VizierHeartbeatAck) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Status != 0 {
		n += 1 + sovCloud(uint64(m.Status))
	}
	if m.Time != 0 {
		n += 1 + sovCloud(uint64(m.Time))
	}
	if m.SequenceNumber != 0 {
		n += 1 + sovCloud(uint64(m.SequenceNumber))
	}
	l = len(m.ErrorMessage)
	if l > 0 {
		n += 1 + l + sovCloud(uint64(l))
	}
	return n
}

func sovCloud(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCloud(x uint64) (n int) {
	return sovCloud(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *RegisterVizierRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&RegisterVizierRequest{`,
		`VizierID:` + strings.Replace(fmt.Sprintf("%v", this.VizierID), "UUID", "proto1.UUID", 1) + `,`,
		`JwtKey:` + fmt.Sprintf("%v", this.JwtKey) + `,`,
		`}`,
	}, "")
	return s
}
func (this *RegisterVizierAck) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&RegisterVizierAck{`,
		`Status:` + fmt.Sprintf("%v", this.Status) + `,`,
		`}`,
	}, "")
	return s
}
func (this *VizierHeartbeat) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&VizierHeartbeat{`,
		`Time:` + fmt.Sprintf("%v", this.Time) + `,`,
		`SequenceNumber:` + fmt.Sprintf("%v", this.SequenceNumber) + `,`,
		`}`,
	}, "")
	return s
}
func (this *VizierHeartbeatAck) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&VizierHeartbeatAck{`,
		`Status:` + fmt.Sprintf("%v", this.Status) + `,`,
		`Time:` + fmt.Sprintf("%v", this.Time) + `,`,
		`SequenceNumber:` + fmt.Sprintf("%v", this.SequenceNumber) + `,`,
		`ErrorMessage:` + fmt.Sprintf("%v", this.ErrorMessage) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringCloud(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *RegisterVizierRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCloud
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RegisterVizierRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisterVizierRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VizierID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCloud
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCloud
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCloud
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.VizierID == nil {
				m.VizierID = &proto1.UUID{}
			}
			if err := m.VizierID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JwtKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCloud
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCloud
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCloud
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.JwtKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCloud(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCloud
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCloud
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RegisterVizierAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCloud
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RegisterVizierAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisterVizierAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCloud
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= RegisterVizierAck_RegistrationStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCloud(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCloud
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCloud
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *VizierHeartbeat) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCloud
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: VizierHeartbeat: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VizierHeartbeat: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			m.Time = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCloud
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Time |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SequenceNumber", wireType)
			}
			m.SequenceNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCloud
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SequenceNumber |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCloud(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCloud
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCloud
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *VizierHeartbeatAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCloud
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: VizierHeartbeatAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VizierHeartbeatAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCloud
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= VizierHeartbeatAck_HeartbeatStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			m.Time = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCloud
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Time |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SequenceNumber", wireType)
			}
			m.SequenceNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCloud
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SequenceNumber |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrorMessage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCloud
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCloud
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCloud
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ErrorMessage = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCloud(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCloud
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCloud
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCloud(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCloud
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCloud
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCloud
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthCloud
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthCloud
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCloud
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipCloud(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthCloud
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthCloud = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCloud   = fmt.Errorf("proto: integer overflow")
)
