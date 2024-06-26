package surf

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"surf/x/surf/keeper"
	"surf/x/surf/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the worldIdtoUser
	for _, elem := range genState.WorldIdtoUserList {
		k.SetWorldIdtoUser(ctx, elem)
	}
	// Set all the userToWorldId
	for _, elem := range genState.UserToWorldIdList {
		k.SetUserToWorldId(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetPort(ctx, genState.PortId)
	// Only try to bind to port if it is not already bound, since we may already own
	// port capability from capability InitGenesis
	if k.ShouldBound(ctx, genState.PortId) {
		// module binds to the port on InitChain
		// and claims the returned capability
		err := k.BindPort(ctx, genState.PortId)
		if err != nil {
			panic("could not claim port capability: " + err.Error())
		}
	}
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.PortId = k.GetPort(ctx)
	genesis.WorldIdtoUserList = k.GetAllWorldIdtoUser(ctx)
	genesis.UserToWorldIdList = k.GetAllUserToWorldId(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
