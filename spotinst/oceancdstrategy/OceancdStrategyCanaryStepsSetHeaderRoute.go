// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceancdstrategy


type OceancdStrategyCanaryStepsSetHeaderRoute struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.1/docs/resources/oceancd_strategy#header_route_name OceancdStrategy#header_route_name}.
	HeaderRouteName *string `field:"required" json:"headerRouteName" yaml:"headerRouteName"`
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.1/docs/resources/oceancd_strategy#match OceancdStrategy#match}
	Match interface{} `field:"required" json:"match" yaml:"match"`
}

