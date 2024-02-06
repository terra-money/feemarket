package keeper_test

import (
	"testing"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/skip-mev/chaintestutil/encoding"
	"github.com/stretchr/testify/suite"

	"github.com/skip-mev/feemarket/tests/app"
	testkeeper "github.com/skip-mev/feemarket/testutils/keeper"
	"github.com/skip-mev/feemarket/x/feemarket/keeper"
	"github.com/skip-mev/feemarket/x/feemarket/types"
	"github.com/skip-mev/feemarket/x/feemarket/types/mocks"
)

type KeeperTestSuite struct {
	suite.Suite

	accountKeeper    *mocks.AccountKeeper
	feeMarketKeeper  *keeper.Keeper
	encCfg           encoding.TestEncodingConfig
	ctx              sdk.Context
	authorityAccount sdk.AccAddress

	// Message server variables
	msgServer types.MsgServer

	// Query server variables
	queryServer types.QueryServer
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (s *KeeperTestSuite) SetupTest() {
	s.encCfg = encoding.MakeTestEncodingConfig(app.ModuleBasics.RegisterInterfaces)
	s.authorityAccount = authtypes.NewModuleAddress(govtypes.ModuleName)
	s.accountKeeper = mocks.NewAccountKeeper(s.T())
	ctx, tk, tm := testkeeper.NewTestSetup(s.T())

	s.ctx = ctx
	s.feeMarketKeeper = tk.FeeMarketKeeper
	s.msgServer = tm.FeeMarketMsgServer
	s.queryServer = keeper.NewQueryServer(*s.feeMarketKeeper)
}

func (s *KeeperTestSuite) TestFeeDenomParam() {
	s.Run("set and get default eip1559 feeDenomParam", func() {
		fdp := types.DefaultFeeDenomParam()[0]

		err := s.feeMarketKeeper.SetFeeDenomParam(s.ctx, fdp)
		s.Require().NoError(err)

		gotFdp, err := s.feeMarketKeeper.GetFeeDenomParam(s.ctx, fdp.FeeDenom)
		s.Require().NoError(err)

		s.Require().EqualValues(fdp, gotFdp)
	})

	s.Run("set and get aimd eip1559 feeDenomParam", func() {
		fdp := types.DefaultAIMDFeeDenomParam()[0]

		err := s.feeMarketKeeper.SetFeeDenomParam(s.ctx, fdp)
		s.Require().NoError(err)

		gotFdp, err := s.feeMarketKeeper.GetFeeDenomParam(s.ctx, fdp.FeeDenom)
		s.Require().NoError(err)

		s.Require().Equal(fdp, gotFdp)
	})
}

func (s *KeeperTestSuite) TestParams() {
	s.Run("set and get default params", func() {
		params := types.DefaultParams()

		err := s.feeMarketKeeper.SetParams(s.ctx, params)
		s.Require().NoError(err)

		gotParams, err := s.feeMarketKeeper.GetParams(s.ctx)
		s.Require().NoError(err)

		s.Require().EqualValues(params, gotParams)
	})

	s.Run("set and get custom params", func() {
		params := types.Params{
			Alpha:                  math.LegacyMustNewDecFromStr("0.1"),
			Beta:                   math.LegacyMustNewDecFromStr("0.1"),
			Theta:                  math.LegacyMustNewDecFromStr("0.1"),
			MinLearningRate:        math.LegacyMustNewDecFromStr("0.1"),
			MaxLearningRate:        math.LegacyMustNewDecFromStr("0.1"),
			TargetBlockUtilization: 5,
			MaxBlockUtilization:    10,
			Window:                 1,
			Enabled:                true,
		}

		err := s.feeMarketKeeper.SetParams(s.ctx, params)
		s.Require().NoError(err)

		gotParams, err := s.feeMarketKeeper.GetParams(s.ctx)
		s.Require().NoError(err)

		s.Require().EqualValues(params, gotParams)
	})
}
