package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// ErrInvalid is error msg
var (
	ErrInvalid = sdkerrors.Register(ModuleName, 1, "custom error message")
)
