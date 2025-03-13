// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaws


type OceanAwsResourceTagSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.212.0/docs/resources/ocean_aws#should_tag_volumes OceanAws#should_tag_volumes}.
	ShouldTagVolumes interface{} `field:"optional" json:"shouldTagVolumes" yaml:"shouldTagVolumes"`
}

