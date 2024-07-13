// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupaws


type ElastigroupAwsScalingDownPolicyStepAdjustments struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.181.0/docs/resources/elastigroup_aws#action ElastigroupAws#action}
	Action *ElastigroupAwsScalingDownPolicyStepAdjustmentsAction `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.181.0/docs/resources/elastigroup_aws#threshold ElastigroupAws#threshold}.
	Threshold *float64 `field:"required" json:"threshold" yaml:"threshold"`
}

