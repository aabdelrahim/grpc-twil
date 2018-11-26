// Code generated by protoc-gen-go. DO NOT EDIT.
// source: twil.proto

package twil

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RecognizeRequest struct {
	Audio                []byte   `protobuf:"bytes,1,opt,name=audio,proto3" json:"audio,omitempty"`
	Language             string   `protobuf:"bytes,2,opt,name=language,proto3" json:"language,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecognizeRequest) Reset()         { *m = RecognizeRequest{} }
func (m *RecognizeRequest) String() string { return proto.CompactTextString(m) }
func (*RecognizeRequest) ProtoMessage()    {}
func (*RecognizeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5bb3739b531c9e5, []int{0}
}

func (m *RecognizeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecognizeRequest.Unmarshal(m, b)
}
func (m *RecognizeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecognizeRequest.Marshal(b, m, deterministic)
}
func (m *RecognizeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecognizeRequest.Merge(m, src)
}
func (m *RecognizeRequest) XXX_Size() int {
	return xxx_messageInfo_RecognizeRequest.Size(m)
}
func (m *RecognizeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RecognizeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RecognizeRequest proto.InternalMessageInfo

func (m *RecognizeRequest) GetAudio() []byte {
	if m != nil {
		return m.Audio
	}
	return nil
}

func (m *RecognizeRequest) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

type RecognizeResponse struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecognizeResponse) Reset()         { *m = RecognizeResponse{} }
func (m *RecognizeResponse) String() string { return proto.CompactTextString(m) }
func (*RecognizeResponse) ProtoMessage()    {}
func (*RecognizeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5bb3739b531c9e5, []int{1}
}

func (m *RecognizeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecognizeResponse.Unmarshal(m, b)
}
func (m *RecognizeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecognizeResponse.Marshal(b, m, deterministic)
}
func (m *RecognizeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecognizeResponse.Merge(m, src)
}
func (m *RecognizeResponse) XXX_Size() int {
	return xxx_messageInfo_RecognizeResponse.Size(m)
}
func (m *RecognizeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RecognizeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RecognizeResponse proto.InternalMessageInfo

func (m *RecognizeResponse) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*RecognizeRequest)(nil), "twil.RecognizeRequest")
	proto.RegisterType((*RecognizeResponse)(nil), "twil.RecognizeResponse")
}

func init() { proto.RegisterFile("twil.proto", fileDescriptor_f5bb3739b531c9e5) }

var fileDescriptor_f5bb3739b531c9e5 = []byte{
	// 167 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x29, 0xcf, 0xcc,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0x5c, 0xb8, 0x04, 0x82, 0x52,
	0x93, 0xf3, 0xd3, 0xf3, 0x32, 0xab, 0x52, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x44,
	0xb8, 0x58, 0x13, 0x4b, 0x53, 0x32, 0xf3, 0x25, 0x18, 0x15, 0x18, 0x35, 0x78, 0x82, 0x20, 0x1c,
	0x21, 0x29, 0x2e, 0x8e, 0x9c, 0xc4, 0xbc, 0xf4, 0xd2, 0xc4, 0xf4, 0x54, 0x09, 0x26, 0x05, 0x46,
	0x0d, 0xce, 0x20, 0x38, 0x5f, 0x49, 0x9d, 0x4b, 0x10, 0xc9, 0x94, 0xe2, 0x82, 0xfc, 0xbc, 0xe2,
	0x54, 0x21, 0x21, 0x2e, 0x96, 0x92, 0xd4, 0x8a, 0x12, 0xb0, 0x29, 0x9c, 0x41, 0x60, 0xb6, 0x91,
	0x1f, 0x17, 0x4f, 0x70, 0x41, 0x6a, 0x6a, 0x72, 0x46, 0x48, 0x7e, 0x48, 0x6a, 0x45, 0x89, 0x90,
	0x1d, 0x17, 0x27, 0x5c, 0xa3, 0x90, 0x98, 0x1e, 0xd8, 0x79, 0xe8, 0xee, 0x91, 0x12, 0xc7, 0x10,
	0x87, 0xd8, 0xa0, 0xc4, 0x90, 0xc4, 0x06, 0xf6, 0x8b, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x83,
	0xe8, 0x5b, 0xda, 0xd9, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SpeechToTextClient is the client API for SpeechToText service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SpeechToTextClient interface {
	Recognize(ctx context.Context, in *RecognizeRequest, opts ...grpc.CallOption) (*RecognizeResponse, error)
}

type speechToTextClient struct {
	cc *grpc.ClientConn
}

func NewSpeechToTextClient(cc *grpc.ClientConn) SpeechToTextClient {
	return &speechToTextClient{cc}
}

func (c *speechToTextClient) Recognize(ctx context.Context, in *RecognizeRequest, opts ...grpc.CallOption) (*RecognizeResponse, error) {
	out := new(RecognizeResponse)
	err := c.cc.Invoke(ctx, "/twil.SpeechToText/Recognize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SpeechToTextServer is the server API for SpeechToText service.
type SpeechToTextServer interface {
	Recognize(context.Context, *RecognizeRequest) (*RecognizeResponse, error)
}

func RegisterSpeechToTextServer(s *grpc.Server, srv SpeechToTextServer) {
	s.RegisterService(&_SpeechToText_serviceDesc, srv)
}

func _SpeechToText_Recognize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecognizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpeechToTextServer).Recognize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/twil.SpeechToText/Recognize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpeechToTextServer).Recognize(ctx, req.(*RecognizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SpeechToText_serviceDesc = grpc.ServiceDesc{
	ServiceName: "twil.SpeechToText",
	HandlerType: (*SpeechToTextServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Recognize",
			Handler:    _SpeechToText_Recognize_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "twil.proto",
}
