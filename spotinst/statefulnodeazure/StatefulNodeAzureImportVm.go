// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package statefulnodeazure


type StatefulNodeAzureImportVm struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.165.0/docs/resources/stateful_node_azure#original_vm_name StatefulNodeAzure#original_vm_name}.
	OriginalVmName *string `field:"required" json:"originalVmName" yaml:"originalVmName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.165.0/docs/resources/stateful_node_azure#resource_group_name StatefulNodeAzure#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.165.0/docs/resources/stateful_node_azure#draining_timeout StatefulNodeAzure#draining_timeout}.
	DrainingTimeout *float64 `field:"optional" json:"drainingTimeout" yaml:"drainingTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.165.0/docs/resources/stateful_node_azure#resources_retention_time StatefulNodeAzure#resources_retention_time}.
	ResourcesRetentionTime *float64 `field:"optional" json:"resourcesRetentionTime" yaml:"resourcesRetentionTime"`
}

