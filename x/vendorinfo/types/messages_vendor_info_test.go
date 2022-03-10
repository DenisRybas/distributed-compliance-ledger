package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	tmrand "github.com/tendermint/tendermint/libs/rand"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/testutil/sample"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/validator"
)

func TestMsgCreateVendorInfo_ValidateBasic(t *testing.T) {
	negative_tests := []struct {
		name string
		msg  MsgCreateVendorInfo
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateVendorInfo{
				Creator:              "invalid_address",
				Vid:                  testconstants.VendorID1,
				VendorName:           testconstants.VendorName,
				CompanyLegalName:     testconstants.CompanyLegalName,
				CompanyPrefferedName: testconstants.CompanyPreferredName,
				VendorLandingPageUrl: testconstants.VendorLandingPageUrl,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "vendor name is not set",
			msg: MsgCreateVendorInfo{
				Creator:              sample.AccAddress(),
				Vid:                  testconstants.VendorID1,
				VendorName:           "",
				CompanyLegalName:     testconstants.CompanyLegalName,
				CompanyPrefferedName: testconstants.CompanyPreferredName,
				VendorLandingPageUrl: testconstants.VendorLandingPageUrl,
			},
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "company legal name is not set",
			msg: MsgCreateVendorInfo{
				Creator:              sample.AccAddress(),
				Vid:                  testconstants.VendorID1,
				VendorName:           testconstants.VendorName,
				CompanyLegalName:     "",
				CompanyPrefferedName: testconstants.CompanyPreferredName,
				VendorLandingPageUrl: testconstants.VendorLandingPageUrl,
			},
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "vid less than 0",
			msg: MsgCreateVendorInfo{
				Creator:              sample.AccAddress(),
				Vid:                  -1,
				VendorName:           testconstants.VendorName,
				CompanyLegalName:     testconstants.CompanyLegalName,
				CompanyPrefferedName: testconstants.CompanyPreferredName,
				VendorLandingPageUrl: testconstants.VendorLandingPageUrl,
			},
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "vid is 0",
			msg: MsgCreateVendorInfo{
				Creator:              sample.AccAddress(),
				Vid:                  0,
				VendorName:           testconstants.VendorName,
				CompanyLegalName:     testconstants.CompanyLegalName,
				CompanyPrefferedName: testconstants.CompanyPreferredName,
				VendorLandingPageUrl: testconstants.VendorLandingPageUrl,
			},
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "vid is bigger than 65535",
			msg: MsgCreateVendorInfo{
				Creator:              sample.AccAddress(),
				Vid:                  65536,
				VendorName:           testconstants.VendorName,
				CompanyLegalName:     testconstants.CompanyLegalName,
				CompanyPrefferedName: testconstants.CompanyPreferredName,
				VendorLandingPageUrl: testconstants.VendorLandingPageUrl,
			},
			err: validator.ErrFieldUpperBoundViolated,
		},
		{
			name: "vendor name len > 128",
			msg: MsgCreateVendorInfo{
				Creator:              sample.AccAddress(),
				Vid:                  testconstants.VendorID1,
				VendorName:           tmrand.Str(129),
				CompanyLegalName:     testconstants.CompanyLegalName,
				CompanyPrefferedName: testconstants.CompanyPreferredName,
				VendorLandingPageUrl: testconstants.VendorLandingPageUrl,
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "company legal name len > 256",
			msg: MsgCreateVendorInfo{
				Creator:              sample.AccAddress(),
				Vid:                  testconstants.VendorID1,
				VendorName:           testconstants.VendorName,
				CompanyLegalName:     tmrand.Str(257),
				CompanyPrefferedName: testconstants.CompanyPreferredName,
				VendorLandingPageUrl: testconstants.VendorLandingPageUrl,
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "company preffered name len > 256",
			msg: MsgCreateVendorInfo{
				Creator:              sample.AccAddress(),
				Vid:                  testconstants.VendorID1,
				VendorName:           testconstants.VendorName,
				CompanyLegalName:     testconstants.CompanyPreferredName,
				CompanyPrefferedName: tmrand.Str(257),
				VendorLandingPageUrl: testconstants.VendorLandingPageUrl,
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "vendor landing page URL is not URL",
			msg: MsgCreateVendorInfo{
				Creator:              sample.AccAddress(),
				Vid:                  testconstants.VendorID1,
				VendorName:           testconstants.VendorName,
				CompanyLegalName:     testconstants.CompanyPreferredName,
				VendorLandingPageUrl: "ABC",
				CompanyPrefferedName: testconstants.CompanyPreferredName,
			},
			err: validator.ErrFieldNotValid,
		},
		{
			name: "vendor landing page URL len > 256",
			msg: MsgCreateVendorInfo{
				Creator:              sample.AccAddress(),
				Vid:                  testconstants.VendorID1,
				VendorName:           testconstants.VendorName,
				CompanyLegalName:     testconstants.CompanyPreferredName,
				VendorLandingPageUrl: "https://www.example.com/ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789",
				CompanyPrefferedName: testconstants.CompanyPreferredName,
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
	}

	positive_tests := []struct {
		name string
		msg  MsgCreateVendorInfo
	}{
		{
			name: "valid create vendorinfo message",
			msg: MsgCreateVendorInfo{
				Creator:              sample.AccAddress(),
				Vid:                  testconstants.VendorID1,
				VendorName:           testconstants.VendorName,
				CompanyLegalName:     testconstants.CompanyLegalName,
				CompanyPrefferedName: testconstants.CompanyPreferredName,
				VendorLandingPageUrl: testconstants.VendorLandingPageUrl,
			},
		},
	}
	for _, tt := range positive_tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			require.NoError(t, err)
		})
	}

	for _, tt := range negative_tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			require.Error(t, err)
			require.ErrorIs(t, err, tt.err)
		})
	}
}

