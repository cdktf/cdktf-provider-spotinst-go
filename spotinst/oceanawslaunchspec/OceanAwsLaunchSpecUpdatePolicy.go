// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanawslaunchspec


type OceanAwsLaunchSpecUpdatePolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.140.0/docs/resources/ocean_aws_launch_spec#should_roll OceanAwsLaunchSpec#should_roll}.
	ShouldRoll interface{} `field:"required" json:"shouldRoll" yaml:"shouldRoll"`
	// roll_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.140.0/docs/resources/ocean_aws_launch_spec#roll_config OceanAwsLaunchSpec#roll_config}
	RollConfig *OceanAwsLaunchSpecUpdatePolicyRollConfig `field:"optional" json:"rollConfig" yaml:"rollConfig"`
}

