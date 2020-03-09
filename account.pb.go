// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account.proto

package account_grpc

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

// Payment message defines the payment amount being received by sent by a client
type Payment struct {
	//    all scalar types
	Amount               int32    `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	AccountHolder        []byte   `protobuf:"bytes,2,opt,name=account_holder,json=accountHolder,proto3" json:"account_holder,omitempty"`
	SenderBank           string   `protobuf:"bytes,3,opt,name=sender_bank,json=senderBank,proto3" json:"sender_bank,omitempty"`
	PaymentId            string   `protobuf:"bytes,4,opt,name=payment_id,json=paymentId,proto3" json:"payment_id,omitempty"`
	AccountReceiver      []byte   `protobuf:"bytes,5,opt,name=account_receiver,json=accountReceiver,proto3" json:"account_receiver,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Payment) Reset()         { *m = Payment{} }
func (m *Payment) String() string { return proto.CompactTextString(m) }
func (*Payment) ProtoMessage()    {}
func (*Payment) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{0}
}

func (m *Payment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Payment.Unmarshal(m, b)
}
func (m *Payment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Payment.Marshal(b, m, deterministic)
}
func (m *Payment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Payment.Merge(m, src)
}
func (m *Payment) XXX_Size() int {
	return xxx_messageInfo_Payment.Size(m)
}
func (m *Payment) XXX_DiscardUnknown() {
	xxx_messageInfo_Payment.DiscardUnknown(m)
}

var xxx_messageInfo_Payment proto.InternalMessageInfo

func (m *Payment) GetAmount() int32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Payment) GetAccountHolder() []byte {
	if m != nil {
		return m.AccountHolder
	}
	return nil
}

func (m *Payment) GetSenderBank() string {
	if m != nil {
		return m.SenderBank
	}
	return ""
}

func (m *Payment) GetPaymentId() string {
	if m != nil {
		return m.PaymentId
	}
	return ""
}

func (m *Payment) GetAccountReceiver() []byte {
	if m != nil {
		return m.AccountReceiver
	}
	return nil
}

