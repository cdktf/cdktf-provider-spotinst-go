// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceangkeimport


type OceanGkeImportAutoscaler struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.3/docs/resources/ocean_gke_import#auto_headroom_percentage OceanGkeImport#auto_headroom_percentage}.
	AutoHeadroomPercentage *float64 `field:"optional" json:"autoHeadroomPercentage" yaml:"autoHeadroomPercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.3/docs/resources/ocean_gke_import#cooldown OceanGkeImport#cooldown}.
	Cooldown *float64 `field:"optional" json:"cooldown" yaml:"cooldown"`
	// down block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.3/docs/resources/ocean_gke_import#down OceanGkeImport#down}
	Down *OceanGkeImportAutoscalerDown `field:"optional" json:"down" yaml:"down"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.3/docs/resources/ocean_gke_import#enable_automatic_and_manual_headroom OceanGkeImport#enable_automatic_and_manual_headroom}.
	EnableAutomaticAndManualHeadroom interface{} `field:"optional" json:"enableAutomaticAndManualHeadroom" yaml:"enableAutomaticAndManualHeadroom"`
	// headroom block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.3/docs/resources/ocean_gke_import#headroom OceanGkeImport#headroom}
	Headroom *OceanGkeImportAutoscalerHeadroom `field:"optional" json:"headroom" yaml:"headroom"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.3/docs/resources/ocean_gke_import#is_auto_config OceanGkeImport#is_auto_config}.
	IsAutoConfig interface{} `field:"optional" json:"isAutoConfig" yaml:"isAutoConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.3/docs/resources/ocean_gke_import#is_enabled OceanGkeImport#is_enabled}.
	IsEnabled interface{} `field:"optional" json:"isEnabled" yaml:"isEnabled"`
	// resource_limits block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.3/docs/resources/ocean_gke_import#resource_limits OceanGkeImport#resource_limits}
	ResourceLimits *OceanGkeImportAutoscalerResourceLimits `field:"optional" json:"resourceLimits" yaml:"resourceLimits"`
}

