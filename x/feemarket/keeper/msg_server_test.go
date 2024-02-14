package keeper_test

import (
	sdkmath "cosmossdk.io/math"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

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
			Authority:  s.authorityAccount.String(),
			MinBaseFee: fdp.MinBaseFee,
		}
		_, err := s.msgServer.FeeDenomParam(s.ctx, req)
		s.Require().Error(err)
	})

	s.Run("accepts a req with all required fields", func() {
		reqFdp := types.DefaultFeeDenomParam()[0]
		feeDenom := reqFdp.FeeDenom

		reqFdp.MinBaseFee = sdkmath.LegacyOneDec()

		req := &types.MsgFeeDenomParam{
			Authority:  s.authorityAccount.String(),
			FeeDenom:   feeDenom,
			MinBaseFee: reqFdp.MinBaseFee,
		}
		_, err := s.msgServer.FeeDenomParam(s.ctx, req)
		s.Require().NoError(err)

		fdp, err := s.feeMarketKeeper.GetFeeDenomParam(s.ctx, feeDenom)
		s.Require().NoError(err)
		s.Require().Equal(fdp, reqFdp)
	})
}

func (s *KeeperTestSuite) TestMsgRemoveFeeDenomParam() {
	s.Run("rejects a req with invalid signer", func() {
		req := &types.MsgRemoveFeeDenomParam{
			Authority: "invalid",
		}
		_, err := s.msgServer.RemoveFeeDenomParam(s.ctx, req)
		s.Require().Error(err)
	})

	s.Run("rejects a req with no feeDenom", func() {
		req := &types.MsgRemoveFeeDenomParam{
			Authority: s.authorityAccount.String(),
		}
		_, err := s.msgServer.RemoveFeeDenomParam(s.ctx, req)
		s.Require().Error(err)
	})

	s.Run("rejects a req that has empty feeDenom", func() {
		req := &types.MsgRemoveFeeDenomParam{
			Authority: s.authorityAccount.String(),
			FeeDenom:  "",
		}
		_, err := s.msgServer.RemoveFeeDenomParam(s.ctx, req)
		s.Require().Error(err)
	})

	s.Run("accepts a req with all required fields", func() {
		reqFdp := types.DefaultFeeDenomParam()[0]
		feeDenom := reqFdp.FeeDenom

		reqFdp.MinBaseFee = sdkmath.LegacyOneDec()

		req := &types.MsgRemoveFeeDenomParam{
			Authority: s.authorityAccount.String(),
			FeeDenom:  feeDenom,
		}
		_, err := s.msgServer.RemoveFeeDenomParam(s.ctx, req)
		s.Require().NoError(err)

		_, err = s.feeMarketKeeper.GetFeeDenomParam(s.ctx, feeDenom)
		s.Require().Error(err)
		s.Require().True(sdkerrors.ErrKeyNotFound.Is(err))
	})
}
