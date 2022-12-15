package api

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/go-resty/resty/v2"
)

// API is a struct implementing DSC API iteraction.
type API struct {
	client *resty.Client

	// Parameters
	chainID   string
	baseDenom string
}

// NewAPI creates Decimal API instance.
func NewAPI(apiURL string) *API {
	initConfig()
	// denom is detected by apiURL
	var baseDenom = DevnetBaseCoin
	if apiURL == MainnetGate {
		baseDenom = MainnetBaseCoin
	}
	if apiURL == TestnetGate {
		baseDenom = TestnetBaseCoin
	}
	return &API{
		client:    resty.New().SetBaseURL(apiURL).SetTimeout(time.Minute),
		baseDenom: baseDenom,
	}
}

func (api *API) GetParameters() error {
	// request
	res, err := api.client.R().Get("/rpc/genesis/chain")
	if err = processConnectionError(res, err); err != nil {
		return err
	}
	// json decode
	respValue := string(res.Body())
	// process results
	api.chainID = respValue
	return nil
}

// Address requests full information about specified address
func (api *API) GetAccountNumberAndSequence(address string) (uint64, uint64, error) {
	type respDirectAddress struct {
		Account struct {
			BaseAccount struct {
				Number   string `json:"account_number"`
				Sequence string `json:"sequence"`
			} `json:"base_account"`
		} `json:"account"`
	}
	// request
	res, err := api.client.R().Get(fmt.Sprintf("/rpc/auth/accounts/%s", address))
	if err = processConnectionError(res, err); err != nil {
		return 0, 0, err
	}
	// json decode
	respValue, respErr := respDirectAddress{}, Error{}
	err = universalJSONDecode(res.Body(), &respValue, &respErr, func() (bool, bool) {
		return respValue.Account.BaseAccount.Number > "", respErr.StatusCode != 0
	})
	//empty account (no transactions), it's normal
	if err != nil {
		return 0, 0, joinErrors(err, respErr)
	}
	// process result
	accNumber, _ := strconv.ParseUint(respValue.Account.BaseAccount.Number, 10, 64)
	seq, _ := strconv.ParseUint(respValue.Account.BaseAccount.Sequence, 10, 64)

	return accNumber, seq, nil
}

// Address requests full information about specified address
func (api *API) GetAccountBalance(address string) (sdk.Coins, error) {
	type respBalance struct {
		Ok     bool `json:"ok"`
		Result map[string]struct {
			Amount string `json:"amount"`
		} `json:"result"`
	}
	// request
	res, err := api.client.R().Get(fmt.Sprintf("/address/%s/balances", address))
	if err = processConnectionError(res, err); err != nil {
		return sdk.NewCoins(), err
	}
	// json decode
	respValue, respErr := respBalance{}, Error{}
	err = universalJSONDecode(res.Body(), &respValue, &respErr, func() (bool, bool) {
		return respValue.Ok, respErr.StatusCode != 0
	})
	if err != nil {
		return sdk.NewCoins(), joinErrors(err, respErr)
	}
	// process result
	result := sdk.NewCoins()
	for denom, v := range respValue.Result {
		amount, ok := sdk.NewIntFromString(v.Amount)
		if !ok {
			return sdk.NewCoins(), fmt.Errorf("can't convert amount to int")
		}
		result = result.Add(sdk.NewCoin(denom, amount))
	}
	return result, nil
}

// BaseCoin() returns base coin symbol from genesis. Need for correct transaction building
func (api *API) BaseCoin() string {
	return api.baseDenom
}

// ChainID() returns blockchain network chain id
func (api *API) ChainID() string {
	return api.chainID
}

// Response of broadcast_tx_sync
type TxResponse struct {
	// transaction hash
	Hash string
	// error info. Code = 0 mean no error
	Code      int
	Log       string
	Codespace string
}

