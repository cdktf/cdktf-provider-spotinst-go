// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupgke


type ElastigroupGkeDiskInitializeParams struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.188.0/docs/resources/elastigroup_gke#source_image ElastigroupGke#source_image}.
	SourceImage *string `field:"required" json:"sourceImage" yaml:"sourceImage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.188.0/docs/resources/elastigroup_gke#disk_size_gb ElastigroupGke#disk_size_gb}.
	DiskSizeGb *string `field:"optional" json:"diskSizeGb" yaml:"diskSizeGb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.188.0/docs/resources/elastigroup_gke#disk_type ElastigroupGke#disk_type}.
	DiskType *string `field:"optional" json:"diskType" yaml:"diskType"`
}

