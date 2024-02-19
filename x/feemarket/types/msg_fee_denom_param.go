package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"cosmossdk.io/math"
)

var _ sdk.Msg = &MsgFeeDenomParam{}

// NewMsgFeeDenomParam returns a new message to update the x/feemarket module's feeDenomParam for feeDenom.
func NewMsgFeeDenomParam(authority string, feeDenom string, minBaseFee math.LegacyDec) MsgFeeDenomParam {
	return MsgFeeDenomParam{
		Authority:  authority,
		FeeDenom:   feeDenom,
		MinBaseFee: minBaseFee,
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

	if m.MinBaseFee.IsNil() || m.MinBaseFee.LTE(math.LegacyZeroDec()) {
		return fmt.Errorf("min base fee must be positive")
	}

	if m.FeeDenom == "" {
		return fmt.Errorf("fee denom cannot be empty")
	}

	return nil
}
