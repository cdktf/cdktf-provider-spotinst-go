// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package organizationusergroup


type OrganizationUserGroupPolicies struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.213.0/docs/resources/organization_user_group#account_ids OrganizationUserGroup#account_ids}.
	AccountIds *[]*string `field:"required" json:"accountIds" yaml:"accountIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.213.0/docs/resources/organization_user_group#policy_id OrganizationUserGroup#policy_id}.
	PolicyId *string `field:"required" json:"policyId" yaml:"policyId"`
}

