package surf

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "surf/api/surf/surf"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "AccountToUserAll",
					Use:       "list-account-to-user",
					Short:     "List all accountToUser",
				},
				{
					RpcMethod:      "AccountToUser",
					Use:            "show-account-to-user [id]",
					Short:          "Shows a accountToUser",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod: "WorldIdToUserAll",
					Use:       "list-world-id-to-user",
					Short:     "List all worldIdToUser",
				},
				{
					RpcMethod:      "WorldIdToUser",
					Use:            "show-world-id-to-user [id]",
					Short:          "Shows a worldIdToUser",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
