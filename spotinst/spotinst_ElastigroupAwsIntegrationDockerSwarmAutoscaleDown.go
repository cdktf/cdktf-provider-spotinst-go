// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type ElastigroupAwsIntegrationDockerSwarmAutoscaleDown struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_aws#evaluation_periods ElastigroupAws#evaluation_periods}.
	EvaluationPeriods *float64 `field:"optional" json:"evaluationPeriods" yaml:"evaluationPeriods"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_aws#max_scale_down_percentage ElastigroupAws#max_scale_down_percentage}.
	MaxScaleDownPercentage *float64 `field:"optional" json:"maxScaleDownPercentage" yaml:"maxScaleDownPercentage"`
}

