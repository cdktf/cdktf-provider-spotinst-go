// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceancdstrategy


type OceancdStrategyCanaryStepsSetHeaderRouteMatchHeaderValue struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.173.0/docs/resources/oceancd_strategy#exact OceancdStrategy#exact}.
	Exact *string `field:"optional" json:"exact" yaml:"exact"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.173.0/docs/resources/oceancd_strategy#prefix OceancdStrategy#prefix}.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.173.0/docs/resources/oceancd_strategy#regex OceancdStrategy#regex}.
	Regex *string `field:"optional" json:"regex" yaml:"regex"`
}

