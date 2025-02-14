// Copyright 2020 DSR Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/conversions"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/pagination"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/model/internal/types"
)

func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	modelQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the model module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	modelQueryCmd.AddCommand(client.GetCommands(
		GetCmdModel(storeKey, cdc),
		GetCmdAllModels(storeKey, cdc),
		GetCmdVendors(storeKey, cdc),
		GetCmdVendorModels(storeKey, cdc),
		GetCmdModelVersion(storeKey, cdc),
		GetCmdAllModelVersions(storeKey, cdc),
	)...)

	return modelQueryCmd
}

//nolint:dupl
func GetCmdModel(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-model",
		Short: "Query Model by combination of Vendor ID and Product ID",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			vid, err_ := conversions.ParseVID(viper.GetString(FlagVID))
			if err_ != nil {
				return err_
			}

			pid, err_ := conversions.ParsePID(viper.GetString(FlagPID))
			if err_ != nil {
				return err_
			}

			res, height, err := cliCtx.QueryStore(types.GetModelKey(vid, pid), queryRoute)
			if err != nil || res == nil {
				return types.ErrModelDoesNotExist(vid, pid)
			}

			var model types.Model
			cdc.MustUnmarshalBinaryBare(res, &model)

			return cliCtx.EncodeAndPrintWithHeight(model, height)
		},
	}

	cmd.Flags().String(FlagVID, "", "Model vendor ID")
	cmd.Flags().String(FlagPID, "", "Model product ID")
	cmd.Flags().Bool(cli.FlagPreviousHeight, false, cli.FlagPreviousHeightUsage)

	_ = cmd.MarkFlagRequired(FlagVID)
	_ = cmd.MarkFlagRequired(FlagPID)

	return cmd
}

//nolint:dupl
func GetCmdAllModels(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-models",
		Short: "Query the list of all Models",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)
			params := pagination.ParsePaginationParamsFromFlags()

			return cliCtx.QueryList(fmt.Sprintf("custom/%s/all_models", queryRoute), params)
		},
	}

	cmd.Flags().Int(pagination.FlagSkip, 0, "amount of models to skip")
	cmd.Flags().Int(pagination.FlagTake, 0, "amount of models to take")

	return cmd
}

func GetCmdVendors(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vendors",
		Short: "Query the list of Vendors",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)
			params := pagination.ParsePaginationParamsFromFlags()

			return cliCtx.QueryList(fmt.Sprintf("custom/%s/vendors", queryRoute), params)
		},
	}

	cmd.Flags().Int(pagination.FlagSkip, 0, "amount of vendors to skip")
	cmd.Flags().Int(pagination.FlagTake, 0, "amount of vendors to take")

	return cmd
}

func GetCmdVendorModels(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vendor-models",
		Short: "Query the list of Models for the given Vendor",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			vid, err_ := conversions.ParseVID(viper.GetString(FlagVID))
			if err_ != nil {
				return err_
			}

			res, height, err := cliCtx.QueryStore(types.GetVendorProductsKey(vid), queryRoute)
			if err != nil || res == nil {
				return types.ErrVendorProductsDoNotExist(vid)
			}

			var vendorProducts types.VendorProducts
			cdc.MustUnmarshalBinaryBare(res, &vendorProducts)

			return cliCtx.EncodeAndPrintWithHeight(vendorProducts, height)
		},
	}

	cmd.Flags().String(FlagVID, "", "Model vendor ID")
	cmd.Flags().Bool(cli.FlagPreviousHeight, false, cli.FlagPreviousHeightUsage)

	_ = cmd.MarkFlagRequired(FlagVID)

	return cmd
}

//nolint:dupl
func GetCmdModelVersion(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-model-version",
		Short: "Query Model Version by combination of Vendor ID, Product ID and Software Version",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			vid, err_ := conversions.ParseVID(viper.GetString(FlagVID))
			if err_ != nil {
				return err_
			}

			pid, err_ := conversions.ParsePID(viper.GetString(FlagPID))
			if err_ != nil {
				return err_
			}

			softwareVersion, err_ := conversions.ParseUInt32FromString(FlagSoftwareVersion, viper.GetString(FlagSoftwareVersion))
			if err_ != nil {
				return err_
			}

			res, height, err := cliCtx.QueryStore(types.GetModelVersionKey(vid, pid, softwareVersion), queryRoute)
			if err != nil || res == nil {
				return types.ErrModelVersionDoesNotExist(vid, pid, softwareVersion)
			}

			var modelVersion types.ModelVersion
			cdc.MustUnmarshalBinaryBare(res, &modelVersion)

			return cliCtx.EncodeAndPrintWithHeight(modelVersion, height)
		},
	}

	cmd.Flags().String(FlagVID, "", "Model vendor ID")
	cmd.Flags().String(FlagPID, "", "Model product ID")
	cmd.Flags().String(FlagSoftwareVersion, "", "Model Software Version")
	cmd.Flags().Bool(cli.FlagPreviousHeight, false, cli.FlagPreviousHeightUsage)

	_ = cmd.MarkFlagRequired(FlagVID)
	_ = cmd.MarkFlagRequired(FlagPID)
	_ = cmd.MarkFlagRequired(FlagSoftwareVersion)

	return cmd
}

//nolint:dupl
func GetCmdAllModelVersions(queryRoute string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-model-versions",
		Short: "Query the list of all versions for a given Device Model",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			vid, err_ := conversions.ParseVID(viper.GetString(FlagVID))
			if err_ != nil {
				return err_
			}

			pid, err_ := conversions.ParsePID(viper.GetString(FlagPID))
			if err_ != nil {
				return err_
			}

			res, height, err := cliCtx.QueryStore(types.GetModelVersionsKey(vid, pid), queryRoute)
			if err != nil || res == nil {
				return types.ErrNoModelVersionsExist(vid, pid)
			}

			var modelVersions types.ModelVersions
			cdc.MustUnmarshalBinaryBare(res, &modelVersions)

			return cliCtx.EncodeAndPrintWithHeight(modelVersions, height)
		},
	}

	cmd.Flags().String(FlagVID, "", "Model Vendor ID")
	cmd.Flags().String(FlagPID, "", "Model Product ID")
	cmd.Flags().Bool(cli.FlagPreviousHeight, false, cli.FlagPreviousHeightUsage)

	_ = cmd.MarkFlagRequired(FlagVID)
	_ = cmd.MarkFlagRequired(FlagPID)

	return cmd
}
