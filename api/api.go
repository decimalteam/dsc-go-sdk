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
	rpc    *resty.Client
	rest   *resty.Client

	// Parameters
	chainID        string
	baseCoinSymbol string
}

// NewAPI creates Decimal API instance.
func NewAPI(apiURL, rpcURL, restURL string) *API {
	initConfig()
	return &API{
		client: resty.New().SetHostURL(apiURL).SetTimeout(time.Minute),
		rpc:    resty.New().SetHostURL(rpcURL).SetTimeout(time.Minute),
		rest:   resty.New().SetHostURL(restURL).SetTimeout(time.Minute),
	}
}

func (api *API) GetParameters() error {
	type respDirectGenesis struct {
		Result struct {
			Genesis struct {
				ChainID  string `json:"chain_id"`
				AppState struct {
					Coin struct {
						Params struct {
							BaseSymbol string `json:"base_symbol"`
						} `json:"params"`
					} `json:"coin"`
				} `json:"app_state"`
			} `json:"genesis"`
		} `json:"result"`
	}
	// request
	res, err := api.rpc.R().Get("/genesis")
	if err = processConnectionError(res, err); err != nil {
		return err
	}
	// json decode
	respValue := respDirectGenesis{}
	err = universalJSONDecode(res.Body(), &respValue, nil, func() (bool, bool) {
		return respValue.Result.Genesis.ChainID > "", false
	})
	if err != nil {
		return err
	}
	// process results
	api.chainID = respValue.Result.Genesis.ChainID
	api.baseCoinSymbol = respValue.Result.Genesis.AppState.Coin.Params.BaseSymbol
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
	res, err := api.rest.R().Get(fmt.Sprintf("/cosmos/auth/v1beta1/accounts/%s", address))
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

// BaseCoin() returns base coin symbol from genesis. Need for correct transaction building
func (api *API) BaseCoin() string {
	return api.baseCoinSymbol
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
	res, err := api.rpc.R().Get("/broadcast_tx_sync?tx=0x" + hex.EncodeToString(data))
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

// Init global cosmos sdk config
// Do not seal config or rework to use sealed config
func initConfig() {
	cfg := sdk.GetConfig()
	cfg.SetBech32PrefixForAccount("dx", "dxpub")
}
