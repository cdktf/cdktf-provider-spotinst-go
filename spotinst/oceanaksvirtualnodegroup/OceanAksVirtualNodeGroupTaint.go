// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaksvirtualnodegroup


type OceanAksVirtualNodeGroupTaint struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.142.0/docs/resources/ocean_aks_virtual_node_group#effect OceanAksVirtualNodeGroup#effect}.
	Effect *string `field:"required" json:"effect" yaml:"effect"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.142.0/docs/resources/ocean_aks_virtual_node_group#key OceanAksVirtualNodeGroup#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.142.0/docs/resources/ocean_aks_virtual_node_group#value OceanAksVirtualNodeGroup#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

