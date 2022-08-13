// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type ElastigroupGkeGpu struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_gke#count ElastigroupGke#count}.
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_gke#type ElastigroupGke#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

