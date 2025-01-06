// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package statefulnodeazure


type StatefulNodeAzureOsDisk struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.205.0/docs/resources/stateful_node_azure#type StatefulNodeAzure#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.205.0/docs/resources/stateful_node_azure#caching StatefulNodeAzure#caching}.
	Caching *string `field:"optional" json:"caching" yaml:"caching"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.205.0/docs/resources/stateful_node_azure#size_gb StatefulNodeAzure#size_gb}.
	SizeGb *float64 `field:"optional" json:"sizeGb" yaml:"sizeGb"`
}

