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

func CmdMintBookverse() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mintBookverse [sender] [recipient] [sID] [denom] [name] [description] [image] [tokenUri]",
		Short: "Broadcast message mintBookverse",
		Args:  cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsSender := string(args[0])
			argsRecipient := string(args[1])
			argsSID := string(args[2])
			argsDenom := string(args[3])
			argsName := string(args[4])
			argsDescription := string(args[5])
			argsImage := string(args[6])
			argsTokenUri := string(args[7])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgMintBookverse(clientCtx.GetFromAddress().String(), string(argsSender), string(argsRecipient), string(argsSID), string(argsDenom), string(argsName), string(argsDescription), string(argsImage), string(argsTokenUri))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
