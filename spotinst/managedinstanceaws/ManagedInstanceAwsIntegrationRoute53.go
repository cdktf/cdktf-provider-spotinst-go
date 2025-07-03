// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package managedinstanceaws


type ManagedInstanceAwsIntegrationRoute53 struct {
	// domains block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.221.0/docs/resources/managed_instance_aws#domains ManagedInstanceAws#domains}
	Domains interface{} `field:"required" json:"domains" yaml:"domains"`
}

