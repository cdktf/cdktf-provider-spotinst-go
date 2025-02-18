// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupaws


type ElastigroupAwsIntegrationBeanstalkManagedActionsPlatformUpdate struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.209.2/docs/resources/elastigroup_aws#perform_at ElastigroupAws#perform_at}.
	PerformAt *string `field:"optional" json:"performAt" yaml:"performAt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.209.2/docs/resources/elastigroup_aws#time_window ElastigroupAws#time_window}.
	TimeWindow *string `field:"optional" json:"timeWindow" yaml:"timeWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.209.2/docs/resources/elastigroup_aws#update_level ElastigroupAws#update_level}.
	UpdateLevel *string `field:"optional" json:"updateLevel" yaml:"updateLevel"`
}

