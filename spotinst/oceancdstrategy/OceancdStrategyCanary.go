// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceancdstrategy


type OceancdStrategyCanary struct {
	// steps block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.198.0/docs/resources/oceancd_strategy#steps OceancdStrategy#steps}
	Steps interface{} `field:"required" json:"steps" yaml:"steps"`
	// background_verification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.198.0/docs/resources/oceancd_strategy#background_verification OceancdStrategy#background_verification}
	BackgroundVerification *OceancdStrategyCanaryBackgroundVerification `field:"optional" json:"backgroundVerification" yaml:"backgroundVerification"`
}

