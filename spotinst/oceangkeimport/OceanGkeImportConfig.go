// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceangkeimport

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OceanGkeImportConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.227.0/docs/resources/ocean_gke_import#cluster_name OceanGkeImport#cluster_name}.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.227.0/docs/resources/ocean_gke_import#location OceanGkeImport#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// autoscaler block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.227.0/docs/resources/ocean_gke_import#autoscaler OceanGkeImport#autoscaler}
	Autoscaler *OceanGkeImportAutoscaler `field:"optional" json:"autoscaler" yaml:"autoscaler"`
	// auto_update block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.227.0/docs/resources/ocean_gke_import#auto_update OceanGkeImport#auto_update}
	AutoUpdate interface{} `field:"optional" json:"autoUpdate" yaml:"autoUpdate"`
	// backend_services block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.227.0/docs/resources/ocean_gke_import#backend_services OceanGkeImport#backend_services}
	BackendServices interface{} `field:"optional" json:"backendServices" yaml:"backendServices"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.227.0/docs/resources/ocean_gke_import#blacklist OceanGkeImport#blacklist}.
	Blacklist *[]*string `field:"optional" json:"blacklist" yaml:"blacklist"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.227.0/docs/resources/ocean_gke_import#controller_cluster_id OceanGkeImport#controller_cluster_id}.
	ControllerClusterId *string `field:"optional" json:"controllerClusterId" yaml:"controllerClusterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.227.0/docs/resources/ocean_gke_import#desired_capacity OceanGkeImport#desired_capacity}.
	DesiredCapacity *float64 `field:"optional" json:"desiredCapacity" yaml:"desiredCapacity"`
	// filters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.227.0/docs/resources/ocean_gke_import#filters OceanGkeImport#filters}
	Filters *OceanGkeImportFilters `field:"optional" json:"filters" yaml:"filters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.227.0/docs/resources/ocean_gke_import#id OceanGkeImport#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.227.0/docs/resources/ocean_gke_import#max_size OceanGkeImport#max_size}.
	MaxSize *float64 `field:"optional" json:"maxSize" yaml:"maxSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.227.0/docs/resources/ocean_gke_import#min_size OceanGkeImport#min_size}.
	MinSize *float64 `field:"optional" json:"minSize" yaml:"minSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.227.0/docs/resources/ocean_gke_import#root_volume_type OceanGkeImport#root_volume_type}.
	RootVolumeType *string `field:"optional" json:"rootVolumeType" yaml:"rootVolumeType"`
	// scheduled_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.227.0/docs/resources/ocean_gke_import#scheduled_task OceanGkeImport#scheduled_task}
	ScheduledTask interface{} `field:"optional" json:"scheduledTask" yaml:"scheduledTask"`
	// shielded_instance_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.227.0/docs/resources/ocean_gke_import#shielded_instance_config OceanGkeImport#shielded_instance_config}
	ShieldedInstanceConfig *OceanGkeImportShieldedInstanceConfig `field:"optional" json:"shieldedInstanceConfig" yaml:"shieldedInstanceConfig"`
	// strategy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.227.0/docs/resources/ocean_gke_import#strategy OceanGkeImport#strategy}
	Strategy interface{} `field:"optional" json:"strategy" yaml:"strategy"`
	// update_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.227.0/docs/resources/ocean_gke_import#update_policy OceanGkeImport#update_policy}
	UpdatePolicy *OceanGkeImportUpdatePolicy `field:"optional" json:"updatePolicy" yaml:"updatePolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.227.0/docs/resources/ocean_gke_import#use_as_template_only OceanGkeImport#use_as_template_only}.
	UseAsTemplateOnly interface{} `field:"optional" json:"useAsTemplateOnly" yaml:"useAsTemplateOnly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.227.0/docs/resources/ocean_gke_import#whitelist OceanGkeImport#whitelist}.
	Whitelist *[]*string `field:"optional" json:"whitelist" yaml:"whitelist"`
}

