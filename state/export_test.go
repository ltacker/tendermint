package state

import (
	dbm "github.com/tendermint/tm-db"

	abci "github.com/tendermint/tendermint/abci/types"
	tmstate "github.com/tendermint/tendermint/proto/state"
	tmproto "github.com/tendermint/tendermint/proto/types"
	"github.com/tendermint/tendermint/types"
)

//
// TODO: Remove dependence on all entities exported from this file.
//
// Every entity exported here is dependent on a private entity from the `state`
// package. Currently, these functions are only made available to tests in the
// `state_test` package, but we should not be relying on them for our testing.
// Instead, we should be exclusively relying on exported entities for our
// testing, and should be refactoring exported entities to make them more
// easily testable from outside of the package.
//

const ValSetCheckpointInterval = valSetCheckpointInterval

// UpdateState is an alias for updateState exported from execution.go,
// exclusively and explicitly for testing.
func UpdateState(
	state tmstate.State,
	blockID types.BlockID,
	header *types.Header,
	abciResponses *tmstate.ABCIResponses,
	validatorUpdates []*types.Validator,
) (tmstate.State, error) {
	bi := blockID.ToProto()
	return updateState(state, *bi, header, abciResponses, validatorUpdates)
}

// ValidateValidatorUpdates is an alias for validateValidatorUpdates exported
// from execution.go, exclusively and explicitly for testing.
func ValidateValidatorUpdates(abciUpdates []abci.ValidatorUpdate, params tmproto.ValidatorParams) error {
	return validateValidatorUpdates(abciUpdates, params)
}

// SaveConsensusParamsInfo is an alias for the private saveConsensusParamsInfo
// method in store.go, exported exclusively and explicitly for testing.
func SaveConsensusParamsInfo(db dbm.DB, nextHeight, changeHeight int64, params tmproto.ConsensusParams) {
	saveConsensusParamsInfo(db, nextHeight, changeHeight, params)
}

// SaveValidatorsInfo is an alias for the private saveValidatorsInfo method in
// store.go, exported exclusively and explicitly for testing.
func SaveValidatorsInfo(db dbm.DB, height, lastHeightChanged int64, valSet *tmproto.ValidatorSet) {
	saveValidatorsInfo(db, height, lastHeightChanged, valSet)
}
