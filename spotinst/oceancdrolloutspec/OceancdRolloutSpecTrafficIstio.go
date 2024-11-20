// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceancdrolloutspec


type OceancdRolloutSpecTrafficIstio struct {
	// virtual_services block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.198.0/docs/resources/oceancd_rollout_spec#virtual_services OceancdRolloutSpec#virtual_services}
	VirtualServices interface{} `field:"required" json:"virtualServices" yaml:"virtualServices"`
	// destination_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.198.0/docs/resources/oceancd_rollout_spec#destination_rule OceancdRolloutSpec#destination_rule}
	DestinationRule *OceancdRolloutSpecTrafficIstioDestinationRule `field:"optional" json:"destinationRule" yaml:"destinationRule"`
}

