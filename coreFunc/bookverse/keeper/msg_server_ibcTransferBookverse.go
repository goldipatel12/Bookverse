package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/cosmos-sdkcoreFuncibc/core/02-client/types"
	"github.com/goldipatel12/marketplacecoreFuncBookverses/types"
)

func (k msgServer) SendIbcTransferBookverse(goCtx context.Context, msg *types.MsgSendIbcTransferBookverse) (*types.MsgSendIbcTransferBookverseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: logic before transmitting the packet

	// Construct the packet
	var packet types.IbcTransferBookversePacketData

	packet.SID = msg.SID
	packet.NftStorageCID = msg.NftStorageCID
	packet.Creator = msg.Sender

	// Transmit the packet
	err := k.TransmitIbcTransferBookversePacket(
		ctx,
		packet,
		msg.Port,
		msg.ChannelID,
		clienttypes.ZeroHeight(),
		msg.TimeoutTimestamp,
	)
	if err != nil {
		return nil, err
	}

	return &types.MsgSendIbcTransferBookverseResponse{}, nil
}
