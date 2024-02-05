package keeper_test

import (
	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/skip-mev/feemarket/x/feemarket/types"
)

func (s *KeeperTestSuite) TestUpdateFeeMarket() {
	s.Run("empty block with default eip1559 with min base fee", func() {
		state := types.DefaultState()
		params := types.DefaultParams()
		fdp := types.DefaultFeeDenomParam()[0]

		s.setGenesis(params, state, fdp)

		s.Require().NoError(s.feeMarketKeeper.UpdateFeeMarket(s.ctx))

		fee, err := s.feeMarketKeeper.GetBaseFee(s.ctx, fdp.FeeDenom)
		s.Require().NoError(err)
		s.Require().Equal(fee, fdp.MinBaseFee)

		s.Require().Equal(math.LegacyMustNewDecFromStr("0.125"), state.LearningRate)
	})

	s.Run("empty block with default eip1559 with preset base fee", func() {
		state := types.DefaultState()
		params := types.DefaultParams()
		fdp := types.DefaultFeeDenomParam()[0]

		fdp.BaseFee = fdp.BaseFee.MulInt64(2)
		s.setGenesis(params, state, fdp)

		s.Require().NoError(s.feeMarketKeeper.UpdateFeeMarket(s.ctx))

		// We expect the base fee to decrease by 1/8th.
		fee, err := s.feeMarketKeeper.GetBaseFee(s.ctx, fdp.FeeDenom)
		s.Require().NoError(err)

		factor := math.LegacyMustNewDecFromStr("0.875")
		expectedFee := fdp.BaseFee.Mul(factor)
		s.Require().Equal(fee, expectedFee)

		lr, err := s.feeMarketKeeper.GetLearningRate(s.ctx)
		s.Require().NoError(err)
		s.Require().Equal(math.LegacyMustNewDecFromStr("0.125"), lr)
	})

	s.Run("empty block default eip1559 with preset base fee that should default to min", func() {
		// Set the base fee to just below the expected threshold decrease of 1/8th. This means it
		// should default to the minimum base fee.
		state := types.DefaultState()
		params := types.DefaultParams()
		fdp := types.DefaultFeeDenomParam()[0]

		factor := math.LegacyMustNewDecFromStr("0.125")
		increase := fdp.BaseFee.Mul(factor).TruncateInt()
		fdp.BaseFee = types.DefaultMinBaseFee.Add(increase.ToLegacyDec()).Sub(math.LegacyOneDec())

		s.setGenesis(params, state, fdp)

		s.Require().NoError(s.feeMarketKeeper.UpdateFeeMarket(s.ctx))

		// We expect the base fee to decrease by 1/8th.
		fee, err := s.feeMarketKeeper.GetBaseFee(s.ctx, fdp.FeeDenom)
		s.Require().NoError(err)
		s.Require().Equal(fee, fdp.MinBaseFee)

		lr, err := s.feeMarketKeeper.GetLearningRate(s.ctx)
		s.Require().NoError(err)
		s.Require().Equal(math.LegacyMustNewDecFromStr("0.125"), lr)
	})

	s.Run("target block with default eip1559 at min base fee", func() {
		state := types.DefaultState()
		params := types.DefaultParams()
		fdp := types.DefaultFeeDenomParam()[0]

		// Reaching the target block size means that we expect this to not
		// increase.
		err := state.Update(params.TargetBlockUtilization, params)
		s.Require().NoError(err)

		s.setGenesis(params, state, fdp)

		s.Require().NoError(s.feeMarketKeeper.UpdateFeeMarket(s.ctx))

		// We expect the base fee to remain the same.
		fee, err := s.feeMarketKeeper.GetBaseFee(s.ctx, fdp.FeeDenom)
		s.Require().NoError(err)
		s.Require().Equal(fee, fdp.MinBaseFee)

		lr, err := s.feeMarketKeeper.GetLearningRate(s.ctx)
		s.Require().NoError(err)
		s.Require().Equal(math.LegacyMustNewDecFromStr("0.125"), lr)
	})

	s.Run("target block with default eip1559 at preset base fee", func() {
		state := types.DefaultState()
		params := types.DefaultParams()
		fdp := types.DefaultFeeDenomParam()[0]

		fdp.BaseFee = fdp.BaseFee.MulInt64(2)
		// Reaching the target block size means that we expect this to not
		// increase.
		err := state.Update(params.TargetBlockUtilization, params)
		s.Require().NoError(err)

		s.setGenesis(params, state, fdp)

		s.Require().NoError(s.feeMarketKeeper.UpdateFeeMarket(s.ctx))

		// We expect the base fee to remain the same.
		fee, err := s.feeMarketKeeper.GetBaseFee(s.ctx, fdp.FeeDenom)
		s.Require().NoError(err)
		s.Require().Equal(fdp.BaseFee, fee)

		lr, err := s.feeMarketKeeper.GetLearningRate(s.ctx)
		s.Require().NoError(err)
		s.Require().Equal(math.LegacyMustNewDecFromStr("0.125"), lr)
	})

	s.Run("max block with default eip1559 at min base fee", func() {
		state := types.DefaultState()
		params := types.DefaultParams()
		fdp := types.DefaultFeeDenomParam()[0]

		// Reaching the target block size means that we expect this to not
		// increase.
		err := state.Update(params.MaxBlockUtilization, params)
		s.Require().NoError(err)

		s.setGenesis(params, state, fdp)

		s.Require().NoError(s.feeMarketKeeper.UpdateFeeMarket(s.ctx))

		// We expect the base fee to increase by 1/8th.
		fee, err := s.feeMarketKeeper.GetBaseFee(s.ctx, fdp.FeeDenom)
		s.Require().NoError(err)

		factor := math.LegacyMustNewDecFromStr("1.125")
		expectedFee := fdp.BaseFee.Mul(factor)
		s.Require().Equal(fee, expectedFee)

		lr, err := s.feeMarketKeeper.GetLearningRate(s.ctx)
		s.Require().NoError(err)
		s.Require().Equal(math.LegacyMustNewDecFromStr("0.125"), lr)
	})

	s.Run("max block with default eip1559 at preset base fee", func() {
		state := types.DefaultState()
		params := types.DefaultParams()
		fdp := types.DefaultFeeDenomParam()[0]

		fdp.BaseFee = fdp.BaseFee.MulInt64(2)
		// Reaching the target block size means that we expect this to not
		// increase.
		err := state.Update(params.MaxBlockUtilization, params)
		s.Require().NoError(err)

		s.setGenesis(params, state, fdp)

		s.Require().NoError(s.feeMarketKeeper.UpdateFeeMarket(s.ctx))

		// We expect the base fee to increase by 1/8th.
		fee, err := s.feeMarketKeeper.GetBaseFee(s.ctx, fdp.FeeDenom)
		s.Require().NoError(err)

		factor := math.LegacyMustNewDecFromStr("1.125")
		expectedFee := fdp.BaseFee.Mul(factor)
		s.Require().Equal(fee, expectedFee)

		lr, err := s.feeMarketKeeper.GetLearningRate(s.ctx)
		s.Require().NoError(err)
		s.Require().Equal(math.LegacyMustNewDecFromStr("0.125"), lr)
	})

	s.Run("in-between min and target block with default eip1559 at min base fee", func() {
		state := types.DefaultState()
		params := types.DefaultParams()
		fdp := types.DefaultFeeDenomParam()[0]

		params.MaxBlockUtilization = 100
		params.TargetBlockUtilization = 50

		err := state.Update(25, params)
		s.Require().NoError(err)

		s.setGenesis(params, state, fdp)

		s.Require().NoError(s.feeMarketKeeper.UpdateFeeMarket(s.ctx))

		// We expect the base fee to remain the same since it is at min base fee.
		fee, err := s.feeMarketKeeper.GetBaseFee(s.ctx, fdp.FeeDenom)
		s.Require().NoError(err)
		s.Require().Equal(fee, fdp.MinBaseFee)

		lr, err := s.feeMarketKeeper.GetLearningRate(s.ctx)
		s.Require().NoError(err)
		s.Require().Equal(math.LegacyMustNewDecFromStr("0.125"), lr)
	})

	s.Run("in-between min and target block with default eip1559 at preset base fee", func() {
		state := types.DefaultState()
		params := types.DefaultParams()
		fdp := types.DefaultFeeDenomParam()[0]

		fdp.BaseFee = fdp.BaseFee.MulInt64(2)

		params.MaxBlockUtilization = 100
		params.TargetBlockUtilization = 50
		err := state.Update(25, params)

		s.Require().NoError(err)

		s.setGenesis(params, state, fdp)

		s.Require().NoError(s.feeMarketKeeper.UpdateFeeMarket(s.ctx))

		// We expect the base fee to decrease by 1/8th * 1/2.
		fee, err := s.feeMarketKeeper.GetBaseFee(s.ctx, fdp.FeeDenom)
		s.Require().NoError(err)

		factor := math.LegacyMustNewDecFromStr("0.9375")
		expectedFee := fdp.BaseFee.Mul(factor)
		s.Require().Equal(fee, expectedFee)

		lr, err := s.feeMarketKeeper.GetLearningRate(s.ctx)
		s.Require().NoError(err)
		s.Require().Equal(math.LegacyMustNewDecFromStr("0.125"), lr)
	})

	s.Run("in-between target and max block with default eip1559 at min base fee", func() {
		state := types.DefaultState()
		params := types.DefaultParams()
		fdp := types.DefaultFeeDenomParam()[0]

		params.MaxBlockUtilization = 100
		params.TargetBlockUtilization = 50

		err := state.Update(75, params)
		s.Require().NoError(err)

		s.setGenesis(params, state, fdp)

		s.Require().NoError(s.feeMarketKeeper.UpdateFeeMarket(s.ctx))

		// We expect the base fee to increase by 1/8th * 1/2.
		fee, err := s.feeMarketKeeper.GetBaseFee(s.ctx, fdp.FeeDenom)
		s.Require().NoError(err)

		factor := math.LegacyMustNewDecFromStr("1.0625")
		expectedFee := fdp.BaseFee.Mul(factor)
		s.Require().Equal(fee, expectedFee)

		lr, err := s.feeMarketKeeper.GetLearningRate(s.ctx)
		s.Require().NoError(err)
		s.Require().Equal(math.LegacyMustNewDecFromStr("0.125"), lr)
	})

	s.Run("in-between target and max block with default eip1559 at preset base fee", func() {
		state := types.DefaultState()
		params := types.DefaultParams()
		fdp := types.DefaultFeeDenomParam()[0]

		fdp.BaseFee = fdp.BaseFee.MulInt64(2)

		params.MaxBlockUtilization = 100
		params.TargetBlockUtilization = 50

		err := state.Update(75, params)
		s.Require().NoError(err)

		s.setGenesis(params, state, fdp)

		s.Require().NoError(s.feeMarketKeeper.UpdateFeeMarket(s.ctx))

		// We expect the base fee to increase by 1/8th * 1/2.
		fee, err := s.feeMarketKeeper.GetBaseFee(s.ctx, fdp.FeeDenom)
		s.Require().NoError(err)

		factor := math.LegacyMustNewDecFromStr("1.0625")
		expectedFee := fdp.BaseFee.Mul(factor)
		s.Require().Equal(fee, expectedFee)

		lr, err := s.feeMarketKeeper.GetLearningRate(s.ctx)
		s.Require().NoError(err)
		s.Require().Equal(math.LegacyMustNewDecFromStr("0.125"), lr)
	})

	s.Run("empty blocks with aimd eip1559 with min base fee", func() {
		state := types.DefaultState()
		params := types.DefaultParams()
		fdp := types.DefaultFeeDenomParam()[0]

		s.setGenesis(params, state, fdp)

		s.Require().NoError(s.feeMarketKeeper.UpdateFeeMarket(s.ctx))

		fee, err := s.feeMarketKeeper.GetBaseFee(s.ctx, fdp.FeeDenom)
		s.Require().NoError(err)
		s.Require().Equal(fee, fdp.MinBaseFee)

		lr, err := s.feeMarketKeeper.GetLearningRate(s.ctx)
		s.Require().NoError(err)
		expectedLR := state.LearningRate.Add(params.Alpha)
		s.Require().Equal(expectedLR, lr)
	})

	s.Run("empty blocks with aimd eip1559 with preset base fee", func() {
		state := types.DefaultAIMDState()
		params := types.DefaultAIMDParams()
		fdp := types.DefaultAIMDFeeDenomParam()[1]

		fdp.BaseFee = fdp.BaseFee.MulInt64(2)
		s.setGenesis(params, state, fdp)

		s.Require().NoError(s.feeMarketKeeper.UpdateFeeMarket(s.ctx))

		// We expect the base fee to decrease by 1/8th and the learning rate to
		// increase by alpha.
		lr, err := s.feeMarketKeeper.GetLearningRate(s.ctx)
		s.Require().NoError(err)
		expectedLR := state.LearningRate.Add(params.Alpha)
		s.Require().Equal(expectedLR, lr)

		fee, err := s.feeMarketKeeper.GetBaseFee(s.ctx, fdp.FeeDenom)
		s.Require().NoError(err)
		factor := math.LegacyOneDec().Add(math.LegacyMustNewDecFromStr("-1.0").Mul(lr))
		expectedFee := fdp.BaseFee.Mul(factor)
		s.Require().Equal(fee, expectedFee)
	})

	s.Run("empty blocks aimd eip1559 with preset base fee that should default to min", func() {
		state := types.DefaultAIMDState()
		params := types.DefaultAIMDParams()
		fdp := types.DefaultAIMDFeeDenomParam()[0]

		lr := math.LegacyMustNewDecFromStr("0.125")
		increase := fdp.BaseFee.Mul(lr).TruncateInt()

		fdp.BaseFee = types.DefaultMinBaseFee.Add(increase.ToLegacyDec()).Sub(math.LegacyOneDec())
		state.LearningRate = lr

		s.setGenesis(params, state, fdp)

		s.Require().NoError(s.feeMarketKeeper.UpdateFeeMarket(s.ctx))

		lr, err := s.feeMarketKeeper.GetLearningRate(s.ctx)
		s.Require().NoError(err)
		expectedLR := state.LearningRate.Add(params.Alpha)
		s.Require().Equal(expectedLR, lr)

		// We expect the base fee to decrease by 1/8th and the learning rate to
		// increase by alpha.
		fee, err := s.feeMarketKeeper.GetBaseFee(s.ctx, fdp.FeeDenom)
		s.Require().NoError(err)
		s.Require().Equal(fee, fdp.MinBaseFee)
	})

	s.Run("target block with aimd eip1559 at min base fee + LR", func() {
		state := types.DefaultAIMDState()
		params := types.DefaultAIMDParams()
		fdp := types.DefaultAIMDFeeDenomParam()[1]

		// Reaching the target block size means that we expect this to not
		// increase.
		for i := 0; i < len(state.Window); i++ {
			state.Window[i] = params.TargetBlockUtilization
		}

		s.setGenesis(params, state, fdp)

		s.Require().NoError(s.feeMarketKeeper.UpdateFeeMarket(s.ctx))

		// We expect the base fee to remain the same and the learning rate to
		// remain at minimum.
		lr, err := s.feeMarketKeeper.GetLearningRate(s.ctx)
		s.Require().NoError(err)
		s.Require().Equal(params.MinLearningRate, lr)

		fee, err := s.feeMarketKeeper.GetBaseFee(s.ctx, fdp.FeeDenom)
		s.Require().NoError(err)
		s.Require().Equal(fdp.BaseFee, fee)
	})

	s.Run("target block with aimd eip1559 at preset base fee + LR", func() {
		state := types.DefaultAIMDState()
		params := types.DefaultAIMDParams()
		fdp := types.DefaultAIMDFeeDenomParam()[1]

		fdp.BaseFee = fdp.BaseFee.MulInt64(2)
		state.LearningRate = math.LegacyMustNewDecFromStr("0.125")

		// Reaching the target block size means that we expect this to not
		// increase.
		for i := 0; i < len(state.Window); i++ {
			state.Window[i] = params.TargetBlockUtilization
		}

		s.setGenesis(params, state, fdp)

		s.Require().NoError(s.feeMarketKeeper.UpdateFeeMarket(s.ctx))

		// We expect the base fee to decrease by 1/8th and the learning rate to
		// decrease by lr * beta.
		lr, err := s.feeMarketKeeper.GetLearningRate(s.ctx)
		s.Require().NoError(err)
		expectedLR := state.LearningRate.Mul(params.Beta)
		s.Require().Equal(expectedLR, lr)

		fee, err := s.feeMarketKeeper.GetBaseFee(s.ctx, fdp.FeeDenom)
		s.Require().NoError(err)
		s.Require().Equal(fdp.BaseFee, fee)
	})
}

