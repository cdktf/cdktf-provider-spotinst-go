// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanrightsizingrule


type OceanRightSizingRuleRecommendationApplicationIntervalsMonthlyRepetitionBasisWeeklyRepetitionBasis struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.182.0/docs/resources/ocean_right_sizing_rule#interval_days OceanRightSizingRule#interval_days}.
	IntervalDays *[]*string `field:"required" json:"intervalDays" yaml:"intervalDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.182.0/docs/resources/ocean_right_sizing_rule#interval_hours_end_time OceanRightSizingRule#interval_hours_end_time}.
	IntervalHoursEndTime *string `field:"required" json:"intervalHoursEndTime" yaml:"intervalHoursEndTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.182.0/docs/resources/ocean_right_sizing_rule#interval_hours_start_time OceanRightSizingRule#interval_hours_start_time}.
	IntervalHoursStartTime *string `field:"required" json:"intervalHoursStartTime" yaml:"intervalHoursStartTime"`
}

