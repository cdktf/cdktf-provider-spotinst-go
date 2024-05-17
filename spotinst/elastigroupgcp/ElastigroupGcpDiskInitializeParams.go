// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupgcp


type ElastigroupGcpDiskInitializeParams struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.172.2/docs/resources/elastigroup_gcp#source_image ElastigroupGcp#source_image}.
	SourceImage *string `field:"required" json:"sourceImage" yaml:"sourceImage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.172.2/docs/resources/elastigroup_gcp#disk_size_gb ElastigroupGcp#disk_size_gb}.
	DiskSizeGb *string `field:"optional" json:"diskSizeGb" yaml:"diskSizeGb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.172.2/docs/resources/elastigroup_gcp#disk_type ElastigroupGcp#disk_type}.
	DiskType *string `field:"optional" json:"diskType" yaml:"diskType"`
}

