// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package statefulnodeazure


type StatefulNodeAzureDelete struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.224.0/docs/resources/stateful_node_azure#should_terminate_vm StatefulNodeAzure#should_terminate_vm}.
	ShouldTerminateVm interface{} `field:"required" json:"shouldTerminateVm" yaml:"shouldTerminateVm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.224.0/docs/resources/stateful_node_azure#disk_should_deallocate StatefulNodeAzure#disk_should_deallocate}.
	DiskShouldDeallocate interface{} `field:"optional" json:"diskShouldDeallocate" yaml:"diskShouldDeallocate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.224.0/docs/resources/stateful_node_azure#disk_ttl_in_hours StatefulNodeAzure#disk_ttl_in_hours}.
	DiskTtlInHours *float64 `field:"optional" json:"diskTtlInHours" yaml:"diskTtlInHours"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.224.0/docs/resources/stateful_node_azure#network_should_deallocate StatefulNodeAzure#network_should_deallocate}.
	NetworkShouldDeallocate interface{} `field:"optional" json:"networkShouldDeallocate" yaml:"networkShouldDeallocate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.224.0/docs/resources/stateful_node_azure#network_ttl_in_hours StatefulNodeAzure#network_ttl_in_hours}.
	NetworkTtlInHours *float64 `field:"optional" json:"networkTtlInHours" yaml:"networkTtlInHours"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.224.0/docs/resources/stateful_node_azure#public_ip_should_deallocate StatefulNodeAzure#public_ip_should_deallocate}.
	PublicIpShouldDeallocate interface{} `field:"optional" json:"publicIpShouldDeallocate" yaml:"publicIpShouldDeallocate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.224.0/docs/resources/stateful_node_azure#public_ip_ttl_in_hours StatefulNodeAzure#public_ip_ttl_in_hours}.
	PublicIpTtlInHours *float64 `field:"optional" json:"publicIpTtlInHours" yaml:"publicIpTtlInHours"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.224.0/docs/resources/stateful_node_azure#should_deregister_from_lb StatefulNodeAzure#should_deregister_from_lb}.
	ShouldDeregisterFromLb interface{} `field:"optional" json:"shouldDeregisterFromLb" yaml:"shouldDeregisterFromLb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.224.0/docs/resources/stateful_node_azure#should_revert_to_od StatefulNodeAzure#should_revert_to_od}.
	ShouldRevertToOd interface{} `field:"optional" json:"shouldRevertToOd" yaml:"shouldRevertToOd"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.224.0/docs/resources/stateful_node_azure#snapshot_should_deallocate StatefulNodeAzure#snapshot_should_deallocate}.
	SnapshotShouldDeallocate interface{} `field:"optional" json:"snapshotShouldDeallocate" yaml:"snapshotShouldDeallocate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.224.0/docs/resources/stateful_node_azure#snapshot_ttl_in_hours StatefulNodeAzure#snapshot_ttl_in_hours}.
	SnapshotTtlInHours *float64 `field:"optional" json:"snapshotTtlInHours" yaml:"snapshotTtlInHours"`
}

