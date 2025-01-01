// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceancdverificationprovider


type OceancdVerificationProviderNewRelic struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.203.0/docs/resources/oceancd_verification_provider#account_id OceancdVerificationProvider#account_id}.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.203.0/docs/resources/oceancd_verification_provider#personal_api_key OceancdVerificationProvider#personal_api_key}.
	PersonalApiKey *string `field:"required" json:"personalApiKey" yaml:"personalApiKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.203.0/docs/resources/oceancd_verification_provider#base_url_nerd_graph OceancdVerificationProvider#base_url_nerd_graph}.
	BaseUrlNerdGraph *string `field:"optional" json:"baseUrlNerdGraph" yaml:"baseUrlNerdGraph"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.203.0/docs/resources/oceancd_verification_provider#base_url_rest OceancdVerificationProvider#base_url_rest}.
	BaseUrlRest *string `field:"optional" json:"baseUrlRest" yaml:"baseUrlRest"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.203.0/docs/resources/oceancd_verification_provider#region OceancdVerificationProvider#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

