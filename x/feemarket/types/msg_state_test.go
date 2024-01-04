package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/skip-mev/feemarket/x/feemarket/types"
)

func TestMsgState(t *testing.T) {
	t.Run("should reject a message with an invalid authority address", func(t *testing.T) {
		msg := types.NewMsgState("invalid", types.DefaultState()[0])
		err := msg.ValidateBasic()
		require.Error(t, err)
	})

	t.Run("should accept an empty message with a valid authority address", func(t *testing.T) {
		msg := types.NewMsgState(sdk.AccAddress("test").String(), types.DefaultState()[0])
		err := msg.ValidateBasic()
		require.NoError(t, err)
	})
}
