// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupaws


type ElastigroupAwsScalingDownPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.142.0/docs/resources/elastigroup_aws#metric_name ElastigroupAws#metric_name}.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.142.0/docs/resources/elastigroup_aws#namespace ElastigroupAws#namespace}.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.142.0/docs/resources/elastigroup_aws#policy_name ElastigroupAws#policy_name}.
	PolicyName *string `field:"required" json:"policyName" yaml:"policyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.142.0/docs/resources/elastigroup_aws#action_type ElastigroupAws#action_type}.
	ActionType *string `field:"optional" json:"actionType" yaml:"actionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.142.0/docs/resources/elastigroup_aws#adjustment ElastigroupAws#adjustment}.
	Adjustment *string `field:"optional" json:"adjustment" yaml:"adjustment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.142.0/docs/resources/elastigroup_aws#cooldown ElastigroupAws#cooldown}.
	Cooldown *float64 `field:"optional" json:"cooldown" yaml:"cooldown"`
	// dimensions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.142.0/docs/resources/elastigroup_aws#dimensions ElastigroupAws#dimensions}
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.142.0/docs/resources/elastigroup_aws#evaluation_periods ElastigroupAws#evaluation_periods}.
	EvaluationPeriods *float64 `field:"optional" json:"evaluationPeriods" yaml:"evaluationPeriods"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.142.0/docs/resources/elastigroup_aws#is_enabled ElastigroupAws#is_enabled}.
	IsEnabled interface{} `field:"optional" json:"isEnabled" yaml:"isEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.142.0/docs/resources/elastigroup_aws#maximum ElastigroupAws#maximum}.
	Maximum *string `field:"optional" json:"maximum" yaml:"maximum"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.142.0/docs/resources/elastigroup_aws#max_target_capacity ElastigroupAws#max_target_capacity}.
	MaxTargetCapacity *string `field:"optional" json:"maxTargetCapacity" yaml:"maxTargetCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.142.0/docs/resources/elastigroup_aws#minimum ElastigroupAws#minimum}.
	Minimum *string `field:"optional" json:"minimum" yaml:"minimum"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.142.0/docs/resources/elastigroup_aws#min_target_capacity ElastigroupAws#min_target_capacity}.
	MinTargetCapacity *string `field:"optional" json:"minTargetCapacity" yaml:"minTargetCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.142.0/docs/resources/elastigroup_aws#operator ElastigroupAws#operator}.
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.142.0/docs/resources/elastigroup_aws#period ElastigroupAws#period}.
	Period *float64 `field:"optional" json:"period" yaml:"period"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.142.0/docs/resources/elastigroup_aws#source ElastigroupAws#source}.
	Source *string `field:"optional" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.142.0/docs/resources/elastigroup_aws#statistic ElastigroupAws#statistic}.
	Statistic *string `field:"optional" json:"statistic" yaml:"statistic"`
	// step_adjustments block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.142.0/docs/resources/elastigroup_aws#step_adjustments ElastigroupAws#step_adjustments}
	StepAdjustments interface{} `field:"optional" json:"stepAdjustments" yaml:"stepAdjustments"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.142.0/docs/resources/elastigroup_aws#target ElastigroupAws#target}.
	Target *string `field:"optional" json:"target" yaml:"target"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.142.0/docs/resources/elastigroup_aws#threshold ElastigroupAws#threshold}.
	Threshold *float64 `field:"optional" json:"threshold" yaml:"threshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.142.0/docs/resources/elastigroup_aws#unit ElastigroupAws#unit}.
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

