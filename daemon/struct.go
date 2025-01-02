package daemon

type GetTopoheightRangeParams struct {
	StartTopoheight uint64 `json:"start_topoheight"`
	EndTopoheight   uint64 `json:"end_topoheight"`
}

type GetBlockAtTopoheightParams struct {
	Topoheight uint64 `json:"topoheight"`
	IncludeTxs bool   `json:"include_txs"`
}

type GetBlocksAtHeightParams struct {
	Height     uint64 `json:"height"`
	IncludeTxs bool   `json:"include_txs"`
}

type GetBlockByHashParams struct {
	Hash       string `json:"hash"`
	IncludeTxs bool   `json:"include_txs"`
}

type GetTopBlockParams struct {
	IncludeTxs bool `json:"include_txs"`
}

type GetBalanceParams struct {
	Address string `json:"address"`
	Asset   string `json:"asset"`
}

type BalanceType string

var (
	BalanceInput  BalanceType = `input`
	BalanceOutput BalanceType = `output`
	BalanceBoth   BalanceType = `both`
)

type EncryptedBalance struct {
	Commitment []byte `json:"commitment"`
	Handle     []byte `json:"handle"`
}

type VersionedBalance struct {
	BalanceType        BalanceType       `json:"balance_type"`
	FinalBalance       EncryptedBalance  `json:"final_balance"`
	OutputBalance      *EncryptedBalance `json:"output_balance"`
	PreviousTopoheight *uint64           `json:"previous_topoheight"`
}

type GetBalanceResult struct {
	Version    VersionedBalance `json:"version"`
	Topoheight uint64           `json:"topoheight"`
}

type GetStableBalanceResult struct {
	StableTopoheight uint64           `json:"stable_topoheight"`
	StableBlockHash  string           `json:"stable_block_hash"`
	Version          VersionedBalance `json:"version"`
}

type GetNonceAtTopoheightParams struct {
	Address    string `json:"address"`
	Topoheight uint64 `json:"topoheight"`
}

type GetBalanceAtTopoheightParams struct {
	Address    string `json:"address"`
	Asset      string `json:"asset"`
	Topoheight uint64 `json:"topoheight"`
}

type GetHeightRangeParams struct {
	StartHeight uint64 `json:"start_height"`
	EndHeight   uint64 `json:"end_height"`
}

type GetTransactionsParams struct {
	TxHashes []string `json:"tx_hashes"`
}

type P2PStatusResult struct {
	BestTopoheight uint64 `json:"best_topoheight"`
	MaxPeers       uint64 `json:"max_peers"`
	OurTopoheight  uint64 `json:"our_topoheight"`
	PeerCount      uint64 `json:"peer_count"`
	PeerId         uint64 `json:"peer_id"`
	Tag            string `json:"tag"`
}

type GetAssetsParams = GetAccountsParams

type GetAccountsParams struct {
	Skip              uint64 `json:"skip,omitempty"`
	Maximum           uint64 `json:"maximum,omitempty"`
	MinimumTopoheight uint64 `json:"minimum_topoheight,omitempty"`
	MaximumTopoheight uint64 `json:"maximum_topoheight,omitempty"`
}

type Block struct {
	BlockType            string   `json:"block_type"`
	CumulativeDifficulty string   `json:"cumulative_difficulty"`
	Difficulty           string   `json:"difficulty"`
	ExtraNonce           string   `json:"extra_nonce"`
	Hash                 string   `json:"hash"`
	Height               uint64   `json:"height"`
	Miner                string   `json:"miner"`
	Nonce                uint64   `json:"nonce"`
	Reward               *uint64  `json:"reward"` // full reward miner_reward + dev_reward
	MinerReward          *uint64  `json:"miner_reward"`
	DevReward            *uint64  `json:"dev_reward"`
	Supply               *uint64  `json:"supply"`
	Timestamp            uint64   `json:"timestamp"`
	Tips                 []string `json:"tips"`
	Topoheight           *uint64  `json:"topoheight"`
	TotalFees            *uint64  `json:"total_fees"`
	TotalSizeInBytes     uint64   `json:"total_size_in_bytes"`
	TxsHashes            []string `json:"txs_hashes"`
	Version              uint64   `json:"version"`
}

