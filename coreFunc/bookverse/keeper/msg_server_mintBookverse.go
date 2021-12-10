package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/goldipatel12/marketplace/x/Bookverses/types"
)

func (k msgServer) MintBookverse(goCtx context.Context, msg *types.MsgMintBookverse) (*types.MsgMintBookverseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: validation
	if len(msg.Sender) == 0 || len(msg.Name) == 0 || len(msg.Image) == 0 {
		return nil, sdkerrors.Wrap(types.ErrRequiredFields, "some fields are required")
	}

	// create new NFT stay with generated stay ID
	var Bookverse = types.Bookverse{
		Creator:     msg.Sender,
		Owner:       msg.Recipient,
		SID:         GenerateStaySID(),
		Name:        msg.Name,
		Description: msg.Description,
		Image:       msg.Image,
		TokenUri:    msg.TokenUri,
	}

	k.AppendBookverse(
		ctx,
		Bookverse,
	)

	return &types.MsgMintBookverseResponse{}, nil
}
