// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceancdrolloutspec


type OceancdRolloutSpecStrategyArgs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.2/docs/resources/oceancd_rollout_spec#arg_name OceancdRolloutSpec#arg_name}.
	ArgName *string `field:"required" json:"argName" yaml:"argName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.2/docs/resources/oceancd_rollout_spec#arg_value OceancdRolloutSpec#arg_value}.
	ArgValue *string `field:"optional" json:"argValue" yaml:"argValue"`
	// value_from block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.2/docs/resources/oceancd_rollout_spec#value_from OceancdRolloutSpec#value_from}
	ValueFrom *OceancdRolloutSpecStrategyArgsValueFrom `field:"optional" json:"valueFrom" yaml:"valueFrom"`
}

