package post_test

import (
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	"github.com/stretchr/testify/mock"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	antesuite "github.com/skip-mev/feemarket/x/feemarket/ante/suite"
	"github.com/skip-mev/feemarket/x/feemarket/types"
)

func TestMultiFeePostHandle(t *testing.T) {
	// Same data for every test case
	gasLimit := antesuite.NewTestGasLimit()
	validFeeAmount := types.DefaultMinBaseFee.MulInt64(int64(gasLimit)).TruncateInt()
	validFee1 := sdk.NewCoins(sdk.NewCoin(types.DefaultFeeDenom, validFeeAmount))
	validFee2 := sdk.NewCoins(sdk.NewCoin(types.TestFeeDenom, validFeeAmount))
	invalidFee := sdk.NewCoins(sdk.NewCoin("invalid", validFeeAmount))

	testCases := []antesuite.TestCase{
		{
			Name: "signer has enough funds, should pass, no tip",
			Malleate: func(s *antesuite.TestSuite) antesuite.TestCaseArgs {
				accs := s.CreateTestAccounts(1)
				s.MockBankKeeper.On("SendCoinsFromAccountToModule", mock.Anything, accs[0].Account.GetAddress(), types.FeeCollectorName, mock.Anything).Return(nil)

				return antesuite.TestCaseArgs{
					Msgs:      []sdk.Msg{testdata.NewTestMsg(accs[0].Account.GetAddress())},
					GasLimit:  gasLimit,
					FeeAmount: validFee1,
				}
			},
			RunAnte:  true,
			RunPost:  true,
			Simulate: false,
			ExpPass:  true,
			ExpErr:   nil,
		},
		{
			Name: "signer has enough funds, should pass, no tip",
			Malleate: func(s *antesuite.TestSuite) antesuite.TestCaseArgs {
				accs := s.CreateTestAccounts(1)
				s.MockBankKeeper.On("SendCoinsFromAccountToModule", mock.Anything, accs[0].Account.GetAddress(), types.FeeCollectorName, mock.Anything).Return(nil)

				return antesuite.TestCaseArgs{
					Msgs:      []sdk.Msg{testdata.NewTestMsg(accs[0].Account.GetAddress())},
					GasLimit:  gasLimit,
					FeeAmount: validFee2,
				}
			},
			RunAnte:  true,
			RunPost:  true,
			Simulate: false,
			ExpPass:  true,
			ExpErr:   nil,
		},
		{
			Name: "signer supplies invalid fee token, should fail",
			Malleate: func(s *antesuite.TestSuite) antesuite.TestCaseArgs {
				accs := s.CreateTestAccounts(1)

				return antesuite.TestCaseArgs{
					Msgs:      []sdk.Msg{testdata.NewTestMsg(accs[0].Account.GetAddress())},
					GasLimit:  gasLimit,
					FeeAmount: invalidFee,
				}
			},
			RunAnte:  true,
			RunPost:  false,
			Simulate: false,
			ExpPass:  false,
			ExpErr:   sdkerrors.ErrKeyNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Case %s", tc.Name), func(t *testing.T) {
			s := antesuite.SetupTestSuite(t, true)
			s.TxBuilder = s.ClientCtx.TxConfig.NewTxBuilder()
			args := tc.Malleate(s)

			s.RunTestCase(t, tc, args)
		})
	}
}
