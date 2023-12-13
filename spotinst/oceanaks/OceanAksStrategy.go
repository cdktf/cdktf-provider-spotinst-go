// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaks


type OceanAksStrategy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.155.0/docs/resources/ocean_aks#fallback_to_ondemand OceanAks#fallback_to_ondemand}.
	FallbackToOndemand interface{} `field:"optional" json:"fallbackToOndemand" yaml:"fallbackToOndemand"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.155.0/docs/resources/ocean_aks#spot_percentage OceanAks#spot_percentage}.
	SpotPercentage *float64 `field:"optional" json:"spotPercentage" yaml:"spotPercentage"`
}

