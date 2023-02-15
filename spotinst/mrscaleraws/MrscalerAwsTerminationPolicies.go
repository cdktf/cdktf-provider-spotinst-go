package mrscaleraws


type MrscalerAwsTerminationPolicies struct {
	// statements block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/mrscaler_aws#statements MrscalerAws#statements}
	Statements interface{} `field:"required" json:"statements" yaml:"statements"`
}

