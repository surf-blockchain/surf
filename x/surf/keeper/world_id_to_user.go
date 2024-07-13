package keeper

import (
	"context"

	"surf/x/surf/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetWorldIdToUser set a specific worldIdToUser in the store from its index
func (k Keeper) SetWorldIdToUser(ctx context.Context, worldIdToUser types.WorldIdToUser) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.WorldIdToUserKeyPrefix))
	b := k.cdc.MustMarshal(&worldIdToUser)
	store.Set(types.WorldIdToUserKey(
		worldIdToUser.Index,
	), b)
}

// GetWorldIdToUser returns a worldIdToUser from its index
func (k Keeper) GetWorldIdToUser(
	ctx context.Context,
	index string,

) (val types.WorldIdToUser, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.WorldIdToUserKeyPrefix))

	b := store.Get(types.WorldIdToUserKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveWorldIdToUser removes a worldIdToUser from the store
func (k Keeper) RemoveWorldIdToUser(
	ctx context.Context,
	index string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.WorldIdToUserKeyPrefix))
	store.Delete(types.WorldIdToUserKey(
		index,
	))
}

// GetAllWorldIdToUser returns all worldIdToUser
func (k Keeper) GetAllWorldIdToUser(ctx context.Context) (list []types.WorldIdToUser) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.WorldIdToUserKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.WorldIdToUser
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
