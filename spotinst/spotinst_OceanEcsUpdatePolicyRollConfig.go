// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type OceanEcsUpdatePolicyRollConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_ecs#batch_size_percentage OceanEcs#batch_size_percentage}.
	BatchSizePercentage *float64 `field:"required" json:"batchSizePercentage" yaml:"batchSizePercentage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_ecs#batch_min_healthy_percentage OceanEcs#batch_min_healthy_percentage}.
	BatchMinHealthyPercentage *float64 `field:"optional" json:"batchMinHealthyPercentage" yaml:"batchMinHealthyPercentage"`
}

