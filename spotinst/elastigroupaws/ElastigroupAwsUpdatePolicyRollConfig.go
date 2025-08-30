// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupaws


type ElastigroupAwsUpdatePolicyRollConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.226.0/docs/resources/elastigroup_aws#batch_size_percentage ElastigroupAws#batch_size_percentage}.
	BatchSizePercentage *float64 `field:"required" json:"batchSizePercentage" yaml:"batchSizePercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.226.0/docs/resources/elastigroup_aws#grace_period ElastigroupAws#grace_period}.
	GracePeriod *float64 `field:"optional" json:"gracePeriod" yaml:"gracePeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.226.0/docs/resources/elastigroup_aws#health_check_type ElastigroupAws#health_check_type}.
	HealthCheckType *string `field:"optional" json:"healthCheckType" yaml:"healthCheckType"`
	// strategy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.226.0/docs/resources/elastigroup_aws#strategy ElastigroupAws#strategy}
	Strategy *ElastigroupAwsUpdatePolicyRollConfigStrategy `field:"optional" json:"strategy" yaml:"strategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.226.0/docs/resources/elastigroup_aws#wait_for_roll_percentage ElastigroupAws#wait_for_roll_percentage}.
	WaitForRollPercentage *float64 `field:"optional" json:"waitForRollPercentage" yaml:"waitForRollPercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.226.0/docs/resources/elastigroup_aws#wait_for_roll_timeout ElastigroupAws#wait_for_roll_timeout}.
	WaitForRollTimeout *float64 `field:"optional" json:"waitForRollTimeout" yaml:"waitForRollTimeout"`
}

