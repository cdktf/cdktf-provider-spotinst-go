// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaws


type OceanAwsFilters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.230.0/docs/resources/ocean_aws#architectures OceanAws#architectures}.
	Architectures *[]*string `field:"optional" json:"architectures" yaml:"architectures"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.230.0/docs/resources/ocean_aws#categories OceanAws#categories}.
	Categories *[]*string `field:"optional" json:"categories" yaml:"categories"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.230.0/docs/resources/ocean_aws#disk_types OceanAws#disk_types}.
	DiskTypes *[]*string `field:"optional" json:"diskTypes" yaml:"diskTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.230.0/docs/resources/ocean_aws#exclude_families OceanAws#exclude_families}.
	ExcludeFamilies *[]*string `field:"optional" json:"excludeFamilies" yaml:"excludeFamilies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.230.0/docs/resources/ocean_aws#exclude_metal OceanAws#exclude_metal}.
	ExcludeMetal interface{} `field:"optional" json:"excludeMetal" yaml:"excludeMetal"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.230.0/docs/resources/ocean_aws#hypervisor OceanAws#hypervisor}.
	Hypervisor *[]*string `field:"optional" json:"hypervisor" yaml:"hypervisor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.230.0/docs/resources/ocean_aws#include_families OceanAws#include_families}.
	IncludeFamilies *[]*string `field:"optional" json:"includeFamilies" yaml:"includeFamilies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.230.0/docs/resources/ocean_aws#is_ena_supported OceanAws#is_ena_supported}.
	IsEnaSupported *string `field:"optional" json:"isEnaSupported" yaml:"isEnaSupported"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.230.0/docs/resources/ocean_aws#max_gpu OceanAws#max_gpu}.
	MaxGpu *float64 `field:"optional" json:"maxGpu" yaml:"maxGpu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.230.0/docs/resources/ocean_aws#max_memory_gib OceanAws#max_memory_gib}.
	MaxMemoryGib *float64 `field:"optional" json:"maxMemoryGib" yaml:"maxMemoryGib"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.230.0/docs/resources/ocean_aws#max_network_performance OceanAws#max_network_performance}.
	MaxNetworkPerformance *float64 `field:"optional" json:"maxNetworkPerformance" yaml:"maxNetworkPerformance"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.230.0/docs/resources/ocean_aws#max_vcpu OceanAws#max_vcpu}.
	MaxVcpu *float64 `field:"optional" json:"maxVcpu" yaml:"maxVcpu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.230.0/docs/resources/ocean_aws#min_enis OceanAws#min_enis}.
	MinEnis *float64 `field:"optional" json:"minEnis" yaml:"minEnis"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.230.0/docs/resources/ocean_aws#min_gpu OceanAws#min_gpu}.
	MinGpu *float64 `field:"optional" json:"minGpu" yaml:"minGpu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.230.0/docs/resources/ocean_aws#min_memory_gib OceanAws#min_memory_gib}.
	MinMemoryGib *float64 `field:"optional" json:"minMemoryGib" yaml:"minMemoryGib"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.230.0/docs/resources/ocean_aws#min_network_performance OceanAws#min_network_performance}.
	MinNetworkPerformance *float64 `field:"optional" json:"minNetworkPerformance" yaml:"minNetworkPerformance"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.230.0/docs/resources/ocean_aws#min_vcpu OceanAws#min_vcpu}.
	MinVcpu *float64 `field:"optional" json:"minVcpu" yaml:"minVcpu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.230.0/docs/resources/ocean_aws#root_device_types OceanAws#root_device_types}.
	RootDeviceTypes *[]*string `field:"optional" json:"rootDeviceTypes" yaml:"rootDeviceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.230.0/docs/resources/ocean_aws#virtualization_types OceanAws#virtualization_types}.
	VirtualizationTypes *[]*string `field:"optional" json:"virtualizationTypes" yaml:"virtualizationTypes"`
}

