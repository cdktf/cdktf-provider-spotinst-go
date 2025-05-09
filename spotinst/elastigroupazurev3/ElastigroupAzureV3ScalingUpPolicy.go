// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupazurev3


type ElastigroupAzureV3ScalingUpPolicy struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.219.0/docs/resources/elastigroup_azure_v3#action ElastigroupAzureV3#action}
	Action *ElastigroupAzureV3ScalingUpPolicyAction `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.219.0/docs/resources/elastigroup_azure_v3#cooldown ElastigroupAzureV3#cooldown}.
	Cooldown *float64 `field:"required" json:"cooldown" yaml:"cooldown"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.219.0/docs/resources/elastigroup_azure_v3#evaluation_periods ElastigroupAzureV3#evaluation_periods}.
	EvaluationPeriods *float64 `field:"required" json:"evaluationPeriods" yaml:"evaluationPeriods"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.219.0/docs/resources/elastigroup_azure_v3#metric_name ElastigroupAzureV3#metric_name}.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.219.0/docs/resources/elastigroup_azure_v3#namespace ElastigroupAzureV3#namespace}.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.219.0/docs/resources/elastigroup_azure_v3#operator ElastigroupAzureV3#operator}.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.219.0/docs/resources/elastigroup_azure_v3#period ElastigroupAzureV3#period}.
	Period *float64 `field:"required" json:"period" yaml:"period"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.219.0/docs/resources/elastigroup_azure_v3#policy_name ElastigroupAzureV3#policy_name}.
	PolicyName *string `field:"required" json:"policyName" yaml:"policyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.219.0/docs/resources/elastigroup_azure_v3#statistic ElastigroupAzureV3#statistic}.
	Statistic *string `field:"required" json:"statistic" yaml:"statistic"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.219.0/docs/resources/elastigroup_azure_v3#threshold ElastigroupAzureV3#threshold}.
	Threshold *float64 `field:"required" json:"threshold" yaml:"threshold"`
	// dimensions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.219.0/docs/resources/elastigroup_azure_v3#dimensions ElastigroupAzureV3#dimensions}
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.219.0/docs/resources/elastigroup_azure_v3#is_enabled ElastigroupAzureV3#is_enabled}.
	IsEnabled interface{} `field:"optional" json:"isEnabled" yaml:"isEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.219.0/docs/resources/elastigroup_azure_v3#source ElastigroupAzureV3#source}.
	Source *string `field:"optional" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.219.0/docs/resources/elastigroup_azure_v3#unit ElastigroupAzureV3#unit}.
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

