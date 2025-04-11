// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceancdstrategy


type OceancdStrategyCanaryStepsSetCanaryScale struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.216.1/docs/resources/oceancd_strategy#match_traffic_weight OceancdStrategy#match_traffic_weight}.
	MatchTrafficWeight interface{} `field:"optional" json:"matchTrafficWeight" yaml:"matchTrafficWeight"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.216.1/docs/resources/oceancd_strategy#replicas OceancdStrategy#replicas}.
	Replicas *float64 `field:"optional" json:"replicas" yaml:"replicas"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.216.1/docs/resources/oceancd_strategy#weight OceancdStrategy#weight}.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

