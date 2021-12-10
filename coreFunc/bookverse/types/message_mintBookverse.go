package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgMintBookverse{}

func NewMsgMintBookverse(creator string, sender string, recipient string, sID string, denom string, name string, description string, image string, tokenUri string) *MsgMintBookverse {
	return &MsgMintBookverse{
		Creator:     creator,
		Sender:      sender,
		Recipient:   recipient,
		SID:         sID,
		Denom:       denom,
		Name:        name,
		Description: description,
		Image:       image,
		TokenUri:    tokenUri,
	}
}

func (msg *MsgMintBookverse) Route() string {
	return RouterKey
}

func (msg *MsgMintBookverse) Type() string {
	return "MintBookverse"
}

func (msg *MsgMintBookverse) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgMintBookverse) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMintBookverse) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
