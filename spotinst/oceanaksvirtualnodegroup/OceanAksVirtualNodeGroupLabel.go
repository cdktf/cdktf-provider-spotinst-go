// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaksvirtualnodegroup


type OceanAksVirtualNodeGroupLabel struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.165.1/docs/resources/ocean_aks_virtual_node_group#key OceanAksVirtualNodeGroup#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.165.1/docs/resources/ocean_aks_virtual_node_group#value OceanAksVirtualNodeGroup#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

