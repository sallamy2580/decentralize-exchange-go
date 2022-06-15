// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: block.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// BlockSyncMethod define block sync method.
type BlockSyncMethod int32

const (
	BlockSyncMethod_CONTRACTVM BlockSyncMethod = 0
	BlockSyncMethod_SQLDML     BlockSyncMethod = 1
)

var BlockSyncMethod_name = map[int32]string{
	0: "CONTRACTVM",
	1: "SQLDML",
}

var BlockSyncMethod_value = map[string]int32{
	"CONTRACTVM": 0,
	"SQLDML":     1,
}

func (x BlockSyncMethod) String() string {
	return proto.EnumName(BlockSyncMethod_name, int32(x))
}

func (BlockSyncMethod) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8e550b1f5926e92d, []int{0}
}

//BlockHeader is a structure of the block's header
type BlockHeader struct {
	BlockId      int64  `protobuf:"varint,1,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	Timestamp    int64  `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	EcosystemId  int64  `protobuf:"varint,3,opt,name=ecosystem_id,json=ecosystemId,proto3" json:"ecosystem_id,omitempty"`
	KeyId        int64  `protobuf:"varint,4,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	NodePosition int64  `protobuf:"varint,5,opt,name=node_position,json=nodePosition,proto3" json:"node_position,omitempty"`
	Sign         []byte `protobuf:"bytes,6,opt,name=sign,proto3" json:"sign,omitempty"`
	BlockHash    []byte `protobuf:"bytes,7,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`
	//differences with before and after in tx modification table
	RollbacksHash  []byte `protobuf:"bytes,8,opt,name=rollbacks_hash,json=rollbacksHash,proto3" json:"rollbacks_hash,omitempty"`
	Version        int32  `protobuf:"varint,9,opt,name=version,proto3" json:"version,omitempty"`
	ConsensusMode  int32  `protobuf:"varint,10,opt,name=consensus_mode,json=consensusMode,proto3" json:"consensus_mode,omitempty"`
	CandidateNodes []byte `protobuf:"bytes,11,opt,name=candidate_nodes,json=candidateNodes,proto3" json:"candidate_nodes,omitempty"`
}

func (m *BlockHeader) Reset()         { *m = BlockHeader{} }
func (m *BlockHeader) String() string { return proto.CompactTextString(m) }
func (*BlockHeader) ProtoMessage()    {}
func (*BlockHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e550b1f5926e92d, []int{0}
}
func (m *BlockHeader) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockHeader.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlockHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockHeader.Merge(m, src)
}
func (m *BlockHeader) XXX_Size() int {
	return m.Size()
}
func (m *BlockHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockHeader.DiscardUnknown(m)
}

var xxx_messageInfo_BlockHeader proto.InternalMessageInfo

func (m *BlockHeader) GetBlockId() int64 {
	if m != nil {
		return m.BlockId
	}
	return 0
}

func (m *BlockHeader) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *BlockHeader) GetEcosystemId() int64 {
	if m != nil {
		return m.EcosystemId
	}
	return 0
}

func (m *BlockHeader) GetKeyId() int64 {
	if m != nil {
		return m.KeyId
	}
	return 0
}

func (m *BlockHeader) GetNodePosition() int64 {
	if m != nil {
		return m.NodePosition
	}
	return 0
}

func (m *BlockHeader) GetSign() []byte {
	if m != nil {
		return m.Sign
	}
	return nil
}

func (m *BlockHeader) GetBlockHash() []byte {
	if m != nil {
		return m.BlockHash
	}
	return nil
}

func (m *BlockHeader) GetRollbacksHash() []byte {
	if m != nil {
		return m.RollbacksHash
	}
	return nil
}

func (m *BlockHeader) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *BlockHeader) GetConsensusMode() int32 {
	if m != nil {
		return m.ConsensusMode
	}
	return 0
}

func (m *BlockHeader) GetCandidateNodes() []byte {
	if m != nil {
		return m.CandidateNodes
	}
	return nil
}

