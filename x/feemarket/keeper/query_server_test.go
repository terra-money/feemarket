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

func (s *KeeperTestSuite) TestFeeDenomParamRequest() {
	s.Run("can get default feeDenomParam", func() {
		req := &types.FeeDenomParamRequest{
			FeeDenom: "",
		}
		resp, err := s.queryServer.FeeDenomParam(s.ctx, req)
		s.Require().NoError(err)
		s.Require().NotNil(resp)

		s.Require().Equal(types.DefaultFeeDenomParam(), resp.FeeDenomParams)

		fdps, err := s.feeMarketKeeper.GetFeeDenomParams(s.ctx)
		s.Require().NoError(err)

		s.Require().Equal(resp.FeeDenomParams, fdps)
	})

	s.Run("can get updated feeDenomParam", func() {
		fdp := types.FeeDenomParam{
			MinBaseFee: math.LegacyOneDec(),
			BaseFee:    math.LegacyOneDec(),
			FeeDenom:   types.DefaultFeeDenom,
		}
		err := s.feeMarketKeeper.SetFeeDenomParam(s.ctx, fdp)
		s.Require().NoError(err)

		req := &types.FeeDenomParamRequest{
			FeeDenom: types.DefaultFeeDenom,
		}
		resp, err := s.queryServer.FeeDenomParam(s.ctx, req)
		s.Require().NoError(err)
		s.Require().NotNil(resp)

		s.Require().Equal(fdp, resp.FeeDenomParams[0])

		fdp, err = s.feeMarketKeeper.GetFeeDenomParam(s.ctx, types.DefaultFeeDenom)
		s.Require().NoError(err)

		s.Require().Equal(resp.FeeDenomParams[0], fdp)
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

		fees, err := s.feeMarketKeeper.GetMinGasPrice(s.ctx, types.DefaultFeeDenom)
		s.Require().NoError(err)

		s.Require().Equal(resp.Fee, fees)
	})

	s.Run("can get updated base fee", func() {
		fdp, err := s.feeMarketKeeper.GetFeeDenomParam(s.ctx, types.DefaultFeeDenom)
		s.Require().NoError(err)

		fdp.BaseFee = math.LegacyOneDec()

		err = s.feeMarketKeeper.SetFeeDenomParam(s.ctx, fdp)
		s.Require().NoError(err)

		req := &types.BaseFeeRequest{
			FeeDenom: types.DefaultFeeDenom,
		}
		resp, err := s.queryServer.BaseFee(s.ctx, req)
		s.Require().NoError(err)
		s.Require().NotNil(resp)

		fee, err := s.feeMarketKeeper.GetMinGasPrice(s.ctx, types.DefaultFeeDenom)
		s.Require().NoError(err)

		s.Require().Equal(resp.Fee, fee)
	})
}
