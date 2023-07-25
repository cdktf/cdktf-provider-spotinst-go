package elastigroupaws


type ElastigroupAwsEphemeralBlockDevice struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.128.0/docs/resources/elastigroup_aws#device_name ElastigroupAws#device_name}.
	DeviceName *string `field:"required" json:"deviceName" yaml:"deviceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.128.0/docs/resources/elastigroup_aws#virtual_name ElastigroupAws#virtual_name}.
	VirtualName *string `field:"required" json:"virtualName" yaml:"virtualName"`
}

