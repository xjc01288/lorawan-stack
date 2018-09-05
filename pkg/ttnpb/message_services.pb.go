// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/message_services.proto

package ttnpb // import "go.thethings.network/lorawan-stack/pkg/ttnpb"

import proto "github.com/gogo/protobuf/proto"
import golang_proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import context "context"
import grpc "google.golang.org/grpc"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ProcessUplinkMessageRequest struct {
	EndDeviceVersionIdentifiers EndDeviceVersionIdentifiers `protobuf:"bytes,1,opt,name=end_device_version_identifiers,json=endDeviceVersionIdentifiers" json:"end_device_version_identifiers"`
	Message                     UplinkMessage               `protobuf:"bytes,2,opt,name=message" json:"message"`
	Parameter                   string                      `protobuf:"bytes,3,opt,name=parameter,proto3" json:"parameter,omitempty"`
	XXX_NoUnkeyedLiteral        struct{}                    `json:"-"`
	XXX_sizecache               int32                       `json:"-"`
}

func (m *ProcessUplinkMessageRequest) Reset()      { *m = ProcessUplinkMessageRequest{} }
func (*ProcessUplinkMessageRequest) ProtoMessage() {}
func (*ProcessUplinkMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_services_b3ffb3b59e8cb296, []int{0}
}
func (m *ProcessUplinkMessageRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProcessUplinkMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProcessUplinkMessageRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *ProcessUplinkMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessUplinkMessageRequest.Merge(dst, src)
}
func (m *ProcessUplinkMessageRequest) XXX_Size() int {
	return m.Size()
}
func (m *ProcessUplinkMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessUplinkMessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessUplinkMessageRequest proto.InternalMessageInfo

func (m *ProcessUplinkMessageRequest) GetEndDeviceVersionIdentifiers() EndDeviceVersionIdentifiers {
	if m != nil {
		return m.EndDeviceVersionIdentifiers
	}
	return EndDeviceVersionIdentifiers{}
}

func (m *ProcessUplinkMessageRequest) GetMessage() UplinkMessage {
	if m != nil {
		return m.Message
	}
	return UplinkMessage{}
}

func (m *ProcessUplinkMessageRequest) GetParameter() string {
	if m != nil {
		return m.Parameter
	}
	return ""
}

type ProcessDownlinkMessageRequest struct {
	EndDeviceVersionIdentifiers EndDeviceVersionIdentifiers `protobuf:"bytes,1,opt,name=end_device_version_identifiers,json=endDeviceVersionIdentifiers" json:"end_device_version_identifiers"`
	Message                     DownlinkMessage             `protobuf:"bytes,2,opt,name=message" json:"message"`
	Parameter                   string                      `protobuf:"bytes,3,opt,name=parameter,proto3" json:"parameter,omitempty"`
	XXX_NoUnkeyedLiteral        struct{}                    `json:"-"`
	XXX_sizecache               int32                       `json:"-"`
}

func (m *ProcessDownlinkMessageRequest) Reset()      { *m = ProcessDownlinkMessageRequest{} }
func (*ProcessDownlinkMessageRequest) ProtoMessage() {}
func (*ProcessDownlinkMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_services_b3ffb3b59e8cb296, []int{1}
}
func (m *ProcessDownlinkMessageRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProcessDownlinkMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProcessDownlinkMessageRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *ProcessDownlinkMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessDownlinkMessageRequest.Merge(dst, src)
}
func (m *ProcessDownlinkMessageRequest) XXX_Size() int {
	return m.Size()
}
func (m *ProcessDownlinkMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessDownlinkMessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessDownlinkMessageRequest proto.InternalMessageInfo

func (m *ProcessDownlinkMessageRequest) GetEndDeviceVersionIdentifiers() EndDeviceVersionIdentifiers {
	if m != nil {
		return m.EndDeviceVersionIdentifiers
	}
	return EndDeviceVersionIdentifiers{}
}

func (m *ProcessDownlinkMessageRequest) GetMessage() DownlinkMessage {
	if m != nil {
		return m.Message
	}
	return DownlinkMessage{}
}

func (m *ProcessDownlinkMessageRequest) GetParameter() string {
	if m != nil {
		return m.Parameter
	}
	return ""
}

func init() {
	proto.RegisterType((*ProcessUplinkMessageRequest)(nil), "ttn.lorawan.v3.ProcessUplinkMessageRequest")
	golang_proto.RegisterType((*ProcessUplinkMessageRequest)(nil), "ttn.lorawan.v3.ProcessUplinkMessageRequest")
	proto.RegisterType((*ProcessDownlinkMessageRequest)(nil), "ttn.lorawan.v3.ProcessDownlinkMessageRequest")
	golang_proto.RegisterType((*ProcessDownlinkMessageRequest)(nil), "ttn.lorawan.v3.ProcessDownlinkMessageRequest")
}
func (this *ProcessUplinkMessageRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ProcessUplinkMessageRequest)
	if !ok {
		that2, ok := that.(ProcessUplinkMessageRequest)
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
	if !this.EndDeviceVersionIdentifiers.Equal(&that1.EndDeviceVersionIdentifiers) {
		return false
	}
	if !this.Message.Equal(&that1.Message) {
		return false
	}
	if this.Parameter != that1.Parameter {
		return false
	}
	return true
}
func (this *ProcessDownlinkMessageRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ProcessDownlinkMessageRequest)
	if !ok {
		that2, ok := that.(ProcessDownlinkMessageRequest)
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
	if !this.EndDeviceVersionIdentifiers.Equal(&that1.EndDeviceVersionIdentifiers) {
		return false
	}
	if !this.Message.Equal(&that1.Message) {
		return false
	}
	if this.Parameter != that1.Parameter {
		return false
	}
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for UplinkMessageProcessor service

type UplinkMessageProcessorClient interface {
	Process(ctx context.Context, in *ProcessUplinkMessageRequest, opts ...grpc.CallOption) (*UplinkMessage, error)
}

type uplinkMessageProcessorClient struct {
	cc *grpc.ClientConn
}

func NewUplinkMessageProcessorClient(cc *grpc.ClientConn) UplinkMessageProcessorClient {
	return &uplinkMessageProcessorClient{cc}
}

func (c *uplinkMessageProcessorClient) Process(ctx context.Context, in *ProcessUplinkMessageRequest, opts ...grpc.CallOption) (*UplinkMessage, error) {
	out := new(UplinkMessage)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.UplinkMessageProcessor/Process", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UplinkMessageProcessor service

type UplinkMessageProcessorServer interface {
	Process(context.Context, *ProcessUplinkMessageRequest) (*UplinkMessage, error)
}

func RegisterUplinkMessageProcessorServer(s *grpc.Server, srv UplinkMessageProcessorServer) {
	s.RegisterService(&_UplinkMessageProcessor_serviceDesc, srv)
}

func _UplinkMessageProcessor_Process_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessUplinkMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UplinkMessageProcessorServer).Process(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.UplinkMessageProcessor/Process",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UplinkMessageProcessorServer).Process(ctx, req.(*ProcessUplinkMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UplinkMessageProcessor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.UplinkMessageProcessor",
	HandlerType: (*UplinkMessageProcessorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Process",
			Handler:    _UplinkMessageProcessor_Process_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/message_services.proto",
}

// Client API for DownlinkMessageProcessor service

type DownlinkMessageProcessorClient interface {
	Process(ctx context.Context, in *ProcessDownlinkMessageRequest, opts ...grpc.CallOption) (*DownlinkMessage, error)
}

type downlinkMessageProcessorClient struct {
	cc *grpc.ClientConn
}

func NewDownlinkMessageProcessorClient(cc *grpc.ClientConn) DownlinkMessageProcessorClient {
	return &downlinkMessageProcessorClient{cc}
}

func (c *downlinkMessageProcessorClient) Process(ctx context.Context, in *ProcessDownlinkMessageRequest, opts ...grpc.CallOption) (*DownlinkMessage, error) {
	out := new(DownlinkMessage)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.DownlinkMessageProcessor/Process", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DownlinkMessageProcessor service

type DownlinkMessageProcessorServer interface {
	Process(context.Context, *ProcessDownlinkMessageRequest) (*DownlinkMessage, error)
}

func RegisterDownlinkMessageProcessorServer(s *grpc.Server, srv DownlinkMessageProcessorServer) {
	s.RegisterService(&_DownlinkMessageProcessor_serviceDesc, srv)
}

func _DownlinkMessageProcessor_Process_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessDownlinkMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DownlinkMessageProcessorServer).Process(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.DownlinkMessageProcessor/Process",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DownlinkMessageProcessorServer).Process(ctx, req.(*ProcessDownlinkMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DownlinkMessageProcessor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.DownlinkMessageProcessor",
	HandlerType: (*DownlinkMessageProcessorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Process",
			Handler:    _DownlinkMessageProcessor_Process_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/message_services.proto",
}

func (m *ProcessUplinkMessageRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProcessUplinkMessageRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintMessageServices(dAtA, i, uint64(m.EndDeviceVersionIdentifiers.Size()))
	n1, err := m.EndDeviceVersionIdentifiers.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x12
	i++
	i = encodeVarintMessageServices(dAtA, i, uint64(m.Message.Size()))
	n2, err := m.Message.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	if len(m.Parameter) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintMessageServices(dAtA, i, uint64(len(m.Parameter)))
		i += copy(dAtA[i:], m.Parameter)
	}
	return i, nil
}

func (m *ProcessDownlinkMessageRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProcessDownlinkMessageRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintMessageServices(dAtA, i, uint64(m.EndDeviceVersionIdentifiers.Size()))
	n3, err := m.EndDeviceVersionIdentifiers.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	dAtA[i] = 0x12
	i++
	i = encodeVarintMessageServices(dAtA, i, uint64(m.Message.Size()))
	n4, err := m.Message.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	if len(m.Parameter) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintMessageServices(dAtA, i, uint64(len(m.Parameter)))
		i += copy(dAtA[i:], m.Parameter)
	}
	return i, nil
}

func encodeVarintMessageServices(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedProcessUplinkMessageRequest(r randyMessageServices, easy bool) *ProcessUplinkMessageRequest {
	this := &ProcessUplinkMessageRequest{}
	v1 := NewPopulatedEndDeviceVersionIdentifiers(r, easy)
	this.EndDeviceVersionIdentifiers = *v1
	v2 := NewPopulatedUplinkMessage(r, easy)
	this.Message = *v2
	this.Parameter = randStringMessageServices(r)
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedProcessDownlinkMessageRequest(r randyMessageServices, easy bool) *ProcessDownlinkMessageRequest {
	this := &ProcessDownlinkMessageRequest{}
	v3 := NewPopulatedEndDeviceVersionIdentifiers(r, easy)
	this.EndDeviceVersionIdentifiers = *v3
	v4 := NewPopulatedDownlinkMessage(r, easy)
	this.Message = *v4
	this.Parameter = randStringMessageServices(r)
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyMessageServices interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneMessageServices(r randyMessageServices) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringMessageServices(r randyMessageServices) string {
	v5 := r.Intn(100)
	tmps := make([]rune, v5)
	for i := 0; i < v5; i++ {
		tmps[i] = randUTF8RuneMessageServices(r)
	}
	return string(tmps)
}
func randUnrecognizedMessageServices(r randyMessageServices, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldMessageServices(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldMessageServices(dAtA []byte, r randyMessageServices, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateMessageServices(dAtA, uint64(key))
		v6 := r.Int63()
		if r.Intn(2) == 0 {
			v6 *= -1
		}
		dAtA = encodeVarintPopulateMessageServices(dAtA, uint64(v6))
	case 1:
		dAtA = encodeVarintPopulateMessageServices(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateMessageServices(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateMessageServices(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateMessageServices(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateMessageServices(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(v&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *ProcessUplinkMessageRequest) Size() (n int) {
	var l int
	_ = l
	l = m.EndDeviceVersionIdentifiers.Size()
	n += 1 + l + sovMessageServices(uint64(l))
	l = m.Message.Size()
	n += 1 + l + sovMessageServices(uint64(l))
	l = len(m.Parameter)
	if l > 0 {
		n += 1 + l + sovMessageServices(uint64(l))
	}
	return n
}

func (m *ProcessDownlinkMessageRequest) Size() (n int) {
	var l int
	_ = l
	l = m.EndDeviceVersionIdentifiers.Size()
	n += 1 + l + sovMessageServices(uint64(l))
	l = m.Message.Size()
	n += 1 + l + sovMessageServices(uint64(l))
	l = len(m.Parameter)
	if l > 0 {
		n += 1 + l + sovMessageServices(uint64(l))
	}
	return n
}

func sovMessageServices(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMessageServices(x uint64) (n int) {
	return sovMessageServices((x << 1) ^ uint64((int64(x) >> 63)))
}
func (this *ProcessUplinkMessageRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ProcessUplinkMessageRequest{`,
		`EndDeviceVersionIdentifiers:` + strings.Replace(strings.Replace(this.EndDeviceVersionIdentifiers.String(), "EndDeviceVersionIdentifiers", "EndDeviceVersionIdentifiers", 1), `&`, ``, 1) + `,`,
		`Message:` + strings.Replace(strings.Replace(this.Message.String(), "UplinkMessage", "UplinkMessage", 1), `&`, ``, 1) + `,`,
		`Parameter:` + fmt.Sprintf("%v", this.Parameter) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ProcessDownlinkMessageRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ProcessDownlinkMessageRequest{`,
		`EndDeviceVersionIdentifiers:` + strings.Replace(strings.Replace(this.EndDeviceVersionIdentifiers.String(), "EndDeviceVersionIdentifiers", "EndDeviceVersionIdentifiers", 1), `&`, ``, 1) + `,`,
		`Message:` + strings.Replace(strings.Replace(this.Message.String(), "DownlinkMessage", "DownlinkMessage", 1), `&`, ``, 1) + `,`,
		`Parameter:` + fmt.Sprintf("%v", this.Parameter) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringMessageServices(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ProcessUplinkMessageRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessageServices
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ProcessUplinkMessageRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProcessUplinkMessageRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndDeviceVersionIdentifiers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessageServices
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessageServices
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.EndDeviceVersionIdentifiers.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessageServices
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessageServices
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Message.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Parameter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessageServices
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessageServices
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Parameter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessageServices(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessageServices
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
func (m *ProcessDownlinkMessageRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessageServices
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ProcessDownlinkMessageRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProcessDownlinkMessageRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndDeviceVersionIdentifiers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessageServices
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessageServices
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.EndDeviceVersionIdentifiers.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessageServices
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessageServices
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Message.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Parameter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessageServices
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessageServices
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Parameter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessageServices(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessageServices
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
func skipMessageServices(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessageServices
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
					return 0, ErrIntOverflowMessageServices
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
					return 0, ErrIntOverflowMessageServices
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthMessageServices
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMessageServices
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
				next, err := skipMessageServices(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthMessageServices = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessageServices   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("lorawan-stack/api/message_services.proto", fileDescriptor_message_services_b3ffb3b59e8cb296)
}
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/message_services.proto", fileDescriptor_message_services_b3ffb3b59e8cb296)
}

var fileDescriptor_message_services_b3ffb3b59e8cb296 = []byte{
	// 474 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x93, 0x31, 0x6c, 0xd3, 0x40,
	0x14, 0x86, 0xef, 0x01, 0xa2, 0xea, 0x21, 0x31, 0x78, 0x40, 0x51, 0x4a, 0x5f, 0xa3, 0x4c, 0x91,
	0x20, 0xb6, 0x94, 0xce, 0x08, 0x09, 0x95, 0x81, 0x01, 0x09, 0x45, 0x2a, 0x12, 0x2c, 0x91, 0x93,
	0xbc, 0x3a, 0xa7, 0x34, 0x77, 0xe6, 0xee, 0x92, 0xac, 0x1d, 0x3b, 0x32, 0x32, 0x22, 0xa6, 0x8e,
	0x1d, 0x3b, 0x76, 0xcc, 0xd8, 0xb1, 0x13, 0xaa, 0xef, 0x96, 0x8e, 0x15, 0x03, 0xea, 0x88, 0x70,
	0x8c, 0xa2, 0xb8, 0xa4, 0x88, 0xad, 0x9b, 0x9f, 0xfd, 0xfe, 0xfb, 0xfe, 0xff, 0x9d, 0x1f, 0x6f,
	0xec, 0x2b, 0x1d, 0x4f, 0x63, 0xd9, 0x34, 0x36, 0xee, 0x0d, 0xa3, 0x38, 0x15, 0xd1, 0x88, 0x8c,
	0x89, 0x13, 0xea, 0x18, 0xd2, 0x13, 0xd1, 0x23, 0x13, 0xa6, 0x5a, 0x59, 0x15, 0x3c, 0xb6, 0x56,
	0x86, 0x45, 0x77, 0x38, 0xd9, 0xae, 0x36, 0x13, 0x61, 0x07, 0xe3, 0x6e, 0xd8, 0x53, 0xa3, 0x28,
	0x51, 0x89, 0x8a, 0xf2, 0xb6, 0xee, 0x78, 0x2f, 0xaf, 0xf2, 0x22, 0x7f, 0x9a, 0xcb, 0xab, 0xf5,
	0x9b, 0x20, 0x92, 0xfd, 0x4e, 0x9f, 0x7e, 0x33, 0x8a, 0x9e, 0xda, 0x4a, 0x33, 0x85, 0x89, 0xfa,
	0x0f, 0xe0, 0x1b, 0xef, 0xb4, 0xea, 0x91, 0x31, 0xbb, 0xe9, 0xbe, 0x90, 0xc3, 0xb7, 0xf3, 0xef,
	0x6d, 0xfa, 0x34, 0x26, 0x63, 0x83, 0x09, 0xc7, 0xc5, 0xa9, 0x9d, 0x09, 0x69, 0x23, 0x94, 0xec,
	0x88, 0x3e, 0x49, 0x2b, 0xf6, 0x04, 0x69, 0x53, 0x81, 0x1a, 0x34, 0x1e, 0xb5, 0x9e, 0x85, 0xcb,
	0x69, 0xc2, 0xd7, 0xb2, 0xbf, 0x93, 0x8b, 0xde, 0xcf, 0x35, 0x6f, 0x16, 0x92, 0x57, 0x0f, 0x66,
	0xdf, 0xb7, 0x58, 0x7b, 0x83, 0x56, 0xb7, 0x04, 0x2f, 0xf8, 0x5a, 0xe1, 0xb4, 0x72, 0x2f, 0x07,
	0x6c, 0x96, 0x01, 0x4b, 0x76, 0x8b, 0x23, 0xff, 0x68, 0x82, 0xa7, 0x7c, 0x3d, 0x8d, 0x75, 0x3c,
	0x22, 0x4b, 0xba, 0x72, 0xbf, 0x06, 0x8d, 0xf5, 0xf6, 0xe2, 0x45, 0xfd, 0x27, 0xf0, 0xcd, 0x22,
	0xf4, 0x8e, 0x9a, 0xca, 0x3b, 0x14, 0xfb, 0x65, 0x39, 0xf6, 0x56, 0x19, 0x50, 0x32, 0xfc, 0x5f,
	0xc1, 0x5b, 0x8a, 0x3f, 0x59, 0x1a, 0x5b, 0x31, 0x04, 0xa5, 0x83, 0x5d, 0xbe, 0x56, 0x14, 0xc1,
	0x8d, 0x4c, 0xb7, 0xfc, 0x1f, 0xd5, 0xdb, 0xaf, 0xa5, 0x35, 0xe6, 0x95, 0x92, 0xe1, 0x05, 0xf2,
	0xc3, 0x02, 0xd9, 0x5c, 0x81, 0xfc, 0xfb, 0xed, 0x54, 0xff, 0x39, 0x94, 0x6f, 0x30, 0xcb, 0x10,
	0xce, 0x32, 0x84, 0xf3, 0x0c, 0xd9, 0x45, 0x86, 0xec, 0x32, 0x43, 0x76, 0x95, 0x21, 0xbb, 0xce,
	0x10, 0x0e, 0x1c, 0xc2, 0xa1, 0x43, 0x76, 0xe4, 0x10, 0x8e, 0x1d, 0xb2, 0x13, 0x87, 0xec, 0xd4,
	0x21, 0x9b, 0x39, 0x84, 0x33, 0x87, 0x70, 0xee, 0x90, 0x5d, 0x38, 0x84, 0x4b, 0x87, 0xec, 0xca,
	0x21, 0x5c, 0x3b, 0x64, 0x07, 0x1e, 0xd9, 0xa1, 0x47, 0xf8, 0xec, 0x91, 0x7d, 0xf1, 0x08, 0x5f,
	0x3d, 0xb2, 0x23, 0x8f, 0xec, 0xd8, 0x23, 0x9c, 0x78, 0x84, 0x53, 0x8f, 0xf0, 0xf1, 0x79, 0xa2,
	0x42, 0x3b, 0x20, 0x3b, 0x10, 0x32, 0x31, 0xa1, 0x24, 0x3b, 0x55, 0x7a, 0x18, 0x2d, 0xef, 0x60,
	0x3a, 0x4c, 0x22, 0x6b, 0x65, 0xda, 0xed, 0x3e, 0xcc, 0x37, 0x70, 0xfb, 0x57, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x4d, 0xe6, 0x50, 0xd4, 0x32, 0x04, 0x00, 0x00,
}