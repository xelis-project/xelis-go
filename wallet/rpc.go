package wallet

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/xelis-project/xelis-go-sdk/data"
	"github.com/xelis-project/xelis-go-sdk/rpc"
	"github.com/xelis-project/xelis-go-sdk/wallet/methods"
)

type RPC struct {
	http *rpc.Http
}

func setAuthHeader(header http.Header, username string, password string) {
	auth := fmt.Sprintf("%s:%s", username, password)
	buf := bytes.Buffer{}
	encoder := base64.NewEncoder(base64.StdEncoding, &buf)
	encoder.Write([]byte(auth))
	encoder.Close()

	header.Set("Authorization", fmt.Sprintf("Basic %s", buf.String()))
}

func NewRPC(url string, username string, password string) (*RPC, error) {
	header := http.Header{}
	setAuthHeader(header, username, password)
	http, err := rpc.NewHttp(url, header)
	if err != nil {
		return nil, err
	}

	daemon := &RPC{
		http,
	}

	return daemon, nil
}

func (d *RPC) GetVersion() (version string, err error) {
	_, err = d.http.Request(methods.GetVersion, nil, &version)
	return
}

func (d *RPC) GetNetwork() (network string, err error) {
	_, err = d.http.Request(methods.GetNetwork, nil, &network)
	return
}

func (d *RPC) GetNonce() (nonce uint64, err error) {
	_, err = d.http.Request(methods.GetNonce, nil, &nonce)
	return
}

func (d *RPC) GetTopoheight() (topoheight uint64, err error) {
	_, err = d.http.Request(methods.GetTopoheight, nil, &topoheight)
	return
}

func (d *RPC) GetAddress(params GetAddressParams) (address string, err error) {
	_, err = d.http.Request(methods.GetAddress, params, &address)
	return
}

func (d *RPC) SplitAddress(params SplitAddressParams) (result SplitAddressResult, err error) {
	_, err = d.http.Request(methods.SplitAddress, params, &result)
	return
}

func (d *RPC) Rescan(params RescanParams) (success bool, err error) {
	_, err = d.http.Request(methods.Rescan, params, &success)
	return
}

func (d *RPC) GetBalance(params GetBalanceParams) (balance uint64, err error) {
	_, err = d.http.Request(methods.GetBalance, params, &balance)
	return
}

func (d *RPC) HasBalance(params GetBalanceParams) (exists bool, err error) {
	_, err = d.http.Request(methods.HasBalance, params, &exists)
	return
}

func (d *RPC) GetTrackedAssets() (assets []string, err error) {
	_, err = d.http.Request(methods.GetTrackedAssets, nil, &assets)
	return
}

func (d *RPC) GetAssetPrecision(params GetAssetPrecisionParams) (decimals int, err error) {
	_, err = d.http.Request(methods.GetAssetPrecision, params, &decimals)
	return
}

func (d *RPC) GetAssets() (assets map[string]Asset, err error) {
	_, err = d.http.Request(methods.GetAssets, nil, &assets)
	return
}

func (d *RPC) GetAsset(params GetAssetParams) (asset Asset, err error) {
	_, err = d.http.Request(methods.GetAsset, params, &asset)
	return
}

func (d *RPC) GetTransaction(params GetTransactionParams) (transaction TransactionEntry, err error) {
	_, err = d.http.Request(methods.GetTransaction, params, &transaction)
	return
}

func (d *RPC) BuildTransaction(params BuildTransactionParams) (result TransactionResponse, err error) {
	if err = checkFeeBuilder(params.Fee); err != nil {
		return
	}

	_, err = d.http.Request(methods.BuildTransaction, params, &result)
	return
}

func (d *RPC) BuildTransactionOffline(params BuildTransactionOfflineParams) (result TransactionResponse, err error) {
	_, err = d.http.Request(methods.BuildTransactionOffline, params, &result)
	return
}

func (d *RPC) BuildUnsignedTransaction(params BuildUnsignedTransactionParams) (result UnsignedTransactionResponse, err error) {
	_, err = d.http.Request(methods.BuildUnsignedTransaction, params, &result)
	return
}

