// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaksvirtualnodegroup


type OceanAksVirtualNodeGroupLaunchSpecificationOsDisk struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.165.1/docs/resources/ocean_aks_virtual_node_group#size_gb OceanAksVirtualNodeGroup#size_gb}.
	SizeGb *float64 `field:"required" json:"sizeGb" yaml:"sizeGb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.165.1/docs/resources/ocean_aks_virtual_node_group#type OceanAksVirtualNodeGroup#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.165.1/docs/resources/ocean_aks_virtual_node_group#utilize_ephemeral_storage OceanAksVirtualNodeGroup#utilize_ephemeral_storage}.
	UtilizeEphemeralStorage interface{} `field:"optional" json:"utilizeEphemeralStorage" yaml:"utilizeEphemeralStorage"`
}

