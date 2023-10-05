// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaks


type OceanAksTag struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.143.0/docs/resources/ocean_aks#key OceanAks#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.143.0/docs/resources/ocean_aks#value OceanAks#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

