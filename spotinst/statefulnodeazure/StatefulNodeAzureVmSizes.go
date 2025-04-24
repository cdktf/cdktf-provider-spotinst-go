// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package statefulnodeazure


type StatefulNodeAzureVmSizes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.216.2/docs/resources/stateful_node_azure#od_sizes StatefulNodeAzure#od_sizes}.
	OdSizes *[]*string `field:"required" json:"odSizes" yaml:"odSizes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.216.2/docs/resources/stateful_node_azure#excluded_vm_sizes StatefulNodeAzure#excluded_vm_sizes}.
	ExcludedVmSizes *[]*string `field:"optional" json:"excludedVmSizes" yaml:"excludedVmSizes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.216.2/docs/resources/stateful_node_azure#preferred_spot_sizes StatefulNodeAzure#preferred_spot_sizes}.
	PreferredSpotSizes *[]*string `field:"optional" json:"preferredSpotSizes" yaml:"preferredSpotSizes"`
	// spot_size_attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.216.2/docs/resources/stateful_node_azure#spot_size_attributes StatefulNodeAzure#spot_size_attributes}
	SpotSizeAttributes *StatefulNodeAzureVmSizesSpotSizeAttributes `field:"optional" json:"spotSizeAttributes" yaml:"spotSizeAttributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.216.2/docs/resources/stateful_node_azure#spot_sizes StatefulNodeAzure#spot_sizes}.
	SpotSizes *[]*string `field:"optional" json:"spotSizes" yaml:"spotSizes"`
}

