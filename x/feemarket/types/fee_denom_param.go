package types

import (
	fmt "fmt"

	"cosmossdk.io/math"
)

// NewFeeDenomParam instantiates a new FeeDenomParam instance.
func NewFeeDenomParam(
	feeDenom string,
	minBaseFee,
	baseFee math.LegacyDec,
) *FeeDenomParam {
	return &FeeDenomParam{
		FeeDenom:   feeDenom,
		MinBaseFee: minBaseFee,
		BaseFee:    baseFee,
	}
}

// ValidateBasic performs basic validation on the feeDenomParam.
func (s *FeeDenomParam) ValidateBasic() error {
	if s.MinBaseFee.IsNil() || s.MinBaseFee.LTE(math.LegacyZeroDec()) {
		return fmt.Errorf("min base fee must be positive")
	}

	if s.BaseFee.IsNil() || s.BaseFee.LT(math.LegacyZeroDec()) {
		return fmt.Errorf("base fee must be positive")
	}

	if s.MinBaseFee.GT(s.BaseFee) {
		return fmt.Errorf("min base fee cannot be greater than base fee")
	}

	if s.FeeDenom == "" {
		return fmt.Errorf("fee denom cannot be empty")
	}

	return nil
}
