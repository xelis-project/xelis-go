package wallet

import (
	"github.com/xelis-project/xelis-go-sdk/daemon"
)

type GetAddressParams struct {
	IntegratedData *interface{} `json:"integrated_data,omitempty"`
}

type SplitAddressParams struct {
	Address string `json:"address"`
}

type SplitAddressResult struct {
	Address        string      `json:"address"`
	IntegratedData interface{} `json:"integrated_data"`
}

type GetBalanceParams struct {
	Asset string `json:"asset"`
}

type GetTransactionParams struct {
	Hash string `json:"hash"`
}

type RescanParams struct {
	UntilTopoheight uint64 `json:"until_topoheight"`
}

type GetAssetPrecisionParams struct {
	Asset string `json:"asset"`
}

type TransferIn struct {
	Amount    uint64       `json:"amount"`
	Asset     string       `json:"asset"`
	ExtraData *interface{} `json:"extra_data"`
}

/*
// we use *interface{} instead of DataElement so user can serialize/deserialize how he wants
type DataElement struct {
	Value  interface{}     `json:"value,omitempty"`
	Array  []DataElement   `json:"array,omitempty"`
	Fields json.RawMessage `json:"fields,omitempty"` // can't do map[interface{}]DataElement json unsupported parsing
}
*/

type TransferOut struct {
	Amount      uint64       `json:"amount"`
	Asset       string       `json:"asset"`
	Destination string       `json:"destination"`
	ExtraData   *interface{} `json:"extra_data,omitempty"`
}

type FeeBuilder struct {
	Multiplier *float64 `json:"multiplier,omitempty"`
	Value      *uint64  `json:"value,omitempty"`
}

type MutliSigBuilder struct {
	Participants []string `json:"participants"`
	Threshold    uint8    `json:"threshold"`
}

type ContractDepositBuilder struct {
	Amount  uint64 `json:"amount"`
	Private bool   `json:"private"`
}

type InvokeContractBuilder struct {
	Contract   string                            `json:"contract"`
	MaxGas     uint64                            `json:"max_gas"`
	ChunkId    uint16                            `json:"chunk_id"`
	Parameters []interface{}                     `json:"parameters"`
	Deposits   map[string]ContractDepositBuilder `json:"deposits"`
}

type SignerId struct {
	Id         uint8  `json:"id"`
	PrivateKey []byte `json:"private_key"`
}

type BuildTransactionParams struct {
	Transfers      []TransferOut          `json:"transfers"`
	Burn           *daemon.Burn           `json:"burn,omitempty"`
	MultiSig       *MutliSigBuilder       `json:"multi_sig,omitempty"`
	InvokeContract *InvokeContractBuilder `json:"invoke_contract,omitempty"`
	DeployContract *string                `json:"deploy_contract,omitempty"`
	Fee            *FeeBuilder            `json:"fee,omitempty"`
	Nonce          *uint64                `json:"nonce,omitempty"`
	TxVersion      *uint8                 `json:"tx_version,omitempty"`
	Broadcast      bool                   `json:"broadcast"`
	TxAsHex        bool                   `json:"tx_as_hex"`
	Signers        *[]SignerId            `json:"signers,omitempty"`
}

// !!! not the same as daemon.Transfer
// the destination is []byte and the other it's string
type Transfer struct {
	Asset           string       `json:"asset"`
	ExtraData       *[]byte      `json:"extra_data"`
	Destination     []byte       `json:"destination"`
	Commitment      []byte       `json:"commitment"`
	SenderHandle    []byte       `json:"sender_handle"`
	ReceiverHandle  []byte       `json:"receiver_handle"`
	CTValidityProof daemon.Proof `json:"ct_validity_proof"`
}

type TransactionData struct {
	Transfers []Transfer   `json:"transfers"`
	Burn      *daemon.Burn `json:"burn"`
}

type TransactionResponse struct {
	Data              TransactionData           `json:"data"`
	Fee               uint64                    `json:"fee"`
	Hash              string                    `json:"hash"`
	Nonce             uint64                    `json:"nonce"`
	RangeProof        []byte                    `json:"range_proof"`
	Reference         daemon.Reference          `json:"reference"`
	Signature         string                    `json:"signature"`
	Source            []byte                    `json:"source"`
	SourceCommitments []daemon.SourceCommitment `json:"source_commitments"`
	TxAsHex           string                    `json:"tx_as_hex"`
	Version           uint64                    `json:"version"`
}

type Outgoing struct {
	Fee       uint64        `json:"fee"`
	Nonce     uint64        `json:"nonce"`
	Transfers []TransferOut `json:"transfers"`
}

type Incoming struct {
	From      string       `json:"from"`
	Transfers []TransferIn `json:"transfers"`
}

type Coinbase struct {
	Reward uint64 `json:"reward"`
}

