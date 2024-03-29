// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupazure


type ElastigroupAzureImageMarketplace struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.165.1/docs/resources/elastigroup_azure#offer ElastigroupAzure#offer}.
	Offer *string `field:"required" json:"offer" yaml:"offer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.165.1/docs/resources/elastigroup_azure#publisher ElastigroupAzure#publisher}.
	Publisher *string `field:"required" json:"publisher" yaml:"publisher"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.165.1/docs/resources/elastigroup_azure#sku ElastigroupAzure#sku}.
	Sku *string `field:"required" json:"sku" yaml:"sku"`
}

