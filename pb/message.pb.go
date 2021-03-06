// Code generated by protoc-gen-go.
// source: message.proto
// DO NOT EDIT!

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf1 "github.com/golang/protobuf/ptypes/any"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Message_MessageType int32

const (
	Message_PING               Message_MessageType = 0
	Message_CHAT               Message_MessageType = 1
	Message_FOLLOW             Message_MessageType = 2
	Message_UNFOLLOW           Message_MessageType = 3
	Message_ORDER              Message_MessageType = 4
	Message_ORDER_REJECT       Message_MessageType = 5
	Message_ORDER_CANCEL       Message_MessageType = 6
	Message_ORDER_CONFIRMATION Message_MessageType = 7
	Message_ORDER_FULFILLMENT  Message_MessageType = 8
	Message_ORDER_COMPLETION   Message_MessageType = 9
	Message_DISPUTE_OPEN       Message_MessageType = 10
	Message_DISPUTE_UPDATE     Message_MessageType = 11
	Message_DISPUTE_CLOSE      Message_MessageType = 12
	Message_REFUND             Message_MessageType = 13
	Message_OFFLINE_ACK        Message_MessageType = 14
	Message_OFFLINE_RELAY      Message_MessageType = 15
	Message_MODERATOR_ADD      Message_MessageType = 16
	Message_MODERATOR_REMOVE   Message_MessageType = 17
	Message_ERROR              Message_MessageType = 500
)

var Message_MessageType_name = map[int32]string{
	0:   "PING",
	1:   "CHAT",
	2:   "FOLLOW",
	3:   "UNFOLLOW",
	4:   "ORDER",
	5:   "ORDER_REJECT",
	6:   "ORDER_CANCEL",
	7:   "ORDER_CONFIRMATION",
	8:   "ORDER_FULFILLMENT",
	9:   "ORDER_COMPLETION",
	10:  "DISPUTE_OPEN",
	11:  "DISPUTE_UPDATE",
	12:  "DISPUTE_CLOSE",
	13:  "REFUND",
	14:  "OFFLINE_ACK",
	15:  "OFFLINE_RELAY",
	16:  "MODERATOR_ADD",
	17:  "MODERATOR_REMOVE",
	500: "ERROR",
}
var Message_MessageType_value = map[string]int32{
	"PING":               0,
	"CHAT":               1,
	"FOLLOW":             2,
	"UNFOLLOW":           3,
	"ORDER":              4,
	"ORDER_REJECT":       5,
	"ORDER_CANCEL":       6,
	"ORDER_CONFIRMATION": 7,
	"ORDER_FULFILLMENT":  8,
	"ORDER_COMPLETION":   9,
	"DISPUTE_OPEN":       10,
	"DISPUTE_UPDATE":     11,
	"DISPUTE_CLOSE":      12,
	"REFUND":             13,
	"OFFLINE_ACK":        14,
	"OFFLINE_RELAY":      15,
	"MODERATOR_ADD":      16,
	"MODERATOR_REMOVE":   17,
	"ERROR":              500,
}

func (x Message_MessageType) String() string {
	return proto.EnumName(Message_MessageType_name, int32(x))
}
func (Message_MessageType) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 0} }

type Chat_Flag int32

const (
	Chat_MESSAGE Chat_Flag = 0
	Chat_TYPING  Chat_Flag = 1
	Chat_READ    Chat_Flag = 2
)

var Chat_Flag_name = map[int32]string{
	0: "MESSAGE",
	1: "TYPING",
	2: "READ",
}
var Chat_Flag_value = map[string]int32{
	"MESSAGE": 0,
	"TYPING":  1,
	"READ":    2,
}

func (x Chat_Flag) String() string {
	return proto.EnumName(Chat_Flag_name, int32(x))
}
func (Chat_Flag) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{2, 0} }

type Message struct {
	MessageType Message_MessageType   `protobuf:"varint,1,opt,name=messageType,enum=Message_MessageType" json:"messageType,omitempty"`
	Payload     *google_protobuf1.Any `protobuf:"bytes,2,opt,name=payload" json:"payload,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *Message) GetMessageType() Message_MessageType {
	if m != nil {
		return m.MessageType
	}
	return Message_PING
}

func (m *Message) GetPayload() *google_protobuf1.Any {
	if m != nil {
		return m.Payload
	}
	return nil
}

type Envelope struct {
	Message   *Message `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	Pubkey    []byte   `protobuf:"bytes,2,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	Signature []byte   `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *Envelope) Reset()                    { *m = Envelope{} }
