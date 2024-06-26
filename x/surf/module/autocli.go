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
					RpcMethod: "WorldIdtoUserAll",
					Use:       "list-world-id-to-user",
					Short:     "List all worldIDToUser",
				},
				{
					RpcMethod:      "WorldIdtoUser",
					Use:            "show-world-id-to-user [id]",
					Short:          "Shows a worldIDToUser",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod: "UserToWorldIdAll",
					Use:       "list-user-to-world-id",
					Short:     "List all userToWorldID",
				},
				{
					RpcMethod:      "UserToWorldId",
					Use:            "show-user-to-world-id [id]",
					Short:          "Shows a userToWorldID",
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