// BlockData is a structure of the block's
type BlockData struct {
	Header     *BlockHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	PrevHeader *BlockHeader `protobuf:"bytes,2,opt,name=prev_header,json=prevHeader,proto3" json:"prev_header,omitempty"`
	MerkleRoot []byte       `protobuf:"bytes,3,opt,name=merkle_root,json=merkleRoot,proto3" json:"merkle_root,omitempty"`
	BinData    []byte       `protobuf:"bytes,4,opt,name=bin_data,json=binData,proto3" json:"bin_data,omitempty"`
	TxFullData [][]byte     `protobuf:"bytes,5,rep,name=tx_full_data,json=txFullData,proto3" json:"tx_full_data,omitempty"`
	AfterTxs   *AfterTxs    `protobuf:"bytes,6,opt,name=after_txs,json=afterTxs,proto3" json:"after_txs,omitempty"`
	SysUpdate  bool         `protobuf:"varint,7,opt,name=sys_update,json=sysUpdate,proto3" json:"sys_update,omitempty"`
}

func (m *BlockData) Reset()         { *m = BlockData{} }
func (m *BlockData) String() string { return proto.CompactTextString(m) }
func (*BlockData) ProtoMessage()    {}
func (*BlockData) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e550b1f5926e92d, []int{1}
}
func (m *BlockData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlockData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockData.Merge(m, src)
}
func (m *BlockData) XXX_Size() int {
	return m.Size()
}
func (m *BlockData) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockData.DiscardUnknown(m)
}

var xxx_messageInfo_BlockData proto.InternalMessageInfo

func (m *BlockData) GetHeader() *BlockHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *BlockData) GetPrevHeader() *BlockHeader {
	if m != nil {
		return m.PrevHeader
	}
	return nil
}

func (m *BlockData) GetMerkleRoot() []byte {
	if m != nil {
		return m.MerkleRoot
	}
	return nil
}

func (m *BlockData) GetBinData() []byte {
	if m != nil {
		return m.BinData
	}
	return nil
}

func (m *BlockData) GetTxFullData() [][]byte {
	if m != nil {
		return m.TxFullData
	}
	return nil
}

func (m *BlockData) GetAfterTxs() *AfterTxs {
	if m != nil {
		return m.AfterTxs
	}
	return nil
}

func (m *BlockData) GetSysUpdate() bool {
	if m != nil {
		return m.SysUpdate
	}
	return false
}

func init() {
	proto.RegisterEnum("types.BlockSyncMethod", BlockSyncMethod_name, BlockSyncMethod_value)
	proto.RegisterType((*BlockHeader)(nil), "types.BlockHeader")
	proto.RegisterType((*BlockData)(nil), "types.BlockData")
}

func init() { proto.RegisterFile("block.proto", fileDescriptor_8e550b1f5926e92d) }

