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

func CmdCreateBookverse() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-Bookverse [sID] [owner] [name] [description] [image] [tokenUri]",
		Short: "Create a new Bookverse",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsSID, err := cast.ToStringE(args[0])
			if err != nil {
				return err
			}
			argsOwner, err := cast.ToStringE(args[1])
			if err != nil {
				return err
			}
			argsName, err := cast.ToStringE(args[2])
			if err != nil {
				return err
			}
			argsDescription, err := cast.ToStringE(args[3])
			if err != nil {
				return err
			}
			argsImage, err := cast.ToStringE(args[4])
			if err != nil {
				return err
			}
			argsTokenUri, err := cast.ToStringE(args[5])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateBookverse(clientCtx.GetFromAddress().String(), argsSID, argsOwner, argsName, argsDescription, argsImage, argsTokenUri)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateBookverse() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-Bookverse [id] [sID] [owner] [name] [description] [image] [tokenUri]",
		Short: "Update a Bookverse",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argsSID, err := cast.ToStringE(args[1])
			if err != nil {
				return err
			}

			argsOwner, err := cast.ToStringE(args[2])
			if err != nil {
				return err
			}

			argsName, err := cast.ToStringE(args[3])
			if err != nil {
				return err
			}

			argsDescription, err := cast.ToStringE(args[4])
			if err != nil {
				return err
			}

			argsImage, err := cast.ToStringE(args[5])
			if err != nil {
				return err
			}

			argsTokenUri, err := cast.ToStringE(args[6])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateBookverse(clientCtx.GetFromAddress().String(), id, argsSID, argsOwner, argsName, argsDescription, argsImage, argsTokenUri)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteBookverse() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-Bookverse [id]",
		Short: "Delete a Bookverse by id",
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

			msg := types.NewMsgDeleteBookverse(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
