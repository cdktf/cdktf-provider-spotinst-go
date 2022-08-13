// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type MultaiTargetSetTags struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/multai_target_set#key MultaiTargetSet#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/multai_target_set#value MultaiTargetSet#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

