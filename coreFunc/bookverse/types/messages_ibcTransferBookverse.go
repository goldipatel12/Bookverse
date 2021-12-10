package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSendIbcTransferBookverse{}

func NewMsgSendIbcTransferBookverse(
	sender string,
	port string,
	channelID string,
	timeoutTimestamp uint64,
	sID string,
	nftStorageCID string,
) *MsgSendIbcTransferBookverse {
	return &MsgSendIbcTransferBookverse{
		Sender:           sender,
		Port:             port,
		ChannelID:        channelID,
		TimeoutTimestamp: timeoutTimestamp,
		SID:              sID,
		NftStorageCID:    nftStorageCID,
	}
}

func (msg *MsgSendIbcTransferBookverse) Route() string {
	return RouterKey
}

func (msg *MsgSendIbcTransferBookverse) Type() string {
	return "SendIbcTransferBookverse"
}

func (msg *MsgSendIbcTransferBookverse) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgSendIbcTransferBookverse) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendIbcTransferBookverse) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}
