// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaks


type OceanAksNetwork struct {
	// network_interface block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.170.0/docs/resources/ocean_aks#network_interface OceanAks#network_interface}
	NetworkInterface interface{} `field:"optional" json:"networkInterface" yaml:"networkInterface"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.170.0/docs/resources/ocean_aks#resource_group_name OceanAks#resource_group_name}.
	ResourceGroupName *string `field:"optional" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.170.0/docs/resources/ocean_aks#virtual_network_name OceanAks#virtual_network_name}.
	VirtualNetworkName *string `field:"optional" json:"virtualNetworkName" yaml:"virtualNetworkName"`
}

