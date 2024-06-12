// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceancdverificationprovider


type OceancdVerificationProviderJenkins struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.176.0/docs/resources/oceancd_verification_provider#api_token OceancdVerificationProvider#api_token}.
	ApiToken *string `field:"required" json:"apiToken" yaml:"apiToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.176.0/docs/resources/oceancd_verification_provider#base_url OceancdVerificationProvider#base_url}.
	BaseUrl *string `field:"required" json:"baseUrl" yaml:"baseUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.176.0/docs/resources/oceancd_verification_provider#username OceancdVerificationProvider#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
}

