package wallet

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	netUrl "net/url"

	"github.com/creachadair/jrpc2"
	"github.com/creachadair/jrpc2/jhttp"
	"github.com/xelis-project/xelis-go-sdk/wallet/methods"
)

type RPC struct {
	ctx    context.Context
	Client *jrpc2.Client
}

type AuthTransport struct {
	Transport http.RoundTripper
	Username  string
	Password  string
}

func (t *AuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	setAuthHeader(req.Header, t.Username, t.Password)
	return t.Transport.RoundTrip(req)
}

func setAuthHeader(header http.Header, username string, password string) {
	auth := fmt.Sprintf("%s:%s", username, password)
	buf := bytes.Buffer{}
	encoder := base64.NewEncoder(base64.StdEncoding, &buf)
	encoder.Write([]byte(auth))
	encoder.Close()

	header.Set("Authorization", fmt.Sprintf("Basic %s", buf.String()))
}

func NewRPC(ctx context.Context, url string, username string, password string) (*RPC, error) {
	daemonUrl, err := netUrl.Parse(url)
	if err != nil {
		return nil, err
	}

	channel := jhttp.NewChannel(daemonUrl.String(), &jhttp.ChannelOptions{
		Client: &http.Client{
			Transport: &AuthTransport{
				Transport: http.DefaultTransport,
				Username:  username,
				Password:  password,
			},
		},
	})
	rpcClient := jrpc2.NewClient(channel, nil)

	daemon := &RPC{
		ctx:    ctx,
		Client: rpcClient,
	}

	return daemon, nil
}

func (d *RPC) GetVersion() (version string, err error) {
	err = d.Client.CallResult(d.ctx, methods.GetVersion, nil, &version)
	return
}

func (d *RPC) GetNetwork() (network string, err error) {
	err = d.Client.CallResult(d.ctx, methods.GetNetwork, nil, &network)
	return
}

func (d *RPC) GetNonce() (nonce uint64, err error) {
	err = d.Client.CallResult(d.ctx, methods.GetNonce, nil, &nonce)
	return
}

func (d *RPC) GetTopoheight() (topoheight uint64, err error) {
	err = d.Client.CallResult(d.ctx, methods.GetTopoheight, nil, &topoheight)
	return
}

func (d *RPC) GetAddress(params GetAddressParams) (address string, err error) {
	err = d.Client.CallResult(d.ctx, methods.GetAddress, params, &address)
	return
}

func (d *RPC) SplitAddress(params SplitAddressParams) (result SplitAddressResult, err error) {
	err = d.Client.CallResult(d.ctx, methods.SplitAddress, params, &result)
	return
}

func (d *RPC) Rescan(params RescanParams) (success bool, err error) {
	err = d.Client.CallResult(d.ctx, methods.Rescan, params, &success)
	return
}

func (d *RPC) GetBalance(params GetBalanceParams) (balance uint64, err error) {
	err = d.Client.CallResult(d.ctx, methods.GetBalance, params, &balance)
	return
}

func (d *RPC) HasBalance(params GetBalanceParams) (exists bool, err error) {
	err = d.Client.CallResult(d.ctx, methods.HasBalance, params, &exists)
	return
}

func (d *RPC) GetTrackedAssets() (assets []string, err error) {
	err = d.Client.CallResult(d.ctx, methods.GetTrackedAssets, nil, &assets)
	return
}

func (d *RPC) GetAssetPrecision(params GetAssetPrecisionParams) (decimals int, err error) {
	err = d.Client.CallResult(d.ctx, methods.GetAssetPrecision, params, &decimals)
	return
}

func (d *RPC) GetAssets() (assets []string, err error) {
	err = d.Client.CallResult(d.ctx, methods.GetAssets, nil, &assets)
	return
}

func (d *RPC) GetAsset(params GetAssetPrecisionParams) (asset string, err error) {
	err = d.Client.CallResult(d.ctx, methods.GetAsset, params, &asset)
	return
}

func (d *RPC) GetTransaction(params GetTransactionParams) (transaction TransactionEntry, err error) {
	err = d.Client.CallResult(d.ctx, methods.GetTransaction, params, &transaction)
	return
}

func (d *RPC) BuildTransaction(params BuildTransactionParams) (result TransactionResponse, err error) {
	if err = checkFeeBuilder(params.Fee); err != nil {
		return
	}

	err = d.Client.CallResult(d.ctx, methods.BuildTransaction, params, &result)
	return
}

func (d *RPC) BuildTransactionOffline(params BuildTransactionOfflineParams) (result TransactionResponse, err error) {
	err = d.Client.CallResult(d.ctx, methods.BuildTransactionOffline, params, &result)
	return
}

