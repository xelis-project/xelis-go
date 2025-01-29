package wallet

import (
	"net/http"

	"github.com/xelis-project/xelis-go-sdk/daemon"
	"github.com/xelis-project/xelis-go-sdk/data"
	"github.com/xelis-project/xelis-go-sdk/rpc"
	"github.com/xelis-project/xelis-go-sdk/wallet/events"
	"github.com/xelis-project/xelis-go-sdk/wallet/methods"
)

type WebSocket struct {
	Prefix string
	WS     *rpc.WebSocket
}

func NewWebSocket(endpoint string, username string, password string) (*WebSocket, error) {
	header := make(http.Header)
	setAuthHeader(header, username, password)
	ws, err := rpc.NewWebSocket(endpoint, header)
	if err != nil {
		return nil, err
	}

	daemonWS := &WebSocket{
		WS: ws,
	}

	return daemonWS, nil
}

func (w *WebSocket) BatchCall(requests []rpc.RPCRequest, result []interface{}) (res []rpc.RPCResponse, errs []error) {
	return w.WS.BatchCall(requests, result)
}

func (w *WebSocket) Close() error {
	return w.WS.Close()
}

func (w *WebSocket) CloseEvent(event string) error {
	return w.WS.CloseEvent(event)
}

func (w *WebSocket) ConnectionErr() chan error {
	return w.WS.ConnectionErr
}

func (w *WebSocket) NewTopoheightChannel() (chan uint64, chan error, error) {
	chanResult := make(chan uint64)
	chanErr := make(chan error)

	err := w.WS.ListenEventFunc(w.Prefix+events.NewTopoheight, func(res rpc.RPCResponse) {
		var result map[string]interface{}
		err := rpc.ParseResponseResult(res, &result)
		if err != nil {
			chanErr <- err
		} else {
			chanResult <- uint64(result["topoheight"].(float64))
		}
	})

	return chanResult, chanErr, err
}

func (w *WebSocket) NewTopoheightFunc(onData func(uint64, error)) error {
	return w.WS.ListenEventFunc(w.Prefix+events.NewTopoheight, func(res rpc.RPCResponse) {
		var result map[string]interface{}
		err := rpc.ParseResponseResult(res, &result)
		topoheight := uint64(result["topoheight"].(float64))
		onData(topoheight, err)
	})
}

func (w *WebSocket) NewAssetChannel() (chan daemon.AssetData, chan error, error) {
	chanResult := make(chan daemon.AssetData)
	chanErr := make(chan error)

	err := w.WS.ListenEventFunc(w.Prefix+events.NewAsset, func(res rpc.RPCResponse) {
		var result daemon.AssetData
		err := rpc.ParseResponseResult(res, &result)
		if err != nil {
			chanErr <- err
		} else {
			chanResult <- result
		}
	})

	return chanResult, chanErr, err
}

func (w *WebSocket) NewAssetFunc(onData func(daemon.AssetData, error)) error {
	return w.WS.ListenEventFunc(w.Prefix+events.NewAsset, func(res rpc.RPCResponse) {
		var result daemon.AssetData
		err := rpc.ParseResponseResult(res, &result)
		onData(result, err)
	})
}

func (w *WebSocket) NewTransactionChannel() (chan TransactionEntry, chan error, error) {
	chanResult := make(chan TransactionEntry)
	chanErr := make(chan error)

	err := w.WS.ListenEventFunc(w.Prefix+events.NewTransaction, func(res rpc.RPCResponse) {
		var result TransactionEntry
		err := rpc.ParseResponseResult(res, &result)
		if err != nil {
			chanErr <- err
		} else {
			chanResult <- result
		}
	})

	return chanResult, chanErr, err
}

func (w *WebSocket) NewTransactionFunc(onData func(TransactionEntry, error)) error {
	return w.WS.ListenEventFunc(w.Prefix+events.NewTransaction, func(res rpc.RPCResponse) {
		var result TransactionEntry
		err := rpc.ParseResponseResult(res, &result)
		onData(result, err)
	})
}

