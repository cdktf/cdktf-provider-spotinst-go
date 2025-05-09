// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type SpotinstProviderConfig struct {
	// Spotinst Account ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.219.0/docs#account SpotinstProvider#account}
	Account *string `field:"optional" json:"account" yaml:"account"`
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.219.0/docs#alias SpotinstProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// Enable or disable the Spotinst provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.219.0/docs#enabled SpotinstProvider#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Spotinst SDK Feature Flags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.219.0/docs#feature_flags SpotinstProvider#feature_flags}
	FeatureFlags *string `field:"optional" json:"featureFlags" yaml:"featureFlags"`
	// Spotinst Personal API Access Token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.219.0/docs#token SpotinstProvider#token}
	Token *string `field:"optional" json:"token" yaml:"token"`
}

