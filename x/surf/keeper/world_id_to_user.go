package keeper

import (
	"context"

	"surf/x/surf/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetWorldIdtoUser set a specific worldIdtoUser in the store from its index
func (k Keeper) SetWorldIdtoUser(ctx context.Context, worldIdtoUser types.WorldIdtoUser) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.WorldIdtoUserKeyPrefix))
	b := k.cdc.MustMarshal(&worldIdtoUser)
	store.Set(types.WorldIdtoUserKey(
		worldIdtoUser.Index,
	), b)
}

// GetWorldIdtoUser returns a worldIdtoUser from its index
func (k Keeper) GetWorldIdtoUser(
	ctx context.Context,
	index string,

) (val types.WorldIdtoUser, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.WorldIdtoUserKeyPrefix))

	b := store.Get(types.WorldIdtoUserKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveWorldIdtoUser removes a worldIdtoUser from the store
func (k Keeper) RemoveWorldIdtoUser(
	ctx context.Context,
	index string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.WorldIdtoUserKeyPrefix))
	store.Delete(types.WorldIdtoUserKey(
		index,
	))
}

// GetAllWorldIdtoUser returns all worldIdtoUser
func (k Keeper) GetAllWorldIdtoUser(ctx context.Context) (list []types.WorldIdtoUser) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.WorldIdtoUserKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.WorldIdtoUser
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
