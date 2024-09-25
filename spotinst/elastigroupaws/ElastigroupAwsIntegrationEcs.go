// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupaws


type ElastigroupAwsIntegrationEcs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.194.1/docs/resources/elastigroup_aws#cluster_name ElastigroupAws#cluster_name}.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// autoscale_attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.194.1/docs/resources/elastigroup_aws#autoscale_attributes ElastigroupAws#autoscale_attributes}
	AutoscaleAttributes interface{} `field:"optional" json:"autoscaleAttributes" yaml:"autoscaleAttributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.194.1/docs/resources/elastigroup_aws#autoscale_cooldown ElastigroupAws#autoscale_cooldown}.
	AutoscaleCooldown *float64 `field:"optional" json:"autoscaleCooldown" yaml:"autoscaleCooldown"`
	// autoscale_down block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.194.1/docs/resources/elastigroup_aws#autoscale_down ElastigroupAws#autoscale_down}
	AutoscaleDown *ElastigroupAwsIntegrationEcsAutoscaleDown `field:"optional" json:"autoscaleDown" yaml:"autoscaleDown"`
	// autoscale_headroom block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.194.1/docs/resources/elastigroup_aws#autoscale_headroom ElastigroupAws#autoscale_headroom}
	AutoscaleHeadroom *ElastigroupAwsIntegrationEcsAutoscaleHeadroom `field:"optional" json:"autoscaleHeadroom" yaml:"autoscaleHeadroom"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.194.1/docs/resources/elastigroup_aws#autoscale_is_auto_config ElastigroupAws#autoscale_is_auto_config}.
	AutoscaleIsAutoConfig interface{} `field:"optional" json:"autoscaleIsAutoConfig" yaml:"autoscaleIsAutoConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.194.1/docs/resources/elastigroup_aws#autoscale_is_enabled ElastigroupAws#autoscale_is_enabled}.
	AutoscaleIsEnabled interface{} `field:"optional" json:"autoscaleIsEnabled" yaml:"autoscaleIsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.194.1/docs/resources/elastigroup_aws#autoscale_scale_down_non_service_tasks ElastigroupAws#autoscale_scale_down_non_service_tasks}.
	AutoscaleScaleDownNonServiceTasks interface{} `field:"optional" json:"autoscaleScaleDownNonServiceTasks" yaml:"autoscaleScaleDownNonServiceTasks"`
	// batch block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.194.1/docs/resources/elastigroup_aws#batch ElastigroupAws#batch}
	Batch *ElastigroupAwsIntegrationEcsBatch `field:"optional" json:"batch" yaml:"batch"`
}

