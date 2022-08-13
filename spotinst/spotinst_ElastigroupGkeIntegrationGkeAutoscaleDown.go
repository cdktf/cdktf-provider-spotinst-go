// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type ElastigroupGkeIntegrationGkeAutoscaleDown struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_gke#evaluation_periods ElastigroupGke#evaluation_periods}.
	EvaluationPeriods *float64 `field:"optional" json:"evaluationPeriods" yaml:"evaluationPeriods"`
}

