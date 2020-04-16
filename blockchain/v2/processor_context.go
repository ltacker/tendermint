package v2

import (
	"fmt"

	tmstate "github.com/tendermint/tendermint/proto/state"
	"github.com/tendermint/tendermint/types"
)

type processorContext interface {
	applyBlock(blockID types.BlockID, block *types.Block) error
	verifyCommit(chainID string, blockID types.BlockID, height int64, commit *types.Commit) error
	saveBlock(block *types.Block, blockParts *types.PartSet, seenCommit *types.Commit)
	tmState() tmstate.State
}

type pContext struct {
	store   blockStore
	applier blockApplier
	state   tmstate.State
}

func newProcessorContext(st blockStore, ex blockApplier, s tmstate.State) *pContext {
	return &pContext{
		store:   st,
		applier: ex,
		state:   s,
	}
}

func (pc *pContext) applyBlock(blockID types.BlockID, block *types.Block) error {
	newState, _, err := pc.applier.ApplyBlock(pc.state, blockID, block)
	pc.state = newState
	return err
}

func (pc pContext) tmState() tmstate.State {
	return pc.state
}

func (pc pContext) verifyCommit(chainID string, blockID types.BlockID, height int64, commit *types.Commit) error {
	return pc.state.Validators.VerifyCommit(chainID, blockID, height, commit)
}

func (pc *pContext) saveBlock(block *types.Block, blockParts *types.PartSet, seenCommit *types.Commit) {
	pc.store.SaveBlock(block, blockParts, seenCommit)
}

type mockPContext struct {
	applicationBL  []int64
	verificationBL []int64
	state          tmstate.State
}

func newMockProcessorContext(
	state tmstate.State,
	verificationBlackList []int64,
	applicationBlackList []int64) *mockPContext {
	return &mockPContext{
		applicationBL:  applicationBlackList,
		verificationBL: verificationBlackList,
		state:          state,
	}
}

func (mpc *mockPContext) applyBlock(blockID types.BlockID, block *types.Block) error {
	for _, h := range mpc.applicationBL {
		if h == block.Height {
			return fmt.Errorf("generic application error")
		}
	}
	mpc.state.LastBlockHeight = block.Height
	return nil
}

func (mpc *mockPContext) verifyCommit(chainID string, blockID types.BlockID, height int64, commit *types.Commit) error {
	for _, h := range mpc.verificationBL {
		if h == height {
			return fmt.Errorf("generic verification error")
		}
	}
	return nil
}

func (mpc *mockPContext) saveBlock(block *types.Block, blockParts *types.PartSet, seenCommit *types.Commit) {

}

func (mpc *mockPContext) tmState() tmstate.State {
	return mpc.state
}
