// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaws


type OceanAwsScheduledTaskTasks struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.181.1/docs/resources/ocean_aws#cron_expression OceanAws#cron_expression}.
	CronExpression *string `field:"required" json:"cronExpression" yaml:"cronExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.181.1/docs/resources/ocean_aws#is_enabled OceanAws#is_enabled}.
	IsEnabled interface{} `field:"required" json:"isEnabled" yaml:"isEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.181.1/docs/resources/ocean_aws#task_type OceanAws#task_type}.
	TaskType *string `field:"required" json:"taskType" yaml:"taskType"`
	// parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.181.1/docs/resources/ocean_aws#parameters OceanAws#parameters}
	Parameters *OceanAwsScheduledTaskTasksParameters `field:"optional" json:"parameters" yaml:"parameters"`
}

