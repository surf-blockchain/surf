package keeper_test

import (
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "surf/testutil/keeper"
	"surf/testutil/nullify"
	"surf/x/surf/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestWorldIdtoUserQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.SurfKeeper(t)
	msgs := createNWorldIdtoUser(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetWorldIdtoUserRequest
		response *types.QueryGetWorldIdtoUserResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetWorldIdtoUserRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetWorldIdtoUserResponse{WorldIdtoUser: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetWorldIdtoUserRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetWorldIdtoUserResponse{WorldIdtoUser: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetWorldIdtoUserRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.WorldIdtoUser(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestWorldIdtoUserQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.SurfKeeper(t)
	msgs := createNWorldIdtoUser(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllWorldIdtoUserRequest {
		return &types.QueryAllWorldIdtoUserRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.WorldIdtoUserAll(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.WorldIdtoUser), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.WorldIdtoUser),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.WorldIdtoUserAll(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.WorldIdtoUser), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.WorldIdtoUser),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.WorldIdtoUserAll(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.WorldIdtoUser),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.WorldIdtoUserAll(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
