// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanawslaunchspec

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OceanAwsLaunchSpecConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/ocean_aws_launch_spec#ocean_id OceanAwsLaunchSpec#ocean_id}.
	OceanId *string `field:"required" json:"oceanId" yaml:"oceanId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/ocean_aws_launch_spec#associate_public_ip_address OceanAwsLaunchSpec#associate_public_ip_address}.
	AssociatePublicIpAddress interface{} `field:"optional" json:"associatePublicIpAddress" yaml:"associatePublicIpAddress"`
	// autoscale_down block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/ocean_aws_launch_spec#autoscale_down OceanAwsLaunchSpec#autoscale_down}
	AutoscaleDown interface{} `field:"optional" json:"autoscaleDown" yaml:"autoscaleDown"`
	// autoscale_headrooms block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/ocean_aws_launch_spec#autoscale_headrooms OceanAwsLaunchSpec#autoscale_headrooms}
	AutoscaleHeadrooms interface{} `field:"optional" json:"autoscaleHeadrooms" yaml:"autoscaleHeadrooms"`
	// autoscale_headrooms_automatic block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/ocean_aws_launch_spec#autoscale_headrooms_automatic OceanAwsLaunchSpec#autoscale_headrooms_automatic}
	AutoscaleHeadroomsAutomatic interface{} `field:"optional" json:"autoscaleHeadroomsAutomatic" yaml:"autoscaleHeadroomsAutomatic"`
	// block_device_mappings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/ocean_aws_launch_spec#block_device_mappings OceanAwsLaunchSpec#block_device_mappings}
	BlockDeviceMappings interface{} `field:"optional" json:"blockDeviceMappings" yaml:"blockDeviceMappings"`
	// create_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/ocean_aws_launch_spec#create_options OceanAwsLaunchSpec#create_options}
	CreateOptions *OceanAwsLaunchSpecCreateOptions `field:"optional" json:"createOptions" yaml:"createOptions"`
	// delete_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/ocean_aws_launch_spec#delete_options OceanAwsLaunchSpec#delete_options}
	DeleteOptions *OceanAwsLaunchSpecDeleteOptions `field:"optional" json:"deleteOptions" yaml:"deleteOptions"`
	// elastic_ip_pool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/ocean_aws_launch_spec#elastic_ip_pool OceanAwsLaunchSpec#elastic_ip_pool}
	ElasticIpPool interface{} `field:"optional" json:"elasticIpPool" yaml:"elasticIpPool"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/ocean_aws_launch_spec#iam_instance_profile OceanAwsLaunchSpec#iam_instance_profile}.
	IamInstanceProfile *string `field:"optional" json:"iamInstanceProfile" yaml:"iamInstanceProfile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/ocean_aws_launch_spec#id OceanAwsLaunchSpec#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/ocean_aws_launch_spec#image_id OceanAwsLaunchSpec#image_id}.
	ImageId *string `field:"optional" json:"imageId" yaml:"imageId"`
	// images block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/ocean_aws_launch_spec#images OceanAwsLaunchSpec#images}
	Images interface{} `field:"optional" json:"images" yaml:"images"`
	// instance_metadata_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/ocean_aws_launch_spec#instance_metadata_options OceanAwsLaunchSpec#instance_metadata_options}
	InstanceMetadataOptions *OceanAwsLaunchSpecInstanceMetadataOptions `field:"optional" json:"instanceMetadataOptions" yaml:"instanceMetadataOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/ocean_aws_launch_spec#instance_types OceanAwsLaunchSpec#instance_types}.
	InstanceTypes *[]*string `field:"optional" json:"instanceTypes" yaml:"instanceTypes"`
	// instance_types_filters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/ocean_aws_launch_spec#instance_types_filters OceanAwsLaunchSpec#instance_types_filters}
	InstanceTypesFilters *OceanAwsLaunchSpecInstanceTypesFilters `field:"optional" json:"instanceTypesFilters" yaml:"instanceTypesFilters"`
	// labels block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/ocean_aws_launch_spec#labels OceanAwsLaunchSpec#labels}
	Labels interface{} `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/ocean_aws_launch_spec#name OceanAwsLaunchSpec#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/ocean_aws_launch_spec#preferred_spot_types OceanAwsLaunchSpec#preferred_spot_types}.
	PreferredSpotTypes *[]*string `field:"optional" json:"preferredSpotTypes" yaml:"preferredSpotTypes"`
	// resource_limits block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/ocean_aws_launch_spec#resource_limits OceanAwsLaunchSpec#resource_limits}
	ResourceLimits interface{} `field:"optional" json:"resourceLimits" yaml:"resourceLimits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/ocean_aws_launch_spec#restrict_scale_down OceanAwsLaunchSpec#restrict_scale_down}.
	RestrictScaleDown interface{} `field:"optional" json:"restrictScaleDown" yaml:"restrictScaleDown"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/ocean_aws_launch_spec#root_volume_size OceanAwsLaunchSpec#root_volume_size}.
	RootVolumeSize *float64 `field:"optional" json:"rootVolumeSize" yaml:"rootVolumeSize"`
	// scheduling_shutdown_hours block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/ocean_aws_launch_spec#scheduling_shutdown_hours OceanAwsLaunchSpec#scheduling_shutdown_hours}
	SchedulingShutdownHours *OceanAwsLaunchSpecSchedulingShutdownHours `field:"optional" json:"schedulingShutdownHours" yaml:"schedulingShutdownHours"`
	// scheduling_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/ocean_aws_launch_spec#scheduling_task OceanAwsLaunchSpec#scheduling_task}
	SchedulingTask interface{} `field:"optional" json:"schedulingTask" yaml:"schedulingTask"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/ocean_aws_launch_spec#security_groups OceanAwsLaunchSpec#security_groups}.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// strategy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/ocean_aws_launch_spec#strategy OceanAwsLaunchSpec#strategy}
	Strategy interface{} `field:"optional" json:"strategy" yaml:"strategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/ocean_aws_launch_spec#subnet_ids OceanAwsLaunchSpec#subnet_ids}.
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/ocean_aws_launch_spec#tags OceanAwsLaunchSpec#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// taints block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/ocean_aws_launch_spec#taints OceanAwsLaunchSpec#taints}
	Taints interface{} `field:"optional" json:"taints" yaml:"taints"`
	// update_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/ocean_aws_launch_spec#update_policy OceanAwsLaunchSpec#update_policy}
	UpdatePolicy *OceanAwsLaunchSpecUpdatePolicy `field:"optional" json:"updatePolicy" yaml:"updatePolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/ocean_aws_launch_spec#user_data OceanAwsLaunchSpec#user_data}.
	UserData *string `field:"optional" json:"userData" yaml:"userData"`
}

