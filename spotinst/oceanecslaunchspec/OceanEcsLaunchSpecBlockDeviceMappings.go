// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanecslaunchspec


type OceanEcsLaunchSpecBlockDeviceMappings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.207.0/docs/resources/ocean_ecs_launch_spec#device_name OceanEcsLaunchSpec#device_name}.
	DeviceName *string `field:"required" json:"deviceName" yaml:"deviceName"`
	// ebs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.207.0/docs/resources/ocean_ecs_launch_spec#ebs OceanEcsLaunchSpec#ebs}
	Ebs *OceanEcsLaunchSpecBlockDeviceMappingsEbs `field:"optional" json:"ebs" yaml:"ebs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.207.0/docs/resources/ocean_ecs_launch_spec#no_device OceanEcsLaunchSpec#no_device}.
	NoDevice *string `field:"optional" json:"noDevice" yaml:"noDevice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.207.0/docs/resources/ocean_ecs_launch_spec#virtual_name OceanEcsLaunchSpec#virtual_name}.
	VirtualName *string `field:"optional" json:"virtualName" yaml:"virtualName"`
}

