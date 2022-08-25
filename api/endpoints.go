package api

// This file is autogenerated. DO NOT EDIT

import (
	"fmt"
	"strings"
)

///////////////

type resultGetCoins struct {
	Ok     bool `json:"ok"`
	Result struct {
		CountAllCoin int64  `json:"countAllCoin"`
		TotalReserve string `json:"totalReserve"`
		Coins        []struct {
			Symbol      string `json:"symbol"`
			Title       string `json:"title"`
			Volume      string `json:"volume"`
			Reserve     string `json:"reserve"`
			Crr         int64  `json:"crr"`
			LimitVolume string `json:"limitVolume"`
			Creator     string `json:"creator"`
			TxHash      string `json:"txHash"`
			BlockId     string `json:"blockId"`
			Price       string `json:"price"`
			Delegated   string `json:"delegated"`
			Avatar      string `json:"avatar"`
		} `json:"coins"`
	} `json:"result"`
}

// /coins
// Get all coins
func (api *API) GetCoins() ([]CoinInfo, error) {

	var link = "/coins"

	// request
	res, err := api.client.R().Get(link)
	if err = processConnectionError(res, err); err != nil {
		return nil, err
	}
	// json decode
	respValue, respErr := resultGetCoins{}, Error{}
	err = universalJSONDecode(res.Body(), &respValue, &respErr, func() (bool, bool) {
		return respValue.Ok, respErr.StatusCode != 0
	})
	if err != nil {
		return nil, joinErrors(err, respErr)
	}
	// process result
	return converterGetCoins(respValue)
}

///////////////

type resultGetEvmAccounts struct {
	Ok     bool `json:"ok"`
	Result struct {
		Count       uint64 `json:"count"`
		EvmAccounts []struct {
			CreatedAt                  string `json:"createdAt"`
			UpdatedAt                  string `json:"updatedAt"`
			Address                    string `json:"address"`
			CreationEvmBlockHeight     uint64 `json:"creationEvmBlockHeight"`
			CreationEvmTransactionHash string `json:"creationEvmTransactionHash"`
		} `json:"evmAccounts"`
	} `json:"result"`
}

// /evm-accounts
// Get evm accounts
func (api *API) GetEvmAccounts() ([]EvmAccount, error) {

	var link = "/evm-accounts"

	// request
	res, err := api.client.R().Get(link)
	if err = processConnectionError(res, err); err != nil {
		return nil, err
	}
	// json decode
	respValue, respErr := resultGetEvmAccounts{}, Error{}
	err = universalJSONDecode(res.Body(), &respValue, &respErr, func() (bool, bool) {
		return respValue.Ok, respErr.StatusCode != 0
	})
	if err != nil {
		return nil, joinErrors(err, respErr)
	}
	// process result
	return converterGetEvmAccounts(respValue)
}

///////////////

type resultGetEvmContracts struct {
	Ok     bool `json:"ok"`
	Result struct {
		Count        uint64 `json:"count"`
		EvmContracts []struct {
			CreatedAt                    string      `json:"createdAt"`
			UpdatedAt                    string      `json:"updatedAt"`
			Address                      string      `json:"address"`
			Status                       string      `json:"status"`
			Abi                          interface{} `json:"abi"`
			ByteCode                     string      `json:"byteCode"`
			DeploymentEvmAccountAddress  string      `json:"deploymentEvmAccountAddress"`
			DeploymentEvmBlockHeight     uint64      `json:"deploymentEvmBlockHeight"`
			DeploymentEvmTransactionHash string      `json:"deploymentEvmTransactionHash"`
			DeploymentEvmReceiptId       uint64      `json:"deploymentEvmReceiptId"`
		} `json:"evmContracts"`
	} `json:"result"`
}

// /evm-contracts
// Get evm contracts
func (api *API) GetEvmContracts() ([]EvmContract, error) {

	var link = "/evm-contracts"

	// request
	res, err := api.client.R().Get(link)
	if err = processConnectionError(res, err); err != nil {
		return nil, err
	}
	// json decode
	respValue, respErr := resultGetEvmContracts{}, Error{}
	err = universalJSONDecode(res.Body(), &respValue, &respErr, func() (bool, bool) {
		return respValue.Ok, respErr.StatusCode != 0
	})
	if err != nil {
		return nil, joinErrors(err, respErr)
	}
	// process result
	return converterGetEvmContracts(respValue)
}

