package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = &MsgFeeDenomParam{}

// NewMsgFeeDenomParam returns a new message to update the x/feemarket module's feeDenomParam for feeDenom.
func NewMsgFeeDenomParam(authority string, fdp FeeDenomParam) MsgFeeDenomParam {
	return MsgFeeDenomParam{
		Authority:     authority,
		FeeDenomParam: fdp,
	}
}

// GetSigners implements GetSigners for the msg.
func (m *MsgFeeDenomParam) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.MustAccAddressFromBech32(m.Authority)}
}

// ValidateBasic determines whether the information in the message is formatted correctly, specifically
// whether the authority is a valid acc-address.
func (m *MsgFeeDenomParam) ValidateBasic() error {
	// validate authority address
	_, err := sdk.AccAddressFromBech32(m.Authority)
	if err != nil {
		return err
	}

	return nil
}
