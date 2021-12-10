package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgTransferBookverse{}

func NewMsgTransferBookverse(creator string, sender string, recipient string, sID string, denom string) *MsgTransferBookverse {
	return &MsgTransferBookverse{
		Creator:   creator,
		Sender:    sender,
		Recipient: recipient,
		SID:       sID,
		Denom:     denom,
	}
}

func (msg *MsgTransferBookverse) Route() string {
	return RouterKey
}

func (msg *MsgTransferBookverse) Type() string {
	return "TransferBookverse"
}

func (msg *MsgTransferBookverse) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgTransferBookverse) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgTransferBookverse) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