///////////////

type resultGetEvmTransactions struct {
	Ok     bool `json:"ok"`
	Result struct {
		Count           uint64 `json:"count"`
		EvmTransactions []struct {
			CreatedAt            string      `json:"createdAt"`
			UpdatedAt            string      `json:"updatedAt"`
			Hash                 string      `json:"hash"`
			V                    string      `json:"v"`
			R                    string      `json:"r"`
			S                    string      `json:"s"`
			Gas                  uint64      `json:"gas"`
			Type                 uint64      `json:"type"`
			Input                string      `json:"input"`
			Nonce                uint64      `json:"nonce"`
			Value                string      `json:"value"`
			ChainId              uint64      `json:"chainId"`
			GasPrice             uint64      `json:"gasPrice"`
			AccessList           interface{} `json:"accessList"`
			MaxFeePerGas         uint64      `json:"maxFeePerGas"`
			MaxPriorityFeePerGas uint64      `json:"maxPriorityFeePerGas"`
			ExtraData            interface{} `json:"extraData"`
			From                 string      `json:"from"`
			To                   string      `json:"to"`
			EvmBlockHeight       uint64      `json:"evmBlockHeight"`
		} `json:"evmTransactions"`
	} `json:"result"`
}

// /evm-transactions
// Get evm transactions
func (api *API) GetEvmTransactions() ([]EvmTransaction, error) {

	var link = "/evm-transactions"

	// request
	res, err := api.client.R().Get(link)
	if err = processConnectionError(res, err); err != nil {
		return nil, err
	}
	// json decode
	respValue, respErr := resultGetEvmTransactions{}, Error{}
	err = universalJSONDecode(res.Body(), &respValue, &respErr, func() (bool, bool) {
		return respValue.Ok, respErr.StatusCode != 0
	})
	if err != nil {
		return nil, joinErrors(err, respErr)
	}
	// process result
	return converterGetEvmTransactions(respValue)
}

///////////////

type resultGetAllNFT struct {
	Ok     bool `json:"ok"`
	Result struct {
		Count int64 `json:"count"`
		Nfts  []struct {
			NftCollection string `json:"nftCollection"`
			NftId         string `json:"nftId"`
			Quantity      string `json:"quantity"`
			Reserve       string `json:"reserve"`
			Sender        string `json:"sender"`
			Recipient     string `json:"recipient"`
			TxHash        string `json:"txHash"`
		} `json:"nfts"`
	} `json:"result"`
}

// /nfts
// Get all nfts
func (api *API) GetAllNFT() ([]NFT, error) {

	var link = "/nfts"

	// request
	res, err := api.client.R().Get(link)
	if err = processConnectionError(res, err); err != nil {
		return nil, err
	}
	// json decode
	respValue, respErr := resultGetAllNFT{}, Error{}
	err = universalJSONDecode(res.Body(), &respValue, &respErr, func() (bool, bool) {
		return respValue.Ok, respErr.StatusCode != 0
	})
	if err != nil {
		return nil, joinErrors(err, respErr)
	}
	// process result
	return converterGetAllNFT(respValue)
}

///////////////

type resultGetTxs struct {
	Ok     bool `json:"ok"`
	Result struct {
		Count int64 `json:"count"`
		Txs   []struct {
			CreatedAt string `json:"createdAt"`
			UpdatedAt string `json:"updatedAt"`
			Hash      string `json:"hash"`
			Timestamp string `json:"timestamp"`
			Status    uint64 `json:"status"`
			Type      string `json:"type"`
			Fee       struct {
				Coin   string `json:"coin"`
				Amount string `json:"amount"`
			} `json:"fee"`
			Data    interface{} `json:"data"`
			Nonce   int64       `json:"nonce"`
			BlockId int64       `json:"blockId"`
			Message string      `json:"message"`
			From    string      `json:"from"`
			To      string      `json:"to"`
		} `json:"txs"`
	} `json:"result"`
}

