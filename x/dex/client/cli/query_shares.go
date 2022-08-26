package cli

import (
	"context"

	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListShares() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-shares",
		Short: "list all Shares",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllSharesRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.SharesAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowShares() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-shares [address] [price] [fee] [order-type]",
		Short: "shows a Shares",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argAddress := args[0]
			argPrice := args[1]
			argFee := args[2]
			argOrderType := args[3]

			params := &types.QueryGetSharesRequest{
				Address:   argAddress,
				Price:     argPrice,
				Fee:       argFee,
				OrderType: argOrderType,
			}

			res, err := queryClient.Shares(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
