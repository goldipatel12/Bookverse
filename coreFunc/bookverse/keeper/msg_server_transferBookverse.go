package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/goldipatel12/marketplace/x/Bookverses/types"
)

func (k msgServer) TransferBookverse(goCtx context.Context, msg *types.MsgTransferBookverse) (*types.MsgTransferBookverseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check validations
	stayID := GetBookverseIDFromSID(msg.SID)
	if !k.HasBookverse(ctx, stayID) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", stayID))
	}
	if len(msg.Sender) == 0 || len(msg.Recipient) == 0 {
		return nil, sdkerrors.Wrap(types.ErrRequiredFields, "Sender and Recipient are required")
	}

	_, err := sdk.AccAddressFromBech32(msg.Recipient)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid Recipient address (%s)", err)
	}

	// cannot send yourself
	if msg.Sender == msg.Recipient {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("cannot send yourself"))
	}

	// update stay owner
	stay := k.Keeper.GetBookverse(ctx, stayID)
	stay.Owner = msg.Recipient
	k.SetBookverse(ctx, stay)

	return &types.MsgTransferBookverseResponse{}, nil
}
