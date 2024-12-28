package methods

const (
	GetVersion                  string = "get_version"
	GetNetwork                  string = "get_network"
	GetNonce                    string = "get_nonce"
	GetTopoheight               string = "get_topoheight"
	GetAddress                  string = "get_address"
	SplitAddress                string = "split_address"
	Rescan                      string = "rescan"
	GetBalance                  string = "get_balance"
	HasBalance                  string = "has_balance"
	GetTrackedAssets            string = "get_tracked_assets"
	GetAssetPrecision           string = "get_asset_precision"
	GetAssets                   string = "get_assets"
	GetAsset                    string = "get_asset"
	GetTransaction              string = "get_transaction"
	BuildTransaction            string = "build_transaction"
	BuildTransactionOffline     string = "build_transaction_offline"
	BuildUnsignedTransaction    string = "build_unsigned_transaction"
	SignUnsignedTransaction     string = "sign_unsigned_transaction"
	FinalizeUnsignedTransaction string = "finalize_unsigned_transaction"
	ClearTxCache                string = "clear_tx_cache"
	ListTransactions            string = "list_transactions"
	IsOnline                    string = "is_online"
	SetOnlineMode               string = "set_online_mode"
	SetOfflineMode              string = "set_offline_mode"
	SignData                    string = "sign_data"
	EstimateFees                string = "estimate_fees"
	EstimateExtraDataSize       string = "estimate_extra_data_size"
	NetworkInfo                 string = "network_info"
	DecryptExtraData            string = "decrypt_extra_data"
	DecryptCiphertext           string = "decrypt_ciphertext"

	// Interact with wallet encrypted database
	GetMatchingKeys      string = "get_matching_keys"
	CountMatchingEntries string = "count_matching_entries"
	GetValueFromKey      string = "get_value_from_key"
	Store                string = "store"
	Delete               string = "delete"
	DeleteTreeEntries    string = "delete_tree_entries"
	HasKey               string = "has_key"
	QueryDB              string = "query_db"
)
