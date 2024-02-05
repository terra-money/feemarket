package types_test

import (
	"testing"

	"cosmossdk.io/math"
	"github.com/stretchr/testify/require"

	"github.com/skip-mev/feemarket/x/feemarket/types"
)

func FuzzDefaultFeeMarket(f *testing.F) {
	testCases := []uint64{
		0,
		1_000,
		10_000,
		100_000,
		1_000_000,
		10_000_000,
		100_000_000,
	}

	for _, tc := range testCases {
		f.Add(tc)
	}

	defaultLR := math.LegacyMustNewDecFromStr("0.125")

	// Default fee market.
	f.Fuzz(func(t *testing.T, blockGasUsed uint64) {
		state := types.DefaultState()
		params := types.DefaultParams()
		fdp := types.DefaultFeeDenomParam()[0]

		fdp.MinBaseFee = math.LegacyNewDec(100)
		fdp.BaseFee = math.LegacyNewDec(200)
		err := state.Update(blockGasUsed, params)

		if blockGasUsed > params.MaxBlockUtilization {
			require.Error(t, err)
			return
		}

		require.NoError(t, err)
		require.Equal(t, blockGasUsed, state.Window[state.Index])

		// Ensure the learning rate is always the default learning rate.
		lr := state.UpdateLearningRate(
			params,
		)
		require.Equal(t, defaultLR, lr)

		oldFee := fdp.BaseFee
		newFee := fdp.UpdateBaseFee(params, state)

		if blockGasUsed > params.TargetBlockUtilization {
			require.True(t, newFee.GT(oldFee))
		} else {
			require.True(t, newFee.LT(oldFee))
		}
	})
}

func FuzzAIMDFeeMarket(f *testing.F) {
	testCases := []uint64{
		0,
		1_000,
		10_000,
		100_000,
		1_000_000,
		10_000_000,
		100_000_000,
	}

	for _, tc := range testCases {
		f.Add(tc)
	}

	// Fee market with adjustable learning rate.
	f.Fuzz(func(t *testing.T, blockGasUsed uint64) {
		state := types.DefaultAIMDState()
		params := types.DefaultAIMDParams()
		fdp := types.DefaultAIMDFeeDenomParam()[0]

		fdp.MinBaseFee = math.LegacyNewDec(100)
		fdp.BaseFee = math.LegacyNewDec(200)
		state.Window = make([]uint64, 1)
		err := state.Update(blockGasUsed, params)

		if blockGasUsed > params.MaxBlockUtilization {
			require.Error(t, err)
			return
		}

		require.NoError(t, err)
		require.Equal(t, blockGasUsed, state.Window[state.Index])

		oldFee := fdp.BaseFee
		newFee := fdp.UpdateBaseFee(params, state)

		if blockGasUsed > params.TargetBlockUtilization {
			require.True(t, newFee.GT(oldFee))
		} else {
			require.True(t, newFee.LT(oldFee))
		}
	})
}