func (m *Envelope) String() string            { return proto.CompactTextString(m) }
func (*Envelope) ProtoMessage()               {}
func (*Envelope) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *Envelope) GetMessage() *Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *Envelope) GetPubkey() []byte {
	if m != nil {
		return m.Pubkey
	}
	return nil
}

func (m *Envelope) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type Chat struct {
	MessageId string                     `protobuf:"bytes,1,opt,name=messageId" json:"messageId,omitempty"`
	Subject   string                     `protobuf:"bytes,2,opt,name=subject" json:"subject,omitempty"`
	Message   string                     `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=timestamp" json:"timestamp,omitempty"`
	Flag      Chat_Flag                  `protobuf:"varint,5,opt,name=flag,enum=Chat_Flag" json:"flag,omitempty"`
}

func (m *Chat) Reset()                    { *m = Chat{} }
func (m *Chat) String() string            { return proto.CompactTextString(m) }
func (*Chat) ProtoMessage()               {}
func (*Chat) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *Chat) GetMessageId() string {
	if m != nil {
		return m.MessageId
	}
	return ""
}

func (m *Chat) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *Chat) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Chat) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Chat) GetFlag() Chat_Flag {
	if m != nil {
		return m.Flag
	}
	return Chat_MESSAGE
}

func init() {
	proto.RegisterType((*Message)(nil), "Message")
	proto.RegisterType((*Envelope)(nil), "Envelope")
	proto.RegisterType((*Chat)(nil), "Chat")
	proto.RegisterEnum("Message_MessageType", Message_MessageType_name, Message_MessageType_value)
	proto.RegisterEnum("Chat_Flag", Chat_Flag_name, Chat_Flag_value)
}

