package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSellBookverse{}

func NewMsgSellBookverse(creator string, sID string, seller string, price int32) *MsgSellBookverse {
	return &MsgSellBookverse{
		Creator:   creator,
		SID:       sID,
		Seller:    seller,
		Price:     price,
	}
}

func (msg *MsgSellBookverse) Route() string {
	return RouterKey
}

func (msg *MsgSellBookverse) Type() string {
	return "SellBookverse"
}

func (msg *MsgSellBookverse) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Seller)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSellBookverse) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSellBookverse) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Seller)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