func (w *WebSocket) BalanceChangedChannel() (chan BalanceChangedResult, chan error, error) {
	chanResult := make(chan BalanceChangedResult)
	chanErr := make(chan error)

	err := w.WS.ListenEventFunc(w.Prefix+events.BalanceChanged, func(res rpc.RPCResponse) {
		var result BalanceChangedResult
		err := rpc.ParseResponseResult(res, &result)
		if err != nil {
			chanErr <- err
		} else {
			chanResult <- result
		}
	})

	return chanResult, chanErr, err
}

func (w *WebSocket) BalanceChangedFunc(onData func(BalanceChangedResult, error)) error {
	return w.WS.ListenEventFunc(w.Prefix+events.BalanceChanged, func(res rpc.RPCResponse) {
		var result BalanceChangedResult
		err := rpc.ParseResponseResult(res, &result)
		onData(result, err)
	})
}

func (w *WebSocket) RescanChannel() (chan uint64, chan error, error) {
	chanResult := make(chan uint64)
	chanErr := make(chan error)

	err := w.WS.ListenEventFunc(w.Prefix+events.Rescan, func(res rpc.RPCResponse) {
		var result map[string]interface{}
		err := rpc.ParseResponseResult(res, &result)
		if err != nil {
			chanErr <- err
		} else {
			chanResult <- uint64(result["start_topoheight"].(float64))
		}
	})

	return chanResult, chanErr, err
}

func (w *WebSocket) RescanFunc(onData func(uint64, error)) error {
	return w.WS.ListenEventFunc(w.Prefix+events.Rescan, func(res rpc.RPCResponse) {
		var result map[string]interface{}
		err := rpc.ParseResponseResult(res, &result)
		startTopoheight := uint64(result["start_topoheight"].(float64))
		onData(startTopoheight, err)
	})
}

func (w *WebSocket) HistorySyncedChannel() (chan uint64, chan error, error) {
	chanResult := make(chan uint64)
	chanErr := make(chan error)

	err := w.WS.ListenEventFunc(w.Prefix+events.HistorySynced, func(res rpc.RPCResponse) {
		var result map[string]interface{}
		err := rpc.ParseResponseResult(res, &result)
		if err != nil {
			chanErr <- err
		} else {
			chanResult <- uint64(result["topoheight"].(float64))
		}
	})

	return chanResult, chanErr, err
}

func (w *WebSocket) HistorySyncedFunc(onData func(uint64, error)) error {
	return w.WS.ListenEventFunc(w.Prefix+events.HistorySynced, func(res rpc.RPCResponse) {
		var result map[string]interface{}
		err := rpc.ParseResponseResult(res, &result)
		startTopoheight := uint64(result["topoheight"].(float64))
		onData(startTopoheight, err)
	})
}

func (w *WebSocket) OnlineChannel() (chan bool, chan error, error) {
	chanResult := make(chan bool)
	chanErr := make(chan error)

	err := w.WS.ListenEventFunc(w.Prefix+events.Online, func(res rpc.RPCResponse) {
		chanResult <- true
	})

	return chanResult, chanErr, err
}

func (w *WebSocket) OnlineFunc(onData func()) error {
	return w.WS.ListenEventFunc(w.Prefix+events.Online, func(res rpc.RPCResponse) {
		onData()
	})
}

func (w *WebSocket) OfflineChannel() (chan bool, chan error, error) {
	chanResult := make(chan bool)
	chanErr := make(chan error)

	err := w.WS.ListenEventFunc(w.Prefix+events.Offline, func(res rpc.RPCResponse) {
		chanResult <- true
	})

	return chanResult, chanErr, err
}

func (w *WebSocket) OfflineFunc(onData func()) error {
	return w.WS.ListenEventFunc(w.Prefix+events.Offline, func(res rpc.RPCResponse) {
		onData()
	})
}

func (w *WebSocket) GetVersion() (version string, err error) {
	_, err = w.WS.Call(w.Prefix+methods.GetVersion, nil, &version)
	return
}