func init() { proto.RegisterFile("message.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 539 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x92, 0xd1, 0x6e, 0xd3, 0x4c,
	0x10, 0x85, 0xeb, 0xc4, 0x89, 0xe3, 0x71, 0x9a, 0x6e, 0x56, 0xf9, 0x2b, 0xff, 0x15, 0x82, 0xca,
	0x57, 0xe5, 0xc6, 0x95, 0x82, 0x84, 0xb8, 0x5d, 0xec, 0x75, 0x31, 0xd8, 0xde, 0x68, 0xe3, 0x80,
	0xca, 0x4d, 0xe4, 0x10, 0xd7, 0x14, 0x92, 0xd8, 0x6a, 0x1c, 0xa4, 0x3c, 0x1e, 0x0f, 0xc0, 0x23,
	0xf0, 0x16, 0x3c, 0x00, 0xda, 0x8d, 0x97, 0x54, 0x70, 0xe5, 0x9d, 0xef, 0x1c, 0x9d, 0x19, 0x6b,
	0x06, 0x4e, 0xd7, 0xf9, 0x76, 0x9b, 0x15, 0xb9, 0x5b, 0x3d, 0x94, 0x75, 0x79, 0xf1, 0x7f, 0x51,
	0x96, 0xc5, 0x2a, 0xbf, 0x96, 0xd5, 0x62, 0x77, 0x77, 0x9d, 0x6d, 0xf6, 0x8d, 0xf4, 0xec, 0x6f,
	0xa9, 0xbe, 0x5f, 0xe7, 0xdb, 0x3a, 0x5b, 0x57, 0x07, 0x83, 0xf3, 0xbd, 0x0d, 0x46, 0x7c, 0x48,
	0xc3, 0x2f, 0xc1, 0x6a, 0x82, 0xd3, 0x7d, 0x95, 0xdb, 0xda, 0xa5, 0x76, 0x35, 0x18, 0x8f, 0xdc,
	0x46, 0x56, 0x5f, 0xa1, 0xf1, 0xc7, 0x46, 0xec, 0x82, 0x51, 0x65, 0xfb, 0x55, 0x99, 0x2d, 0xed,
	0xd6, 0xa5, 0x76, 0x65, 0x8d, 0x47, 0xee, 0xa1, 0xad, 0xab, 0xda, 0xba, 0x64, 0xb3, 0xe7, 0xca,
	0xe4, 0xfc, 0x68, 0x81, 0xf5, 0x28, 0x0c, 0xf7, 0x40, 0x9f, 0x84, 0xc9, 0x0d, 0x3a, 0x11, 0x2f,
	0xef, 0x0d, 0x49, 0x91, 0x86, 0x01, 0xba, 0x01, 0x8b, 0x22, 0xf6, 0x01, 0xb5, 0x70, 0x1f, 0x7a,
	0xb3, 0xa4, 0xa9, 0xda, 0xd8, 0x84, 0x0e, 0xe3, 0x3e, 0xe5, 0x48, 0xc7, 0x08, 0xfa, 0xf2, 0x39,
	0xe7, 0xf4, 0x2d, 0xf5, 0x52, 0xd4, 0x39, 0x12, 0x8f, 0x24, 0x1e, 0x8d, 0x50, 0x17, 0x9f, 0x03,
	0x6e, 0x08, 0x4b, 0x82, 0x90, 0xc7, 0x24, 0x0d, 0x59, 0x82, 0x0c, 0xfc, 0x1f, 0x0c, 0x0f, 0x3c,
	0x98, 0x45, 0x41, 0x18, 0x45, 0x31, 0x4d, 0x52, 0xd4, 0xc3, 0x23, 0x40, 0xca, 0x1e, 0x4f, 0x22,
	0x2a, 0xcd, 0xa6, 0x88, 0xf5, 0xc3, 0xe9, 0x64, 0x96, 0xd2, 0x39, 0x9b, 0xd0, 0x04, 0x01, 0xc6,
	0x30, 0x50, 0x64, 0x36, 0xf1, 0x49, 0x4a, 0x91, 0x85, 0x87, 0x70, 0xaa, 0x98, 0x17, 0xb1, 0x29,
	0x45, 0x7d, 0xf1, 0x1b, 0x9c, 0x06, 0xb3, 0xc4, 0x47, 0xa7, 0xf8, 0x0c, 0x2c, 0x16, 0x04, 0x51,
	0x98, 0xd0, 0x39, 0xf1, 0xde, 0xa1, 0x81, 0xf0, 0x2b, 0xc0, 0x69, 0x44, 0x6e, 0xd1, 0x99, 0x40,
	0x31, 0xf3, 0x29, 0x27, 0x29, 0xe3, 0x73, 0xe2, 0xfb, 0x08, 0x89, 0x89, 0x8e, 0x88, 0xd3, 0x98,
	0xbd, 0xa7, 0x68, 0x88, 0x01, 0x3a, 0x94, 0x73, 0xc6, 0xd1, 0xaf, 0xb6, 0xb3, 0x84, 0x1e, 0xdd,
	0x7c, 0xcb, 0x57, 0x65, 0x95, 0x63, 0x07, 0x8c, 0x66, 0x35, 0x72, 0x7f, 0xd6, 0xb8, 0xa7, 0xf6,
	0xc6, 0x95, 0x80, 0xcf, 0xa1, 0x5b, 0xed, 0x16, 0x5f, 0xf3, 0xbd, 0x5c, 0x57, 0x9f, 0x37, 0x15,
	0x7e, 0x02, 0xe6, 0xf6, 0xbe, 0xd8, 0x64, 0xf5, 0xee, 0x21, 0xb7, 0xdb, 0x52, 0x3a, 0x02, 0xe7,
	0xa7, 0x06, 0xba, 0xf7, 0x39, 0xab, 0x85, 0xad, 0x49, 0x0a, 0x97, 0xb2, 0x89, 0xc9, 0x8f, 0x00,
	0xdb, 0x60, 0x6c, 0x77, 0x8b, 0x2f, 0xf9, 0xa7, 0x5a, 0xa6, 0x9b, 0x5c, 0x95, 0x42, 0x51, 0xa3,
	0xb5, 0x0f, 0x8a, 0x1a, 0xe8, 0x15, 0x98, 0x7f, 0xee, 0xd2, 0xd6, 0xe5, 0xd8, 0x17, 0xff, 0x9c,
	0x50, 0xaa, 0x1c, 0xfc, 0x68, 0xc6, 0x4f, 0x41, 0xbf, 0x5b, 0x65, 0x85, 0xdd, 0x91, 0xb7, 0x0a,
	0xae, 0x18, 0xd0, 0x0d, 0x56, 0x59, 0xc1, 0x25, 0x77, 0x9e, 0x83, 0x2e, 0x2a, 0x6c, 0x81, 0x11,
	0xd3, 0xe9, 0x94, 0xdc, 0x50, 0x74, 0x22, 0x96, 0x92, 0xde, 0xca, 0x8b, 0xd3, 0xc4, 0xc5, 0x71,
	0x4a, 0x7c, 0xd4, 0x7a, 0xad, 0x7f, 0x6c, 0x55, 0x8b, 0x45, 0x57, 0xf6, 0x7b, 0xf1, 0x3b, 0x00,
	0x00, 0xff, 0xff, 0x87, 0xc6, 0xb8, 0x5b, 0x63, 0x03, 0x00, 0x00,
}
