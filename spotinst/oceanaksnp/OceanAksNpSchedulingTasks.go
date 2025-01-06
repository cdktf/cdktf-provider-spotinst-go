// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaksnp


type OceanAksNpSchedulingTasks struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.205.0/docs/resources/ocean_aks_np#cron_expression OceanAksNp#cron_expression}.
	CronExpression *string `field:"required" json:"cronExpression" yaml:"cronExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.205.0/docs/resources/ocean_aks_np#is_enabled OceanAksNp#is_enabled}.
	IsEnabled interface{} `field:"required" json:"isEnabled" yaml:"isEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.205.0/docs/resources/ocean_aks_np#task_type OceanAksNp#task_type}.
	TaskType *string `field:"required" json:"taskType" yaml:"taskType"`
	// parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.205.0/docs/resources/ocean_aks_np#parameters OceanAksNp#parameters}
	Parameters *OceanAksNpSchedulingTasksParameters `field:"optional" json:"parameters" yaml:"parameters"`
}

