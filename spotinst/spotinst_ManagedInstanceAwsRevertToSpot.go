// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type ManagedInstanceAwsRevertToSpot struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/managed_instance_aws#perform_at ManagedInstanceAws#perform_at}.
	PerformAt *string `field:"required" json:"performAt" yaml:"performAt"`
}

