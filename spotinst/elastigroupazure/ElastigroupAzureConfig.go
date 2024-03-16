// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupazure

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ElastigroupAzureConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/elastigroup_azure#low_priority_sizes ElastigroupAzure#low_priority_sizes}.
	LowPrioritySizes *[]*string `field:"required" json:"lowPrioritySizes" yaml:"lowPrioritySizes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/elastigroup_azure#name ElastigroupAzure#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// network block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/elastigroup_azure#network ElastigroupAzure#network}
	Network *ElastigroupAzureNetwork `field:"required" json:"network" yaml:"network"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/elastigroup_azure#od_sizes ElastigroupAzure#od_sizes}.
	OdSizes *[]*string `field:"required" json:"odSizes" yaml:"odSizes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/elastigroup_azure#product ElastigroupAzure#product}.
	Product *string `field:"required" json:"product" yaml:"product"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/elastigroup_azure#region ElastigroupAzure#region}.
	Region *string `field:"required" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/elastigroup_azure#resource_group_name ElastigroupAzure#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// strategy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/elastigroup_azure#strategy ElastigroupAzure#strategy}
	Strategy *ElastigroupAzureStrategy `field:"required" json:"strategy" yaml:"strategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/elastigroup_azure#custom_data ElastigroupAzure#custom_data}.
	CustomData *string `field:"optional" json:"customData" yaml:"customData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/elastigroup_azure#desired_capacity ElastigroupAzure#desired_capacity}.
	DesiredCapacity *float64 `field:"optional" json:"desiredCapacity" yaml:"desiredCapacity"`
	// health_check block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/elastigroup_azure#health_check ElastigroupAzure#health_check}
	HealthCheck *ElastigroupAzureHealthCheck `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/elastigroup_azure#id ElastigroupAzure#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// image block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/elastigroup_azure#image ElastigroupAzure#image}
	Image interface{} `field:"optional" json:"image" yaml:"image"`
	// integration_kubernetes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/elastigroup_azure#integration_kubernetes ElastigroupAzure#integration_kubernetes}
	IntegrationKubernetes *ElastigroupAzureIntegrationKubernetes `field:"optional" json:"integrationKubernetes" yaml:"integrationKubernetes"`
	// integration_multai_runtime block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/elastigroup_azure#integration_multai_runtime ElastigroupAzure#integration_multai_runtime}
	IntegrationMultaiRuntime *ElastigroupAzureIntegrationMultaiRuntime `field:"optional" json:"integrationMultaiRuntime" yaml:"integrationMultaiRuntime"`
	// load_balancers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/elastigroup_azure#load_balancers ElastigroupAzure#load_balancers}
	LoadBalancers interface{} `field:"optional" json:"loadBalancers" yaml:"loadBalancers"`
	// login block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/elastigroup_azure#login ElastigroupAzure#login}
	Login *ElastigroupAzureLogin `field:"optional" json:"login" yaml:"login"`
	// managed_service_identities block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/elastigroup_azure#managed_service_identities ElastigroupAzure#managed_service_identities}
	ManagedServiceIdentities interface{} `field:"optional" json:"managedServiceIdentities" yaml:"managedServiceIdentities"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/elastigroup_azure#max_size ElastigroupAzure#max_size}.
	MaxSize *float64 `field:"optional" json:"maxSize" yaml:"maxSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/elastigroup_azure#min_size ElastigroupAzure#min_size}.
	MinSize *float64 `field:"optional" json:"minSize" yaml:"minSize"`
	// scaling_down_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/elastigroup_azure#scaling_down_policy ElastigroupAzure#scaling_down_policy}
	ScalingDownPolicy interface{} `field:"optional" json:"scalingDownPolicy" yaml:"scalingDownPolicy"`
	// scaling_up_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/elastigroup_azure#scaling_up_policy ElastigroupAzure#scaling_up_policy}
	ScalingUpPolicy interface{} `field:"optional" json:"scalingUpPolicy" yaml:"scalingUpPolicy"`
	// scheduled_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/elastigroup_azure#scheduled_task ElastigroupAzure#scheduled_task}
	ScheduledTask interface{} `field:"optional" json:"scheduledTask" yaml:"scheduledTask"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/elastigroup_azure#shutdown_script ElastigroupAzure#shutdown_script}.
	ShutdownScript *string `field:"optional" json:"shutdownScript" yaml:"shutdownScript"`
	// update_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/elastigroup_azure#update_policy ElastigroupAzure#update_policy}
	UpdatePolicy *ElastigroupAzureUpdatePolicy `field:"optional" json:"updatePolicy" yaml:"updatePolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.164.0/docs/resources/elastigroup_azure#user_data ElastigroupAzure#user_data}.
	UserData *string `field:"optional" json:"userData" yaml:"userData"`
}

