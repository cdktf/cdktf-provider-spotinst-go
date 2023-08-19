package oceanaws


type OceanAwsTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.134.0/docs/resources/ocean_aws#key OceanAws#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.134.0/docs/resources/ocean_aws#value OceanAws#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

