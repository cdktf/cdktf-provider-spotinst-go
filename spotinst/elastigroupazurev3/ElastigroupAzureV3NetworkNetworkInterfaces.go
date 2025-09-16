// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupazurev3


type ElastigroupAzureV3NetworkNetworkInterfaces struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.228.0/docs/resources/elastigroup_azure_v3#assign_public_ip ElastigroupAzureV3#assign_public_ip}.
	AssignPublicIp interface{} `field:"required" json:"assignPublicIp" yaml:"assignPublicIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.228.0/docs/resources/elastigroup_azure_v3#is_primary ElastigroupAzureV3#is_primary}.
	IsPrimary interface{} `field:"required" json:"isPrimary" yaml:"isPrimary"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.228.0/docs/resources/elastigroup_azure_v3#subnet_name ElastigroupAzureV3#subnet_name}.
	SubnetName *string `field:"required" json:"subnetName" yaml:"subnetName"`
	// additional_ip_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.228.0/docs/resources/elastigroup_azure_v3#additional_ip_configs ElastigroupAzureV3#additional_ip_configs}
	AdditionalIpConfigs interface{} `field:"optional" json:"additionalIpConfigs" yaml:"additionalIpConfigs"`
	// application_security_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.228.0/docs/resources/elastigroup_azure_v3#application_security_group ElastigroupAzureV3#application_security_group}
	ApplicationSecurityGroup interface{} `field:"optional" json:"applicationSecurityGroup" yaml:"applicationSecurityGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.228.0/docs/resources/elastigroup_azure_v3#enable_ip_forwarding ElastigroupAzureV3#enable_ip_forwarding}.
	EnableIpForwarding interface{} `field:"optional" json:"enableIpForwarding" yaml:"enableIpForwarding"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.228.0/docs/resources/elastigroup_azure_v3#private_ip_addresses ElastigroupAzureV3#private_ip_addresses}.
	PrivateIpAddresses *[]*string `field:"optional" json:"privateIpAddresses" yaml:"privateIpAddresses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.228.0/docs/resources/elastigroup_azure_v3#public_ip_sku ElastigroupAzureV3#public_ip_sku}.
	PublicIpSku *string `field:"optional" json:"publicIpSku" yaml:"publicIpSku"`
	// security_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.228.0/docs/resources/elastigroup_azure_v3#security_group ElastigroupAzureV3#security_group}
	SecurityGroup interface{} `field:"optional" json:"securityGroup" yaml:"securityGroup"`
}

