// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceancdverificationtemplate


type OceancdVerificationTemplateMetricsBaselineBaselineProvider struct {
	// datadog block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.205.1/docs/resources/oceancd_verification_template#datadog OceancdVerificationTemplate#datadog}
	Datadog *OceancdVerificationTemplateMetricsBaselineBaselineProviderDatadog `field:"optional" json:"datadog" yaml:"datadog"`
	// new_relic block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.205.1/docs/resources/oceancd_verification_template#new_relic OceancdVerificationTemplate#new_relic}
	NewRelic *OceancdVerificationTemplateMetricsBaselineBaselineProviderNewRelic `field:"optional" json:"newRelic" yaml:"newRelic"`
	// prometheus block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.205.1/docs/resources/oceancd_verification_template#prometheus OceancdVerificationTemplate#prometheus}
	Prometheus *OceancdVerificationTemplateMetricsBaselineBaselineProviderPrometheus `field:"optional" json:"prometheus" yaml:"prometheus"`
}

