// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: network/node/internal/node/node.proto

package node

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"
	io "io"
	math "math"
	reflect "reflect"
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

type Node struct {
	NodeID         []byte `protobuf:"bytes,1,opt,name=NodeID,proto3" json:"NodeID,omitempty"`
	NodeShortID    uint32 `protobuf:"varint,2,opt,name=NodeShortID,proto3" json:"NodeShortID,omitempty"`
	NodeRole       uint32 `protobuf:"varint,3,opt,name=NodeRole,proto3" json:"NodeRole,omitempty"`
	NodePublicKey  []byte `protobuf:"bytes,4,opt,name=NodePublicKey,proto3" json:"NodePublicKey,omitempty"`
	NodeAddress    string `protobuf:"bytes,5,opt,name=NodeAddress,proto3" json:"NodeAddress,omitempty"`
	NodeVersion    string `protobuf:"bytes,6,opt,name=NodeVersion,proto3" json:"NodeVersion,omitempty"`
	NodeLeavingETA uint32 `protobuf:"varint,7,opt,name=NodeLeavingETA,proto3" json:"NodeLeavingETA,omitempty"`
	State          uint32 `protobuf:"varint,8,opt,name=state,proto3" json:"state,omitempty"`
}

func (m *Node) Reset()      { *m = Node{} }
func (*Node) ProtoMessage() {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_54a5c157c9a4f0ee, []int{0}
}
func (m *Node) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Node.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return m.Size()
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

type NodeList struct {
	List []*Node `protobuf:"bytes,1,rep,name=List,proto3" json:"List,omitempty"`
}

func (m *NodeList) Reset()      { *m = NodeList{} }
func (*NodeList) ProtoMessage() {}
func (*NodeList) Descriptor() ([]byte, []int) {
	return fileDescriptor_54a5c157c9a4f0ee, []int{1}
}
func (m *NodeList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NodeList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NodeList.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NodeList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeList.Merge(m, src)
}
func (m *NodeList) XXX_Size() int {
	return m.Size()
}
func (m *NodeList) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeList.DiscardUnknown(m)
}

var xxx_messageInfo_NodeList proto.InternalMessageInfo

