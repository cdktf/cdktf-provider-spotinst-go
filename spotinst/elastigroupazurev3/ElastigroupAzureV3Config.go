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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.143.0/docs/resources/elastigroup_azure_v3#fallback_to_on_demand ElastigroupAzureV3#fallback_to_on_demand}.
	FallbackToOnDemand interface{} `field:"required" json:"fallbackToOnDemand" yaml:"fallbackToOnDemand"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.143.0/docs/resources/elastigroup_azure_v3#name ElastigroupAzureV3#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// network block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.143.0/docs/resources/elastigroup_azure_v3#network ElastigroupAzureV3#network}
	Network *ElastigroupAzureV3Network `field:"required" json:"network" yaml:"network"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.143.0/docs/resources/elastigroup_azure_v3#od_sizes ElastigroupAzureV3#od_sizes}.
	OdSizes *[]*string `field:"required" json:"odSizes" yaml:"odSizes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.143.0/docs/resources/elastigroup_azure_v3#os ElastigroupAzureV3#os}.
	Os *string `field:"required" json:"os" yaml:"os"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.143.0/docs/resources/elastigroup_azure_v3#region ElastigroupAzureV3#region}.
	Region *string `field:"required" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.143.0/docs/resources/elastigroup_azure_v3#resource_group_name ElastigroupAzureV3#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.143.0/docs/resources/elastigroup_azure_v3#spot_sizes ElastigroupAzureV3#spot_sizes}.
	SpotSizes *[]*string `field:"required" json:"spotSizes" yaml:"spotSizes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.143.0/docs/resources/elastigroup_azure_v3#custom_data ElastigroupAzureV3#custom_data}.
	CustomData *string `field:"optional" json:"customData" yaml:"customData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.143.0/docs/resources/elastigroup_azure_v3#desired_capacity ElastigroupAzureV3#desired_capacity}.
	DesiredCapacity *float64 `field:"optional" json:"desiredCapacity" yaml:"desiredCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.143.0/docs/resources/elastigroup_azure_v3#draining_timeout ElastigroupAzureV3#draining_timeout}.
	DrainingTimeout *float64 `field:"optional" json:"drainingTimeout" yaml:"drainingTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.143.0/docs/resources/elastigroup_azure_v3#id ElastigroupAzureV3#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// image block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.143.0/docs/resources/elastigroup_azure_v3#image ElastigroupAzureV3#image}
	Image interface{} `field:"optional" json:"image" yaml:"image"`
	// login block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.143.0/docs/resources/elastigroup_azure_v3#login ElastigroupAzureV3#login}
	Login *ElastigroupAzureV3Login `field:"optional" json:"login" yaml:"login"`
	// managed_service_identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.143.0/docs/resources/elastigroup_azure_v3#managed_service_identity ElastigroupAzureV3#managed_service_identity}
	ManagedServiceIdentity interface{} `field:"optional" json:"managedServiceIdentity" yaml:"managedServiceIdentity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.143.0/docs/resources/elastigroup_azure_v3#max_size ElastigroupAzureV3#max_size}.
	MaxSize *float64 `field:"optional" json:"maxSize" yaml:"maxSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.143.0/docs/resources/elastigroup_azure_v3#min_size ElastigroupAzureV3#min_size}.
	MinSize *float64 `field:"optional" json:"minSize" yaml:"minSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.143.0/docs/resources/elastigroup_azure_v3#on_demand_count ElastigroupAzureV3#on_demand_count}.
	OnDemandCount *float64 `field:"optional" json:"onDemandCount" yaml:"onDemandCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.143.0/docs/resources/elastigroup_azure_v3#spot_percentage ElastigroupAzureV3#spot_percentage}.
	SpotPercentage *float64 `field:"optional" json:"spotPercentage" yaml:"spotPercentage"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.143.0/docs/resources/elastigroup_azure_v3#tags ElastigroupAzureV3#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

