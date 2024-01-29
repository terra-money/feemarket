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

	s.Run("rejects a req with no state", func() {
		req := &types.MsgState{
			Authority: s.authorityAccount.String(),
		}
		_, err := s.msgServer.State(s.ctx, req)
		s.Require().Error(err)
	})

	s.Run("rejects a req with state that has no feeDenom", func() {
		state := types.DefaultState()[0]
		state.FeeDenom = ""
		req := &types.MsgState{
			Authority: s.authorityAccount.String(),
			State:     state,
		}
		_, err := s.msgServer.State(s.ctx, req)
		s.Require().Error(err)
	})

	s.Run("rejects a req with state that has no feeDenom", func() {
		reqState := types.DefaultState()[0]
		feeDenom := reqState.FeeDenom

		reqState.MinBaseFee = sdkmath.LegacyNewDec(2)
		reqState.BaseFee = reqState.MinBaseFee

		req := &types.MsgState{
			Authority: s.authorityAccount.String(),
			State:     reqState,
		}
		_, err := s.msgServer.State(s.ctx, req)
		s.Require().NoError(err)

		state, err := s.feeMarketKeeper.GetState(s.ctx, feeDenom)
		s.Require().NoError(err)
		s.Require().Equal(state, reqState)
	})
}
