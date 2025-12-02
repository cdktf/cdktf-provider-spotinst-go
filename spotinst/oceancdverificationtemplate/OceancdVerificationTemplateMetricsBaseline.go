// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceancdverificationtemplate


type OceancdVerificationTemplateMetricsBaseline struct {
	// baseline_provider block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.231.0/docs/resources/oceancd_verification_template#baseline_provider OceancdVerificationTemplate#baseline_provider}
	BaselineProvider interface{} `field:"required" json:"baselineProvider" yaml:"baselineProvider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.231.0/docs/resources/oceancd_verification_template#threshold OceancdVerificationTemplate#threshold}.
	Threshold *string `field:"required" json:"threshold" yaml:"threshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.231.0/docs/resources/oceancd_verification_template#max_range OceancdVerificationTemplate#max_range}.
	MaxRange *float64 `field:"optional" json:"maxRange" yaml:"maxRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.231.0/docs/resources/oceancd_verification_template#min_range OceancdVerificationTemplate#min_range}.
	MinRange *float64 `field:"optional" json:"minRange" yaml:"minRange"`
}

