// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceangkeimport


type OceanGkeImportFilters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/ocean_gke_import#exclude_families OceanGkeImport#exclude_families}.
	ExcludeFamilies *[]*string `field:"optional" json:"excludeFamilies" yaml:"excludeFamilies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/ocean_gke_import#include_families OceanGkeImport#include_families}.
	IncludeFamilies *[]*string `field:"optional" json:"includeFamilies" yaml:"includeFamilies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/ocean_gke_import#max_memory_gib OceanGkeImport#max_memory_gib}.
	MaxMemoryGib *float64 `field:"optional" json:"maxMemoryGib" yaml:"maxMemoryGib"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/ocean_gke_import#max_vcpu OceanGkeImport#max_vcpu}.
	MaxVcpu *float64 `field:"optional" json:"maxVcpu" yaml:"maxVcpu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/ocean_gke_import#min_memory_gib OceanGkeImport#min_memory_gib}.
	MinMemoryGib *float64 `field:"optional" json:"minMemoryGib" yaml:"minMemoryGib"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/ocean_gke_import#min_vcpu OceanGkeImport#min_vcpu}.
	MinVcpu *float64 `field:"optional" json:"minVcpu" yaml:"minVcpu"`
}

