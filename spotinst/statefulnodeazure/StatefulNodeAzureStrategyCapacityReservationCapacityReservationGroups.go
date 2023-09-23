// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package statefulnodeazure


type StatefulNodeAzureStrategyCapacityReservationCapacityReservationGroups struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.140.0/docs/resources/stateful_node_azure#crg_name StatefulNodeAzure#crg_name}.
	CrgName *string `field:"required" json:"crgName" yaml:"crgName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.140.0/docs/resources/stateful_node_azure#crg_resource_group_name StatefulNodeAzure#crg_resource_group_name}.
	CrgResourceGroupName *string `field:"required" json:"crgResourceGroupName" yaml:"crgResourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.140.0/docs/resources/stateful_node_azure#crg_should_prioritize StatefulNodeAzure#crg_should_prioritize}.
	CrgShouldPrioritize interface{} `field:"optional" json:"crgShouldPrioritize" yaml:"crgShouldPrioritize"`
}

