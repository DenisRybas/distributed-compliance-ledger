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

const (
	// Flags for Model.
	FlagVID                                        = "vid"
	FlagPID                                        = "pid"
	FlagDeviceTypeID                               = "deviceTypeID"
	FlagProductName                                = "productName"
	FlagProductNameShortcut                        = "n"
	FlagProductLabel                               = "productLabel"
	FlagProductLabelShortcut                       = "d"
	FlagPartNumber                                 = "partNumber"
	FlagCommissioningCustomFlow                    = "commissioningCustomFlow"
	FlagCommissioningCustomFlowURL                 = "commissioningCustomFlowURL"
	FlagCommissioningModeInitialStepsHint          = "commissioningModeInitialStepsHint"
	FlagCommissioningModeInitialStepsInstruction   = "commissioningModeInitialStepsInstruction"
	FlagCommissioningModeSecondaryStepsHint        = "commissioningModeSecondaryStepsHint"
	FlagCommissioningModeSecondaryStepsInstruction = "commissioningModeSecondaryStepsInstruction"
	FlagUserManualURL                              = "userManualURL"
	FlagSupportURL                                 = "supportURL"
	FlagProductURL                                 = "productURL"

	// Flags for Model Software Version.
	FlagSoftwareVersion              = "softwareVersion"
	FlagSoftwareVersionShortcut      = "v"
	FlagSoftwareVersionString        = "softwareVersionString"
	FlagCDVersionNumber              = "cdVersionNumber"
	FlagFirmwareDigests              = "firmwareDigests"
	FlagSoftwareVersionValid         = "softwareVersionValid"
	FlagOtaURL                       = "otaURL"
	FlagOtaFileSize                  = "otaFileSize"
	FlagOtaChecksum                  = "otaChecksum"
	FlagOtaChecksumType              = "otaChecksumType"
	FlagMinApplicableSoftwareVersion = "minApplicableSoftwareVersion"
	FlagMaxApplicableSoftwareVersion = "maxApplicableSoftwareVersion"
	FlagReleaseNotesURL              = "releaseNotesURL"
)
