package wallet

import (
	"context"
	"testing"

	"github.com/xelis-project/xelis-go-sdk/config"
	"github.com/xelis-project/xelis-go-sdk/daemon"
)

const TESTING_ADDR = "xet:qf5u2p46jpgqmypqc2xwtq25yek2t7qhnqtdhw5kpfwcrlavs5asq0r83r7"
const MAINNET_ADDR = "xel:as3mgjlevw5ve6k70evzz8lwmsa5p0lgws2d60fulxylnmeqrp9qqukwdfg"

// cargo run --bin xelis_wallet -- --wallet-path ./wallets/test --password test --network dev --rpc-password test --rpc-bind-address 127.0.0.1:8081 --rpc-username test

func prepareRPC(t *testing.T) (wallet *RPC, ctx context.Context) {
	ctx = context.Background()
	wallet, err := NewRPC(ctx, config.LOCAL_WALLET_RPC, "test", "test")
	if err != nil {
		t.Fatal(err)
	}

	return
}

func TestRPCGetVersion(t *testing.T) {
	wallet, _ := prepareRPC(t)

	version, err := wallet.GetVersion()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", version)
}

func TestRPCGetNetwork(t *testing.T) {
	wallet, _ := prepareRPC(t)

	network, err := wallet.GetNetwork()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", network)
}

func TestRPCGetNonce(t *testing.T) {
	wallet, _ := prepareRPC(t)

	nonce, err := wallet.GetNonce()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", nonce)
}

func TestRPCGetTopoheight(t *testing.T) {
	wallet, _ := prepareRPC(t)

	topo, err := wallet.GetTopoheight()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", topo)
}

func TestRPCGetAddress(t *testing.T) {
	wallet, _ := prepareRPC(t)

	address, err := wallet.GetAddress(GetAddressParams{})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", address)
}

func TestRPCIsOnline(t *testing.T) {
	wallet, _ := prepareRPC(t)

	isOnline, err := wallet.IsOnline()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", isOnline)
}

