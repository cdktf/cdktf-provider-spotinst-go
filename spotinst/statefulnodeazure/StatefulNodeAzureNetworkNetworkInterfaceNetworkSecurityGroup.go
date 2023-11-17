// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package statefulnodeazure


type StatefulNodeAzureNetworkNetworkInterfaceNetworkSecurityGroup struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.150.0/docs/resources/stateful_node_azure#name StatefulNodeAzure#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.150.0/docs/resources/stateful_node_azure#network_resource_group_name StatefulNodeAzure#network_resource_group_name}.
	NetworkResourceGroupName *string `field:"optional" json:"networkResourceGroupName" yaml:"networkResourceGroupName"`
}

