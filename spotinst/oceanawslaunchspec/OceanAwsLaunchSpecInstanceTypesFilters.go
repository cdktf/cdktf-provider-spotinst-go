// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanawslaunchspec


type OceanAwsLaunchSpecInstanceTypesFilters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/ocean_aws_launch_spec#categories OceanAwsLaunchSpec#categories}.
	Categories *[]*string `field:"optional" json:"categories" yaml:"categories"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/ocean_aws_launch_spec#disk_types OceanAwsLaunchSpec#disk_types}.
	DiskTypes *[]*string `field:"optional" json:"diskTypes" yaml:"diskTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/ocean_aws_launch_spec#exclude_families OceanAwsLaunchSpec#exclude_families}.
	ExcludeFamilies *[]*string `field:"optional" json:"excludeFamilies" yaml:"excludeFamilies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/ocean_aws_launch_spec#exclude_metal OceanAwsLaunchSpec#exclude_metal}.
	ExcludeMetal interface{} `field:"optional" json:"excludeMetal" yaml:"excludeMetal"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/ocean_aws_launch_spec#hypervisor OceanAwsLaunchSpec#hypervisor}.
	Hypervisor *[]*string `field:"optional" json:"hypervisor" yaml:"hypervisor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/ocean_aws_launch_spec#include_families OceanAwsLaunchSpec#include_families}.
	IncludeFamilies *[]*string `field:"optional" json:"includeFamilies" yaml:"includeFamilies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/ocean_aws_launch_spec#is_ena_supported OceanAwsLaunchSpec#is_ena_supported}.
	IsEnaSupported *string `field:"optional" json:"isEnaSupported" yaml:"isEnaSupported"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/ocean_aws_launch_spec#max_gpu OceanAwsLaunchSpec#max_gpu}.
	MaxGpu *float64 `field:"optional" json:"maxGpu" yaml:"maxGpu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/ocean_aws_launch_spec#max_memory_gib OceanAwsLaunchSpec#max_memory_gib}.
	MaxMemoryGib *float64 `field:"optional" json:"maxMemoryGib" yaml:"maxMemoryGib"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/ocean_aws_launch_spec#max_network_performance OceanAwsLaunchSpec#max_network_performance}.
	MaxNetworkPerformance *float64 `field:"optional" json:"maxNetworkPerformance" yaml:"maxNetworkPerformance"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/ocean_aws_launch_spec#max_vcpu OceanAwsLaunchSpec#max_vcpu}.
	MaxVcpu *float64 `field:"optional" json:"maxVcpu" yaml:"maxVcpu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/ocean_aws_launch_spec#min_enis OceanAwsLaunchSpec#min_enis}.
	MinEnis *float64 `field:"optional" json:"minEnis" yaml:"minEnis"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/ocean_aws_launch_spec#min_gpu OceanAwsLaunchSpec#min_gpu}.
	MinGpu *float64 `field:"optional" json:"minGpu" yaml:"minGpu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/ocean_aws_launch_spec#min_memory_gib OceanAwsLaunchSpec#min_memory_gib}.
	MinMemoryGib *float64 `field:"optional" json:"minMemoryGib" yaml:"minMemoryGib"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/ocean_aws_launch_spec#min_network_performance OceanAwsLaunchSpec#min_network_performance}.
	MinNetworkPerformance *float64 `field:"optional" json:"minNetworkPerformance" yaml:"minNetworkPerformance"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/ocean_aws_launch_spec#min_vcpu OceanAwsLaunchSpec#min_vcpu}.
	MinVcpu *float64 `field:"optional" json:"minVcpu" yaml:"minVcpu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/ocean_aws_launch_spec#root_device_types OceanAwsLaunchSpec#root_device_types}.
	RootDeviceTypes *[]*string `field:"optional" json:"rootDeviceTypes" yaml:"rootDeviceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/ocean_aws_launch_spec#virtualization_types OceanAwsLaunchSpec#virtualization_types}.
	VirtualizationTypes *[]*string `field:"optional" json:"virtualizationTypes" yaml:"virtualizationTypes"`
}

