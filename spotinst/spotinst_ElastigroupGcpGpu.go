// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type ElastigroupGcpGpu struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_gcp#count ElastigroupGcp#count}.
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_gcp#type ElastigroupGcp#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

