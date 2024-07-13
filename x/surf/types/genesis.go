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
		AccountToUserList: []AccountToUser{},
		WorldIdToUserList: []WorldIdToUser{},
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
	// Check for duplicated index in accountToUser
	accountToUserIndexMap := make(map[string]struct{})

	for _, elem := range gs.AccountToUserList {
		index := string(AccountToUserKey(elem.Index))
		if _, ok := accountToUserIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for accountToUser")
		}
		accountToUserIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in worldIdToUser
	worldIdToUserIndexMap := make(map[string]struct{})

	for _, elem := range gs.WorldIdToUserList {
		index := string(WorldIdToUserKey(elem.Index))
		if _, ok := worldIdToUserIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for worldIdToUser")
		}
		worldIdToUserIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
