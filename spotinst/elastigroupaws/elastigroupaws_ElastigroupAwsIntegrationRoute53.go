package elastigroupaws


type ElastigroupAwsIntegrationRoute53 struct {
	// domains block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_aws#domains ElastigroupAws#domains}
	Domains interface{} `field:"required" json:"domains" yaml:"domains"`
}

