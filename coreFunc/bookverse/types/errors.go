package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/Bookverses module sentinel errors
var (
	ErrSample               = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrInvalidPacketTimeout = sdkerrors.Register(ModuleName, 1500, "invalid packet timeout")
	ErrInvalidVersion       = sdkerrors.Register(ModuleName, 1501, "invalid version")

	ErrRequiredFields = sdkerrors.Register(ModuleName, 8001, "Required fields")
	ErrInvalidDate    = sdkerrors.Register(ModuleName, 8002, "Invalid date")
	ErrInvalidPrice   = sdkerrors.Register(ModuleName, 8003, "Invalid price")
	ErrExistedData    = sdkerrors.Register(ModuleName, 8004, "Existed data")
	ErrInvalidOffers  = sdkerrors.Register(ModuleName, 8005, "Invalid offers")
	ErrInvalidBuyer   = sdkerrors.Register(ModuleName, 8006, "Invalid buyer")
)
