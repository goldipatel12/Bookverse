package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/goldipatel12/marketplace/x/Bookverses/types"
)

func (k msgServer) CreateBookverse(goCtx context.Context, msg *types.MsgCreateBookverse) (*types.MsgCreateBookverseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var Bookverse = types.Bookverse{
		Creator:     msg.Creator,
		SID:         msg.SID,
		Owner:       msg.Owner,
		Name:        msg.Name,
		Description: msg.Description,
		Image:       msg.Image,
		TokenUri:    msg.TokenUri,
	}

	id := k.AppendBookverse(
		ctx,
		Bookverse,
	)

	return &types.MsgCreateBookverseResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateBookverse(goCtx context.Context, msg *types.MsgUpdateBookverse) (*types.MsgUpdateBookverseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var Bookverse = types.Bookverse{
		Creator:     msg.Creator,
		Id:          msg.Id,
		SID:         msg.SID,
		Owner:       msg.Owner,
		Name:        msg.Name,
		Description: msg.Description,
		Image:       msg.Image,
		TokenUri:    msg.TokenUri,
	}

	// Checks that the element exists
	if !k.HasBookverse(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetBookverseOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetBookverse(ctx, Bookverse)

	return &types.MsgUpdateBookverseResponse{}, nil
}

func (k msgServer) DeleteBookverse(goCtx context.Context, msg *types.MsgDeleteBookverse) (*types.MsgDeleteBookverseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasBookverse(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != k.GetBookverseOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveBookverse(ctx, msg.Id)

	return &types.MsgDeleteBookverseResponse{}, nil
}
