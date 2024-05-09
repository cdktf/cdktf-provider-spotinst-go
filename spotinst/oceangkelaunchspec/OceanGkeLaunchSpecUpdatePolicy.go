// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceangkelaunchspec


type OceanGkeLaunchSpecUpdatePolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.172.0/docs/resources/ocean_gke_launch_spec#should_roll OceanGkeLaunchSpec#should_roll}.
	ShouldRoll interface{} `field:"required" json:"shouldRoll" yaml:"shouldRoll"`
	// roll_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.172.0/docs/resources/ocean_gke_launch_spec#roll_config OceanGkeLaunchSpec#roll_config}
	RollConfig *OceanGkeLaunchSpecUpdatePolicyRollConfig `field:"optional" json:"rollConfig" yaml:"rollConfig"`
}

