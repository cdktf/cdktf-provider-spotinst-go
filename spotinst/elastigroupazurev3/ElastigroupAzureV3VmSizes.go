// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupazurev3


type ElastigroupAzureV3VmSizes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.221.0/docs/resources/elastigroup_azure_v3#od_sizes ElastigroupAzureV3#od_sizes}.
	OdSizes *[]*string `field:"required" json:"odSizes" yaml:"odSizes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.221.0/docs/resources/elastigroup_azure_v3#excluded_vm_sizes ElastigroupAzureV3#excluded_vm_sizes}.
	ExcludedVmSizes *[]*string `field:"optional" json:"excludedVmSizes" yaml:"excludedVmSizes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.221.0/docs/resources/elastigroup_azure_v3#preferred_spot_sizes ElastigroupAzureV3#preferred_spot_sizes}.
	PreferredSpotSizes *[]*string `field:"optional" json:"preferredSpotSizes" yaml:"preferredSpotSizes"`
	// spot_size_attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.221.0/docs/resources/elastigroup_azure_v3#spot_size_attributes ElastigroupAzureV3#spot_size_attributes}
	SpotSizeAttributes *ElastigroupAzureV3VmSizesSpotSizeAttributes `field:"optional" json:"spotSizeAttributes" yaml:"spotSizeAttributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.221.0/docs/resources/elastigroup_azure_v3#spot_sizes ElastigroupAzureV3#spot_sizes}.
	SpotSizes *[]*string `field:"optional" json:"spotSizes" yaml:"spotSizes"`
}

