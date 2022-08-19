package keeper

import (
	"context"

	"github.com/NicholasDotSol/duality/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AddLiquidity(goCtx context.Context, msg *types.MsgAddLiquidity) (*types.MsgAddLiquidityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	token0, token1, callerAdr, receiverAdr, amounts, price, err := k.AddLiquidityVerification(goCtx, msg)

	if err != nil {
		return nil, err
	}

	err = k.SingleDeposit(goCtx, token0, token1, amounts, price, msg, callerAdr, receiverAdr)

	if err != nil {
		return nil, err
	}

	//else {
	// MultiDepsoit
	//}
	_ = token0
	_ = token1
	_ = callerAdr
	_ = receiverAdr
	_ = amounts
	_ = ctx

	return &types.MsgAddLiquidityResponse{}, nil
}