type Transfer struct {
	Asset           string  `json:"asset"`
	ExtraData       *[]byte `json:"extra_data"`
	Destination     string  `json:"destination"`
	Commitment      []byte  `json:"commitment"`
	SenderHandle    []byte  `json:"sender_handle"`
	ReceiverHandle  []byte  `json:"receiver_handle"`
	CTValidityProof Proof   `json:"ct_validity_proof"`
}

type Burn struct {
	Asset  string `json:"asset"`
	Amount uint64 `json:"amount"`
}

type CallContract struct {
	Contract string `json:"contract"`
}

type TransactionData struct {
	Transfers []Transfer `json:"transfers"`
	Burn      *Burn      `json:"burn"`
	// CallContract   string     `json:"call_contract"`
	// DeployContract string     `json:"deploy_contract"`
}

type Reference struct {
	Hash       string `json:"hash"`
	Topoheight uint64 `json:"topoheight"`
}

type Proof struct {
	Y_0 []byte `json:"Y_0"`
	Y_1 []byte `json:"Y_1"`
	Z_R []byte `json:"z_r"`
	Z_X []byte `json:"z_x"`
}

type EqProof struct {
	Y_0 []byte `json:"Y_0"`
	Y_1 []byte `json:"Y_1"`
	Y_2 []byte `json:"Y_2"`
	Z_R []byte `json:"z_r"`
	Z_S []byte `json:"z_s"`
	Z_X []byte `json:"z_x"`
}

type SourceCommitment struct {
	Commitment []byte  `json:"commitment"`
	Proof      EqProof `json:"proof"`
	Asset      string  `json:"asset"`
}

type Transaction struct {
	Blocks            []string           `json:"blocks"`
	Hash              string             `json:"hash"`
	Data              TransactionData    `json:"data"`
	Fee               uint64             `json:"fee"`
	Nonce             uint64             `json:"nonce"`
	Source            string             `json:"source"`
	Reference         Reference          `json:"reference"`
	SourceCommitments []SourceCommitment `json:"source_commitments"`
	RangeProof        []byte             `json:"range_proof"`
	Signature         string             `json:"signature"`
	ExecutedInBlock   *string            `json:"executed_in_block"`
	Version           uint64             `json:"version"`
	FirstSeen         *uint64            `json:"first_seen"`
	InMempool         bool               `json:"in_mempool"`
	Size              uint64             `json:"size"`
}

type GetInfoResult struct {
	Height            uint64 `json:"height"`
	Topoheight        uint64 `json:"topoheight"`
	Stableheight      uint64 `json:"stableheight"`
	PrunedTopoheight  uint64 `json:"pruned_topoheight"`
	TopBlockHash      string `json:"top_block_hash"`
	CirculatingSupply uint64 `json:"circulating_supply"`
	BurnedSupply      uint64 `json:"burned_supply"`
	EmittedSupply     uint64 `json:"emitted_supply"`
	MaximumSupply     uint64 `json:"maximum_supply"`
	Difficulty        string `json:"difficulty"`
	BlockTimeTarget   uint64 `json:"block_time_target"`
	AverageBlockTime  uint64 `json:"average_block_time"`
	BlockReward       uint64 `json:"block_reward"`
	DevReward         uint64 `json:"dev_reward"`
	MinerReward       uint64 `json:"miner_reward"`
	MempoolSize       uint64 `json:"mempool_size"`
	Version           string `json:"version"`
	Network           string `json:"network"`
}

type GetBlockTemplateResult struct {
	Template   string `json:"template"`
	Algorithm  string `json:"algorithm"`
	Height     uint64 `json:"height"`
	Topoheight uint64 `json:"topoheight"`
	Difficulty string `json:"difficulty"`
}

type SubmitBlockParams struct {
	BlockTemplate string  `json:"block_template"`
	MinerWork     *string `json:"miner_work,omitempty"`
}

