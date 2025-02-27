package types

import "context"

//go:generate ../../scripts/mockery_generate.sh Application

// Application is an interface that enables any finite, deterministic state machine
// to be driven by a blockchain-based replication engine via the ABCI.
type Application interface {
	// Info/Query Connection
	Info(context.Context, *InfoRequest) (*InfoResponse, error)    // Return application info
	Query(context.Context, *QueryRequest) (*QueryResponse, error) // Query for state

	// Mempool Connection
	CheckTx(context.Context, *CheckTxRequest) (*CheckTxResponse, error) // Validate a tx for the mempool

	// Consensus Connection
	InitChain(context.Context, *InitChainRequest) (*InitChainResponse, error) // Initialize blockchain w validators/other info from CometBFT
	PrepareProposal(context.Context, *PrepareProposalRequest) (*PrepareProposalResponse, error)
	ProcessProposal(context.Context, *ProcessProposalRequest) (*ProcessProposalResponse, error)
	// Deliver the decided block with its txs to the Application
	FinalizeBlock(context.Context, *FinalizeBlockRequest) (*FinalizeBlockResponse, error)
	// Create application specific vote extension
	ExtendVote(context.Context, *ExtendVoteRequest) (*ExtendVoteResponse, error)
	// Verify application's vote extension data
	VerifyVoteExtension(context.Context, *VerifyVoteExtensionRequest) (*VerifyVoteExtensionResponse, error)
	// Commit the state and return the application Merkle root hash
	Commit(context.Context, *CommitRequest) (*CommitResponse, error)

	// State Sync Connection
	ListSnapshots(context.Context, *ListSnapshotsRequest) (*ListSnapshotsResponse, error)                // List available snapshots
	OfferSnapshot(context.Context, *OfferSnapshotRequest) (*OfferSnapshotResponse, error)                // Offer a snapshot to the application
	LoadSnapshotChunk(context.Context, *LoadSnapshotChunkRequest) (*LoadSnapshotChunkResponse, error)    // Load a snapshot chunk
	ApplySnapshotChunk(context.Context, *ApplySnapshotChunkRequest) (*ApplySnapshotChunkResponse, error) // Apply a shapshot chunk
}

//-------------------------------------------------------
// BaseApplication is a base form of Application

var _ Application = (*BaseApplication)(nil)

type BaseApplication struct{}

func NewBaseApplication() *BaseApplication {
	return &BaseApplication{}
}

func (BaseApplication) Info(context.Context, *InfoRequest) (*InfoResponse, error) {
	return &InfoResponse{}, nil
}

func (BaseApplication) CheckTx(context.Context, *CheckTxRequest) (*CheckTxResponse, error) {
	return &CheckTxResponse{Code: CodeTypeOK}, nil
}

func (BaseApplication) Commit(context.Context, *CommitRequest) (*CommitResponse, error) {
	return &CommitResponse{}, nil
}

func (BaseApplication) Query(context.Context, *QueryRequest) (*QueryResponse, error) {
	return &QueryResponse{Code: CodeTypeOK}, nil
}

func (BaseApplication) InitChain(context.Context, *InitChainRequest) (*InitChainResponse, error) {
	return &InitChainResponse{}, nil
}

func (BaseApplication) ListSnapshots(context.Context, *ListSnapshotsRequest) (*ListSnapshotsResponse, error) {
	return &ListSnapshotsResponse{}, nil
}

func (BaseApplication) OfferSnapshot(context.Context, *OfferSnapshotRequest) (*OfferSnapshotResponse, error) {
	return &OfferSnapshotResponse{}, nil
}

func (BaseApplication) LoadSnapshotChunk(context.Context, *LoadSnapshotChunkRequest) (*LoadSnapshotChunkResponse, error) {
	return &LoadSnapshotChunkResponse{}, nil
}

func (BaseApplication) ApplySnapshotChunk(context.Context, *ApplySnapshotChunkRequest) (*ApplySnapshotChunkResponse, error) {
	return &ApplySnapshotChunkResponse{}, nil
}

func (BaseApplication) PrepareProposal(_ context.Context, req *PrepareProposalRequest) (*PrepareProposalResponse, error) {
	txs := make([][]byte, 0, len(req.Txs))
	var totalBytes int64
	for _, tx := range req.Txs {
		totalBytes += int64(len(tx))
		if totalBytes > req.MaxTxBytes {
			break
		}
		txs = append(txs, tx)
	}
	return &PrepareProposalResponse{Txs: txs}, nil
}

func (BaseApplication) ProcessProposal(context.Context, *ProcessProposalRequest) (*ProcessProposalResponse, error) {
	return &ProcessProposalResponse{Status: PROCESS_PROPOSAL_STATUS_ACCEPT}, nil
}

func (BaseApplication) ExtendVote(context.Context, *ExtendVoteRequest) (*ExtendVoteResponse, error) {
	return &ExtendVoteResponse{}, nil
}

func (BaseApplication) VerifyVoteExtension(context.Context, *VerifyVoteExtensionRequest) (*VerifyVoteExtensionResponse, error) {
	return &VerifyVoteExtensionResponse{
		Status: VERIFY_VOTE_EXTENSION_STATUS_ACCEPT,
	}, nil
}

func (BaseApplication) FinalizeBlock(_ context.Context, req *FinalizeBlockRequest) (*FinalizeBlockResponse, error) {
	txs := make([]*ExecTxResult, len(req.Txs))
	for i := range req.Txs {
		txs[i] = &ExecTxResult{Code: CodeTypeOK}
	}
	return &FinalizeBlockResponse{
		TxResults: txs,
	}, nil
}
