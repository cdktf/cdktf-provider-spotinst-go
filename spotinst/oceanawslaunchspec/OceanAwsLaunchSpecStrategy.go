// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanawslaunchspec


type OceanAwsLaunchSpecStrategy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.216.1/docs/resources/ocean_aws_launch_spec#draining_timeout OceanAwsLaunchSpec#draining_timeout}.
	DrainingTimeout *float64 `field:"optional" json:"drainingTimeout" yaml:"drainingTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.216.1/docs/resources/ocean_aws_launch_spec#spot_percentage OceanAwsLaunchSpec#spot_percentage}.
	SpotPercentage *float64 `field:"optional" json:"spotPercentage" yaml:"spotPercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.216.1/docs/resources/ocean_aws_launch_spec#utilize_commitments OceanAwsLaunchSpec#utilize_commitments}.
	UtilizeCommitments interface{} `field:"optional" json:"utilizeCommitments" yaml:"utilizeCommitments"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.216.1/docs/resources/ocean_aws_launch_spec#utilize_reserved_instances OceanAwsLaunchSpec#utilize_reserved_instances}.
	UtilizeReservedInstances interface{} `field:"optional" json:"utilizeReservedInstances" yaml:"utilizeReservedInstances"`
}

