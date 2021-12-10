package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgClaimBookverse{}

func NewMsgClaimBookverse(creator string, marketId string, buyer string) *MsgClaimBookverse {
	return &MsgClaimBookverse{
		Creator:  creator,
		MarketId: marketId,
		Buyer:    buyer,
	}
}

func (msg *MsgClaimBookverse) Route() string {
	return RouterKey
}

func (msg *MsgClaimBookverse) Type() string {
	return "ClaimBookverse"
}

func (msg *MsgClaimBookverse) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Buyer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgClaimBookverse) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgClaimBookverse) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Buyer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
