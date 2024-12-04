// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupazurev3

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ElastigroupAzureV3Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.3/docs/resources/elastigroup_azure_v3#fallback_to_on_demand ElastigroupAzureV3#fallback_to_on_demand}.
	FallbackToOnDemand interface{} `field:"required" json:"fallbackToOnDemand" yaml:"fallbackToOnDemand"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.3/docs/resources/elastigroup_azure_v3#name ElastigroupAzureV3#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// network block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.3/docs/resources/elastigroup_azure_v3#network ElastigroupAzureV3#network}
	Network *ElastigroupAzureV3Network `field:"required" json:"network" yaml:"network"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.3/docs/resources/elastigroup_azure_v3#os ElastigroupAzureV3#os}.
	Os *string `field:"required" json:"os" yaml:"os"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.3/docs/resources/elastigroup_azure_v3#region ElastigroupAzureV3#region}.
	Region *string `field:"required" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.3/docs/resources/elastigroup_azure_v3#resource_group_name ElastigroupAzureV3#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// vm_sizes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.3/docs/resources/elastigroup_azure_v3#vm_sizes ElastigroupAzureV3#vm_sizes}
	VmSizes *ElastigroupAzureV3VmSizes `field:"required" json:"vmSizes" yaml:"vmSizes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.3/docs/resources/elastigroup_azure_v3#availability_vs_cost ElastigroupAzureV3#availability_vs_cost}.
	AvailabilityVsCost *float64 `field:"optional" json:"availabilityVsCost" yaml:"availabilityVsCost"`
	// boot_diagnostics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.3/docs/resources/elastigroup_azure_v3#boot_diagnostics ElastigroupAzureV3#boot_diagnostics}
	BootDiagnostics interface{} `field:"optional" json:"bootDiagnostics" yaml:"bootDiagnostics"`
	// capacity_reservation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.3/docs/resources/elastigroup_azure_v3#capacity_reservation ElastigroupAzureV3#capacity_reservation}
	CapacityReservation *ElastigroupAzureV3CapacityReservation `field:"optional" json:"capacityReservation" yaml:"capacityReservation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.3/docs/resources/elastigroup_azure_v3#custom_data ElastigroupAzureV3#custom_data}.
	CustomData *string `field:"optional" json:"customData" yaml:"customData"`
	// data_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.3/docs/resources/elastigroup_azure_v3#data_disk ElastigroupAzureV3#data_disk}
	DataDisk interface{} `field:"optional" json:"dataDisk" yaml:"dataDisk"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.3/docs/resources/elastigroup_azure_v3#description ElastigroupAzureV3#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.3/docs/resources/elastigroup_azure_v3#desired_capacity ElastigroupAzureV3#desired_capacity}.
	DesiredCapacity *float64 `field:"optional" json:"desiredCapacity" yaml:"desiredCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.3/docs/resources/elastigroup_azure_v3#draining_timeout ElastigroupAzureV3#draining_timeout}.
	DrainingTimeout *float64 `field:"optional" json:"drainingTimeout" yaml:"drainingTimeout"`
	// extensions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.3/docs/resources/elastigroup_azure_v3#extensions ElastigroupAzureV3#extensions}
	Extensions interface{} `field:"optional" json:"extensions" yaml:"extensions"`
	// health block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.3/docs/resources/elastigroup_azure_v3#health ElastigroupAzureV3#health}
	Health *ElastigroupAzureV3Health `field:"optional" json:"health" yaml:"health"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.3/docs/resources/elastigroup_azure_v3#id ElastigroupAzureV3#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// image block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.3/docs/resources/elastigroup_azure_v3#image ElastigroupAzureV3#image}
	Image interface{} `field:"optional" json:"image" yaml:"image"`
	// load_balancer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.3/docs/resources/elastigroup_azure_v3#load_balancer ElastigroupAzureV3#load_balancer}
	LoadBalancer interface{} `field:"optional" json:"loadBalancer" yaml:"loadBalancer"`
	// login block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.3/docs/resources/elastigroup_azure_v3#login ElastigroupAzureV3#login}
	Login *ElastigroupAzureV3Login `field:"optional" json:"login" yaml:"login"`
	// managed_service_identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.3/docs/resources/elastigroup_azure_v3#managed_service_identity ElastigroupAzureV3#managed_service_identity}
	ManagedServiceIdentity interface{} `field:"optional" json:"managedServiceIdentity" yaml:"managedServiceIdentity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.3/docs/resources/elastigroup_azure_v3#max_size ElastigroupAzureV3#max_size}.
	MaxSize *float64 `field:"optional" json:"maxSize" yaml:"maxSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.3/docs/resources/elastigroup_azure_v3#min_size ElastigroupAzureV3#min_size}.
	MinSize *float64 `field:"optional" json:"minSize" yaml:"minSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.3/docs/resources/elastigroup_azure_v3#on_demand_count ElastigroupAzureV3#on_demand_count}.
	OnDemandCount *float64 `field:"optional" json:"onDemandCount" yaml:"onDemandCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.3/docs/resources/elastigroup_azure_v3#optimization_windows ElastigroupAzureV3#optimization_windows}.
	OptimizationWindows *[]*string `field:"optional" json:"optimizationWindows" yaml:"optimizationWindows"`
	// os_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.3/docs/resources/elastigroup_azure_v3#os_disk ElastigroupAzureV3#os_disk}
	OsDisk *ElastigroupAzureV3OsDisk `field:"optional" json:"osDisk" yaml:"osDisk"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.3/docs/resources/elastigroup_azure_v3#preferred_zones ElastigroupAzureV3#preferred_zones}.
	PreferredZones *[]*string `field:"optional" json:"preferredZones" yaml:"preferredZones"`
	// proximity_placement_groups block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.3/docs/resources/elastigroup_azure_v3#proximity_placement_groups ElastigroupAzureV3#proximity_placement_groups}
	ProximityPlacementGroups interface{} `field:"optional" json:"proximityPlacementGroups" yaml:"proximityPlacementGroups"`
	// revert_to_spot block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.3/docs/resources/elastigroup_azure_v3#revert_to_spot ElastigroupAzureV3#revert_to_spot}
	RevertToSpot *ElastigroupAzureV3RevertToSpot `field:"optional" json:"revertToSpot" yaml:"revertToSpot"`
	// scaling_down_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.3/docs/resources/elastigroup_azure_v3#scaling_down_policy ElastigroupAzureV3#scaling_down_policy}
	ScalingDownPolicy interface{} `field:"optional" json:"scalingDownPolicy" yaml:"scalingDownPolicy"`
	// scaling_up_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.3/docs/resources/elastigroup_azure_v3#scaling_up_policy ElastigroupAzureV3#scaling_up_policy}
	ScalingUpPolicy interface{} `field:"optional" json:"scalingUpPolicy" yaml:"scalingUpPolicy"`
	// scheduling_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.3/docs/resources/elastigroup_azure_v3#scheduling_task ElastigroupAzureV3#scheduling_task}
	SchedulingTask interface{} `field:"optional" json:"schedulingTask" yaml:"schedulingTask"`
	// secret block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.3/docs/resources/elastigroup_azure_v3#secret ElastigroupAzureV3#secret}
	Secret interface{} `field:"optional" json:"secret" yaml:"secret"`
	// security block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.3/docs/resources/elastigroup_azure_v3#security ElastigroupAzureV3#security}
	Security *ElastigroupAzureV3Security `field:"optional" json:"security" yaml:"security"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.3/docs/resources/elastigroup_azure_v3#shutdown_script ElastigroupAzureV3#shutdown_script}.
	ShutdownScript *string `field:"optional" json:"shutdownScript" yaml:"shutdownScript"`
	// signal block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.3/docs/resources/elastigroup_azure_v3#signal ElastigroupAzureV3#signal}
	Signal interface{} `field:"optional" json:"signal" yaml:"signal"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.3/docs/resources/elastigroup_azure_v3#spot_percentage ElastigroupAzureV3#spot_percentage}.
	SpotPercentage *float64 `field:"optional" json:"spotPercentage" yaml:"spotPercentage"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.3/docs/resources/elastigroup_azure_v3#tags ElastigroupAzureV3#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.3/docs/resources/elastigroup_azure_v3#user_data ElastigroupAzureV3#user_data}.
	UserData *string `field:"optional" json:"userData" yaml:"userData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.3/docs/resources/elastigroup_azure_v3#vm_name_prefix ElastigroupAzureV3#vm_name_prefix}.
	VmNamePrefix *string `field:"optional" json:"vmNamePrefix" yaml:"vmNamePrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.3/docs/resources/elastigroup_azure_v3#zones ElastigroupAzureV3#zones}.
	Zones *[]*string `field:"optional" json:"zones" yaml:"zones"`
}

