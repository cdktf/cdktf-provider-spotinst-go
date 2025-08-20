// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupazurev3


type ElastigroupAzureV3LoadBalancer struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.225.1/docs/resources/elastigroup_azure_v3#name ElastigroupAzureV3#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.225.1/docs/resources/elastigroup_azure_v3#resource_group_name ElastigroupAzureV3#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.225.1/docs/resources/elastigroup_azure_v3#type ElastigroupAzureV3#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.225.1/docs/resources/elastigroup_azure_v3#backend_pool_names ElastigroupAzureV3#backend_pool_names}.
	BackendPoolNames *[]*string `field:"optional" json:"backendPoolNames" yaml:"backendPoolNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.225.1/docs/resources/elastigroup_azure_v3#sku ElastigroupAzureV3#sku}.
	Sku *string `field:"optional" json:"sku" yaml:"sku"`
}

