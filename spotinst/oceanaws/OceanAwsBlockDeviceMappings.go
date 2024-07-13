// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaws


type OceanAwsBlockDeviceMappings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.181.0/docs/resources/ocean_aws#device_name OceanAws#device_name}.
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// ebs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.181.0/docs/resources/ocean_aws#ebs OceanAws#ebs}
	Ebs *OceanAwsBlockDeviceMappingsEbs `field:"optional" json:"ebs" yaml:"ebs"`
}

