// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanrightsizingrule


type OceanRightSizingRuleRecommendationApplicationIntervalsMonthlyRepetitionBasis struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.230.0/docs/resources/ocean_right_sizing_rule#interval_months OceanRightSizingRule#interval_months}.
	IntervalMonths *[]*float64 `field:"required" json:"intervalMonths" yaml:"intervalMonths"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.230.0/docs/resources/ocean_right_sizing_rule#week_of_the_month OceanRightSizingRule#week_of_the_month}.
	WeekOfTheMonth *[]*string `field:"required" json:"weekOfTheMonth" yaml:"weekOfTheMonth"`
	// weekly_repetition_basis block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.230.0/docs/resources/ocean_right_sizing_rule#weekly_repetition_basis OceanRightSizingRule#weekly_repetition_basis}
	WeeklyRepetitionBasis interface{} `field:"optional" json:"weeklyRepetitionBasis" yaml:"weeklyRepetitionBasis"`
}