var fileDescriptor_8e550b1f5926e92d = []byte{
	// 541 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xcd, 0x6e, 0xd3, 0x4e,
	0x14, 0xc5, 0xe3, 0xb4, 0xf9, 0xf0, 0xb5, 0xfb, 0xa1, 0x91, 0xfe, 0x92, 0xff, 0x08, 0x4c, 0x28,
	0x42, 0x94, 0x8a, 0x24, 0x52, 0xfb, 0x04, 0x49, 0x2a, 0xd4, 0x48, 0x4d, 0x01, 0x37, 0x20, 0xc4,
	0xc6, 0x1a, 0x7b, 0xa6, 0xb1, 0xe5, 0x8f, 0xb1, 0x3c, 0x93, 0x2a, 0x7e, 0x0b, 0x76, 0xbc, 0x12,
	0xcb, 0x2e, 0x59, 0xa2, 0x64, 0xc5, 0x5b, 0xa0, 0xb9, 0x31, 0x81, 0x0d, 0xbb, 0x99, 0xdf, 0x39,
	0x37, 0x73, 0x72, 0x8f, 0x0c, 0x56, 0x90, 0x8a, 0x30, 0x19, 0x14, 0xa5, 0x50, 0x82, 0xb4, 0x54,
	0x55, 0x70, 0xf9, 0x08, 0x8a, 0x94, 0x56, 0x5b, 0x74, 0xf2, 0xb3, 0x09, 0xd6, 0x58, 0x5b, 0xae,
	0x38, 0x65, 0xbc, 0x24, 0xff, 0x43, 0x17, 0x27, 0xfc, 0x98, 0x39, 0x46, 0xcf, 0x38, 0xdd, 0xf3,
	0x3a, 0x78, 0x9f, 0x32, 0xf2, 0x18, 0x4c, 0x15, 0x67, 0x5c, 0x2a, 0x9a, 0x15, 0x4e, 0x13, 0xb5,
	0x3f, 0x80, 0x3c, 0x03, 0x9b, 0x87, 0x42, 0x56, 0x52, 0xf1, 0x4c, 0x0f, 0xef, 0xa1, 0xc1, 0xda,
	0xb1, 0x29, 0x23, 0xff, 0x41, 0x3b, 0xe1, 0x95, 0x16, 0xf7, 0x51, 0x6c, 0x25, 0xbc, 0x9a, 0x32,
	0xf2, 0x1c, 0x0e, 0x72, 0xc1, 0xb8, 0x5f, 0x08, 0x19, 0xab, 0x58, 0xe4, 0x4e, 0x0b, 0x55, 0x5b,
	0xc3, 0x77, 0x35, 0x23, 0x04, 0xf6, 0x65, 0xbc, 0xc8, 0x9d, 0x76, 0xcf, 0x38, 0xb5, 0x3d, 0x3c,
	0x93, 0x27, 0x00, 0xdb, 0xac, 0x11, 0x95, 0x91, 0xd3, 0x41, 0xc5, 0x44, 0x72, 0x45, 0x65, 0x44,
	0x5e, 0xc0, 0x61, 0x29, 0xd2, 0x34, 0xa0, 0x61, 0x22, 0xb7, 0x96, 0x2e, 0x5a, 0x0e, 0x76, 0x14,
	0x6d, 0x0e, 0x74, 0xee, 0x79, 0x29, 0xf5, 0xc3, 0x66, 0xcf, 0x38, 0x6d, 0x79, 0xbf, 0xaf, 0xfa,
	0x07, 0x42, 0x91, 0x4b, 0x9e, 0xcb, 0xa5, 0xf4, 0x33, 0xc1, 0xb8, 0x03, 0x68, 0x38, 0xd8, 0xd1,
	0x99, 0x60, 0x9c, 0xbc, 0x84, 0xa3, 0x90, 0xe6, 0x2c, 0x66, 0x54, 0x71, 0x5f, 0x87, 0x96, 0x8e,
	0x85, 0x0f, 0x1d, 0xee, 0xf0, 0x8d, 0xa6, 0x27, 0x5f, 0x9b, 0x60, 0xe2, 0xae, 0x2f, 0xa9, 0xa2,
	0xe4, 0x0c, 0xda, 0x11, 0xee, 0x1c, 0xf7, 0x6c, 0x9d, 0x93, 0x01, 0xb6, 0x33, 0xf8, 0xab, 0x0d,
	0xaf, 0x76, 0x90, 0x0b, 0xb0, 0x8a, 0x92, 0xdf, 0xfb, 0xf5, 0x40, 0xf3, 0x9f, 0x03, 0xa0, 0x6d,
	0x75, 0x95, 0x4f, 0xc1, 0xca, 0x78, 0x99, 0xa4, 0xdc, 0x2f, 0x85, 0x50, 0x58, 0x88, 0xed, 0xc1,
	0x16, 0x79, 0x42, 0x28, 0xec, 0x3a, 0xce, 0x7d, 0x46, 0x15, 0xc5, 0x46, 0x6c, 0xaf, 0x13, 0xc4,
	0x39, 0x86, 0xeb, 0x81, 0xad, 0x56, 0xfe, 0xdd, 0x32, 0x4d, 0xb7, 0x72, 0xab, 0xb7, 0xa7, 0x87,
	0xd5, 0xea, 0xcd, 0x32, 0x4d, 0xd1, 0xf1, 0x1a, 0x4c, 0x7a, 0xa7, 0x78, 0xe9, 0xab, 0x95, 0xc4,
	0x56, 0xac, 0xf3, 0xa3, 0x3a, 0xd0, 0x48, 0xf3, 0xf9, 0x4a, 0x7a, 0x5d, 0x5a, 0x9f, 0x74, 0x55,
	0xb2, 0x92, 0xfe, 0xb2, 0xd0, 0xdb, 0xc0, 0xaa, 0xba, 0x9e, 0x29, 0x2b, 0xf9, 0x01, 0xc1, 0x59,
	0x1f, 0x8e, 0xf0, 0x5f, 0xdc, 0x56, 0x79, 0x38, 0xe3, 0x2a, 0x12, 0x8c, 0x1c, 0x02, 0x4c, 0xde,
	0xde, 0xcc, 0xbd, 0xd1, 0x64, 0xfe, 0x71, 0x76, 0xdc, 0x20, 0x00, 0xed, 0xdb, 0xf7, 0xd7, 0x97,
	0xb3, 0xeb, 0x63, 0x63, 0x3c, 0xf9, 0xb6, 0x76, 0x8d, 0x87, 0xb5, 0x6b, 0xfc, 0x58, 0xbb, 0xc6,
	0x97, 0x8d, 0xdb, 0x78, 0xd8, 0xb8, 0x8d, 0xef, 0x1b, 0xb7, 0xf1, 0xf9, 0xd5, 0x22, 0x56, 0xd1,
	0x32, 0x18, 0x84, 0x22, 0x1b, 0x4e, 0xc7, 0xa3, 0x4f, 0xfd, 0x58, 0x0c, 0x17, 0xa2, 0x1f, 0x07,
	0x74, 0x35, 0x2c, 0x68, 0x98, 0xd0, 0x05, 0x97, 0x43, 0x4c, 0x19, 0xb4, 0xf1, 0x03, 0xb8, 0xf8,
	0x15, 0x00, 0x00, 0xff, 0xff, 0x12, 0x76, 0x2c, 0xbf, 0x22, 0x03, 0x00, 0x00,
}