type GetMinerWorkParams struct {
	Template string  `json:"template"`
	Address  *string `json:"address,omitempty"`
}

type GetMinerWorkResult struct {
	MinerWork  string `json:"miner_work"`
	Algorithm  string `json:"algorithm"`
	Height     uint64 `json:"height"`
	Difficulty string `json:"difficulty"`
	Topoheight uint64 `json:"topoheight"`
}

type GetNonceResult struct {
	Nonce              uint64  `json:"nonce"`
	PreviousTopoheight *uint64 `json:"previous_topoheight"`
	Topoheight         uint64  `json:"topoheight"`
}

type VersionedNonce struct {
	Nonce              uint64  `json:"nonce"`
	PreviousTopoheight *uint64 `json:"previous_topoheight"`
}

type MiningHistory struct {
	Reward uint64 `json:"reward"`
}

type BurnHistory struct {
	Amount uint64 `json:"amount"`
}

type OutgoingHistory struct {
	To string `json:"to"`
}

type IncomingHistory struct {
	From string `json:"from"`
}

type Asset struct {
	Topoheight uint64 `json:"topoheight"`
	Decimals   int    `json:"decimals"`
}

type AssetWithData struct {
	Asset      string `json:"asset"`
	Topoheight uint64 `json:"topoheight"`
	Decimals   int    `json:"decimals"`
}

type Fee struct {
	FeePercentage int    `json:"fee_percentage"`
	Height        uint64 `json:"height"`
}

type SizeOnDisk struct {
	SizeBytes     uint64 `json:"size_bytes"`
	SizeFormatted string `json:"size_formatted"`
}

type IsTxExecutedInBlockParams struct {
	TxHash    string `json:"tx_hash"`
	BlockHash string `json:"block_hash"`
}

type AccountHistory struct {
	Topoheight     uint64           `json:"topoheight"`
	BlockTimestamp uint64           `json:"block_timestamp"`
	Hash           string           `json:"hash"`
	Mining         *MiningHistory   `json:"mining"`
	Burn           *BurnHistory     `json:"burn"`
	Outgoing       *OutgoingHistory `json:"outgoing"`
	Incoming       *IncomingHistory `json:"incoming"`
	DevFee         *MiningHistory   `json:"dev_fee"`
}

type TransactionExecutedResult struct {
	BlockHash  string `json:"block_hash"`
	Topoheight uint64 `json:"topoheight"`
	TxHash     string `json:"tx_hash"`
}

type PeerDirection string

const (
	PeerIn   PeerDirection = "In"
	PeerOut  PeerDirection = "Out"
	PeerBoth PeerDirection = "Both"
)

type Peer struct {
	Id                   uint64                   `json:"id"`
	CumulativeDifficulty string                   `json:"cumulative_difficulty"`
	PrunedTopoheight     uint64                   `json:"pruned_topoheight"`
	ConnectedOn          uint64                   `json:"connected_on"`
	Height               uint64                   `json:"height"`
	LocalPort            int                      `json:"local_port"`
	TopBlockHash         string                   `json:"top_block_hash"`
	Addr                 string                   `json:"addr"`
	LastPing             uint64                   `json:"last_ping"`
	Tag                  string                   `json:"tag"`
	Topoheight           uint64                   `json:"topoheight"`
	Peers                map[string]PeerDirection `json:"peers"`
	Version              string                   `json:"version"`
}

type GetPeersResult struct {
	Peers       []Peer `json:"peers"`
	TotalPeers  int    `json:"total_peers"`
	HiddenPeers int    `json:"hidden_peers"`
}

type IsAccountRegisteredParams struct {
	Address        string `json:"address"`
	InStableHeight bool   `json:"in_stable_height"`
}

type GetDifficultyResult struct {
	Difficulty        string `json:"difficulty"`
	Hashrate          string `json:"hashrate"`
	HashrateFormatted string `json:"hashrate_formatted"`
}

type ValidateAddressParams struct {
	Address         string `json:"address"`
	AllowIntegrated bool   `json:"allow_integrated"`
}

