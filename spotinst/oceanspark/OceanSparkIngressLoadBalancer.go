// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanspark


type OceanSparkIngressLoadBalancer struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.171.4/docs/resources/ocean_spark#managed OceanSpark#managed}.
	Managed interface{} `field:"optional" json:"managed" yaml:"managed"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.171.4/docs/resources/ocean_spark#service_annotations OceanSpark#service_annotations}.
	ServiceAnnotations *map[string]*string `field:"optional" json:"serviceAnnotations" yaml:"serviceAnnotations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.171.4/docs/resources/ocean_spark#target_group_arn OceanSpark#target_group_arn}.
	TargetGroupArn *string `field:"optional" json:"targetGroupArn" yaml:"targetGroupArn"`
}

