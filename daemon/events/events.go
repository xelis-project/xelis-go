package events

const (
	NewBlock                  string = `new_block`
	TransactionAddedInMempool string = `transaction_added_in_mempool`
	TransactionExecuted       string = `transaction_executed`
	StableHeightChanged       string = `stable_height_changed`
	StableTopoHeightChanged   string = `stable_topo_height_changed`
	TransactionOrphaned       string = `transaction_orphaned`
	DeployContract            string = `deploy_contract`
	BlockOrdered              string = `block_ordered`
	BlockOrphaned             string = `block_orphaned`
	NewAsset                  string = `new_asset`
	PeerConnected             string = `peer_connected`
	PeerDisconnected          string = `peer_disconnect`
	PeerStateUpdated          string = `peer_state_updated`
	PeerPeerListUpdated       string = `peer_peer_list_updated`
	PeerPeerDisconnected      string = `peer_peer_disconnected`
)
