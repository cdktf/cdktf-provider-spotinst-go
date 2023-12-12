// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanecs

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OceanEcsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.153.1/docs/resources/ocean_ecs#cluster_name OceanEcs#cluster_name}.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.153.1/docs/resources/ocean_ecs#name OceanEcs#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.153.1/docs/resources/ocean_ecs#region OceanEcs#region}.
	Region *string `field:"required" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.153.1/docs/resources/ocean_ecs#security_group_ids OceanEcs#security_group_ids}.
	SecurityGroupIds *[]*string `field:"required" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.153.1/docs/resources/ocean_ecs#subnet_ids OceanEcs#subnet_ids}.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.153.1/docs/resources/ocean_ecs#associate_public_ip_address OceanEcs#associate_public_ip_address}.
	AssociatePublicIpAddress interface{} `field:"optional" json:"associatePublicIpAddress" yaml:"associatePublicIpAddress"`
	// autoscaler block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.153.1/docs/resources/ocean_ecs#autoscaler OceanEcs#autoscaler}
	Autoscaler *OceanEcsAutoscaler `field:"optional" json:"autoscaler" yaml:"autoscaler"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.153.1/docs/resources/ocean_ecs#blacklist OceanEcs#blacklist}.
	Blacklist *[]*string `field:"optional" json:"blacklist" yaml:"blacklist"`
	// block_device_mappings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.153.1/docs/resources/ocean_ecs#block_device_mappings OceanEcs#block_device_mappings}
	BlockDeviceMappings interface{} `field:"optional" json:"blockDeviceMappings" yaml:"blockDeviceMappings"`
	// cluster_orientation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.153.1/docs/resources/ocean_ecs#cluster_orientation OceanEcs#cluster_orientation}
	ClusterOrientation interface{} `field:"optional" json:"clusterOrientation" yaml:"clusterOrientation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.153.1/docs/resources/ocean_ecs#desired_capacity OceanEcs#desired_capacity}.
	DesiredCapacity *float64 `field:"optional" json:"desiredCapacity" yaml:"desiredCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.153.1/docs/resources/ocean_ecs#draining_timeout OceanEcs#draining_timeout}.
	DrainingTimeout *float64 `field:"optional" json:"drainingTimeout" yaml:"drainingTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.153.1/docs/resources/ocean_ecs#ebs_optimized OceanEcs#ebs_optimized}.
	EbsOptimized interface{} `field:"optional" json:"ebsOptimized" yaml:"ebsOptimized"`
	// filters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.153.1/docs/resources/ocean_ecs#filters OceanEcs#filters}
	Filters *OceanEcsFilters `field:"optional" json:"filters" yaml:"filters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.153.1/docs/resources/ocean_ecs#iam_instance_profile OceanEcs#iam_instance_profile}.
	IamInstanceProfile *string `field:"optional" json:"iamInstanceProfile" yaml:"iamInstanceProfile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.153.1/docs/resources/ocean_ecs#id OceanEcs#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.153.1/docs/resources/ocean_ecs#image_id OceanEcs#image_id}.
	ImageId *string `field:"optional" json:"imageId" yaml:"imageId"`
	// instance_metadata_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.153.1/docs/resources/ocean_ecs#instance_metadata_options OceanEcs#instance_metadata_options}
	InstanceMetadataOptions *OceanEcsInstanceMetadataOptions `field:"optional" json:"instanceMetadataOptions" yaml:"instanceMetadataOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.153.1/docs/resources/ocean_ecs#key_pair OceanEcs#key_pair}.
	KeyPair *string `field:"optional" json:"keyPair" yaml:"keyPair"`
	// logging block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.153.1/docs/resources/ocean_ecs#logging OceanEcs#logging}
	Logging *OceanEcsLogging `field:"optional" json:"logging" yaml:"logging"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.153.1/docs/resources/ocean_ecs#max_size OceanEcs#max_size}.
	MaxSize *float64 `field:"optional" json:"maxSize" yaml:"maxSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.153.1/docs/resources/ocean_ecs#min_size OceanEcs#min_size}.
	MinSize *float64 `field:"optional" json:"minSize" yaml:"minSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.153.1/docs/resources/ocean_ecs#monitoring OceanEcs#monitoring}.
	Monitoring interface{} `field:"optional" json:"monitoring" yaml:"monitoring"`
	// optimize_images block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.153.1/docs/resources/ocean_ecs#optimize_images OceanEcs#optimize_images}
	OptimizeImages *OceanEcsOptimizeImages `field:"optional" json:"optimizeImages" yaml:"optimizeImages"`
	// scheduled_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.153.1/docs/resources/ocean_ecs#scheduled_task OceanEcs#scheduled_task}
	ScheduledTask interface{} `field:"optional" json:"scheduledTask" yaml:"scheduledTask"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.153.1/docs/resources/ocean_ecs#spot_percentage OceanEcs#spot_percentage}.
	SpotPercentage *float64 `field:"optional" json:"spotPercentage" yaml:"spotPercentage"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.153.1/docs/resources/ocean_ecs#tags OceanEcs#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// update_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.153.1/docs/resources/ocean_ecs#update_policy OceanEcs#update_policy}
	UpdatePolicy *OceanEcsUpdatePolicy `field:"optional" json:"updatePolicy" yaml:"updatePolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.153.1/docs/resources/ocean_ecs#use_as_template_only OceanEcs#use_as_template_only}.
	UseAsTemplateOnly interface{} `field:"optional" json:"useAsTemplateOnly" yaml:"useAsTemplateOnly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.153.1/docs/resources/ocean_ecs#user_data OceanEcs#user_data}.
	UserData *string `field:"optional" json:"userData" yaml:"userData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.153.1/docs/resources/ocean_ecs#utilize_commitments OceanEcs#utilize_commitments}.
	UtilizeCommitments interface{} `field:"optional" json:"utilizeCommitments" yaml:"utilizeCommitments"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.153.1/docs/resources/ocean_ecs#utilize_reserved_instances OceanEcs#utilize_reserved_instances}.
	UtilizeReservedInstances interface{} `field:"optional" json:"utilizeReservedInstances" yaml:"utilizeReservedInstances"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.153.1/docs/resources/ocean_ecs#whitelist OceanEcs#whitelist}.
	Whitelist *[]*string `field:"optional" json:"whitelist" yaml:"whitelist"`
}

