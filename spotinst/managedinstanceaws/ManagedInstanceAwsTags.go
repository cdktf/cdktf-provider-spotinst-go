// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package managedinstanceaws


type ManagedInstanceAwsTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/managed_instance_aws#key ManagedInstanceAws#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/managed_instance_aws#value ManagedInstanceAws#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

