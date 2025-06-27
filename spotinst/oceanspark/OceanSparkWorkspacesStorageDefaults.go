// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanspark


type OceanSparkWorkspacesStorageDefaults struct {
	// The name of the persistent volume storage class to use by default for new workspaces.
	//
	// Leave it empty to use the cluster defaults.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.220.4/docs/resources/ocean_spark#storage_class_name OceanSpark#storage_class_name}
	StorageClassName *string `field:"optional" json:"storageClassName" yaml:"storageClassName"`
}

