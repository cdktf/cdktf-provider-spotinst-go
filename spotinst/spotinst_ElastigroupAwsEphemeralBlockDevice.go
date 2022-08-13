// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type ElastigroupAwsEphemeralBlockDevice struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_aws#device_name ElastigroupAws#device_name}.
	DeviceName *string `field:"required" json:"deviceName" yaml:"deviceName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_aws#virtual_name ElastigroupAws#virtual_name}.
	VirtualName *string `field:"required" json:"virtualName" yaml:"virtualName"`
}

