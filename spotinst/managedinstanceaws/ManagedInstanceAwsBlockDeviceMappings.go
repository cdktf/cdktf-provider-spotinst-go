// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package managedinstanceaws


type ManagedInstanceAwsBlockDeviceMappings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.220.0/docs/resources/managed_instance_aws#device_name ManagedInstanceAws#device_name}.
	DeviceName *string `field:"required" json:"deviceName" yaml:"deviceName"`
	// ebs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.220.0/docs/resources/managed_instance_aws#ebs ManagedInstanceAws#ebs}
	Ebs *ManagedInstanceAwsBlockDeviceMappingsEbs `field:"optional" json:"ebs" yaml:"ebs"`
}

