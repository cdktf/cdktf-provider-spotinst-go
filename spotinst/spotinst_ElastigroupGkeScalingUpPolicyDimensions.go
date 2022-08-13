// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type ElastigroupGkeScalingUpPolicyDimensions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_gke#name ElastigroupGke#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_gke#value ElastigroupGke#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

