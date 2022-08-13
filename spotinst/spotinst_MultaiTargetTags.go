// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type MultaiTargetTags struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/multai_target#key MultaiTarget#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/multai_target#value MultaiTarget#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

