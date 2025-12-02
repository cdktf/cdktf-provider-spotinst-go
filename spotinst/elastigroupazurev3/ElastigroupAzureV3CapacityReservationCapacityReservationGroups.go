// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupazurev3


type ElastigroupAzureV3CapacityReservationCapacityReservationGroups struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.231.0/docs/resources/elastigroup_azure_v3#crg_name ElastigroupAzureV3#crg_name}.
	CrgName *string `field:"required" json:"crgName" yaml:"crgName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.231.0/docs/resources/elastigroup_azure_v3#crg_resource_group_name ElastigroupAzureV3#crg_resource_group_name}.
	CrgResourceGroupName *string `field:"required" json:"crgResourceGroupName" yaml:"crgResourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.231.0/docs/resources/elastigroup_azure_v3#crg_should_prioritize ElastigroupAzureV3#crg_should_prioritize}.
	CrgShouldPrioritize interface{} `field:"optional" json:"crgShouldPrioritize" yaml:"crgShouldPrioritize"`
}

