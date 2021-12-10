package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/goldipatel12/marketplacecoreFuncBookverses/types"
	"strconv"
)

// GetMarketCount get the total number of market
func (k Keeper) GetMarketCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MarketCountKey))
	byteKey := types.KeyPrefix(types.MarketCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseUint(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to iint64
		panic("cannot decode count")
	}

	return count
}

// SetMarketCount set the total number of market
func (k Keeper) SetMarketCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MarketCountKey))
	byteKey := types.KeyPrefix(types.MarketCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendMarket appends a market in the store with a new id and update the count
func (k Keeper) AppendMarket(
	ctx sdk.Context,
	market types.Market,
) uint64 {
	// Create the market
	count := k.GetMarketCount(ctx)

	// Set the ID of the appended value
	market.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MarketKey))
	appendedValue := k.cdc.MustMarshalBinaryBare(&market)
	store.Set(GetMarketIDBytes(market.Id), appendedValue)

	// Update market count
	k.SetMarketCount(ctx, count+1)

	return count
}

// SetMarket set a specific market in the store
func (k Keeper) SetMarket(ctx sdk.Context, market types.Market) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MarketKey))
	b := k.cdc.MustMarshalBinaryBare(&market)
	store.Set(GetMarketIDBytes(market.Id), b)
}

// GetMarket returns a market from its id
func (k Keeper) GetMarket(ctx sdk.Context, id uint64) types.Market {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MarketKey))
	var market types.Market
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetMarketIDBytes(id)), &market)
	return market
}

// HasMarket checks if the market exists in the store
func (k Keeper) HasMarket(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MarketKey))
	return store.Has(GetMarketIDBytes(id))
}

// GetMarketOwner returns the creator of the market
func (k Keeper) GetMarketOwner(ctx sdk.Context, id uint64) string {
	return k.GetMarket(ctx, id).Creator
}

// RemoveMarket removes a market from the store
func (k Keeper) RemoveMarket(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MarketKey))
	store.Delete(GetMarketIDBytes(id))
}

// GetAllMarket returns all market
func (k Keeper) GetAllMarket(ctx sdk.Context) (list []types.Market) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MarketKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Market
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetMarketIDBytes returns the byte representation of the ID
func GetMarketIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetMarketIDFromBytes returns ID in uint64 format from a byte array
func GetMarketIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
