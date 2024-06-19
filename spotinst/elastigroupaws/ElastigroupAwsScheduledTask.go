// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupaws


type ElastigroupAwsScheduledTask struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.177.0/docs/resources/elastigroup_aws#task_type ElastigroupAws#task_type}.
	TaskType *string `field:"required" json:"taskType" yaml:"taskType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.177.0/docs/resources/elastigroup_aws#adjustment ElastigroupAws#adjustment}.
	Adjustment *string `field:"optional" json:"adjustment" yaml:"adjustment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.177.0/docs/resources/elastigroup_aws#adjustment_percentage ElastigroupAws#adjustment_percentage}.
	AdjustmentPercentage *string `field:"optional" json:"adjustmentPercentage" yaml:"adjustmentPercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.177.0/docs/resources/elastigroup_aws#batch_size_percentage ElastigroupAws#batch_size_percentage}.
	BatchSizePercentage *string `field:"optional" json:"batchSizePercentage" yaml:"batchSizePercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.177.0/docs/resources/elastigroup_aws#cron_expression ElastigroupAws#cron_expression}.
	CronExpression *string `field:"optional" json:"cronExpression" yaml:"cronExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.177.0/docs/resources/elastigroup_aws#frequency ElastigroupAws#frequency}.
	Frequency *string `field:"optional" json:"frequency" yaml:"frequency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.177.0/docs/resources/elastigroup_aws#grace_period ElastigroupAws#grace_period}.
	GracePeriod *string `field:"optional" json:"gracePeriod" yaml:"gracePeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.177.0/docs/resources/elastigroup_aws#is_enabled ElastigroupAws#is_enabled}.
	IsEnabled interface{} `field:"optional" json:"isEnabled" yaml:"isEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.177.0/docs/resources/elastigroup_aws#max_capacity ElastigroupAws#max_capacity}.
	MaxCapacity *string `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.177.0/docs/resources/elastigroup_aws#min_capacity ElastigroupAws#min_capacity}.
	MinCapacity *string `field:"optional" json:"minCapacity" yaml:"minCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.177.0/docs/resources/elastigroup_aws#scale_max_capacity ElastigroupAws#scale_max_capacity}.
	ScaleMaxCapacity *string `field:"optional" json:"scaleMaxCapacity" yaml:"scaleMaxCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.177.0/docs/resources/elastigroup_aws#scale_min_capacity ElastigroupAws#scale_min_capacity}.
	ScaleMinCapacity *string `field:"optional" json:"scaleMinCapacity" yaml:"scaleMinCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.177.0/docs/resources/elastigroup_aws#scale_target_capacity ElastigroupAws#scale_target_capacity}.
	ScaleTargetCapacity *string `field:"optional" json:"scaleTargetCapacity" yaml:"scaleTargetCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.177.0/docs/resources/elastigroup_aws#start_time ElastigroupAws#start_time}.
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.177.0/docs/resources/elastigroup_aws#target_capacity ElastigroupAws#target_capacity}.
	TargetCapacity *string `field:"optional" json:"targetCapacity" yaml:"targetCapacity"`
}

