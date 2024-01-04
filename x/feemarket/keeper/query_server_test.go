package keeper_test

import (
	"cosmossdk.io/math"

	"github.com/skip-mev/feemarket/x/feemarket/types"
)

func (s *KeeperTestSuite) TestParamsRequest() {
	s.Run("can get default params", func() {
		req := &types.ParamsRequest{}
		resp, err := s.queryServer.Params(s.ctx, req)
		s.Require().NoError(err)
		s.Require().NotNil(resp)

		s.Require().Equal(types.DefaultParams(), resp.Params)

		params, err := s.feeMarketKeeper.GetParams(s.ctx)
		s.Require().NoError(err)

		s.Require().Equal(resp.Params, params)
	})

	s.Run("can get updated params", func() {
		params := types.Params{
			Alpha:                  math.LegacyMustNewDecFromStr("0.1"),
			Beta:                   math.LegacyMustNewDecFromStr("0.1"),
			Theta:                  math.LegacyMustNewDecFromStr("0.1"),
			Delta:                  math.LegacyMustNewDecFromStr("0.1"),
			MinLearningRate:        math.LegacyMustNewDecFromStr("0.1"),
			MaxLearningRate:        math.LegacyMustNewDecFromStr("0.1"),
			TargetBlockUtilization: 5,
			MaxBlockUtilization:    10,
			Window:                 1,
			Enabled:                true,
		}
		err := s.feeMarketKeeper.SetParams(s.ctx, params)
		s.Require().NoError(err)

		req := &types.ParamsRequest{}
		resp, err := s.queryServer.Params(s.ctx, req)
		s.Require().NoError(err)
		s.Require().NotNil(resp)

		s.Require().Equal(params, resp.Params)

		params, err = s.feeMarketKeeper.GetParams(s.ctx)
		s.Require().NoError(err)

		s.Require().Equal(resp.Params, params)
	})
}

func (s *KeeperTestSuite) TestStateRequest() {
	s.Run("can get default state", func() {
		req := &types.StateRequest{
			FeeDenom: "",
		}
		resp, err := s.queryServer.State(s.ctx, req)
		s.Require().NoError(err)
		s.Require().NotNil(resp)

		s.Require().Equal(types.DefaultState(), resp.States)

		states, err := s.feeMarketKeeper.GetStates(s.ctx)
		s.Require().NoError(err)

		s.Require().Equal(resp.States, states)
	})

	s.Run("can get updated state", func() {
		state := types.State{
			FeeDenom:     types.DefaultFeeDenom,
			MinBaseFee:   math.OneInt(),
			BaseFee:      math.OneInt(),
			LearningRate: math.LegacyOneDec(),
			Window:       []uint64{1},
			Index:        0,
		}
		err := s.feeMarketKeeper.SetState(s.ctx, state)
		s.Require().NoError(err)

		req := &types.StateRequest{
			FeeDenom: types.DefaultFeeDenom,
		}
		resp, err := s.queryServer.State(s.ctx, req)
		s.Require().NoError(err)
		s.Require().NotNil(resp)

		s.Require().Equal(state, resp.States[0])

		states, err := s.feeMarketKeeper.GetState(s.ctx, types.DefaultFeeDenom)
		s.Require().NoError(err)

		s.Require().Equal(resp.States[0], states)
	})
}

func (s *KeeperTestSuite) TestBaseFeeRequest() {
	s.Run("can get default base fee", func() {
		req := &types.BaseFeeRequest{
			FeeDenom: types.DefaultFeeDenom,
		}
		resp, err := s.queryServer.BaseFee(s.ctx, req)
		s.Require().NoError(err)
		s.Require().NotNil(resp)

		fees, err := s.feeMarketKeeper.GetMinGasPrices(s.ctx, types.DefaultFeeDenom)
		s.Require().NoError(err)

		s.Require().Equal(resp.Fees, fees)
	})

	// TODO: fix test, after adding update to state
	// s.Run("can get updated base fee", func() {
	// 	state := types.State{
	// 		BaseFee: math.OneInt(),
	// 	}
	// 	err := s.feeMarketKeeper.SetState(s.ctx, state)
	// 	s.Require().NoError(err)

	// 	params := types.Params{
	// 		FeeDenom: "test",
	// 	}
	// 	err = s.feeMarketKeeper.SetParams(s.ctx, params)
	// 	s.Require().NoError(err)

	// 	req := &types.BaseFeeRequest{}
	// 	resp, err := s.queryServer.BaseFee(s.ctx, req)
	// 	s.Require().NoError(err)
	// 	s.Require().NotNil(resp)

	// 	fees, err := s.feeMarketKeeper.GetMinGasPrices(s.ctx)
	// 	s.Require().NoError(err)

	// 	s.Require().Equal(resp.Fees, fees)
	// })
}
