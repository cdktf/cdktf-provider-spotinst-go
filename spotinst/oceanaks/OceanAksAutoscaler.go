// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaks


type OceanAksAutoscaler struct {
	// autoscale_down block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.157.0/docs/resources/ocean_aks#autoscale_down OceanAks#autoscale_down}
	AutoscaleDown *OceanAksAutoscalerAutoscaleDown `field:"optional" json:"autoscaleDown" yaml:"autoscaleDown"`
	// autoscale_headroom block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.157.0/docs/resources/ocean_aks#autoscale_headroom OceanAks#autoscale_headroom}
	AutoscaleHeadroom *OceanAksAutoscalerAutoscaleHeadroom `field:"optional" json:"autoscaleHeadroom" yaml:"autoscaleHeadroom"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.157.0/docs/resources/ocean_aks#autoscale_is_enabled OceanAks#autoscale_is_enabled}.
	AutoscaleIsEnabled interface{} `field:"optional" json:"autoscaleIsEnabled" yaml:"autoscaleIsEnabled"`
	// resource_limits block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.157.0/docs/resources/ocean_aks#resource_limits OceanAks#resource_limits}
	ResourceLimits *OceanAksAutoscalerResourceLimits `field:"optional" json:"resourceLimits" yaml:"resourceLimits"`
}

