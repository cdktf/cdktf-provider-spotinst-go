// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanawslaunchspec


type OceanAwsLaunchSpecLoadBalancers struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.230.0/docs/resources/ocean_aws_launch_spec#type OceanAwsLaunchSpec#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.230.0/docs/resources/ocean_aws_launch_spec#arn OceanAwsLaunchSpec#arn}.
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.230.0/docs/resources/ocean_aws_launch_spec#name OceanAwsLaunchSpec#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

