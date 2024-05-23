// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceancdverificationtemplate


type OceancdVerificationTemplateMetricsBaselineBaselineProviderNewRelic struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.175.0/docs/resources/oceancd_verification_template#new_relic_query OceancdVerificationTemplate#new_relic_query}.
	NewRelicQuery *string `field:"required" json:"newRelicQuery" yaml:"newRelicQuery"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.175.0/docs/resources/oceancd_verification_template#profile OceancdVerificationTemplate#profile}.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
}

