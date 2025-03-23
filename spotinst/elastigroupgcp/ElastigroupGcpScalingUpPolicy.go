// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupgcp


type ElastigroupGcpScalingUpPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.214.0/docs/resources/elastigroup_gcp#metric_name ElastigroupGcp#metric_name}.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.214.0/docs/resources/elastigroup_gcp#namespace ElastigroupGcp#namespace}.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.214.0/docs/resources/elastigroup_gcp#policy_name ElastigroupGcp#policy_name}.
	PolicyName *string `field:"required" json:"policyName" yaml:"policyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.214.0/docs/resources/elastigroup_gcp#threshold ElastigroupGcp#threshold}.
	Threshold *float64 `field:"required" json:"threshold" yaml:"threshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.214.0/docs/resources/elastigroup_gcp#unit ElastigroupGcp#unit}.
	Unit *string `field:"required" json:"unit" yaml:"unit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.214.0/docs/resources/elastigroup_gcp#action_type ElastigroupGcp#action_type}.
	ActionType *string `field:"optional" json:"actionType" yaml:"actionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.214.0/docs/resources/elastigroup_gcp#adjustment ElastigroupGcp#adjustment}.
	Adjustment *float64 `field:"optional" json:"adjustment" yaml:"adjustment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.214.0/docs/resources/elastigroup_gcp#cooldown ElastigroupGcp#cooldown}.
	Cooldown *float64 `field:"optional" json:"cooldown" yaml:"cooldown"`
	// dimensions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.214.0/docs/resources/elastigroup_gcp#dimensions ElastigroupGcp#dimensions}
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.214.0/docs/resources/elastigroup_gcp#evaluation_periods ElastigroupGcp#evaluation_periods}.
	EvaluationPeriods *float64 `field:"optional" json:"evaluationPeriods" yaml:"evaluationPeriods"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.214.0/docs/resources/elastigroup_gcp#operator ElastigroupGcp#operator}.
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.214.0/docs/resources/elastigroup_gcp#period ElastigroupGcp#period}.
	Period *float64 `field:"optional" json:"period" yaml:"period"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.214.0/docs/resources/elastigroup_gcp#source ElastigroupGcp#source}.
	Source *string `field:"optional" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.214.0/docs/resources/elastigroup_gcp#statistic ElastigroupGcp#statistic}.
	Statistic *string `field:"optional" json:"statistic" yaml:"statistic"`
}