func (m *BlockHeader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockHeader) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlockHeader) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CandidateNodes) > 0 {
		i -= len(m.CandidateNodes)
		copy(dAtA[i:], m.CandidateNodes)
		i = encodeVarintBlock(dAtA, i, uint64(len(m.CandidateNodes)))
		i--
		dAtA[i] = 0x5a
	}
	if m.ConsensusMode != 0 {
		i = encodeVarintBlock(dAtA, i, uint64(m.ConsensusMode))
		i--
		dAtA[i] = 0x50
	}
	if m.Version != 0 {
		i = encodeVarintBlock(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x48
	}
	if len(m.RollbacksHash) > 0 {
		i -= len(m.RollbacksHash)
		copy(dAtA[i:], m.RollbacksHash)
		i = encodeVarintBlock(dAtA, i, uint64(len(m.RollbacksHash)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.BlockHash) > 0 {
		i -= len(m.BlockHash)
		copy(dAtA[i:], m.BlockHash)
		i = encodeVarintBlock(dAtA, i, uint64(len(m.BlockHash)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Sign) > 0 {
		i -= len(m.Sign)
		copy(dAtA[i:], m.Sign)
		i = encodeVarintBlock(dAtA, i, uint64(len(m.Sign)))
		i--
		dAtA[i] = 0x32
	}
	if m.NodePosition != 0 {
		i = encodeVarintBlock(dAtA, i, uint64(m.NodePosition))
		i--
		dAtA[i] = 0x28
	}
	if m.KeyId != 0 {
		i = encodeVarintBlock(dAtA, i, uint64(m.KeyId))
		i--
		dAtA[i] = 0x20
	}
	if m.EcosystemId != 0 {
		i = encodeVarintBlock(dAtA, i, uint64(m.EcosystemId))
		i--
		dAtA[i] = 0x18
	}
	if m.Timestamp != 0 {
		i = encodeVarintBlock(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x10
	}
	if m.BlockId != 0 {
		i = encodeVarintBlock(dAtA, i, uint64(m.BlockId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *BlockData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlockData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SysUpdate {
		i--
		if m.SysUpdate {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if m.AfterTxs != nil {
		{
			size, err := m.AfterTxs.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBlock(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if len(m.TxFullData) > 0 {
		for iNdEx := len(m.TxFullData) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.TxFullData[iNdEx])
			copy(dAtA[i:], m.TxFullData[iNdEx])
			i = encodeVarintBlock(dAtA, i, uint64(len(m.TxFullData[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.BinData) > 0 {
		i -= len(m.BinData)
		copy(dAtA[i:], m.BinData)
		i = encodeVarintBlock(dAtA, i, uint64(len(m.BinData)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.MerkleRoot) > 0 {
		i -= len(m.MerkleRoot)
		copy(dAtA[i:], m.MerkleRoot)
		i = encodeVarintBlock(dAtA, i, uint64(len(m.MerkleRoot)))
		i--
		dAtA[i] = 0x1a
	}
	if m.PrevHeader != nil {
		{
			size, err := m.PrevHeader.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBlock(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBlock(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBlock(dAtA []byte, offset int, v uint64) int {
	offset -= sovBlock(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BlockHeader) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BlockId != 0 {
		n += 1 + sovBlock(uint64(m.BlockId))
	}
	if m.Timestamp != 0 {
		n += 1 + sovBlock(uint64(m.Timestamp))
	}
	if m.EcosystemId != 0 {
		n += 1 + sovBlock(uint64(m.EcosystemId))
	}
	if m.KeyId != 0 {
		n += 1 + sovBlock(uint64(m.KeyId))
	}
	if m.NodePosition != 0 {
		n += 1 + sovBlock(uint64(m.NodePosition))
	}
	l = len(m.Sign)
	if l > 0 {
		n += 1 + l + sovBlock(uint64(l))
	}
	l = len(m.BlockHash)
	if l > 0 {
		n += 1 + l + sovBlock(uint64(l))
	}
	l = len(m.RollbacksHash)
	if l > 0 {
		n += 1 + l + sovBlock(uint64(l))
	}
	if m.Version != 0 {
		n += 1 + sovBlock(uint64(m.Version))
	}
	if m.ConsensusMode != 0 {
		n += 1 + sovBlock(uint64(m.ConsensusMode))
	}
	l = len(m.CandidateNodes)
	if l > 0 {
		n += 1 + l + sovBlock(uint64(l))
	}
	return n
}

func (m *BlockData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovBlock(uint64(l))
	}
	if m.PrevHeader != nil {
		l = m.PrevHeader.Size()
		n += 1 + l + sovBlock(uint64(l))
	}
	l = len(m.MerkleRoot)
	if l > 0 {
		n += 1 + l + sovBlock(uint64(l))
	}
	l = len(m.BinData)
	if l > 0 {
		n += 1 + l + sovBlock(uint64(l))
	}
	if len(m.TxFullData) > 0 {
		for _, b := range m.TxFullData {
			l = len(b)
			n += 1 + l + sovBlock(uint64(l))
		}
	}
	if m.AfterTxs != nil {
		l = m.AfterTxs.Size()
		n += 1 + l + sovBlock(uint64(l))
	}
	if m.SysUpdate {
		n += 2
	}
	return n
}

func sovBlock(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBlock(x uint64) (n int) {
	return sovBlock(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BlockHeader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlock
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
			return fmt.Errorf("proto: BlockHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockId", wireType)
			}
			m.BlockId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EcosystemId", wireType)
			}
			m.EcosystemId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EcosystemId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyId", wireType)
			}
			m.KeyId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.KeyId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodePosition", wireType)
			}
			m.NodePosition = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NodePosition |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sign", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
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
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sign = append(m.Sign[:0], dAtA[iNdEx:postIndex]...)
			if m.Sign == nil {
				m.Sign = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
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
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BlockHash = append(m.BlockHash[:0], dAtA[iNdEx:postIndex]...)
			if m.BlockHash == nil {
				m.BlockHash = []byte{}
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RollbacksHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
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
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RollbacksHash = append(m.RollbacksHash[:0], dAtA[iNdEx:postIndex]...)
			if m.RollbacksHash == nil {
				m.RollbacksHash = []byte{}
			}
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsensusMode", wireType)
			}
			m.ConsensusMode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ConsensusMode |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CandidateNodes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
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
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CandidateNodes = append(m.CandidateNodes[:0], dAtA[iNdEx:postIndex]...)
			if m.CandidateNodes == nil {
				m.CandidateNodes = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBlock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBlock
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
func (m *BlockData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlock
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
			return fmt.Errorf("proto: BlockData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
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
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &BlockHeader{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrevHeader", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
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
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PrevHeader == nil {
				m.PrevHeader = &BlockHeader{}
			}
			if err := m.PrevHeader.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MerkleRoot", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
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
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MerkleRoot = append(m.MerkleRoot[:0], dAtA[iNdEx:postIndex]...)
			if m.MerkleRoot == nil {
				m.MerkleRoot = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BinData", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
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
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BinData = append(m.BinData[:0], dAtA[iNdEx:postIndex]...)
			if m.BinData == nil {
				m.BinData = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxFullData", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
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
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxFullData = append(m.TxFullData, make([]byte, postIndex-iNdEx))
			copy(m.TxFullData[len(m.TxFullData)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AfterTxs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
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
				return ErrInvalidLengthBlock
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBlock
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AfterTxs == nil {
				m.AfterTxs = &AfterTxs{}
			}
			if err := m.AfterTxs.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SysUpdate", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.SysUpdate = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipBlock(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBlock
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
func skipBlock(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBlock
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
					return 0, ErrIntOverflowBlock
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowBlock
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
				return 0, ErrInvalidLengthBlock
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBlock
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBlock
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBlock        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBlock          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBlock = fmt.Errorf("proto: unexpected end of group")
)
