# Decimal Go SDK

For detailed explanation on how things work, checkout the:

- [Decimal SDK docs](https://help.decimalchain.com/api-sdk/).
- [Decimal Console site](https://console.decimalchain.com/).
- [Swagger documentation](https://devnet-dec2-explorer-api.decimalchain.com/api/documentation).

# Install
```
go get bitbucket.org/decimalteam/dsc-go-sdk
```

# Usage

You can see working example in `example/main.go`

## I. Action

Actions are creating wallet, creating and sending transactions.

### 1. Create wallet

```go
package ...

import (
    "fmt"

    dscWallet "bitbucket.org/decimalteam/dsc-go-sdk/wallet"
)

const (
    // PLEASE, DON'T USE THIS MNEMONIC OR ANY PUBLIC EXPOSED MNEMONIC IN MAINNET
    testMnemonicWords      = "repair furnace west loud peasant false six hockey poem tube now alien service phone hazard winter favorite away sand fuel describe version tragic vendor"
	testMnemonicPassphrase = ""
)

func main() {
    // Option 1. Generate private key (account) by mnemonic words (bip39)
    account, err := dscWallet.NewAccountFromMnemonicWords(testMnemonicWords, testMnemonicPassphrase)
	if err != nil {
		// Error handling
	}
    // Output: dx1...
    fmt.Println(account.Address())

    ...
    // Option 2. Generate mnemonic for future use
    mnemonicObject err := NewMnemonic(testMnemonicPassphrase)
	if err != nil {
		// Error handling
	}    
    // print mnemonic words
    fmt.Println(mnemonicObject.Words())

    mnemonic := mnemonicObject.Words()
    account, err := dscWallet.NewAccountFromMnemonicWords(mnemonic, testMnemonicPassphrase)
    ...
}
```

### 2. Create and send transaction

To send transaction you need:

*) Create account (wallet) from known mnemonic

*) Create API instance for account binding

*) Bind account: set chan id, account number, account sequence (nonce)

*) Create transaction message

*) Sign transaction by account, send transaction

*) Verify transaction delivery for sync mode

```go
package ...

import (
    "fmt"

    // Required imports
    dscApi "bitbucket.org/decimalteam/dsc-go-sdk/api"
    dscTx "bitbucket.org/decimalteam/dsc-go-sdk/tx"
    dscWallet "bitbucket.org/decimalteam/dsc-go-sdk/wallet"
    // optional cosmos sdk to work with sdk.Coin and math.Int
    "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func main() {
    // 1. Create wallet (see above)
    account := ...

    // 2. Create API instance for account binding
    // TODO: constructor may change
    api := dscApi.NewAPI(endpoints...)
    err := api.GetParameters()
    // ...error handling
    // now api has valid results for api.ChainID(), api.BaseCoin()

    // 3. Bind account
    accNumber, accSequence, err := api.GetAccountNumberAndSequence(account.Address())
    // ...error handling
    account = account.WithChainID(api.ChainID()).WithSequence(accSequence).WithAccountNumber(accNumber)

    // 4. Create transaction message
    // For possible transaction messages see tx/types.go and DSC source
	msg := dscTx.NewMsgSendCoin(
		account.SdkAddress(),
		sdk.NewCoin("del", dscApi.EtherToWei(math.NewInt(1))),
		receiverAddress,
	)
    // or you can use message type directly
    msg := dscTx.MsgSendCoin{
        Sender: account.Address(),
        Receiver: receiver,
        Coin: sdk.NewCoin("del", dscApi.EtherToWei(math.NewInt(1))),
    }

	tx, err := dscTx.BuildTransaction(
		account,
		[]sdk.Msg{msg},
		"some tx memo",
		// fee to pay for transaction
		// if amount = 0, amount will be calculated and collected automaticaly by validator
		sdk.NewCoin("del", sdk.NewInt(0)),
	)
    // ...error handling

    // 5. Sign and send
    err = tx.SignTransaction(account)
	// ...error handling
	bz, err := tx.BytesToSend()
	// ...error handling

	// 1) BroadcastTxSync: send transaction in SYNC mode and get transaction hash and
	// possible error of transaction check
	// You can check later transaction delivery by hash
	// 2) BroadcastTxCommit: same as BroadcastTxSync, but WAIT
	// for delivery at end of block (about 5 seconds)    
    result, err := api.BroadcastTxSync(bz)
    result, err := api.BroadcastTxCommit(bz)

    // wait for block
    // 6. Verify transaction delivery
    // NOTE: if transaction not in block already, you can get HTTP 404 error
    // If you want to be sure after every transaction, use BroadcastTxCommit
    time.Sleep(time.Second * 6)
    txInfo, err := api.GetTxByHash(result.Hash)
    // ...error handling
}
```

## II. Views

To get some information about blocks, transactions, accounts, etc use GetXXX methods.

All GetXXX methods declared in `api/endpoint.go`, return types declared in `api/types.go`

`OptionalParams` is structure with fields Limit and Offset.
`opt *OptionalParams` can be nil.
If `opt *OptionalParams` is not nil, GetXXX gets part of data, specified by Limit and Offset.

List of known methods:

- GetAddress(id string) (*AddressInfo, error)
- GetAddressTxs(id string, opt *OptionalParams) ([]TxInfo, error)
- GetAllNFT(opt *OptionalParams) ([]NFT, error)
- GetNFTCollection(collection string) (interface{}, error)
- GetTxByHash(hash string) (*TxInfo, error)
- GetTxs(opt *OptionalParams) ([]TxInfo, error)
- GetCoins(opt *OptionalParams) ([]CoinInfo, error)
- GetCoin(coin string) (*CoinInfo, error)
- GetBlockByHeight(height uint64) (*BlockInfo, error)
- GetBlockTransactions(height uint64) ([]TxInfo, error)
- GetEvmContracts(opt *OptionalParams) ([]EvmContract, error)
- GetEvmTransactions(opt *OptionalParams) ([]EvmTransaction, error)
- GetEvmAccounts(opt *OptionalParams) ([]EvmAccount, error)
- GetEvmContract(address string) (*EvmContract, error)
- GetEvmTransaction(hash string) (*EvmTransaction, error)
- GetEvmAccount(address string) (*EvmAccount, error)
- GetEvmContractTransactions(address string, opt *OptionalParams) ([]EvmTransaction, error)
- GetEvmContractEvents(address string, opt *OptionalParams) ([]EvmEvent, error)
- GetEvmAccountBalances(address string, opt *OptionalParams) ([]EvmAccountBalance, error)
- GetValidatorsByKind(kind string) ([]Validator, error)
- GetValidatorsCoins(address string, opt *OptionalParams) ([]ValidatorStakedCoin, error)
- GetValidator(address string) (*Validator, error)
- GetValidatorStakes(address string, opt *OptionalParams) ([]ValidatorStake, error)
- GetValidatorStakesNFT(address string) ([]ValidatorStakeNFT, error)