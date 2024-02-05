package keeper

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// UpdateFeeMarket updates the base fee and learning rate based on the
// AIMD learning rate adjustment algorithm. Note that if the fee market
// is disabled, this function will return without updating the fee market.
// This is executed in EndBlock which allows the next block's base fee to
// be readily available for wallets to estimate gas prices.
func (k *Keeper) UpdateFeeMarket(ctx sdk.Context) error {
	params, err := k.GetParams(ctx)
	if err != nil {
		return err
	}

	if !params.Enabled {
		return nil
	}

	state, err := k.GetState(ctx)
	if err != nil {
		panic(err)
	}

	// Update the learning rate based on the block utilization seen in the
	// current block. This is the AIMD learning rate adjustment algorithm.
	newLR := state.UpdateLearningRate(
		params,
	)

	// Increment the height of the state and set the new state.
	state.IncrementHeight()
	if err := k.SetState(ctx, state); err != nil {
		return err
	}

	k.Logger(ctx).Info(
		"updated the fee market state",
		"height", ctx.BlockHeight(),
		"new_learning_rate", newLR,
		"average_block_utilization", state.GetAverageUtilization(params),
		"net_block_utilization", state.GetNetUtilization(params),
	)

	fdps, err := k.GetFeeDenomParams(ctx)
	if err != nil {
		return err
	}

	for _, fdp := range fdps {
		// Update the base fee based with the new learning rate and delta adjustment.
		newBaseFee := fdp.UpdateBaseFee(params, state)

		k.Logger(ctx).Info(
			"updated the feeDenomParam",
			"height", ctx.BlockHeight(),
			"denom", fdp.FeeDenom,
			"new_base_fee", newBaseFee,
		)

		// Set the new feeDenomParam.
		if err := k.SetFeeDenomParam(ctx, fdp); err != nil {
			return err
		}
	}
	return nil
}

// GetBaseFee returns the base fee from the fee market state.
func (k *Keeper) GetBaseFee(ctx sdk.Context, feeDenom string) (math.LegacyDec, error) {
	fdp, err := k.GetFeeDenomParam(ctx, feeDenom)
	if err != nil {
		return math.LegacyDec{}, err
	}

	return fdp.BaseFee, nil
}

// GetLearningRate returns the learning rate from the fee market state.
func (k *Keeper) GetLearningRate(ctx sdk.Context) (math.LegacyDec, error) {
	state, err := k.GetState(ctx)
	if err != nil {
		return math.LegacyDec{}, err
	}

	return state.LearningRate, nil
}

// GetMinGasPrice returns the mininum gas prices as sdk.DecCoin from the fee market state.
func (k *Keeper) GetMinGasPrice(ctx sdk.Context, feeDenom string) (sdk.DecCoin, error) {
	baseFee, err := k.GetBaseFee(ctx, feeDenom)
	if err != nil {
		return sdk.DecCoin{}, err
	}

	minGasPrice := sdk.NewDecCoinFromDec(feeDenom, baseFee)

	return minGasPrice, nil
}
