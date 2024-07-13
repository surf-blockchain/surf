package keeper

import (
	"context"

	"surf/x/surf/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AccountToUserAll(ctx context.Context, req *types.QueryAllAccountToUserRequest) (*types.QueryAllAccountToUserResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var accountToUsers []types.AccountToUser

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	accountToUserStore := prefix.NewStore(store, types.KeyPrefix(types.AccountToUserKeyPrefix))

	pageRes, err := query.Paginate(accountToUserStore, req.Pagination, func(key []byte, value []byte) error {
		var accountToUser types.AccountToUser
		if err := k.cdc.Unmarshal(value, &accountToUser); err != nil {
			return err
		}

		accountToUsers = append(accountToUsers, accountToUser)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAccountToUserResponse{AccountToUser: accountToUsers, Pagination: pageRes}, nil
}

func (k Keeper) AccountToUser(ctx context.Context, req *types.QueryGetAccountToUserRequest) (*types.QueryGetAccountToUserResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetAccountToUser(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetAccountToUserResponse{AccountToUser: val}, nil
}
