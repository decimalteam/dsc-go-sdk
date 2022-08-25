package api

import (
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	bigE15 = new(big.Int).Exp(big.NewInt(10), big.NewInt(15), nil)
	bigE18 = new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
	sdkE15 = sdk.NewIntFromBigInt(bigE15)
	sdkE18 = sdk.NewIntFromBigInt(bigE18)
)

func BipToPip(bip sdk.Int) sdk.Int {
	return bip.Mul(sdk.NewIntFromBigInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)))
}

// EtherToWei convert number 1 to 1 * 10^18
func EtherToWei(ether sdk.Int) sdk.Int {
	return ether.Mul(sdkE18)
}

// FinneyToWei convert number 1 to 1 * 10^15
func FinneyToWei(finney sdk.Int) sdk.Int {
	return finney.Mul(sdkE15)
}

// WeiToFinney convert 1 * 10^15 to 1
func WeiToFinney(wei sdk.Int) sdk.Int {
	return wei.Quo(sdkE15)
}

// WeiToEther convert 1 * 10^18 to 1
func WeiToEther(wei sdk.Int) sdk.Int {
	return wei.Quo(sdkE18)
}
