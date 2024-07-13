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

func createNAccountToUser(keeper keeper.Keeper, ctx context.Context, n int) []types.AccountToUser {
	items := make([]types.AccountToUser, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetAccountToUser(ctx, items[i])
	}
	return items
}

func TestAccountToUserGet(t *testing.T) {
	keeper, ctx := keepertest.SurfKeeper(t)
	items := createNAccountToUser(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetAccountToUser(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestAccountToUserRemove(t *testing.T) {
	keeper, ctx := keepertest.SurfKeeper(t)
	items := createNAccountToUser(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveAccountToUser(ctx,
			item.Index,
		)
		_, found := keeper.GetAccountToUser(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestAccountToUserGetAll(t *testing.T) {
	keeper, ctx := keepertest.SurfKeeper(t)
	items := createNAccountToUser(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllAccountToUser(ctx)),
	)
}
