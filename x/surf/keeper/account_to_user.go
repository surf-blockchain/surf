package keeper

import (
	"context"

	"surf/x/surf/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetAccountToUser set a specific accountToUser in the store from its index
func (k Keeper) SetAccountToUser(ctx context.Context, accountToUser types.AccountToUser) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AccountToUserKeyPrefix))
	b := k.cdc.MustMarshal(&accountToUser)
	store.Set(types.AccountToUserKey(
		accountToUser.Index,
	), b)
}

// GetAccountToUser returns a accountToUser from its index
func (k Keeper) GetAccountToUser(
	ctx context.Context,
	index string,

) (val types.AccountToUser, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AccountToUserKeyPrefix))

	b := store.Get(types.AccountToUserKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAccountToUser removes a accountToUser from the store
func (k Keeper) RemoveAccountToUser(
	ctx context.Context,
	index string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AccountToUserKeyPrefix))
	store.Delete(types.AccountToUserKey(
		index,
	))
}

// GetAllAccountToUser returns all accountToUser
func (k Keeper) GetAllAccountToUser(ctx context.Context) (list []types.AccountToUser) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AccountToUserKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.AccountToUser
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
