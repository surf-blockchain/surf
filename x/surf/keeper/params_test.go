package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "surf/testutil/keeper"
	"surf/x/surf/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.SurfKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
