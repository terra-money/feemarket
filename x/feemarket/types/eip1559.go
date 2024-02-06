package types

import (
	"cosmossdk.io/math"
)

// Note: We use the same default values as Ethereum for the EIP-1559
// fee market implementation. These parameters do not implement the
// AIMD learning rate adjustment algorithm.

var (
	// DefaultWindow is the default window size for the sliding window
	// used to calculate the base fee. In the base EIP-1559 implementation,
	// only the previous block is considered.
	DefaultWindowSize uint64 = 1

	// DefaultAlpha is not used in the base EIP-1559 implementation.
	DefaultAlpha = math.LegacyMustNewDecFromStr("0.0")

	// DefaultBeta is not used in the base EIP-1559 implementation.
	DefaultBeta = math.LegacyMustNewDecFromStr("1.0")

	// DefaultTheta is not used in the base EIP-1559 implementation.
	DefaultTheta = math.LegacyMustNewDecFromStr("0.0")

	// DefaultTargetBlockUtilization is the default target block utilization. This is the default
	// on Ethereum. This denominated in units of gas consumed in a block.
	DefaultTargetBlockUtilization uint64 = 15_000_000

	// DefaultMaxBlockUtilization is the default maximum block utilization. This is the default
	// on Ethereum. This denominated in units of gas consumed in a block.
	DefaultMaxBlockUtilization uint64 = 30_000_000

	// DefaultMinBaseFee is the default minimum base fee. This is the default
	// on Ethereum. Note that Ethereum is denominated in 1e18 wei. Cosmos chains will
	// likely want to change this to 1e6.
	DefaultMinBaseFee = math.LegacyNewDec(1_000_000)

	// DefaultMinLearningRate is not used in the base EIP-1559 implementation.
	DefaultMinLearningRate = math.LegacyMustNewDecFromStr("0.125")

	// DefaultMaxLearningRate is not used in the base EIP-1559 implementation.
	DefaultMaxLearningRate = math.LegacyMustNewDecFromStr("0.125")

	// DefaultFeeDenom is set to uluna
	DefaultFeeDenom = "uluna"

	// TestFeeDenom is the other fee denom for testing purpose.
	TestFeeDenom = "fee"

	// DefaultIndex is the default index for state
	DefaultIndex uint64 = 0
)

// DefaultParams returns a default set of parameters that implements
// the EIP-1559 fee market implementation without the AIMD learning
// rate adjustment algorithm.
func DefaultParams() Params {
	return NewParams(
		DefaultWindowSize,
		DefaultAlpha,
		DefaultBeta,
		DefaultTheta,
		DefaultTargetBlockUtilization,
		DefaultMaxBlockUtilization,
		DefaultMinLearningRate,
		DefaultMaxLearningRate,
		true,
		DefaultFeeDenom,
	)
}

// DefaultState returns the default state for the EIP-1559 fee market
// implementation without the AIMD learning rate adjustment algorithm.
func DefaultState() State {
	return NewState(
		DefaultWindowSize,
		DefaultMinLearningRate,
		DefaultIndex,
	)
}

// DefaultFeeDenomParam returns the default state for the EIP-1559 fee market
// implementation without the AIMD learning rate adjustment algorithm.
func DefaultFeeDenomParam() []FeeDenomParam {
	return []FeeDenomParam{
		NewFeeDenomParam(
			TestFeeDenom,
			DefaultMinBaseFee,
			DefaultMinBaseFee,
		),
		NewFeeDenomParam(
			DefaultFeeDenom,
			math.LegacyMustNewDecFromStr("0.0015"),
			math.LegacyMustNewDecFromStr("0.0015"),
		),
	}
}

// DefaultGenesisState returns a default genesis state that implements
// the EIP-1559 fee market implementation without the AIMD learning
// rate adjustment algorithm.
func DefaultGenesisState() *GenesisState {
	return NewGenesisState(DefaultParams(), DefaultState(), DefaultFeeDenomParam())
}
