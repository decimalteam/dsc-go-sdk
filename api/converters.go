package api

import (
	"fmt"
	"strconv"

	"cosmossdk.io/math"
)

func converterGetCoin(resp resultGetCoin) (*CoinInfo, error) {
	return &CoinInfo{}, nil
}

func converterGetCoins(resp resultGetCoins) ([]CoinInfo, error) {
	var ok bool
	res := make([]CoinInfo, len(resp.Result.Coins))
	for i, coin := range resp.Result.Coins {
		res[i].Symbol = coin.Symbol
		res[i].Title = coin.Title
		res[i].CRR = coin.Crr
		res[i].Volume, ok = math.NewIntFromString(coin.Volume)
		if !ok {
			return nil, fmt.Errorf("cannot convert volume '%s' to math.Int", coin.Volume)
		}
		res[i].Reserve, ok = math.NewIntFromString(coin.Reserve)
		if !ok {
			return nil, fmt.Errorf("cannot convert reserve '%s' to math.Int", coin.Reserve)
		}
		res[i].LimitVolume, ok = math.NewIntFromString(coin.LimitVolume)
		if !ok {
			return nil, fmt.Errorf("cannot convert limit volume '%s' to math.Int", coin.LimitVolume)
		}
		res[i].Creator = coin.Creator
	}
	return res, nil
}

func converterGetAddress(resp resultGetAddress) (*AddressInfo, error) {
	var res AddressInfo
	res.Address = resp.Result.Address
	res.Type = resp.Result.Type
	res.Nonce = resp.Result.Nonce
	return &res, nil
}

func converterGetBlockByHeight(resp resultGetBlock) (*BlockInfo, error) {
	var res BlockInfo
	res.Height = resp.Result.Height
	res.Reward = resp.Result.Reward
	res.TxsCount = resp.Result.TxsCount
	res.EvmTxsCount = int64(resp.Result.EvmBlock.TransactionsCount)
	res.ValidatorsCount = resp.Result.ValidatorsCount
	return &res, nil
}

func converterGetNFTCollection(resp resultGetNFTCollection) (interface{}, error) {
	return resp.Result, nil
}

func converterGetTxByHash(resp resultGetTx) (*TxInfo, error) {
	var res TxInfo
	res.Block = resp.Result.BlockId
	res.From = resp.Result.From
	res.To = resp.Result.To
	res.Hash = resp.Result.Hash
	res.Status = resp.Result.Status
	return &res, nil
}

func converterGetTxs(resp resultGetTxs) ([]TxInfo, error) {
	var res = make([]TxInfo, len(resp.Result.Txs))
	for i, tx := range resp.Result.Txs {
		res[i].Block = tx.BlockId
		res[i].From = tx.From
		res[i].To = tx.To
		res[i].Hash = tx.Hash
		res[i].Status = tx.Status
	}
	return res, nil
}

func converterGetAddressTxs(resp resultGetAddressTxs) ([]TxInfo, error) {
	var res = make([]TxInfo, len(resp.Result.Txs))
	for i, tx := range resp.Result.Txs {
		res[i].Block = tx.BlockId
		res[i].From = tx.From
		res[i].To = tx.To
		res[i].Hash = tx.Hash
		res[i].Status = tx.Status
	}
	return res, nil
}

func converterGetEvmContracts(resp resultGetEvmContracts) ([]EvmContract, error) {
	var res = make([]EvmContract, len(resp.Result.EvmContracts))
	for i, cntr := range resp.Result.EvmContracts {
		res[i].Address = cntr.Address
		res[i].Status = cntr.Status
		res[i].DeploymentEvmAccountAddress = cntr.DeploymentEvmAccountAddress
		res[i].DeploymentEvmTransactionHash = cntr.DeploymentEvmTransactionHash
	}
	return res, nil
}

func converterGetEvmTransactions(resp resultGetEvmTransactions) ([]EvmTransaction, error) {
	var res = make([]EvmTransaction, len(resp.Result.EvmTransactions))
	for i, tx := range resp.Result.EvmTransactions {
		res[i].Hash = tx.Hash
		res[i].Gas = tx.Gas
		res[i].Type = tx.Type
		res[i].Input = tx.Input
		res[i].Nonce = tx.Nonce
		res[i].Value = tx.Value
		res[i].ChainId = tx.ChainId
		res[i].GasPrice = tx.GasPrice
		res[i].MaxFeePerGas = tx.MaxFeePerGas
		res[i].From = tx.From
		res[i].To = tx.To
		res[i].EvmBlockHeight = tx.EvmBlockHeight
	}
	return res, nil
}

func converterGetEvmAccounts(resp resultGetEvmAccounts) ([]EvmAccount, error) {
	var res = make([]EvmAccount, len(resp.Result.EvmAccounts))
	for i, acc := range resp.Result.EvmAccounts {
		res[i].Address = acc.Address
	}
	return res, nil
}

func converterGetAllNFT(resp resultGetAllNFT) ([]NFT, error) {
	var err error
	var ok bool
	var res = make([]NFT, len(resp.Result.Nfts))
	for i, nft := range resp.Result.Nfts {
		res[i].NftCollection = nft.NftCollection
		res[i].NftId = nft.NftId
		res[i].Quantity, err = strconv.ParseUint(nft.Quantity, 10, 64)
		if err != nil {
			return nil, err
		}
		res[i].Reserve, ok = math.NewIntFromString(nft.Reserve)
		if !ok {
			return nil, fmt.Errorf("cannot convert reserve '%s' to math.Int", nft.Reserve)
		}
		res[i].Sender = nft.Sender
		res[i].Recipient = nft.Recipient
	}
	return res, nil
}

func converterGetValidatorsByKind(resp resultGetValidatorsByKind) ([]Validator, error) {
	var res = make([]Validator, len(resp.Result.Validators))
	for i, val := range resp.Result.Validators {
		res[i].Address = val.Address
	}
	return res, nil
}

func converterGetValidator(resp resultGetValidator) (*Validator, error) {
	return &Validator{Address: resp.Result.Address}, nil
}

func converterGetValidatorStakes(resp resultGetValidatorStakes) ([]ValidatorStake, error) {
	var ok bool
	var res = make([]ValidatorStake, len(resp.Result.Stakes))
	for i, stake := range resp.Result.Stakes {
		res[i].ValidatorId = stake.ValidatorId
		res[i].AddressId = stake.AddressId
		res[i].Amount, ok = math.NewIntFromString(stake.Amount)
		if !ok {
			return nil, fmt.Errorf("cannot convert amount '%s' to math.Int", stake.Amount)
		}
		res[i].CoinSymbol = stake.CoinSymbol
	}
	return res, nil
}
