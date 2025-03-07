// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanecs


type OceanEcsScheduledTask struct {
	// shutdown_hours block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.211.0/docs/resources/ocean_ecs#shutdown_hours OceanEcs#shutdown_hours}
	ShutdownHours *OceanEcsScheduledTaskShutdownHours `field:"optional" json:"shutdownHours" yaml:"shutdownHours"`
	// tasks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.211.0/docs/resources/ocean_ecs#tasks OceanEcs#tasks}
	Tasks interface{} `field:"optional" json:"tasks" yaml:"tasks"`
}

