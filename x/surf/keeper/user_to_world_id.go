package keeper

import (
	"context"

	"surf/x/surf/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetUserToWorldId set a specific userToWorldId in the store from its index
func (k Keeper) SetUserToWorldId(ctx context.Context, userToWorldId types.UserToWorldId) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.UserToWorldIdKeyPrefix))
	b := k.cdc.MustMarshal(&userToWorldId)
	store.Set(types.UserToWorldIdKey(
		userToWorldId.Index,
	), b)
}

// GetUserToWorldId returns a userToWorldId from its index
func (k Keeper) GetUserToWorldId(
	ctx context.Context,
	index string,

) (val types.UserToWorldId, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.UserToWorldIdKeyPrefix))

	b := store.Get(types.UserToWorldIdKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveUserToWorldId removes a userToWorldId from the store
func (k Keeper) RemoveUserToWorldId(
	ctx context.Context,
	index string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.UserToWorldIdKeyPrefix))
	store.Delete(types.UserToWorldIdKey(
		index,
	))
}

// GetAllUserToWorldId returns all userToWorldId
func (k Keeper) GetAllUserToWorldId(ctx context.Context) (list []types.UserToWorldId) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.UserToWorldIdKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.UserToWorldId
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