func (d *RPC) SignUnsignedTransaction(params SignUnsignedTransactionParams) (result SignatureId, err error) {
	_, err = d.http.Request(methods.SignUnsignedTransaction, params, &result)
	return
}

func (d *RPC) FinalizeUnsignedTransaction(params FinalizeUnsignedTransactionParams) (result TransactionResponse, err error) {
	_, err = d.http.Request(methods.FinalizeUnsignedTransaction, params, &result)
	return
}

func (d *RPC) ClearTxCache() (result bool, err error) {
	_, err = d.http.Request(methods.ClearTxCache, nil, &result)
	return
}

func (d *RPC) ListTransactions(params ListTransactionsParams) (txs []TransactionEntry, err error) {
	_, err = d.http.Request(methods.ListTransactions, params, &txs)
	return
}

func (d *RPC) IsOnline() (online bool, err error) {
	_, err = d.http.Request(methods.IsOnline, nil, &online)
	return
}

func (d *RPC) SetOnlineMode() (success bool, err error) {
	_, err = d.http.Request(methods.SetOnlineMode, nil, &success)
	return
}

func (d *RPC) SetOfflineMode() (success bool, err error) {
	_, err = d.http.Request(methods.SetOfflineMode, nil, &success)
	return
}

func (d *RPC) SignData(data data.Element) (signature string, err error) {
	_, err = d.http.Request(methods.SignData, data, &signature)
	return
}

func (d *RPC) EstimateFees(params EstimateFeesParams) (amount uint64, err error) {
	_, err = d.http.Request(methods.EstimateFees, params, &amount)
	return
}

func (d *RPC) EstimateExtraDataSize(params EstimateExtraDataSizeParams) (result EstimateExtraDataSizeResult, err error) {
	_, err = d.http.Request(methods.EstimateExtraDataSize, params, &result)
	return
}

func (d *RPC) NetworkInfo() (result NetworkInfoResult, err error) {
	_, err = d.http.Request(methods.NetworkInfo, nil, &result)
	return
}

func (d *RPC) DecryptExtraData(params DecryptExtraDataParams) (result PlaintextExtraData, err error) {
	_, err = d.http.Request(methods.DecryptExtraData, params, &result)
	return
}

func (d *RPC) DecryptCiphertext(params DecryptCiphertextParams) (result uint64, err error) {
	_, err = d.http.Request(methods.DecryptCiphertext, params, &result)
	return
}

func (d *RPC) GetMatchingKeys(params GetMatchingKeysParams) (result []interface{}, err error) {
	_, err = d.http.Request(methods.GetMatchingKeys, params, &result)
	return
}

func (d *RPC) CountMatchingEntries(params CountMatchingEntriesParams) (result uint64, err error) {
	_, err = d.http.Request(methods.CountMatchingEntries, params, &result)
	return
}

func (d *RPC) GetValueFromKey(params GetValueFromKeyParams) (result interface{}, err error) {
	_, err = d.http.Request(methods.GetValueFromKey, params, &result)
	return
}

func (d *RPC) Store(params StoreParams) (result bool, err error) {
	_, err = d.http.Request(methods.Store, params, &result)
	return
}

func (d *RPC) Delete(params StoreParams) (result bool, err error) {
	_, err = d.http.Request(methods.Delete, params, &result)
	return
}

func (d *RPC) DeleteTreeEntries(params DeleteTreeEntriesParams) (result bool, err error) {
	_, err = d.http.Request(methods.DeleteTreeEntries, params, &result)
	return
}

func (d *RPC) HasKey(params HasKeyParams) (result bool, err error) {
	_, err = d.http.Request(methods.HasKey, params, &result)
	return
}

func (d *RPC) QueryDB(params QueryDBParams) (result QueryResult, err error) {
	_, err = d.http.Request(methods.QueryDB, params, &result)
	return
}

func checkFeeBuilder(fee *FeeBuilder) error {
	if fee != nil && fee.Multiplier != nil && fee.Value != nil {
		return fmt.Errorf("you cannot set both Multiplier and Value in FeeBuilder")
	}

	return nil
}
