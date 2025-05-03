// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceangkeimport


type OceanGkeImportScheduledTask struct {
	// shutdown_hours block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/ocean_gke_import#shutdown_hours OceanGkeImport#shutdown_hours}
	ShutdownHours *OceanGkeImportScheduledTaskShutdownHours `field:"optional" json:"shutdownHours" yaml:"shutdownHours"`
	// tasks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/ocean_gke_import#tasks OceanGkeImport#tasks}
	Tasks interface{} `field:"optional" json:"tasks" yaml:"tasks"`
}

