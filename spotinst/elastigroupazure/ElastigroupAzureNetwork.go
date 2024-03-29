// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupazure


type ElastigroupAzureNetwork struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.165.1/docs/resources/elastigroup_azure#resource_group_name ElastigroupAzure#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.165.1/docs/resources/elastigroup_azure#subnet_name ElastigroupAzure#subnet_name}.
	SubnetName *string `field:"required" json:"subnetName" yaml:"subnetName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.165.1/docs/resources/elastigroup_azure#virtual_network_name ElastigroupAzure#virtual_network_name}.
	VirtualNetworkName *string `field:"required" json:"virtualNetworkName" yaml:"virtualNetworkName"`
	// additional_ip_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.165.1/docs/resources/elastigroup_azure#additional_ip_configs ElastigroupAzure#additional_ip_configs}
	AdditionalIpConfigs interface{} `field:"optional" json:"additionalIpConfigs" yaml:"additionalIpConfigs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.165.1/docs/resources/elastigroup_azure#assign_public_ip ElastigroupAzure#assign_public_ip}.
	AssignPublicIp interface{} `field:"optional" json:"assignPublicIp" yaml:"assignPublicIp"`
}

