// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanawslaunchspec


type OceanAwsLaunchSpecBlockDeviceMappingsEbsDynamicVolumeSize struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.197.1/docs/resources/ocean_aws_launch_spec#base_size OceanAwsLaunchSpec#base_size}.
	BaseSize *float64 `field:"required" json:"baseSize" yaml:"baseSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.197.1/docs/resources/ocean_aws_launch_spec#resource OceanAwsLaunchSpec#resource}.
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.197.1/docs/resources/ocean_aws_launch_spec#size_per_resource_unit OceanAwsLaunchSpec#size_per_resource_unit}.
	SizePerResourceUnit *float64 `field:"required" json:"sizePerResourceUnit" yaml:"sizePerResourceUnit"`
}

