// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceancdrolloutspec


type OceancdRolloutSpecTraffic struct {
	// alb block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.222.1/docs/resources/oceancd_rollout_spec#alb OceancdRolloutSpec#alb}
	Alb *OceancdRolloutSpecTrafficAlb `field:"optional" json:"alb" yaml:"alb"`
	// ambassador block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.222.1/docs/resources/oceancd_rollout_spec#ambassador OceancdRolloutSpec#ambassador}
	Ambassador *OceancdRolloutSpecTrafficAmbassador `field:"optional" json:"ambassador" yaml:"ambassador"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.222.1/docs/resources/oceancd_rollout_spec#canary_service OceancdRolloutSpec#canary_service}.
	CanaryService *string `field:"optional" json:"canaryService" yaml:"canaryService"`
	// istio block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.222.1/docs/resources/oceancd_rollout_spec#istio OceancdRolloutSpec#istio}
	Istio *OceancdRolloutSpecTrafficIstio `field:"optional" json:"istio" yaml:"istio"`
	// nginx block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.222.1/docs/resources/oceancd_rollout_spec#nginx OceancdRolloutSpec#nginx}
	Nginx *OceancdRolloutSpecTrafficNginx `field:"optional" json:"nginx" yaml:"nginx"`
	// ping_pong block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.222.1/docs/resources/oceancd_rollout_spec#ping_pong OceancdRolloutSpec#ping_pong}
	PingPong *OceancdRolloutSpecTrafficPingPong `field:"optional" json:"pingPong" yaml:"pingPong"`
	// smi block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.222.1/docs/resources/oceancd_rollout_spec#smi OceancdRolloutSpec#smi}
	Smi *OceancdRolloutSpecTrafficSmi `field:"optional" json:"smi" yaml:"smi"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.222.1/docs/resources/oceancd_rollout_spec#stable_service OceancdRolloutSpec#stable_service}.
	StableService *string `field:"optional" json:"stableService" yaml:"stableService"`
}

