// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupgke


type ElastigroupGkeNetworkInterfaceAliasIpRanges struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.184.0/docs/resources/elastigroup_gke#ip_cidr_range ElastigroupGke#ip_cidr_range}.
	IpCidrRange *string `field:"required" json:"ipCidrRange" yaml:"ipCidrRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.184.0/docs/resources/elastigroup_gke#subnetwork_range_name ElastigroupGke#subnetwork_range_name}.
	SubnetworkRangeName *string `field:"required" json:"subnetworkRangeName" yaml:"subnetworkRangeName"`
}

