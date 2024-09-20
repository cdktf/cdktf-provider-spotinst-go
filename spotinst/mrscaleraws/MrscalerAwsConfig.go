// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mrscaleraws

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MrscalerAwsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#name MrscalerAws#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#strategy MrscalerAws#strategy}.
	Strategy *string `field:"required" json:"strategy" yaml:"strategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#additional_info MrscalerAws#additional_info}.
	AdditionalInfo *string `field:"optional" json:"additionalInfo" yaml:"additionalInfo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#additional_primary_security_groups MrscalerAws#additional_primary_security_groups}.
	AdditionalPrimarySecurityGroups *[]*string `field:"optional" json:"additionalPrimarySecurityGroups" yaml:"additionalPrimarySecurityGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#additional_replica_security_groups MrscalerAws#additional_replica_security_groups}.
	AdditionalReplicaSecurityGroups *[]*string `field:"optional" json:"additionalReplicaSecurityGroups" yaml:"additionalReplicaSecurityGroups"`
	// applications block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#applications MrscalerAws#applications}
	Applications interface{} `field:"optional" json:"applications" yaml:"applications"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#availability_zones MrscalerAws#availability_zones}.
	AvailabilityZones *[]*string `field:"optional" json:"availabilityZones" yaml:"availabilityZones"`
	// bootstrap_actions_file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#bootstrap_actions_file MrscalerAws#bootstrap_actions_file}
	BootstrapActionsFile interface{} `field:"optional" json:"bootstrapActionsFile" yaml:"bootstrapActionsFile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#cluster_id MrscalerAws#cluster_id}.
	ClusterId *string `field:"optional" json:"clusterId" yaml:"clusterId"`
	// configurations_file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#configurations_file MrscalerAws#configurations_file}
	ConfigurationsFile interface{} `field:"optional" json:"configurationsFile" yaml:"configurationsFile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#core_desired_capacity MrscalerAws#core_desired_capacity}.
	CoreDesiredCapacity *float64 `field:"optional" json:"coreDesiredCapacity" yaml:"coreDesiredCapacity"`
	// core_ebs_block_device block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#core_ebs_block_device MrscalerAws#core_ebs_block_device}
	CoreEbsBlockDevice interface{} `field:"optional" json:"coreEbsBlockDevice" yaml:"coreEbsBlockDevice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#core_ebs_optimized MrscalerAws#core_ebs_optimized}.
	CoreEbsOptimized interface{} `field:"optional" json:"coreEbsOptimized" yaml:"coreEbsOptimized"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#core_instance_types MrscalerAws#core_instance_types}.
	CoreInstanceTypes *[]*string `field:"optional" json:"coreInstanceTypes" yaml:"coreInstanceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#core_lifecycle MrscalerAws#core_lifecycle}.
	CoreLifecycle *string `field:"optional" json:"coreLifecycle" yaml:"coreLifecycle"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#core_max_size MrscalerAws#core_max_size}.
	CoreMaxSize *float64 `field:"optional" json:"coreMaxSize" yaml:"coreMaxSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#core_min_size MrscalerAws#core_min_size}.
	CoreMinSize *float64 `field:"optional" json:"coreMinSize" yaml:"coreMinSize"`
	// core_scaling_down_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#core_scaling_down_policy MrscalerAws#core_scaling_down_policy}
	CoreScalingDownPolicy interface{} `field:"optional" json:"coreScalingDownPolicy" yaml:"coreScalingDownPolicy"`
	// core_scaling_up_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#core_scaling_up_policy MrscalerAws#core_scaling_up_policy}
	CoreScalingUpPolicy interface{} `field:"optional" json:"coreScalingUpPolicy" yaml:"coreScalingUpPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#core_unit MrscalerAws#core_unit}.
	CoreUnit *string `field:"optional" json:"coreUnit" yaml:"coreUnit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#custom_ami_id MrscalerAws#custom_ami_id}.
	CustomAmiId *string `field:"optional" json:"customAmiId" yaml:"customAmiId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#description MrscalerAws#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#ebs_root_volume_size MrscalerAws#ebs_root_volume_size}.
	EbsRootVolumeSize *float64 `field:"optional" json:"ebsRootVolumeSize" yaml:"ebsRootVolumeSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#ec2_key_name MrscalerAws#ec2_key_name}.
	Ec2KeyName *string `field:"optional" json:"ec2KeyName" yaml:"ec2KeyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#expose_cluster_id MrscalerAws#expose_cluster_id}.
	ExposeClusterId interface{} `field:"optional" json:"exposeClusterId" yaml:"exposeClusterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#id MrscalerAws#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// instance_weights block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#instance_weights MrscalerAws#instance_weights}
	InstanceWeights interface{} `field:"optional" json:"instanceWeights" yaml:"instanceWeights"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#job_flow_role MrscalerAws#job_flow_role}.
	JobFlowRole *string `field:"optional" json:"jobFlowRole" yaml:"jobFlowRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#keep_job_flow_alive MrscalerAws#keep_job_flow_alive}.
	KeepJobFlowAlive interface{} `field:"optional" json:"keepJobFlowAlive" yaml:"keepJobFlowAlive"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#log_uri MrscalerAws#log_uri}.
	LogUri *string `field:"optional" json:"logUri" yaml:"logUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#managed_primary_security_group MrscalerAws#managed_primary_security_group}.
	ManagedPrimarySecurityGroup *string `field:"optional" json:"managedPrimarySecurityGroup" yaml:"managedPrimarySecurityGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#managed_replica_security_group MrscalerAws#managed_replica_security_group}.
	ManagedReplicaSecurityGroup *string `field:"optional" json:"managedReplicaSecurityGroup" yaml:"managedReplicaSecurityGroup"`
	// master_ebs_block_device block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#master_ebs_block_device MrscalerAws#master_ebs_block_device}
	MasterEbsBlockDevice interface{} `field:"optional" json:"masterEbsBlockDevice" yaml:"masterEbsBlockDevice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#master_ebs_optimized MrscalerAws#master_ebs_optimized}.
	MasterEbsOptimized interface{} `field:"optional" json:"masterEbsOptimized" yaml:"masterEbsOptimized"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#master_instance_types MrscalerAws#master_instance_types}.
	MasterInstanceTypes *[]*string `field:"optional" json:"masterInstanceTypes" yaml:"masterInstanceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#master_lifecycle MrscalerAws#master_lifecycle}.
	MasterLifecycle *string `field:"optional" json:"masterLifecycle" yaml:"masterLifecycle"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#master_target MrscalerAws#master_target}.
	MasterTarget *float64 `field:"optional" json:"masterTarget" yaml:"masterTarget"`
	// provisioning_timeout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#provisioning_timeout MrscalerAws#provisioning_timeout}
	ProvisioningTimeout *MrscalerAwsProvisioningTimeout `field:"optional" json:"provisioningTimeout" yaml:"provisioningTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#region MrscalerAws#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#release_label MrscalerAws#release_label}.
	ReleaseLabel *string `field:"optional" json:"releaseLabel" yaml:"releaseLabel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#repo_upgrade_on_boot MrscalerAws#repo_upgrade_on_boot}.
	RepoUpgradeOnBoot *string `field:"optional" json:"repoUpgradeOnBoot" yaml:"repoUpgradeOnBoot"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#retries MrscalerAws#retries}.
	Retries *float64 `field:"optional" json:"retries" yaml:"retries"`
	// scheduled_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#scheduled_task MrscalerAws#scheduled_task}
	ScheduledTask interface{} `field:"optional" json:"scheduledTask" yaml:"scheduledTask"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#security_config MrscalerAws#security_config}.
	SecurityConfig *string `field:"optional" json:"securityConfig" yaml:"securityConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#service_access_security_group MrscalerAws#service_access_security_group}.
	ServiceAccessSecurityGroup *string `field:"optional" json:"serviceAccessSecurityGroup" yaml:"serviceAccessSecurityGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#service_role MrscalerAws#service_role}.
	ServiceRole *string `field:"optional" json:"serviceRole" yaml:"serviceRole"`
	// steps_file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#steps_file MrscalerAws#steps_file}
	StepsFile interface{} `field:"optional" json:"stepsFile" yaml:"stepsFile"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#tags MrscalerAws#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#task_desired_capacity MrscalerAws#task_desired_capacity}.
	TaskDesiredCapacity *float64 `field:"optional" json:"taskDesiredCapacity" yaml:"taskDesiredCapacity"`
	// task_ebs_block_device block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#task_ebs_block_device MrscalerAws#task_ebs_block_device}
	TaskEbsBlockDevice interface{} `field:"optional" json:"taskEbsBlockDevice" yaml:"taskEbsBlockDevice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#task_ebs_optimized MrscalerAws#task_ebs_optimized}.
	TaskEbsOptimized interface{} `field:"optional" json:"taskEbsOptimized" yaml:"taskEbsOptimized"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#task_instance_types MrscalerAws#task_instance_types}.
	TaskInstanceTypes *[]*string `field:"optional" json:"taskInstanceTypes" yaml:"taskInstanceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#task_lifecycle MrscalerAws#task_lifecycle}.
	TaskLifecycle *string `field:"optional" json:"taskLifecycle" yaml:"taskLifecycle"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#task_max_size MrscalerAws#task_max_size}.
	TaskMaxSize *float64 `field:"optional" json:"taskMaxSize" yaml:"taskMaxSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#task_min_size MrscalerAws#task_min_size}.
	TaskMinSize *float64 `field:"optional" json:"taskMinSize" yaml:"taskMinSize"`
	// task_scaling_down_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#task_scaling_down_policy MrscalerAws#task_scaling_down_policy}
	TaskScalingDownPolicy interface{} `field:"optional" json:"taskScalingDownPolicy" yaml:"taskScalingDownPolicy"`
	// task_scaling_up_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#task_scaling_up_policy MrscalerAws#task_scaling_up_policy}
	TaskScalingUpPolicy interface{} `field:"optional" json:"taskScalingUpPolicy" yaml:"taskScalingUpPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#task_unit MrscalerAws#task_unit}.
	TaskUnit *string `field:"optional" json:"taskUnit" yaml:"taskUnit"`
	// termination_policies block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#termination_policies MrscalerAws#termination_policies}
	TerminationPolicies interface{} `field:"optional" json:"terminationPolicies" yaml:"terminationPolicies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#termination_protected MrscalerAws#termination_protected}.
	TerminationProtected interface{} `field:"optional" json:"terminationProtected" yaml:"terminationProtected"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.193.0/docs/resources/mrscaler_aws#visible_to_all_users MrscalerAws#visible_to_all_users}.
	VisibleToAllUsers interface{} `field:"optional" json:"visibleToAllUsers" yaml:"visibleToAllUsers"`
}

