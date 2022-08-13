// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type ElastigroupAwsScalingUpPolicyStepAdjustments struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_aws#action ElastigroupAws#action}
	Action *ElastigroupAwsScalingUpPolicyStepAdjustmentsAction `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_aws#threshold ElastigroupAws#threshold}.
	Threshold *float64 `field:"required" json:"threshold" yaml:"threshold"`
}

