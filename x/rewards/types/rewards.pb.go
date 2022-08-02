// DONTCOVER
// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: archway/rewards/v1beta1/rewards.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
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

// Params defines the module parameters.
type Params struct {
	// inflation_rewards_ratio defines the percentage of minted inflation tokens that are used for dApp rewards [0.0, 1.0].
	// If set to 0.0, no inflation rewards are distributed.
	InflationRewardsRatio github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=inflation_rewards_ratio,json=inflationRewardsRatio,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"inflation_rewards_ratio"`
	// tx_fee_rebate_ratio defines the percentage of tx fees that are used for dApp rewards [0.0, 1.0].
	// If set to 0.0, no fee rewards are distributed.
	TxFeeRebateRatio github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=tx_fee_rebate_ratio,json=txFeeRebateRatio,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"tx_fee_rebate_ratio"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_50f478faffe74434, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

// ContractMetadata defines the contract rewards distribution options for a particular contract.
type ContractMetadata struct {
	// contract_address defines the contract address (bech32 encoded).
	ContractAddress string `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	// owner_address is the contract owner address that can modify contract reward options (bech32 encoded).
	// That could be the contract admin or the contract itself.
	// If owner_address is set to contract address, contract can modify the metadata on its own using WASM bindings.
	OwnerAddress string `protobuf:"bytes,2,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	// rewards_address is an address to distribute rewards to (bech32 encoded).
	// If not set (empty), rewards are not distributed for this contract.
	RewardsAddress string `protobuf:"bytes,3,opt,name=rewards_address,json=rewardsAddress,proto3" json:"rewards_address,omitempty"`
}

func (m *ContractMetadata) Reset()      { *m = ContractMetadata{} }
func (*ContractMetadata) ProtoMessage() {}
func (*ContractMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_50f478faffe74434, []int{1}
}
func (m *ContractMetadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContractMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContractMetadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContractMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractMetadata.Merge(m, src)
}
func (m *ContractMetadata) XXX_Size() int {
	return m.Size()
}
func (m *ContractMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_ContractMetadata proto.InternalMessageInfo

func (m *ContractMetadata) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *ContractMetadata) GetOwnerAddress() string {
	if m != nil {
		return m.OwnerAddress
	}
	return ""
}

func (m *ContractMetadata) GetRewardsAddress() string {
	if m != nil {
		return m.RewardsAddress
	}
	return ""
}

// BlockRewards defines block related rewards distribution data.
type BlockRewards struct {
	// height defines the block height.
	Height int64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	// inflation_rewards is the rewards to be distributed.
	InflationRewards types.Coin `protobuf:"bytes,2,opt,name=inflation_rewards,json=inflationRewards,proto3" json:"inflation_rewards"`
	// max_gas defines the maximum gas for the block that is used to distribute inflation rewards (consensus parameter).
	MaxGas uint64 `protobuf:"varint,3,opt,name=max_gas,json=maxGas,proto3" json:"max_gas,omitempty"`
}

func (m *BlockRewards) Reset()      { *m = BlockRewards{} }
func (*BlockRewards) ProtoMessage() {}
func (*BlockRewards) Descriptor() ([]byte, []int) {
	return fileDescriptor_50f478faffe74434, []int{2}
}
func (m *BlockRewards) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockRewards) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockRewards.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlockRewards) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockRewards.Merge(m, src)
}
func (m *BlockRewards) XXX_Size() int {
	return m.Size()
}
func (m *BlockRewards) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockRewards.DiscardUnknown(m)
}

var xxx_messageInfo_BlockRewards proto.InternalMessageInfo

func (m *BlockRewards) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *BlockRewards) GetInflationRewards() types.Coin {
	if m != nil {
		return m.InflationRewards
	}
	return types.Coin{}
}

func (m *BlockRewards) GetMaxGas() uint64 {
	if m != nil {
		return m.MaxGas
	}
	return 0
}

// TxRewards defines transaction related rewards distribution data.
type TxRewards struct {
	// tx_id is the tracking transaction ID (x/tracking is the data source for this value).
	TxId uint64 `protobuf:"varint,1,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
	// height defines the block height.
	Height int64 `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	// fee_rewards is the rewards to be distributed.
	FeeRewards []types.Coin `protobuf:"bytes,3,rep,name=fee_rewards,json=feeRewards,proto3" json:"fee_rewards"`
}

func (m *TxRewards) Reset()      { *m = TxRewards{} }
func (*TxRewards) ProtoMessage() {}
func (*TxRewards) Descriptor() ([]byte, []int) {
	return fileDescriptor_50f478faffe74434, []int{3}
}
func (m *TxRewards) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TxRewards) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TxRewards.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TxRewards) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxRewards.Merge(m, src)
}
func (m *TxRewards) XXX_Size() int {
	return m.Size()
}
func (m *TxRewards) XXX_DiscardUnknown() {
	xxx_messageInfo_TxRewards.DiscardUnknown(m)
}

var xxx_messageInfo_TxRewards proto.InternalMessageInfo

func (m *TxRewards) GetTxId() uint64 {
	if m != nil {
		return m.TxId
	}
	return 0
}

func (m *TxRewards) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *TxRewards) GetFeeRewards() []types.Coin {
	if m != nil {
		return m.FeeRewards
	}
	return nil
}

// BlockTracking is the tracking information for a block.
type BlockTracking struct {
	// inflation_rewards defines the inflation rewards for the block.
	InflationRewards BlockRewards `protobuf:"bytes,1,opt,name=inflation_rewards,json=inflationRewards,proto3" json:"inflation_rewards"`
	// tx_rewards defines the transaction rewards for the block.
	TxRewards []TxRewards `protobuf:"bytes,2,rep,name=tx_rewards,json=txRewards,proto3" json:"tx_rewards"`
}

func (m *BlockTracking) Reset()      { *m = BlockTracking{} }
func (*BlockTracking) ProtoMessage() {}
func (*BlockTracking) Descriptor() ([]byte, []int) {
	return fileDescriptor_50f478faffe74434, []int{4}
}
func (m *BlockTracking) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockTracking) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockTracking.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlockTracking) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockTracking.Merge(m, src)
}
func (m *BlockTracking) XXX_Size() int {
	return m.Size()
}
func (m *BlockTracking) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockTracking.DiscardUnknown(m)
}

var xxx_messageInfo_BlockTracking proto.InternalMessageInfo

func (m *BlockTracking) GetInflationRewards() BlockRewards {
	if m != nil {
		return m.InflationRewards
	}
	return BlockRewards{}
}

func (m *BlockTracking) GetTxRewards() []TxRewards {
	if m != nil {
		return m.TxRewards
	}
	return nil
}

func init() {
	proto.RegisterType((*Params)(nil), "archway.rewards.v1beta1.Params")
	proto.RegisterType((*ContractMetadata)(nil), "archway.rewards.v1beta1.ContractMetadata")
	proto.RegisterType((*BlockRewards)(nil), "archway.rewards.v1beta1.BlockRewards")
	proto.RegisterType((*TxRewards)(nil), "archway.rewards.v1beta1.TxRewards")
	proto.RegisterType((*BlockTracking)(nil), "archway.rewards.v1beta1.BlockTracking")
}

func init() {
	proto.RegisterFile("archway/rewards/v1beta1/rewards.proto", fileDescriptor_50f478faffe74434)
}

var fileDescriptor_50f478faffe74434 = []byte{
	// 533 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xcd, 0x6e, 0xd3, 0x40,
	0x14, 0x85, 0x33, 0x4d, 0x08, 0xca, 0xb4, 0xa5, 0xc1, 0x05, 0x02, 0x5d, 0x38, 0x55, 0x50, 0xf9,
	0x59, 0x74, 0xac, 0x96, 0x1d, 0x2b, 0x48, 0x11, 0x15, 0x52, 0x91, 0xd0, 0xa8, 0x0b, 0x84, 0x84,
	0xac, 0x1b, 0x7b, 0xe2, 0x58, 0xa9, 0x3d, 0xd5, 0x78, 0x20, 0xd3, 0x1d, 0x8f, 0x00, 0x62, 0xc3,
	0x92, 0xc7, 0xe0, 0x11, 0xba, 0xec, 0x06, 0x09, 0xb1, 0xa8, 0x50, 0xf2, 0x22, 0xc8, 0xf3, 0x63,
	0xa5, 0x40, 0x24, 0xd4, 0x55, 0x72, 0xaf, 0x8f, 0xcf, 0x7c, 0x3e, 0x77, 0x2e, 0xde, 0x02, 0x11,
	0x8d, 0x26, 0x70, 0x12, 0x08, 0x36, 0x01, 0x11, 0x17, 0xc1, 0xfb, 0x9d, 0x01, 0x93, 0xb0, 0xe3,
	0x6a, 0x72, 0x2c, 0xb8, 0xe4, 0x5e, 0xc7, 0xca, 0x88, 0x6b, 0x5b, 0xd9, 0xc6, 0x8d, 0x84, 0x27,
	0x5c, 0x6b, 0x82, 0xf2, 0x9f, 0x91, 0x6f, 0xf8, 0x11, 0x2f, 0x32, 0x5e, 0x04, 0x03, 0x28, 0x58,
	0xe5, 0x18, 0xf1, 0x34, 0x37, 0xcf, 0x7b, 0xdf, 0x11, 0x6e, 0xbe, 0x02, 0x01, 0x59, 0xe1, 0x0d,
	0x71, 0x27, 0xcd, 0x87, 0x47, 0x20, 0x53, 0x9e, 0x87, 0xd6, 0x3d, 0x14, 0x65, 0x79, 0x1b, 0x6d,
	0xa2, 0x07, 0xad, 0x3e, 0x39, 0x3d, 0xef, 0xd6, 0x7e, 0x9e, 0x77, 0xef, 0x25, 0xa9, 0x1c, 0xbd,
	0x1b, 0x90, 0x88, 0x67, 0x81, 0xb5, 0x37, 0x3f, 0xdb, 0x45, 0x3c, 0x0e, 0xe4, 0xc9, 0x31, 0x2b,
	0xc8, 0x33, 0x16, 0xd1, 0x9b, 0x95, 0x1d, 0x35, 0x6e, 0xb4, 0x2c, 0xbc, 0xb7, 0x78, 0x5d, 0xaa,
	0x70, 0xc8, 0x58, 0x28, 0xd8, 0x00, 0x24, 0xb3, 0x67, 0x2c, 0x5d, 0xea, 0x8c, 0xb6, 0x54, 0xcf,
	0x19, 0xa3, 0xda, 0x48, 0xdb, 0x3f, 0x6e, 0x7c, 0xf9, 0xda, 0xad, 0xf5, 0x3e, 0x21, 0xdc, 0xde,
	0xe3, 0xb9, 0x14, 0x10, 0xc9, 0x97, 0x4c, 0x42, 0x0c, 0x12, 0xbc, 0x87, 0xb8, 0x1d, 0xd9, 0x5e,
	0x08, 0x71, 0x2c, 0x58, 0x51, 0x98, 0x4f, 0xa3, 0x6b, 0xae, 0xff, 0xd4, 0xb4, 0xbd, 0xbb, 0x78,
	0x95, 0x4f, 0x72, 0x26, 0x2a, 0x9d, 0xc6, 0xa3, 0x2b, 0xba, 0xe9, 0x44, 0xf7, 0xf1, 0x9a, 0xcb,
	0xc9, 0xc9, 0xea, 0x5a, 0x76, 0xcd, 0xb6, 0xad, 0xd0, 0x32, 0x7d, 0x46, 0x78, 0xa5, 0x7f, 0xc4,
	0xa3, 0xb1, 0x8d, 0xc3, 0xbb, 0x85, 0x9b, 0x23, 0x96, 0x26, 0x23, 0xa9, 0x29, 0xea, 0xd4, 0x56,
	0xde, 0x01, 0xbe, 0xfe, 0xd7, 0x24, 0x34, 0xc0, 0xf2, 0xee, 0x1d, 0x62, 0x62, 0x20, 0xe5, 0x40,
	0xdd, 0xec, 0xc9, 0x1e, 0x4f, 0xf3, 0x7e, 0xa3, 0x8c, 0x8e, 0xb6, 0xff, 0x0c, 0xdd, 0xeb, 0xe0,
	0xab, 0x19, 0xa8, 0x30, 0x01, 0x43, 0xd7, 0xa0, 0xcd, 0x0c, 0xd4, 0x3e, 0x38, 0xaa, 0x0f, 0x08,
	0xb7, 0x0e, 0x95, 0x13, 0xaf, 0xe3, 0x2b, 0x52, 0x85, 0x69, 0xac, 0x89, 0x1a, 0xb4, 0x21, 0xd5,
	0x8b, 0x78, 0x8e, 0x73, 0xe9, 0x02, 0xe7, 0x13, 0xbc, 0x6c, 0xc6, 0x68, 0x08, 0xeb, 0x9b, 0xf5,
	0xff, 0x21, 0xc4, 0xc3, 0x72, 0x60, 0xfa, 0x15, 0x8b, 0xf0, 0x0d, 0xe1, 0x55, 0x1d, 0xcc, 0xa1,
	0x80, 0x68, 0x9c, 0xe6, 0x89, 0xf7, 0xfa, 0x5f, 0x09, 0x20, 0x9d, 0xc0, 0x16, 0x59, 0xb0, 0x01,
	0x64, 0x3e, 0xdb, 0x85, 0x69, 0xec, 0x63, 0x2c, 0xd5, 0x5c, 0xa8, 0x25, 0x72, 0x6f, 0xa1, 0x65,
	0x15, 0x8c, 0xf5, 0x6b, 0x49, 0x75, 0x01, 0xbd, 0x7f, 0x70, 0x3a, 0xf5, 0xd1, 0xd9, 0xd4, 0x47,
	0xbf, 0xa6, 0x3e, 0xfa, 0x38, 0xf3, 0x6b, 0x67, 0x33, 0xbf, 0xf6, 0x63, 0xe6, 0xd7, 0xde, 0xec,
	0xce, 0xdd, 0x60, 0x6b, 0xbf, 0x9d, 0x33, 0x39, 0xe1, 0x62, 0xec, 0xea, 0x40, 0x55, 0xcb, 0xae,
	0x6f, 0xf4, 0xa0, 0xa9, 0x97, 0xf2, 0xd1, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1f, 0x44, 0x4b,
	0xf2, 0x0c, 0x04, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.TxFeeRebateRatio.Size()
		i -= size
		if _, err := m.TxFeeRebateRatio.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintRewards(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.InflationRewardsRatio.Size()
		i -= size
		if _, err := m.InflationRewardsRatio.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintRewards(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ContractMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContractMetadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ContractMetadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RewardsAddress) > 0 {
		i -= len(m.RewardsAddress)
		copy(dAtA[i:], m.RewardsAddress)
		i = encodeVarintRewards(dAtA, i, uint64(len(m.RewardsAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.OwnerAddress) > 0 {
		i -= len(m.OwnerAddress)
		copy(dAtA[i:], m.OwnerAddress)
		i = encodeVarintRewards(dAtA, i, uint64(len(m.OwnerAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintRewards(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BlockRewards) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockRewards) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlockRewards) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MaxGas != 0 {
		i = encodeVarintRewards(dAtA, i, uint64(m.MaxGas))
		i--
		dAtA[i] = 0x18
	}
	{
		size, err := m.InflationRewards.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintRewards(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Height != 0 {
		i = encodeVarintRewards(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *TxRewards) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TxRewards) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TxRewards) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FeeRewards) > 0 {
		for iNdEx := len(m.FeeRewards) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FeeRewards[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRewards(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Height != 0 {
		i = encodeVarintRewards(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x10
	}
	if m.TxId != 0 {
		i = encodeVarintRewards(dAtA, i, uint64(m.TxId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *BlockTracking) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockTracking) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlockTracking) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TxRewards) > 0 {
		for iNdEx := len(m.TxRewards) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TxRewards[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRewards(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.InflationRewards.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintRewards(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintRewards(dAtA []byte, offset int, v uint64) int {
	offset -= sovRewards(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.InflationRewardsRatio.Size()
	n += 1 + l + sovRewards(uint64(l))
	l = m.TxFeeRebateRatio.Size()
	n += 1 + l + sovRewards(uint64(l))
	return n
}

func (m *ContractMetadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovRewards(uint64(l))
	}
	l = len(m.OwnerAddress)
	if l > 0 {
		n += 1 + l + sovRewards(uint64(l))
	}
	l = len(m.RewardsAddress)
	if l > 0 {
		n += 1 + l + sovRewards(uint64(l))
	}
	return n
}

func (m *BlockRewards) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovRewards(uint64(m.Height))
	}
	l = m.InflationRewards.Size()
	n += 1 + l + sovRewards(uint64(l))
	if m.MaxGas != 0 {
		n += 1 + sovRewards(uint64(m.MaxGas))
	}
	return n
}

func (m *TxRewards) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TxId != 0 {
		n += 1 + sovRewards(uint64(m.TxId))
	}
	if m.Height != 0 {
		n += 1 + sovRewards(uint64(m.Height))
	}
	if len(m.FeeRewards) > 0 {
		for _, e := range m.FeeRewards {
			l = e.Size()
			n += 1 + l + sovRewards(uint64(l))
		}
	}
	return n
}

func (m *BlockTracking) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.InflationRewards.Size()
	n += 1 + l + sovRewards(uint64(l))
	if len(m.TxRewards) > 0 {
		for _, e := range m.TxRewards {
			l = e.Size()
			n += 1 + l + sovRewards(uint64(l))
		}
	}
	return n
}

func sovRewards(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRewards(x uint64) (n int) {
	return sovRewards(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRewards
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InflationRewardsRatio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRewards
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
				return ErrInvalidLengthRewards
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRewards
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InflationRewardsRatio.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxFeeRebateRatio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRewards
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
				return ErrInvalidLengthRewards
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRewards
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TxFeeRebateRatio.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRewards(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRewards
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
func (m *ContractMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRewards
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
			return fmt.Errorf("proto: ContractMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContractMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRewards
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
				return ErrInvalidLengthRewards
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRewards
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRewards
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
				return ErrInvalidLengthRewards
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRewards
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OwnerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardsAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRewards
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
				return ErrInvalidLengthRewards
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRewards
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RewardsAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRewards(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRewards
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
func (m *BlockRewards) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRewards
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
			return fmt.Errorf("proto: BlockRewards: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockRewards: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRewards
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InflationRewards", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRewards
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
				return ErrInvalidLengthRewards
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRewards
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InflationRewards.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxGas", wireType)
			}
			m.MaxGas = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRewards
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxGas |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRewards(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRewards
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
func (m *TxRewards) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRewards
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
			return fmt.Errorf("proto: TxRewards: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TxRewards: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxId", wireType)
			}
			m.TxId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRewards
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TxId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRewards
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeRewards", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRewards
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
				return ErrInvalidLengthRewards
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRewards
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeeRewards = append(m.FeeRewards, types.Coin{})
			if err := m.FeeRewards[len(m.FeeRewards)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRewards(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRewards
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
func (m *BlockTracking) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRewards
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
			return fmt.Errorf("proto: BlockTracking: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockTracking: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InflationRewards", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRewards
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
				return ErrInvalidLengthRewards
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRewards
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InflationRewards.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxRewards", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRewards
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
				return ErrInvalidLengthRewards
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRewards
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxRewards = append(m.TxRewards, TxRewards{})
			if err := m.TxRewards[len(m.TxRewards)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRewards(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRewards
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
func skipRewards(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRewards
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
					return 0, ErrIntOverflowRewards
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
					return 0, ErrIntOverflowRewards
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
				return 0, ErrInvalidLengthRewards
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRewards
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRewards
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRewards        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRewards          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRewards = fmt.Errorf("proto: unexpected end of group")
)
