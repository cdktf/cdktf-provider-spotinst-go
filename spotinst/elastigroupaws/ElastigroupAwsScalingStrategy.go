// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupaws


type ElastigroupAwsScalingStrategy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.169.0/docs/resources/elastigroup_aws#terminate_at_end_of_billing_hour ElastigroupAws#terminate_at_end_of_billing_hour}.
	TerminateAtEndOfBillingHour interface{} `field:"optional" json:"terminateAtEndOfBillingHour" yaml:"terminateAtEndOfBillingHour"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.169.0/docs/resources/elastigroup_aws#termination_policy ElastigroupAws#termination_policy}.
	TerminationPolicy *string `field:"optional" json:"terminationPolicy" yaml:"terminationPolicy"`
}

