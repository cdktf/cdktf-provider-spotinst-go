// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupaws


type ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailure struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.220.4/docs/resources/elastigroup_aws#action_type ElastigroupAws#action_type}.
	ActionType *string `field:"required" json:"actionType" yaml:"actionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.220.4/docs/resources/elastigroup_aws#batch_num ElastigroupAws#batch_num}.
	BatchNum *float64 `field:"optional" json:"batchNum" yaml:"batchNum"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.220.4/docs/resources/elastigroup_aws#draining_timeout ElastigroupAws#draining_timeout}.
	DrainingTimeout *float64 `field:"optional" json:"drainingTimeout" yaml:"drainingTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.220.4/docs/resources/elastigroup_aws#should_decrement_target_capacity ElastigroupAws#should_decrement_target_capacity}.
	ShouldDecrementTargetCapacity interface{} `field:"optional" json:"shouldDecrementTargetCapacity" yaml:"shouldDecrementTargetCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.220.4/docs/resources/elastigroup_aws#should_handle_all_batches ElastigroupAws#should_handle_all_batches}.
	ShouldHandleAllBatches interface{} `field:"optional" json:"shouldHandleAllBatches" yaml:"shouldHandleAllBatches"`
}

