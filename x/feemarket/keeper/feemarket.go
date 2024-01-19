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

	states, err := k.GetStates(ctx)
	if err != nil {
		return err
	}

	for _, state := range states {

		// Update the learning rate based on the block utilization seen in the
		// current block. This is the AIMD learning rate adjustment algorithm.
		newLR := state.UpdateLearningRate(
			params,
		)

		// Update the base fee based with the new learning rate and delta adjustment.
		newBaseFee := state.UpdateBaseFee(params)

		k.Logger(ctx).Info(
			"updated the fee market",
			"height", ctx.BlockHeight(),
			"denom", state.FeeDenom,
			"new_base_fee", newBaseFee,
			"new_learning_rate", newLR,
			"average_block_utilization", state.GetAverageUtilization(params),
			"net_block_utilization", state.GetNetUtilization(params),
		)

		// Increment the height of the state and set the new state.
		state.IncrementHeight()
		if err := k.SetState(ctx, state); err != nil {
			return err
		}
	}
	return nil
}

// GetBaseFee returns the base fee from the fee market state.
func (k *Keeper) GetBaseFee(ctx sdk.Context, feeDenom string) (math.LegacyDec, error) {
	state, err := k.GetState(ctx, feeDenom)
	if err != nil {
		return math.LegacyDec{}, err
	}

	return state.BaseFee, nil
}

// GetLearningRate returns the learning rate from the fee market state.
func (k *Keeper) GetLearningRate(ctx sdk.Context, feeDenom string) (math.LegacyDec, error) {
	state, err := k.GetState(ctx, feeDenom)
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
