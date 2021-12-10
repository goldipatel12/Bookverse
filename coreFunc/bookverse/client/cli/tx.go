package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/goldipatel12/marketplacecoreFuncBookverses/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// this line is used by starport scaffolding # 1
	cmd.AddCommand(CmdSendIbcTransferBookverse())

	cmd.AddCommand(CmdUnsellBookverse())

	cmd.AddCommand(CmdClaimBookverse())

	cmd.AddCommand(CmdBidBookverse())

	cmd.AddCommand(CmdBuyBookverse())

	cmd.AddCommand(CmdSellBookverse())

	cmd.AddCommand(CmdCreateMarket())
	cmd.AddCommand(CmdUpdateMarket())
	cmd.AddCommand(CmdDeleteMarket())

	cmd.AddCommand(CmdTransferBookverse())

	cmd.AddCommand(CmdBurnBookverse())

	cmd.AddCommand(CmdMintBookverse())

	cmd.AddCommand(CmdCreateBookverse())
	cmd.AddCommand(CmdUpdateBookverse())
	cmd.AddCommand(CmdDeleteBookverse())

	return cmd
}
