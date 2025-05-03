// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceancdrolloutspec


type OceancdRolloutSpecTrafficPingPong struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/oceancd_rollout_spec#ping_service OceancdRolloutSpec#ping_service}.
	PingService *string `field:"required" json:"pingService" yaml:"pingService"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.218.0/docs/resources/oceancd_rollout_spec#pong_service OceancdRolloutSpec#pong_service}.
	PongService *string `field:"required" json:"pongService" yaml:"pongService"`
}

