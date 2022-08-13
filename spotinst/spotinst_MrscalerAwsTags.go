// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type MrscalerAwsTags struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/mrscaler_aws#key MrscalerAws#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/mrscaler_aws#value MrscalerAws#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

