// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaws


type OceanAwsStartupTaints struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.224.0/docs/resources/ocean_aws#effect OceanAws#effect}.
	Effect *string `field:"optional" json:"effect" yaml:"effect"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.224.0/docs/resources/ocean_aws#key OceanAws#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.224.0/docs/resources/ocean_aws#value OceanAws#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

