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

func (k Keeper) UserToWorldIdAll(ctx context.Context, req *types.QueryAllUserToWorldIdRequest) (*types.QueryAllUserToWorldIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var userToWorldIds []types.UserToWorldId

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	userToWorldIdStore := prefix.NewStore(store, types.KeyPrefix(types.UserToWorldIdKeyPrefix))

	pageRes, err := query.Paginate(userToWorldIdStore, req.Pagination, func(key []byte, value []byte) error {
		var userToWorldId types.UserToWorldId
		if err := k.cdc.Unmarshal(value, &userToWorldId); err != nil {
			return err
		}

		userToWorldIds = append(userToWorldIds, userToWorldId)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllUserToWorldIdResponse{UserToWorldId: userToWorldIds, Pagination: pageRes}, nil
}

func (k Keeper) UserToWorldId(ctx context.Context, req *types.QueryGetUserToWorldIdRequest) (*types.QueryGetUserToWorldIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetUserToWorldId(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetUserToWorldIdResponse{UserToWorldId: val}, nil
}
