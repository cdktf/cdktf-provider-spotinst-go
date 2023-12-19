// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaks


type OceanAksLoadBalancer struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.156.0/docs/resources/ocean_aks#backend_pool_names OceanAks#backend_pool_names}.
	BackendPoolNames *[]*string `field:"optional" json:"backendPoolNames" yaml:"backendPoolNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.156.0/docs/resources/ocean_aks#load_balancer_sku OceanAks#load_balancer_sku}.
	LoadBalancerSku *string `field:"optional" json:"loadBalancerSku" yaml:"loadBalancerSku"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.156.0/docs/resources/ocean_aks#name OceanAks#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.156.0/docs/resources/ocean_aks#resource_group_name OceanAks#resource_group_name}.
	ResourceGroupName *string `field:"optional" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.156.0/docs/resources/ocean_aks#type OceanAks#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