// /txs
// Get all transactions
func (api *API) GetTxs() ([]TxInfo, error) {

	var link = "/txs"

	// request
	res, err := api.client.R().Get(link)
	if err = processConnectionError(res, err); err != nil {
		return nil, err
	}
	// json decode
	respValue, respErr := resultGetTxs{}, Error{}
	err = universalJSONDecode(res.Body(), &respValue, &respErr, func() (bool, bool) {
		return respValue.Ok, respErr.StatusCode != 0
	})
	if err != nil {
		return nil, joinErrors(err, respErr)
	}
	// process result
	return converterGetTxs(respValue)
}

///////////////

type resultGetAddress struct {
	Ok     bool `json:"ok"`
	Result struct {
		CreatedAt string `json:"createdAt"`
		UpdatedAt string `json:"updatedAt"`
		Address   string `json:"address"`
		Type      string `json:"type"`
		Balance   struct {
			Del string `json:"del"`
		} `json:"balance"`
		BalanceNft []struct {
			NftId      string `json:"nftId"`
			Amount     string `json:"amount"`
			Collection string `json:"collection"`
		} `json:"balanceNft"`
		Nonce         uint64 `json:"nonce"`
		Txes          uint64 `json:"txes"`
		UnbondBalance struct {
			Del string `json:"del"`
		} `json:"unbondBalance"`
		StakeBalance struct {
			Del string `json:"del"`
		} `json:"stakeBalance"`
	} `json:"result"`
}

// /address/{id}
// Get address by id
func (api *API) GetAddress(id string) (*AddressInfo, error) {

	var r = strings.NewReplacer(
		"{id}", fmt.Sprintf("%s", id),
	)
	var link = r.Replace("/address/{id}")

	// request
	res, err := api.client.R().Get(link)
	if err = processConnectionError(res, err); err != nil {
		return nil, err
	}
	// json decode
	respValue, respErr := resultGetAddress{}, Error{}
	err = universalJSONDecode(res.Body(), &respValue, &respErr, func() (bool, bool) {
		return respValue.Ok, respErr.StatusCode != 0
	})
	if err != nil {
		return nil, joinErrors(err, respErr)
	}
	// process result
	return converterGetAddress(respValue)
}

///////////////

type resultGetBlock struct {
	Ok     bool `json:"ok"`
	Result struct {
		CreatedAt       string `json:"createdAt"`
		UpdatedAt       string `json:"updatedAt"`
		Height          int64  `json:"height"`
		Date            string `json:"date"`
		Hash            string `json:"hash"`
		Size            int64  `json:"size"`
		Reward          int64  `json:"reward"`
		BlockTime       int64  `json:"blockTime"`
		TxsCount        int64  `json:"txsCount"`
		ValidatorsCount int64  `json:"validatorsCount"`
		ProposerId      string `json:"proposerId"`
		EvmBlock        struct {
			CreatedAt         string      `json:"createdAt"`
			UpdatedAt         string      `json:"updatedAt"`
			Height            uint64      `json:"height"`
			Hash              string      `json:"hash"`
			Date              string      `json:"date"`
			Miner             string      `json:"miner"`
			BaseFeePerGas     uint64      `json:"baseFeePerGas"`
			GasUsed           uint64      `json:"gasUsed"`
			GasLimit          uint64      `json:"gasLimit"`
			Data              interface{} `json:"data"`
			TransactionsCount uint64      `json:"transactionsCount"`
			ReceiptsCount     uint64      `json:"receiptsCount"`
		} `json:"evmBlock"`
	} `json:"result"`
}

// /block/{height}
// Get block by id
func (api *API) GetBlockByHeight(height uint64) (*BlockInfo, error) {

	var r = strings.NewReplacer(
		"{height}", fmt.Sprintf("%d", height),
	)
	var link = r.Replace("/block/{height}")

	// request
	res, err := api.client.R().Get(link)
	if err = processConnectionError(res, err); err != nil {
		return nil, err
	}
	// json decode
	respValue, respErr := resultGetBlock{}, Error{}
	err = universalJSONDecode(res.Body(), &respValue, &respErr, func() (bool, bool) {
		return respValue.Ok, respErr.StatusCode != 0
	})
	if err != nil {
		return nil, joinErrors(err, respErr)
	}
	// process result
	return converterGetBlockByHeight(respValue)
}

///////////////

