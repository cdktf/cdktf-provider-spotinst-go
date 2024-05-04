// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mrscaleraws


type MrscalerAwsScheduledTask struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.171.3/docs/resources/mrscaler_aws#cron MrscalerAws#cron}.
	Cron *string `field:"required" json:"cron" yaml:"cron"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.171.3/docs/resources/mrscaler_aws#instance_group_type MrscalerAws#instance_group_type}.
	InstanceGroupType *string `field:"required" json:"instanceGroupType" yaml:"instanceGroupType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.171.3/docs/resources/mrscaler_aws#task_type MrscalerAws#task_type}.
	TaskType *string `field:"required" json:"taskType" yaml:"taskType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.171.3/docs/resources/mrscaler_aws#desired_capacity MrscalerAws#desired_capacity}.
	DesiredCapacity *string `field:"optional" json:"desiredCapacity" yaml:"desiredCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.171.3/docs/resources/mrscaler_aws#is_enabled MrscalerAws#is_enabled}.
	IsEnabled interface{} `field:"optional" json:"isEnabled" yaml:"isEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.171.3/docs/resources/mrscaler_aws#max_capacity MrscalerAws#max_capacity}.
	MaxCapacity *string `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.171.3/docs/resources/mrscaler_aws#min_capacity MrscalerAws#min_capacity}.
	MinCapacity *string `field:"optional" json:"minCapacity" yaml:"minCapacity"`
}

