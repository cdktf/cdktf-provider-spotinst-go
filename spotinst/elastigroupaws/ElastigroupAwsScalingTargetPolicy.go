// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupaws


type ElastigroupAwsScalingTargetPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.168.2/docs/resources/elastigroup_aws#metric_name ElastigroupAws#metric_name}.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.168.2/docs/resources/elastigroup_aws#namespace ElastigroupAws#namespace}.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.168.2/docs/resources/elastigroup_aws#policy_name ElastigroupAws#policy_name}.
	PolicyName *string `field:"required" json:"policyName" yaml:"policyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.168.2/docs/resources/elastigroup_aws#target ElastigroupAws#target}.
	Target *float64 `field:"required" json:"target" yaml:"target"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.168.2/docs/resources/elastigroup_aws#cooldown ElastigroupAws#cooldown}.
	Cooldown *float64 `field:"optional" json:"cooldown" yaml:"cooldown"`
	// dimensions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.168.2/docs/resources/elastigroup_aws#dimensions ElastigroupAws#dimensions}
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.168.2/docs/resources/elastigroup_aws#evaluation_periods ElastigroupAws#evaluation_periods}.
	EvaluationPeriods *float64 `field:"optional" json:"evaluationPeriods" yaml:"evaluationPeriods"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.168.2/docs/resources/elastigroup_aws#max_capacity_per_scale ElastigroupAws#max_capacity_per_scale}.
	MaxCapacityPerScale *string `field:"optional" json:"maxCapacityPerScale" yaml:"maxCapacityPerScale"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.168.2/docs/resources/elastigroup_aws#period ElastigroupAws#period}.
	Period *float64 `field:"optional" json:"period" yaml:"period"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.168.2/docs/resources/elastigroup_aws#predictive_mode ElastigroupAws#predictive_mode}.
	PredictiveMode *string `field:"optional" json:"predictiveMode" yaml:"predictiveMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.168.2/docs/resources/elastigroup_aws#source ElastigroupAws#source}.
	Source *string `field:"optional" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.168.2/docs/resources/elastigroup_aws#statistic ElastigroupAws#statistic}.
	Statistic *string `field:"optional" json:"statistic" yaml:"statistic"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.168.2/docs/resources/elastigroup_aws#unit ElastigroupAws#unit}.
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

