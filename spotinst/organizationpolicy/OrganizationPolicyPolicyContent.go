// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package organizationpolicy


type OrganizationPolicyPolicyContent struct {
	// statements block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.151.0/docs/resources/organization_policy#statements OrganizationPolicy#statements}
	Statements interface{} `field:"required" json:"statements" yaml:"statements"`
}

