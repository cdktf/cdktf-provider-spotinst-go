// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceancdverificationprovider


type OceancdVerificationProviderDatadog struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.1/docs/resources/oceancd_verification_provider#address OceancdVerificationProvider#address}.
	Address *string `field:"required" json:"address" yaml:"address"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.1/docs/resources/oceancd_verification_provider#api_key OceancdVerificationProvider#api_key}.
	ApiKey *string `field:"required" json:"apiKey" yaml:"apiKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.1/docs/resources/oceancd_verification_provider#app_key OceancdVerificationProvider#app_key}.
	AppKey *string `field:"required" json:"appKey" yaml:"appKey"`
}

