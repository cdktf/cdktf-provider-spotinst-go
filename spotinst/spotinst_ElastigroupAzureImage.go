// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type ElastigroupAzureImage struct {
	// custom block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_azure#custom ElastigroupAzure#custom}
	Custom interface{} `field:"optional" json:"custom" yaml:"custom"`
	// marketplace block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_azure#marketplace ElastigroupAzure#marketplace}
	Marketplace interface{} `field:"optional" json:"marketplace" yaml:"marketplace"`
}
