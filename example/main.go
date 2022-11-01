package main

// This is example of using Decimal Smart Chain API

import (
	"encoding/json"
	"fmt"
	"time"

	dscApi "bitbucket.org/decimalteam/dsc-go-sdk/api"
	dscSwagger "bitbucket.org/decimalteam/dsc-go-sdk/swagger"
	dscTx "bitbucket.org/decimalteam/dsc-go-sdk/tx"
	dscWallet "bitbucket.org/decimalteam/dsc-go-sdk/wallet"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func main() {
	verifyEndpoints()

	api := dscApi.NewAPI(
		"https://devnet-explorer-api.decimalchain.com/api/",
		"https://devnet-val.decimalchain.com/rpc/",
		"https://devnet-val.decimalchain.com/rest/",
	)

	err := api.GetParameters()
	if err != nil {
		fmt.Printf("GetParameters() error: %v\n", err)
		return
	}

	/*
		we, err := dscApi.CreateTxSubscription("wss://devnet-dec2-explorer-api.decimalchain.com/api")
		if err != nil {
			fmt.Printf("CreateTxSubscription() error: %v\n", err)
			return
		}
		go we.ReadCycle()
	*/

	//printBlockchainInfo(api)

	//sampleSendCoins(api)
	//time.Sleep(time.Second * 10)

}

func verifyEndpoints() {
	const address = "dx1fatzsagt96pfglxlq245th252mv3neckvkmf68"

	var res []string
	var err error
	apiVerificator := dscSwagger.NewAPI("https://devnet-explorer-api.decimalchain.com/api/")
	api := dscApi.NewAPI(
		"https://devnet-explorer-api.decimalchain.com/api/",
		"https://devnet-val.decimalchain.com/rpc/",
		"https://devnet-val.decimalchain.com/rest/",
	)

	res, err = apiVerificator.VerificationGetAddress(address)
	fmt.Printf("VerificationGetAddress: err = %v, result = %s\n", err, formatAsJSON(res))

	// res, err = apiVerificator.VerificationGetAddressTxs(address, nil)
	// fmt.Printf("VerificationGetAddressTxs: err = %v, result = %s\n", err, formatAsJSON(res))

	// res, err = apiVerificator.VerificationGetAddressStakes(address, nil)
	// fmt.Printf("VerificationGetAddressStakes: err = %v, result = %s\n", err, formatAsJSON(res))

	// res, err = apiVerificator.VerificationGetAddressRewards(address, nil)
	// fmt.Printf("VerificationGetAddressRewards: err = %v, result = %s\n", err, formatAsJSON(res))

	// res, err = apiVerificator.VerificationGetAllNFT(nil)
	// fmt.Printf("VerificationGetAllNFT: err = %v, result = %s\n", err, formatAsJSON(res))

	// nfts, err := api.GetAllNFT(&dscApi.OptionalParams{Limit: 1})
	// if err == nil && len(nfts) > 0 {
	// 	res, err = apiVerificator.VerificationGetNFTCollection(nfts[0].NftCollection)
	// 	fmt.Printf("VerificationGetNFTCollection: err = %v, result = %s\n", err, formatAsJSON(res))

	// 	res, err := apiVerificator.VerificationGetNFTTransactions(nfts[0].NftCollection, nfts[0].NftId, &dscSwagger.OptionalParams{Limit: 1})
	// 	fmt.Printf("VerificationGetNFTTransactions: err = %v, result = %s\n", err, formatAsJSON(res))
	// }

	// TODO: test after unmarshaling fix

	res, err = apiVerificator.VerificationGetTxs(&dscSwagger.OptionalParams{Limit: 1})
	fmt.Printf("VerificationGetTxs: err = %v, result = %s\n", err, formatAsJSON(res))

	if err == nil && len(res) > 0 {
		res, err = apiVerificator.VerificationGetTxByHash("")
		fmt.Printf("VerificationGetTxByHash: err = %v, result = %s\n", err, formatAsJSON(res))
	}

	res, err = apiVerificator.VerificationGetCoins(&dscSwagger.OptionalParams{Limit: 1})
	fmt.Printf("VerificationGetCoins: err = %v, result = %s\n", err, formatAsJSON(res))

	res, err = apiVerificator.VerificationGetCoin("del")
	fmt.Printf("VerificationGetCoin: err = %v, result = %s\n", err, formatAsJSON(res))

	res, err = apiVerificator.VerificationGetBlocks(&dscSwagger.OptionalParams{Limit: 1})
	fmt.Printf("VerificationGetBlocks: err = %v, result = %s\n", err, formatAsJSON(res))

	res, err = apiVerificator.VerificationGetBlockByHeight(1)
	fmt.Printf("VerificationGetBlockByHeight: err = %v, result = %s\n", err, formatAsJSON(res))

	res, err = apiVerificator.VerificationGetBlockTransactions(1)
	fmt.Printf("VerificationGetBlockTransactions: err = %v, result = %s\n", err, formatAsJSON(res))

	res, err = apiVerificator.VerificationGetEvmContracts(&dscSwagger.OptionalParams{Limit: 1})
	fmt.Printf("VerificationGetEvmContracts: err = %v, result = %s\n", err, formatAsJSON(res))

	contracts, err := api.GetEvmContracts(&dscApi.OptionalParams{Limit: 1})
	if err == nil && len(contracts) > 0 {
		res, err = apiVerificator.VerificationGetEvmContract(contracts[0].Address)
		fmt.Printf("VerificationGetEvmContract: err = %v, result = %s\n", err, formatAsJSON(res))

		res, err = apiVerificator.VerificationGetEvmContractTransactions(contracts[0].Address, &dscSwagger.OptionalParams{Limit: 1})
		fmt.Printf("VerificationGetEvmContractTransactions: err = %v, result = %s\n", err, formatAsJSON(res))
	}

	res, err = apiVerificator.VerificationGetEvmTransactions(&dscSwagger.OptionalParams{Limit: 1})
	fmt.Printf("VerificationGetEvmTransactions: err = %v, result = %s\n", err, formatAsJSON(res))

	txs, err := api.GetEvmTransactions(&dscApi.OptionalParams{Limit: 1})
	if err == nil && len(txs) > 0 {
		res, err = apiVerificator.VerificationGetEvmTransaction(txs[0].Hash)
		fmt.Printf("VerificationGetEvmTransaction: err = %v, result = %s\n", err, formatAsJSON(res))
	}

	res, err = apiVerificator.VerificationGetEvmAccounts(&dscSwagger.OptionalParams{Limit: 1})
	fmt.Printf("VerificationGetEvmAccounts: err = %v, result = %s\n", err, formatAsJSON(res))

	accs, err := api.GetEvmAccounts(&dscApi.OptionalParams{Limit: 1})
	if err == nil && len(accs) > 0 {
		res, err = apiVerificator.VerificationGetEvmAccount(accs[0].Address)
		fmt.Printf("VerificationGetEvmAccount: err = %v, result = %s\n", err, formatAsJSON(res))
	}

	res, err = apiVerificator.VerificationGetEvmContractEvents("", &dscSwagger.OptionalParams{Limit: 1})
	fmt.Printf("VerificationGetEvmContractEvents: err = %v, result = %s\n", err, formatAsJSON(res))

	res, err = apiVerificator.VerificationGetEvmAccountBalances("", &dscSwagger.OptionalParams{Limit: 1})
	fmt.Printf("VerificationGetEvmAccountBalances: err = %v, result = %s\n", err, formatAsJSON(res))

	res, err = apiVerificator.VerificationGetValidatorsByKind("validator")
	fmt.Printf("VerificationGetValidatorsByKind: err = %v, result = %s\n", err, formatAsJSON(res))

	// res, err = apiVerificator.VerificationGetValidatorsCoins("del", &dscSwagger.OptionalParams{Limit: 1})
	// fmt.Printf("VerificationGetValidatorsCoins: err = %v, result = %s\n", err, formatAsJSON(res))

	vals, err := api.GetValidatorsByKind("validator")
	if err == nil && len(vals) > 0 {
		res, err = apiVerificator.VerificationGetValidator(vals[0].Address)
		fmt.Printf("VerificationGetValidator: err = %v, result = %s\n", err, formatAsJSON(res))

		res, err = apiVerificator.VerificationGetValidatorStakes(vals[0].Address, &dscSwagger.OptionalParams{Limit: 1})
		fmt.Printf("VerificationGetValidatorStakes: err = %v, result = %s\n", err, formatAsJSON(res))

		//res, err = apiVerificator.VerificationGetValidatorStakesNFT(vals[0].Address)
		//fmt.Printf("VerificationGetValidatorStakesNFT: err = %v, result = %s\n", err, formatAsJSON(res))
	}
}

// helper function
func formatAsJSON(obj interface{}) string {
	objStr, err := json.MarshalIndent(obj, "", "    ")
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%s\n", objStr)
}

func sampleSendCoins(api *dscApi.API) {
	const mnemonic1 = "quarter cook oven describe orient clip clay credit degree purpose disease depart collect bonus crane hover key accuse scare afford settle tourist sing humor"
	const mnemonic2 = "zone math funny unfold burger achieve foot uncover guilt vivid load bind pizza space silk void judge hub wild slot gossip stem plate enable"

	acc1, err := dscWallet.NewAccountFromMnemonicWords(mnemonic1, "")
	if err != nil {
		fmt.Printf("Create acc 1 error: %v\n", err)
		return
	}

	acc2, err := dscWallet.NewAccountFromMnemonicWords(mnemonic2, "")
	if err != nil {
		fmt.Printf("Create acc 2 error: %v\n", err)
		return
	}

	sendCoin(api, acc1, acc2.Address())
	sendCoin(api, acc2, acc1.Address())
}

// This is the sample of transaction sending
func sendCoin(api *dscApi.API, senderWallet *dscWallet.Account, receiver string) {
	// 1. set valid chain ID, account number, account sequence (nonce) for Sender
	num, seq, err := api.GetAccountNumberAndSequence(senderWallet.Address())
	if err != nil {
		fmt.Printf("GetAccountNumberAndSequence(%s) error: %v\n", senderWallet.Address(), err)
		return
	}
	senderWallet = senderWallet.WithChainID(api.ChainID()).WithSequence(seq).WithAccountNumber(num)

	// 2. prepare message
	// example of use Cosmos SDK standart functions: sdk.NewCoin, math.NewInt

	receiverAddress, err := sdk.AccAddressFromBech32(receiver)
	if err != nil {
		fmt.Printf("sdk.AccAddressFromBech32(%s) error: %v\n", receiver, err)
		return
	}

	msg := dscTx.NewMsgSendCoin(
		senderWallet.SdkAddress(),
		sdk.NewCoin("del", dscApi.EtherToWei(math.NewInt(1))),
		receiverAddress,
	)

	// 3. build transaction
	tx, err := dscTx.BuildTransaction(
		senderWallet,
		[]sdk.Msg{msg},
		"go sdk test", // any transaction memo
		// fee to pay for transaction
		// if amount = 0, amount will be calculated and collected automaticaly by validator
		sdk.NewCoin("del", sdk.NewInt(0)),
	)
	if err != nil {
		fmt.Printf("Create tx error: %v\n", err)
		return
	}

	// 4. sign transaction and serialize to bytes
	err = tx.SignTransaction(senderWallet)
	if err != nil {
		fmt.Printf("Sign tx error: %v\n", err)
		return
	}
	bz, err := tx.BytesToSend()
	if err != nil {
		fmt.Printf("Bytes tx error: %v\n", err)
		return
	}

	// 5. send transaction bytes to blockchain node
	// 1) BroadcastTxSync: send transaction and get transaction hash and
	// possible error of transaction check
	// You can check transaction delivery by hash
	// 2) BroadcastTxCommit: same as BroadcastTxSync, but WAIT
	// for delivery and end of block (about 5 seconds)
	result, err := api.BroadcastTxSync(bz)
	if err != nil {
		fmt.Printf("BroadcastTxSync error: %v\n", err)
		return
	}
	fmt.Printf("Send result: %s\n", formatAsJSON(result))

	// This is dumb method to wait for delivery
	// You can send multiple transactions, accumulate hashes and check
	// all transactions later
	for i := 0; i < 6; i++ {
		txInfo, err := api.GetTxByHash(result.Hash)
		if err != nil {
			fmt.Printf("GetTxByHash error: %v\n", err)
			time.Sleep(time.Second)
		} else {
			fmt.Printf("TxInfo: %v\n", formatAsJSON(txInfo))
			break
		}
	}
}

func printBlockchainInfo(api *dscApi.API) {
	printCoins(api)
	printTx(api)
	printBlocks(api)
	printAddressInfo(api)
	printTxInfo(api)
	printEvmAccounts(api)
	printEvmContracts(api)
	printEvmTransactions(api)
	printGetValidatorsByKind(api)
}

func printCoins(api *dscApi.API) {
	coins, err := api.GetCoins(nil)
	if err != nil {
		fmt.Printf("GetCoins() error: %v\n", err)
		return
	}
	for _, coin := range coins {
		fmt.Printf("%s\n", formatAsJSON(coin))
	}
}

func printTx(api *dscApi.API) {
	txs, err := api.GetTxs(nil)
	if err != nil {
		fmt.Printf("GetTxs() error: %v\n", err)
	} else {
		for i, tx := range txs {
			fmt.Printf("GetTxs() %d result:\n%s\n", i, formatAsJSON(tx))
		}
	}
}

func printBlocks(api *dscApi.API) {
	for block := uint64(0); block < 10000; block += 1000 {
		blockInfo, err := api.GetBlockByHeight(block)
		if err != nil {
			fmt.Printf("GetBlockByHeight() error: %v\n", err)
		} else {
			fmt.Printf("GetBlockByHeight() result:\n%s\n", formatAsJSON(blockInfo))
		}
	}
}

func printAddressInfo(api *dscApi.API) {
	for _, adr := range []string{"dx15cv03c4e2dvnc8cg72eaec4fv08pxzgkmr255d", "dx184qe86tyhurv5fxlxgvcwa6znfg3ugk8ajn4r3"} {
		inf, err := api.GetAddress(adr)
		if err != nil {
			fmt.Printf("GetAddress() error: %v\n", err)
		} else {
			fmt.Printf("GetAddress() %s result:\n%s\n", adr, formatAsJSON(inf))
		}
	}
}

func printTxInfo(api *dscApi.API) {
	for _, hash := range []string{"0236FD82E1CAA67C7C3023B26E27F8EBDA3475C47936A4E5F61C7D655D5B39B2",
		"D355B19F7958DC76454BCD057715D232DED4634458AF1A7F64FDADB0FBBB6699"} {
		tx, err := api.GetTxByHash(hash)
		if err != nil {
			fmt.Printf("GetTxByHash() error: %v\n", err)
		} else {
			fmt.Printf("GetTxByHash() %s result:\n%s\n", hash, formatAsJSON(tx))
		}
	}
}

func printEvmAccounts(api *dscApi.API) {
	accs, err := api.GetEvmAccounts(nil)
	if err != nil {
		fmt.Printf("GetEvmAccounts() error: %v\n", err)
	} else {
		for i, acc := range accs {
			fmt.Printf("GetEvmAccounts() %d result:\n%s\n", i, formatAsJSON(acc))
		}
	}
}

func printEvmContracts(api *dscApi.API) {
	contracts, err := api.GetEvmContracts(nil)
	if err != nil {
		fmt.Printf("GetEvmContracts() error: %v\n", err)
	} else {
		for i, cntr := range contracts {
			fmt.Printf("GetEvmContracts() %d result:\n%s\n", i, formatAsJSON(cntr))
		}
	}
}

func printEvmTransactions(api *dscApi.API) {
	txs, err := api.GetEvmTransactions(nil)
	if err != nil {
		fmt.Printf("GetEvmTransactions() error: %v\n", err)
	} else {
		for i, tx := range txs {
			fmt.Printf("GetEvmTransactions() %d result:\n%s\n", i, formatAsJSON(tx))
		}
	}
}

func printGetValidatorsByKind(api *dscApi.API) {
	validators, err := api.GetValidatorsByKind("validator")
	if err != nil {
		fmt.Printf("GetValidatorsByKind() error: %v\n", err)
	} else {
		for _, val := range validators {
			stakes, err := api.GetValidatorStakes(val.Address, nil)
			if err != nil {
				fmt.Printf("GetValidatorStakes(%s) error: %v\n", val.Address, err)
			} else {
				for _, st := range stakes {
					fmt.Printf("stake = %s\n", formatAsJSON(st))
				}
			}
		}
	}
}
