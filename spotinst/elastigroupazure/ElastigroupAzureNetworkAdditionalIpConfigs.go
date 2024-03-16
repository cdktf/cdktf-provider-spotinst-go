// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupazure


type ElastigroupAzureNetworkAdditionalIpConfigs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/elastigroup_azure#name ElastigroupAzure#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/elastigroup_azure#private_ip_version ElastigroupAzure#private_ip_version}.
	PrivateIpVersion *string `field:"optional" json:"privateIpVersion" yaml:"privateIpVersion"`
}

