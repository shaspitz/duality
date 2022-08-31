package keeper

import (
	"context"

	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) AddLiquidityVerification(goCtx context.Context, msg *types.MsgAddLiquidity) (string, string, sdk.AccAddress, sdk.AccAddress, sdk.Dec, sdk.Dec, sdk.Dec, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	price, err := sdk.NewDecFromStr(msg.Price)
	// Error checking for valid sdk.Dec
	if err != nil {
		return "", "", nil, nil, sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec(), sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "Not a valid decimal type: %s", err)
	}

	token0, token1, priceDec, err := k.SortTokens(ctx, msg.TokenA, msg.TokenB, price)

	if err != nil {
		return "", "", nil, nil, sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec(), sdkerrors.Wrapf(types.ErrInvalidTokenPair, "Not a valid Token Pair: tokenA and tokenB cannot be the same")
	}

	// Converts input address (string) to sdk.AccAddress
	callerAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	// Error checking for the calling address
	if err != nil {
		return "", "", nil, nil, sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec(), sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	// Converts receiver address (string) to sdk.AccAddress
	receiverAddr, err := sdk.AccAddressFromBech32(msg.Receiver)
	// Error checking for the valid receiver address
	if err != nil {
		return "", "", nil, nil, sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec(), sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid receiver address (%s)", err)
	}

	amount0, err := sdk.NewDecFromStr(msg.AmountA)
	amount1, err := sdk.NewDecFromStr(msg.AmountB)

	if token0 != msg.TokenA {
		tmp := amount0
		amount0 = amount1
		amount1 = tmp
	}

	// Error checking for valid sdk.Dec
	if err != nil {
		return "", "", nil, nil, sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec(), sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "Not a valid decimal type: %s", err)
	}

	if msg.TokenDirection != msg.TokenA && msg.TokenB != msg.TokenDirection {
		return "", "", nil, nil, sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec(), sdkerrors.Wrapf(types.ErrValidPairNotFound, "Token Direction must be the same as either Token A or Token B")
	}

	AccountsToken0Balance := sdk.NewDecFromInt(k.bankKeeper.GetBalance(ctx, callerAddr, token0).Amount)

	// Error handling to verify the amount wished to deposit is NOT more then the msg.creator holds in their accounts
	if AccountsToken0Balance.LT(amount0) {
		return "", "", nil, nil, sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec(), sdkerrors.Wrapf(types.ErrNotEnoughCoins, "Address %s  does not have enough of token 0", callerAddr)
	}

	AccountsToken1Balance := sdk.NewDecFromInt(k.bankKeeper.GetBalance(ctx, callerAddr, token1).Amount)

	// Error handling to verify the amount wished to deposit is NOT more then the msg.creator holds in their accounts
	if AccountsToken1Balance.LT(amount1) {
		return "", "", nil, nil, sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec(), sdkerrors.Wrapf(types.ErrNotEnoughCoins, "Address %s  does not have enough of token 0", callerAddr)
	}

	return token0, token1, callerAddr, receiverAddr, amount0, amount1, priceDec, nil
}

func (k msgServer) RemoveLiquidityVerification(goCtx context.Context, msg *types.MsgRemoveLiquidity) (string, string, sdk.AccAddress, sdk.AccAddress, sdk.Dec, sdk.Dec, sdk.Dec, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	price, err := sdk.NewDecFromStr(msg.Price)
	// Error checking for valid sdk.Dec
	if err != nil {
		return "", "", nil, nil, sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec(), sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "Not a valid decimal type: %s", err)
	}

	token0, token1, priceDec, err := k.SortTokens(ctx, msg.TokenA, msg.TokenB, price)

	if err != nil {
		return "", "", nil, nil, sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec(), sdkerrors.Wrapf(types.ErrInvalidTokenPair, "Not a valid Token Pair: tokenA and tokenB cannot be the same")
	}

	// Converts input address (string) to sdk.AccAddress
	callerAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	// Error checking for the calling address
	if err != nil {
		return "", "", nil, nil, sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec(), sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	// Converts receiver address (string) to sdk.AccAddress
	receiverAddr, err := sdk.AccAddressFromBech32(msg.Receiver)
	// Error checking for the valid receiver address
	if err != nil {
		return "", "", nil, nil, sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec(), sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid receiver address (%s)", err)
	}

	sharesToRemove, err := sdk.NewDecFromStr(msg.Shares)
	// Error checking for valid sdk.Dec
	if err != nil {
		return "", "", nil, nil, sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec(), sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "Not a valid decimal type: %s", err)
	}

	sharesOwned, sharesFound := k.GetShares(ctx, token0, token1, msg.Creator, msg.Price, msg.Fee, msg.OrderType)

	if !sharesFound || sharesToRemove.GT(sharesOwned.SharesOwned) {
		return "", "", nil, nil, sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec(), sdkerrors.Wrapf(types.ErrNotEnoughShares, "Not enough shares owned")
	}

	return token0, token1, callerAddr, receiverAddr, sharesToRemove, sharesOwned.SharesOwned, priceDec, nil
}