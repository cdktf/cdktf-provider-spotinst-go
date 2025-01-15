// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaksnp

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OceanAksNpConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.207.0/docs/resources/ocean_aks_np#aks_cluster_name OceanAksNp#aks_cluster_name}.
	AksClusterName *string `field:"required" json:"aksClusterName" yaml:"aksClusterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.207.0/docs/resources/ocean_aks_np#aks_infrastructure_resource_group_name OceanAksNp#aks_infrastructure_resource_group_name}.
	AksInfrastructureResourceGroupName *string `field:"required" json:"aksInfrastructureResourceGroupName" yaml:"aksInfrastructureResourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.207.0/docs/resources/ocean_aks_np#aks_region OceanAksNp#aks_region}.
	AksRegion *string `field:"required" json:"aksRegion" yaml:"aksRegion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.207.0/docs/resources/ocean_aks_np#aks_resource_group_name OceanAksNp#aks_resource_group_name}.
	AksResourceGroupName *string `field:"required" json:"aksResourceGroupName" yaml:"aksResourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.207.0/docs/resources/ocean_aks_np#availability_zones OceanAksNp#availability_zones}.
	AvailabilityZones *[]*string `field:"required" json:"availabilityZones" yaml:"availabilityZones"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.207.0/docs/resources/ocean_aks_np#controller_cluster_id OceanAksNp#controller_cluster_id}.
	ControllerClusterId *string `field:"required" json:"controllerClusterId" yaml:"controllerClusterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.207.0/docs/resources/ocean_aks_np#name OceanAksNp#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// autoscaler block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.207.0/docs/resources/ocean_aks_np#autoscaler OceanAksNp#autoscaler}
	Autoscaler *OceanAksNpAutoscaler `field:"optional" json:"autoscaler" yaml:"autoscaler"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.207.0/docs/resources/ocean_aks_np#enable_node_public_ip OceanAksNp#enable_node_public_ip}.
	EnableNodePublicIp interface{} `field:"optional" json:"enableNodePublicIp" yaml:"enableNodePublicIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.207.0/docs/resources/ocean_aks_np#fallback_to_ondemand OceanAksNp#fallback_to_ondemand}.
	FallbackToOndemand interface{} `field:"optional" json:"fallbackToOndemand" yaml:"fallbackToOndemand"`
	// filters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.207.0/docs/resources/ocean_aks_np#filters OceanAksNp#filters}
	Filters *OceanAksNpFilters `field:"optional" json:"filters" yaml:"filters"`
	// headrooms block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.207.0/docs/resources/ocean_aks_np#headrooms OceanAksNp#headrooms}
	Headrooms interface{} `field:"optional" json:"headrooms" yaml:"headrooms"`
	// health block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.207.0/docs/resources/ocean_aks_np#health OceanAksNp#health}
	Health *OceanAksNpHealth `field:"optional" json:"health" yaml:"health"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.207.0/docs/resources/ocean_aks_np#id OceanAksNp#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.207.0/docs/resources/ocean_aks_np#kubernetes_version OceanAksNp#kubernetes_version}.
	KubernetesVersion *string `field:"optional" json:"kubernetesVersion" yaml:"kubernetesVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.207.0/docs/resources/ocean_aks_np#labels OceanAksNp#labels}.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// linux_os_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.207.0/docs/resources/ocean_aks_np#linux_os_config OceanAksNp#linux_os_config}
	LinuxOsConfig interface{} `field:"optional" json:"linuxOsConfig" yaml:"linuxOsConfig"`
	// logging block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.207.0/docs/resources/ocean_aks_np#logging OceanAksNp#logging}
	Logging *OceanAksNpLogging `field:"optional" json:"logging" yaml:"logging"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.207.0/docs/resources/ocean_aks_np#max_count OceanAksNp#max_count}.
	MaxCount *float64 `field:"optional" json:"maxCount" yaml:"maxCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.207.0/docs/resources/ocean_aks_np#max_pods_per_node OceanAksNp#max_pods_per_node}.
	MaxPodsPerNode *float64 `field:"optional" json:"maxPodsPerNode" yaml:"maxPodsPerNode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.207.0/docs/resources/ocean_aks_np#min_count OceanAksNp#min_count}.
	MinCount *float64 `field:"optional" json:"minCount" yaml:"minCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.207.0/docs/resources/ocean_aks_np#os_disk_size_gb OceanAksNp#os_disk_size_gb}.
	OsDiskSizeGb *float64 `field:"optional" json:"osDiskSizeGb" yaml:"osDiskSizeGb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.207.0/docs/resources/ocean_aks_np#os_disk_type OceanAksNp#os_disk_type}.
	OsDiskType *string `field:"optional" json:"osDiskType" yaml:"osDiskType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.207.0/docs/resources/ocean_aks_np#os_sku OceanAksNp#os_sku}.
	OsSku *string `field:"optional" json:"osSku" yaml:"osSku"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.207.0/docs/resources/ocean_aks_np#os_type OceanAksNp#os_type}.
	OsType *string `field:"optional" json:"osType" yaml:"osType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.207.0/docs/resources/ocean_aks_np#pod_subnet_ids OceanAksNp#pod_subnet_ids}.
	PodSubnetIds *[]*string `field:"optional" json:"podSubnetIds" yaml:"podSubnetIds"`
	// scheduling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.207.0/docs/resources/ocean_aks_np#scheduling OceanAksNp#scheduling}
	Scheduling *OceanAksNpScheduling `field:"optional" json:"scheduling" yaml:"scheduling"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.207.0/docs/resources/ocean_aks_np#spot_percentage OceanAksNp#spot_percentage}.
	SpotPercentage *float64 `field:"optional" json:"spotPercentage" yaml:"spotPercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.207.0/docs/resources/ocean_aks_np#tags OceanAksNp#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// taints block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.207.0/docs/resources/ocean_aks_np#taints OceanAksNp#taints}
	Taints interface{} `field:"optional" json:"taints" yaml:"taints"`
	// update_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.207.0/docs/resources/ocean_aks_np#update_policy OceanAksNp#update_policy}
	UpdatePolicy *OceanAksNpUpdatePolicy `field:"optional" json:"updatePolicy" yaml:"updatePolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.207.0/docs/resources/ocean_aks_np#vnet_subnet_ids OceanAksNp#vnet_subnet_ids}.
	VnetSubnetIds *[]*string `field:"optional" json:"vnetSubnetIds" yaml:"vnetSubnetIds"`
	// vng_template_scheduling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.207.0/docs/resources/ocean_aks_np#vng_template_scheduling OceanAksNp#vng_template_scheduling}
	VngTemplateScheduling *OceanAksNpVngTemplateScheduling `field:"optional" json:"vngTemplateScheduling" yaml:"vngTemplateScheduling"`
}

