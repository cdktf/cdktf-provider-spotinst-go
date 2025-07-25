// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package organizationprogrammaticuser


type OrganizationProgrammaticUserPolicies struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.224.1/docs/resources/organization_programmatic_user#policy_id OrganizationProgrammaticUser#policy_id}.
	PolicyId *string `field:"required" json:"policyId" yaml:"policyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.224.1/docs/resources/organization_programmatic_user#policy_account_ids OrganizationProgrammaticUser#policy_account_ids}.
	PolicyAccountIds *[]*string `field:"optional" json:"policyAccountIds" yaml:"policyAccountIds"`
}

