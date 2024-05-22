// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupaws


type ElastigroupAwsResourceTagSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.173.0/docs/resources/elastigroup_aws#should_tag_amis ElastigroupAws#should_tag_amis}.
	ShouldTagAmis interface{} `field:"optional" json:"shouldTagAmis" yaml:"shouldTagAmis"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.173.0/docs/resources/elastigroup_aws#should_tag_enis ElastigroupAws#should_tag_enis}.
	ShouldTagEnis interface{} `field:"optional" json:"shouldTagEnis" yaml:"shouldTagEnis"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.173.0/docs/resources/elastigroup_aws#should_tag_snapshots ElastigroupAws#should_tag_snapshots}.
	ShouldTagSnapshots interface{} `field:"optional" json:"shouldTagSnapshots" yaml:"shouldTagSnapshots"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.173.0/docs/resources/elastigroup_aws#should_tag_volumes ElastigroupAws#should_tag_volumes}.
	ShouldTagVolumes interface{} `field:"optional" json:"shouldTagVolumes" yaml:"shouldTagVolumes"`
}

