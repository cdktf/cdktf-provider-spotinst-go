// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type OceanAwsLoggingExport struct {
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aws#s3 OceanAws#s3}
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
}

