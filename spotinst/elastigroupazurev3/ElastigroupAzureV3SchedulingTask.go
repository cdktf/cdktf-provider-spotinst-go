// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupazurev3


type ElastigroupAzureV3SchedulingTask struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/elastigroup_azure_v3#cron_expression ElastigroupAzureV3#cron_expression}.
	CronExpression *string `field:"required" json:"cronExpression" yaml:"cronExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/elastigroup_azure_v3#is_enabled ElastigroupAzureV3#is_enabled}.
	IsEnabled interface{} `field:"required" json:"isEnabled" yaml:"isEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/elastigroup_azure_v3#type ElastigroupAzureV3#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/elastigroup_azure_v3#adjustment ElastigroupAzureV3#adjustment}.
	Adjustment *string `field:"optional" json:"adjustment" yaml:"adjustment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/elastigroup_azure_v3#adjustment_percentage ElastigroupAzureV3#adjustment_percentage}.
	AdjustmentPercentage *string `field:"optional" json:"adjustmentPercentage" yaml:"adjustmentPercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/elastigroup_azure_v3#batch_size_percentage ElastigroupAzureV3#batch_size_percentage}.
	BatchSizePercentage *string `field:"optional" json:"batchSizePercentage" yaml:"batchSizePercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/elastigroup_azure_v3#grace_period ElastigroupAzureV3#grace_period}.
	GracePeriod *string `field:"optional" json:"gracePeriod" yaml:"gracePeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/elastigroup_azure_v3#scale_max_capacity ElastigroupAzureV3#scale_max_capacity}.
	ScaleMaxCapacity *string `field:"optional" json:"scaleMaxCapacity" yaml:"scaleMaxCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/elastigroup_azure_v3#scale_min_capacity ElastigroupAzureV3#scale_min_capacity}.
	ScaleMinCapacity *string `field:"optional" json:"scaleMinCapacity" yaml:"scaleMinCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/elastigroup_azure_v3#scale_target_capacity ElastigroupAzureV3#scale_target_capacity}.
	ScaleTargetCapacity *string `field:"optional" json:"scaleTargetCapacity" yaml:"scaleTargetCapacity"`
}

