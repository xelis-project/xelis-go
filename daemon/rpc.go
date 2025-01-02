package daemon

import (
	"context"
	netUrl "net/url"

	"github.com/creachadair/jrpc2"
	"github.com/creachadair/jrpc2/jhttp"
	"github.com/xelis-project/xelis-go-sdk/daemon/methods"
)

type RPC struct {
	ctx    context.Context
	Client *jrpc2.Client
}

func NewRPC(ctx context.Context, url string) (*RPC, error) {
	daemonUrl, err := netUrl.Parse(url)
	if err != nil {
		return nil, err
	}

	channel := jhttp.NewChannel(daemonUrl.String(), nil)
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

func (d *RPC) GetInfo() (result GetInfoResult, err error) {
	err = d.Client.CallResult(d.ctx, methods.GetInfo, nil, &result)
	return
}

func (d *RPC) GetHeight() (height uint64, err error) {
	err = d.Client.CallResult(d.ctx, methods.GetHeight, nil, &height)
	return
}

func (d *RPC) GetTopoheight() (topoheight uint64, err error) {
	err = d.Client.CallResult(d.ctx, methods.GetTopoHeight, nil, &topoheight)
	return
}

func (d *RPC) GetStableHeight() (stableheight uint64, err error) {
	err = d.Client.CallResult(d.ctx, methods.GetStableHeight, nil, &stableheight)
	return
}

func (d *RPC) GetStableTopoheight() (topoheight uint64, err error) {
	err = d.Client.CallResult(d.ctx, methods.GetStableTopoheight, nil, &topoheight)
	return
}

func (d *RPC) GetStableBalance(params GetBalanceParams) (result GetStableBalanceResult, err error) {
	err = d.Client.CallResult(d.ctx, methods.GetStableBalance, params, &result)
	return
}

func (d *RPC) GetBlockTemplate(addr string) (result GetBlockTemplateResult, err error) {
	params := map[string]string{"address": addr}
	err = d.Client.CallResult(d.ctx, methods.GetBlockTemplate, params, &result)
	return
}

func (d *RPC) GetBlockAtTopoheight(params GetBlockAtTopoheightParams) (block Block, err error) {
	err = d.Client.CallResult(d.ctx, methods.GetBlockAtTopoheight, params, &block)
	return
}

func (d *RPC) GetBlocksAtHeight(params GetBlocksAtHeightParams) (blocks []Block, err error) {
	err = d.Client.CallResult(d.ctx, methods.GetBlocksAtHeight, params, &blocks)
	return
}

func (d *RPC) GetBlockByHash(params GetBlockByHashParams) (block Block, err error) {
	err = d.Client.CallResult(d.ctx, methods.GetBlockByHash, params, &block)
	return
}

func (d *RPC) GetTopBlock(params GetTopBlockParams) (block Block, err error) {
	err = d.Client.CallResult(d.ctx, methods.GetTopBlock, params, &block)
	return
}

func (d *RPC) GetNonce(addr string) (nonce GetNonceResult, err error) {
	params := map[string]string{"address": addr}
	err = d.Client.CallResult(d.ctx, methods.GetNonce, params, &nonce)
	return
}

func (d *RPC) HasNonce(addr string) (hasNonce bool, err error) {
	params := map[string]string{"address": addr}
	var result map[string]bool
	err = d.Client.CallResult(d.ctx, methods.HasNonce, params, &result)
	hasNonce = result["exist"]
	return
}

func (d *RPC) GetNonceAtTopoheight(params GetNonceAtTopoheightParams) (nonce VersionedNonce, err error) {
	err = d.Client.CallResult(d.ctx, methods.GetNonceAtTopoheight, params, &nonce)
	return
}

func (d *RPC) GetBalance(params GetBalanceParams) (balance GetBalanceResult, err error) {
	err = d.Client.CallResult(d.ctx, methods.GetBalance, params, &balance)
	return
}

func (d *RPC) HasBalance(params GetBalanceParams) (hasBalance bool, err error) {
	var result map[string]bool
	err = d.Client.CallResult(d.ctx, methods.HasBalance, params, &result)
	hasBalance = result["exists"]
	return
}

func (d *RPC) GetBalanceAtTopoheight(params GetBalanceAtTopoheightParams) (balance VersionedBalance, err error) {
	err = d.Client.CallResult(d.ctx, methods.GetBalanceAtTopoheight, params, &balance)
	return
}

func (d *RPC) GetAsset(assetId string) (asset Asset, err error) {
	params := map[string]string{"asset": assetId}
	err = d.Client.CallResult(d.ctx, methods.GetAsset, params, &asset)
	return
}

func (d *RPC) GetAssets(params GetAssetsParams) (assets []AssetWithData, err error) {
	err = d.Client.CallResult(d.ctx, methods.GetAssets, params, &assets)
	return
}

func (d *RPC) CountAssets() (count uint64, err error) {
	err = d.Client.CallResult(d.ctx, methods.CountAssets, nil, &count)
	return
}

func (d *RPC) CountTransactions() (count uint64, err error) {
	err = d.Client.CallResult(d.ctx, methods.CountTransactions, nil, &count)
	return
}

func (d *RPC) CountAccounts() (count uint64, err error) {
	err = d.Client.CallResult(d.ctx, methods.CountAccounts, nil, &count)
	return
}

func (d *RPC) GetTips() (tips []string, err error) {
	err = d.Client.CallResult(d.ctx, methods.GetTips, nil, &tips)
	return
}

func (d *RPC) P2PStatus() (status P2PStatusResult, err error) {
	err = d.Client.CallResult(d.ctx, methods.P2PStatus, nil, &status)
	return
}

func (d *RPC) GetDAGOrder(params GetTopoheightRangeParams) (hashes []string, err error) {
	err = d.Client.CallResult(d.ctx, methods.GetDAGOrder, params, &hashes)
	return
}

func (d *RPC) SubmitBlock(params SubmitBlockParams) (result bool, err error) {
	err = d.Client.CallResult(d.ctx, methods.SubmitBlock, params, &result)
	return
}

func (d *RPC) SubmitTransaction(data string) (result bool, err error) {
	params := map[string]string{"data": data}
	err = d.Client.CallResult(d.ctx, methods.SubmitTransaction, params, &result)
	return
}

func (d *RPC) GetMempool() (txs []Transaction, err error) {
	err = d.Client.CallResult(d.ctx, methods.GetMempool, nil, &txs)
	return
}

func (d *RPC) GetTransaction(hash string) (tx Transaction, err error) {
	params := map[string]string{"hash": hash}
	err = d.Client.CallResult(d.ctx, methods.GetTransaction, params, &tx)
	return
}

func (d *RPC) GetTransactions(params GetTransactionsParams) (txs []Transaction, err error) {
	err = d.Client.CallResult(d.ctx, methods.GetTransactions, params, &txs)
	return
}

func (d *RPC) GetBlocksRangeByTopoheight(params GetTopoheightRangeParams) (blocks []Block, err error) {
	err = d.Client.CallResult(d.ctx, methods.GetBlocksRangeByTopoheight, params, &blocks)
	return
}

func (d *RPC) GetBlocksRangeByHeight(params GetHeightRangeParams) (blocks []Block, err error) {
	err = d.Client.CallResult(d.ctx, methods.GetBlocksRangeByHeight, params, &blocks)
	return
}

func (d *RPC) GetAccounts(params GetAccountsParams) (addresses []string, err error) {
	err = d.Client.CallResult(d.ctx, methods.GetAccounts, params, &addresses)
	return
}

func (d *RPC) GetAccountHistory(addr string) (history []AccountHistory, err error) {
	params := map[string]string{"address": addr}
	err = d.Client.CallResult(d.ctx, methods.GetAccountHistory, params, &history)
	return
}

func (d *RPC) GetAccountAssets(addr string) (assets []string, err error) {
	params := map[string]string{"address": addr}
	err = d.Client.CallResult(d.ctx, methods.GetAccountAssets, params, &assets)
	return
}

func (d *RPC) GetPeers() (result GetPeersResult, err error) {
	err = d.Client.CallResult(d.ctx, methods.GetPeers, nil, &result)
	return
}

func (d *RPC) GetDevFeeThresholds() (fees []Fee, err error) {
	err = d.Client.CallResult(d.ctx, methods.GetDevFeeThresholds, nil, &fees)
	return
}

func (d *RPC) GetSizeOnDisk() (sizeOnDisk SizeOnDisk, err error) {
	err = d.Client.CallResult(d.ctx, methods.GetSizeOnDisk, nil, &sizeOnDisk)
	return
}

func (d *RPC) IsTxExecutedInBlock(params IsTxExecutedInBlockParams) (executed bool, err error) {
	err = d.Client.CallResult(d.ctx, methods.IsTxExecutedInBlock, params, &executed)
	return
}

func (d *RPC) GetAccountRegistrationTopoheight(addr string) (topoheight uint64, err error) {
	params := map[string]string{"address": addr}
	err = d.Client.CallResult(d.ctx, methods.GetAccountRegistrationTopoheight, params, &topoheight)
	return
}

func (d *RPC) IsAccountRegistered(params IsAccountRegisteredParams) (exists bool, err error) {
	err = d.Client.CallResult(d.ctx, methods.IsAccountRegistered, params, &exists)
	return
}

func (d *RPC) GetDifficulty() (result GetDifficultyResult, err error) {
	err = d.Client.CallResult(d.ctx, methods.GetDifficulty, nil, &result)
	return
}

func (d *RPC) ValidateAddress(params ValidateAddressParams) (result ValidateAddressResult, err error) {
	err = d.Client.CallResult(d.ctx, methods.ValidateAddress, params, &result)
	return
}

func (d *RPC) ExtractKeyFromAddress(params ExtractKeyFromAddressParams) (key interface{}, err error) {
	err = d.Client.CallResult(d.ctx, methods.ExtractKeyFromAddress, params, &key)
	return
}

func (d *RPC) GetMinerWork(params GetMinerWorkParams) (result GetMinerWorkResult, err error) {
	err = d.Client.CallResult(d.ctx, methods.GetMinerWork, params, &result)
	return
}

func (d *RPC) SplitAddress(params SplitAddressParams) (result SplitAddressResult, err error) {
	err = d.Client.CallResult(d.ctx, methods.SplitAddress, params, &result)
	return
}

func (d *RPC) GetHardForks() (result []HardFork, err error) {
	err = d.Client.CallResult(d.ctx, methods.GetHardForks, nil, &result)
	return
}

func (d *RPC) GetEstimatedFeeRates() (result FeeRatesEstimated, err error) {
	err = d.Client.CallResult(d.ctx, methods.GetEstimatedFeeRates, nil, &result)
	return
}

func (d *RPC) GetPrunedTopoheight() (result uint64, err error) {
	err = d.Client.CallResult(d.ctx, methods.GetPrunedTopoheight, nil, &result)
	return
}

func (d *RPC) GetTransactionExecutor(params GetTransactionExecutorParams) (result GetTransactionExecutorResult, err error) {
	err = d.Client.CallResult(d.ctx, methods.GetTransactionExecutor, params, &result)
	return
}

func (d *RPC) HasMultisigAtTopoheight(params HasMultisigAtTopoheightParams) (result bool, err error) {
	err = d.Client.CallResult(d.ctx, methods.HasMultisigAtTopoheight, params, &result)
	return
}

func (d *RPC) GetMultisigAtTopoheight(params GetMultisigAtTopoheightParams) (result GetMultisigAtTopoHeightResult, err error) {
	err = d.Client.CallResult(d.ctx, methods.GetMultisigAtTopoheight, params, &result)
	return
}

func (d *RPC) GetMultisig(params GetMultisigParams) (result GetMultisigResult, err error) {
	err = d.Client.CallResult(d.ctx, methods.GetMultisig, params, &result)
	return
}

func (d *RPC) HasMultisig(params HasMultisigParams) (result bool, err error) {
	err = d.Client.CallResult(d.ctx, methods.HasMultisig, params, &result)
	return
}

func (d *RPC) GetContractOutputs(params GetContractOutputsParams) (result []ContractOutput, err error) {
	var outputs []interface{}
	err = d.Client.CallResult(d.ctx, methods.GetContractOutputs, params, &outputs)
	result = parseContractOutputs(outputs)
	return
}

func (d *RPC) GetContractModule(params GetContractModuleParams) (result GetContractModuleResult, err error) {
	err = d.Client.CallResult(d.ctx, methods.GetContractModule, params, &result)
	return
}

func (d *RPC) GetContractData(params GetContractDataParams) (result interface{}, err error) {
	err = d.Client.CallResult(d.ctx, methods.GetContractData, params, &result)
	return
}

func (d *RPC) GetContractDataAtTopoheight(params GetContractDataAtTopoheightParams) (result interface{}, err error) {
	err = d.Client.CallResult(d.ctx, methods.GetContractDataAtTopoheight, params, &result)
	return
}

func (d *RPC) GetContractBalance(params GetContractBalanceParams) (result interface{}, err error) {
	err = d.Client.CallResult(d.ctx, methods.GetContractBalance, params, &result)
	return
}

func (d *RPC) GetContractBalanceAtTopoheight(params GetContractBalanceAtTopoheightParams) (result interface{}, err error) {
	err = d.Client.CallResult(d.ctx, methods.GetContractBalanceAtTopoheight, params, &result)
	return
}

func (d *RPC) CountContracts() (result uint64, err error) {
	err = d.Client.CallResult(d.ctx, methods.CountAccounts, nil, &result)
	return
}

func (d *RPC) MakeIntegratedAddress(params MakeIntegratedAddressParams) (result uint64, err error) {
	err = d.Client.CallResult(d.ctx, methods.MakeIntegratedAddress, params, &result)
	return
}

func (d *RPC) DecryptExtraData(params DecryptExtraDataParams) (result interface{}, err error) {
	err = d.Client.CallResult(d.ctx, methods.DecryptExtraData, params, &result)
	return
}

func parseContractOutputs(outputs []interface{}) (result []ContractOutput) {
	for _, output := range outputs {
		switch out := output.(type) {
		case map[string]interface{}:
			for key, value := range out {
				switch key {
				case "exit_code":
					exit_code, ok := value.(float64)
					if !ok {
						break
					}

					result = append(result, ContractOutputExitCode{
						ExitCode: uint64(exit_code),
					})
				case "refund_gas":
					refund_gas, ok := value.(map[string]interface{})
					if !ok {
						break
					}

					amount, ok := refund_gas["amount"].(float64)
					if !ok {
						break
					}

					result = append(result, ContractOutputRefundGas{
						Amount: uint64(amount),
					})
				case "transfer":
					transfer, ok := value.(map[string]interface{})
					if !ok {
						break
					}

					amount, ok := transfer["amount"].(float64)
					if !ok {
						break
					}

					asset, ok := transfer["asset"].(string)
					if !ok {
						break
					}

					destination, ok := transfer["destination"].(string)
					if !ok {
						break
					}

					result = append(result, ContractOutputTransfer{
						Amount:      uint64(amount),
						Asset:       asset,
						Destination: destination,
					})
				}
			}
		case string:
			switch out {
			case "refund_deposits":
				result = append(result, ContractOutputRefundDeposits{})
			}
		}
	}

	return
}
