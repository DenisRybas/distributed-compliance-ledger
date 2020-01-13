package cli

import (
	"fmt"

	"github.com/spf13/viper"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/compliance/internal/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

const (
	FlagSkip = "skip"
	FlagTake = "take"
)

func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	complianceQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the compliance module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	complianceQueryCmd.AddCommand(client.GetCommands(
		GetCmdModelInfo(storeKey, cdc),
		GetCmdModelInfoWithProof(storeKey, cdc),
		GetCmdModelInfoHeaders(storeKey, cdc),
	)...)

	return complianceQueryCmd
}

func GetCmdModelInfo(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "model-info [id]",
		Short: "Query ModelInfo by ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			id := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/model_info/%s", queryRoute, id), nil)
			if err != nil {
				fmt.Printf("could not query ModelInfo - %s \n", id)
				return nil
			}

			var out types.ModelInfo
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdModelInfoWithProof(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "model-info-with-proof [id]",
		Short: "Query ModelInfo with proof by ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			id := args[0]

			res, _, err := cliCtx.QueryStore([]byte(id), queryRoute)
			if err != nil {
				fmt.Printf("could not query ModelInfo - %s \n", id)
				return nil
			}

			var out types.ModelInfo
			cdc.MustUnmarshalBinaryBare(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdModelInfoHeaders(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "model-info-headers",
		Short: "List all ModelInfo IDs",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			params := types.NewQueryModelInfoHeadersParams(
				viper.GetInt(FlagSkip),
				viper.GetInt(FlagTake),
			)

			data := cdc.MustMarshalJSON(params)

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/model_info_headers", queryRoute), data)
			if err != nil {
				fmt.Printf("could not get query names\n")
				return nil
			}

			var out types.QueryModelInfoHeadersResult
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}

	cmd.Flags().Int(FlagSkip, 0, "amount of accounts to skip")
	cmd.Flags().Int(FlagTake, 0, "amount of accounts to take")

	return cmd
}
