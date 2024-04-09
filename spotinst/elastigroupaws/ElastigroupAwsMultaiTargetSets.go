// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupaws


type ElastigroupAwsMultaiTargetSets struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.169.0/docs/resources/elastigroup_aws#balancer_id ElastigroupAws#balancer_id}.
	BalancerId *string `field:"required" json:"balancerId" yaml:"balancerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.169.0/docs/resources/elastigroup_aws#target_set_id ElastigroupAws#target_set_id}.
	TargetSetId *string `field:"required" json:"targetSetId" yaml:"targetSetId"`
}

