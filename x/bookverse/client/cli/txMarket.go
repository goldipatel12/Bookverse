package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/spf13/cast"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/goldipatel12/marketplace/x/Bookverses/types"
)

func CmdCreateMarket() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-market [sID] [price] [status] [seller]",
		Short: "Create a new market",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsSID, err := cast.ToStringE(args[0])
			if err != nil {
				return err
			}
			argsPrice, err := cast.ToInt32E(args[1])
			if err != nil {
				return err
			}
			argsStatus, err := cast.ToStringE(args[2])
			if err != nil {
				return err
			}
			argsSeller, err := cast.ToStringE(args[3])
			if err != nil {
				return err
			}
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateMarket(clientCtx.GetFromAddress().String(), argsSID, argsPrice, argsStatus, argsSeller)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateMarket() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-market [id] [sID] [price] [status] [seller]",
		Short: "Update a market",
		Args:  cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argsSID, err := cast.ToStringE(args[1])
			if err != nil {
				return err
			}

			argsPrice, err := cast.ToInt32E(args[2])
			if err != nil {
				return err
			}

			argsStatus, err := cast.ToStringE(args[3])
			if err != nil {
				return err
			}

			argsSeller, err := cast.ToStringE(args[4])
			if err != nil {
				return err
			}
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateMarket(clientCtx.GetFromAddress().String(), id, argsSID, argsPrice, argsStatus, argsSeller)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteMarket() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-market [id]",
		Short: "Delete a market by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteMarket(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