func (w *WebSocket) GetNetwork() (network string, err error) {
	_, err = w.WS.Call(w.Prefix+methods.GetNetwork, nil, &network)
	return
}

func (w *WebSocket) GetNonce() (nonce uint64, err error) {
	_, err = w.WS.Call(w.Prefix+methods.GetNonce, nil, &nonce)
	return
}

func (w *WebSocket) GetTopoheight() (topoheight uint64, err error) {
	_, err = w.WS.Call(w.Prefix+methods.GetTopoheight, nil, &topoheight)
	return
}

func (w *WebSocket) GetAddress(params GetAddressParams) (address string, err error) {
	_, err = w.WS.Call(w.Prefix+methods.GetAddress, params, &address)
	return
}

func (w *WebSocket) SplitAddress(params SplitAddressParams) (result SplitAddressResult, err error) {
	_, err = w.WS.Call(w.Prefix+methods.SplitAddress, params, &result)
	return
}

func (w *WebSocket) Rescan(params RescanParams) (success bool, err error) {
	_, err = w.WS.Call(w.Prefix+methods.Rescan, params, &success)
	return
}

func (w *WebSocket) GetBalance(params GetBalanceParams) (balance uint64, err error) {
	_, err = w.WS.Call(w.Prefix+methods.Rescan, params, &balance)
	return
}

func (w *WebSocket) HasBalance(params GetBalanceParams) (exists bool, err error) {
	_, err = w.WS.Call(w.Prefix+methods.HasBalance, params, &exists)
	return
}

func (w *WebSocket) GetTrackedAssets() (assets []string, err error) {
	_, err = w.WS.Call(w.Prefix+methods.GetTrackedAssets, nil, &assets)
	return
}

func (w *WebSocket) GetAssetPrecision(params GetAssetPrecisionParams) (decimals int, err error) {
	_, err = w.WS.Call(w.Prefix+methods.GetAssetPrecision, nil, &decimals)
	return
}

func (w *WebSocket) GetAssets() (assets map[string]Asset, err error) {
	_, err = w.WS.Call(w.Prefix+methods.GetAssets, nil, &assets)
	return
}

func (w *WebSocket) GetAsset(params GetAssetPrecisionParams) (asset Asset, err error) {
	_, err = w.WS.Call(w.Prefix+methods.GetAsset, nil, &asset)
	return
}

func (w *WebSocket) GetTransaction(params GetTransactionParams) (transaction TransactionEntry, err error) {
	_, err = w.WS.Call(w.Prefix+methods.GetTransaction, params, &transaction)
	return
}

func (w *WebSocket) BuildTransaction(params BuildTransactionParams) (result TransactionResponse, err error) {
	if err = checkFeeBuilder(params.Fee); err != nil {
		return
	}

	_, err = w.WS.Call(w.Prefix+methods.BuildTransaction, params, &result)
	return
}

func (w *WebSocket) BuildTransactionOffline(params BuildTransactionOfflineParams) (result TransactionResponse, err error) {
	if err = checkFeeBuilder(params.Fee); err != nil {
		return
	}

	_, err = w.WS.Call(w.Prefix+methods.BuildTransactionOffline, params, &result)
	return
}

func (w *WebSocket) BuildUnsignedTransaction(params BuildUnsignedTransactionParams) (result UnsignedTransactionResponse, err error) {
	if err = checkFeeBuilder(params.Fee); err != nil {
		return
	}

	_, err = w.WS.Call(w.Prefix+methods.BuildUnsignedTransaction, params, &result)
	return
}

func (w *WebSocket) SignUnsignedTransaction(params SignUnsignedTransactionParams) (result SignatureId, err error) {
	_, err = w.WS.Call(w.Prefix+methods.SignUnsignedTransaction, params, &result)
	return
}

