// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type ElastigroupAwsScalingStrategy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_aws#terminate_at_end_of_billing_hour ElastigroupAws#terminate_at_end_of_billing_hour}.
	TerminateAtEndOfBillingHour interface{} `field:"optional" json:"terminateAtEndOfBillingHour" yaml:"terminateAtEndOfBillingHour"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_aws#termination_policy ElastigroupAws#termination_policy}.
	TerminationPolicy *string `field:"optional" json:"terminationPolicy" yaml:"terminationPolicy"`
}

