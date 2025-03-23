// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceancdrolloutspec


type OceancdRolloutSpecTrafficAlb struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.214.0/docs/resources/oceancd_rollout_spec#alb_ingress OceancdRolloutSpec#alb_ingress}.
	AlbIngress *string `field:"required" json:"albIngress" yaml:"albIngress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.214.0/docs/resources/oceancd_rollout_spec#alb_root_service OceancdRolloutSpec#alb_root_service}.
	AlbRootService *string `field:"required" json:"albRootService" yaml:"albRootService"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.214.0/docs/resources/oceancd_rollout_spec#service_port OceancdRolloutSpec#service_port}.
	ServicePort *float64 `field:"required" json:"servicePort" yaml:"servicePort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.214.0/docs/resources/oceancd_rollout_spec#alb_annotation_prefix OceancdRolloutSpec#alb_annotation_prefix}.
	AlbAnnotationPrefix *string `field:"optional" json:"albAnnotationPrefix" yaml:"albAnnotationPrefix"`
	// stickiness_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.214.0/docs/resources/oceancd_rollout_spec#stickiness_config OceancdRolloutSpec#stickiness_config}
	StickinessConfig *OceancdRolloutSpecTrafficAlbStickinessConfig `field:"optional" json:"stickinessConfig" yaml:"stickinessConfig"`
}

