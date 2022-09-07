// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type ElastigroupAwsImages struct {
	// image block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_aws#image ElastigroupAws#image}
	Image interface{} `field:"required" json:"image" yaml:"image"`
}

