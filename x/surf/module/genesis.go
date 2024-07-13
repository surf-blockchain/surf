package surf

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"surf/x/surf/keeper"
	"surf/x/surf/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the accountToUser
	for _, elem := range genState.AccountToUserList {
		k.SetAccountToUser(ctx, elem)
	}
	// Set all the worldIdToUser
	for _, elem := range genState.WorldIdToUserList {
		k.SetWorldIdToUser(ctx, elem)
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
	genesis.AccountToUserList = k.GetAllAccountToUser(ctx)
	genesis.WorldIdToUserList = k.GetAllWorldIdToUser(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
