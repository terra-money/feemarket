package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/skip-mev/feemarket/x/feemarket/types"
)

func TestMsgRemoveFeeDenomParam(t *testing.T) {
	t.Run("should reject a message with an invalid authority address", func(t *testing.T) {
		fdp := types.DefaultFeeDenomParam()[0]
		msg := types.NewMsgRemoveFeeDenomParam("invalid", fdp.FeeDenom)
		err := msg.ValidateBasic()
		require.Error(t, err)
	})

	t.Run("should accept an empty message with a valid authority address", func(t *testing.T) {
		fdp := types.DefaultFeeDenomParam()[0]
		msg := types.NewMsgRemoveFeeDenomParam(sdk.AccAddress("test").String(), fdp.FeeDenom)
		err := msg.ValidateBasic()
		require.NoError(t, err)
	})
}
