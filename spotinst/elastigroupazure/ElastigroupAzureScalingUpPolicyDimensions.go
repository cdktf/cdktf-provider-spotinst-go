package elastigroupazure


type ElastigroupAzureScalingUpPolicyDimensions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.0/docs/resources/elastigroup_azure#name ElastigroupAzure#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.0/docs/resources/elastigroup_azure#value ElastigroupAzure#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

