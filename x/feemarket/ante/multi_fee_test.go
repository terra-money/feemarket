package ante_test

import (
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	_ "github.com/cosmos/cosmos-sdk/x/auth"

	antesuite "github.com/skip-mev/feemarket/x/feemarket/ante/suite"
	"github.com/skip-mev/feemarket/x/feemarket/types"
)

func TestMultiFeeAnteHandle(t *testing.T) {
	// Same data for every test case
	gasLimit := antesuite.NewTestGasLimit()
	validFeeAmount := types.DefaultMinBaseFee.MulRaw(int64(gasLimit))
	invalidFee := sdk.NewCoins(sdk.NewCoin("stake", validFeeAmount), sdk.NewCoin(types.TestFeeDenom, validFeeAmount))

	testCases := []antesuite.TestCase{
		{
			Name: "signer supplies multiple fee tokens should fail",
			Malleate: func(suite *antesuite.TestSuite) antesuite.TestCaseArgs {
				accs := suite.CreateTestAccounts(1)

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
			ExpErr:   sdkerrors.ErrInsufficientFee,
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
