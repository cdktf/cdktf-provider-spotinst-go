// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupgke


type ElastigroupGkeInstanceTypesCustom struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.201.0/docs/resources/elastigroup_gke#memory_gib ElastigroupGke#memory_gib}.
	MemoryGib *float64 `field:"required" json:"memoryGib" yaml:"memoryGib"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.201.0/docs/resources/elastigroup_gke#vcpu ElastigroupGke#vcpu}.
	Vcpu *float64 `field:"required" json:"vcpu" yaml:"vcpu"`
}

