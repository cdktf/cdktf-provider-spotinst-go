// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanspark


type OceanSparkIngress struct {
	// controller block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.148.0/docs/resources/ocean_spark#controller OceanSpark#controller}
	Controller *OceanSparkIngressController `field:"optional" json:"controller" yaml:"controller"`
	// custom_endpoint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.148.0/docs/resources/ocean_spark#custom_endpoint OceanSpark#custom_endpoint}
	CustomEndpoint *OceanSparkIngressCustomEndpoint `field:"optional" json:"customEndpoint" yaml:"customEndpoint"`
	// load_balancer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.148.0/docs/resources/ocean_spark#load_balancer OceanSpark#load_balancer}
	LoadBalancer *OceanSparkIngressLoadBalancer `field:"optional" json:"loadBalancer" yaml:"loadBalancer"`
	// private_link block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.148.0/docs/resources/ocean_spark#private_link OceanSpark#private_link}
	PrivateLink *OceanSparkIngressPrivateLink `field:"optional" json:"privateLink" yaml:"privateLink"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.148.0/docs/resources/ocean_spark#service_annotations OceanSpark#service_annotations}.
	ServiceAnnotations *map[string]*string `field:"optional" json:"serviceAnnotations" yaml:"serviceAnnotations"`
}

