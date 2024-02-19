package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/skip-mev/feemarket/x/feemarket/types"
)

var _ types.MsgServer = (*MsgServer)(nil)

// MsgServer is the server API for x/feemarket Msg service.
type MsgServer struct {
	k Keeper
}

// NewMsgServer returns the MsgServer implementation.
func NewMsgServer(k Keeper) types.MsgServer {
	return &MsgServer{k}
}

// Params defines a method that updates the module's parameters. The signer of the message must
// be the module authority.
func (ms MsgServer) Params(goCtx context.Context, msg *types.MsgParams) (*types.MsgParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Authority != ms.k.GetAuthority() {
		return nil, fmt.Errorf("invalid authority to execute message")
	}

	params := msg.Params
	if err := ms.k.SetParams(ctx, params); err != nil {
		return nil, fmt.Errorf("error setting params: %w", err)
	}
	state := ms.k.GetState(ctx)

	if int(params.Window) != len(state.Window) {
		window := make([]uint64, params.Window)
		for i := range window {
			if i < len(state.Window) {
				window[i] = state.Window[i]
			} else {
				window[i] = 0
			}
		}
		state.Window = window
		ms.k.SetState(ctx, state)
	}

	// newState := types.NewState(params.Window, params.MinBaseFee, params.MinLearningRate)
	// if err := ms.k.SetState(ctx, newState); err != nil {
	// 	return nil, fmt.Errorf("error setting state: %w", err)
	// }

	return &types.MsgParamsResponse{}, nil
}

// FeeDenomParam defines a method that adds/updates the module's state for a feeDenom. The signer of the message must
// be the module authority.
func (ms MsgServer) FeeDenomParam(goCtx context.Context, msg *types.MsgFeeDenomParam) (*types.MsgFeeDenomParamResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Authority != ms.k.GetAuthority() {
		return nil, fmt.Errorf("invalid authority to execute message")
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, fmt.Errorf("invalid MsgFeeDenomParam: %w", err)
	}

	fdp, err := ms.k.GetFeeDenomParam(ctx, msg.FeeDenom)
	if sdkerrors.ErrKeyNotFound.Is(err) {
		fdp = types.NewFeeDenomParam(msg.FeeDenom, msg.MinBaseFee, msg.MinBaseFee)
	} else if err != nil {
		return nil, err
	}

	fdp.MinBaseFee = msg.MinBaseFee
	if fdp.BaseFee.LT(msg.MinBaseFee) {
		fdp.BaseFee = msg.MinBaseFee
	}

	if err := ms.k.SetFeeDenomParam(ctx, *fdp); err != nil {
		return nil, err
	}

	return &types.MsgFeeDenomParamResponse{}, nil
}

// RemoveFeeDenomParam defines a method that removes the module's state for a feeDenom. The signer of the message must
// be the module authority.
func (ms MsgServer) RemoveFeeDenomParam(goCtx context.Context, msg *types.MsgRemoveFeeDenomParam) (*types.MsgRemoveFeeDenomParamResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Authority != ms.k.GetAuthority() {
		return nil, fmt.Errorf("invalid authority to execute message")
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, fmt.Errorf("invalid MsgRemoveFeeDenomParam: %w", err)
	}

	if err := ms.k.DeleteFeeDenomParam(ctx, msg.FeeDenom); err != nil {
		return nil, err
	}

	return &types.MsgRemoveFeeDenomParamResponse{}, nil
}
