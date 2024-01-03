// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package organizationprogrammaticuser


type OrganizationProgrammaticUserAccounts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.157.0/docs/resources/organization_programmatic_user#account_id OrganizationProgrammaticUser#account_id}.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.157.0/docs/resources/organization_programmatic_user#account_role OrganizationProgrammaticUser#account_role}.
	AccountRole *string `field:"required" json:"accountRole" yaml:"accountRole"`
}

