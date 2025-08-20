// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupaws


type ElastigroupAwsIntegrationKubernetesAutoscaleDown struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.225.1/docs/resources/elastigroup_aws#evaluation_periods ElastigroupAws#evaluation_periods}.
	EvaluationPeriods *float64 `field:"optional" json:"evaluationPeriods" yaml:"evaluationPeriods"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.225.1/docs/resources/elastigroup_aws#max_scale_down_percentage ElastigroupAws#max_scale_down_percentage}.
	MaxScaleDownPercentage *float64 `field:"optional" json:"maxScaleDownPercentage" yaml:"maxScaleDownPercentage"`
}

