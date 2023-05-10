package mrscaleraws


type MrscalerAwsTerminationPolicies struct {
	// statements block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.117.0/docs/resources/mrscaler_aws#statements MrscalerAws#statements}
	Statements interface{} `field:"required" json:"statements" yaml:"statements"`
}

