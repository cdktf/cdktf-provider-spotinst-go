// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceancdstrategy


type OceancdStrategyRolling struct {
	// steps block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.189.0/docs/resources/oceancd_strategy#steps OceancdStrategy#steps}
	Steps interface{} `field:"required" json:"steps" yaml:"steps"`
}

