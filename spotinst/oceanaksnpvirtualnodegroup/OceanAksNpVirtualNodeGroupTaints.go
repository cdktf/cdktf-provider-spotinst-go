// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaksnpvirtualnodegroup


type OceanAksNpVirtualNodeGroupTaints struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.219.0/docs/resources/ocean_aks_np_virtual_node_group#effect OceanAksNpVirtualNodeGroup#effect}.
	Effect *string `field:"required" json:"effect" yaml:"effect"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.219.0/docs/resources/ocean_aks_np_virtual_node_group#key OceanAksNpVirtualNodeGroup#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.219.0/docs/resources/ocean_aks_np_virtual_node_group#value OceanAksNpVirtualNodeGroup#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

