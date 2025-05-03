// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package statefulnodeazure


type StatefulNodeAzureNetworkNetworkInterface struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/stateful_node_azure#is_primary StatefulNodeAzure#is_primary}.
	IsPrimary interface{} `field:"required" json:"isPrimary" yaml:"isPrimary"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/stateful_node_azure#subnet_name StatefulNodeAzure#subnet_name}.
	SubnetName *string `field:"required" json:"subnetName" yaml:"subnetName"`
	// additional_ip_configurations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/stateful_node_azure#additional_ip_configurations StatefulNodeAzure#additional_ip_configurations}
	AdditionalIpConfigurations interface{} `field:"optional" json:"additionalIpConfigurations" yaml:"additionalIpConfigurations"`
	// application_security_groups block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/stateful_node_azure#application_security_groups StatefulNodeAzure#application_security_groups}
	ApplicationSecurityGroups interface{} `field:"optional" json:"applicationSecurityGroups" yaml:"applicationSecurityGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/stateful_node_azure#assign_public_ip StatefulNodeAzure#assign_public_ip}.
	AssignPublicIp interface{} `field:"optional" json:"assignPublicIp" yaml:"assignPublicIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/stateful_node_azure#enable_ip_forwarding StatefulNodeAzure#enable_ip_forwarding}.
	EnableIpForwarding interface{} `field:"optional" json:"enableIpForwarding" yaml:"enableIpForwarding"`
	// network_security_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/stateful_node_azure#network_security_group StatefulNodeAzure#network_security_group}
	NetworkSecurityGroup interface{} `field:"optional" json:"networkSecurityGroup" yaml:"networkSecurityGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/stateful_node_azure#private_ip_addresses StatefulNodeAzure#private_ip_addresses}.
	PrivateIpAddresses *[]*string `field:"optional" json:"privateIpAddresses" yaml:"privateIpAddresses"`
	// public_ips block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/stateful_node_azure#public_ips StatefulNodeAzure#public_ips}
	PublicIps interface{} `field:"optional" json:"publicIps" yaml:"publicIps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/stateful_node_azure#public_ip_sku StatefulNodeAzure#public_ip_sku}.
	PublicIpSku *string `field:"optional" json:"publicIpSku" yaml:"publicIpSku"`
}

