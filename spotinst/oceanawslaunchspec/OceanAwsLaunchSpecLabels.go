// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanawslaunchspec


type OceanAwsLaunchSpecLabels struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.146.0/docs/resources/ocean_aws_launch_spec#key OceanAwsLaunchSpec#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.146.0/docs/resources/ocean_aws_launch_spec#value OceanAwsLaunchSpec#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

