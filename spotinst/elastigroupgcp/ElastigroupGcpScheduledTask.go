// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupgcp


type ElastigroupGcpScheduledTask struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.171.4/docs/resources/elastigroup_gcp#task_type ElastigroupGcp#task_type}.
	TaskType *string `field:"required" json:"taskType" yaml:"taskType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.171.4/docs/resources/elastigroup_gcp#cron_expression ElastigroupGcp#cron_expression}.
	CronExpression *string `field:"optional" json:"cronExpression" yaml:"cronExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.171.4/docs/resources/elastigroup_gcp#is_enabled ElastigroupGcp#is_enabled}.
	IsEnabled interface{} `field:"optional" json:"isEnabled" yaml:"isEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.171.4/docs/resources/elastigroup_gcp#max_capacity ElastigroupGcp#max_capacity}.
	MaxCapacity *string `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.171.4/docs/resources/elastigroup_gcp#min_capacity ElastigroupGcp#min_capacity}.
	MinCapacity *string `field:"optional" json:"minCapacity" yaml:"minCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.171.4/docs/resources/elastigroup_gcp#target_capacity ElastigroupGcp#target_capacity}.
	TargetCapacity *string `field:"optional" json:"targetCapacity" yaml:"targetCapacity"`
}

