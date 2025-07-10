// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanrightsizingrule


type OceanRightSizingRuleRecommendationApplicationIntervals struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/ocean_right_sizing_rule#repetition_basis OceanRightSizingRule#repetition_basis}.
	RepetitionBasis *string `field:"required" json:"repetitionBasis" yaml:"repetitionBasis"`
	// monthly_repetition_basis block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/ocean_right_sizing_rule#monthly_repetition_basis OceanRightSizingRule#monthly_repetition_basis}
	MonthlyRepetitionBasis interface{} `field:"optional" json:"monthlyRepetitionBasis" yaml:"monthlyRepetitionBasis"`
	// weekly_repetition_basis block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/ocean_right_sizing_rule#weekly_repetition_basis OceanRightSizingRule#weekly_repetition_basis}
	WeeklyRepetitionBasis interface{} `field:"optional" json:"weeklyRepetitionBasis" yaml:"weeklyRepetitionBasis"`
}