// Reference message defines the receiver bank and whether the payment was processed
type Reference struct {
	//    all scalar types
	ReceiverBank         string   `protobuf:"bytes,1,opt,name=receiver_bank,json=receiverBank,proto3" json:"receiver_bank,omitempty"`
	Confirmed            bool     `protobuf:"varint,2,opt,name=confirmed,proto3" json:"confirmed,omitempty"`
	AccountPaidTo        []byte   `protobuf:"bytes,3,opt,name=account_paid_to,json=accountPaidTo,proto3" json:"account_paid_to,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Reference) Reset()         { *m = Reference{} }
func (m *Reference) String() string { return proto.CompactTextString(m) }
func (*Reference) ProtoMessage()    {}
func (*Reference) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{1}
}

func (m *Reference) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Reference.Unmarshal(m, b)
}
func (m *Reference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Reference.Marshal(b, m, deterministic)
}
func (m *Reference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reference.Merge(m, src)
}
func (m *Reference) XXX_Size() int {
	return xxx_messageInfo_Reference.Size(m)
}
func (m *Reference) XXX_DiscardUnknown() {
	xxx_messageInfo_Reference.DiscardUnknown(m)
}

var xxx_messageInfo_Reference proto.InternalMessageInfo

func (m *Reference) GetReceiverBank() string {
	if m != nil {
		return m.ReceiverBank
	}
	return ""
}

func (m *Reference) GetConfirmed() bool {
	if m != nil {
		return m.Confirmed
	}
	return false
}

func (m *Reference) GetAccountPaidTo() []byte {
	if m != nil {
		return m.AccountPaidTo
	}
	return nil
}

func init() {
	proto.RegisterType((*Payment)(nil), "account_grpc.Payment")
	proto.RegisterType((*Reference)(nil), "account_grpc.Reference")
}

func init() { proto.RegisterFile("account.proto", fileDescriptor_8e28828dcb8d24f0) }

var fileDescriptor_8e28828dcb8d24f0 = []byte{
	// 273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xc1, 0x4e, 0xf4, 0x20,
	0x10, 0xc7, 0x3f, 0x3e, 0xdd, 0xd5, 0xce, 0xb6, 0x6a, 0x48, 0xd4, 0xc6, 0x68, 0x6c, 0x6a, 0x34,
	0xf5, 0xd2, 0x83, 0x3e, 0x80, 0xd1, 0x93, 0xde, 0x1a, 0xf4, 0xde, 0xb0, 0x30, 0xab, 0x64, 0x2d,
	0x34, 0x2c, 0xdb, 0xc4, 0xe7, 0xf2, 0x05, 0x8d, 0x14, 0x5c, 0x3d, 0xce, 0x6f, 0x60, 0x7e, 0x7f,
	0x06, 0xc8, 0xb8, 0x10, 0x66, 0xad, 0x5d, 0xdd, 0x5b, 0xe3, 0x0c, 0x4d, 0x43, 0xd9, 0xbe, 0xda,
	0x5e, 0x94, 0x9f, 0x04, 0x76, 0x1a, 0xfe, 0xd1, 0xa1, 0x76, 0xf4, 0x08, 0xa6, 0xbc, 0xfb, 0x6e,
	0xe5, 0xa4, 0x20, 0xd5, 0x84, 0x85, 0x8a, 0x5e, 0xc2, 0x5e, 0xbc, 0xf3, 0x66, 0xde, 0x25, 0xda,
	0xfc, 0x7f, 0x41, 0xaa, 0x94, 0xc5, 0xc1, 0x8f, 0x1e, 0xd2, 0x73, 0x98, 0xad, 0x50, 0x4b, 0xb4,
	0xed, 0x9c, 0xeb, 0x65, 0xbe, 0x55, 0x90, 0x2a, 0x61, 0x30, 0xa2, 0x07, 0xae, 0x97, 0xf4, 0x0c,
	0xa0, 0x1f, 0x55, 0xad, 0x92, 0xf9, 0xb6, 0xef, 0x27, 0x81, 0x3c, 0x49, 0x7a, 0x0d, 0x07, 0x51,
	0x63, 0x51, 0xa0, 0x1a, 0xd0, 0xe6, 0x13, 0x2f, 0xda, 0x0f, 0x9c, 0x05, 0x5c, 0x0e, 0x90, 0x30,
	0x5c, 0xa0, 0x45, 0x2d, 0x90, 0x5e, 0x40, 0x16, 0xcf, 0x8f, 0x66, 0xe2, 0x27, 0xa7, 0x11, 0x7a,
	0xf7, 0x29, 0x24, 0xc2, 0xe8, 0x85, 0xb2, 0x1d, 0x4a, 0x1f, 0x7f, 0x97, 0x6d, 0x00, 0xbd, 0x82,
	0xa8, 0x68, 0x7b, 0xae, 0x64, 0xeb, 0x8c, 0x8f, 0xbf, 0x79, 0x62, 0xc3, 0x95, 0x7c, 0x31, 0x37,
	0x0d, 0x64, 0xf7, 0x21, 0x8a, 0x59, 0x3b, 0x5c, 0xd1, 0x3b, 0x98, 0x3d, 0xf3, 0x01, 0xe3, 0x06,
	0x0f, 0xeb, 0xdf, 0xcb, 0xad, 0x03, 0x3e, 0x39, 0xfe, 0x8b, 0x7f, 0xa2, 0x97, 0xff, 0xe6, 0x53,
	0xff, 0x29, 0xb7, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8b, 0x5e, 0x7c, 0xa3, 0xa5, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccountRoutesClient is the client API for AccountRoutes service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountRoutesClient interface {
	//    SavePayment rpc responsible for processing payment with encrypted user data
	SavePayment(ctx context.Context, in *Payment, opts ...grpc.CallOption) (*Reference, error)
}

type accountRoutesClient struct {
	cc *grpc.ClientConn
}

func NewAccountRoutesClient(cc *grpc.ClientConn) AccountRoutesClient {
	return &accountRoutesClient{cc}
}

func (c *accountRoutesClient) SavePayment(ctx context.Context, in *Payment, opts ...grpc.CallOption) (*Reference, error) {
	out := new(Reference)
	err := c.cc.Invoke(ctx, "/account_grpc.AccountRoutes/SavePayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountRoutesServer is the server API for AccountRoutes service.
type AccountRoutesServer interface {
	//    SavePayment rpc responsible for processing payment with encrypted user data
	SavePayment(context.Context, *Payment) (*Reference, error)
}

// UnimplementedAccountRoutesServer can be embedded to have forward compatible implementations.
type UnimplementedAccountRoutesServer struct {
}

func (*UnimplementedAccountRoutesServer) SavePayment(ctx context.Context, req *Payment) (*Reference, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SavePayment not implemented")
}

func RegisterAccountRoutesServer(s *grpc.Server, srv AccountRoutesServer) {
	s.RegisterService(&_AccountRoutes_serviceDesc, srv)
}

func _AccountRoutes_SavePayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Payment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountRoutesServer).SavePayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account_grpc.AccountRoutes/SavePayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountRoutesServer).SavePayment(ctx, req.(*Payment))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccountRoutes_serviceDesc = grpc.ServiceDesc{
	ServiceName: "account_grpc.AccountRoutes",
	HandlerType: (*AccountRoutesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SavePayment",
			Handler:    _AccountRoutes_SavePayment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "account.proto",
}
