// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package organizationpolicy


type OrganizationPolicyPolicyContentStatements struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.189.0/docs/resources/organization_policy#actions OrganizationPolicy#actions}.
	Actions *[]*string `field:"required" json:"actions" yaml:"actions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.189.0/docs/resources/organization_policy#effect OrganizationPolicy#effect}.
	Effect *string `field:"required" json:"effect" yaml:"effect"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.189.0/docs/resources/organization_policy#resources OrganizationPolicy#resources}.
	Resources *[]*string `field:"required" json:"resources" yaml:"resources"`
}

