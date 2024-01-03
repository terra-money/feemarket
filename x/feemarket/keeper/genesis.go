package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/skip-mev/feemarket/x/feemarket/types"
)

// InitGenesis initializes the feemarket module's state from a given genesis state.
func (k *Keeper) InitGenesis(ctx sdk.Context, gs types.GenesisState) {
	if err := gs.ValidateBasic(); err != nil {
		panic(err)
	}

	for _, state := range gs.States {
		if gs.Params.Window != uint64(len(state.Window)) {
			panic(fmt.Sprintf("genesis state and parameters do not match for window for denom: %s", state.FeeDenom))
		}
	}

	// Initialize the fee market state and parameters.
	if err := k.SetParams(ctx, gs.Params); err != nil {
		panic(err)
	}

	for _, state := range gs.States {
		if err := k.SetState(ctx, state); err != nil {
			panic(err)
		}
	}
}

// ExportGenesis returns a GenesisState for a given context.
func (k *Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	// Get the feemarket module's parameters.
	params, err := k.GetParams(ctx)
	if err != nil {
		panic(err)
	}

	// Get the feemarket module's state.
	states, err := k.GetStates(ctx)
	if err != nil {
		panic(err)
	}

	return types.NewGenesisState(params, states)
}
