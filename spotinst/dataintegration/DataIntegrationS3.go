// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataintegration


type DataIntegrationS3 struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.220.4/docs/resources/data_integration#bucket_name DataIntegration#bucket_name}.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.220.4/docs/resources/data_integration#subdir DataIntegration#subdir}.
	Subdir *string `field:"optional" json:"subdir" yaml:"subdir"`
}

