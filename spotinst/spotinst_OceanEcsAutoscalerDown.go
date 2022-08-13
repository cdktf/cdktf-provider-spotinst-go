// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type OceanEcsAutoscalerDown struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_ecs#max_scale_down_percentage OceanEcs#max_scale_down_percentage}.
	MaxScaleDownPercentage *float64 `field:"optional" json:"maxScaleDownPercentage" yaml:"maxScaleDownPercentage"`
}

