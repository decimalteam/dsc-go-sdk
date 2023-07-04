package api

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"io"
	"math/big"
	"net/http"
	"net/url"
)

const (
	byteCodePath      = DevnetGate + "/evm-token/data?"
	devNetValPath     = "https://devnet-val.decimalchain.com/web3/"
	privateKeyAddress = ""
)

type Payload struct {
	Name      string
	Symbol    string
	Supply    string
	MaxSupply string
	Mintable  string
	Burnable  string
	Capped    string
}

type Response struct {
	Ok     bool
	Result string
}

func getPayload() *Payload {
	return &Payload{
		Name:      "testR",
		Symbol:    "RTTA",
		Supply:    "100",
		MaxSupply: "1000000",
		Mintable:  "true",
		Burnable:  "true",
		Capped:    "false",
	}
}

func getBytecode(path string, payload *Payload) (*Response, error) {
	request, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, fmt.Errorf("get request error: %s", err)
	}

	request.Header = http.Header{
		"Content-type": {"application/json"},
	}

	urlValues := url.Values{}
	urlValues.Add("name", payload.Name)
	urlValues.Add("symbol", payload.Symbol)
	urlValues.Add("supply", payload.Supply)
	urlValues.Add("maxSupply", payload.MaxSupply)
	urlValues.Add("mintable", payload.Mintable)
	urlValues.Add("burnable", payload.Burnable)
	urlValues.Add("capped", payload.Capped)

	request.URL.RawQuery = urlValues.Encode()
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("get request error: %s", err)
	}
	defer response.Body.Close()

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body error: %s", err)
	}

	var result Response
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func sendTokenPayload() {
	payload := getPayload()

	txData, err := getBytecode(byteCodePath, payload)
	if err != nil {
		fmt.Printf("getBytecode error: %v\n", err)
	}

	client, err := ethclient.Dial(devNetValPath)
	if err != nil {
		fmt.Printf("ethclient.Dial error: %v\n", err)
	}

	sendTx(client, txData.Result)
}

func sendTx(client *ethclient.Client, txData string) {

	privateKey, err := crypto.HexToECDSA(privateKeyAddress)
	if err != nil {
		fmt.Printf("crypto.HexToECDSA error: %v\n", err)
		return
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Printf("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		fmt.Printf("crypto.HexToECDSA error: %v\n", err)
		return
	}

	value := big.NewInt(0)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Printf("client.SuggestGasPrice error: %v\n", err)
		return
	}

	amount := new(big.Int)
	amount.SetString("100000000000000000000", 10)

	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      1000000000,
		Value:    value,
		Data:     []byte(txData),
	})

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		fmt.Printf("client.NetworkID error: %v\n", err)
		return
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		fmt.Printf("types.SignTx error: %v\n", err)
		return
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		fmt.Printf("client.SendTransaction error: %v\n", err)
		return
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}