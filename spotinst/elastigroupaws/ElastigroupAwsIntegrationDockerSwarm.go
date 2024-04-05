// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupaws


type ElastigroupAwsIntegrationDockerSwarm struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.168.2/docs/resources/elastigroup_aws#master_host ElastigroupAws#master_host}.
	MasterHost *string `field:"required" json:"masterHost" yaml:"masterHost"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.168.2/docs/resources/elastigroup_aws#master_port ElastigroupAws#master_port}.
	MasterPort *float64 `field:"required" json:"masterPort" yaml:"masterPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.168.2/docs/resources/elastigroup_aws#autoscale_cooldown ElastigroupAws#autoscale_cooldown}.
	AutoscaleCooldown *float64 `field:"optional" json:"autoscaleCooldown" yaml:"autoscaleCooldown"`
	// autoscale_down block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.168.2/docs/resources/elastigroup_aws#autoscale_down ElastigroupAws#autoscale_down}
	AutoscaleDown *ElastigroupAwsIntegrationDockerSwarmAutoscaleDown `field:"optional" json:"autoscaleDown" yaml:"autoscaleDown"`
	// autoscale_headroom block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.168.2/docs/resources/elastigroup_aws#autoscale_headroom ElastigroupAws#autoscale_headroom}
	AutoscaleHeadroom *ElastigroupAwsIntegrationDockerSwarmAutoscaleHeadroom `field:"optional" json:"autoscaleHeadroom" yaml:"autoscaleHeadroom"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.168.2/docs/resources/elastigroup_aws#autoscale_is_enabled ElastigroupAws#autoscale_is_enabled}.
	AutoscaleIsEnabled interface{} `field:"optional" json:"autoscaleIsEnabled" yaml:"autoscaleIsEnabled"`
}