func (d *RPC) BuildUnsignedTransaction(params BuildUnsignedTransactionParams) (result UnsignedTransactionResponse, err error) {
	err = d.Client.CallResult(d.ctx, methods.BuildUnsignedTransaction, params, &result)
	return
}

func (d *RPC) SignUnsignedTransaction(params SignUnsignedTransactionParams) (result SignatureId, err error) {
	err = d.Client.CallResult(d.ctx, methods.SignUnsignedTransaction, params, &result)
	return
}

func (d *RPC) FinalizeUnsignedTransaction(params FinalizeUnsignedTransactionParams) (result TransactionResponse, err error) {
	err = d.Client.CallResult(d.ctx, methods.FinalizeUnsignedTransaction, params, &result)
	return
}

func (d *RPC) ClearTxCache() (result bool, err error) {
	err = d.Client.CallResult(d.ctx, methods.ClearTxCache, nil, &result)
	return
}

func (d *RPC) ListTransactions(params ListTransactionsParams) (txs []TransactionEntry, err error) {
	err = d.Client.CallResult(d.ctx, methods.ListTransactions, params, &txs)
	return
}

func (d *RPC) IsOnline() (online bool, err error) {
	err = d.Client.CallResult(d.ctx, methods.IsOnline, nil, &online)
	return
}

func (d *RPC) SetOnlineMode() (success bool, err error) {
	err = d.Client.CallResult(d.ctx, methods.SetOnlineMode, nil, &success)
	return
}

func (d *RPC) SetOfflineMode() (success bool, err error) {
	err = d.Client.CallResult(d.ctx, methods.SetOfflineMode, nil, &success)
	return
}

func (d *RPC) SignData(data interface{}) (signature string, err error) {
	err = d.Client.CallResult(d.ctx, methods.SignData, data, &signature)
	return
}

func (d *RPC) EstimateFees(params EstimateFeesParams) (amount uint64, err error) {
	err = d.Client.CallResult(d.ctx, methods.EstimateFees, params, &amount)
	return
}

func (d *RPC) EstimateExtraDataSize(params EstimateExtraDataSizeParams) (result EstimateExtraDataSizeResult, err error) {
	err = d.Client.CallResult(d.ctx, methods.EstimateExtraDataSize, params, &result)
	return
}

func (d *RPC) NetworkInfo() (result NetworkInfoResult, err error) {
	err = d.Client.CallResult(d.ctx, methods.NetworkInfo, nil, &result)
	return
}

func (d *RPC) DecryptExtraData(params DecryptExtraDataParams) (result PlaintextExtraData, err error) {
	err = d.Client.CallResult(d.ctx, methods.DecryptExtraData, params, &result)
	return
}

func (d *RPC) DecryptCiphertext(params DecryptCiphertextParams) (result uint64, err error) {
	err = d.Client.CallResult(d.ctx, methods.DecryptCiphertext, nil, &result)
	return
}

func (d *RPC) GetMatchingKeys(params GetMatchingKeysParams) (result []interface{}, err error) {
	err = d.Client.CallResult(d.ctx, methods.GetMatchingKeys, params, &result)
	return
}

func (d *RPC) CountMatchingEntries(params CountMatchingEntriesParams) (result uint64, err error) {
	err = d.Client.CallResult(d.ctx, methods.CountMatchingEntries, params, &result)
	return
}

func (d *RPC) GetValueFromKey(params GetValueFromKeyParams) (result interface{}, err error) {
	err = d.Client.CallResult(d.ctx, methods.GetValueFromKey, params, &result)
	return
}

func (d *RPC) Store(params StoreParams) (result bool, err error) {
	err = d.Client.CallResult(d.ctx, methods.Store, params, &result)
	return
}

func (d *RPC) Delete(params StoreParams) (result bool, err error) {
	err = d.Client.CallResult(d.ctx, methods.Delete, params, &result)
	return
}

func (d *RPC) DeleteTreeEntries(params DeleteTreeEntriesParams) (result bool, err error) {
	err = d.Client.CallResult(d.ctx, methods.DeleteTreeEntries, params, &result)
	return
}

func (d *RPC) HasKey(params HasKeyParams) (result bool, err error) {
	err = d.Client.CallResult(d.ctx, methods.HasKey, params, &result)
	return
}

func (d *RPC) QueryDB(params QueryDBParams) (result QueryResult, err error) {
	err = d.Client.CallResult(d.ctx, methods.QueryDB, params, &result)
	return
}

func checkFeeBuilder(fee *FeeBuilder) error {
	if fee != nil && fee.Multiplier != nil && fee.Value != nil {
		return fmt.Errorf("you cannot set both Multiplier and Value in FeeBuilder")
	}

	return nil
}