type ValidateAddressResult struct {
	IsIntegrated bool `json:"is_integrated"`
	IsValid      bool `json:"is_valid"`
}

type ExtractKeyFromAddressParams struct {
	Address string `json:"address"`
	AsHex   bool   `json:"as_hex"`
}

type SplitAddressParams struct {
	Address string `json:"address"`
}

type SplitAddressResult struct {
	Address        string      `json:"address"`
	IntegratedData interface{} `json:"integrated_data"`
}

type GetTransactionExecutorParams struct {
	Hash string `json:"hash"`
}

type GetTransactionExecutorResult struct {
	BlockTopoheight uint64 `json:"block_topoheight"`
	BlockHash       string `json:"block_hash"`
	BlockTimestamp  uint64 `json:"block_timestamp"`
}

type HasMultisigAtTopoheightParams struct {
	Address    string `json:"address"`
	Topoheight uint32 `json:"topoheight"`
}

type GetMultisigAtTopoheightParams struct {
	Address    string `json:"address"`
	Topoheight uint32 `json:"topoheight"`
}

type GetMultisigAtTopoHeightResult struct {
	State string `json:"state"`
}

type GetMultisigParams struct {
	Address string `json:"address"`
}

type GetMultisigResult struct {
	State      string `json:"state"`
	Topoheight uint64 `json:"topoheight"`
}

type HasMultisigParams struct {
	Address    string  `json:"address"`
	Topoheight *uint32 `json:"topoheight,omitempty"`
}

type GetContractOutputsParams struct {
	Transaction string `json:"transaction"`
}

type GetContractModuleParams struct {
	Contract string `json:"contract"`
}

type GetContractDataParams struct {
	Contract string      `json:"contract"`
	Key      interface{} `json:"key"`
}

type GetContractDataAtTopoheightParams struct {
	Contract   string      `json:"contract"`
	Key        interface{} `json:"key"`
	Topoheight *uint64     `json:"topoheight,omitempty"`
}

type GetContractBalanceParams struct {
	Contract string `json:"contract"`
	Asset    string `json:"asset"`
}

type GetContractBalanceAtTopoheightParams struct {
	Contract   string `json:"contract"`
	Asset      string `json:"asset"`
	Topoheight uint64 `json:"topoheight"`
}

type ContractOutputRefundGas struct {
	Amount uint64
}

type ContractOutputTransfer struct {
	Amount      uint64
	Asset       string
	Destination string
}

type ContractOutputExitCode struct {
	ExitCode uint64
}

type ContractOutputRefundDeposits struct{}

type ContractOutput interface{}

type Chunk struct {
	Instructions []byte `json:"instructions"`
}

type Module struct {
	Constants     []interface{} `json:"constants"`
	Chunks        []Chunk       `json:"chunks"`
	EntryChunkIds []uint64      `json:"entry_chunk_ids"`
	Structs       []interface{} `json:"structs"`
	Enums         []interface{} `json:"enums"`
}

type GetContractModuleResult struct {
	PreviousTopoheight *uint64 `json:"previous_topoheight"`
	Data               *Module `json:"data"`
}

type GetContractDataResult struct {
	PreviousTopoheight *uint64      `json:"previous_topoheight"`
	Data               *interface{} `json:"data"`
}

type HardFork struct {
	Height             uint64  `json:"height"`
	Version            uint8   `json:"version"`
	Changelog          string  `json:"changelog"`
	VersionRequirement *string `json:"version_requirement"`
}

type FeeRatesEstimated struct {
	Low     uint64 `json:"low"`
	Medium  uint64 `json:"medium"`
	High    uint64 `json:"high"`
	Default uint64 `json:"default"`
}

type MakeIntegratedAddressParams struct {
	Address        string      `json:"address"`
	IntegratedData interface{} `json:"integrated_data"`
}

type DecryptExtraDataParams struct {
	SharedKey []byte `json:"shared_key"`
	ExtraData []byte `json:"extra_data"`
}
