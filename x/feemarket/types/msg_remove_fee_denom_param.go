package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = &MsgRemoveFeeDenomParam{}

// MsgRemoveFeeDenomParam returns a new message to update the x/feemarket module's feeDenomParam for feeDenom.
func NewMsgRemoveFeeDenomParam(authority string, feeDenom string) MsgRemoveFeeDenomParam {
	return MsgRemoveFeeDenomParam{
		Authority: authority,
		FeeDenom:  feeDenom,
	}
}

// GetSigners implements GetSigners for the msg.
func (m *MsgRemoveFeeDenomParam) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Authority)}
}

// ValidateBasic determines whether the information in the message is formatted correctly, specifically
// whether the authority is a valid acc-address.
func (m *MsgRemoveFeeDenomParam) ValidateBasic() error {
	// validate authority address
	_, err := sdk.AccAddressFromBech32(m.Authority)
	if err != nil {
		return err
	}

	if m.FeeDenom == "" {
		return fmt.Errorf("fee denom cannot be empty")
	}

	return nil
}
