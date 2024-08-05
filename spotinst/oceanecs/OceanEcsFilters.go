// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanecs


type OceanEcsFilters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.182.0/docs/resources/ocean_ecs#architectures OceanEcs#architectures}.
	Architectures *[]*string `field:"optional" json:"architectures" yaml:"architectures"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.182.0/docs/resources/ocean_ecs#categories OceanEcs#categories}.
	Categories *[]*string `field:"optional" json:"categories" yaml:"categories"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.182.0/docs/resources/ocean_ecs#disk_types OceanEcs#disk_types}.
	DiskTypes *[]*string `field:"optional" json:"diskTypes" yaml:"diskTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.182.0/docs/resources/ocean_ecs#exclude_families OceanEcs#exclude_families}.
	ExcludeFamilies *[]*string `field:"optional" json:"excludeFamilies" yaml:"excludeFamilies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.182.0/docs/resources/ocean_ecs#exclude_metal OceanEcs#exclude_metal}.
	ExcludeMetal interface{} `field:"optional" json:"excludeMetal" yaml:"excludeMetal"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.182.0/docs/resources/ocean_ecs#hypervisor OceanEcs#hypervisor}.
	Hypervisor *[]*string `field:"optional" json:"hypervisor" yaml:"hypervisor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.182.0/docs/resources/ocean_ecs#include_families OceanEcs#include_families}.
	IncludeFamilies *[]*string `field:"optional" json:"includeFamilies" yaml:"includeFamilies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.182.0/docs/resources/ocean_ecs#is_ena_supported OceanEcs#is_ena_supported}.
	IsEnaSupported *string `field:"optional" json:"isEnaSupported" yaml:"isEnaSupported"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.182.0/docs/resources/ocean_ecs#max_gpu OceanEcs#max_gpu}.
	MaxGpu *float64 `field:"optional" json:"maxGpu" yaml:"maxGpu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.182.0/docs/resources/ocean_ecs#max_memory_gib OceanEcs#max_memory_gib}.
	MaxMemoryGib *float64 `field:"optional" json:"maxMemoryGib" yaml:"maxMemoryGib"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.182.0/docs/resources/ocean_ecs#max_network_performance OceanEcs#max_network_performance}.
	MaxNetworkPerformance *float64 `field:"optional" json:"maxNetworkPerformance" yaml:"maxNetworkPerformance"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.182.0/docs/resources/ocean_ecs#max_vcpu OceanEcs#max_vcpu}.
	MaxVcpu *float64 `field:"optional" json:"maxVcpu" yaml:"maxVcpu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.182.0/docs/resources/ocean_ecs#min_enis OceanEcs#min_enis}.
	MinEnis *float64 `field:"optional" json:"minEnis" yaml:"minEnis"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.182.0/docs/resources/ocean_ecs#min_gpu OceanEcs#min_gpu}.
	MinGpu *float64 `field:"optional" json:"minGpu" yaml:"minGpu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.182.0/docs/resources/ocean_ecs#min_memory_gib OceanEcs#min_memory_gib}.
	MinMemoryGib *float64 `field:"optional" json:"minMemoryGib" yaml:"minMemoryGib"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.182.0/docs/resources/ocean_ecs#min_network_performance OceanEcs#min_network_performance}.
	MinNetworkPerformance *float64 `field:"optional" json:"minNetworkPerformance" yaml:"minNetworkPerformance"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.182.0/docs/resources/ocean_ecs#min_vcpu OceanEcs#min_vcpu}.
	MinVcpu *float64 `field:"optional" json:"minVcpu" yaml:"minVcpu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.182.0/docs/resources/ocean_ecs#root_device_types OceanEcs#root_device_types}.
	RootDeviceTypes *[]*string `field:"optional" json:"rootDeviceTypes" yaml:"rootDeviceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.182.0/docs/resources/ocean_ecs#virtualization_types OceanEcs#virtualization_types}.
	VirtualizationTypes *[]*string `field:"optional" json:"virtualizationTypes" yaml:"virtualizationTypes"`
}

