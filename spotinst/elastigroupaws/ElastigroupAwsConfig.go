// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupaws

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ElastigroupAwsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#fallback_to_ondemand ElastigroupAws#fallback_to_ondemand}.
	FallbackToOndemand interface{} `field:"required" json:"fallbackToOndemand" yaml:"fallbackToOndemand"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#name ElastigroupAws#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#orientation ElastigroupAws#orientation}.
	Orientation *string `field:"required" json:"orientation" yaml:"orientation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#product ElastigroupAws#product}.
	Product *string `field:"required" json:"product" yaml:"product"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#security_groups ElastigroupAws#security_groups}.
	SecurityGroups *[]*string `field:"required" json:"securityGroups" yaml:"securityGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#availability_zones ElastigroupAws#availability_zones}.
	AvailabilityZones *[]*string `field:"optional" json:"availabilityZones" yaml:"availabilityZones"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#block_devices_mode ElastigroupAws#block_devices_mode}.
	BlockDevicesMode *string `field:"optional" json:"blockDevicesMode" yaml:"blockDevicesMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#capacity_unit ElastigroupAws#capacity_unit}.
	CapacityUnit *string `field:"optional" json:"capacityUnit" yaml:"capacityUnit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#consider_od_pricing ElastigroupAws#consider_od_pricing}.
	ConsiderOdPricing interface{} `field:"optional" json:"considerOdPricing" yaml:"considerOdPricing"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#cpu_credits ElastigroupAws#cpu_credits}.
	CpuCredits *string `field:"optional" json:"cpuCredits" yaml:"cpuCredits"`
	// cpu_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#cpu_options ElastigroupAws#cpu_options}
	CpuOptions *ElastigroupAwsCpuOptions `field:"optional" json:"cpuOptions" yaml:"cpuOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#description ElastigroupAws#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#desired_capacity ElastigroupAws#desired_capacity}.
	DesiredCapacity *float64 `field:"optional" json:"desiredCapacity" yaml:"desiredCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#draining_timeout ElastigroupAws#draining_timeout}.
	DrainingTimeout *float64 `field:"optional" json:"drainingTimeout" yaml:"drainingTimeout"`
	// ebs_block_device block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#ebs_block_device ElastigroupAws#ebs_block_device}
	EbsBlockDevice interface{} `field:"optional" json:"ebsBlockDevice" yaml:"ebsBlockDevice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#ebs_optimized ElastigroupAws#ebs_optimized}.
	EbsOptimized interface{} `field:"optional" json:"ebsOptimized" yaml:"ebsOptimized"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#elastic_ips ElastigroupAws#elastic_ips}.
	ElasticIps *[]*string `field:"optional" json:"elasticIps" yaml:"elasticIps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#elastic_load_balancers ElastigroupAws#elastic_load_balancers}.
	ElasticLoadBalancers *[]*string `field:"optional" json:"elasticLoadBalancers" yaml:"elasticLoadBalancers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#enable_monitoring ElastigroupAws#enable_monitoring}.
	EnableMonitoring interface{} `field:"optional" json:"enableMonitoring" yaml:"enableMonitoring"`
	// ephemeral_block_device block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#ephemeral_block_device ElastigroupAws#ephemeral_block_device}
	EphemeralBlockDevice interface{} `field:"optional" json:"ephemeralBlockDevice" yaml:"ephemeralBlockDevice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#health_check_grace_period ElastigroupAws#health_check_grace_period}.
	HealthCheckGracePeriod *float64 `field:"optional" json:"healthCheckGracePeriod" yaml:"healthCheckGracePeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#health_check_type ElastigroupAws#health_check_type}.
	HealthCheckType *string `field:"optional" json:"healthCheckType" yaml:"healthCheckType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#health_check_unhealthy_duration_before_replacement ElastigroupAws#health_check_unhealthy_duration_before_replacement}.
	HealthCheckUnhealthyDurationBeforeReplacement *float64 `field:"optional" json:"healthCheckUnhealthyDurationBeforeReplacement" yaml:"healthCheckUnhealthyDurationBeforeReplacement"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#iam_instance_profile ElastigroupAws#iam_instance_profile}.
	IamInstanceProfile *string `field:"optional" json:"iamInstanceProfile" yaml:"iamInstanceProfile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#id ElastigroupAws#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#image_id ElastigroupAws#image_id}.
	ImageId *string `field:"optional" json:"imageId" yaml:"imageId"`
	// images block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#images ElastigroupAws#images}
	Images interface{} `field:"optional" json:"images" yaml:"images"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#immediate_od_recover_threshold ElastigroupAws#immediate_od_recover_threshold}.
	ImmediateOdRecoverThreshold *float64 `field:"optional" json:"immediateOdRecoverThreshold" yaml:"immediateOdRecoverThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#instance_types_ondemand ElastigroupAws#instance_types_ondemand}.
	InstanceTypesOndemand *string `field:"optional" json:"instanceTypesOndemand" yaml:"instanceTypesOndemand"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#instance_types_preferred_spot ElastigroupAws#instance_types_preferred_spot}.
	InstanceTypesPreferredSpot *[]*string `field:"optional" json:"instanceTypesPreferredSpot" yaml:"instanceTypesPreferredSpot"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#instance_types_spot ElastigroupAws#instance_types_spot}.
	InstanceTypesSpot *[]*string `field:"optional" json:"instanceTypesSpot" yaml:"instanceTypesSpot"`
	// instance_types_weights block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#instance_types_weights ElastigroupAws#instance_types_weights}
	InstanceTypesWeights interface{} `field:"optional" json:"instanceTypesWeights" yaml:"instanceTypesWeights"`
	// integration_beanstalk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#integration_beanstalk ElastigroupAws#integration_beanstalk}
	IntegrationBeanstalk *ElastigroupAwsIntegrationBeanstalk `field:"optional" json:"integrationBeanstalk" yaml:"integrationBeanstalk"`
	// integration_codedeploy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#integration_codedeploy ElastigroupAws#integration_codedeploy}
	IntegrationCodedeploy *ElastigroupAwsIntegrationCodedeploy `field:"optional" json:"integrationCodedeploy" yaml:"integrationCodedeploy"`
	// integration_docker_swarm block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#integration_docker_swarm ElastigroupAws#integration_docker_swarm}
	IntegrationDockerSwarm *ElastigroupAwsIntegrationDockerSwarm `field:"optional" json:"integrationDockerSwarm" yaml:"integrationDockerSwarm"`
	// integration_ecs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#integration_ecs ElastigroupAws#integration_ecs}
	IntegrationEcs *ElastigroupAwsIntegrationEcs `field:"optional" json:"integrationEcs" yaml:"integrationEcs"`
	// integration_gitlab block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#integration_gitlab ElastigroupAws#integration_gitlab}
	IntegrationGitlab *ElastigroupAwsIntegrationGitlab `field:"optional" json:"integrationGitlab" yaml:"integrationGitlab"`
	// integration_kubernetes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#integration_kubernetes ElastigroupAws#integration_kubernetes}
	IntegrationKubernetes *ElastigroupAwsIntegrationKubernetes `field:"optional" json:"integrationKubernetes" yaml:"integrationKubernetes"`
	// integration_mesosphere block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#integration_mesosphere ElastigroupAws#integration_mesosphere}
	IntegrationMesosphere *ElastigroupAwsIntegrationMesosphere `field:"optional" json:"integrationMesosphere" yaml:"integrationMesosphere"`
	// integration_nomad block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#integration_nomad ElastigroupAws#integration_nomad}
	IntegrationNomad *ElastigroupAwsIntegrationNomad `field:"optional" json:"integrationNomad" yaml:"integrationNomad"`
	// integration_rancher block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#integration_rancher ElastigroupAws#integration_rancher}
	IntegrationRancher *ElastigroupAwsIntegrationRancher `field:"optional" json:"integrationRancher" yaml:"integrationRancher"`
	// integration_route53 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#integration_route53 ElastigroupAws#integration_route53}
	IntegrationRoute53 *ElastigroupAwsIntegrationRoute53 `field:"optional" json:"integrationRoute53" yaml:"integrationRoute53"`
	// itf block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#itf ElastigroupAws#itf}
	Itf interface{} `field:"optional" json:"itf" yaml:"itf"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#key_name ElastigroupAws#key_name}.
	KeyName *string `field:"optional" json:"keyName" yaml:"keyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#lifetime_period ElastigroupAws#lifetime_period}.
	LifetimePeriod *string `field:"optional" json:"lifetimePeriod" yaml:"lifetimePeriod"`
	// logging block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#logging ElastigroupAws#logging}
	Logging *ElastigroupAwsLogging `field:"optional" json:"logging" yaml:"logging"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#max_size ElastigroupAws#max_size}.
	MaxSize *float64 `field:"optional" json:"maxSize" yaml:"maxSize"`
	// metadata_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#metadata_options ElastigroupAws#metadata_options}
	MetadataOptions *ElastigroupAwsMetadataOptions `field:"optional" json:"metadataOptions" yaml:"metadataOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#minimum_instance_lifetime ElastigroupAws#minimum_instance_lifetime}.
	MinimumInstanceLifetime *float64 `field:"optional" json:"minimumInstanceLifetime" yaml:"minimumInstanceLifetime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#min_size ElastigroupAws#min_size}.
	MinSize *float64 `field:"optional" json:"minSize" yaml:"minSize"`
	// multiple_metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#multiple_metrics ElastigroupAws#multiple_metrics}
	MultipleMetrics *ElastigroupAwsMultipleMetrics `field:"optional" json:"multipleMetrics" yaml:"multipleMetrics"`
	// network_interface block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#network_interface ElastigroupAws#network_interface}
	NetworkInterface interface{} `field:"optional" json:"networkInterface" yaml:"networkInterface"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#ondemand_count ElastigroupAws#ondemand_count}.
	OndemandCount *float64 `field:"optional" json:"ondemandCount" yaml:"ondemandCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#on_demand_types ElastigroupAws#on_demand_types}.
	OnDemandTypes *[]*string `field:"optional" json:"onDemandTypes" yaml:"onDemandTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#persist_block_devices ElastigroupAws#persist_block_devices}.
	PersistBlockDevices interface{} `field:"optional" json:"persistBlockDevices" yaml:"persistBlockDevices"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#persist_private_ip ElastigroupAws#persist_private_ip}.
	PersistPrivateIp interface{} `field:"optional" json:"persistPrivateIp" yaml:"persistPrivateIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#persist_root_device ElastigroupAws#persist_root_device}.
	PersistRootDevice interface{} `field:"optional" json:"persistRootDevice" yaml:"persistRootDevice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#placement_tenancy ElastigroupAws#placement_tenancy}.
	PlacementTenancy *string `field:"optional" json:"placementTenancy" yaml:"placementTenancy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#preferred_availability_zones ElastigroupAws#preferred_availability_zones}.
	PreferredAvailabilityZones *[]*string `field:"optional" json:"preferredAvailabilityZones" yaml:"preferredAvailabilityZones"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#private_ips ElastigroupAws#private_ips}.
	PrivateIps *[]*string `field:"optional" json:"privateIps" yaml:"privateIps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#region ElastigroupAws#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// resource_requirements block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#resource_requirements ElastigroupAws#resource_requirements}
	ResourceRequirements interface{} `field:"optional" json:"resourceRequirements" yaml:"resourceRequirements"`
	// resource_tag_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#resource_tag_specification ElastigroupAws#resource_tag_specification}
	ResourceTagSpecification interface{} `field:"optional" json:"resourceTagSpecification" yaml:"resourceTagSpecification"`
	// revert_to_spot block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#revert_to_spot ElastigroupAws#revert_to_spot}
	RevertToSpot *ElastigroupAwsRevertToSpot `field:"optional" json:"revertToSpot" yaml:"revertToSpot"`
	// scaling_down_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#scaling_down_policy ElastigroupAws#scaling_down_policy}
	ScalingDownPolicy interface{} `field:"optional" json:"scalingDownPolicy" yaml:"scalingDownPolicy"`
	// scaling_strategy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#scaling_strategy ElastigroupAws#scaling_strategy}
	ScalingStrategy interface{} `field:"optional" json:"scalingStrategy" yaml:"scalingStrategy"`
	// scaling_target_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#scaling_target_policy ElastigroupAws#scaling_target_policy}
	ScalingTargetPolicy interface{} `field:"optional" json:"scalingTargetPolicy" yaml:"scalingTargetPolicy"`
	// scaling_up_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#scaling_up_policy ElastigroupAws#scaling_up_policy}
	ScalingUpPolicy interface{} `field:"optional" json:"scalingUpPolicy" yaml:"scalingUpPolicy"`
	// scheduled_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#scheduled_task ElastigroupAws#scheduled_task}
	ScheduledTask interface{} `field:"optional" json:"scheduledTask" yaml:"scheduledTask"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#shutdown_script ElastigroupAws#shutdown_script}.
	ShutdownScript *string `field:"optional" json:"shutdownScript" yaml:"shutdownScript"`
	// signal block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#signal ElastigroupAws#signal}
	Signal interface{} `field:"optional" json:"signal" yaml:"signal"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#spot_percentage ElastigroupAws#spot_percentage}.
	SpotPercentage *float64 `field:"optional" json:"spotPercentage" yaml:"spotPercentage"`
	// stateful_deallocation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#stateful_deallocation ElastigroupAws#stateful_deallocation}
	StatefulDeallocation *ElastigroupAwsStatefulDeallocation `field:"optional" json:"statefulDeallocation" yaml:"statefulDeallocation"`
	// stateful_instance_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#stateful_instance_action ElastigroupAws#stateful_instance_action}
	StatefulInstanceAction interface{} `field:"optional" json:"statefulInstanceAction" yaml:"statefulInstanceAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#subnet_ids ElastigroupAws#subnet_ids}.
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#tags ElastigroupAws#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#target_group_arns ElastigroupAws#target_group_arns}.
	TargetGroupArns *[]*string `field:"optional" json:"targetGroupArns" yaml:"targetGroupArns"`
	// update_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#update_policy ElastigroupAws#update_policy}
	UpdatePolicy *ElastigroupAwsUpdatePolicy `field:"optional" json:"updatePolicy" yaml:"updatePolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#user_data ElastigroupAws#user_data}.
	UserData *string `field:"optional" json:"userData" yaml:"userData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#utilize_commitments ElastigroupAws#utilize_commitments}.
	UtilizeCommitments interface{} `field:"optional" json:"utilizeCommitments" yaml:"utilizeCommitments"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#utilize_reserved_instances ElastigroupAws#utilize_reserved_instances}.
	UtilizeReservedInstances interface{} `field:"optional" json:"utilizeReservedInstances" yaml:"utilizeReservedInstances"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#wait_for_capacity ElastigroupAws#wait_for_capacity}.
	WaitForCapacity *float64 `field:"optional" json:"waitForCapacity" yaml:"waitForCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/elastigroup_aws#wait_for_capacity_timeout ElastigroupAws#wait_for_capacity_timeout}.
	WaitForCapacityTimeout *float64 `field:"optional" json:"waitForCapacityTimeout" yaml:"waitForCapacityTimeout"`
}

