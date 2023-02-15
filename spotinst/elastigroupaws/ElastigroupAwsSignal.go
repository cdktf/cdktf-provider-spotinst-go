package elastigroupaws


type ElastigroupAwsSignal struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_aws#name ElastigroupAws#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_aws#timeout ElastigroupAws#timeout}.
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
}