func TestMsgUpdateVendorInfo_ValidateBasic(t *testing.T) {
	negative_tests := []struct {
		name string
		msg  MsgUpdateVendorInfo
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateVendorInfo{
				Creator:              "invalid_address",
				Vid:                  testconstants.VendorID1,
				VendorName:           testconstants.VendorName,
				CompanyLegalName:     testconstants.CompanyLegalName,
				CompanyPrefferedName: testconstants.CompanyPreferredName,
				VendorLandingPageUrl: testconstants.VendorLandingPageUrl,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "vendor name is not set",
			msg: MsgUpdateVendorInfo{
				Creator:              sample.AccAddress(),
				Vid:                  testconstants.VendorID1,
				VendorName:           "",
				CompanyLegalName:     testconstants.CompanyLegalName,
				CompanyPrefferedName: testconstants.CompanyPreferredName,
				VendorLandingPageUrl: testconstants.VendorLandingPageUrl,
			},
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "company legal name is not set",
			msg: MsgUpdateVendorInfo{
				Creator:              sample.AccAddress(),
				Vid:                  testconstants.VendorID1,
				VendorName:           testconstants.VendorName,
				CompanyLegalName:     "",
				CompanyPrefferedName: testconstants.CompanyPreferredName,
				VendorLandingPageUrl: testconstants.VendorLandingPageUrl,
			},
			err: validator.ErrRequiredFieldMissing,
		},
		{
			name: "vid less than 0",
			msg: MsgUpdateVendorInfo{
				Creator:              sample.AccAddress(),
				Vid:                  -1,
				VendorName:           testconstants.VendorName,
				CompanyLegalName:     testconstants.CompanyLegalName,
				CompanyPrefferedName: testconstants.CompanyPreferredName,
				VendorLandingPageUrl: testconstants.VendorLandingPageUrl,
			},
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "vid is 0",
			msg: MsgUpdateVendorInfo{
				Creator:              sample.AccAddress(),
				Vid:                  0,
				VendorName:           testconstants.VendorName,
				CompanyLegalName:     testconstants.CompanyLegalName,
				CompanyPrefferedName: testconstants.CompanyPreferredName,
				VendorLandingPageUrl: testconstants.VendorLandingPageUrl,
			},
			err: validator.ErrFieldLowerBoundViolated,
		},
		{
			name: "vid is bigger than 65535",
			msg: MsgUpdateVendorInfo{
				Creator:              sample.AccAddress(),
				Vid:                  65536,
				VendorName:           testconstants.VendorName,
				CompanyLegalName:     testconstants.CompanyLegalName,
				CompanyPrefferedName: testconstants.CompanyPreferredName,
				VendorLandingPageUrl: testconstants.VendorLandingPageUrl,
			},
			err: validator.ErrFieldUpperBoundViolated,
		},
		{
			name: "vendor name len > 128",
			msg: MsgUpdateVendorInfo{
				Creator:              sample.AccAddress(),
				Vid:                  testconstants.VendorID1,
				VendorName:           tmrand.Str(129),
				CompanyLegalName:     testconstants.CompanyLegalName,
				CompanyPrefferedName: testconstants.CompanyPreferredName,
				VendorLandingPageUrl: testconstants.VendorLandingPageUrl,
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "company legal name len > 256",
			msg: MsgUpdateVendorInfo{
				Creator:              sample.AccAddress(),
				Vid:                  testconstants.VendorID1,
				VendorName:           testconstants.VendorName,
				CompanyLegalName:     tmrand.Str(257),
				CompanyPrefferedName: testconstants.CompanyPreferredName,
				VendorLandingPageUrl: testconstants.VendorLandingPageUrl,
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "company preffered name len > 256",
			msg: MsgUpdateVendorInfo{
				Creator:              sample.AccAddress(),
				Vid:                  testconstants.VendorID1,
				VendorName:           testconstants.VendorName,
				CompanyLegalName:     testconstants.CompanyPreferredName,
				CompanyPrefferedName: tmrand.Str(257),
				VendorLandingPageUrl: testconstants.VendorLandingPageUrl,
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
		{
			name: "vendor landing page URL is not URL",
			msg: MsgUpdateVendorInfo{
				Creator:              sample.AccAddress(),
				Vid:                  testconstants.VendorID1,
				VendorName:           testconstants.VendorName,
				CompanyLegalName:     testconstants.CompanyPreferredName,
				VendorLandingPageUrl: "ABC",
				CompanyPrefferedName: testconstants.CompanyPreferredName,
			},
			err: validator.ErrFieldNotValid,
		},
		{
			name: "vendor landing page URL len > 256",
			msg: MsgUpdateVendorInfo{
				Creator:              sample.AccAddress(),
				Vid:                  testconstants.VendorID1,
				VendorName:           testconstants.VendorName,
				CompanyLegalName:     testconstants.CompanyPreferredName,
				VendorLandingPageUrl: "https://www.example.com/ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789",
				CompanyPrefferedName: testconstants.CompanyPreferredName,
			},
			err: validator.ErrFieldMaxLengthExceeded,
		},
	}

	positive_tests := []struct {
		name string
		msg  MsgUpdateVendorInfo
	}{
		{
			name: "valid create vendorinfo message",
			msg: MsgUpdateVendorInfo{
				Creator:              sample.AccAddress(),
				Vid:                  testconstants.VendorID1,
				VendorName:           testconstants.VendorName,
				CompanyLegalName:     testconstants.CompanyLegalName,
				CompanyPrefferedName: testconstants.CompanyPreferredName,
				VendorLandingPageUrl: testconstants.VendorLandingPageUrl,
			},
		},
		{
			name: "valid create vendorinfo message",
			msg: MsgUpdateVendorInfo{
				Creator:              sample.AccAddress(),
				Vid:                  testconstants.VendorID1,
				VendorName:           testconstants.VendorName,
				CompanyLegalName:     testconstants.CompanyLegalName,
				VendorLandingPageUrl: testconstants.VendorLandingPageUrl,
			},
		},
	}
	for _, tt := range positive_tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			require.NoError(t, err)
		})
	}

	for _, tt := range negative_tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			require.Error(t, err)
			require.ErrorIs(t, err, tt.err)
		})
	}
}
