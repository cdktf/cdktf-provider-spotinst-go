// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaws


type OceanAwsAutoscalerAutoscaleDown struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.136.0/docs/resources/ocean_aws#evaluation_periods OceanAws#evaluation_periods}.
	EvaluationPeriods *float64 `field:"optional" json:"evaluationPeriods" yaml:"evaluationPeriods"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.136.0/docs/resources/ocean_aws#max_scale_down_percentage OceanAws#max_scale_down_percentage}.
	MaxScaleDownPercentage *float64 `field:"optional" json:"maxScaleDownPercentage" yaml:"maxScaleDownPercentage"`
}