type Snapshot struct {
	PulseNumber uint32               `protobuf:"varint,1,opt,name=pulseNumber,proto3" json:"pulseNumber,omitempty"`
	State       uint32               `protobuf:"varint,2,opt,name=state,proto3" json:"state,omitempty"`
	Nodes       map[uint32]*NodeList `protobuf:"bytes,3,rep,name=nodes,proto3" json:"nodes,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *Snapshot) Reset()      { *m = Snapshot{} }
func (*Snapshot) ProtoMessage() {}
func (*Snapshot) Descriptor() ([]byte, []int) {
	return fileDescriptor_54a5c157c9a4f0ee, []int{2}
}
func (m *Snapshot) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Snapshot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Snapshot.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Snapshot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Snapshot.Merge(m, src)
}
func (m *Snapshot) XXX_Size() int {
	return m.Size()
}
func (m *Snapshot) XXX_DiscardUnknown() {
	xxx_messageInfo_Snapshot.DiscardUnknown(m)
}

var xxx_messageInfo_Snapshot proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Node)(nil), "node.Node")
	proto.RegisterType((*NodeList)(nil), "node.NodeList")
	proto.RegisterType((*Snapshot)(nil), "node.Snapshot")
	proto.RegisterMapType((map[uint32]*NodeList)(nil), "node.Snapshot.NodesEntry")
}

func init() {
	proto.RegisterFile("network/node/internal/node/node.proto", fileDescriptor_54a5c157c9a4f0ee)
}

var fileDescriptor_54a5c157c9a4f0ee = []byte{
	// 435 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x52, 0xbf, 0x6f, 0x13, 0x31,
	0x14, 0x3e, 0xe7, 0x17, 0xe1, 0x85, 0x54, 0xc8, 0x42, 0xc8, 0x64, 0x78, 0x3a, 0x45, 0x05, 0x45,
	0x48, 0xe4, 0xa4, 0xb2, 0x20, 0xc4, 0x52, 0xd4, 0x4a, 0x54, 0x54, 0x15, 0xba, 0x22, 0xf6, 0xbb,
	0xc6, 0x24, 0xa7, 0x5e, 0xed, 0xc8, 0xf6, 0x15, 0x65, 0x63, 0xe0, 0x0f, 0xe0, 0xcf, 0xe0, 0x9f,
	0x60, 0xef, 0x98, 0xb1, 0x23, 0x77, 0x59, 0x18, 0x33, 0x32, 0x22, 0xdb, 0x09, 0x39, 0xba, 0x9c,
	0xbf, 0xef, 0x7b, 0xef, 0xde, 0xf7, 0xf9, 0xc9, 0xf0, 0x54, 0x70, 0xf3, 0x45, 0xaa, 0xcb, 0x48,
	0xc8, 0x09, 0x8f, 0x32, 0x61, 0xb8, 0x12, 0x49, 0xee, 0x99, 0xfd, 0x8c, 0xe7, 0x4a, 0x1a, 0x49,
	0x5b, 0x16, 0x0f, 0x5e, 0x4c, 0x33, 0x33, 0x2b, 0xd2, 0xf1, 0x85, 0xbc, 0x8a, 0xa6, 0x72, 0x2a,
	0x23, 0x57, 0x4c, 0x8b, 0xcf, 0x8e, 0x39, 0xe2, 0x90, 0xff, 0x69, 0xf8, 0xad, 0x01, 0xad, 0x33,
	0x39, 0xe1, 0xf4, 0x31, 0x74, 0xec, 0x79, 0x72, 0xc4, 0x48, 0x48, 0x46, 0x0f, 0xe2, 0x0d, 0xa3,
	0x21, 0xf4, 0x2c, 0x3a, 0x9f, 0x49, 0x65, 0x4e, 0x8e, 0x58, 0x23, 0x24, 0xa3, 0x7e, 0x5c, 0x97,
	0xe8, 0x00, 0xba, 0x96, 0xc6, 0x32, 0xe7, 0xac, 0xe9, 0xca, 0xff, 0x38, 0xdd, 0x87, 0xbe, 0xc5,
	0x1f, 0x8a, 0x34, 0xcf, 0x2e, 0xde, 0xf3, 0x05, 0x6b, 0xb9, 0xe1, 0xff, 0x8b, 0x5b, 0x8f, 0xc3,
	0xc9, 0x44, 0x71, 0xad, 0x59, 0x3b, 0x24, 0xa3, 0xfb, 0x71, 0x5d, 0xda, 0x76, 0x7c, 0xe2, 0x4a,
	0x67, 0x52, 0xb0, 0xce, 0xae, 0x63, 0x23, 0xd1, 0x67, 0xb0, 0x67, 0xe9, 0x29, 0x4f, 0xae, 0x33,
	0x31, 0x3d, 0xfe, 0x78, 0xc8, 0xee, 0xb9, 0x2c, 0x77, 0x54, 0xfa, 0x08, 0xda, 0xda, 0x24, 0x86,
	0xb3, 0xae, 0x2b, 0x7b, 0x32, 0x7c, 0xee, 0xef, 0x70, 0x9a, 0x69, 0x43, 0x11, 0x5a, 0xf6, 0x64,
	0x24, 0x6c, 0x8e, 0x7a, 0x07, 0x30, 0x76, 0x2b, 0x76, 0x37, 0x72, 0xfa, 0xf0, 0x27, 0x81, 0xee,
	0xb9, 0x48, 0xe6, 0x7a, 0x26, 0x8d, 0x0d, 0x36, 0x2f, 0x72, 0xcd, 0xcf, 0x8a, 0xab, 0x94, 0x2b,
	0xb7, 0xbb, 0x7e, 0x5c, 0x97, 0x76, 0x86, 0x8d, 0x9a, 0x21, 0x8d, 0xa0, 0x6d, 0xe7, 0x6a, 0xd6,
	0x74, 0x2e, 0x4f, 0xbc, 0xcb, 0x76, 0xac, 0xb3, 0xd3, 0xc7, 0xc2, 0xa8, 0x45, 0xec, 0xfb, 0x06,
	0xef, 0x00, 0x76, 0x22, 0x7d, 0x08, 0xcd, 0x4b, 0xbe, 0xd8, 0xd8, 0x59, 0x48, 0xf7, 0xa1, 0x7d,
	0x9d, 0xe4, 0x85, 0xb7, 0xe9, 0x1d, 0xec, 0xed, 0x62, 0xdb, 0xd0, 0xb1, 0x2f, 0xbe, 0x6e, 0xbc,
	0x22, 0x6f, 0xdf, 0xdc, 0x94, 0x18, 0x2c, 0x4b, 0x0c, 0x6e, 0x4b, 0x0c, 0xd6, 0x25, 0x92, 0x3f,
	0x25, 0x06, 0x5f, 0x2b, 0x24, 0x3f, 0x2a, 0x24, 0x37, 0x15, 0x92, 0x65, 0x85, 0xe4, 0x57, 0x85,
	0xe4, 0x77, 0x85, 0xc1, 0xba, 0x42, 0xf2, 0x7d, 0x85, 0xc1, 0x72, 0x85, 0xc1, 0xed, 0x0a, 0x83,
	0xb4, 0xe3, 0xde, 0xcd, 0xcb, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xbf, 0x4d, 0xc8, 0x84, 0x95,
	0x02, 0x00, 0x00,
}

func (this *Node) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Node)
	if !ok {
		that2, ok := that.(Node)
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
	if !bytes.Equal(this.NodeID, that1.NodeID) {
		return false
	}
	if this.NodeShortID != that1.NodeShortID {
		return false
	}
	if this.NodeRole != that1.NodeRole {
		return false
	}
	if !bytes.Equal(this.NodePublicKey, that1.NodePublicKey) {
		return false
	}
	if this.NodeAddress != that1.NodeAddress {
		return false
	}
	if this.NodeVersion != that1.NodeVersion {
		return false
	}
	if this.NodeLeavingETA != that1.NodeLeavingETA {
		return false
	}
	if this.State != that1.State {
		return false
	}
	return true
}
func (this *NodeList) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*NodeList)
	if !ok {
		that2, ok := that.(NodeList)
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
	if len(this.List) != len(that1.List) {
		return false
	}
	for i := range this.List {
		if !this.List[i].Equal(that1.List[i]) {
			return false
		}
	}
	return true
}
func (this *Snapshot) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Snapshot)
	if !ok {
		that2, ok := that.(Snapshot)
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
	if this.PulseNumber != that1.PulseNumber {
		return false
	}
	if this.State != that1.State {
		return false
	}
	if len(this.Nodes) != len(that1.Nodes) {
		return false
	}
	for i := range this.Nodes {
		if !this.Nodes[i].Equal(that1.Nodes[i]) {
			return false
		}
	}
	return true
}
func (this *Node) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 12)
	s = append(s, "&node.Node{")
	s = append(s, "NodeID: "+fmt.Sprintf("%#v", this.NodeID)+",\n")
	s = append(s, "NodeShortID: "+fmt.Sprintf("%#v", this.NodeShortID)+",\n")
	s = append(s, "NodeRole: "+fmt.Sprintf("%#v", this.NodeRole)+",\n")
	s = append(s, "NodePublicKey: "+fmt.Sprintf("%#v", this.NodePublicKey)+",\n")
	s = append(s, "NodeAddress: "+fmt.Sprintf("%#v", this.NodeAddress)+",\n")
	s = append(s, "NodeVersion: "+fmt.Sprintf("%#v", this.NodeVersion)+",\n")
	s = append(s, "NodeLeavingETA: "+fmt.Sprintf("%#v", this.NodeLeavingETA)+",\n")
	s = append(s, "State: "+fmt.Sprintf("%#v", this.State)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *NodeList) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&node.NodeList{")
	if this.List != nil {
		s = append(s, "List: "+fmt.Sprintf("%#v", this.List)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Snapshot) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&node.Snapshot{")
	s = append(s, "PulseNumber: "+fmt.Sprintf("%#v", this.PulseNumber)+",\n")
	s = append(s, "State: "+fmt.Sprintf("%#v", this.State)+",\n")
	keysForNodes := make([]uint32, 0, len(this.Nodes))
	for k, _ := range this.Nodes {
		keysForNodes = append(keysForNodes, k)
	}
	github_com_gogo_protobuf_sortkeys.Uint32s(keysForNodes)
	mapStringForNodes := "map[uint32]*NodeList{"
	for _, k := range keysForNodes {
		mapStringForNodes += fmt.Sprintf("%#v: %#v,", k, this.Nodes[k])
	}
	mapStringForNodes += "}"
	if this.Nodes != nil {
		s = append(s, "Nodes: "+mapStringForNodes+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringNode(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Node) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Node) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.NodeID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintNode(dAtA, i, uint64(len(m.NodeID)))
		i += copy(dAtA[i:], m.NodeID)
	}
	if m.NodeShortID != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintNode(dAtA, i, uint64(m.NodeShortID))
	}
	if m.NodeRole != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintNode(dAtA, i, uint64(m.NodeRole))
	}
	if len(m.NodePublicKey) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintNode(dAtA, i, uint64(len(m.NodePublicKey)))
		i += copy(dAtA[i:], m.NodePublicKey)
	}
	if len(m.NodeAddress) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintNode(dAtA, i, uint64(len(m.NodeAddress)))
		i += copy(dAtA[i:], m.NodeAddress)
	}
	if len(m.NodeVersion) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintNode(dAtA, i, uint64(len(m.NodeVersion)))
		i += copy(dAtA[i:], m.NodeVersion)
	}
	if m.NodeLeavingETA != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintNode(dAtA, i, uint64(m.NodeLeavingETA))
	}
	if m.State != 0 {
		dAtA[i] = 0x40
		i++
		i = encodeVarintNode(dAtA, i, uint64(m.State))
	}
	return i, nil
}

func (m *NodeList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NodeList) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.List) > 0 {
		for _, msg := range m.List {
			dAtA[i] = 0xa
			i++
			i = encodeVarintNode(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *Snapshot) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Snapshot) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.PulseNumber != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintNode(dAtA, i, uint64(m.PulseNumber))
	}
	if m.State != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintNode(dAtA, i, uint64(m.State))
	}
	if len(m.Nodes) > 0 {
		for k, _ := range m.Nodes {
			dAtA[i] = 0x1a
			i++
			v := m.Nodes[k]
			msgSize := 0
			if v != nil {
				msgSize = v.Size()
				msgSize += 1 + sovNode(uint64(msgSize))
			}
			mapSize := 1 + sovNode(uint64(k)) + msgSize
			i = encodeVarintNode(dAtA, i, uint64(mapSize))
			dAtA[i] = 0x8
			i++
			i = encodeVarintNode(dAtA, i, uint64(k))
			if v != nil {
				dAtA[i] = 0x12
				i++
				i = encodeVarintNode(dAtA, i, uint64(v.Size()))
				n1, err1 := v.MarshalTo(dAtA[i:])
				if err1 != nil {
					return 0, err1
				}
				i += n1
			}
		}
	}
	return i, nil
}

func encodeVarintNode(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Node) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.NodeID)
	if l > 0 {
		n += 1 + l + sovNode(uint64(l))
	}
	if m.NodeShortID != 0 {
		n += 1 + sovNode(uint64(m.NodeShortID))
	}
	if m.NodeRole != 0 {
		n += 1 + sovNode(uint64(m.NodeRole))
	}
	l = len(m.NodePublicKey)
	if l > 0 {
		n += 1 + l + sovNode(uint64(l))
	}
	l = len(m.NodeAddress)
	if l > 0 {
		n += 1 + l + sovNode(uint64(l))
	}
	l = len(m.NodeVersion)
	if l > 0 {
		n += 1 + l + sovNode(uint64(l))
	}
	if m.NodeLeavingETA != 0 {
		n += 1 + sovNode(uint64(m.NodeLeavingETA))
	}
	if m.State != 0 {
		n += 1 + sovNode(uint64(m.State))
	}
	return n
}

func (m *NodeList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.List) > 0 {
		for _, e := range m.List {
			l = e.Size()
			n += 1 + l + sovNode(uint64(l))
		}
	}
	return n
}

func (m *Snapshot) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PulseNumber != 0 {
		n += 1 + sovNode(uint64(m.PulseNumber))
	}
	if m.State != 0 {
		n += 1 + sovNode(uint64(m.State))
	}
	if len(m.Nodes) > 0 {
		for k, v := range m.Nodes {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovNode(uint64(l))
			}
			mapEntrySize := 1 + sovNode(uint64(k)) + l
			n += mapEntrySize + 1 + sovNode(uint64(mapEntrySize))
		}
	}
	return n
}

func sovNode(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozNode(x uint64) (n int) {
	return sovNode(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Node) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Node{`,
		`NodeID:` + fmt.Sprintf("%v", this.NodeID) + `,`,
		`NodeShortID:` + fmt.Sprintf("%v", this.NodeShortID) + `,`,
		`NodeRole:` + fmt.Sprintf("%v", this.NodeRole) + `,`,
		`NodePublicKey:` + fmt.Sprintf("%v", this.NodePublicKey) + `,`,
		`NodeAddress:` + fmt.Sprintf("%v", this.NodeAddress) + `,`,
		`NodeVersion:` + fmt.Sprintf("%v", this.NodeVersion) + `,`,
		`NodeLeavingETA:` + fmt.Sprintf("%v", this.NodeLeavingETA) + `,`,
		`State:` + fmt.Sprintf("%v", this.State) + `,`,
		`}`,
	}, "")
	return s
}
func (this *NodeList) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForList := "[]*Node{"
	for _, f := range this.List {
		repeatedStringForList += strings.Replace(f.String(), "Node", "Node", 1) + ","
	}
	repeatedStringForList += "}"
	s := strings.Join([]string{`&NodeList{`,
		`List:` + repeatedStringForList + `,`,
		`}`,
	}, "")
	return s
}
func (this *Snapshot) String() string {
	if this == nil {
		return "nil"
	}
	keysForNodes := make([]uint32, 0, len(this.Nodes))
	for k, _ := range this.Nodes {
		keysForNodes = append(keysForNodes, k)
	}
	github_com_gogo_protobuf_sortkeys.Uint32s(keysForNodes)
	mapStringForNodes := "map[uint32]*NodeList{"
	for _, k := range keysForNodes {
		mapStringForNodes += fmt.Sprintf("%v: %v,", k, this.Nodes[k])
	}
	mapStringForNodes += "}"
	s := strings.Join([]string{`&Snapshot{`,
		`PulseNumber:` + fmt.Sprintf("%v", this.PulseNumber) + `,`,
		`State:` + fmt.Sprintf("%v", this.State) + `,`,
		`Nodes:` + mapStringForNodes + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringNode(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Node) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNode
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
			return fmt.Errorf("proto: Node: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Node: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthNode
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthNode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeID = append(m.NodeID[:0], dAtA[iNdEx:postIndex]...)
			if m.NodeID == nil {
				m.NodeID = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeShortID", wireType)
			}
			m.NodeShortID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NodeShortID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeRole", wireType)
			}
			m.NodeRole = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NodeRole |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodePublicKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthNode
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthNode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodePublicKey = append(m.NodePublicKey[:0], dAtA[iNdEx:postIndex]...)
			if m.NodePublicKey == nil {
				m.NodePublicKey = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
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
				return ErrInvalidLengthNode
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
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
				return ErrInvalidLengthNode
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeLeavingETA", wireType)
			}
			m.NodeLeavingETA = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NodeLeavingETA |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipNode(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNode
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNode
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
func (m *NodeList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNode
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
			return fmt.Errorf("proto: NodeList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NodeList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field List", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
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
				return ErrInvalidLengthNode
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.List = append(m.List, &Node{})
			if err := m.List[len(m.List)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNode(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNode
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNode
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
func (m *Snapshot) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNode
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
			return fmt.Errorf("proto: Snapshot: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Snapshot: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PulseNumber", wireType)
			}
			m.PulseNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PulseNumber |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nodes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
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
				return ErrInvalidLengthNode
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Nodes == nil {
				m.Nodes = make(map[uint32]*NodeList)
			}
			var mapkey uint32
			var mapvalue *NodeList
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowNode
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
				if fieldNum == 1 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowNode
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapkey |= uint32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowNode
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthNode
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthNode
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &NodeList{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipNode(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthNode
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Nodes[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNode(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNode
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNode
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
func skipNode(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNode
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
					return 0, ErrIntOverflowNode
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
					return 0, ErrIntOverflowNode
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
				return 0, ErrInvalidLengthNode
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthNode
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowNode
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
				next, err := skipNode(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthNode
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
	ErrInvalidLengthNode = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNode   = fmt.Errorf("proto: integer overflow")
)
