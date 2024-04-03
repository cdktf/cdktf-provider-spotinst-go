// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanecs


type OceanEcsTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.168.1/docs/resources/ocean_ecs#key OceanEcs#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.168.1/docs/resources/ocean_ecs#value OceanEcs#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

