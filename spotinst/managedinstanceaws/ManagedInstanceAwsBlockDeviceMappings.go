package managedinstanceaws


type ManagedInstanceAwsBlockDeviceMappings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/managed_instance_aws#device_name ManagedInstanceAws#device_name}.
	DeviceName *string `field:"required" json:"deviceName" yaml:"deviceName"`
	// ebs block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/managed_instance_aws#ebs ManagedInstanceAws#ebs}
	Ebs *ManagedInstanceAwsBlockDeviceMappingsEbs `field:"optional" json:"ebs" yaml:"ebs"`
}