func (w *WebSocket) FinalizeUnsignedTransaction(params FinalizeUnsignedTransactionParams) (result TransactionResponse, err error) {
	_, err = w.WS.Call(w.Prefix+methods.FinalizeUnsignedTransaction, params, &result)
	return
}

func (w *WebSocket) ClearTxCache() (result bool, err error) {
	_, err = w.WS.Call(w.Prefix+methods.ListTransactions, nil, &result)
	return
}

func (w *WebSocket) ListTransactions(params ListTransactionsParams) (txs []TransactionEntry, err error) {
	_, err = w.WS.Call(w.Prefix+methods.ListTransactions, params, &txs)
	return
}

func (w *WebSocket) IsOnline() (online bool, err error) {
	_, err = w.WS.Call(w.Prefix+methods.IsOnline, nil, &online)
	return
}

func (w *WebSocket) SetOnlineMode() (success bool, err error) {
	_, err = w.WS.Call(w.Prefix+methods.SetOnlineMode, nil, &success)
	return
}

func (w *WebSocket) SetOfflineMode() (success bool, err error) {
	_, err = w.WS.Call(w.Prefix+methods.SetOfflineMode, nil, &success)
	return
}

func (w *WebSocket) SignData(data data.Element) (signature string, err error) {
	_, err = w.WS.Call(w.Prefix+methods.SignData, data, &signature)
	return
}

func (w *WebSocket) EstimateFees(params EstimateFeesParams) (amount uint64, err error) {
	_, err = w.WS.Call(w.Prefix+methods.EstimateFees, params, &amount)
	return
}

func (w *WebSocket) EstimateExtraDataSize(params EstimateExtraDataSizeParams) (result EstimateExtraDataSizeResult, err error) {
	_, err = w.WS.Call(w.Prefix+methods.EstimateExtraDataSize, params, &result)
	return
}

func (w *WebSocket) NetworkInfo() (result NetworkInfoResult, err error) {
	_, err = w.WS.Call(w.Prefix+methods.NetworkInfo, nil, &result)
	return
}

func (w *WebSocket) DecryptExtraData(params DecryptExtraDataParams) (result interface{}, err error) {
	_, err = w.WS.Call(w.Prefix+methods.DecryptExtraData, params, &result)
	return
}

func (w *WebSocket) DecryptCiphertext(params DecryptCiphertextParams) (result uint64, err error) {
	_, err = w.WS.Call(w.Prefix+methods.DecryptCiphertext, params, &result)
	return
}

func (w *WebSocket) GetMatchingKeys(params GetMatchingKeysParams) (result []interface{}, err error) {
	_, err = w.WS.Call(w.Prefix+methods.GetMatchingKeys, params, &result)
	return
}

func (w *WebSocket) CountMatchingEntries(params CountMatchingEntriesParams) (result uint64, err error) {
	_, err = w.WS.Call(w.Prefix+methods.CountMatchingEntries, params, &result)
	return
}

func (w *WebSocket) GetValueFromKey(params GetValueFromKeyParams) (result interface{}, err error) {
	_, err = w.WS.Call(w.Prefix+methods.GetValueFromKey, params, &result)
	return
}

func (w *WebSocket) Store(params StoreParams) (result bool, err error) {
	_, err = w.WS.Call(w.Prefix+methods.Store, params, &result)
	return
}

func (w *WebSocket) Delete(params DeleteParams) (result bool, err error) {
	_, err = w.WS.Call(w.Prefix+methods.Delete, params, &result)
	return
}

func (w *WebSocket) DeleteTreeEntries(params DeleteTreeEntriesParams) (result bool, err error) {
	_, err = w.WS.Call(w.Prefix+methods.DeleteTreeEntries, params, &result)
	return
}

func (w *WebSocket) HasKey(params HasKeyParams) (result bool, err error) {
	_, err = w.WS.Call(w.Prefix+methods.HasKey, params, &result)
	return
}

func (w *WebSocket) QueryDB(params QueryDBParams) (result QueryResult, err error) {
	_, err = w.WS.Call(w.Prefix+methods.QueryDB, params, &result)
	return
}
