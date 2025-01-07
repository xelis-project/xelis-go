package daemon

import (
	"context"
	"testing"

	"github.com/xelis-project/xelis-go-sdk/config"
)

const WALLET_ADDR = "xet:62wnkswt0rmrdd9d2lawgpzuh87fkpmp4gx9j3g4u24yrdkdxgksqnuuucf"

func prepareRPC(t *testing.T) (daemon *RPC, ctx context.Context) {
	ctx = context.Background()
	daemon, err := NewRPC(ctx, config.LOCAL_NODE_RPC)
	if err != nil {
		t.Fatal(err)
	}

	return
}

func TestGetVersion(t *testing.T) {
	daemon, _ := prepareRPC(t)

	version, err := daemon.GetVersion()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", version)
}

func TestGetHeight(t *testing.T) {
	daemon, _ := prepareRPC(t)

	height, err := daemon.GetHeight()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", height)
}

func TestGetTopoheight(t *testing.T) {
	daemon, _ := prepareRPC(t)

	topoheight, err := daemon.GetTopoheight()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", topoheight)
}

func TestGetStableHeight(t *testing.T) {
	daemon, _ := prepareRPC(t)

	stableheight, err := daemon.GetStableHeight()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", stableheight)
}

func TestGetStableTopoheight(t *testing.T) {
	daemon, _ := prepareRPC(t)

	stabletopoheight, err := daemon.GetStableTopoheight()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", stabletopoheight)
}

func TestGetBlockTemplate(t *testing.T) {
	daemon, _ := prepareRPC(t)

	template, err := daemon.GetBlockTemplate(WALLET_ADDR)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", template)
}

func TestGetBlockAtTopoheight(t *testing.T) {
	daemon, _ := prepareRPC(t)

	genesisBlock, err := daemon.GetBlockAtTopoheight(GetBlockAtTopoheightParams{Topoheight: 0})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", genesisBlock)
}

func TestGetBlocksAtHeight(t *testing.T) {
	daemon, _ := prepareRPC(t)

	blocks, err := daemon.GetBlocksAtHeight(GetBlocksAtHeightParams{Height: 0})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", blocks)
}

func TestGetBlockByHash(t *testing.T) {
	daemon, _ := prepareRPC(t)

	block, err := daemon.GetBlockByHash(GetBlockByHashParams{Hash: `23827b240a9e6aeb0e7164a4e402838ffc383efdc92789d705921fccfed516b5`})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", block)
}

func TestGetTopBlock(t *testing.T) {
	daemon, _ := prepareRPC(t)

	topBlock, err := daemon.GetTopBlock(GetTopBlockParams{})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", topBlock)
}

func TestGetInfo(t *testing.T) {
	daemon, _ := prepareRPC(t)

	info, err := daemon.GetInfo()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", info)
}

func TestGetAsset(t *testing.T) {
	daemon, _ := prepareRPC(t)

	asset, err := daemon.GetAsset(config.XELIS_ASSET)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", asset)
}

func TestGetAssets(t *testing.T) {
	daemon, _ := prepareRPC(t)

	assets, err := daemon.GetAssets(GetAssetsParams{})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", assets)
}

func TestCountAssets(t *testing.T) {
	daemon, _ := prepareRPC(t)

	countAssets, err := daemon.CountAssets()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", countAssets)
}

func TestCountAccounts(t *testing.T) {
	daemon, _ := prepareRPC(t)

	countAccounts, err := daemon.CountAccounts()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", countAccounts)
}

func TestCountTransactions(t *testing.T) {
	daemon, _ := prepareRPC(t)

	countTransactions, err := daemon.CountTransactions()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", countTransactions)
}

func TestCountContracts(t *testing.T) {
	daemon, _ := prepareRPC(t)

	countContracts, err := daemon.CountContracts()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", countContracts)
}

func TestP2PStatus(t *testing.T) {
	daemon, _ := prepareRPC(t)

	status, err := daemon.P2PStatus()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", status)
}

func TestGetPeers(t *testing.T) {
	daemon, _ := prepareRPC(t)

	result, err := daemon.GetPeers()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(result)
}

