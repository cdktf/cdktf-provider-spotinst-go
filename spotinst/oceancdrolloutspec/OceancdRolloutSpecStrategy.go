// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceancdrolloutspec


type OceancdRolloutSpecStrategy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.198.0/docs/resources/oceancd_rollout_spec#strategy_name OceancdRolloutSpec#strategy_name}.
	StrategyName *string `field:"required" json:"strategyName" yaml:"strategyName"`
	// args block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.198.0/docs/resources/oceancd_rollout_spec#args OceancdRolloutSpec#args}
	Args interface{} `field:"optional" json:"args" yaml:"args"`
}

