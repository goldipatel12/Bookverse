package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/goldipatel12/marketplacecoreFuncBookverses/types"
	"github.com/stretchr/testify/assert"
)

func createNBookverse(keeper *Keeper, ctx sdk.Context, n int) []types.Bookverse {
	items := make([]types.Bookverse, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Id = keeper.AppendBookverse(ctx, items[i])
	}
	return items
}

func TestBookverseGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNBookverse(keeper, ctx, 10)
	for _, item := range items {
		assert.Equal(t, item, keeper.GetBookverse(ctx, item.Id))
	}
}

func TestBookverseExist(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNBookverse(keeper, ctx, 10)
	for _, item := range items {
		assert.True(t, keeper.HasBookverse(ctx, item.Id))
	}
}

func TestBookverseRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNBookverse(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveBookverse(ctx, item.Id)
		assert.False(t, keeper.HasBookverse(ctx, item.Id))
	}
}

func TestBookverseGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNBookverse(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllBookverse(ctx))
}

func TestBookverseCount(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNBookverse(keeper, ctx, 10)
	count := uint64(len(items))
	assert.Equal(t, count, keeper.GetBookverseCount(ctx))
}
