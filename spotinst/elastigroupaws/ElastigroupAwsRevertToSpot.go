// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupaws


type ElastigroupAwsRevertToSpot struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.135.0/docs/resources/elastigroup_aws#perform_at ElastigroupAws#perform_at}.
	PerformAt *string `field:"required" json:"performAt" yaml:"performAt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.135.0/docs/resources/elastigroup_aws#time_windows ElastigroupAws#time_windows}.
	TimeWindows *[]*string `field:"optional" json:"timeWindows" yaml:"timeWindows"`
}

