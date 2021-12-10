package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/goldipatel12/marketplacecoreFuncBookverses/types"
	"github.com/stretchr/testify/assert"
)

func createNMarket(keeper *Keeper, ctx sdk.Context, n int) []types.Market {
	items := make([]types.Market, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Id = keeper.AppendMarket(ctx, items[i])
	}
	return items
}

func TestMarketGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNMarket(keeper, ctx, 10)
	for _, item := range items {
		assert.Equal(t, item, keeper.GetMarket(ctx, item.Id))
	}
}

func TestMarketExist(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNMarket(keeper, ctx, 10)
	for _, item := range items {
		assert.True(t, keeper.HasMarket(ctx, item.Id))
	}
}

func TestMarketRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNMarket(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveMarket(ctx, item.Id)
		assert.False(t, keeper.HasMarket(ctx, item.Id))
	}
}

func TestMarketGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNMarket(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllMarket(ctx))
}

func TestMarketCount(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNMarket(keeper, ctx, 10)
	count := uint64(len(items))
	assert.Equal(t, count, keeper.GetMarketCount(ctx))
}
