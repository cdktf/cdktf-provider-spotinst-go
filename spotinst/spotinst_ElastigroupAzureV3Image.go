// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type ElastigroupAzureV3Image struct {
	// custom block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_azure_v3#custom ElastigroupAzureV3#custom}
	Custom interface{} `field:"optional" json:"custom" yaml:"custom"`
	// marketplace block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_azure_v3#marketplace ElastigroupAzureV3#marketplace}
	Marketplace interface{} `field:"optional" json:"marketplace" yaml:"marketplace"`
}

