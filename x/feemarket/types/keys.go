package types

import (
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

const (
	// ModuleName is the name of the feemarket module.
	ModuleName = "feemarket"
	// StoreKey is the store key string for the feemarket module.
	StoreKey = ModuleName

	// FeeCollectorName the root string for the fee market fee collector account address.
	FeeCollectorName = authtypes.FeeCollectorName
)

const (
	prefixParams        = iota + 1
	prefixState         = iota + 2
	prefixFeeDenomParam = iota + 3
)

var (
	// KeyParams is the store key for the feemarket module's parameters.
	KeyParams = []byte{prefixParams}

	// KeyPrefixState is the prefix of store key for the feemarket module's data.
	KeyState = []byte{prefixState}

	// KeyPrefixFeeDenomParam is the prefix of store key for the fee denom's data.
	KeyPrefixFeeDenomParam = []byte{prefixFeeDenomParam}

	EventTypeFeePay      = "fee_pay"
	EventTypeTipPay      = "tip_pay"
	AttributeKeyTip      = "tip"
	AttributeKeyTipPayer = "tip_payer"
	AttributeKeyTipPayee = "tip_payee"
)

// GetKeyPrefixState returns the KVStore key prefix for storing
// registered feeshare contract for a withdrawer
func GetKeyPrefixFeeDenomParam(feeDenom string) []byte {
	return append(KeyPrefixFeeDenomParam, feeDenom...)
}
