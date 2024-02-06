package keeper_test

import (
	sdkmath "cosmossdk.io/math"
	"github.com/skip-mev/feemarket/x/feemarket/types"
)

func (s *KeeperTestSuite) TestMsgParams() {
	s.Run("accepts a req with no params", func() {
		req := &types.MsgParams{
			Authority: s.authorityAccount.String(),
		}
		resp, err := s.msgServer.Params(s.ctx, req)
		s.Require().NoError(err)
		s.Require().NotNil(resp)

		params, err := s.feeMarketKeeper.GetParams(s.ctx)
		s.Require().NoError(err)
		s.Require().False(params.Enabled)
	})

	s.Run("accepts a req with params", func() {
		req := &types.MsgParams{
			Authority: s.authorityAccount.String(),
			Params:    types.DefaultParams(),
		}
		resp, err := s.msgServer.Params(s.ctx, req)
		s.Require().NoError(err)
		s.Require().NotNil(resp)

		params, err := s.feeMarketKeeper.GetParams(s.ctx)
		s.Require().NoError(err)
		s.Require().Equal(req.Params, params)
	})

	s.Run("rejects a req with invalid signer", func() {
		req := &types.MsgParams{
			Authority: "invalid",
		}
		_, err := s.msgServer.Params(s.ctx, req)
		s.Require().Error(err)
	})
}

func (s *KeeperTestSuite) TestMsgFeeDenomParam() {
	s.Run("rejects a req with invalid signer", func() {
		req := &types.MsgFeeDenomParam{
			Authority: "invalid",
		}
		_, err := s.msgServer.FeeDenomParam(s.ctx, req)
		s.Require().Error(err)
	})

	s.Run("rejects a req with no feeDenomParam", func() {
		req := &types.MsgFeeDenomParam{
			Authority: s.authorityAccount.String(),
		}
		_, err := s.msgServer.FeeDenomParam(s.ctx, req)
		s.Require().Error(err)
	})

	s.Run("rejects a req with feeDenomParam that has no feeDenom", func() {
		fdp := types.DefaultFeeDenomParam()[0]
		fdp.FeeDenom = ""
		req := &types.MsgFeeDenomParam{
			Authority:     s.authorityAccount.String(),
			FeeDenomParam: fdp,
		}
		_, err := s.msgServer.FeeDenomParam(s.ctx, req)
		s.Require().Error(err)
	})

	s.Run("accepts a req with all required fields", func() {
		reqFdp := types.DefaultFeeDenomParam()[0]
		feeDenom := reqFdp.FeeDenom

		reqFdp.MinBaseFee = sdkmath.LegacyNewDec(2)
		reqFdp.BaseFee = reqFdp.MinBaseFee

		req := &types.MsgFeeDenomParam{
			Authority:     s.authorityAccount.String(),
			FeeDenomParam: reqFdp,
		}
		_, err := s.msgServer.FeeDenomParam(s.ctx, req)
		s.Require().NoError(err)

		fdp, err := s.feeMarketKeeper.GetFeeDenomParam(s.ctx, feeDenom)
		s.Require().NoError(err)
		s.Require().Equal(fdp, reqFdp)
	})
}
