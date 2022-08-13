// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type OceanGkeImportAutoscalerDown struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_gke_import#evaluation_periods OceanGkeImport#evaluation_periods}.
	EvaluationPeriods *float64 `field:"optional" json:"evaluationPeriods" yaml:"evaluationPeriods"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_gke_import#max_scale_down_percentage OceanGkeImport#max_scale_down_percentage}.
	MaxScaleDownPercentage *float64 `field:"optional" json:"maxScaleDownPercentage" yaml:"maxScaleDownPercentage"`
}

