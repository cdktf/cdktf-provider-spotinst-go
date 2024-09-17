// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package statefulnodeazure


type StatefulNodeAzureVmSizes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.191.0/docs/resources/stateful_node_azure#od_sizes StatefulNodeAzure#od_sizes}.
	OdSizes *[]*string `field:"required" json:"odSizes" yaml:"odSizes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.191.0/docs/resources/stateful_node_azure#spot_sizes StatefulNodeAzure#spot_sizes}.
	SpotSizes *[]*string `field:"required" json:"spotSizes" yaml:"spotSizes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.191.0/docs/resources/stateful_node_azure#preferred_spot_sizes StatefulNodeAzure#preferred_spot_sizes}.
	PreferredSpotSizes *[]*string `field:"optional" json:"preferredSpotSizes" yaml:"preferredSpotSizes"`
}

