// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceangkelaunchspec


type OceanGkeLaunchSpecSchedulingTaskTaskHeadroom struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.157.0/docs/resources/ocean_gke_launch_spec#num_of_units OceanGkeLaunchSpec#num_of_units}.
	NumOfUnits *float64 `field:"required" json:"numOfUnits" yaml:"numOfUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.157.0/docs/resources/ocean_gke_launch_spec#cpu_per_unit OceanGkeLaunchSpec#cpu_per_unit}.
	CpuPerUnit *float64 `field:"optional" json:"cpuPerUnit" yaml:"cpuPerUnit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.157.0/docs/resources/ocean_gke_launch_spec#gpu_per_unit OceanGkeLaunchSpec#gpu_per_unit}.
	GpuPerUnit *float64 `field:"optional" json:"gpuPerUnit" yaml:"gpuPerUnit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.157.0/docs/resources/ocean_gke_launch_spec#memory_per_unit OceanGkeLaunchSpec#memory_per_unit}.
	MemoryPerUnit *float64 `field:"optional" json:"memoryPerUnit" yaml:"memoryPerUnit"`
}

