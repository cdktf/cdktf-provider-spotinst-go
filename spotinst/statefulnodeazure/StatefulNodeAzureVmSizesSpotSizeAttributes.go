// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package statefulnodeazure


type StatefulNodeAzureVmSizesSpotSizeAttributes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.216.0/docs/resources/stateful_node_azure#max_cpu StatefulNodeAzure#max_cpu}.
	MaxCpu *float64 `field:"optional" json:"maxCpu" yaml:"maxCpu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.216.0/docs/resources/stateful_node_azure#max_memory StatefulNodeAzure#max_memory}.
	MaxMemory *float64 `field:"optional" json:"maxMemory" yaml:"maxMemory"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.216.0/docs/resources/stateful_node_azure#max_storage StatefulNodeAzure#max_storage}.
	MaxStorage *float64 `field:"optional" json:"maxStorage" yaml:"maxStorage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.216.0/docs/resources/stateful_node_azure#min_cpu StatefulNodeAzure#min_cpu}.
	MinCpu *float64 `field:"optional" json:"minCpu" yaml:"minCpu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.216.0/docs/resources/stateful_node_azure#min_memory StatefulNodeAzure#min_memory}.
	MinMemory *float64 `field:"optional" json:"minMemory" yaml:"minMemory"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.216.0/docs/resources/stateful_node_azure#min_storage StatefulNodeAzure#min_storage}.
	MinStorage *float64 `field:"optional" json:"minStorage" yaml:"minStorage"`
}

