// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceancdstrategy


type OceancdStrategyRollingStepsVerification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.206.0/docs/resources/oceancd_strategy#template_names OceancdStrategy#template_names}.
	TemplateNames *[]*string `field:"required" json:"templateNames" yaml:"templateNames"`
}