func TestRPCIntegratedAddress(t *testing.T) {
	wallet, _ := prepareRPC(t)

	var integratedData interface{} = map[string]interface{}{"hello": "world"}

	integratedAddress, err := wallet.GetAddress(GetAddressParams{IntegratedData: &integratedData})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", integratedAddress)

	split, err := wallet.SplitAddress(SplitAddressParams{
		Address: integratedAddress,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", split)
}

func TestRPCRescan(t *testing.T) {
	wallet, _ := prepareRPC(t)
	_, err := wallet.Rescan(RescanParams{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestRPCSignData(t *testing.T) {
	wallet, _ := prepareRPC(t)
	data, err := wallet.SignData(map[string]interface{}{"hello": "world"})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", data)
}

func TestRPCBalanceAndAsset(t *testing.T) {
	wallet, _ := prepareRPC(t)

	balance, err := wallet.GetBalance(GetBalanceParams{
		Asset: config.XELIS_ASSET,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", balance)

	hasBalance, err := wallet.HasBalance(GetBalanceParams{
		Asset: config.XELIS_ASSET,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", hasBalance)

	precision, err := wallet.GetAssetPrecision(GetAssetPrecisionParams{
		Asset: config.XELIS_ASSET,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", precision)

	assets, err := wallet.GetTrackedAssets()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", assets)
}

func TestRPCGetTransaction(t *testing.T) {
	wallet, _ := prepareRPC(t)

	// Send: 50ebdb059e5c9ad0f9fc7ac5d970b17ec2fc81bf197c9e737cf2d3ca14c5ae84
	tx, err := wallet.GetTransaction(GetTransactionParams{
		Hash: "50ebdb059e5c9ad0f9fc7ac5d970b17ec2fc81bf197c9e737cf2d3ca14c5ae84",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", tx)

	// Burn: 8383b7027694615e790bea812a3385af8f140b55734b8eb89bf8a42d0671aec7
	tx, err = wallet.GetTransaction(GetTransactionParams{
		Hash: "8383b7027694615e790bea812a3385af8f140b55734b8eb89bf8a42d0671aec7",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", tx)

	// Receive: 37ecec82d39ea38d94240335d9fc1de01d039d52c764709f37766d36e3f5c336
	tx, err = wallet.GetTransaction(GetTransactionParams{
		Hash: "37ecec82d39ea38d94240335d9fc1de01d039d52c764709f37766d36e3f5c336",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", tx)

	// Coinbase: 000000001fd2bf51d9c895bc200bd3e17597edd9827ac616d66884b75b55ddab
	tx, err = wallet.GetTransaction(GetTransactionParams{
		Hash: "000000001fd2bf51d9c895bc200bd3e17597edd9827ac616d66884b75b55ddab",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", tx)
}

func TestRPCGetTransactionWithExtraData(t *testing.T) {
	wallet, _ := prepareRPC(t)

	tx, err := wallet.GetTransaction(GetTransactionParams{
		Hash: "5459a2567c7666d902fa5042db601d50b8353cd73927d6b5c3ad4f99a1368206",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", tx)
}

func TestRPCListTransactions(t *testing.T) {
	wallet, _ := prepareRPC(t)

	txs, err := wallet.ListTransactions(ListTransactionsParams{
		AcceptOutgoing: true,
		AcceptIncoming: true,
		AcceptCoinbase: false,
		AcceptBurn:     false,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", txs)
}

func TestRPCBurn(t *testing.T) {
	wallet, _ := prepareRPC(t)

	result, err := wallet.BuildTransaction(BuildTransactionParams{
		Burn: &daemon.Burn{
			Asset:  config.XELIS_ASSET,
			Amount: 1,
		},
		Broadcast: false,
		TxAsHex:   true,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", result)
}

func TestRPCTransfer(t *testing.T) {
	wallet, _ := prepareRPC(t)

	result, err := wallet.BuildTransaction(BuildTransactionParams{
		Transfers: []TransferBuilder{
			{Amount: 1, Asset: config.XELIS_ASSET, Destination: MAINNET_ADDR},
		},
		Broadcast: false,
		TxAsHex:   true,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", result)
}

func TestRPCSendExtraData(t *testing.T) {
	wallet, _ := prepareRPC(t)

	var extraData interface{} = map[string]interface{}{"hello": "world"}

	result, err := wallet.BuildTransaction(BuildTransactionParams{
		Transfers: []TransferBuilder{
			{Amount: 0, Asset: config.XELIS_ASSET, Destination: TESTING_ADDR, ExtraData: &extraData},
		},
		Broadcast: false,
		TxAsHex:   true,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", result)

	first_transfer := (*result.Data.Transfers)[0]

	result2, err := wallet.DecryptExtraData(DecryptExtraDataParams{
		ExtraData: *first_transfer.ExtraData,
		Role:      TxSenderRole,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", result2)
}

func TestRPCSendWithFeeBuilder(t *testing.T) {
	wallet, _ := prepareRPC(t)

	// you can only have one of both
	// either use Multiplier or Value
	//feeMultiplier := float64(1)
	feeValue := uint64(1)

	result, err := wallet.BuildTransaction(BuildTransactionParams{
		Transfers: []TransferBuilder{
			{Amount: 0, Asset: config.XELIS_ASSET, Destination: MAINNET_ADDR},
		},
		Fee: &FeeBuilder{
			//Multiplier: &feeMultiplier,
			Value: &feeValue,
		},
		Broadcast: false,
		TxAsHex:   true,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", result)
}

func TestRPCDeploySC(t *testing.T) {
	wallet, _ := prepareRPC(t)

	hex_program := "00000000000000010000000e0200000201000100000101001a1000010000"

	result, err := wallet.BuildTransaction(BuildTransactionParams{
		DeployContract: &hex_program,
		Broadcast:      true,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", result.Hash)
}

func TestRPCInvokeSC(t *testing.T) {
	wallet, _ := prepareRPC(t)

	result, err := wallet.BuildTransaction(BuildTransactionParams{
		InvokeContract: &InvokeContractBuilder{
			Contract: "dfed8218ba12cc5e155d3bbbbcff8a2060c2bf0eea0a52e7a33a8a81336b84ab",
			MaxGas:   50,
			ChunkId:  0,
			Parameters: []Constant{
				ConstantDefaultU64(1),
				ConstantDefault(ConstantValueU64, 2),
			},
			Deposits: map[string]ContractDepositBuilder{
				config.XELIS_ASSET: {Amount: 100, Private: false},
			},
		},
		Broadcast: true,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", result.Hash)
}

func TestRPCEstimateFees(t *testing.T) {
	wallet, _ := prepareRPC(t)

	result, err := wallet.EstimateFees(EstimateFeesParams{
		Transfers: []TransferBuilder{
			{Amount: 100, Asset: config.XELIS_ASSET, Destination: TESTING_ADDR},
			{Amount: 200, Asset: config.XELIS_ASSET, Destination: TESTING_ADDR},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", result)
}

func TestRPCNetworkInfo(t *testing.T) {
	wallet, _ := prepareRPC(t)

	result, err := wallet.NetworkInfo()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", result)
}

func TestRPCClearTxCache(t *testing.T) {
	wallet, _ := prepareRPC(t)

	result, err := wallet.ClearTxCache()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", result)
}

func TestRPCEstimateExtraDataSize(t *testing.T) {
	wallet, _ := prepareRPC(t)

	var integratedData interface{} = map[string]interface{}{"hello": "world"}

	addr, err := wallet.GetAddress(GetAddressParams{IntegratedData: &integratedData})
	if err != nil {
		t.Fatal(err)
	}

	result, err := wallet.EstimateExtraDataSize(EstimateExtraDataSizeParams{
		Destinations: []string{addr},
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", result)
}
