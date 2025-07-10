// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanecs


type OceanEcsClusterOrientation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/ocean_ecs#availability_vs_cost OceanEcs#availability_vs_cost}.
	AvailabilityVsCost *string `field:"optional" json:"availabilityVsCost" yaml:"availabilityVsCost"`
}

