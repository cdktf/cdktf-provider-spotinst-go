// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package managedinstanceaws


type ManagedInstanceAwsLoadBalancers struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.0/docs/resources/managed_instance_aws#type ManagedInstanceAws#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.0/docs/resources/managed_instance_aws#arn ManagedInstanceAws#arn}.
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.0/docs/resources/managed_instance_aws#name ManagedInstanceAws#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

