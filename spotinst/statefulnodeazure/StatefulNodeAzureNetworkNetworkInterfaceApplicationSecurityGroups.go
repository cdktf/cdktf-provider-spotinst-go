// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package statefulnodeazure


type StatefulNodeAzureNetworkNetworkInterfaceApplicationSecurityGroups struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.205.1/docs/resources/stateful_node_azure#name StatefulNodeAzure#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.205.1/docs/resources/stateful_node_azure#network_resource_group_name StatefulNodeAzure#network_resource_group_name}.
	NetworkResourceGroupName *string `field:"required" json:"networkResourceGroupName" yaml:"networkResourceGroupName"`
}

