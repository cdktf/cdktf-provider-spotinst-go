// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package managedinstanceaws


type ManagedInstanceAwsRevertToSpot struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.198.0/docs/resources/managed_instance_aws#perform_at ManagedInstanceAws#perform_at}.
	PerformAt *string `field:"required" json:"performAt" yaml:"performAt"`
}

