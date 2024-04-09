// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaws


type OceanAwsLogging struct {
	// export block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.169.0/docs/resources/ocean_aws#export OceanAws#export}
	Export *OceanAwsLoggingExport `field:"optional" json:"export" yaml:"export"`
}

