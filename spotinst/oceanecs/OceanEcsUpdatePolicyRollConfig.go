// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanecs


type OceanEcsUpdatePolicyRollConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.220.2/docs/resources/ocean_ecs#batch_size_percentage OceanEcs#batch_size_percentage}.
	BatchSizePercentage *float64 `field:"required" json:"batchSizePercentage" yaml:"batchSizePercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.220.2/docs/resources/ocean_ecs#batch_min_healthy_percentage OceanEcs#batch_min_healthy_percentage}.
	BatchMinHealthyPercentage *float64 `field:"optional" json:"batchMinHealthyPercentage" yaml:"batchMinHealthyPercentage"`
}

