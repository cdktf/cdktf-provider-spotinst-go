// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupazure


type ElastigroupAzureScheduledTask struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.153.1/docs/resources/elastigroup_azure#cron_expression ElastigroupAzure#cron_expression}.
	CronExpression *string `field:"required" json:"cronExpression" yaml:"cronExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.153.1/docs/resources/elastigroup_azure#task_type ElastigroupAzure#task_type}.
	TaskType *string `field:"required" json:"taskType" yaml:"taskType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.153.1/docs/resources/elastigroup_azure#adjustment ElastigroupAzure#adjustment}.
	Adjustment *string `field:"optional" json:"adjustment" yaml:"adjustment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.153.1/docs/resources/elastigroup_azure#adjustment_percentage ElastigroupAzure#adjustment_percentage}.
	AdjustmentPercentage *string `field:"optional" json:"adjustmentPercentage" yaml:"adjustmentPercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.153.1/docs/resources/elastigroup_azure#batch_size_percentage ElastigroupAzure#batch_size_percentage}.
	BatchSizePercentage *string `field:"optional" json:"batchSizePercentage" yaml:"batchSizePercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.153.1/docs/resources/elastigroup_azure#grace_period ElastigroupAzure#grace_period}.
	GracePeriod *string `field:"optional" json:"gracePeriod" yaml:"gracePeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.153.1/docs/resources/elastigroup_azure#is_enabled ElastigroupAzure#is_enabled}.
	IsEnabled interface{} `field:"optional" json:"isEnabled" yaml:"isEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.153.1/docs/resources/elastigroup_azure#scale_max_capacity ElastigroupAzure#scale_max_capacity}.
	ScaleMaxCapacity *string `field:"optional" json:"scaleMaxCapacity" yaml:"scaleMaxCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.153.1/docs/resources/elastigroup_azure#scale_min_capacity ElastigroupAzure#scale_min_capacity}.
	ScaleMinCapacity *string `field:"optional" json:"scaleMinCapacity" yaml:"scaleMinCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.153.1/docs/resources/elastigroup_azure#scale_target_capacity ElastigroupAzure#scale_target_capacity}.
	ScaleTargetCapacity *string `field:"optional" json:"scaleTargetCapacity" yaml:"scaleTargetCapacity"`
}

