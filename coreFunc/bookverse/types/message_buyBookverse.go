package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgBuyBookverse{}

func NewMsgBuyBookverse(creator string, marketId string, buyer string) *MsgBuyBookverse {
	return &MsgBuyBookverse{
		Creator:  creator,
		MarketId: marketId,
		Buyer:    buyer,
	}
}

func (msg *MsgBuyBookverse) Route() string {
	return RouterKey
}

func (msg *MsgBuyBookverse) Type() string {
	return "BuyBookverse"
}

func (msg *MsgBuyBookverse) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Buyer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgBuyBookverse) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBuyBookverse) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Buyer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
