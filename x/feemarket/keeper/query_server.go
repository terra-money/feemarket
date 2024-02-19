package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/skip-mev/feemarket/x/feemarket/types"
)

var _ types.QueryServer = (*QueryServer)(nil)

// QueryServer defines the gRPC server for the x/feemarket module.
type QueryServer struct {
	k Keeper
}

// NewQueryServer creates a new instance of the x/feemarket QueryServer type.
func NewQueryServer(keeper Keeper) types.QueryServer {
	return &QueryServer{k: keeper}
}

// Params defines a method that returns the current feemarket parameters.
func (q QueryServer) Params(goCtx context.Context, _ *types.ParamsRequest) (*types.ParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	params, err := q.k.GetParams(ctx)
	return &types.ParamsResponse{Params: params}, err
}

// State defines a method that returns the current feemarket states. If feeDenom is nil, return all states
func (q QueryServer) State(goCtx context.Context, query *types.StateRequest) (*types.StateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	state := q.k.GetState(ctx)
	return &types.StateResponse{State: state}, nil
}

// State defines a method that returns the current feemarket states. If feeDenom is nil, return all states
func (q QueryServer) FeeDenomParam(goCtx context.Context, query *types.FeeDenomParamRequest) (*types.FeeDenomParamResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if query.FeeDenom == "" {
		fdps, err := q.k.GetFeeDenomParams(ctx)
		return &types.FeeDenomParamResponse{FeeDenomParams: fdps}, err
	}

	fdp, err := q.k.GetFeeDenomParam(ctx, query.FeeDenom)
	return &types.FeeDenomParamResponse{FeeDenomParams: []types.FeeDenomParam{*fdp}}, err
}
