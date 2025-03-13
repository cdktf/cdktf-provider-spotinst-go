// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanrightsizingrule


type OceanRightSizingRuleRecommendationApplicationOverheadValues struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.212.0/docs/resources/ocean_right_sizing_rule#cpu_percentage OceanRightSizingRule#cpu_percentage}.
	CpuPercentage *float64 `field:"optional" json:"cpuPercentage" yaml:"cpuPercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.212.0/docs/resources/ocean_right_sizing_rule#memory_percentage OceanRightSizingRule#memory_percentage}.
	MemoryPercentage *float64 `field:"optional" json:"memoryPercentage" yaml:"memoryPercentage"`
}

