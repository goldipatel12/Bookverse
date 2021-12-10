package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/goldipatel12/marketplacecoreFuncBookverses/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) BookverseAll(c context.Context, req *types.QueryAllBookverseRequest) (*types.QueryAllBookverseResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var Bookverses []*types.Bookverse
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	BookverseStore := prefix.NewStore(store, types.KeyPrefix(types.BookverseKey))

	pageRes, err := query.Paginate(BookverseStore, req.Pagination, func(key []byte, value []byte) error {
		var Bookverse types.Bookverse
		if err := k.cdc.UnmarshalBinaryBare(value, &Bookverse); err != nil {
			return err
		}

		Bookverses = append(Bookverses, &Bookverse)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllBookverseResponse{Bookverse: Bookverses, Pagination: pageRes}, nil
}

func (k Keeper) Bookverse(c context.Context, req *types.QueryGetBookverseRequest) (*types.QueryGetBookverseResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var Bookverse types.Bookverse
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasBookverse(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BookverseKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetBookverseIDBytes(req.Id)), &Bookverse)

	return &types.QueryGetBookverseResponse{Bookverse: &Bookverse}, nil
}
