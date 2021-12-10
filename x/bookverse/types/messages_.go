package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateBookverse{}

func NewMsgCreateBookverse(creator string, sID string, owner string, name string, description string, image string, tokenUri string) *MsgCreateBookverse {
	return &MsgCreateBookverse{
		Creator:     creator,
		SID:         sID,
		Owner:       owner,
		Name:        name,
		Description: description,
		Image:       image,
		TokenUri:    tokenUri,
	}
}

func (msg *MsgCreateBookverse) Route() string {
	return RouterKey
}

func (msg *MsgCreateBookverse) Type() string {
	return "CreateBookverse"
}

func (msg *MsgCreateBookverse) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateBookverse) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateBookverse) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateBookverse{}

func NewMsgUpdateBookverse(creator string, id uint64, sID string, owner string, name string, description string, image string, tokenUri string) *MsgUpdateBookverse {
	return &MsgUpdateBookverse{
		Id:          id,
		Creator:     creator,
		SID:         sID,
		Owner:       owner,
		Name:        name,
		Description: description,
		Image:       image,
		TokenUri:    tokenUri,
	}
}

func (msg *MsgUpdateBookverse) Route() string {
	return RouterKey
}

func (msg *MsgUpdateBookverse) Type() string {
	return "UpdateBookverse"
}

func (msg *MsgUpdateBookverse) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateBookverse) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateBookverse) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteBookverse{}

func NewMsgDeleteBookverse(creator string, id uint64) *MsgDeleteBookverse {
	return &MsgDeleteBookverse{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteBookverse) Route() string {
	return RouterKey
}

func (msg *MsgDeleteBookverse) Type() string {
	return "DeleteBookverse"
}

func (msg *MsgDeleteBookverse) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteBookverse) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteBookverse) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
