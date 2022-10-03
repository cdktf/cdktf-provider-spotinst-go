package mrscaleraws


type MrscalerAwsStepsFile struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/mrscaler_aws#bucket MrscalerAws#bucket}.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/mrscaler_aws#key MrscalerAws#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
}

