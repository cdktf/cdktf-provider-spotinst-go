// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package statefulnodeazure


type StatefulNodeAzureImageMarketplaceImage struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.176.1/docs/resources/stateful_node_azure#offer StatefulNodeAzure#offer}.
	Offer *string `field:"required" json:"offer" yaml:"offer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.176.1/docs/resources/stateful_node_azure#publisher StatefulNodeAzure#publisher}.
	Publisher *string `field:"required" json:"publisher" yaml:"publisher"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.176.1/docs/resources/stateful_node_azure#sku StatefulNodeAzure#sku}.
	Sku *string `field:"required" json:"sku" yaml:"sku"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.176.1/docs/resources/stateful_node_azure#version StatefulNodeAzure#version}.
	Version *string `field:"required" json:"version" yaml:"version"`
}

