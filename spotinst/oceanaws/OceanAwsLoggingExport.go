package oceanaws


type OceanAwsLoggingExport struct {
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.113.0/docs/resources/ocean_aws#s3 OceanAws#s3}
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
}

