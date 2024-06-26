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

func createNWorldIdtoUser(keeper keeper.Keeper, ctx context.Context, n int) []types.WorldIdtoUser {
	items := make([]types.WorldIdtoUser, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetWorldIdtoUser(ctx, items[i])
	}
	return items
}

func TestWorldIdtoUserGet(t *testing.T) {
	keeper, ctx := keepertest.SurfKeeper(t)
	items := createNWorldIdtoUser(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetWorldIdtoUser(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestWorldIdtoUserRemove(t *testing.T) {
	keeper, ctx := keepertest.SurfKeeper(t)
	items := createNWorldIdtoUser(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveWorldIdtoUser(ctx,
			item.Index,
		)
		_, found := keeper.GetWorldIdtoUser(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestWorldIdtoUserGetAll(t *testing.T) {
	keeper, ctx := keepertest.SurfKeeper(t)
	items := createNWorldIdtoUser(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllWorldIdtoUser(ctx)),
	)
}
