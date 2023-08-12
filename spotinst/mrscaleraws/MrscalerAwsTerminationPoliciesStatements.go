package mrscaleraws


type MrscalerAwsTerminationPoliciesStatements struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.133.0/docs/resources/mrscaler_aws#metric_name MrscalerAws#metric_name}.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.133.0/docs/resources/mrscaler_aws#namespace MrscalerAws#namespace}.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.133.0/docs/resources/mrscaler_aws#threshold MrscalerAws#threshold}.
	Threshold *float64 `field:"required" json:"threshold" yaml:"threshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.133.0/docs/resources/mrscaler_aws#evaluation_periods MrscalerAws#evaluation_periods}.
	EvaluationPeriods *float64 `field:"optional" json:"evaluationPeriods" yaml:"evaluationPeriods"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.133.0/docs/resources/mrscaler_aws#operator MrscalerAws#operator}.
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.133.0/docs/resources/mrscaler_aws#period MrscalerAws#period}.
	Period *float64 `field:"optional" json:"period" yaml:"period"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.133.0/docs/resources/mrscaler_aws#statistic MrscalerAws#statistic}.
	Statistic *string `field:"optional" json:"statistic" yaml:"statistic"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.133.0/docs/resources/mrscaler_aws#unit MrscalerAws#unit}.
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

