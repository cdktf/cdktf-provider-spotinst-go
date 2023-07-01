package statefulnodeazure

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StatefulNodeAzureConfig struct {
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
	// image block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.125.0/docs/resources/stateful_node_azure#image StatefulNodeAzure#image}
	Image *StatefulNodeAzureImage `field:"required" json:"image" yaml:"image"`
	// login block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.125.0/docs/resources/stateful_node_azure#login StatefulNodeAzure#login}
	Login *StatefulNodeAzureLogin `field:"required" json:"login" yaml:"login"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.125.0/docs/resources/stateful_node_azure#name StatefulNodeAzure#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// network block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.125.0/docs/resources/stateful_node_azure#network StatefulNodeAzure#network}
	Network *StatefulNodeAzureNetwork `field:"required" json:"network" yaml:"network"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.125.0/docs/resources/stateful_node_azure#od_sizes StatefulNodeAzure#od_sizes}.
	OdSizes *[]*string `field:"required" json:"odSizes" yaml:"odSizes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.125.0/docs/resources/stateful_node_azure#os StatefulNodeAzure#os}.
	Os *string `field:"required" json:"os" yaml:"os"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.125.0/docs/resources/stateful_node_azure#region StatefulNodeAzure#region}.
	Region *string `field:"required" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.125.0/docs/resources/stateful_node_azure#resource_group_name StatefulNodeAzure#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.125.0/docs/resources/stateful_node_azure#should_persist_data_disks StatefulNodeAzure#should_persist_data_disks}.
	ShouldPersistDataDisks interface{} `field:"required" json:"shouldPersistDataDisks" yaml:"shouldPersistDataDisks"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.125.0/docs/resources/stateful_node_azure#should_persist_network StatefulNodeAzure#should_persist_network}.
	ShouldPersistNetwork interface{} `field:"required" json:"shouldPersistNetwork" yaml:"shouldPersistNetwork"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.125.0/docs/resources/stateful_node_azure#should_persist_os_disk StatefulNodeAzure#should_persist_os_disk}.
	ShouldPersistOsDisk interface{} `field:"required" json:"shouldPersistOsDisk" yaml:"shouldPersistOsDisk"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.125.0/docs/resources/stateful_node_azure#spot_sizes StatefulNodeAzure#spot_sizes}.
	SpotSizes *[]*string `field:"required" json:"spotSizes" yaml:"spotSizes"`
	// strategy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.125.0/docs/resources/stateful_node_azure#strategy StatefulNodeAzure#strategy}
	Strategy *StatefulNodeAzureStrategy `field:"required" json:"strategy" yaml:"strategy"`
	// attach_data_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.125.0/docs/resources/stateful_node_azure#attach_data_disk StatefulNodeAzure#attach_data_disk}
	AttachDataDisk interface{} `field:"optional" json:"attachDataDisk" yaml:"attachDataDisk"`
	// boot_diagnostics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.125.0/docs/resources/stateful_node_azure#boot_diagnostics StatefulNodeAzure#boot_diagnostics}
	BootDiagnostics interface{} `field:"optional" json:"bootDiagnostics" yaml:"bootDiagnostics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.125.0/docs/resources/stateful_node_azure#custom_data StatefulNodeAzure#custom_data}.
	CustomData *string `field:"optional" json:"customData" yaml:"customData"`
	// data_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.125.0/docs/resources/stateful_node_azure#data_disk StatefulNodeAzure#data_disk}
	DataDisk interface{} `field:"optional" json:"dataDisk" yaml:"dataDisk"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.125.0/docs/resources/stateful_node_azure#data_disks_persistence_mode StatefulNodeAzure#data_disks_persistence_mode}.
	DataDisksPersistenceMode *string `field:"optional" json:"dataDisksPersistenceMode" yaml:"dataDisksPersistenceMode"`
	// delete block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.125.0/docs/resources/stateful_node_azure#delete StatefulNodeAzure#delete}
	Delete interface{} `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.125.0/docs/resources/stateful_node_azure#description StatefulNodeAzure#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// detach_data_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.125.0/docs/resources/stateful_node_azure#detach_data_disk StatefulNodeAzure#detach_data_disk}
	DetachDataDisk interface{} `field:"optional" json:"detachDataDisk" yaml:"detachDataDisk"`
	// extension block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.125.0/docs/resources/stateful_node_azure#extension StatefulNodeAzure#extension}
	Extension interface{} `field:"optional" json:"extension" yaml:"extension"`
	// health block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.125.0/docs/resources/stateful_node_azure#health StatefulNodeAzure#health}
	Health *StatefulNodeAzureHealth `field:"optional" json:"health" yaml:"health"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.125.0/docs/resources/stateful_node_azure#id StatefulNodeAzure#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// import_vm block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.125.0/docs/resources/stateful_node_azure#import_vm StatefulNodeAzure#import_vm}
	ImportVm interface{} `field:"optional" json:"importVm" yaml:"importVm"`
	// load_balancer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.125.0/docs/resources/stateful_node_azure#load_balancer StatefulNodeAzure#load_balancer}
	LoadBalancer interface{} `field:"optional" json:"loadBalancer" yaml:"loadBalancer"`
	// managed_service_identities block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.125.0/docs/resources/stateful_node_azure#managed_service_identities StatefulNodeAzure#managed_service_identities}
	ManagedServiceIdentities interface{} `field:"optional" json:"managedServiceIdentities" yaml:"managedServiceIdentities"`
	// os_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.125.0/docs/resources/stateful_node_azure#os_disk StatefulNodeAzure#os_disk}
	OsDisk *StatefulNodeAzureOsDisk `field:"optional" json:"osDisk" yaml:"osDisk"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.125.0/docs/resources/stateful_node_azure#os_disk_persistence_mode StatefulNodeAzure#os_disk_persistence_mode}.
	OsDiskPersistenceMode *string `field:"optional" json:"osDiskPersistenceMode" yaml:"osDiskPersistenceMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.125.0/docs/resources/stateful_node_azure#preferred_spot_sizes StatefulNodeAzure#preferred_spot_sizes}.
	PreferredSpotSizes *[]*string `field:"optional" json:"preferredSpotSizes" yaml:"preferredSpotSizes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.125.0/docs/resources/stateful_node_azure#preferred_zones StatefulNodeAzure#preferred_zones}.
	PreferredZones *string `field:"optional" json:"preferredZones" yaml:"preferredZones"`
	// scheduling_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.125.0/docs/resources/stateful_node_azure#scheduling_task StatefulNodeAzure#scheduling_task}
	SchedulingTask interface{} `field:"optional" json:"schedulingTask" yaml:"schedulingTask"`
	// secret block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.125.0/docs/resources/stateful_node_azure#secret StatefulNodeAzure#secret}
	Secret interface{} `field:"optional" json:"secret" yaml:"secret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.125.0/docs/resources/stateful_node_azure#should_persist_vm StatefulNodeAzure#should_persist_vm}.
	ShouldPersistVm interface{} `field:"optional" json:"shouldPersistVm" yaml:"shouldPersistVm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.125.0/docs/resources/stateful_node_azure#shutdown_script StatefulNodeAzure#shutdown_script}.
	ShutdownScript *string `field:"optional" json:"shutdownScript" yaml:"shutdownScript"`
	// signal block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.125.0/docs/resources/stateful_node_azure#signal StatefulNodeAzure#signal}
	Signal interface{} `field:"optional" json:"signal" yaml:"signal"`
	// tag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.125.0/docs/resources/stateful_node_azure#tag StatefulNodeAzure#tag}
	Tag interface{} `field:"optional" json:"tag" yaml:"tag"`
	// update_state block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.125.0/docs/resources/stateful_node_azure#update_state StatefulNodeAzure#update_state}
	UpdateState interface{} `field:"optional" json:"updateState" yaml:"updateState"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.125.0/docs/resources/stateful_node_azure#user_data StatefulNodeAzure#user_data}.
	UserData *string `field:"optional" json:"userData" yaml:"userData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.125.0/docs/resources/stateful_node_azure#zones StatefulNodeAzure#zones}.
	Zones *[]*string `field:"optional" json:"zones" yaml:"zones"`
}

