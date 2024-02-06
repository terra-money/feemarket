package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/skip-mev/feemarket/x/feemarket/types"
)

func TestMsgState(t *testing.T) {
	t.Run("should reject a message with an invalid authority address", func(t *testing.T) {
		fdp := types.DefaultFeeDenomParam()[0]
		msg := types.NewMsgFeeDenomParam("invalid", fdp.FeeDenom, fdp.MinBaseFee)
		err := msg.ValidateBasic()
		require.Error(t, err)
	})

	t.Run("should accept an empty message with a valid authority address", func(t *testing.T) {
		fdp := types.DefaultFeeDenomParam()[0]
		msg := types.NewMsgFeeDenomParam(sdk.AccAddress("test").String(), fdp.FeeDenom, fdp.MinBaseFee)
		err := msg.ValidateBasic()
		require.NoError(t, err)
	})
}
