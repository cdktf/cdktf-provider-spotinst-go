// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupazurev3


type ElastigroupAzureV3VmSizes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.215.0/docs/resources/elastigroup_azure_v3#od_sizes ElastigroupAzureV3#od_sizes}.
	OdSizes *[]*string `field:"required" json:"odSizes" yaml:"odSizes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.215.0/docs/resources/elastigroup_azure_v3#spot_sizes ElastigroupAzureV3#spot_sizes}.
	SpotSizes *[]*string `field:"required" json:"spotSizes" yaml:"spotSizes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.215.0/docs/resources/elastigroup_azure_v3#preferred_spot_sizes ElastigroupAzureV3#preferred_spot_sizes}.
	PreferredSpotSizes *[]*string `field:"optional" json:"preferredSpotSizes" yaml:"preferredSpotSizes"`
}

