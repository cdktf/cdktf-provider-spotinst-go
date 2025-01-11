// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaksnp


type OceanAksNpUpdatePolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.206.0/docs/resources/ocean_aks_np#should_roll OceanAksNp#should_roll}.
	ShouldRoll interface{} `field:"required" json:"shouldRoll" yaml:"shouldRoll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.206.0/docs/resources/ocean_aks_np#conditioned_roll OceanAksNp#conditioned_roll}.
	ConditionedRoll interface{} `field:"optional" json:"conditionedRoll" yaml:"conditionedRoll"`
	// roll_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.206.0/docs/resources/ocean_aks_np#roll_config OceanAksNp#roll_config}
	RollConfig *OceanAksNpUpdatePolicyRollConfig `field:"optional" json:"rollConfig" yaml:"rollConfig"`
}

