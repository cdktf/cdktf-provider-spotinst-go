// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupgcp


type ElastigroupGcpSubnets struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.222.1/docs/resources/elastigroup_gcp#region ElastigroupGcp#region}.
	Region *string `field:"required" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.222.1/docs/resources/elastigroup_gcp#subnet_names ElastigroupGcp#subnet_names}.
	SubnetNames *[]*string `field:"required" json:"subnetNames" yaml:"subnetNames"`
}

