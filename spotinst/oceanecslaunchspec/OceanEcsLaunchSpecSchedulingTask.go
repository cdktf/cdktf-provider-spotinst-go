// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanecslaunchspec


type OceanEcsLaunchSpecSchedulingTask struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.209.1/docs/resources/ocean_ecs_launch_spec#cron_expression OceanEcsLaunchSpec#cron_expression}.
	CronExpression *string `field:"required" json:"cronExpression" yaml:"cronExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.209.1/docs/resources/ocean_ecs_launch_spec#is_enabled OceanEcsLaunchSpec#is_enabled}.
	IsEnabled interface{} `field:"required" json:"isEnabled" yaml:"isEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.209.1/docs/resources/ocean_ecs_launch_spec#task_type OceanEcsLaunchSpec#task_type}.
	TaskType *string `field:"required" json:"taskType" yaml:"taskType"`
	// task_headroom block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.209.1/docs/resources/ocean_ecs_launch_spec#task_headroom OceanEcsLaunchSpec#task_headroom}
	TaskHeadroom interface{} `field:"optional" json:"taskHeadroom" yaml:"taskHeadroom"`
}

