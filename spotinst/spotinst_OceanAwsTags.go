// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type OceanAwsTags struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aws#key OceanAws#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aws#value OceanAws#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

