package vendorinfo_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/zigbee-alliance/distributed-compliance-ledger/testutil/keeper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/vendorinfo/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		VendorInfoList: []types.VendorInfo{
			{
				Vid: 0,
			},
			{
				Vid: 1,
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}
	dclauthKeeper := &vendorinfo.DclauthKeeperMock{}

	k, ctx := keepertest.VendorinfoKeeper(t, dclauthKeeper)
	vendorinfo.InitGenesis(ctx, *k, genesisState)
	got := vendorinfo.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	require.Len(t, got.VendorInfoList, len(genesisState.VendorInfoList))
	require.Subset(t, genesisState.VendorInfoList, got.VendorInfoList)
	// this line is used by starport scaffolding # genesis/test/assert
}
