// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaws


type OceanAwsClusterOrientation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.151.0/docs/resources/ocean_aws#availability_vs_cost OceanAws#availability_vs_cost}.
	AvailabilityVsCost *string `field:"optional" json:"availabilityVsCost" yaml:"availabilityVsCost"`
}

