package methods

const (
	GetVersion          string = "get_version"
	GetHeight           string = "get_height"
	GetTopoHeight       string = "get_topoheight"
	GetPrunedTopoheight string = "get_pruned_topoheight"
	GetInfo             string = "get_info"
	GetDifficulty       string = "get_difficulty"
	GetTips             string = "get_tips"
	GetDevFeeThresholds string = "get_dev_fee_thresholds"
	GetSizeOnDisk       string = "get_size_on_disk"

	GetStableHeight     string = "get_stable_height"
	GetStableTopoheight string = "get_stable_topoheight"
	GetHardForks        string = "get_hard_forks"

	GetBlockAtTopoheight string = "get_block_at_topoheight"
	GetBlocksAtHeight    string = "get_blocks_at_height"
	GetBlockByHash       string = "get_block_by_hash"
	GetTopBlock          string = "get_top_block"

	GetBalance             string = "get_balance"
	GetStableBalance       string = "get_stable_balance"
	HasBalance             string = "has_balance"
	GetBalanceAtTopoheight string = "get_balance_at_topoheight"

	GetNonce             string = "get_nonce"
	HasNonce             string = "has_nonce"
	GetNonceAtTopoheight string = "get_nonce_at_topoheight"

	GetAsset  string = "get_asset"
	GetAssets string = "get_assets"

	CountAssets       string = "count_assets"
	CountTransactions string = "count_transactions"
	CountAccounts     string = "count_accounts"
	CountContracts    string = "count_contracts"

	SubmitTransaction      string = "submit_transaction"
	GetTransactionExecutor string = "get_transaction_executor"
	GetTransaction         string = "get_transaction"
	GetTransactions        string = "get_transactions"
	IsTxExecutedInBlock    string = "is_tx_executed_in_block"

	P2PStatus string = "p2p_status"
	GetPeers  string = "get_peers"

	GetMempool           string = "get_mempool"
	GetMempoolCache      string = "get_mempool_cache"
	GetEstimatedFeeRates string = "get_estimated_fee_rates"

	GetDAGOrder                string = "get_dag_order"
	GetBlocksRangeByTopoheight string = "get_blocks_range_by_topoheight"
	GetBlocksRangeByHeight     string = "get_blocks_range_by_height"

	GetAccountHistory                string = "get_account_history"
	GetAccountAssets                 string = "get_account_assets"
	GetAccounts                      string = "get_accounts"
	IsAccountRegistered              string = "is_account_registered"
	GetAccountRegistrationTopoheight string = "get_account_registration_topoheight"

	ValidateAddress       string = "validate_address"
	SplitAddress          string = "split_address"
	ExtractKeyFromAddress string = "extract_key_from_address"
	MakeIntegratedAddress string = "make_integrated_address"
	DecryptExtraData      string = "decrypt_extra_data"

	GetMultisigAtTopoheight string = "get_multisig_at_topoheight"
	GetMultisig             string = "get_multisig"
	HasMultisig             string = "has_multisig"
	HasMultisigAtTopoheight string = "has_multisig_at_topoheight"

	GetContractOutputs             string = "get_contract_outputs"
	GetContractModule              string = "get_contract_module"
	GetContractData                string = "get_contract_data"
	GetContractDataAtTopoheight    string = "get_contract_data_at_topoheight"
	GetContractBalance             string = "get_contract_balance"
	GetContractBalanceAtTopoheight string = "get_contract_balance_at_topoheight"

	GetBlockTemplate string = "get_block_template"
	GetMinerWork     string = "get_miner_work"
	SubmitBlock      string = "submit_block"
)
