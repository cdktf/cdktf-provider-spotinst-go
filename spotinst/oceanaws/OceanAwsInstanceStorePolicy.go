// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaws


type OceanAwsInstanceStorePolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.224.0/docs/resources/ocean_aws#instance_store_policy_type OceanAws#instance_store_policy_type}.
	InstanceStorePolicyType *string `field:"optional" json:"instanceStorePolicyType" yaml:"instanceStorePolicyType"`
}

