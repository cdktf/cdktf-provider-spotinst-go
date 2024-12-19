// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaksnp


type OceanAksNpFilters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.202.0/docs/resources/ocean_aks_np#accelerated_networking OceanAksNp#accelerated_networking}.
	AcceleratedNetworking *string `field:"optional" json:"acceleratedNetworking" yaml:"acceleratedNetworking"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.202.0/docs/resources/ocean_aks_np#architectures OceanAksNp#architectures}.
	Architectures *[]*string `field:"optional" json:"architectures" yaml:"architectures"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.202.0/docs/resources/ocean_aks_np#disk_performance OceanAksNp#disk_performance}.
	DiskPerformance *string `field:"optional" json:"diskPerformance" yaml:"diskPerformance"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.202.0/docs/resources/ocean_aks_np#exclude_series OceanAksNp#exclude_series}.
	ExcludeSeries *[]*string `field:"optional" json:"excludeSeries" yaml:"excludeSeries"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.202.0/docs/resources/ocean_aks_np#gpu_types OceanAksNp#gpu_types}.
	GpuTypes *[]*string `field:"optional" json:"gpuTypes" yaml:"gpuTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.202.0/docs/resources/ocean_aks_np#max_gpu OceanAksNp#max_gpu}.
	MaxGpu *float64 `field:"optional" json:"maxGpu" yaml:"maxGpu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.202.0/docs/resources/ocean_aks_np#max_memory_gib OceanAksNp#max_memory_gib}.
	MaxMemoryGib *float64 `field:"optional" json:"maxMemoryGib" yaml:"maxMemoryGib"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.202.0/docs/resources/ocean_aks_np#max_vcpu OceanAksNp#max_vcpu}.
	MaxVcpu *float64 `field:"optional" json:"maxVcpu" yaml:"maxVcpu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.202.0/docs/resources/ocean_aks_np#min_disk OceanAksNp#min_disk}.
	MinDisk *float64 `field:"optional" json:"minDisk" yaml:"minDisk"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.202.0/docs/resources/ocean_aks_np#min_gpu OceanAksNp#min_gpu}.
	MinGpu *float64 `field:"optional" json:"minGpu" yaml:"minGpu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.202.0/docs/resources/ocean_aks_np#min_memory_gib OceanAksNp#min_memory_gib}.
	MinMemoryGib *float64 `field:"optional" json:"minMemoryGib" yaml:"minMemoryGib"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.202.0/docs/resources/ocean_aks_np#min_nics OceanAksNp#min_nics}.
	MinNics *float64 `field:"optional" json:"minNics" yaml:"minNics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.202.0/docs/resources/ocean_aks_np#min_vcpu OceanAksNp#min_vcpu}.
	MinVcpu *float64 `field:"optional" json:"minVcpu" yaml:"minVcpu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.202.0/docs/resources/ocean_aks_np#series OceanAksNp#series}.
	Series *[]*string `field:"optional" json:"series" yaml:"series"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.202.0/docs/resources/ocean_aks_np#vm_types OceanAksNp#vm_types}.
	VmTypes *[]*string `field:"optional" json:"vmTypes" yaml:"vmTypes"`
}

