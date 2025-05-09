// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaksnpvirtualnodegroup


type OceanAksNpVirtualNodeGroupUpdatePolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.219.0/docs/resources/ocean_aks_np_virtual_node_group#should_roll OceanAksNpVirtualNodeGroup#should_roll}.
	ShouldRoll interface{} `field:"required" json:"shouldRoll" yaml:"shouldRoll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.219.0/docs/resources/ocean_aks_np_virtual_node_group#conditioned_roll OceanAksNpVirtualNodeGroup#conditioned_roll}.
	ConditionedRoll interface{} `field:"optional" json:"conditionedRoll" yaml:"conditionedRoll"`
	// roll_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.219.0/docs/resources/ocean_aks_np_virtual_node_group#roll_config OceanAksNpVirtualNodeGroup#roll_config}
	RollConfig *OceanAksNpVirtualNodeGroupUpdatePolicyRollConfig `field:"optional" json:"rollConfig" yaml:"rollConfig"`
}

