// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanrightsizingrule


type OceanRightSizingRuleAutoApplyDefinition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.228.0/docs/resources/ocean_right_sizing_rule#enabled OceanRightSizingRule#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.228.0/docs/resources/ocean_right_sizing_rule#labels OceanRightSizingRule#labels}.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.228.0/docs/resources/ocean_right_sizing_rule#namespaces OceanRightSizingRule#namespaces}.
	Namespaces *[]*string `field:"optional" json:"namespaces" yaml:"namespaces"`
}

