// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type MultaiListenerTags struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/multai_listener#key MultaiListener#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/multai_listener#value MultaiListener#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

