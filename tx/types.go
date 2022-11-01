package tx

import (
	coinTypes "bitbucket.org/decimalteam/go-smart-node/x/coin/types"
	multisigTypes "bitbucket.org/decimalteam/go-smart-node/x/multisig/types"
	nftTypes "bitbucket.org/decimalteam/go-smart-node/x/nft/types"
	swapTypes "bitbucket.org/decimalteam/go-smart-node/x/swap/types"
)

type (
	MsgCreateCoin    = coinTypes.MsgCreateCoin
	MsgUpdateCoin    = coinTypes.MsgUpdateCoin
	MsgMultiSendCoin = coinTypes.MsgMultiSendCoin
	MsgBuyCoin       = coinTypes.MsgBuyCoin
	MsgSellCoin      = coinTypes.MsgSellCoin
	MsgSellAllCoin   = coinTypes.MsgSellAllCoin
	MsgSendCoin      = coinTypes.MsgSendCoin
	MsgBurnCoin      = coinTypes.MsgBurnCoin
	MsgRedeemCheck   = coinTypes.MsgRedeemCheck

	MultiSendEntry = coinTypes.MultiSendEntry

	MsgMintToken     = nftTypes.MsgMintToken
	MsgBurnToken     = nftTypes.MsgBurnToken
	MsgUpdateReserve = nftTypes.MsgUpdateReserve
	MsgSendToken     = nftTypes.MsgSendToken
	MsgUpdateToken   = nftTypes.MsgUpdateToken

	MsgCreateWallet      = multisigTypes.MsgCreateWallet
	MsgCreateTransaction = multisigTypes.MsgCreateTransaction
	MsgSignTransaction   = multisigTypes.MsgSignTransaction

	MsgInitializeSwap = swapTypes.MsgInitializeSwap
	MsgRedeemSwap     = swapTypes.MsgRedeemSwap
)

var (
	NewMsgCreateCoin    = coinTypes.NewMsgCreateCoin
	NewMsgUpdateCoin    = coinTypes.NewMsgUpdateCoin
	NewMsgMultiSendCoin = coinTypes.NewMsgMultiSendCoin
	NewMsgBuyCoin       = coinTypes.NewMsgBuyCoin
	NewMsgSellCoin      = coinTypes.NewMsgSellCoin
	NewMsgSellAllCoin   = coinTypes.NewMsgSellAllCoin
	NewMsgSendCoin      = coinTypes.NewMsgSendCoin
	NewMsgBurnCoin      = coinTypes.NewMsgBurnCoin
	NewMsgRedeemCheck   = coinTypes.NewMsgRedeemCheck

	NewMsgMintToken     = nftTypes.NewMsgMintToken
	NewMsgBurnToken     = nftTypes.NewMsgBurnToken
	NewMsgUpdateReserve = nftTypes.NewMsgUpdateReserve
	NewMsgSendToken     = nftTypes.NewMsgSendToken
	NewMsgUpdateToken   = nftTypes.NewMsgUpdateToken

	NewMsgCreateWallet      = multisigTypes.NewMsgCreateWallet
	NewMsgCreateTransaction = multisigTypes.NewMsgCreateTransaction
	NewMsgSignTransaction   = multisigTypes.NewMsgSignTransaction

	NewMsgInitializeSwap = swapTypes.NewMsgInitializeSwap
	NewMsgRedeemSwap     = swapTypes.NewMsgRedeemSwap
)
