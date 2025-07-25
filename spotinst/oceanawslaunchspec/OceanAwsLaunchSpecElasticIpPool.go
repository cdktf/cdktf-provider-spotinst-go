// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanawslaunchspec


type OceanAwsLaunchSpecElasticIpPool struct {
	// tag_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.224.1/docs/resources/ocean_aws_launch_spec#tag_selector OceanAwsLaunchSpec#tag_selector}
	TagSelector *OceanAwsLaunchSpecElasticIpPoolTagSelector `field:"optional" json:"tagSelector" yaml:"tagSelector"`
}