type resultGetCoin struct {
	Ok     bool        `json:"ok"`
	Result interface{} `json:"result"`
}

// /coin/{coin}
// Get coin
func (api *API) GetCoin(coin string) (*CoinInfo, error) {

	var r = strings.NewReplacer(
		"{coin}", fmt.Sprintf("%s", coin),
	)
	var link = r.Replace("/coin/{coin}")

	// request
	res, err := api.client.R().Get(link)
	if err = processConnectionError(res, err); err != nil {
		return nil, err
	}
	// json decode
	respValue, respErr := resultGetCoin{}, Error{}
	err = universalJSONDecode(res.Body(), &respValue, &respErr, func() (bool, bool) {
		return respValue.Ok, respErr.StatusCode != 0
	})
	if err != nil {
		return nil, joinErrors(err, respErr)
	}
	// process result
	return converterGetCoin(respValue)
}

///////////////

type resultGetNFTCollection struct {
	Ok     bool        `json:"ok"`
	Result interface{} `json:"result"`
}

// /nft/{collection}
// Get nft collection
func (api *API) GetNFTCollection(collection string) (interface{}, error) {

	var r = strings.NewReplacer(
		"{collection}", fmt.Sprintf("%s", collection),
	)
	var link = r.Replace("/nft/{collection}")

	// request
	res, err := api.client.R().Get(link)
	if err = processConnectionError(res, err); err != nil {
		return nil, err
	}
	// json decode
	respValue, respErr := resultGetNFTCollection{}, Error{}
	err = universalJSONDecode(res.Body(), &respValue, &respErr, func() (bool, bool) {
		return respValue.Ok, respErr.StatusCode != 0
	})
	if err != nil {
		return nil, joinErrors(err, respErr)
	}
	// process result
	return converterGetNFTCollection(respValue)
}

///////////////

type resultGetTx struct {
	Ok     bool `json:"ok"`
	Result struct {
		CreatedAt string `json:"createdAt"`
		UpdatedAt string `json:"updatedAt"`
		Hash      string `json:"hash"`
		Timestamp string `json:"timestamp"`
		Status    uint64 `json:"status"`
		Type      string `json:"type"`
		Fee       struct {
			Coin   string `json:"coin"`
			Amount string `json:"amount"`
		} `json:"fee"`
		Data    interface{} `json:"data"`
		Nonce   int64       `json:"nonce"`
		BlockId int64       `json:"blockId"`
		Message string      `json:"message"`
		From    string      `json:"from"`
		To      string      `json:"to"`
	} `json:"result"`
}

// /tx/{hash}
// Get transaction by hash data
func (api *API) GetTxByHash(hash string) (*TxInfo, error) {

	var r = strings.NewReplacer(
		"{hash}", fmt.Sprintf("%s", hash),
	)
	var link = r.Replace("/tx/{hash}")

	// request
	res, err := api.client.R().Get(link)
	if err = processConnectionError(res, err); err != nil {
		return nil, err
	}
	// json decode
	respValue, respErr := resultGetTx{}, Error{}
	err = universalJSONDecode(res.Body(), &respValue, &respErr, func() (bool, bool) {
		return respValue.Ok, respErr.StatusCode != 0
	})
	if err != nil {
		return nil, joinErrors(err, respErr)
	}
	// process result
	return converterGetTxByHash(respValue)
}

///////////////

type resultGetValidator struct {
	Ok     bool `json:"ok"`
	Result struct {
		Address string `json:"address"`
		BlockId int64  `json:"blockId"`
	} `json:"result"`
}

// /validator/{address}
// Get validator by address (public key)
func (api *API) GetValidator(address string) (*Validator, error) {

	var r = strings.NewReplacer(
		"{address}", fmt.Sprintf("%s", address),
	)
	var link = r.Replace("/validator/{address}")

	// request
	res, err := api.client.R().Get(link)
	if err = processConnectionError(res, err); err != nil {
		return nil, err
	}
	// json decode
	respValue, respErr := resultGetValidator{}, Error{}
	err = universalJSONDecode(res.Body(), &respValue, &respErr, func() (bool, bool) {
		return respValue.Ok, respErr.StatusCode != 0
	})
	if err != nil {
		return nil, joinErrors(err, respErr)
	}
	// process result
	return converterGetValidator(respValue)
}

