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

func CmdTransferBookverse() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transferBookverse [sender] [recipient] [sID] [denom]",
		Short: "Broadcast message transferBookverse",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsSender := string(args[0])
			argsRecipient := string(args[1])
			argsSID := string(args[2])
			argsDenom := string(args[3])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgTransferBookverse(clientCtx.GetFromAddress().String(), string(argsSender), string(argsRecipient), string(argsSID), string(argsDenom))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
