// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaws


type OceanAwsUpdatePolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.205.0/docs/resources/ocean_aws#should_roll OceanAws#should_roll}.
	ShouldRoll interface{} `field:"required" json:"shouldRoll" yaml:"shouldRoll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.205.0/docs/resources/ocean_aws#auto_apply_tags OceanAws#auto_apply_tags}.
	AutoApplyTags interface{} `field:"optional" json:"autoApplyTags" yaml:"autoApplyTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.205.0/docs/resources/ocean_aws#conditioned_roll OceanAws#conditioned_roll}.
	ConditionedRoll interface{} `field:"optional" json:"conditionedRoll" yaml:"conditionedRoll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.205.0/docs/resources/ocean_aws#conditioned_roll_params OceanAws#conditioned_roll_params}.
	ConditionedRollParams *[]*string `field:"optional" json:"conditionedRollParams" yaml:"conditionedRollParams"`
	// roll_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.205.0/docs/resources/ocean_aws#roll_config OceanAws#roll_config}
	RollConfig *OceanAwsUpdatePolicyRollConfig `field:"optional" json:"rollConfig" yaml:"rollConfig"`
}

