package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/validator/types"
)

var _ = strconv.Itoa(0)

func CmdDisableValidator() *cobra.Command {
	var (
		address string
		info    string
	)

	cmd := &cobra.Command{
		Use:   "disable-validator --address [address] --info [info]",
		Short: "Disables the Validator node by a NodeAdmin.",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			fromAddr := clientCtx.GetFromAddress()
			msg := types.NewMsgDisableValidator(
				sdk.ValAddress(fromAddr),
			)

			// validate basic will be called in GenerateOrBroadcastTxCLI
			err = tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
			if cli.IsWriteInsteadReadRpcError(err) {
				return clientCtx.PrintString(cli.LightClientProxyForWriteRequests)
			}
			return err
		},
	}

	cmd.Flags().StringVar(&address, FlagAddress, "", "Bench32 encoded validator address or owner account")
	cmd.Flags().StringVar(&info, FlagInfo, "", "Optional information/notes for approval, proposal, disable or enable validator")

	flags.AddTxFlagsToCmd(cmd)

	_ = cmd.MarkFlagRequired(FlagAddress)

	return cmd
}
