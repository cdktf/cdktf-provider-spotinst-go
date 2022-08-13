// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type ElastigroupGcpIntegrationGkeAutoscaleLabels struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_gcp#key ElastigroupGcp#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_gcp#value ElastigroupGcp#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

