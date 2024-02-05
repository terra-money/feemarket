package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/skip-mev/feemarket/x/feemarket/types"
)

// InitGenesis initializes the feemarket module's state from a given genesis state.
func (k *Keeper) InitGenesis(ctx sdk.Context, gs types.GenesisState) {
	if err := gs.ValidateBasic(); err != nil {
		panic(err)
	}

	// Initialize the fee market state and parameters.
	if err := k.SetParams(ctx, gs.Params); err != nil {
		panic(err)
	}

	if err := k.SetState(ctx, gs.State); err != nil {
		panic(err)
	}

	for _, fdp := range gs.FeeDenomParams {
		if err := k.SetFeeDenomParam(ctx, fdp); err != nil {
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
	state, err := k.GetState(ctx)
	if err != nil {
		panic(err)
	}

	// Get the feemarket module's fee denom parameters.
	fdps, err := k.GetFeeDenomParams(ctx)
	if err != nil {
		panic(err)
	}

	return types.NewGenesisState(params, state, fdps)
}
