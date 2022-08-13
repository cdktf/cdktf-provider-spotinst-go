// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type MrscalerAwsProvisioningTimeout struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/mrscaler_aws#timeout MrscalerAws#timeout}.
	Timeout *float64 `field:"required" json:"timeout" yaml:"timeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/mrscaler_aws#timeout_action MrscalerAws#timeout_action}.
	TimeoutAction *string `field:"required" json:"timeoutAction" yaml:"timeoutAction"`
}

