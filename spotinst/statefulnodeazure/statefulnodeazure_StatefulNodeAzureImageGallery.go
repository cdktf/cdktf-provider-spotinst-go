package statefulnodeazure


type StatefulNodeAzureImageGallery struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/stateful_node_azure#gallery_name StatefulNodeAzure#gallery_name}.
	GalleryName *string `field:"required" json:"galleryName" yaml:"galleryName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/stateful_node_azure#gallery_resource_group_name StatefulNodeAzure#gallery_resource_group_name}.
	GalleryResourceGroupName *string `field:"required" json:"galleryResourceGroupName" yaml:"galleryResourceGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/stateful_node_azure#image_name StatefulNodeAzure#image_name}.
	ImageName *string `field:"required" json:"imageName" yaml:"imageName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/stateful_node_azure#version_name StatefulNodeAzure#version_name}.
	VersionName *string `field:"required" json:"versionName" yaml:"versionName"`
}
