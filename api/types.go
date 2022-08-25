package api

import (
	"cosmossdk.io/math"
)

type CoinInfo struct {
	Symbol      string
	Title       string
	Volume      math.Int
	Reserve     math.Int
	LimitVolume math.Int
	CRR         int64
	Creator     string
}

type AddressInfo struct {
	Address string
	Type    string
	Nonce   uint64
}

type BlockInfo struct {
	Height          int64
	Reward          int64
	TxsCount        int64
	EvmTxsCount     int64
	ValidatorsCount int64
}

type TxInfo struct {
	Hash   string
	Status uint64
	Type   string
	Block  int64
	From   string
	To     string
}

// EVM

type EvmAccount struct {
	Address string
}

type EvmContract struct {
	Address                      string
	Status                       string
	DeploymentEvmAccountAddress  string
	DeploymentEvmTransactionHash string
}

type EvmTransaction struct {
	Hash           string
	Gas            uint64
	Type           uint64
	Input          string
	Nonce          uint64
	Value          string
	ChainId        uint64
	GasPrice       uint64
	MaxFeePerGas   uint64
	From           string
	To             string
	EvmBlockHeight uint64
}

type NFT struct {
	NftCollection string
	NftId         string
	Quantity      uint64
	Reserve       math.Int
	Sender        string
	Recipient     string
}

type Validator struct {
	Address string
}

type ValidatorStake struct {
	CoinSymbol  string
	Amount      math.Int
	AddressId   string
	ValidatorId string
}
