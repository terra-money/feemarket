package keeper

import (
	"fmt"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/skip-mev/feemarket/x/feemarket/types"
)

type Keeper struct {
	cdc      codec.BinaryCodec
	storeKey storetypes.StoreKey
	ak       types.AccountKeeper

	// The address that is capable of executing a MsgParams message.
	// Typically, this will be the governance module's address.
	authority string

	// Map of feeDenom to bool of if the state has been updated through a MsgState message this block
	// This is used to prevent updating of the state again in the endblocker
	updatedStateMap map[string]bool
}

// NewKeeper constructs a new feemarket keeper.
func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,
	authKeeper types.AccountKeeper,
	authority string,
) *Keeper {
	if _, err := sdk.AccAddressFromBech32(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address: %s", authority))
	}

	k := &Keeper{
		cdc,
		storeKey,
		authKeeper,
		authority,
		make(map[string]bool),
	}

	return k
}

// Logger returns a feemarket module-specific logger.
func (k *Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

// GetAuthority returns the address that is capable of executing a MsgUpdateParams message.
func (k *Keeper) GetAuthority() string {
	return k.authority
}

// GetState returns the feemarket module's state.
func (k *Keeper) GetState(ctx sdk.Context, feeDenom string) (types.State, error) {
	store := ctx.KVStore(k.storeKey)

	key := types.GetKeyPrefixState(feeDenom)
	bz := store.Get(key)

	state := types.State{}
	if bz == nil {
		return state, sdkerrors.ErrKeyNotFound.Wrapf("state not found for feeDenom: %s", feeDenom)
	}

	if err := state.Unmarshal(bz); err != nil {
		return state, err
	}

	return state, nil
}

// GetStateIter returns an iterator for all feemarket module's states.
func (k *Keeper) GetStateIter(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)

	return sdk.KVStorePrefixIterator(store, types.KeyPrefixState)
}

// GetStates returns all feemarket module's states.
func (k *Keeper) GetStates(ctx sdk.Context) ([]types.State, error) {
	iter := k.GetStateIter(ctx)
	defer iter.Close()

	var states []types.State

	for ; iter.Valid(); iter.Next() {
		state := types.State{}
		if err := state.Unmarshal(iter.Value()); err != nil {
			return nil, err
		}

		states = append(states, state)
	}

	return states, nil
}

// SetState sets the feemarket module's state.
func (k *Keeper) SetState(ctx sdk.Context, state types.State) error {
	store := ctx.KVStore(k.storeKey)

	bz, err := state.Marshal()
	if err != nil {
		return err
	}

	store.Set(types.GetKeyPrefixState(state.FeeDenom), bz)

	return nil
}

// GetParams returns the feemarket module's parameters.
func (k *Keeper) GetParams(ctx sdk.Context) (types.Params, error) {
	store := ctx.KVStore(k.storeKey)

	key := types.KeyParams
	bz := store.Get(key)

	params := types.Params{}
	if err := params.Unmarshal(bz); err != nil {
		return types.Params{}, err
	}

	return params, nil
}

// SetParams sets the feemarket module's parameters.
func (k *Keeper) SetParams(ctx sdk.Context, params types.Params) error {
	store := ctx.KVStore(k.storeKey)

	bz, err := params.Marshal()
	if err != nil {
		return err
	}

	store.Set(types.KeyParams, bz)

	return nil
}
