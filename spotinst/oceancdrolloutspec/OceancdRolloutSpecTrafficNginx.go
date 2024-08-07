// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceancdrolloutspec


type OceancdRolloutSpecTrafficNginx struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.184.0/docs/resources/oceancd_rollout_spec#stable_ingress OceancdRolloutSpec#stable_ingress}.
	StableIngress *string `field:"required" json:"stableIngress" yaml:"stableIngress"`
	// additional_ingress_annotation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.184.0/docs/resources/oceancd_rollout_spec#additional_ingress_annotation OceancdRolloutSpec#additional_ingress_annotation}
	AdditionalIngressAnnotation *OceancdRolloutSpecTrafficNginxAdditionalIngressAnnotation `field:"optional" json:"additionalIngressAnnotation" yaml:"additionalIngressAnnotation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.184.0/docs/resources/oceancd_rollout_spec#nginx_annotation_prefix OceancdRolloutSpec#nginx_annotation_prefix}.
	NginxAnnotationPrefix *string `field:"optional" json:"nginxAnnotationPrefix" yaml:"nginxAnnotationPrefix"`
}

