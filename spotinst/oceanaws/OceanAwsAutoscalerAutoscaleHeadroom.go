// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaws


type OceanAwsAutoscalerAutoscaleHeadroom struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.176.0/docs/resources/ocean_aws#cpu_per_unit OceanAws#cpu_per_unit}.
	CpuPerUnit *float64 `field:"optional" json:"cpuPerUnit" yaml:"cpuPerUnit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.176.0/docs/resources/ocean_aws#gpu_per_unit OceanAws#gpu_per_unit}.
	GpuPerUnit *float64 `field:"optional" json:"gpuPerUnit" yaml:"gpuPerUnit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.176.0/docs/resources/ocean_aws#memory_per_unit OceanAws#memory_per_unit}.
	MemoryPerUnit *float64 `field:"optional" json:"memoryPerUnit" yaml:"memoryPerUnit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.176.0/docs/resources/ocean_aws#num_of_units OceanAws#num_of_units}.
	NumOfUnits *float64 `field:"optional" json:"numOfUnits" yaml:"numOfUnits"`
}

