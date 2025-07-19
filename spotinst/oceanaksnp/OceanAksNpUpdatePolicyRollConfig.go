// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaksnp


type OceanAksNpUpdatePolicyRollConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.224.0/docs/resources/ocean_aks_np#batch_min_healthy_percentage OceanAksNp#batch_min_healthy_percentage}.
	BatchMinHealthyPercentage *float64 `field:"optional" json:"batchMinHealthyPercentage" yaml:"batchMinHealthyPercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.224.0/docs/resources/ocean_aks_np#batch_size_percentage OceanAksNp#batch_size_percentage}.
	BatchSizePercentage *float64 `field:"optional" json:"batchSizePercentage" yaml:"batchSizePercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.224.0/docs/resources/ocean_aks_np#comment OceanAksNp#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.224.0/docs/resources/ocean_aks_np#node_names OceanAksNp#node_names}.
	NodeNames *[]*string `field:"optional" json:"nodeNames" yaml:"nodeNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.224.0/docs/resources/ocean_aks_np#node_pool_names OceanAksNp#node_pool_names}.
	NodePoolNames *[]*string `field:"optional" json:"nodePoolNames" yaml:"nodePoolNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.224.0/docs/resources/ocean_aks_np#respect_pdb OceanAksNp#respect_pdb}.
	RespectPdb interface{} `field:"optional" json:"respectPdb" yaml:"respectPdb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.224.0/docs/resources/ocean_aks_np#respect_restrict_scale_down OceanAksNp#respect_restrict_scale_down}.
	RespectRestrictScaleDown interface{} `field:"optional" json:"respectRestrictScaleDown" yaml:"respectRestrictScaleDown"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.224.0/docs/resources/ocean_aks_np#vng_ids OceanAksNp#vng_ids}.
	VngIds *[]*string `field:"optional" json:"vngIds" yaml:"vngIds"`
}

