package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/dex module sentinel errors
var (
	ErrNoSpendableCoins = sdkerrors.Register(ModuleName, 1100, "No Spendable Coins found: %s")
	ErrNotEnoughCoins   = sdkerrors.Register(ModuleName, 1101, "Not enough Spendable Coins found: %s")
	ErrInvalidTokenPair = sdkerrors.Register(ModuleName, 1102, "Invalid Token Pair: (%s, %t)")
	ErrInvalidTokenListSize = sdkerrors.Register(ModuleName, 1103, "Invalid Array: Tokens0 size: %s, Tokens1 size: %t ")
)
