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

func createNWorldIdToUser(keeper keeper.Keeper, ctx context.Context, n int) []types.WorldIdToUser {
	items := make([]types.WorldIdToUser, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetWorldIdToUser(ctx, items[i])
	}
	return items
}

func TestWorldIdToUserGet(t *testing.T) {
	keeper, ctx := keepertest.SurfKeeper(t)
	items := createNWorldIdToUser(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetWorldIdToUser(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestWorldIdToUserRemove(t *testing.T) {
	keeper, ctx := keepertest.SurfKeeper(t)
	items := createNWorldIdToUser(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveWorldIdToUser(ctx,
			item.Index,
		)
		_, found := keeper.GetWorldIdToUser(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestWorldIdToUserGetAll(t *testing.T) {
	keeper, ctx := keepertest.SurfKeeper(t)
	items := createNWorldIdToUser(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllWorldIdToUser(ctx)),
	)
}
