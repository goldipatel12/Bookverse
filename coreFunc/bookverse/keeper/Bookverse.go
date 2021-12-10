package keeper

import (
	"encoding/binary"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/google/uuid"
	"github.com/goldipatel12/marketplacecoreFuncBookverses/types"
)

// GetBookverseCount get the total number of Bookverse
func (k Keeper) GetBookverseCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BookverseCountKey))
	byteKey := types.KeyPrefix(types.BookverseCountKey)
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

// SetBookverseCount set the total number of Bookverse
func (k Keeper) SetBookverseCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BookverseCountKey))
	byteKey := types.KeyPrefix(types.BookverseCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendBookverse appends a Bookverse in the store with a new id and update the count
func (k Keeper) AppendBookverse(
	ctx sdk.Context,
	Bookverse types.Bookverse,
) uint64 {
	// Create the Bookverse
	count := k.GetBookverseCount(ctx)

	// Set the ID of the appended value
	Bookverse.Id = GetBookverseIDFromSID(Bookverse.SID)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BookverseKey))
	appendedValue := k.cdc.MustMarshalBinaryBare(&Bookverse)
	store.Set(GetBookverseIDBytes(Bookverse.Id), appendedValue)

	// Update Bookverse count
	k.SetBookverseCount(ctx, count+1)

	return count
}

// SetBookverse set a specific Bookverse in the store
func (k Keeper) SetBookverse(ctx sdk.Context, Bookverse types.Bookverse) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BookverseKey))
	b := k.cdc.MustMarshalBinaryBare(&Bookverse)
	store.Set(GetBookverseIDBytes(Bookverse.Id), b)
}

// GetBookverse returns a Bookverse from its id
func (k Keeper) GetBookverse(ctx sdk.Context, id uint64) types.Bookverse {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BookverseKey))
	var Bookverse types.Bookverse
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetBookverseIDBytes(id)), &Bookverse)
	return Bookverse
}

// HasBookverse checks if the Bookverse exists in the store
func (k Keeper) HasBookverse(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BookverseKey))
	return store.Has(GetBookverseIDBytes(id))
}

// GetBookverseOwner returns the creator of the Bookverse
func (k Keeper) GetBookverseOwner(ctx sdk.Context, id uint64) string {
	return k.GetBookverse(ctx, id).Owner
}

// RemoveBookverse removes a Bookverse from the store
func (k Keeper) RemoveBookverse(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BookverseKey))
	store.Delete(GetBookverseIDBytes(id))
}

// GetAllBookverse returns all Bookverse
func (k Keeper) GetAllBookverse(ctx sdk.Context) (list []types.Bookverse) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BookverseKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Bookverse
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetBookverseIDBytes returns the byte representation of the ID
func GetBookverseIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetBookverseIDFromBytes returns ID in uint64 format from a byte array
func GetBookverseIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}

func GetBookverseIDFromSID(sID string) uint64 {
	idNum := strings.TrimPrefix(sID, types.BookverseSIDPrefix)
	stayID, err := strconv.ParseUint(idNum, 10, 64)
	if err != nil {
		panic(err)
	}

	return stayID
}

func GenerateStaySID() string {
	newID := uint64(uuid.New().ID())
	sID := types.BookverseSIDPrefix + strconv.FormatUint(newID, 10)
	return sID
}
