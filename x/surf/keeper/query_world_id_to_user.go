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

func (k Keeper) WorldIdToUserAll(ctx context.Context, req *types.QueryAllWorldIdToUserRequest) (*types.QueryAllWorldIdToUserResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var worldIdToUsers []types.WorldIdToUser

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	worldIdToUserStore := prefix.NewStore(store, types.KeyPrefix(types.WorldIdToUserKeyPrefix))

	pageRes, err := query.Paginate(worldIdToUserStore, req.Pagination, func(key []byte, value []byte) error {
		var worldIdToUser types.WorldIdToUser
		if err := k.cdc.Unmarshal(value, &worldIdToUser); err != nil {
			return err
		}

		worldIdToUsers = append(worldIdToUsers, worldIdToUser)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllWorldIdToUserResponse{WorldIdToUser: worldIdToUsers, Pagination: pageRes}, nil
}

func (k Keeper) WorldIdToUser(ctx context.Context, req *types.QueryGetWorldIdToUserRequest) (*types.QueryGetWorldIdToUserResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetWorldIdToUser(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetWorldIdToUserResponse{WorldIdToUser: val}, nil
}
