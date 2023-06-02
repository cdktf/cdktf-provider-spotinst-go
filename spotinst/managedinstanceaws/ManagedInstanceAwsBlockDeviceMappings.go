package managedinstanceaws


type ManagedInstanceAwsBlockDeviceMappings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.121.0/docs/resources/managed_instance_aws#device_name ManagedInstanceAws#device_name}.
	DeviceName *string `field:"required" json:"deviceName" yaml:"deviceName"`
	// ebs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.121.0/docs/resources/managed_instance_aws#ebs ManagedInstanceAws#ebs}
	Ebs *ManagedInstanceAwsBlockDeviceMappingsEbs `field:"optional" json:"ebs" yaml:"ebs"`
}

