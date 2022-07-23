package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) Min(a, b sdk.Dec) sdk.Dec {
	if a.LT(b) {
		return a
	}
	return b
}
