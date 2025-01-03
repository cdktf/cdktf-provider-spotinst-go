// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceancdstrategy


type OceancdStrategyCanaryStepsSetHeaderRouteMatch struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.204.0/docs/resources/oceancd_strategy#header_name OceancdStrategy#header_name}.
	HeaderName *string `field:"required" json:"headerName" yaml:"headerName"`
	// header_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.204.0/docs/resources/oceancd_strategy#header_value OceancdStrategy#header_value}
	HeaderValue *OceancdStrategyCanaryStepsSetHeaderRouteMatchHeaderValue `field:"required" json:"headerValue" yaml:"headerValue"`
}

