// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package statefulnodeazure


type StatefulNodeAzureUpdateState struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.189.0/docs/resources/stateful_node_azure#state StatefulNodeAzure#state}.
	State *string `field:"required" json:"state" yaml:"state"`
}

