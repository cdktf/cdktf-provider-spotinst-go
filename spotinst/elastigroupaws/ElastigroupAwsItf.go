// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupaws


type ElastigroupAwsItf struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.195.0/docs/resources/elastigroup_aws#fixed_target_groups ElastigroupAws#fixed_target_groups}.
	FixedTargetGroups interface{} `field:"required" json:"fixedTargetGroups" yaml:"fixedTargetGroups"`
	// load_balancer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.195.0/docs/resources/elastigroup_aws#load_balancer ElastigroupAws#load_balancer}
	LoadBalancer interface{} `field:"required" json:"loadBalancer" yaml:"loadBalancer"`
	// target_group_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.195.0/docs/resources/elastigroup_aws#target_group_config ElastigroupAws#target_group_config}
	TargetGroupConfig interface{} `field:"required" json:"targetGroupConfig" yaml:"targetGroupConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.195.0/docs/resources/elastigroup_aws#weight_strategy ElastigroupAws#weight_strategy}.
	WeightStrategy *string `field:"required" json:"weightStrategy" yaml:"weightStrategy"`
	// default_static_target_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.195.0/docs/resources/elastigroup_aws#default_static_target_group ElastigroupAws#default_static_target_group}
	DefaultStaticTargetGroup *ElastigroupAwsItfDefaultStaticTargetGroup `field:"optional" json:"defaultStaticTargetGroup" yaml:"defaultStaticTargetGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.195.0/docs/resources/elastigroup_aws#migration_healthiness_threshold ElastigroupAws#migration_healthiness_threshold}.
	MigrationHealthinessThreshold *float64 `field:"optional" json:"migrationHealthinessThreshold" yaml:"migrationHealthinessThreshold"`
}

