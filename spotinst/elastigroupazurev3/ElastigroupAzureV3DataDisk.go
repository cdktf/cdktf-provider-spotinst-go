// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupazurev3


type ElastigroupAzureV3DataDisk struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.3/docs/resources/elastigroup_azure_v3#lun ElastigroupAzureV3#lun}.
	Lun *float64 `field:"required" json:"lun" yaml:"lun"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.3/docs/resources/elastigroup_azure_v3#size_gb ElastigroupAzureV3#size_gb}.
	SizeGb *float64 `field:"required" json:"sizeGb" yaml:"sizeGb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.3/docs/resources/elastigroup_azure_v3#type ElastigroupAzureV3#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

