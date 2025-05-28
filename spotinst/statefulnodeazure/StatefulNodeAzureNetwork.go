// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package statefulnodeazure


type StatefulNodeAzureNetwork struct {
	// network_interface block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.220.2/docs/resources/stateful_node_azure#network_interface StatefulNodeAzure#network_interface}
	NetworkInterface interface{} `field:"required" json:"networkInterface" yaml:"networkInterface"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.220.2/docs/resources/stateful_node_azure#network_resource_group_name StatefulNodeAzure#network_resource_group_name}.
	NetworkResourceGroupName *string `field:"required" json:"networkResourceGroupName" yaml:"networkResourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.220.2/docs/resources/stateful_node_azure#virtual_network_name StatefulNodeAzure#virtual_network_name}.
	VirtualNetworkName *string `field:"required" json:"virtualNetworkName" yaml:"virtualNetworkName"`
}

