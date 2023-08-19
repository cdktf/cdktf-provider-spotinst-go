package elastigroupazure


type ElastigroupAzureImage struct {
	// custom block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.134.0/docs/resources/elastigroup_azure#custom ElastigroupAzure#custom}
	Custom interface{} `field:"optional" json:"custom" yaml:"custom"`
	// marketplace block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.134.0/docs/resources/elastigroup_azure#marketplace ElastigroupAzure#marketplace}
	Marketplace interface{} `field:"optional" json:"marketplace" yaml:"marketplace"`
}

