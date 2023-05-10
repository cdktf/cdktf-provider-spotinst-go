package elastigroupazure


type ElastigroupAzureScalingUpPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.117.0/docs/resources/elastigroup_azure#metric_name ElastigroupAzure#metric_name}.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.117.0/docs/resources/elastigroup_azure#namespace ElastigroupAzure#namespace}.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.117.0/docs/resources/elastigroup_azure#policy_name ElastigroupAzure#policy_name}.
	PolicyName *string `field:"required" json:"policyName" yaml:"policyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.117.0/docs/resources/elastigroup_azure#threshold ElastigroupAzure#threshold}.
	Threshold *float64 `field:"required" json:"threshold" yaml:"threshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.117.0/docs/resources/elastigroup_azure#action_type ElastigroupAzure#action_type}.
	ActionType *string `field:"optional" json:"actionType" yaml:"actionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.117.0/docs/resources/elastigroup_azure#adjustment ElastigroupAzure#adjustment}.
	Adjustment *string `field:"optional" json:"adjustment" yaml:"adjustment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.117.0/docs/resources/elastigroup_azure#cooldown ElastigroupAzure#cooldown}.
	Cooldown *float64 `field:"optional" json:"cooldown" yaml:"cooldown"`
	// dimensions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.117.0/docs/resources/elastigroup_azure#dimensions ElastigroupAzure#dimensions}
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.117.0/docs/resources/elastigroup_azure#evaluation_periods ElastigroupAzure#evaluation_periods}.
	EvaluationPeriods *float64 `field:"optional" json:"evaluationPeriods" yaml:"evaluationPeriods"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.117.0/docs/resources/elastigroup_azure#maximum ElastigroupAzure#maximum}.
	Maximum *string `field:"optional" json:"maximum" yaml:"maximum"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.117.0/docs/resources/elastigroup_azure#max_target_capacity ElastigroupAzure#max_target_capacity}.
	MaxTargetCapacity *string `field:"optional" json:"maxTargetCapacity" yaml:"maxTargetCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.117.0/docs/resources/elastigroup_azure#minimum ElastigroupAzure#minimum}.
	Minimum *string `field:"optional" json:"minimum" yaml:"minimum"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.117.0/docs/resources/elastigroup_azure#min_target_capacity ElastigroupAzure#min_target_capacity}.
	MinTargetCapacity *string `field:"optional" json:"minTargetCapacity" yaml:"minTargetCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.117.0/docs/resources/elastigroup_azure#operator ElastigroupAzure#operator}.
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.117.0/docs/resources/elastigroup_azure#period ElastigroupAzure#period}.
	Period *float64 `field:"optional" json:"period" yaml:"period"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.117.0/docs/resources/elastigroup_azure#statistic ElastigroupAzure#statistic}.
	Statistic *string `field:"optional" json:"statistic" yaml:"statistic"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.117.0/docs/resources/elastigroup_azure#target ElastigroupAzure#target}.
	Target *string `field:"optional" json:"target" yaml:"target"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.117.0/docs/resources/elastigroup_azure#unit ElastigroupAzure#unit}.
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

