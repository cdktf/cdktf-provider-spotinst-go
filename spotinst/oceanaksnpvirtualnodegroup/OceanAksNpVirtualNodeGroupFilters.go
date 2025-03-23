// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaksnpvirtualnodegroup


type OceanAksNpVirtualNodeGroupFilters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.214.0/docs/resources/ocean_aks_np_virtual_node_group#accelerated_networking OceanAksNpVirtualNodeGroup#accelerated_networking}.
	AcceleratedNetworking *string `field:"optional" json:"acceleratedNetworking" yaml:"acceleratedNetworking"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.214.0/docs/resources/ocean_aks_np_virtual_node_group#architectures OceanAksNpVirtualNodeGroup#architectures}.
	Architectures *[]*string `field:"optional" json:"architectures" yaml:"architectures"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.214.0/docs/resources/ocean_aks_np_virtual_node_group#disk_performance OceanAksNpVirtualNodeGroup#disk_performance}.
	DiskPerformance *string `field:"optional" json:"diskPerformance" yaml:"diskPerformance"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.214.0/docs/resources/ocean_aks_np_virtual_node_group#exclude_series OceanAksNpVirtualNodeGroup#exclude_series}.
	ExcludeSeries *[]*string `field:"optional" json:"excludeSeries" yaml:"excludeSeries"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.214.0/docs/resources/ocean_aks_np_virtual_node_group#gpu_types OceanAksNpVirtualNodeGroup#gpu_types}.
	GpuTypes *[]*string `field:"optional" json:"gpuTypes" yaml:"gpuTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.214.0/docs/resources/ocean_aks_np_virtual_node_group#max_gpu OceanAksNpVirtualNodeGroup#max_gpu}.
	MaxGpu *float64 `field:"optional" json:"maxGpu" yaml:"maxGpu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.214.0/docs/resources/ocean_aks_np_virtual_node_group#max_memory_gib OceanAksNpVirtualNodeGroup#max_memory_gib}.
	MaxMemoryGib *float64 `field:"optional" json:"maxMemoryGib" yaml:"maxMemoryGib"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.214.0/docs/resources/ocean_aks_np_virtual_node_group#max_vcpu OceanAksNpVirtualNodeGroup#max_vcpu}.
	MaxVcpu *float64 `field:"optional" json:"maxVcpu" yaml:"maxVcpu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.214.0/docs/resources/ocean_aks_np_virtual_node_group#min_disk OceanAksNpVirtualNodeGroup#min_disk}.
	MinDisk *float64 `field:"optional" json:"minDisk" yaml:"minDisk"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.214.0/docs/resources/ocean_aks_np_virtual_node_group#min_gpu OceanAksNpVirtualNodeGroup#min_gpu}.
	MinGpu *float64 `field:"optional" json:"minGpu" yaml:"minGpu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.214.0/docs/resources/ocean_aks_np_virtual_node_group#min_memory_gib OceanAksNpVirtualNodeGroup#min_memory_gib}.
	MinMemoryGib *float64 `field:"optional" json:"minMemoryGib" yaml:"minMemoryGib"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.214.0/docs/resources/ocean_aks_np_virtual_node_group#min_nics OceanAksNpVirtualNodeGroup#min_nics}.
	MinNics *float64 `field:"optional" json:"minNics" yaml:"minNics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.214.0/docs/resources/ocean_aks_np_virtual_node_group#min_vcpu OceanAksNpVirtualNodeGroup#min_vcpu}.
	MinVcpu *float64 `field:"optional" json:"minVcpu" yaml:"minVcpu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.214.0/docs/resources/ocean_aks_np_virtual_node_group#series OceanAksNpVirtualNodeGroup#series}.
	Series *[]*string `field:"optional" json:"series" yaml:"series"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.214.0/docs/resources/ocean_aks_np_virtual_node_group#vm_types OceanAksNpVirtualNodeGroup#vm_types}.
	VmTypes *[]*string `field:"optional" json:"vmTypes" yaml:"vmTypes"`
}

