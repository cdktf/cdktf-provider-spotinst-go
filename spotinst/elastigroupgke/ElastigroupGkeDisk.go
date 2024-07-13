// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupgke


type ElastigroupGkeDisk struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.181.0/docs/resources/elastigroup_gke#auto_delete ElastigroupGke#auto_delete}.
	AutoDelete interface{} `field:"optional" json:"autoDelete" yaml:"autoDelete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.181.0/docs/resources/elastigroup_gke#boot ElastigroupGke#boot}.
	Boot interface{} `field:"optional" json:"boot" yaml:"boot"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.181.0/docs/resources/elastigroup_gke#device_name ElastigroupGke#device_name}.
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// initialize_params block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.181.0/docs/resources/elastigroup_gke#initialize_params ElastigroupGke#initialize_params}
	InitializeParams interface{} `field:"optional" json:"initializeParams" yaml:"initializeParams"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.181.0/docs/resources/elastigroup_gke#interface ElastigroupGke#interface}.
	Interface *string `field:"optional" json:"interface" yaml:"interface"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.181.0/docs/resources/elastigroup_gke#mode ElastigroupGke#mode}.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.181.0/docs/resources/elastigroup_gke#source ElastigroupGke#source}.
	Source *string `field:"optional" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.181.0/docs/resources/elastigroup_gke#type ElastigroupGke#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

