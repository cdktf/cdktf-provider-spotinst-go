// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanecs


type OceanEcsUpdatePolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.186.0/docs/resources/ocean_ecs#should_roll OceanEcs#should_roll}.
	ShouldRoll interface{} `field:"required" json:"shouldRoll" yaml:"shouldRoll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.186.0/docs/resources/ocean_ecs#auto_apply_tags OceanEcs#auto_apply_tags}.
	AutoApplyTags interface{} `field:"optional" json:"autoApplyTags" yaml:"autoApplyTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.186.0/docs/resources/ocean_ecs#conditioned_roll OceanEcs#conditioned_roll}.
	ConditionedRoll interface{} `field:"optional" json:"conditionedRoll" yaml:"conditionedRoll"`
	// roll_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.186.0/docs/resources/ocean_ecs#roll_config OceanEcs#roll_config}
	RollConfig *OceanEcsUpdatePolicyRollConfig `field:"optional" json:"rollConfig" yaml:"rollConfig"`
}