func TestGetMempool(t *testing.T) {
	daemon, _ := prepareRPC(t)

	mempool, err := daemon.GetMempool()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", mempool)
}

func TestGetMempoolCache(t *testing.T) {
	daemon, _ := prepareRPC(t)

	cache, err := daemon.GetMempoolCache(GetMempoolCacheParams{
		Address: WALLET_ADDR,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", cache)
}

func TestGetTips(t *testing.T) {
	daemon, _ := prepareRPC(t)

	tips, err := daemon.GetTips()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", tips)
}

func TestGetDAGOrder(t *testing.T) {
	daemon, _ := prepareRPC(t)

	dagOrder, err := daemon.GetDAGOrder(GetTopoheightRangeParams{})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", dagOrder)
}

func TestGetAccounts(t *testing.T) {
	daemon, _ := prepareRPC(t)

	accounts, err := daemon.GetAccounts(GetAccountsParams{Maximum: 5})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", accounts)
}

func TestGetDevFeeThresholds(t *testing.T) {
	daemon, _ := prepareRPC(t)

	fees, err := daemon.GetDevFeeThresholds()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", fees)
}

func TestGetSizeOnDisk(t *testing.T) {
	daemon, _ := prepareRPC(t)

	size, err := daemon.GetSizeOnDisk()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", size)
}

func TestGetDifficulty(t *testing.T) {
	daemon, _ := prepareRPC(t)

	diff, err := daemon.GetDifficulty()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", diff)
}

func TestValidateAddress(t *testing.T) {
	daemon, _ := prepareRPC(t)

	validAddr, err := daemon.ValidateAddress(ValidateAddressParams{
		Address:         WALLET_ADDR,
		AllowIntegrated: false,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", validAddr)

	bytePublicKey, err := daemon.ExtractKeyFromAddress(ExtractKeyFromAddressParams{
		Address: WALLET_ADDR,
		AsHex:   false,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", bytePublicKey)

	hexPublicKey, err := daemon.ExtractKeyFromAddress(ExtractKeyFromAddressParams{
		Address: WALLET_ADDR,
		AsHex:   true,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", hexPublicKey)
}

func TestMakeIntegratedAddress(t *testing.T) {
	daemon, _ := prepareRPC(t)

	var integratedData interface{} = map[string]interface{}{"hello": "world"}

	addr, err := daemon.MakeIntegratedAddress(MakeIntegratedAddressParams{
		Address:        WALLET_ADDR,
		IntegratedData: integratedData,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", addr)
}

func TestRPCUnknownMethod(t *testing.T) {
	daemon, ctx := prepareRPC(t)
	res, err := daemon.Client.Call(ctx, "UnknownMethod", nil)
	if err == nil {
		t.Fatal("Expected an error")
	}

	t.Log(res)
}

func TestRPCNonceAndBalance(t *testing.T) {
	daemon, _ := prepareRPC(t)
	has, err := daemon.HasNonce(WALLET_ADDR)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(has)

	has, err = daemon.HasBalance(GetBalanceParams{
		Address: WALLET_ADDR,
		Asset:   config.XELIS_ASSET,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(has)

	nonce, err := daemon.GetNonce(WALLET_ADDR)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(nonce)

	versionedNonce, err := daemon.GetNonceAtTopoheight(GetNonceAtTopoheightParams{
		Address:    WALLET_ADDR,
		Topoheight: nonce.Topoheight,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(versionedNonce)

	balance, err := daemon.GetBalance(GetBalanceParams{
		Address: WALLET_ADDR,
		Asset:   config.XELIS_ASSET,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(balance)

	stableBalance, err := daemon.GetStableBalance(GetBalanceParams{
		Address: WALLET_ADDR,
		Asset:   config.XELIS_ASSET,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(stableBalance)

	versionedBalance, err := daemon.GetBalanceAtTopoheight(GetBalanceAtTopoheightParams{
		Address:    WALLET_ADDR,
		Asset:      config.XELIS_ASSET,
		Topoheight: nonce.Topoheight, // the testing addr does not have a balance before 322
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(versionedBalance)
}

func TestRPCGetBlocksRange(t *testing.T) {
	daemon, _ := prepareRPC(t)

	topoheight, err := daemon.GetTopoheight()
	if err != nil {
		t.Fatal(err)
	}

	blocks, err := daemon.GetBlocksRangeByTopoheight(GetTopoheightRangeParams{
		StartTopoheight: topoheight - 10,
		EndTopoheight:   topoheight,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(blocks)

	height, err := daemon.GetHeight()
	if err != nil {
		t.Fatal(err)
	}

	blocks, err = daemon.GetBlocksRangeByHeight(GetHeightRangeParams{
		StartHeight: height - 10,
		EndHeight:   height,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(blocks)
}

func TestRPCGetTransactions(t *testing.T) {
	daemon, _ := prepareRPC(t)
	txHash := "d9a6810d667c212e499ceb2acf60a8fbc0096da66b1e7a59fb3ae5d412ad58f2"

	txs, err := daemon.GetTransactions(GetTransactionsParams{
		TxHashes: []string{txHash},
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(txs)

	tx, err := daemon.GetTransaction(txHash)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(tx)
}

func TestRPCGetTransaction(t *testing.T) {
	daemon, _ := prepareRPC(t)
	txHash := "5f5e2ff1677860ee1f3e3c58ba188f427fbcb2f344dfb15dd0f7ca60b03f624c"

	tx, err := daemon.GetTransaction(txHash)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(tx)
}

func TestRPCExecutedInBlock(t *testing.T) {
	// https://testnet-explorer.xelis.io/blocks/000000001849d07bbb4165c8ba1d1fc472a0629f56895efb8689e06ce62b3ca8
	daemon, _ := prepareRPC(t)
	executed, err := daemon.IsTxExecutedInBlock(IsTxExecutedInBlockParams{
		TxHash:    "6e4bbd77b305fb68e2cc7576b4846d2db3617e3cbc2eb851cb2ae69b879e9d0f",
		BlockHash: "000000001849d07bbb4165c8ba1d1fc472a0629f56895efb8689e06ce62b3ca8",
	})
	if err != nil {
		t.Fatal(err)
	}

	if !executed {
		t.Errorf("tx should be executed in block")
	}

	t.Log(executed)
}

func TestRPCAccount(t *testing.T) {
	daemon, _ := prepareRPC(t)
	history, err := daemon.GetAccountHistory(WALLET_ADDR)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(history)

	assets, err := daemon.GetAccountAssets(WALLET_ADDR)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(assets)

	topoheight, err := daemon.GetAccountRegistrationTopoheight(WALLET_ADDR)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(topoheight)
}

func TestRPCRegistration(t *testing.T) {
	// using mainnet for this test
	// we need to resync the blockchain to work on testnet
	daemon, _ := prepareRPC(t)

	topoheight, err := daemon.GetAccountRegistrationTopoheight(WALLET_ADDR)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(topoheight)

	exists, err := daemon.IsAccountRegistered(IsAccountRegisteredParams{
		Address:        WALLET_ADDR,
		InStableHeight: true,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(exists)
}

func TestGetMinerWork(t *testing.T) {
	daemon, _ := prepareRPC(t)

	var addr = "xet:w64wu066sq7jq4v9f37a5gy8hgyvc2gvt237u2457mme2m2r7avqqtmufz3"

	blockTemplate, err := daemon.GetBlockTemplate(addr)
	if err != nil {
		t.Fatal(err)
	}

	result, err := daemon.GetMinerWork(GetMinerWorkParams{
		Template: blockTemplate.Template,
		Address:  &addr,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(result)
}

func TestSplitAddress(t *testing.T) {
	daemon, _ := prepareRPC(t)

	result, err := daemon.SplitAddress(SplitAddressParams{Address: "xet:upqflhm65lmjtukavf4de93kphk4j990hw9x9hhrc8rwleduruhqzqqpqvcnydgd3plda"})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(result)
}

func TestGetEstimatedFeeRates(t *testing.T) {
	daemon, _ := prepareRPC(t)

	result, err := daemon.GetEstimatedFeeRates()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(result)
}

func TestGetPrunedTopoheight(t *testing.T) {
	daemon, _ := prepareRPC(t)

	result, err := daemon.GetPrunedTopoheight()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(result)
}

func TestGetHardForks(t *testing.T) {
	daemon, _ := prepareRPC(t)

	result, err := daemon.GetHardForks()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(result)
}

func TestGetTransactionExecutor(t *testing.T) {
	daemon, _ := prepareRPC(t)

	result, err := daemon.GetTransactionExecutor(GetTransactionExecutorParams{
		Hash: "",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(result)
}

func TestHasMultisigAtTopoheight(t *testing.T) {
	daemon, _ := prepareRPC(t)

	result, err := daemon.HasMultisigAtTopoheight(HasMultisigAtTopoheightParams{
		Address:    WALLET_ADDR,
		Topoheight: 0,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(result)
}

func TestGetMultisigAtTopoheight(t *testing.T) {
	daemon, _ := prepareRPC(t)

	result, err := daemon.GetMultisigAtTopoheight(GetMultisigAtTopoheightParams{
		Address:    WALLET_ADDR,
		Topoheight: 0,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(result)
}

func TestGetMultisig(t *testing.T) {
	daemon, _ := prepareRPC(t)

	result, err := daemon.GetMultisig(GetMultisigParams{
		Address: WALLET_ADDR,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(result)
}

func TestHasMultisig(t *testing.T) {
	daemon, _ := prepareRPC(t)

	result, err := daemon.HasMultisig(HasMultisigParams{
		Address: WALLET_ADDR,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(result)
}

func TestGetContractOutputs(t *testing.T) {
	daemon, _ := prepareRPC(t)

	result, err := daemon.GetContractOutputs(GetContractOutputsParams{
		Transaction: "3519b8c72356b8ea4af98b01c1c7e607ccf3770b6ef5bad88634d08a4a58e40f",
	})
	if err != nil {
		t.Fatal(err)
	}

	for _, output := range result {
		switch out := output.(type) {
		case ContractOutputExitCode:
			t.Log("Exit code", out.ExitCode)
		case ContractOutputRefundGas:
			t.Log("Refund gas", out.Amount)
		}
	}

	t.Log(result)
}

func TestGetContractModule(t *testing.T) {
	daemon, _ := prepareRPC(t)

	result, err := daemon.GetContractModule(GetContractModuleParams{
		Contract: "dfed8218ba12cc5e155d3bbbbcff8a2060c2bf0eea0a52e7a33a8a81336b84ab",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(result)
}

func TestGetContractData(t *testing.T) {
	daemon, _ := prepareRPC(t)

	result, err := daemon.GetContractData(GetContractDataParams{
		Contract: "dfed8218ba12cc5e155d3bbbbcff8a2060c2bf0eea0a52e7a33a8a81336b84ab",
		Key:      nil,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(result)
}

func TestGetContractDataAtTopoheight(t *testing.T) {
	daemon, _ := prepareRPC(t)

	topo := uint64(0)
	result, err := daemon.GetContractDataAtTopoheight(GetContractDataAtTopoheightParams{
		Contract:   "dfed8218ba12cc5e155d3bbbbcff8a2060c2bf0eea0a52e7a33a8a81336b84ab",
		Key:        nil,
		Topoheight: &topo,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(result)
}

func TestGetContractBalance(t *testing.T) {
	daemon, _ := prepareRPC(t)

	result, err := daemon.GetContractBalance(GetContractBalanceParams{
		Contract: "dfed8218ba12cc5e155d3bbbbcff8a2060c2bf0eea0a52e7a33a8a81336b84ab",
		Asset:    config.XELIS_ASSET,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(result)
}

func TestGetContractBalanceAtTopoheight(t *testing.T) {
	daemon, _ := prepareRPC(t)

	result, err := daemon.GetContractBalanceAtTopoheight(GetContractBalanceAtTopoheightParams{
		Contract:   "dfed8218ba12cc5e155d3bbbbcff8a2060c2bf0eea0a52e7a33a8a81336b84ab",
		Asset:      config.XELIS_ASSET,
		Topoheight: 0,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(result)
}
