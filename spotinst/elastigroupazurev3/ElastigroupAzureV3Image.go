// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupazurev3


type ElastigroupAzureV3Image struct {
	// custom block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.220.2/docs/resources/elastigroup_azure_v3#custom ElastigroupAzureV3#custom}
	Custom interface{} `field:"optional" json:"custom" yaml:"custom"`
	// gallery_image block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.220.2/docs/resources/elastigroup_azure_v3#gallery_image ElastigroupAzureV3#gallery_image}
	GalleryImage interface{} `field:"optional" json:"galleryImage" yaml:"galleryImage"`
	// marketplace block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.220.2/docs/resources/elastigroup_azure_v3#marketplace ElastigroupAzureV3#marketplace}
	Marketplace interface{} `field:"optional" json:"marketplace" yaml:"marketplace"`
}

