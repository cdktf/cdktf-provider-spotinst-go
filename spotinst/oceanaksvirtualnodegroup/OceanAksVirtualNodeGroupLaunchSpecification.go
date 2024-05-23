// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaksvirtualnodegroup


type OceanAksVirtualNodeGroupLaunchSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.175.0/docs/resources/ocean_aks_virtual_node_group#max_pods OceanAksVirtualNodeGroup#max_pods}.
	MaxPods *float64 `field:"optional" json:"maxPods" yaml:"maxPods"`
	// os_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.175.0/docs/resources/ocean_aks_virtual_node_group#os_disk OceanAksVirtualNodeGroup#os_disk}
	OsDisk *OceanAksVirtualNodeGroupLaunchSpecificationOsDisk `field:"optional" json:"osDisk" yaml:"osDisk"`
	// tag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.175.0/docs/resources/ocean_aks_virtual_node_group#tag OceanAksVirtualNodeGroup#tag}
	Tag interface{} `field:"optional" json:"tag" yaml:"tag"`
}

