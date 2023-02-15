package elastigroupaws


type ElastigroupAwsInstanceTypesWeights struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_aws#instance_type ElastigroupAws#instance_type}.
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_aws#weight ElastigroupAws#weight}.
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
}

