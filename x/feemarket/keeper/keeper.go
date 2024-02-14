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

// GetFeeDenomParam returns the feeDenomParam for feeDenom.
func (k *Keeper) GetFeeDenomParam(ctx sdk.Context, feeDenom string) (types.FeeDenomParam, error) {
	store := ctx.KVStore(k.storeKey)

	key := types.GetKeyPrefixFeeDenomParam(feeDenom)
	bz := store.Get(key)

	fdp := types.FeeDenomParam{}
	if bz == nil {
		return fdp, sdkerrors.ErrKeyNotFound.Wrapf("feeDenomParam not found for feeDenom: %s", feeDenom)
	}

	if err := fdp.Unmarshal(bz); err != nil {
		return fdp, err
	}

	return fdp, nil
}

// GetFeeDenomParamIter returns an iterator for all fee denom feeDenomParam.
func (k *Keeper) GetFeeDenomParamIter(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)

	return sdk.KVStorePrefixIterator(store, types.KeyPrefixFeeDenomParam)
}

// GetFeeDenomParams returns all feemarket feeDenomParams.
func (k *Keeper) GetFeeDenomParams(ctx sdk.Context) ([]types.FeeDenomParam, error) {
	iter := k.GetFeeDenomParamIter(ctx)
	defer iter.Close()

	var fdps []types.FeeDenomParam

	for ; iter.Valid(); iter.Next() {
		fdp := types.FeeDenomParam{}
		if err := fdp.Unmarshal(iter.Value()); err != nil {
			return nil, err
		}

		fdps = append(fdps, fdp)
	}

	return fdps, nil
}

// SetFeeDenomParam sets the feeDenomParam for feeDenom.
func (k *Keeper) SetFeeDenomParam(ctx sdk.Context, fdp types.FeeDenomParam) error {
	store := ctx.KVStore(k.storeKey)

	key := types.GetKeyPrefixFeeDenomParam(fdp.FeeDenom)
	bz, err := fdp.Marshal()
	if err != nil {
		return err
	}

	store.Set(key, bz)

	return nil
}

// DeleteFeeDenomParam removes the feeDenomParam for feeDenom.
func (k *Keeper) DeleteFeeDenomParam(ctx sdk.Context, feeDenom string) error {
	store := ctx.KVStore(k.storeKey)

	key := types.GetKeyPrefixFeeDenomParam(feeDenom)
	store.Delete(key)

	return nil
}

// GetState returns the feemarket module's state.
func (k *Keeper) GetState(ctx sdk.Context) (types.State, error) {
	store := ctx.KVStore(k.storeKey)

	key := types.KeyState
	bz := store.Get(key)

	state := types.State{}
	if err := state.Unmarshal(bz); err != nil {
		return types.State{}, err
	}

	return state, nil
}

// SetState sets the feemarket module's state.
func (k *Keeper) SetState(ctx sdk.Context, state types.State) error {
	store := ctx.KVStore(k.storeKey)

	bz, err := state.Marshal()
	if err != nil {
		return err
	}

	store.Set(types.KeyState, bz)

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
