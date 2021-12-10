package types

import (
	"fmt"
	host "github.com/cosmos/cosmos-sdk/x/ibc/core/24-host"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PortId: PortID,
		// this line is used by starport scaffolding # genesis/types/default
		MarketList:  []*Market{},
		BookverseList: []*Bookverse{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	if err := host.PortIdentifierValidator(gs.PortId); err != nil {
		return err
	}

	// this line is used by starport scaffolding # genesis/types/validate
	// Check for duplicated ID in market
	marketIdMap := make(map[uint64]bool)

	for _, elem := range gs.MarketList {
		if _, ok := marketIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for market")
		}
		marketIdMap[elem.Id] = true
	}
	// Check for duplicated ID in Bookverse
	BookverseIdMap := make(map[uint64]bool)

	for _, elem := range gs.BookverseList {
		if _, ok := BookverseIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for Bookverse")
		}
		BookverseIdMap[elem.Id] = true
	}

	return nil
}
