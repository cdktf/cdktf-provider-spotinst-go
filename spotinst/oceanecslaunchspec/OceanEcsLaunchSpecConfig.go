// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanecslaunchspec

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OceanEcsLaunchSpecConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/ocean_ecs_launch_spec#name OceanEcsLaunchSpec#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/ocean_ecs_launch_spec#ocean_id OceanEcsLaunchSpec#ocean_id}.
	OceanId *string `field:"required" json:"oceanId" yaml:"oceanId"`
	// attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/ocean_ecs_launch_spec#attributes OceanEcsLaunchSpec#attributes}
	Attributes interface{} `field:"optional" json:"attributes" yaml:"attributes"`
	// autoscale_headrooms block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/ocean_ecs_launch_spec#autoscale_headrooms OceanEcsLaunchSpec#autoscale_headrooms}
	AutoscaleHeadrooms interface{} `field:"optional" json:"autoscaleHeadrooms" yaml:"autoscaleHeadrooms"`
	// block_device_mappings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/ocean_ecs_launch_spec#block_device_mappings OceanEcsLaunchSpec#block_device_mappings}
	BlockDeviceMappings interface{} `field:"optional" json:"blockDeviceMappings" yaml:"blockDeviceMappings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/ocean_ecs_launch_spec#iam_instance_profile OceanEcsLaunchSpec#iam_instance_profile}.
	IamInstanceProfile *string `field:"optional" json:"iamInstanceProfile" yaml:"iamInstanceProfile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/ocean_ecs_launch_spec#id OceanEcsLaunchSpec#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/ocean_ecs_launch_spec#image_id OceanEcsLaunchSpec#image_id}.
	ImageId *string `field:"optional" json:"imageId" yaml:"imageId"`
	// images block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/ocean_ecs_launch_spec#images OceanEcsLaunchSpec#images}
	Images interface{} `field:"optional" json:"images" yaml:"images"`
	// instance_metadata_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/ocean_ecs_launch_spec#instance_metadata_options OceanEcsLaunchSpec#instance_metadata_options}
	InstanceMetadataOptions *OceanEcsLaunchSpecInstanceMetadataOptions `field:"optional" json:"instanceMetadataOptions" yaml:"instanceMetadataOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/ocean_ecs_launch_spec#instance_types OceanEcsLaunchSpec#instance_types}.
	InstanceTypes *[]*string `field:"optional" json:"instanceTypes" yaml:"instanceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/ocean_ecs_launch_spec#preferred_spot_types OceanEcsLaunchSpec#preferred_spot_types}.
	PreferredSpotTypes *[]*string `field:"optional" json:"preferredSpotTypes" yaml:"preferredSpotTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/ocean_ecs_launch_spec#restrict_scale_down OceanEcsLaunchSpec#restrict_scale_down}.
	RestrictScaleDown interface{} `field:"optional" json:"restrictScaleDown" yaml:"restrictScaleDown"`
	// scheduling_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/ocean_ecs_launch_spec#scheduling_task OceanEcsLaunchSpec#scheduling_task}
	SchedulingTask interface{} `field:"optional" json:"schedulingTask" yaml:"schedulingTask"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/ocean_ecs_launch_spec#security_group_ids OceanEcsLaunchSpec#security_group_ids}.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// strategy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/ocean_ecs_launch_spec#strategy OceanEcsLaunchSpec#strategy}
	Strategy interface{} `field:"optional" json:"strategy" yaml:"strategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/ocean_ecs_launch_spec#subnet_ids OceanEcsLaunchSpec#subnet_ids}.
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/ocean_ecs_launch_spec#tags OceanEcsLaunchSpec#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/ocean_ecs_launch_spec#user_data OceanEcsLaunchSpec#user_data}.
	UserData *string `field:"optional" json:"userData" yaml:"userData"`
}

