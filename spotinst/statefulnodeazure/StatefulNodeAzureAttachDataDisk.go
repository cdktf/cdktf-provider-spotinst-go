// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package statefulnodeazure


type StatefulNodeAzureAttachDataDisk struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.220.4/docs/resources/stateful_node_azure#data_disk_name StatefulNodeAzure#data_disk_name}.
	DataDiskName *string `field:"required" json:"dataDiskName" yaml:"dataDiskName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.220.4/docs/resources/stateful_node_azure#data_disk_resource_group_name StatefulNodeAzure#data_disk_resource_group_name}.
	DataDiskResourceGroupName *string `field:"required" json:"dataDiskResourceGroupName" yaml:"dataDiskResourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.220.4/docs/resources/stateful_node_azure#size_gb StatefulNodeAzure#size_gb}.
	SizeGb *float64 `field:"required" json:"sizeGb" yaml:"sizeGb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.220.4/docs/resources/stateful_node_azure#storage_account_type StatefulNodeAzure#storage_account_type}.
	StorageAccountType *string `field:"required" json:"storageAccountType" yaml:"storageAccountType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.220.4/docs/resources/stateful_node_azure#lun StatefulNodeAzure#lun}.
	Lun *float64 `field:"optional" json:"lun" yaml:"lun"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.220.4/docs/resources/stateful_node_azure#zone StatefulNodeAzure#zone}.
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

