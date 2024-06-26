package keeper_test

import (
	"context"
	"strconv"
	"testing"

	keepertest "surf/testutil/keeper"
	"surf/testutil/nullify"
	"surf/x/surf/keeper"
	"surf/x/surf/types"

	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNUserToWorldId(keeper keeper.Keeper, ctx context.Context, n int) []types.UserToWorldId {
	items := make([]types.UserToWorldId, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetUserToWorldId(ctx, items[i])
	}
	return items
}

func TestUserToWorldIdGet(t *testing.T) {
	keeper, ctx := keepertest.SurfKeeper(t)
	items := createNUserToWorldId(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetUserToWorldId(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestUserToWorldIdRemove(t *testing.T) {
	keeper, ctx := keepertest.SurfKeeper(t)
	items := createNUserToWorldId(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveUserToWorldId(ctx,
			item.Index,
		)
		_, found := keeper.GetUserToWorldId(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestUserToWorldIdGetAll(t *testing.T) {
	keeper, ctx := keepertest.SurfKeeper(t)
	items := createNUserToWorldId(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllUserToWorldId(ctx)),
	)
}