func (api *API) BroadcastTxSync(data []byte) (*TxResponse, error) {
	type directSyncResponse struct {
		Result struct {
			Code      int    `json:"code"`
			Hash      string `json:"hash"`
			Log       string `json:"log"`
			Codespace string `json:"codespace"`
		} `json:"result"`
	}
	// request
	res, err := api.client.R().SetHeader("Content-Type", "application/json").
		SetBody(map[string]string{
			"hexTx": hex.EncodeToString(data),
		}).Post("/rpc/txs")
	if err = processConnectionError(res, err); err != nil {
		return nil, err
	}
	// json decode
	respValue := directSyncResponse{}
	err = universalJSONDecode(res.Body(), &respValue, nil, func() (bool, bool) {
		return respValue.Result.Hash > "", false
	})
	if err != nil {
		return nil, err
	}
	// process result
	return &TxResponse{
		Code:      respValue.Result.Code,
		Hash:      respValue.Result.Hash,
		Log:       respValue.Result.Log,
		Codespace: respValue.Result.Codespace,
	}, nil
}

func (api *API) CalculateFee(data []byte, denom string) (sdk.Coin, error) {
	type calculateResponse struct {
		Ok     bool `json:"ok"`
		Result struct {
			Commission string `json:"commission"`
		} `json:"result"`
	}
	// request
	res, err := api.client.R().SetHeader("Content-Type", "application/json").
		SetBody(map[string]string{
			"tx_bytes": hex.EncodeToString(data),
			"denom":    denom,
		}).Post("/tx/estimate")
	if err = processConnectionError(res, err); err != nil {
		return sdk.Coin{}, err
	}
	// json decode
	respValue := calculateResponse{}
	err = universalJSONDecode(res.Body(), &respValue, nil, func() (bool, bool) {
		return respValue.Ok, false
	})
	if err != nil {
		return sdk.Coin{}, err
	}
	// process result
	amount, ok := sdk.NewIntFromString(respValue.Result.Commission)
	if !ok {
		return sdk.Coin{}, fmt.Errorf("can't convert commission to int")
	}
	return sdk.NewCoin(denom, amount), nil
}

/*
func (api *API) BroadcastTxCommit(data []byte) (*TxResponse, error) {
	type directSyncResponse struct {
		Result struct {
			CheckTx struct {
				Code      int    `json:"code"`
				Log       string `json:"log"`
				Codespace string `json:"codespace"`
			} `json:"check_tx"`
			DeliverTx struct {
				Code      int    `json:"code"`
				Log       string `json:"log"`
				Codespace string `json:"codespace"`
			} `json:"deliver_tx"`
			Hash string `json:"hash"`
		} `json:"result"`
	}
	// request
	res, err := api.rpc.R().Get("/broadcast_tx_commit?tx=0x" + hex.EncodeToString(data))
	if err = processConnectionError(res, err); err != nil {
		return nil, err
	}
	// json decode
	respValue := directSyncResponse{}
	err = universalJSONDecode(res.Body(), &respValue, nil, func() (bool, bool) {
		return respValue.Result.Hash > "", false
	})
	if err != nil {
		return nil, err
	}
	if respValue.Result.CheckTx.Code != 0 {
		return &TxResponse{
			Code:      respValue.Result.CheckTx.Code,
			Hash:      respValue.Result.Hash,
			Log:       respValue.Result.CheckTx.Log,
			Codespace: respValue.Result.CheckTx.Codespace,
		}, nil
	}
	// process result
	return &TxResponse{
		Code:      respValue.Result.DeliverTx.Code,
		Hash:      respValue.Result.Hash,
		Log:       respValue.Result.DeliverTx.Log,
		Codespace: respValue.Result.DeliverTx.Codespace,
	}, nil
}
*/

// Init global cosmos sdk config
// Do not seal config or rework to use sealed config
func initConfig() {
	cfg := sdk.GetConfig()
	cfg.SetBech32PrefixForAccount("d0", "d0pub")
}
