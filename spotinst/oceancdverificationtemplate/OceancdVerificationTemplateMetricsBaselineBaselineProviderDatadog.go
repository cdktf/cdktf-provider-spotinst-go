// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceancdverificationtemplate


type OceancdVerificationTemplateMetricsBaselineBaselineProviderDatadog struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.220.0/docs/resources/oceancd_verification_template#datadog_query OceancdVerificationTemplate#datadog_query}.
	DatadogQuery *string `field:"required" json:"datadogQuery" yaml:"datadogQuery"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.220.0/docs/resources/oceancd_verification_template#duration OceancdVerificationTemplate#duration}.
	Duration *string `field:"optional" json:"duration" yaml:"duration"`
}

