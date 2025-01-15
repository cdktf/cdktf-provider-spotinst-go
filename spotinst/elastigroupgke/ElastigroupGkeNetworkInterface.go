// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupgke


type ElastigroupGkeNetworkInterface struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.207.0/docs/resources/elastigroup_gke#network ElastigroupGke#network}.
	Network *string `field:"required" json:"network" yaml:"network"`
	// access_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.207.0/docs/resources/elastigroup_gke#access_configs ElastigroupGke#access_configs}
	AccessConfigs interface{} `field:"optional" json:"accessConfigs" yaml:"accessConfigs"`
	// alias_ip_ranges block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.207.0/docs/resources/elastigroup_gke#alias_ip_ranges ElastigroupGke#alias_ip_ranges}
	AliasIpRanges interface{} `field:"optional" json:"aliasIpRanges" yaml:"aliasIpRanges"`
}

