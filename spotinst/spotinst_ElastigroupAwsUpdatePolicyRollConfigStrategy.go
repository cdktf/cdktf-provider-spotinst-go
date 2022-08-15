// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type ElastigroupAwsUpdatePolicyRollConfigStrategy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_aws#action ElastigroupAws#action}.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_aws#batch_min_healthy_percentage ElastigroupAws#batch_min_healthy_percentage}.
	BatchMinHealthyPercentage *float64 `field:"optional" json:"batchMinHealthyPercentage" yaml:"batchMinHealthyPercentage"`
	// on_failure block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_aws#on_failure ElastigroupAws#on_failure}
	OnFailure *ElastigroupAwsUpdatePolicyRollConfigStrategyOnFailure `field:"optional" json:"onFailure" yaml:"onFailure"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_aws#should_drain_instances ElastigroupAws#should_drain_instances}.
	ShouldDrainInstances interface{} `field:"optional" json:"shouldDrainInstances" yaml:"shouldDrainInstances"`
}
