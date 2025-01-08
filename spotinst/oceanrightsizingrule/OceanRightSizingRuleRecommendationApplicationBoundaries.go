// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanrightsizingrule


type OceanRightSizingRuleRecommendationApplicationBoundaries struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.205.1/docs/resources/ocean_right_sizing_rule#cpu_max OceanRightSizingRule#cpu_max}.
	CpuMax *float64 `field:"optional" json:"cpuMax" yaml:"cpuMax"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.205.1/docs/resources/ocean_right_sizing_rule#cpu_min OceanRightSizingRule#cpu_min}.
	CpuMin *float64 `field:"optional" json:"cpuMin" yaml:"cpuMin"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.205.1/docs/resources/ocean_right_sizing_rule#memory_max OceanRightSizingRule#memory_max}.
	MemoryMax *float64 `field:"optional" json:"memoryMax" yaml:"memoryMax"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.205.1/docs/resources/ocean_right_sizing_rule#memory_min OceanRightSizingRule#memory_min}.
	MemoryMin *float64 `field:"optional" json:"memoryMin" yaml:"memoryMin"`
}

