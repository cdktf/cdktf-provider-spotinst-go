// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package statefulnodeazure


type StatefulNodeAzureImageGallery struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.2/docs/resources/stateful_node_azure#gallery_name StatefulNodeAzure#gallery_name}.
	GalleryName *string `field:"required" json:"galleryName" yaml:"galleryName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.2/docs/resources/stateful_node_azure#gallery_resource_group_name StatefulNodeAzure#gallery_resource_group_name}.
	GalleryResourceGroupName *string `field:"required" json:"galleryResourceGroupName" yaml:"galleryResourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.2/docs/resources/stateful_node_azure#image_name StatefulNodeAzure#image_name}.
	ImageName *string `field:"required" json:"imageName" yaml:"imageName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.2/docs/resources/stateful_node_azure#version_name StatefulNodeAzure#version_name}.
	VersionName *string `field:"required" json:"versionName" yaml:"versionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.2/docs/resources/stateful_node_azure#spot_account_id StatefulNodeAzure#spot_account_id}.
	SpotAccountId *string `field:"optional" json:"spotAccountId" yaml:"spotAccountId"`
}

