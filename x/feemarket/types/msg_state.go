package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = &MsgState{}

// NewMsgState returns a new message to update the x/feemarket module's state for feeDenom.
func NewMsgState(authority string, state State) MsgState {
	return MsgState{
		Authority: authority,
		State:     state,
	}
}

// GetSigners implements GetSigners for the msg.
func (m *MsgState) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Authority)}
}

// ValidateBasic determines whether the information in the message is formatted correctly, specifically
// whether the authority is a valid acc-address.
func (m *MsgState) ValidateBasic() error {
	// validate authority address
	_, err := sdk.AccAddressFromBech32(m.Authority)
	if err != nil {
		return err
	}

	return nil
}
