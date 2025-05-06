// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package statefulnodeazure


type StatefulNodeAzureDetachDataDisk struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.1/docs/resources/stateful_node_azure#data_disk_name StatefulNodeAzure#data_disk_name}.
	DataDiskName *string `field:"required" json:"dataDiskName" yaml:"dataDiskName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.1/docs/resources/stateful_node_azure#data_disk_resource_group_name StatefulNodeAzure#data_disk_resource_group_name}.
	DataDiskResourceGroupName *string `field:"required" json:"dataDiskResourceGroupName" yaml:"dataDiskResourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.1/docs/resources/stateful_node_azure#should_deallocate StatefulNodeAzure#should_deallocate}.
	ShouldDeallocate interface{} `field:"required" json:"shouldDeallocate" yaml:"shouldDeallocate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.1/docs/resources/stateful_node_azure#ttl_in_hours StatefulNodeAzure#ttl_in_hours}.
	TtlInHours *float64 `field:"optional" json:"ttlInHours" yaml:"ttlInHours"`
}

