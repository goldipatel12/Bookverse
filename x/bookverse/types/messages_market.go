package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateMarket{}

func NewMsgCreateMarket(creator string, sID string, price int32, status string, seller string, onAuction bool, offers string, expired int32) *MsgCreateMarket {
	return &MsgCreateMarket{
		Creator:   creator,
		SID:       sID,
		Price:     price,
		Status:    status,
		Seller:    seller,
		OnAuction: onAuction,
		Offers:    offers,
		Expired:   expired,
	}
}

func (msg *MsgCreateMarket) Route() string {
	return RouterKey
}

func (msg *MsgCreateMarket) Type() string {
	return "CreateMarket"
}

func (msg *MsgCreateMarket) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateMarket) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateMarket) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateMarket{}

func NewMsgUpdateMarket(creator string, id uint64, sID string, price int32, status string, seller string, onAuction bool, offers string, expired int32) *MsgUpdateMarket {
	return &MsgUpdateMarket{
		Id:        id,
		Creator:   creator,
		SID:       sID,
		Price:     price,
		Status:    status,
		Seller:    seller,
		OnAuction: onAuction,
		Offers:    offers,
		Expired:   expired,
	}
}

func (msg *MsgUpdateMarket) Route() string {
	return RouterKey
}

func (msg *MsgUpdateMarket) Type() string {
	return "UpdateMarket"
}

func (msg *MsgUpdateMarket) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateMarket) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateMarket) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteMarket{}

func NewMsgDeleteMarket(creator string, id uint64) *MsgDeleteMarket {
	return &MsgDeleteMarket{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteMarket) Route() string {
	return RouterKey
}

func (msg *MsgDeleteMarket) Type() string {
	return "DeleteMarket"
}

func (msg *MsgDeleteMarket) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteMarket) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteMarket) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
