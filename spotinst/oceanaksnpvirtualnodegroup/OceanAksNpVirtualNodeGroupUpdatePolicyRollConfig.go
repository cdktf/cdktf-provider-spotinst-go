// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaksnpvirtualnodegroup


type OceanAksNpVirtualNodeGroupUpdatePolicyRollConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.220.3/docs/resources/ocean_aks_np_virtual_node_group#batch_min_healthy_percentage OceanAksNpVirtualNodeGroup#batch_min_healthy_percentage}.
	BatchMinHealthyPercentage *float64 `field:"optional" json:"batchMinHealthyPercentage" yaml:"batchMinHealthyPercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.220.3/docs/resources/ocean_aks_np_virtual_node_group#batch_size_percentage OceanAksNpVirtualNodeGroup#batch_size_percentage}.
	BatchSizePercentage *float64 `field:"optional" json:"batchSizePercentage" yaml:"batchSizePercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.220.3/docs/resources/ocean_aks_np_virtual_node_group#comment OceanAksNpVirtualNodeGroup#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.220.3/docs/resources/ocean_aks_np_virtual_node_group#node_names OceanAksNpVirtualNodeGroup#node_names}.
	NodeNames *[]*string `field:"optional" json:"nodeNames" yaml:"nodeNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.220.3/docs/resources/ocean_aks_np_virtual_node_group#node_pool_names OceanAksNpVirtualNodeGroup#node_pool_names}.
	NodePoolNames *[]*string `field:"optional" json:"nodePoolNames" yaml:"nodePoolNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.220.3/docs/resources/ocean_aks_np_virtual_node_group#respect_pdb OceanAksNpVirtualNodeGroup#respect_pdb}.
	RespectPdb interface{} `field:"optional" json:"respectPdb" yaml:"respectPdb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.220.3/docs/resources/ocean_aks_np_virtual_node_group#respect_restrict_scale_down OceanAksNpVirtualNodeGroup#respect_restrict_scale_down}.
	RespectRestrictScaleDown interface{} `field:"optional" json:"respectRestrictScaleDown" yaml:"respectRestrictScaleDown"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.220.3/docs/resources/ocean_aks_np_virtual_node_group#vng_ids OceanAksNpVirtualNodeGroup#vng_ids}.
	VngIds *[]*string `field:"optional" json:"vngIds" yaml:"vngIds"`
}

