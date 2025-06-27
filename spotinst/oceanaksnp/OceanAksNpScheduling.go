// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaksnp


type OceanAksNpScheduling struct {
	// shutdown_hours block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.220.4/docs/resources/ocean_aks_np#shutdown_hours OceanAksNp#shutdown_hours}
	ShutdownHours *OceanAksNpSchedulingShutdownHours `field:"optional" json:"shutdownHours" yaml:"shutdownHours"`
	// suspension_hours block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.220.4/docs/resources/ocean_aks_np#suspension_hours OceanAksNp#suspension_hours}
	SuspensionHours *OceanAksNpSchedulingSuspensionHours `field:"optional" json:"suspensionHours" yaml:"suspensionHours"`
	// tasks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.220.4/docs/resources/ocean_aks_np#tasks OceanAksNp#tasks}
	Tasks interface{} `field:"optional" json:"tasks" yaml:"tasks"`
}

