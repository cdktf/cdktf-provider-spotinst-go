// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceangkeimport


type OceanGkeImportScheduledTaskShutdownHours struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.159.0/docs/resources/ocean_gke_import#time_windows OceanGkeImport#time_windows}.
	TimeWindows *[]*string `field:"required" json:"timeWindows" yaml:"timeWindows"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.159.0/docs/resources/ocean_gke_import#is_enabled OceanGkeImport#is_enabled}.
	IsEnabled interface{} `field:"optional" json:"isEnabled" yaml:"isEnabled"`
}

