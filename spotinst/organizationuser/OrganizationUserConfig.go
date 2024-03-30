// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package organizationuser

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OrganizationUserConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.167.1/docs/resources/organization_user#email OrganizationUser#email}.
	Email *string `field:"required" json:"email" yaml:"email"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.167.1/docs/resources/organization_user#first_name OrganizationUser#first_name}.
	FirstName *string `field:"required" json:"firstName" yaml:"firstName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.167.1/docs/resources/organization_user#last_name OrganizationUser#last_name}.
	LastName *string `field:"required" json:"lastName" yaml:"lastName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.167.1/docs/resources/organization_user#id OrganizationUser#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.167.1/docs/resources/organization_user#password OrganizationUser#password}.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// policies block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.167.1/docs/resources/organization_user#policies OrganizationUser#policies}
	Policies interface{} `field:"optional" json:"policies" yaml:"policies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.167.1/docs/resources/organization_user#role OrganizationUser#role}.
	Role *string `field:"optional" json:"role" yaml:"role"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.167.1/docs/resources/organization_user#user_group_ids OrganizationUser#user_group_ids}.
	UserGroupIds *[]*string `field:"optional" json:"userGroupIds" yaml:"userGroupIds"`
}

