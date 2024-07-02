// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanawslaunchspec


type OceanAwsLaunchSpecDeleteOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.1/docs/resources/ocean_aws_launch_spec#force_delete OceanAwsLaunchSpec#force_delete}.
	ForceDelete interface{} `field:"required" json:"forceDelete" yaml:"forceDelete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.1/docs/resources/ocean_aws_launch_spec#delete_nodes OceanAwsLaunchSpec#delete_nodes}.
	DeleteNodes interface{} `field:"optional" json:"deleteNodes" yaml:"deleteNodes"`
}

