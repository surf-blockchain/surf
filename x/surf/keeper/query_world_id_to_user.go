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

func (k Keeper) WorldIdtoUserAll(ctx context.Context, req *types.QueryAllWorldIdtoUserRequest) (*types.QueryAllWorldIdtoUserResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var worldIdtoUsers []types.WorldIdtoUser

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	worldIdtoUserStore := prefix.NewStore(store, types.KeyPrefix(types.WorldIdtoUserKeyPrefix))

	pageRes, err := query.Paginate(worldIdtoUserStore, req.Pagination, func(key []byte, value []byte) error {
		var worldIdtoUser types.WorldIdtoUser
		if err := k.cdc.Unmarshal(value, &worldIdtoUser); err != nil {
			return err
		}

		worldIdtoUsers = append(worldIdtoUsers, worldIdtoUser)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllWorldIdtoUserResponse{WorldIdtoUser: worldIdtoUsers, Pagination: pageRes}, nil
}

func (k Keeper) WorldIdtoUser(ctx context.Context, req *types.QueryGetWorldIdtoUserRequest) (*types.QueryGetWorldIdtoUserResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetWorldIdtoUser(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetWorldIdtoUserResponse{WorldIdtoUser: val}, nil
}
