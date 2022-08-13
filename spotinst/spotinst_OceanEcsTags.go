// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type OceanEcsTags struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_ecs#key OceanEcs#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_ecs#value OceanEcs#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

