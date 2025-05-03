// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaksnp


type OceanAksNpAutoscalerAutoscaleHeadroomAutomatic struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/ocean_aks_np#is_enabled OceanAksNp#is_enabled}.
	IsEnabled interface{} `field:"optional" json:"isEnabled" yaml:"isEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/ocean_aks_np#percentage OceanAksNp#percentage}.
	Percentage *float64 `field:"optional" json:"percentage" yaml:"percentage"`
}

