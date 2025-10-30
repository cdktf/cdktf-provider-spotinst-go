// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanawslaunchspec


type OceanAwsLaunchSpecAutoscaleDown struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.230.0/docs/resources/ocean_aws_launch_spec#is_aggressive_scale_down_enabled OceanAwsLaunchSpec#is_aggressive_scale_down_enabled}.
	IsAggressiveScaleDownEnabled interface{} `field:"optional" json:"isAggressiveScaleDownEnabled" yaml:"isAggressiveScaleDownEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.230.0/docs/resources/ocean_aws_launch_spec#max_scale_down_percentage OceanAwsLaunchSpec#max_scale_down_percentage}.
	MaxScaleDownPercentage *float64 `field:"optional" json:"maxScaleDownPercentage" yaml:"maxScaleDownPercentage"`
}

