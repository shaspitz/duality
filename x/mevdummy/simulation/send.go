package simulation

import (
	"math/rand"

	"github.com/NicholasDotSol/duality/x/mevdummy/keeper"
	"github.com/NicholasDotSol/duality/x/mevdummy/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgSend(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgSend{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the Send simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "Send simulation not implemented"), nil, nil
	}
}
