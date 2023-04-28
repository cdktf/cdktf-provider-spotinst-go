package mrscaleraws


type MrscalerAwsTaskScalingDownPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.115.0/docs/resources/mrscaler_aws#metric_name MrscalerAws#metric_name}.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.115.0/docs/resources/mrscaler_aws#namespace MrscalerAws#namespace}.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.115.0/docs/resources/mrscaler_aws#policy_name MrscalerAws#policy_name}.
	PolicyName *string `field:"required" json:"policyName" yaml:"policyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.115.0/docs/resources/mrscaler_aws#threshold MrscalerAws#threshold}.
	Threshold *float64 `field:"required" json:"threshold" yaml:"threshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.115.0/docs/resources/mrscaler_aws#unit MrscalerAws#unit}.
	Unit *string `field:"required" json:"unit" yaml:"unit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.115.0/docs/resources/mrscaler_aws#action_type MrscalerAws#action_type}.
	ActionType *string `field:"optional" json:"actionType" yaml:"actionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.115.0/docs/resources/mrscaler_aws#adjustment MrscalerAws#adjustment}.
	Adjustment *string `field:"optional" json:"adjustment" yaml:"adjustment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.115.0/docs/resources/mrscaler_aws#cooldown MrscalerAws#cooldown}.
	Cooldown *float64 `field:"optional" json:"cooldown" yaml:"cooldown"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.115.0/docs/resources/mrscaler_aws#dimensions MrscalerAws#dimensions}.
	Dimensions *map[string]*string `field:"optional" json:"dimensions" yaml:"dimensions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.115.0/docs/resources/mrscaler_aws#evaluation_periods MrscalerAws#evaluation_periods}.
	EvaluationPeriods *float64 `field:"optional" json:"evaluationPeriods" yaml:"evaluationPeriods"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.115.0/docs/resources/mrscaler_aws#maximum MrscalerAws#maximum}.
	Maximum *string `field:"optional" json:"maximum" yaml:"maximum"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.115.0/docs/resources/mrscaler_aws#max_target_capacity MrscalerAws#max_target_capacity}.
	MaxTargetCapacity *string `field:"optional" json:"maxTargetCapacity" yaml:"maxTargetCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.115.0/docs/resources/mrscaler_aws#minimum MrscalerAws#minimum}.
	Minimum *string `field:"optional" json:"minimum" yaml:"minimum"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.115.0/docs/resources/mrscaler_aws#min_target_capacity MrscalerAws#min_target_capacity}.
	MinTargetCapacity *string `field:"optional" json:"minTargetCapacity" yaml:"minTargetCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.115.0/docs/resources/mrscaler_aws#operator MrscalerAws#operator}.
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.115.0/docs/resources/mrscaler_aws#period MrscalerAws#period}.
	Period *float64 `field:"optional" json:"period" yaml:"period"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.115.0/docs/resources/mrscaler_aws#statistic MrscalerAws#statistic}.
	Statistic *string `field:"optional" json:"statistic" yaml:"statistic"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.115.0/docs/resources/mrscaler_aws#target MrscalerAws#target}.
	Target *string `field:"optional" json:"target" yaml:"target"`
}

