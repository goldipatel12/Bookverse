package Bookverses

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/goldipatel12/marketplace/x/Bookverses/keeper"
	"github.com/goldipatel12/marketplace/x/Bookverses/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	// Set all the market
	for _, elem := range genState.MarketList {
		k.SetMarket(ctx, *elem)
	}

	// Set market count
	k.SetMarketCount(ctx, genState.MarketCount)

	// Set all the Bookverse
	for _, elem := range genState.BookverseList {
		k.SetBookverse(ctx, *elem)
	}

	// Set Bookverse count
	k.SetBookverseCount(ctx, genState.BookverseCount)

	k.SetPort(ctx, genState.PortId)
	// Only try to bind to port if it is not already bound, since we may already own
	// port capability from capability InitGenesis
	if !k.IsBound(ctx, genState.PortId) {
		// module binds to the port on InitChain
		// and claims the returned capability
		err := k.BindPort(ctx, genState.PortId)
		if err != nil {
			panic("could not claim port capability: " + err.Error())
		}
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export
	// Get all market
	marketList := k.GetAllMarket(ctx)
	for _, elem := range marketList {
		elem := elem
		genesis.MarketList = append(genesis.MarketList, &elem)
	}

	// Set the current count
	genesis.MarketCount = k.GetMarketCount(ctx)

	// Get all Bookverse
	BookverseList := k.GetAllBookverse(ctx)
	for _, elem := range BookverseList {
		elem := elem
		genesis.BookverseList = append(genesis.BookverseList, &elem)
	}

	// Set the current count
	genesis.BookverseCount = k.GetBookverseCount(ctx)

	genesis.PortId = k.GetPort(ctx)

	return genesis
}
