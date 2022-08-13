// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type OceanEcsLoggingExport struct {
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_ecs#s3 OceanEcs#s3}
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
}

