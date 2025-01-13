package daemon

import (
	"github.com/xelis-project/xelis-go-sdk/daemon/events"
	"github.com/xelis-project/xelis-go-sdk/daemon/methods"
	"github.com/xelis-project/xelis-go-sdk/rpc"
)

type WebSocket struct {
	Prefix string
	WS     *rpc.WebSocket
}

func NewWebSocket(endpoint string) (*WebSocket, error) {
	ws, err := rpc.NewWebSocket(endpoint, nil)
	if err != nil {
		return nil, err
	}

	return &WebSocket{
		WS: ws,
	}, nil
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

func (w *WebSocket) NewBlockChannel() (chan Block, chan error, error) {
	chanResult := make(chan Block)
	chanErr := make(chan error)

	err := w.WS.ListenEventFunc(events.NewBlock, func(res rpc.RPCResponse) {
		var result Block
		err := rpc.JsonFormatResponse(res, nil, &result)
		if err != nil {
			chanErr <- err
		} else {
			chanResult <- result
		}
	})

	return chanResult, chanErr, err
}

func (w *WebSocket) NewBlockFunc(onData func(Block, error)) error {
	return w.WS.ListenEventFunc(events.NewBlock, func(res rpc.RPCResponse) {
		var result Block
		err := rpc.JsonFormatResponse(res, nil, &result)
		onData(result, err)
	})
}

func (w *WebSocket) TransactionAddedInMempoolChannel() (chan Transaction, chan error, error) {
	chanResult := make(chan Transaction)
	chanErr := make(chan error)

	err := w.WS.ListenEventFunc(events.TransactionAddedInMempool, func(res rpc.RPCResponse) {
		var result Transaction
		err := rpc.JsonFormatResponse(res, nil, &result)
		if err != nil {
			chanErr <- err
		} else {
			chanResult <- result
		}
	})

	return chanResult, chanErr, err
}

func (w *WebSocket) TransactionAddedInMempoolFunc(onData func(Transaction, error)) error {
	return w.WS.ListenEventFunc(events.TransactionAddedInMempool, func(res rpc.RPCResponse) {
		var result Transaction
		err := rpc.JsonFormatResponse(res, nil, &result)
		onData(result, err)
	})
}

func (w *WebSocket) BlockOrderedChannel() (chan BlockOrderedEvent, chan error, error) {
	chanResult := make(chan BlockOrderedEvent)
	chanErr := make(chan error)

	err := w.WS.ListenEventFunc(events.BlockOrdered, func(res rpc.RPCResponse) {
		var result BlockOrderedEvent
		err := rpc.JsonFormatResponse(res, nil, &result)
		if err != nil {
			chanErr <- err
		} else {
			chanResult <- result
		}
	})

	return chanResult, chanErr, err
}

func (w *WebSocket) BlockOrderedFunc(onData func(BlockOrderedEvent, error)) error {
	return w.WS.ListenEventFunc(events.BlockOrdered, func(res rpc.RPCResponse) {
		var result BlockOrderedEvent
		err := rpc.JsonFormatResponse(res, nil, &result)
		onData(result, err)
	})
}

func (w *WebSocket) TransactionExecutedChannel() (chan TransactionExecutedEvent, chan error, error) {
	chanResult := make(chan TransactionExecutedEvent)
	chanErr := make(chan error)

	err := w.WS.ListenEventFunc(events.TransactionExecuted, func(res rpc.RPCResponse) {
		var result TransactionExecutedEvent
		err := rpc.JsonFormatResponse(res, nil, &result)
		if err != nil {
			chanErr <- err
		} else {
			chanResult <- result
		}
	})

	return chanResult, chanErr, err
}

func (w *WebSocket) TransactionExecutedFunc(onData func(TransactionExecutedEvent, error)) error {
	return w.WS.ListenEventFunc(events.TransactionExecuted, func(res rpc.RPCResponse) {
		var result TransactionExecutedEvent
		err := rpc.JsonFormatResponse(res, nil, &result)
		onData(result, err)
	})
}

func (w *WebSocket) PeerConnectedChannel() (chan Peer, chan error, error) {
	chanResult := make(chan Peer)
	chanErr := make(chan error)

	err := w.WS.ListenEventFunc(events.PeerConnected, func(res rpc.RPCResponse) {
		var result Peer
		err := rpc.JsonFormatResponse(res, nil, &result)
		if err != nil {
			chanErr <- err
		} else {
			chanResult <- result
		}
	})

	return chanResult, chanErr, err
}

func (w *WebSocket) PeerConnectedFunc(onData func(Peer, error)) error {
	return w.WS.ListenEventFunc(events.PeerConnected, func(res rpc.RPCResponse) {
		var result Peer
		err := rpc.JsonFormatResponse(res, nil, &result)
		onData(result, err)
	})
}

func (w *WebSocket) PeerDisconnectedChannel() (chan Peer, chan error, error) {
	chanResult := make(chan Peer)
	chanErr := make(chan error)

	err := w.WS.ListenEventFunc(events.PeerDisconnected, func(res rpc.RPCResponse) {
		var result Peer
		err := rpc.JsonFormatResponse(res, nil, &result)
		if err != nil {
			chanErr <- err
		} else {
			chanResult <- result
		}
	})

	return chanResult, chanErr, err
}

func (w *WebSocket) PeerDisconnectedFunc(onData func(Peer, error)) error {
	return w.WS.ListenEventFunc(events.PeerDisconnected, func(res rpc.RPCResponse) {
		var result Peer
		err := rpc.JsonFormatResponse(res, nil, &result)
		onData(result, err)
	})
}

func (w *WebSocket) PeerStateUpdatedChannel() (chan Peer, chan error, error) {
	chanResult := make(chan Peer)
	chanErr := make(chan error)

	err := w.WS.ListenEventFunc(events.PeerStateUpdated, func(res rpc.RPCResponse) {
		var result Peer
		err := rpc.JsonFormatResponse(res, nil, &result)
		if err != nil {
			chanErr <- err
		} else {
			chanResult <- result
		}
	})

	return chanResult, chanErr, err
}

func (w *WebSocket) PeerStateUpdatedFunc(onData func(Peer, error)) error {
	return w.WS.ListenEventFunc(events.PeerStateUpdated, func(res rpc.RPCResponse) {
		var result Peer
		err := rpc.JsonFormatResponse(res, nil, &result)
		onData(result, err)
	})
}

func (w *WebSocket) BlockOrphanedChannel() (chan BlockOrphanedEvent, chan error, error) {
	chanResult := make(chan BlockOrphanedEvent)
	chanErr := make(chan error)

	err := w.WS.ListenEventFunc(events.BlockOrphaned, func(res rpc.RPCResponse) {
		var result BlockOrphanedEvent
		err := rpc.JsonFormatResponse(res, nil, &result)
		if err != nil {
			chanErr <- err
		} else {
			chanResult <- result
		}
	})

	return chanResult, chanErr, err
}

func (w *WebSocket) BlockOrphanedFunc(onData func(BlockOrphanedEvent, error)) error {
	return w.WS.ListenEventFunc(events.BlockOrphaned, func(res rpc.RPCResponse) {
		var result BlockOrphanedEvent
		err := rpc.JsonFormatResponse(res, nil, &result)
		onData(result, err)
	})
}

func (w *WebSocket) StableHeightChangedChannel() (chan StableHeightChangedEvent, chan error, error) {
	chanResult := make(chan StableHeightChangedEvent)
	chanErr := make(chan error)

	err := w.WS.ListenEventFunc(events.StableHeightChanged, func(res rpc.RPCResponse) {
		var result StableHeightChangedEvent
		err := rpc.JsonFormatResponse(res, nil, &result)
		if err != nil {
			chanErr <- err
		} else {
			chanResult <- result
		}
	})

	return chanResult, chanErr, err
}

func (w *WebSocket) StableHeightChangedFunc(onData func(StableHeightChangedEvent, error)) error {
	return w.WS.ListenEventFunc(events.StableHeightChanged, func(res rpc.RPCResponse) {
		var result StableHeightChangedEvent
		err := rpc.JsonFormatResponse(res, nil, &result)
		onData(result, err)
	})
}

func (w *WebSocket) StableTopoHeightChangedChannel() (chan StableTopoHeightChangedEvent, chan error, error) {
	chanResult := make(chan StableTopoHeightChangedEvent)
	chanErr := make(chan error)

	err := w.WS.ListenEventFunc(events.StableTopoHeightChanged, func(res rpc.RPCResponse) {
		var result StableTopoHeightChangedEvent
		err := rpc.JsonFormatResponse(res, nil, &result)
		if err != nil {
			chanErr <- err
		} else {
			chanResult <- result
		}
	})

	return chanResult, chanErr, err
}

func (w *WebSocket) StableTopoHeightChangedFunc(onData func(StableTopoHeightChangedEvent, error)) error {
	return w.WS.ListenEventFunc(events.StableTopoHeightChanged, func(res rpc.RPCResponse) {
		var result StableTopoHeightChangedEvent
		err := rpc.JsonFormatResponse(res, nil, &result)
		onData(result, err)
	})
}

func (w *WebSocket) PeerPeerListUpdatedChannel() (chan PeerPeerListUpdatedEvent, chan error, error) {
	chanResult := make(chan PeerPeerListUpdatedEvent)
	chanErr := make(chan error)

	err := w.WS.ListenEventFunc(events.PeerPeerListUpdated, func(res rpc.RPCResponse) {
		var result PeerPeerListUpdatedEvent
		err := rpc.JsonFormatResponse(res, nil, &result)
		if err != nil {
			chanErr <- err
		} else {
			chanResult <- result
		}
	})

	return chanResult, chanErr, err
}

func (w *WebSocket) PeerPeerListUpdatedFunc(onData func(PeerPeerListUpdatedEvent, error)) error {
	return w.WS.ListenEventFunc(events.PeerPeerListUpdated, func(res rpc.RPCResponse) {
		var result PeerPeerListUpdatedEvent
		err := rpc.JsonFormatResponse(res, nil, &result)
		onData(result, err)
	})
}

func (w *WebSocket) PeerPeerDisconnectedChannel() (chan PeerPeerDisconnectedEvent, chan error, error) {
	chanResult := make(chan PeerPeerDisconnectedEvent)
	chanErr := make(chan error)

	err := w.WS.ListenEventFunc(events.PeerPeerDisconnected, func(res rpc.RPCResponse) {
		var result PeerPeerDisconnectedEvent
		err := rpc.JsonFormatResponse(res, nil, &result)
		if err != nil {
			chanErr <- err
		} else {
			chanResult <- result
		}
	})

	return chanResult, chanErr, err
}

func (w *WebSocket) PeerPeerDisconnectedFunc(onData func(PeerPeerDisconnectedEvent, error)) error {
	return w.WS.ListenEventFunc(events.PeerPeerDisconnected, func(res rpc.RPCResponse) {
		var result PeerPeerDisconnectedEvent
		err := rpc.JsonFormatResponse(res, nil, &result)
		onData(result, err)
	})
}

func (w *WebSocket) GetVersion() (version string, err error) {
	res, err := w.WS.Call(w.Prefix+methods.GetVersion, nil)
	err = rpc.JsonFormatResponse(res, err, &version)
	return
}

func (w *WebSocket) GetInfo() (result GetInfoResult, err error) {
	res, err := w.WS.Call(w.Prefix+methods.GetInfo, nil)
	err = rpc.JsonFormatResponse(res, err, &result)
	return
}

func (w *WebSocket) GetHeight() (height uint64, err error) {
	res, err := w.WS.Call(w.Prefix+methods.GetHeight, nil)
	err = rpc.JsonFormatResponse(res, err, &height)
	return
}

func (w *WebSocket) GetTopoheight() (topoheight uint64, err error) {
	res, err := w.WS.Call(w.Prefix+methods.GetTopoHeight, nil)
	err = rpc.JsonFormatResponse(res, err, &topoheight)
	return
}

func (w *WebSocket) GetStableHeight() (stableheight uint64, err error) {
	res, err := w.WS.Call(w.Prefix+methods.GetStableHeight, nil)
	err = rpc.JsonFormatResponse(res, err, &stableheight)
	return
}

func (w *WebSocket) GetStableTopoheight() (topoheight uint64, err error) {
	res, err := w.WS.Call(w.Prefix+methods.GetStableTopoheight, nil)
	err = rpc.JsonFormatResponse(res, err, &topoheight)
	return
}

func (w *WebSocket) GetStableBalance(params GetBalanceParams) (result GetStableBalanceResult, err error) {
	res, err := w.WS.Call(w.Prefix+methods.GetStableBalance, nil)
	err = rpc.JsonFormatResponse(res, err, &result)
	return
}

func (w *WebSocket) GetBlockTemplate(addr string) (result GetBlockTemplateResult, err error) {
	params := map[string]string{"address": addr}
	res, err := w.WS.Call(w.Prefix+methods.GetBlockTemplate, params)
	err = rpc.JsonFormatResponse(res, err, &result)
	return
}

func (w *WebSocket) GetBlockAtTopoheight(params GetBlockAtTopoheightParams) (block Block, err error) {
	res, err := w.WS.Call(w.Prefix+methods.GetBlockAtTopoheight, params)
	err = rpc.JsonFormatResponse(res, err, &block)
	return
}

func (w *WebSocket) GetBlocksAtHeight(params GetBlockAtTopoheightParams) (blocks []Block, err error) {
	res, err := w.WS.Call(w.Prefix+methods.GetBlocksAtHeight, params)
	err = rpc.JsonFormatResponse(res, err, &blocks)
	return
}

func (w *WebSocket) GetBlockByHash(params GetBlockByHashParams) (block Block, err error) {
	res, err := w.WS.Call(w.Prefix+methods.GetBlockByHash, params)
	err = rpc.JsonFormatResponse(res, err, &block)
	return
}

func (w *WebSocket) GetTopBlock(params GetTopBlockParams) (block Block, err error) {
	res, err := w.WS.Call(w.Prefix+methods.GetTopBlock, params)
	err = rpc.JsonFormatResponse(res, err, &block)
	return
}

func (w *WebSocket) GetNonce(addr string) (nonce GetNonceResult, err error) {
	params := map[string]string{"address": addr}
	res, err := w.WS.Call(w.Prefix+methods.GetNonce, params)
	err = rpc.JsonFormatResponse(res, err, &nonce)
	return
}

func (w *WebSocket) GetNonceAtTopoheight(params GetNonceAtTopoheightParams) (nonce VersionedNonce, err error) {
	res, err := w.WS.Call(w.Prefix+methods.GetNonceAtTopoheight, params)
	err = rpc.JsonFormatResponse(res, err, &nonce)
	return
}

func (w *WebSocket) HasNonce(addr string) (hasNonce bool, err error) {
	params := map[string]string{"address": addr}
	res, err := w.WS.Call(w.Prefix+methods.HasNonce, params)
	err = rpc.JsonFormatResponse(res, err, &hasNonce)
	return
}

func (w *WebSocket) GetBalance(params GetBalanceParams) (balance GetBalanceResult, err error) {
	res, err := w.WS.Call(w.Prefix+methods.GetBalance, params)
	err = rpc.JsonFormatResponse(res, err, &balance)
	return
}

func (w *WebSocket) HasBalance(params GetBalanceParams) (hasBalance bool, err error) {
	res, err := w.WS.Call(w.Prefix+methods.HasBalance, params)
	err = rpc.JsonFormatResponse(res, err, &hasBalance)
	return
}

func (w *WebSocket) GetBalanceAtTopoheight(params GetBalanceAtTopoheightParams) (balance VersionedBalance, err error) {
	res, err := w.WS.Call(w.Prefix+methods.GetBalanceAtTopoheight, params)
	err = rpc.JsonFormatResponse(res, err, &balance)
	return
}

func (w *WebSocket) GetAsset(assetId string) (asset AssetData, err error) {
	params := map[string]string{"asset": assetId}
	res, err := w.WS.Call(w.Prefix+methods.GetAsset, params)
	err = rpc.JsonFormatResponse(res, err, &asset)
	return
}

func (w *WebSocket) GetAssets(params GetAssetsParams) (assets []AssetData, err error) {
	res, err := w.WS.Call(w.Prefix+methods.GetAssets, params)
	err = rpc.JsonFormatResponse(res, err, &assets)
	return
}

func (w *WebSocket) CountAssets() (count uint64, err error) {
	res, err := w.WS.Call(w.Prefix+methods.CountAssets, nil)
	err = rpc.JsonFormatResponse(res, err, &count)
	return
}

func (w *WebSocket) CountTransactions() (count uint64, err error) {
	res, err := w.WS.Call(w.Prefix+methods.CountTransactions, nil)
	err = rpc.JsonFormatResponse(res, err, &count)
	return
}

func (w *WebSocket) CountAccounts() (count uint64, err error) {
	res, err := w.WS.Call(w.Prefix+methods.CountAccounts, nil)
	err = rpc.JsonFormatResponse(res, err, &count)
	return
}

func (w *WebSocket) GetTips() (tips []string, err error) {
	res, err := w.WS.Call(w.Prefix+methods.GetTips, nil)
	err = rpc.JsonFormatResponse(res, err, &tips)
	return
}

func (w *WebSocket) P2PStatus() (status P2PStatusResult, err error) {
	res, err := w.WS.Call(w.Prefix+methods.P2PStatus, nil)
	err = rpc.JsonFormatResponse(res, err, &status)
	return
}

func (w *WebSocket) GetDAGOrder(params GetTopoheightRangeParams) (hashes []string, err error) {
	res, err := w.WS.Call(w.Prefix+methods.GetDAGOrder, params)
	err = rpc.JsonFormatResponse(res, err, &hashes)
	return
}

func (w *WebSocket) SubmitBlock(params SubmitBlockParams) (result bool, err error) {
	res, err := w.WS.Call(w.Prefix+methods.SubmitBlock, params)
	err = rpc.JsonFormatResponse(res, err, &result)
	return
}

func (w *WebSocket) SubmitTransaction(hexData string) (result bool, err error) {
	params := map[string]string{"data": hexData}
	res, err := w.WS.Call(w.Prefix+methods.SubmitTransaction, params)
	err = rpc.JsonFormatResponse(res, err, &hexData)
	return
}

func (w *WebSocket) GetMempool() (txs []Transaction, err error) {
	res, err := w.WS.Call(w.Prefix+methods.GetMempool, nil)
	err = rpc.JsonFormatResponse(res, err, &txs)
	return
}

func (w *WebSocket) GetMempoolCache(params GetMempoolCacheParams) (result GetMempoolCacheResult, err error) {
	res, err := w.WS.Call(w.Prefix+methods.GetMempool, params)
	err = rpc.JsonFormatResponse(res, err, &result)
	return
}

func (w *WebSocket) GetTransaction(hash string) (tx Transaction, err error) {
	params := map[string]string{"hash": hash}
	res, err := w.WS.Call(w.Prefix+methods.GetTransaction, params)
	err = rpc.JsonFormatResponse(res, err, &tx)
	return
}

func (w *WebSocket) GetTransactions(params GetTransactionsParams) (txs []Transaction, err error) {
	res, err := w.WS.Call(w.Prefix+methods.GetTransactions, params)
	err = rpc.JsonFormatResponse(res, err, &txs)
	return
}

func (w *WebSocket) GetBlocksRangeByTopoheight(params GetTopoheightRangeParams) (blocks []Block, err error) {
	res, err := w.WS.Call(w.Prefix+methods.GetBlocksRangeByTopoheight, params)
	err = rpc.JsonFormatResponse(res, err, &blocks)
	return
}

func (w *WebSocket) GetBlocksRangeByHeight(params GetHeightRangeParams) (blocks []Block, err error) {
	res, err := w.WS.Call(w.Prefix+methods.GetBlocksRangeByHeight, params)
	err = rpc.JsonFormatResponse(res, err, &blocks)
	return
}

func (w *WebSocket) GetAccounts(params GetAccountsParams) (addresses []string, err error) {
	res, err := w.WS.Call(w.Prefix+methods.GetAccounts, params)
	err = rpc.JsonFormatResponse(res, err, &addresses)
	return
}

func (w *WebSocket) GetAccountHistory(addr string) (history []AccountHistory, err error) {
	params := map[string]string{"address": addr}
	res, err := w.WS.Call(w.Prefix+methods.GetAccountHistory, params)
	err = rpc.JsonFormatResponse(res, err, &history)
	return
}

func (w *WebSocket) GetAccountAssets(addr string) (assets []string, err error) {
	params := map[string]string{"address": addr}
	res, err := w.WS.Call(w.Prefix+methods.GetAccountAssets, params)
	err = rpc.JsonFormatResponse(res, err, &assets)
	return
}

func (w *WebSocket) GetPeers() (result GetPeersResult, err error) {
	res, err := w.WS.Call(w.Prefix+methods.GetPeers, nil)
	err = rpc.JsonFormatResponse(res, err, &result)
	return
}

func (w *WebSocket) GetDevFeeThresholds() (fees []Fee, err error) {
	res, err := w.WS.Call(w.Prefix+methods.GetDevFeeThresholds, nil)
	err = rpc.JsonFormatResponse(res, err, &fees)
	return
}

func (w *WebSocket) GetSizeOnDisk() (sizeOnDisk SizeOnDisk, err error) {
	res, err := w.WS.Call(w.Prefix+methods.GetSizeOnDisk, nil)
	err = rpc.JsonFormatResponse(res, err, &sizeOnDisk)
	return
}

func (w *WebSocket) IsTxExecutedInBlock(params IsTxExecutedInBlockParams) (executed bool, err error) {
	res, err := w.WS.Call(w.Prefix+methods.IsTxExecutedInBlock, params)
	err = rpc.JsonFormatResponse(res, err, &executed)
	return
}

func (w *WebSocket) GetAccountRegistrationTopoheight(addr string) (topoheight uint64, err error) {
	params := map[string]string{"address": addr}
	res, err := w.WS.Call(w.Prefix+methods.GetAccountRegistrationTopoheight, params)
	err = rpc.JsonFormatResponse(res, err, &topoheight)
	return
}

func (w *WebSocket) IsAccountRegistered(params IsAccountRegisteredParams) (exists bool, err error) {
	res, err := w.WS.Call(w.Prefix+methods.IsAccountRegistered, params)
	err = rpc.JsonFormatResponse(res, err, &exists)
	return
}

func (w *WebSocket) GetDifficulty() (result GetDifficultyResult, err error) {
	res, err := w.WS.Call(w.Prefix+methods.GetDifficulty, nil)
	err = rpc.JsonFormatResponse(res, err, &result)
	return
}

func (w *WebSocket) ValidateAddress(params ValidateAddressParams) (result ValidateAddressResult, err error) {
	res, err := w.WS.Call(w.Prefix+methods.ValidateAddress, params)
	err = rpc.JsonFormatResponse(res, err, &result)
	return
}

func (w *WebSocket) ExtractKeyFromAddress(params ExtractKeyFromAddressParams) (key interface{}, err error) {
	res, err := w.WS.Call(w.Prefix+methods.ExtractKeyFromAddress, params)
	err = rpc.JsonFormatResponse(res, err, &key)
	return
}

func (w *WebSocket) GetMinerWork(params GetMinerWorkParams) (result GetMinerWorkResult, err error) {
	res, err := w.WS.Call(w.Prefix+methods.GetMinerWork, params)
	err = rpc.JsonFormatResponse(res, err, &result)
	return
}

func (w *WebSocket) SplitAddress(params SplitAddressParams) (result SplitAddressResult, err error) {
	res, err := w.WS.Call(w.Prefix+methods.SplitAddress, params)
	err = rpc.JsonFormatResponse(res, err, &result)
	return
}

func (w *WebSocket) GetHardForks() (result []HardFork, err error) {
	res, err := w.WS.Call(w.Prefix+methods.GetHardForks, nil)
	err = rpc.JsonFormatResponse(res, err, &result)
	return
}

func (w *WebSocket) GetEstimatedFeeRates() (result []FeeRatesEstimated, err error) {
	res, err := w.WS.Call(w.Prefix+methods.GetEstimatedFeeRates, nil)
	err = rpc.JsonFormatResponse(res, err, &result)
	return
}

func (w *WebSocket) GetPrunedTopoheight() (result uint64, err error) {
	res, err := w.WS.Call(w.Prefix+methods.GetPrunedTopoheight, nil)
	err = rpc.JsonFormatResponse(res, err, &result)
	return
}

func (w *WebSocket) GetTransactionExecutor(params GetTransactionExecutorParams) (result GetTransactionExecutorResult, err error) {
	res, err := w.WS.Call(w.Prefix+methods.GetTransactionExecutor, params)
	err = rpc.JsonFormatResponse(res, err, &result)
	return
}

func (w *WebSocket) HasMultisigAtTopoheight(params HasMultisigAtTopoheightParams) (result bool, err error) {
	res, err := w.WS.Call(w.Prefix+methods.HasMultisigAtTopoheight, params)
	err = rpc.JsonFormatResponse(res, err, &result)
	return
}

func (w *WebSocket) GetMultisigAtTopoheight(params GetMultisigAtTopoheightParams) (result GetMultisigAtTopoHeightResult, err error) {
	res, err := w.WS.Call(w.Prefix+methods.GetMultisigAtTopoheight, params)
	err = rpc.JsonFormatResponse(res, err, &result)
	return
}

func (w *WebSocket) GetMultisig(params GetMultisigParams) (result GetMultisigResult, err error) {
	res, err := w.WS.Call(w.Prefix+methods.GetMultisig, params)
	err = rpc.JsonFormatResponse(res, err, &result)
	return
}

func (w *WebSocket) HasMultisig(params HasMultisigParams) (result bool, err error) {
	res, err := w.WS.Call(w.Prefix+methods.HasMultisig, params)
	err = rpc.JsonFormatResponse(res, err, &result)
	return
}

func (w *WebSocket) GetContractOutputs(params GetContractOutputsParams) (result []ContractOutput, err error) {
	res, err := w.WS.Call(w.Prefix+methods.GetContractOutputs, params)
	var outputs []interface{}
	err = rpc.JsonFormatResponse(res, err, &outputs)
	result = parseContractOutputs(outputs)
	return
}

func (w *WebSocket) GetContractModule(params GetContractModuleParams) (result GetContractModuleResult, err error) {
	res, err := w.WS.Call(w.Prefix+methods.GetContractModule, params)
	err = rpc.JsonFormatResponse(res, err, &result)
	return
}

func (w *WebSocket) GetContractData(params GetContractDataParams) (result interface{}, err error) {
	res, err := w.WS.Call(w.Prefix+methods.GetContractData, params)
	err = rpc.JsonFormatResponse(res, err, &result)
	return
}

func (w *WebSocket) GetContractDataAtTopoheight(params GetContractDataAtTopoheightParams) (result interface{}, err error) {
	res, err := w.WS.Call(w.Prefix+methods.GetContractDataAtTopoheight, params)
	err = rpc.JsonFormatResponse(res, err, &result)
	return
}

func (w *WebSocket) GetContractBalance(params GetContractBalanceParams) (result interface{}, err error) {
	res, err := w.WS.Call(w.Prefix+methods.GetContractBalance, params)
	err = rpc.JsonFormatResponse(res, err, &result)
	return
}

func (w *WebSocket) GetContractBalanceAtTopoheight(params GetContractBalanceAtTopoheightParams) (result interface{}, err error) {
	res, err := w.WS.Call(w.Prefix+methods.GetContractBalanceAtTopoheight, params)
	err = rpc.JsonFormatResponse(res, err, &result)
	return
}

func (w *WebSocket) CountContracts() (result uint64, err error) {
	res, err := w.WS.Call(w.Prefix+methods.CountContracts, nil)
	err = rpc.JsonFormatResponse(res, err, &result)
	return
}

func (w *WebSocket) MakeIntegratedAddress(params MakeIntegratedAddressParams) (result string, err error) {
	res, err := w.WS.Call(w.Prefix+methods.MakeIntegratedAddress, params)
	err = rpc.JsonFormatResponse(res, err, &result)
	return
}

func (w *WebSocket) DecryptExtraData(params DecryptExtraDataParams) (result interface{}, err error) {
	res, err := w.WS.Call(w.Prefix+methods.DecryptExtraData, params)
	err = rpc.JsonFormatResponse(res, err, &result)
	return
}
