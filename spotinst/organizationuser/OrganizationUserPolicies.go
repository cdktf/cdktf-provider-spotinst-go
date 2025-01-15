// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package organizationuser


type OrganizationUserPolicies struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.207.0/docs/resources/organization_user#policy_account_ids OrganizationUser#policy_account_ids}.
	PolicyAccountIds *[]*string `field:"required" json:"policyAccountIds" yaml:"policyAccountIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.207.0/docs/resources/organization_user#policy_id OrganizationUser#policy_id}.
	PolicyId *string `field:"required" json:"policyId" yaml:"policyId"`
}

