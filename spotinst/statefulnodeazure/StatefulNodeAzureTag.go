// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package statefulnodeazure


type StatefulNodeAzureTag struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.2/docs/resources/stateful_node_azure#tag_key StatefulNodeAzure#tag_key}.
	TagKey *string `field:"required" json:"tagKey" yaml:"tagKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.2/docs/resources/stateful_node_azure#tag_value StatefulNodeAzure#tag_value}.
	TagValue *string `field:"optional" json:"tagValue" yaml:"tagValue"`
}

