// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceancdrolloutspec


type OceancdRolloutSpecTrafficAlbStickinessConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/oceancd_rollout_spec#duration_seconds OceancdRolloutSpec#duration_seconds}.
	DurationSeconds *float64 `field:"optional" json:"durationSeconds" yaml:"durationSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.180.0/docs/resources/oceancd_rollout_spec#enabled OceancdRolloutSpec#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

