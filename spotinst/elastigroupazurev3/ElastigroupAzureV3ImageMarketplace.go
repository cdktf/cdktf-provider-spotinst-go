// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupazurev3


type ElastigroupAzureV3ImageMarketplace struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.172.3/docs/resources/elastigroup_azure_v3#offer ElastigroupAzureV3#offer}.
	Offer *string `field:"required" json:"offer" yaml:"offer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.172.3/docs/resources/elastigroup_azure_v3#publisher ElastigroupAzureV3#publisher}.
	Publisher *string `field:"required" json:"publisher" yaml:"publisher"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.172.3/docs/resources/elastigroup_azure_v3#sku ElastigroupAzureV3#sku}.
	Sku *string `field:"required" json:"sku" yaml:"sku"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.172.3/docs/resources/elastigroup_azure_v3#version ElastigroupAzureV3#version}.
	Version *string `field:"required" json:"version" yaml:"version"`
}

