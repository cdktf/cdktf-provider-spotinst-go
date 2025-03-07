// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceancdstrategy


type OceancdStrategyRollingSteps struct {
	// pause block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.211.0/docs/resources/oceancd_strategy#pause OceancdStrategy#pause}
	Pause *OceancdStrategyRollingStepsPause `field:"optional" json:"pause" yaml:"pause"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.211.0/docs/resources/oceancd_strategy#steps_name OceancdStrategy#steps_name}.
	StepsName *string `field:"optional" json:"stepsName" yaml:"stepsName"`
	// verification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.211.0/docs/resources/oceancd_strategy#verification OceancdStrategy#verification}
	Verification *OceancdStrategyRollingStepsVerification `field:"optional" json:"verification" yaml:"verification"`
}

