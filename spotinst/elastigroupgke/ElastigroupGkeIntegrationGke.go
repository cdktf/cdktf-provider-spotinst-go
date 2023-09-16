// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupgke


type ElastigroupGkeIntegrationGke struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.139.0/docs/resources/elastigroup_gke#autoscale_cooldown ElastigroupGke#autoscale_cooldown}.
	AutoscaleCooldown *float64 `field:"optional" json:"autoscaleCooldown" yaml:"autoscaleCooldown"`
	// autoscale_down block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.139.0/docs/resources/elastigroup_gke#autoscale_down ElastigroupGke#autoscale_down}
	AutoscaleDown *ElastigroupGkeIntegrationGkeAutoscaleDown `field:"optional" json:"autoscaleDown" yaml:"autoscaleDown"`
	// autoscale_headroom block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.139.0/docs/resources/elastigroup_gke#autoscale_headroom ElastigroupGke#autoscale_headroom}
	AutoscaleHeadroom *ElastigroupGkeIntegrationGkeAutoscaleHeadroom `field:"optional" json:"autoscaleHeadroom" yaml:"autoscaleHeadroom"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.139.0/docs/resources/elastigroup_gke#autoscale_is_auto_config ElastigroupGke#autoscale_is_auto_config}.
	AutoscaleIsAutoConfig interface{} `field:"optional" json:"autoscaleIsAutoConfig" yaml:"autoscaleIsAutoConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.139.0/docs/resources/elastigroup_gke#autoscale_is_enabled ElastigroupGke#autoscale_is_enabled}.
	AutoscaleIsEnabled interface{} `field:"optional" json:"autoscaleIsEnabled" yaml:"autoscaleIsEnabled"`
	// autoscale_labels block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.139.0/docs/resources/elastigroup_gke#autoscale_labels ElastigroupGke#autoscale_labels}
	AutoscaleLabels interface{} `field:"optional" json:"autoscaleLabels" yaml:"autoscaleLabels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.139.0/docs/resources/elastigroup_gke#auto_update ElastigroupGke#auto_update}.
	AutoUpdate interface{} `field:"optional" json:"autoUpdate" yaml:"autoUpdate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.139.0/docs/resources/elastigroup_gke#cluster_id ElastigroupGke#cluster_id}.
	ClusterId *string `field:"optional" json:"clusterId" yaml:"clusterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.139.0/docs/resources/elastigroup_gke#location ElastigroupGke#location}.
	Location *string `field:"optional" json:"location" yaml:"location"`
}

