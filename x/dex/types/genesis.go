package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		TickMapList: []TickMap{},
		PairMapList: []PairMap{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in tickMap
	tickMapIndexMap := make(map[string]struct{})

	for _, elem := range gs.TickMapList {
		index := string(TickMapKey(elem.TickIndex))
		if _, ok := tickMapIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for tickMap")
		}
		tickMapIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in pairMap
	pairMapIndexMap := make(map[string]struct{})

	for _, elem := range gs.PairMapList {
		index := string(PairMapKey(elem.PairId))
		if _, ok := pairMapIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for pairMap")
		}
		pairMapIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
