package elastigroupaws


type ElastigroupAwsMultipleMetrics struct {
	// expressions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.119.1/docs/resources/elastigroup_aws#expressions ElastigroupAws#expressions}
	Expressions interface{} `field:"optional" json:"expressions" yaml:"expressions"`
	// metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.119.1/docs/resources/elastigroup_aws#metrics ElastigroupAws#metrics}
	Metrics interface{} `field:"optional" json:"metrics" yaml:"metrics"`
}

