package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgSendIbcTransferBookverse{}, "Bookverses/SendIbcTransferBookverse", nil)

	cdc.RegisterConcrete(&MsgUnsellBookverse{}, "Bookverses/UnsellBookverse", nil)

	cdc.RegisterConcrete(&MsgClaimBookverse{}, "Bookverses/ClaimBookverse", nil)

	cdc.RegisterConcrete(&MsgBidBookverse{}, "Bookverses/BidBookverse", nil)

	cdc.RegisterConcrete(&MsgBuyBookverse{}, "Bookverses/BuyBookverse", nil)

	cdc.RegisterConcrete(&MsgSellBookverse{}, "Bookverses/SellBookverse", nil)

	cdc.RegisterConcrete(&MsgCreateMarket{}, "Bookverses/CreateMarket", nil)
	cdc.RegisterConcrete(&MsgUpdateMarket{}, "Bookverses/UpdateMarket", nil)
	cdc.RegisterConcrete(&MsgDeleteMarket{}, "Bookverses/DeleteMarket", nil)

	cdc.RegisterConcrete(&MsgTransferBookverse{}, "Bookverses/TransferBookverse", nil)

	cdc.RegisterConcrete(&MsgBurnBookverse{}, "Bookverses/BurnBookverse", nil)

	cdc.RegisterConcrete(&MsgMintBookverse{}, "Bookverses/MintBookverse", nil)

	cdc.RegisterConcrete(&MsgCreateBookverse{}, "Bookverses/CreateBookverse", nil)
	cdc.RegisterConcrete(&MsgUpdateBookverse{}, "Bookverses/UpdateBookverse", nil)
	cdc.RegisterConcrete(&MsgDeleteBookverse{}, "Bookverses/DeleteBookverse", nil)

}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendIbcTransferBookverse{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUnsellBookverse{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgClaimBookverse{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBidBookverse{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBuyBookverse{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSellBookverse{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateMarket{},
		&MsgUpdateMarket{},
		&MsgDeleteMarket{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgTransferBookverse{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBurnBookverse{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgMintBookverse{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateBookverse{},
		&MsgUpdateBookverse{},
		&MsgDeleteBookverse{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
