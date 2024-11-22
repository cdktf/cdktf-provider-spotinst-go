// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaksnpvirtualnodegroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OceanAksNpVirtualNodeGroupConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.0/docs/resources/ocean_aks_np_virtual_node_group#name OceanAksNpVirtualNodeGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.0/docs/resources/ocean_aks_np_virtual_node_group#ocean_id OceanAksNpVirtualNodeGroup#ocean_id}.
	OceanId *string `field:"required" json:"oceanId" yaml:"oceanId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.0/docs/resources/ocean_aks_np_virtual_node_group#availability_zones OceanAksNpVirtualNodeGroup#availability_zones}.
	AvailabilityZones *[]*string `field:"optional" json:"availabilityZones" yaml:"availabilityZones"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.0/docs/resources/ocean_aks_np_virtual_node_group#enable_node_public_ip OceanAksNpVirtualNodeGroup#enable_node_public_ip}.
	EnableNodePublicIp interface{} `field:"optional" json:"enableNodePublicIp" yaml:"enableNodePublicIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.0/docs/resources/ocean_aks_np_virtual_node_group#fallback_to_ondemand OceanAksNpVirtualNodeGroup#fallback_to_ondemand}.
	FallbackToOndemand interface{} `field:"optional" json:"fallbackToOndemand" yaml:"fallbackToOndemand"`
	// filters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.0/docs/resources/ocean_aks_np_virtual_node_group#filters OceanAksNpVirtualNodeGroup#filters}
	Filters *OceanAksNpVirtualNodeGroupFilters `field:"optional" json:"filters" yaml:"filters"`
	// headrooms block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.0/docs/resources/ocean_aks_np_virtual_node_group#headrooms OceanAksNpVirtualNodeGroup#headrooms}
	Headrooms interface{} `field:"optional" json:"headrooms" yaml:"headrooms"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.0/docs/resources/ocean_aks_np_virtual_node_group#id OceanAksNpVirtualNodeGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.0/docs/resources/ocean_aks_np_virtual_node_group#kubernetes_version OceanAksNpVirtualNodeGroup#kubernetes_version}.
	KubernetesVersion *string `field:"optional" json:"kubernetesVersion" yaml:"kubernetesVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.0/docs/resources/ocean_aks_np_virtual_node_group#labels OceanAksNpVirtualNodeGroup#labels}.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// linux_os_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.0/docs/resources/ocean_aks_np_virtual_node_group#linux_os_config OceanAksNpVirtualNodeGroup#linux_os_config}
	LinuxOsConfig interface{} `field:"optional" json:"linuxOsConfig" yaml:"linuxOsConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.0/docs/resources/ocean_aks_np_virtual_node_group#max_count OceanAksNpVirtualNodeGroup#max_count}.
	MaxCount *float64 `field:"optional" json:"maxCount" yaml:"maxCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.0/docs/resources/ocean_aks_np_virtual_node_group#max_pods_per_node OceanAksNpVirtualNodeGroup#max_pods_per_node}.
	MaxPodsPerNode *float64 `field:"optional" json:"maxPodsPerNode" yaml:"maxPodsPerNode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.0/docs/resources/ocean_aks_np_virtual_node_group#min_count OceanAksNpVirtualNodeGroup#min_count}.
	MinCount *float64 `field:"optional" json:"minCount" yaml:"minCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.0/docs/resources/ocean_aks_np_virtual_node_group#os_disk_size_gb OceanAksNpVirtualNodeGroup#os_disk_size_gb}.
	OsDiskSizeGb *float64 `field:"optional" json:"osDiskSizeGb" yaml:"osDiskSizeGb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.0/docs/resources/ocean_aks_np_virtual_node_group#os_disk_type OceanAksNpVirtualNodeGroup#os_disk_type}.
	OsDiskType *string `field:"optional" json:"osDiskType" yaml:"osDiskType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.0/docs/resources/ocean_aks_np_virtual_node_group#os_sku OceanAksNpVirtualNodeGroup#os_sku}.
	OsSku *string `field:"optional" json:"osSku" yaml:"osSku"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.0/docs/resources/ocean_aks_np_virtual_node_group#os_type OceanAksNpVirtualNodeGroup#os_type}.
	OsType *string `field:"optional" json:"osType" yaml:"osType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.0/docs/resources/ocean_aks_np_virtual_node_group#pod_subnet_ids OceanAksNpVirtualNodeGroup#pod_subnet_ids}.
	PodSubnetIds *[]*string `field:"optional" json:"podSubnetIds" yaml:"podSubnetIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.0/docs/resources/ocean_aks_np_virtual_node_group#spot_percentage OceanAksNpVirtualNodeGroup#spot_percentage}.
	SpotPercentage *float64 `field:"optional" json:"spotPercentage" yaml:"spotPercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.0/docs/resources/ocean_aks_np_virtual_node_group#tags OceanAksNpVirtualNodeGroup#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// taints block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.0/docs/resources/ocean_aks_np_virtual_node_group#taints OceanAksNpVirtualNodeGroup#taints}
	Taints interface{} `field:"optional" json:"taints" yaml:"taints"`
	// update_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.0/docs/resources/ocean_aks_np_virtual_node_group#update_policy OceanAksNpVirtualNodeGroup#update_policy}
	UpdatePolicy *OceanAksNpVirtualNodeGroupUpdatePolicy `field:"optional" json:"updatePolicy" yaml:"updatePolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.0/docs/resources/ocean_aks_np_virtual_node_group#vnet_subnet_ids OceanAksNpVirtualNodeGroup#vnet_subnet_ids}.
	VnetSubnetIds *[]*string `field:"optional" json:"vnetSubnetIds" yaml:"vnetSubnetIds"`
}

