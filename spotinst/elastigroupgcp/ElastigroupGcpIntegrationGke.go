// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupgcp


type ElastigroupGcpIntegrationGke struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.216.1/docs/resources/elastigroup_gcp#autoscale_cooldown ElastigroupGcp#autoscale_cooldown}.
	AutoscaleCooldown *float64 `field:"optional" json:"autoscaleCooldown" yaml:"autoscaleCooldown"`
	// autoscale_down block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.216.1/docs/resources/elastigroup_gcp#autoscale_down ElastigroupGcp#autoscale_down}
	AutoscaleDown *ElastigroupGcpIntegrationGkeAutoscaleDown `field:"optional" json:"autoscaleDown" yaml:"autoscaleDown"`
	// autoscale_headroom block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.216.1/docs/resources/elastigroup_gcp#autoscale_headroom ElastigroupGcp#autoscale_headroom}
	AutoscaleHeadroom *ElastigroupGcpIntegrationGkeAutoscaleHeadroom `field:"optional" json:"autoscaleHeadroom" yaml:"autoscaleHeadroom"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.216.1/docs/resources/elastigroup_gcp#autoscale_is_auto_config ElastigroupGcp#autoscale_is_auto_config}.
	AutoscaleIsAutoConfig interface{} `field:"optional" json:"autoscaleIsAutoConfig" yaml:"autoscaleIsAutoConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.216.1/docs/resources/elastigroup_gcp#autoscale_is_enabled ElastigroupGcp#autoscale_is_enabled}.
	AutoscaleIsEnabled interface{} `field:"optional" json:"autoscaleIsEnabled" yaml:"autoscaleIsEnabled"`
	// autoscale_labels block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.216.1/docs/resources/elastigroup_gcp#autoscale_labels ElastigroupGcp#autoscale_labels}
	AutoscaleLabels interface{} `field:"optional" json:"autoscaleLabels" yaml:"autoscaleLabels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.216.1/docs/resources/elastigroup_gcp#auto_update ElastigroupGcp#auto_update}.
	AutoUpdate interface{} `field:"optional" json:"autoUpdate" yaml:"autoUpdate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.216.1/docs/resources/elastigroup_gcp#cluster_id ElastigroupGcp#cluster_id}.
	ClusterId *string `field:"optional" json:"clusterId" yaml:"clusterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.216.1/docs/resources/elastigroup_gcp#location ElastigroupGcp#location}.
	Location *string `field:"optional" json:"location" yaml:"location"`
}

