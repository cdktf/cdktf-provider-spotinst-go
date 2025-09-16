// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceangkelaunchspec


type OceanGkeLaunchSpecFilters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.228.0/docs/resources/ocean_gke_launch_spec#exclude_families OceanGkeLaunchSpec#exclude_families}.
	ExcludeFamilies *[]*string `field:"optional" json:"excludeFamilies" yaml:"excludeFamilies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.228.0/docs/resources/ocean_gke_launch_spec#include_families OceanGkeLaunchSpec#include_families}.
	IncludeFamilies *[]*string `field:"optional" json:"includeFamilies" yaml:"includeFamilies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.228.0/docs/resources/ocean_gke_launch_spec#max_memory_gib OceanGkeLaunchSpec#max_memory_gib}.
	MaxMemoryGib *float64 `field:"optional" json:"maxMemoryGib" yaml:"maxMemoryGib"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.228.0/docs/resources/ocean_gke_launch_spec#max_vcpu OceanGkeLaunchSpec#max_vcpu}.
	MaxVcpu *float64 `field:"optional" json:"maxVcpu" yaml:"maxVcpu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.228.0/docs/resources/ocean_gke_launch_spec#min_memory_gib OceanGkeLaunchSpec#min_memory_gib}.
	MinMemoryGib *float64 `field:"optional" json:"minMemoryGib" yaml:"minMemoryGib"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.228.0/docs/resources/ocean_gke_launch_spec#min_vcpu OceanGkeLaunchSpec#min_vcpu}.
	MinVcpu *float64 `field:"optional" json:"minVcpu" yaml:"minVcpu"`
}

