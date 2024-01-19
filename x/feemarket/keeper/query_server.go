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

	if query.FeeDenom == "" {
		states, err := q.k.GetStates(ctx)
		return &types.StateResponse{States: states}, err
	}

	state, err := q.k.GetState(ctx, query.FeeDenom)
	return &types.StateResponse{States: []types.State{state}}, err
}

// BaseFee defines a method that returns the current feemarket base fee.
func (q QueryServer) BaseFee(goCtx context.Context, query *types.BaseFeeRequest) (*types.BaseFeeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	fee, err := q.k.GetMinGasPrice(ctx, query.FeeDenom)
	return &types.BaseFeeResponse{Fee: fee}, err
}
