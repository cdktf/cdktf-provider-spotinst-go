package mrscaleraws


type MrscalerAwsStepsFile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.112.0/docs/resources/mrscaler_aws#bucket MrscalerAws#bucket}.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.112.0/docs/resources/mrscaler_aws#key MrscalerAws#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
}

