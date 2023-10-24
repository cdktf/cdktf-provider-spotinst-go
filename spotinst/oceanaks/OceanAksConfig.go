// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaks

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OceanAksConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.148.0/docs/resources/ocean_aks#acd_identifier OceanAks#acd_identifier}.
	AcdIdentifier *string `field:"required" json:"acdIdentifier" yaml:"acdIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.148.0/docs/resources/ocean_aks#aks_name OceanAks#aks_name}.
	AksName *string `field:"required" json:"aksName" yaml:"aksName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.148.0/docs/resources/ocean_aks#aks_resource_group_name OceanAks#aks_resource_group_name}.
	AksResourceGroupName *string `field:"required" json:"aksResourceGroupName" yaml:"aksResourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.148.0/docs/resources/ocean_aks#name OceanAks#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.148.0/docs/resources/ocean_aks#ssh_public_key OceanAks#ssh_public_key}.
	SshPublicKey *string `field:"required" json:"sshPublicKey" yaml:"sshPublicKey"`
	// autoscaler block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.148.0/docs/resources/ocean_aks#autoscaler OceanAks#autoscaler}
	Autoscaler *OceanAksAutoscaler `field:"optional" json:"autoscaler" yaml:"autoscaler"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.148.0/docs/resources/ocean_aks#controller_cluster_id OceanAks#controller_cluster_id}.
	ControllerClusterId *string `field:"optional" json:"controllerClusterId" yaml:"controllerClusterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.148.0/docs/resources/ocean_aks#custom_data OceanAks#custom_data}.
	CustomData *string `field:"optional" json:"customData" yaml:"customData"`
	// extension block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.148.0/docs/resources/ocean_aks#extension OceanAks#extension}
	Extension interface{} `field:"optional" json:"extension" yaml:"extension"`
	// health block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.148.0/docs/resources/ocean_aks#health OceanAks#health}
	Health *OceanAksHealth `field:"optional" json:"health" yaml:"health"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.148.0/docs/resources/ocean_aks#id OceanAks#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// image block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.148.0/docs/resources/ocean_aks#image OceanAks#image}
	Image interface{} `field:"optional" json:"image" yaml:"image"`
	// load_balancer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.148.0/docs/resources/ocean_aks#load_balancer OceanAks#load_balancer}
	LoadBalancer interface{} `field:"optional" json:"loadBalancer" yaml:"loadBalancer"`
	// managed_service_identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.148.0/docs/resources/ocean_aks#managed_service_identity OceanAks#managed_service_identity}
	ManagedServiceIdentity interface{} `field:"optional" json:"managedServiceIdentity" yaml:"managedServiceIdentity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.148.0/docs/resources/ocean_aks#max_pods OceanAks#max_pods}.
	MaxPods *float64 `field:"optional" json:"maxPods" yaml:"maxPods"`
	// network block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.148.0/docs/resources/ocean_aks#network OceanAks#network}
	Network *OceanAksNetwork `field:"optional" json:"network" yaml:"network"`
	// os_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.148.0/docs/resources/ocean_aks#os_disk OceanAks#os_disk}
	OsDisk *OceanAksOsDisk `field:"optional" json:"osDisk" yaml:"osDisk"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.148.0/docs/resources/ocean_aks#resource_group_name OceanAks#resource_group_name}.
	ResourceGroupName *string `field:"optional" json:"resourceGroupName" yaml:"resourceGroupName"`
	// strategy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.148.0/docs/resources/ocean_aks#strategy OceanAks#strategy}
	Strategy interface{} `field:"optional" json:"strategy" yaml:"strategy"`
	// tag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.148.0/docs/resources/ocean_aks#tag OceanAks#tag}
	Tag interface{} `field:"optional" json:"tag" yaml:"tag"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.148.0/docs/resources/ocean_aks#user_name OceanAks#user_name}.
	UserName *string `field:"optional" json:"userName" yaml:"userName"`
	// vm_sizes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.148.0/docs/resources/ocean_aks#vm_sizes OceanAks#vm_sizes}
	VmSizes interface{} `field:"optional" json:"vmSizes" yaml:"vmSizes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.148.0/docs/resources/ocean_aks#zones OceanAks#zones}.
	Zones *[]*string `field:"optional" json:"zones" yaml:"zones"`
}

