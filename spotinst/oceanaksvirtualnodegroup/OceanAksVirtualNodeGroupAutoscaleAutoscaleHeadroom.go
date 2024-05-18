// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaksvirtualnodegroup


type OceanAksVirtualNodeGroupAutoscaleAutoscaleHeadroom struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.172.3/docs/resources/ocean_aks_virtual_node_group#num_of_units OceanAksVirtualNodeGroup#num_of_units}.
	NumOfUnits *float64 `field:"required" json:"numOfUnits" yaml:"numOfUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.172.3/docs/resources/ocean_aks_virtual_node_group#cpu_per_unit OceanAksVirtualNodeGroup#cpu_per_unit}.
	CpuPerUnit *float64 `field:"optional" json:"cpuPerUnit" yaml:"cpuPerUnit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.172.3/docs/resources/ocean_aks_virtual_node_group#gpu_per_unit OceanAksVirtualNodeGroup#gpu_per_unit}.
	GpuPerUnit *float64 `field:"optional" json:"gpuPerUnit" yaml:"gpuPerUnit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.172.3/docs/resources/ocean_aks_virtual_node_group#memory_per_unit OceanAksVirtualNodeGroup#memory_per_unit}.
	MemoryPerUnit *float64 `field:"optional" json:"memoryPerUnit" yaml:"memoryPerUnit"`
}

