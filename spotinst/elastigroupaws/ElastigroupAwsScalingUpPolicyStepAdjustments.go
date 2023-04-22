package elastigroupaws


type ElastigroupAwsScalingUpPolicyStepAdjustments struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.113.0/docs/resources/elastigroup_aws#action ElastigroupAws#action}
	Action *ElastigroupAwsScalingUpPolicyStepAdjustmentsAction `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.113.0/docs/resources/elastigroup_aws#threshold ElastigroupAws#threshold}.
	Threshold *float64 `field:"required" json:"threshold" yaml:"threshold"`
}

