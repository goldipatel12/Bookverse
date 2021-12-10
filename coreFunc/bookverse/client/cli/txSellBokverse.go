package cli

import (
	"github.com/spf13/cobra"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/goldipatel12/marketplacecoreFuncBookverses/types"
)

var _ = strconv.Itoa(0)

func CmdSellBookverse() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sellBookverse [sID] [seller] [price]",
		Short: "Broadcast message sellBookverse",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsSID := string(args[0])
			argsSeller := string(args[1])
			argsPrice, _ := strconv.ParseInt(args[2], 10, 64)

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSellBookverse(clientCtx.GetFromAddress().String(), string(argsSID), string(argsSeller), int32(argsPrice))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
