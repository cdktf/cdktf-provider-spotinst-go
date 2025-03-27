// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceancdrolloutspec


type OceancdRolloutSpecTrafficIstioDestinationRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.216.0/docs/resources/oceancd_rollout_spec#canary_subset_name OceancdRolloutSpec#canary_subset_name}.
	CanarySubsetName *string `field:"required" json:"canarySubsetName" yaml:"canarySubsetName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.216.0/docs/resources/oceancd_rollout_spec#destination_rule_name OceancdRolloutSpec#destination_rule_name}.
	DestinationRuleName *string `field:"required" json:"destinationRuleName" yaml:"destinationRuleName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.216.0/docs/resources/oceancd_rollout_spec#stable_subset_name OceancdRolloutSpec#stable_subset_name}.
	StableSubsetName *string `field:"required" json:"stableSubsetName" yaml:"stableSubsetName"`
}

