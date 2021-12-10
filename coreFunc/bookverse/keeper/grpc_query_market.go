package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/goldipatel12/marketplace/x/Bookverses/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) MarketAll(c context.Context, req *types.QueryAllMarketRequest) (*types.QueryAllMarketResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var markets []*types.Market
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	marketStore := prefix.NewStore(store, types.KeyPrefix(types.MarketKey))

	pageRes, err := query.Paginate(marketStore, req.Pagination, func(key []byte, value []byte) error {
		var market types.Market
		if err := k.cdc.UnmarshalBinaryBare(value, &market); err != nil {
			return err
		}

		markets = append(markets, &market)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllMarketResponse{Market: markets, Pagination: pageRes}, nil
}

func (k Keeper) Market(c context.Context, req *types.QueryGetMarketRequest) (*types.QueryGetMarketResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var market types.Market
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasMarket(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MarketKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetMarketIDBytes(req.Id)), &market)

	return &types.QueryGetMarketResponse{Market: &market}, nil
}
