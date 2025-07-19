// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanawslaunchspec


type OceanAwsLaunchSpecSchedulingShutdownHours struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.224.0/docs/resources/ocean_aws_launch_spec#time_windows OceanAwsLaunchSpec#time_windows}.
	TimeWindows *[]*string `field:"required" json:"timeWindows" yaml:"timeWindows"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.224.0/docs/resources/ocean_aws_launch_spec#is_enabled OceanAwsLaunchSpec#is_enabled}.
	IsEnabled interface{} `field:"optional" json:"isEnabled" yaml:"isEnabled"`
}

