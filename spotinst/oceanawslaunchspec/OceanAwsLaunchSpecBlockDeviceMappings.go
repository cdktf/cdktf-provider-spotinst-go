// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanawslaunchspec


type OceanAwsLaunchSpecBlockDeviceMappings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.207.0/docs/resources/ocean_aws_launch_spec#device_name OceanAwsLaunchSpec#device_name}.
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// ebs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.207.0/docs/resources/ocean_aws_launch_spec#ebs OceanAwsLaunchSpec#ebs}
	Ebs *OceanAwsLaunchSpecBlockDeviceMappingsEbs `field:"optional" json:"ebs" yaml:"ebs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.207.0/docs/resources/ocean_aws_launch_spec#no_device OceanAwsLaunchSpec#no_device}.
	NoDevice *string `field:"optional" json:"noDevice" yaml:"noDevice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.207.0/docs/resources/ocean_aws_launch_spec#virtual_name OceanAwsLaunchSpec#virtual_name}.
	VirtualName *string `field:"optional" json:"virtualName" yaml:"virtualName"`
}

