package types

import (
	"fmt"

	host "github.com/cosmos/ibc-go/v8/modules/core/24-host"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PortId:            PortID,
		WorldIdtoUserList: []WorldIdtoUser{},
		UserToWorldIdList: []UserToWorldId{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	if err := host.PortIdentifierValidator(gs.PortId); err != nil {
		return err
	}
	// Check for duplicated index in worldIdtoUser
	worldIdtoUserIndexMap := make(map[string]struct{})

	for _, elem := range gs.WorldIdtoUserList {
		index := string(WorldIdtoUserKey(elem.Index))
		if _, ok := worldIdtoUserIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for worldIdtoUser")
		}
		worldIdtoUserIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in userToWorldId
	userToWorldIdIndexMap := make(map[string]struct{})

	for _, elem := range gs.UserToWorldIdList {
		index := string(UserToWorldIdKey(elem.Index))
		if _, ok := userToWorldIdIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for userToWorldId")
		}
		userToWorldIdIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
