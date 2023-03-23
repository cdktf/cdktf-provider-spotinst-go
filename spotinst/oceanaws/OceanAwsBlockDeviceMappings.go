package oceanaws


type OceanAwsBlockDeviceMappings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aws#device_name OceanAws#device_name}.
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// ebs block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aws#ebs OceanAws#ebs}
	Ebs *OceanAwsBlockDeviceMappingsEbs `field:"optional" json:"ebs" yaml:"ebs"`
}

