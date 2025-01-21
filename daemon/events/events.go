package events

const (
	NewBlock                  string = `new_block`
	BlockOrdered              string = `block_ordered`
	BlockOrphaned             string = `block_orphaned`
	StableHeightChanged       string = `stable_height_changed`
	StableTopoheightChanged   string = `stable_topo_height_changed`
	TransactionOrphaned       string = `transaction_orphaned`
	TransactionAddedInMempool string = `transaction_added_in_mempool`
	TransactionExecuted       string = `transaction_executed`
	InvokeContract            string = `invoke_contract`
	DeployContract            string = `deploy_contract`
	NewAsset                  string = `new_asset`
	PeerConnected             string = `peer_connected`
	PeerDisconnected          string = `peer_disconnected`
	PeerStateUpdated          string = `peer_state_updated`
	PeerPeerListUpdated       string = `peer_peer_list_updated`
	PeerPeerDisconnected      string = `peer_peer_disconnected`
)