func (s *KeeperTestSuite) TestGetBaseFee() {
	s.Run("can retrieve base fee with default eip-1559", func() {
		gs := types.DefaultGenesisState()
		s.feeMarketKeeper.InitGenesis(s.ctx, *gs)

		fdp := gs.FeeDenomParams[1]

		fee, err := s.feeMarketKeeper.GetBaseFee(s.ctx, fdp.FeeDenom)
		s.Require().NoError(err)
		s.Require().Equal(fee, fdp.BaseFee)
	})

	s.Run("can retrieve base fee with aimd eip-1559", func() {
		gs := types.DefaultAIMDGenesisState()
		s.feeMarketKeeper.InitGenesis(s.ctx, *gs)

		fdp := gs.FeeDenomParams[1]

		fee, err := s.feeMarketKeeper.GetBaseFee(s.ctx, fdp.FeeDenom)
		s.Require().NoError(err)
		s.Require().Equal(fee, fdp.BaseFee)
	})
}

func (s *KeeperTestSuite) TestGetLearningRate() {
	s.Run("can retrieve learning rate with default eip-1559", func() {
		gs := types.DefaultGenesisState()
		s.feeMarketKeeper.InitGenesis(s.ctx, *gs)

		lr, err := s.feeMarketKeeper.GetLearningRate(s.ctx)
		s.Require().NoError(err)
		s.Require().Equal(lr, gs.State.LearningRate)
	})

	s.Run("can retrieve learning rate with aimd eip-1559", func() {
		gs := types.DefaultAIMDGenesisState()
		s.feeMarketKeeper.InitGenesis(s.ctx, *gs)

		lr, err := s.feeMarketKeeper.GetLearningRate(s.ctx)
		s.Require().NoError(err)
		s.Require().Equal(lr, gs.State.LearningRate)
	})
}

