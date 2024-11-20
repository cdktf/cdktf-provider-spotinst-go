// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceancdstrategy


type OceancdStrategyCanarySteps struct {
	// pause block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.198.0/docs/resources/oceancd_strategy#pause OceancdStrategy#pause}
	Pause *OceancdStrategyCanaryStepsPause `field:"optional" json:"pause" yaml:"pause"`
	// set_canary_scale block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.198.0/docs/resources/oceancd_strategy#set_canary_scale OceancdStrategy#set_canary_scale}
	SetCanaryScale *OceancdStrategyCanaryStepsSetCanaryScale `field:"optional" json:"setCanaryScale" yaml:"setCanaryScale"`
	// set_header_route block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.198.0/docs/resources/oceancd_strategy#set_header_route OceancdStrategy#set_header_route}
	SetHeaderRoute *OceancdStrategyCanaryStepsSetHeaderRoute `field:"optional" json:"setHeaderRoute" yaml:"setHeaderRoute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.198.0/docs/resources/oceancd_strategy#set_weight OceancdStrategy#set_weight}.
	SetWeight *float64 `field:"optional" json:"setWeight" yaml:"setWeight"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.198.0/docs/resources/oceancd_strategy#step_name OceancdStrategy#step_name}.
	StepName *string `field:"optional" json:"stepName" yaml:"stepName"`
	// verification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.198.0/docs/resources/oceancd_strategy#verification OceancdStrategy#verification}
	Verification *OceancdStrategyCanaryStepsVerification `field:"optional" json:"verification" yaml:"verification"`
}

