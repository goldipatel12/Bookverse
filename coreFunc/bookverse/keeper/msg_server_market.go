package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/goldipatel12/marketplace/x/Bookverses/types"
)

func (k msgServer) CreateMarket(goCtx context.Context, msg *types.MsgCreateMarket) (*types.MsgCreateMarketResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var market = types.Market{
		Creator:   msg.Creator,
		SID:       msg.SID,
		Price:     msg.Price,
		Status:    msg.Status,
		Seller:    msg.Seller,
		OnAuction: msg.OnAuction,
		Offers:    msg.Offers,
		Expired:   msg.Expired,
	}

	id := k.AppendMarket(
		ctx,
		market,
	)

	return &types.MsgCreateMarketResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateMarket(goCtx context.Context, msg *types.MsgUpdateMarket) (*types.MsgUpdateMarketResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var market = types.Market{
		Creator:   msg.Creator,
		Id:        msg.Id,
		SID:       msg.SID,
		Price:     msg.Price,
		Status:    msg.Status,
		Seller:    msg.Seller,
		OnAuction: msg.OnAuction,
		Offers:    msg.Offers,
		Expired:   msg.Expired,
	}

	// Checks that the element exists
	if !k.HasMarket(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetMarketOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetMarket(ctx, market)

	return &types.MsgUpdateMarketResponse{}, nil
}

func (k msgServer) DeleteMarket(goCtx context.Context, msg *types.MsgDeleteMarket) (*types.MsgDeleteMarketResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasMarket(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != k.GetMarketOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveMarket(ctx, msg.Id)

	return &types.MsgDeleteMarketResponse{}, nil
}