///////////////

type resultGetValidatorsByKind struct {
	Ok     bool `json:"ok"`
	Result struct {
		Count      int64 `json:"count"`
		Online     int64 `json:"online"`
		Validators []struct {
			Address string `json:"address"`
			BlockId int64  `json:"blockId"`
		} `json:"validators"`
		FreeSlots int64 `json:"freeSlots"`
	} `json:"result"`
}

// /validators/{kind}
// Default route
func (api *API) GetValidatorsByKind(kind string) ([]Validator, error) {

	var r = strings.NewReplacer(
		"{kind}", fmt.Sprintf("%s", kind),
	)
	var link = r.Replace("/validators/{kind}")

	// request
	res, err := api.client.R().Get(link)
	if err = processConnectionError(res, err); err != nil {
		return nil, err
	}
	// json decode
	respValue, respErr := resultGetValidatorsByKind{}, Error{}
	err = universalJSONDecode(res.Body(), &respValue, &respErr, func() (bool, bool) {
		return respValue.Ok, respErr.StatusCode != 0
	})
	if err != nil {
		return nil, joinErrors(err, respErr)
	}
	// process result
	return converterGetValidatorsByKind(respValue)
}

///////////////

type resultGetAddressTxs struct {
	Ok     bool `json:"ok"`
	Result struct {
		Count int64 `json:"count"`
		Txs   []struct {
			CreatedAt string `json:"createdAt"`
			UpdatedAt string `json:"updatedAt"`
			Hash      string `json:"hash"`
			Timestamp string `json:"timestamp"`
			Status    uint64 `json:"status"`
			Type      string `json:"type"`
			Fee       struct {
				Coin   string `json:"coin"`
				Amount string `json:"amount"`
			} `json:"fee"`
			Data    interface{} `json:"data"`
			Nonce   int64       `json:"nonce"`
			BlockId int64       `json:"blockId"`
			Message string      `json:"message"`
			From    string      `json:"from"`
			To      string      `json:"to"`
		} `json:"txs"`
	} `json:"result"`
}

// /address/{id}/txs
// Get address's transactions
func (api *API) GetAddressTxs(id string) ([]TxInfo, error) {

	var r = strings.NewReplacer(
		"{id}", fmt.Sprintf("%s", id),
	)
	var link = r.Replace("/address/{id}/txs")

	// request
	res, err := api.client.R().Get(link)
	if err = processConnectionError(res, err); err != nil {
		return nil, err
	}
	// json decode
	respValue, respErr := resultGetAddressTxs{}, Error{}
	err = universalJSONDecode(res.Body(), &respValue, &respErr, func() (bool, bool) {
		return respValue.Ok, respErr.StatusCode != 0
	})
	if err != nil {
		return nil, joinErrors(err, respErr)
	}
	// process result
	return converterGetAddressTxs(respValue)
}

///////////////

type resultGetValidatorStakes struct {
	Ok     bool `json:"ok"`
	Result struct {
		Count  int64 `json:"count"`
		Stakes []struct {
			CreatedAt   string `json:"createdAt"`
			UpdatedAt   string `json:"updatedAt"`
			CoinSymbol  string `json:"coinSymbol"`
			Amount      string `json:"amount"`
			AddressId   string `json:"addressId"`
			ValidatorId string `json:"validatorId"`
		} `json:"stakes"`
	} `json:"result"`
}

// /validator/{address}/stakes
// Get validator's stake
func (api *API) GetValidatorStakes(address string) ([]ValidatorStake, error) {

	var r = strings.NewReplacer(
		"{address}", fmt.Sprintf("%s", address),
	)
	var link = r.Replace("/validator/{address}/stakes")

	// request
	res, err := api.client.R().Get(link)
	if err = processConnectionError(res, err); err != nil {
		return nil, err
	}
	// json decode
	respValue, respErr := resultGetValidatorStakes{}, Error{}
	err = universalJSONDecode(res.Body(), &respValue, &respErr, func() (bool, bool) {
		return respValue.Ok, respErr.StatusCode != 0
	})
	if err != nil {
		return nil, joinErrors(err, respErr)
	}
	// process result
	return converterGetValidatorStakes(respValue)
}
