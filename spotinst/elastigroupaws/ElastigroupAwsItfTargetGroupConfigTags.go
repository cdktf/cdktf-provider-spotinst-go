// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupaws


type ElastigroupAwsItfTargetGroupConfigTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/elastigroup_aws#tag_key ElastigroupAws#tag_key}.
	TagKey *string `field:"required" json:"tagKey" yaml:"tagKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/elastigroup_aws#tag_value ElastigroupAws#tag_value}.
	TagValue *string `field:"optional" json:"tagValue" yaml:"tagValue"`
}

