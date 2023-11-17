// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupgcp


type ElastigroupGcpDisk struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.150.0/docs/resources/elastigroup_gcp#auto_delete ElastigroupGcp#auto_delete}.
	AutoDelete interface{} `field:"optional" json:"autoDelete" yaml:"autoDelete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.150.0/docs/resources/elastigroup_gcp#boot ElastigroupGcp#boot}.
	Boot interface{} `field:"optional" json:"boot" yaml:"boot"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.150.0/docs/resources/elastigroup_gcp#device_name ElastigroupGcp#device_name}.
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// initialize_params block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.150.0/docs/resources/elastigroup_gcp#initialize_params ElastigroupGcp#initialize_params}
	InitializeParams interface{} `field:"optional" json:"initializeParams" yaml:"initializeParams"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.150.0/docs/resources/elastigroup_gcp#interface ElastigroupGcp#interface}.
	Interface *string `field:"optional" json:"interface" yaml:"interface"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.150.0/docs/resources/elastigroup_gcp#mode ElastigroupGcp#mode}.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.150.0/docs/resources/elastigroup_gcp#source ElastigroupGcp#source}.
	Source *string `field:"optional" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.150.0/docs/resources/elastigroup_gcp#type ElastigroupGcp#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

