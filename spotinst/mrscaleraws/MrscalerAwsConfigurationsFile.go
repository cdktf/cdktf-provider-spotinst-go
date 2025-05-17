// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mrscaleraws


type MrscalerAwsConfigurationsFile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.220.0/docs/resources/mrscaler_aws#bucket MrscalerAws#bucket}.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.220.0/docs/resources/mrscaler_aws#key MrscalerAws#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
}

