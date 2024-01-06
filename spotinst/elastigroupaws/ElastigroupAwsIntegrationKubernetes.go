// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupaws


type ElastigroupAwsIntegrationKubernetes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.158.0/docs/resources/elastigroup_aws#api_server ElastigroupAws#api_server}.
	ApiServer *string `field:"optional" json:"apiServer" yaml:"apiServer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.158.0/docs/resources/elastigroup_aws#autoscale_cooldown ElastigroupAws#autoscale_cooldown}.
	AutoscaleCooldown *float64 `field:"optional" json:"autoscaleCooldown" yaml:"autoscaleCooldown"`
	// autoscale_down block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.158.0/docs/resources/elastigroup_aws#autoscale_down ElastigroupAws#autoscale_down}
	AutoscaleDown *ElastigroupAwsIntegrationKubernetesAutoscaleDown `field:"optional" json:"autoscaleDown" yaml:"autoscaleDown"`
	// autoscale_headroom block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.158.0/docs/resources/elastigroup_aws#autoscale_headroom ElastigroupAws#autoscale_headroom}
	AutoscaleHeadroom *ElastigroupAwsIntegrationKubernetesAutoscaleHeadroom `field:"optional" json:"autoscaleHeadroom" yaml:"autoscaleHeadroom"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.158.0/docs/resources/elastigroup_aws#autoscale_is_auto_config ElastigroupAws#autoscale_is_auto_config}.
	AutoscaleIsAutoConfig interface{} `field:"optional" json:"autoscaleIsAutoConfig" yaml:"autoscaleIsAutoConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.158.0/docs/resources/elastigroup_aws#autoscale_is_enabled ElastigroupAws#autoscale_is_enabled}.
	AutoscaleIsEnabled interface{} `field:"optional" json:"autoscaleIsEnabled" yaml:"autoscaleIsEnabled"`
	// autoscale_labels block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.158.0/docs/resources/elastigroup_aws#autoscale_labels ElastigroupAws#autoscale_labels}
	AutoscaleLabels interface{} `field:"optional" json:"autoscaleLabels" yaml:"autoscaleLabels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.158.0/docs/resources/elastigroup_aws#cluster_identifier ElastigroupAws#cluster_identifier}.
	ClusterIdentifier *string `field:"optional" json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.158.0/docs/resources/elastigroup_aws#integration_mode ElastigroupAws#integration_mode}.
	IntegrationMode *string `field:"optional" json:"integrationMode" yaml:"integrationMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.158.0/docs/resources/elastigroup_aws#token ElastigroupAws#token}.
	Token *string `field:"optional" json:"token" yaml:"token"`
}

