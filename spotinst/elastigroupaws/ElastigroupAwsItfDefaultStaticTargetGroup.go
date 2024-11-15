// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupaws


type ElastigroupAwsItfDefaultStaticTargetGroup struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.197.1/docs/resources/elastigroup_aws#arn ElastigroupAws#arn}.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.197.1/docs/resources/elastigroup_aws#percentage ElastigroupAws#percentage}.
	Percentage *float64 `field:"required" json:"percentage" yaml:"percentage"`
}

