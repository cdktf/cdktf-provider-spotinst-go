// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaws


type OceanAwsScheduledTaskTasksParametersAmiAutoUpdateAmiAutoUpdateClusterRoll struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.176.0/docs/resources/ocean_aws#batch_min_healthy_percentage OceanAws#batch_min_healthy_percentage}.
	BatchMinHealthyPercentage *float64 `field:"optional" json:"batchMinHealthyPercentage" yaml:"batchMinHealthyPercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.176.0/docs/resources/ocean_aws#batch_size_percentage OceanAws#batch_size_percentage}.
	BatchSizePercentage *float64 `field:"optional" json:"batchSizePercentage" yaml:"batchSizePercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.176.0/docs/resources/ocean_aws#comment OceanAws#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.176.0/docs/resources/ocean_aws#respect_pdb OceanAws#respect_pdb}.
	RespectPdb interface{} `field:"optional" json:"respectPdb" yaml:"respectPdb"`
}

