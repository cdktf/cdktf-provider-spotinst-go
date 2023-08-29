// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanspark


type OceanSparkIngressCustomEndpoint struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.134.0/docs/resources/ocean_spark#address OceanSpark#address}.
	Address *string `field:"optional" json:"address" yaml:"address"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.134.0/docs/resources/ocean_spark#enabled OceanSpark#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