type TransactionEntry struct {
	Hash       string       `json:"hash"`
	Topoheight uint64       `json:"topoheight"`
	Outgoing   *Outgoing    `json:"outgoing"`
	Burn       *daemon.Burn `json:"burn"`
	Incoming   *Incoming    `json:"incoming"`
	Coinbase   *Coinbase    `json:"coinbase"`
}

type ListTransactionsParams struct {
	MinTopoheight  *uint64 `json:"min_topoheight"`
	MaxTopoheight  *uint64 `json:"max_topoheight"`
	Address        *string `json:"address"`
	AcceptIncoming bool    `json:"accept_incoming"`
	AcceptOutgoing bool    `json:"accept_outgoing"`
	AcceptCoinbase bool    `json:"accept_coinbase"`
	AcceptBurn     bool    `json:"accept_burn"`
}

type EstimateFeesParams struct {
	Transfers *[]TransferOut `json:"transfers"`
	Burn      *daemon.Burn   `json:"burn"`
}

type BalanceChangedResult struct {
	Asset   string `json:"asset"`
	Balance uint64 `json:"balance"`
}

type BuildTransactionOfflineParams struct {
	Transfers      []TransferOut          `json:"transfers"`
	Burn           *daemon.Burn           `json:"burn,omitempty"`
	MultiSig       *MutliSigBuilder       `json:"multi_sig,omitempty"`
	InvokeContract *InvokeContractBuilder `json:"invoke_contract,omitempty"`
	DeployContract *string                `json:"deploy_contract,omitempty"`
	Fee            *FeeBuilder            `json:"fee,omitempty"`
	TxVersion      *uint8                 `json:"tx_version,omitempty"`
	TxAsHex        bool                   `json:"tx_as_hex"`
	Reference      daemon.Reference       `json:"reference"`
	Nonce          uint64                 `json:"nonce"`
	Signers        *[]SignerId            `json:"signers,omitempty"`
}

type BuildUnsignedTransactionParams struct {
	Transfers      []TransferOut          `json:"transfers"`
	Burn           *daemon.Burn           `json:"burn,omitempty"`
	MultiSig       *MutliSigBuilder       `json:"multi_sig,omitempty"`
	InvokeContract *InvokeContractBuilder `json:"invoke_contract,omitempty"`
	DeployContract *string                `json:"deploy_contract,omitempty"`
	Nonce          *uint64                `json:"nonce,omitempty"`
	Fee            *FeeBuilder            `json:"fee,omitempty"`
	TxVersion      *uint8                 `json:"tx_version,omitempty"`
	TxAsHex        bool                   `json:"tx_as_hex"`
}

type InnerProductProof struct {
}

type RangeProof struct {
	A            []byte            `json:"A"`
	S            []byte            `json:"S"`
	T_1          []byte            `json:"T_1"`
	T_2          []byte            `json:"T_2"`
	T_X          []byte            `json:"t_x"`
	T_X_Blinding []byte            `json:"t_x_blinding"`
	E_Blinding   []byte            `json:"e_blinding"`
	IppProof     InnerProductProof `json:"ipp_proof"`
}

type UnsignedTransactionResponse struct {
	Version           uint8                     `json:"version"`
	Source            string                    `json:"source"`
	Data              interface{}               `json:"data"`
	Fee               uint64                    `json:"fee"`
	Nonce             uint64                    `json:"nonce"`
	SourceCommitments []daemon.SourceCommitment `json:"source_commitments"`
	Reference         daemon.Reference          `json:"reference"`
	RangeProof        RangeProof                `json:"range_proof"`
	MultiSig          []string                  `json:"multisig"`
	Hash              string                    `json:"hash"`
	Threshold         uint8                     `json:"threshold"`
	TxAsHex           bool                      `json:"tx_as_hex"`
}

type SignUnsignedTransactionParams struct {
	Hash     string `json:"hash"`
	SignerId uint8  `json:"signer_id"`
}

type SignatureId struct {
	Id        uint8  `json:"id"`
	Signature string `json:"signature"`
}

type FinalizeUnsignedTransactionParams struct {
	Unsigned   string        `json:"unsigned"`
	Signatures []SignatureId `json:"signatures"`
	Broadcast  bool          `json:"broadcast"`
	TxAsHex    bool          `json:"tx_as_hex"`
}

type Address struct {
	Mainnet  string      `json:"mainnet"`
	AddrType interface{} `json:"addr_type"`
	Key      string      `json:"key"`
}

type EstimateExtraDataSizeParams struct {
	Destinations []Address `json:"destinations"`
}

type EstimateExtraDataSizeResult struct {
	Size uint8 `json:"size"`
}

type NetworkInfoResult struct {
	// TODO
	ConnectedTo string `json:"connected_to"`
}

type DecryptCiphertextParams struct {
	Ciphertext string `json:"ciphertext"`
}

type TxRole string

const (
	TxSenderRole   TxRole = "sender"
	TxReceiverRole TxRole = "receiver"
)

type DecryptExtraDataParams struct {
	ExtraData []byte `json:"extra_data"`
	Role      TxRole `json:"role"`
}
