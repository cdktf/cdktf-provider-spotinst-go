// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaks


type OceanAksAutoscalerAutoscaleHeadroomAutomatic struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.151.0/docs/resources/ocean_aks#is_enabled OceanAks#is_enabled}.
	IsEnabled interface{} `field:"optional" json:"isEnabled" yaml:"isEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.151.0/docs/resources/ocean_aks#percentage OceanAks#percentage}.
	Percentage *float64 `field:"optional" json:"percentage" yaml:"percentage"`
}

