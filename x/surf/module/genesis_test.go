package surf_test

import (
	"testing"

	keepertest "surf/testutil/keeper"
	"surf/testutil/nullify"
	surf "surf/x/surf/module"
	"surf/x/surf/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		WorldIdtoUserList: []types.WorldIdtoUser{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		UserToWorldIdList: []types.UserToWorldId{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SurfKeeper(t)
	surf.InitGenesis(ctx, k, genesisState)
	got := surf.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	require.ElementsMatch(t, genesisState.WorldIdtoUserList, got.WorldIdtoUserList)
	require.ElementsMatch(t, genesisState.UserToWorldIdList, got.UserToWorldIdList)
	// this line is used by starport scaffolding # genesis/test/assert
}