func (s *KeeperTestSuite) TestGetMinGasPrice() {
	s.Run("can retrieve min gas prices with default eip-1559", func() {
		gs := types.DefaultGenesisState()
		s.feeMarketKeeper.InitGenesis(s.ctx, *gs)

		fdp := gs.FeeDenomParams[1]

		expected := sdk.NewDecCoinFromDec(types.DefaultFeeDenom, fdp.BaseFee)

		mgp, err := s.feeMarketKeeper.GetMinGasPrice(s.ctx, fdp.FeeDenom)
		s.Require().NoError(err)
		s.Require().Equal(expected, mgp)
	})

	s.Run("can retrieve min gas prices with aimd eip-1559", func() {
		gs := types.DefaultAIMDGenesisState()
		s.feeMarketKeeper.InitGenesis(s.ctx, *gs)

		fdp := gs.FeeDenomParams[1]

		expected := sdk.NewDecCoinFromDec(types.DefaultFeeDenom, fdp.BaseFee)

		mgp, err := s.feeMarketKeeper.GetMinGasPrice(s.ctx, fdp.FeeDenom)
		s.Require().NoError(err)
		s.Require().Equal(expected, mgp)
	})
}

func (s *KeeperTestSuite) setGenesis(params types.Params, state types.State, fdp types.FeeDenomParam) {
	gs := types.NewGenesisState(params, state, []types.FeeDenomParam{fdp})
	s.NotPanics(func() {
		s.feeMarketKeeper.InitGenesis(s.ctx, *gs)
	})
}
