package mrscaleraws


type MrscalerAwsProvisioningTimeout struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.128.0/docs/resources/mrscaler_aws#timeout MrscalerAws#timeout}.
	Timeout *float64 `field:"required" json:"timeout" yaml:"timeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.128.0/docs/resources/mrscaler_aws#timeout_action MrscalerAws#timeout_action}.
	TimeoutAction *string `field:"required" json:"timeoutAction" yaml:"timeoutAction"`
}

