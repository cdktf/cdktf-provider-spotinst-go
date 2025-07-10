// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupgke

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ElastigroupGkeConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/elastigroup_gke#cluster_zone_name ElastigroupGke#cluster_zone_name}.
	ClusterZoneName *string `field:"required" json:"clusterZoneName" yaml:"clusterZoneName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/elastigroup_gke#desired_capacity ElastigroupGke#desired_capacity}.
	DesiredCapacity *float64 `field:"required" json:"desiredCapacity" yaml:"desiredCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/elastigroup_gke#name ElastigroupGke#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// backend_services block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/elastigroup_gke#backend_services ElastigroupGke#backend_services}
	BackendServices interface{} `field:"optional" json:"backendServices" yaml:"backendServices"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/elastigroup_gke#cluster_id ElastigroupGke#cluster_id}.
	ClusterId *string `field:"optional" json:"clusterId" yaml:"clusterId"`
	// disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/elastigroup_gke#disk ElastigroupGke#disk}
	Disk interface{} `field:"optional" json:"disk" yaml:"disk"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/elastigroup_gke#draining_timeout ElastigroupGke#draining_timeout}.
	DrainingTimeout *float64 `field:"optional" json:"drainingTimeout" yaml:"drainingTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/elastigroup_gke#fallback_to_ondemand ElastigroupGke#fallback_to_ondemand}.
	FallbackToOndemand interface{} `field:"optional" json:"fallbackToOndemand" yaml:"fallbackToOndemand"`
	// gpu block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/elastigroup_gke#gpu ElastigroupGke#gpu}
	Gpu interface{} `field:"optional" json:"gpu" yaml:"gpu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/elastigroup_gke#id ElastigroupGke#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/elastigroup_gke#instance_name_prefix ElastigroupGke#instance_name_prefix}.
	InstanceNamePrefix *string `field:"optional" json:"instanceNamePrefix" yaml:"instanceNamePrefix"`
	// instance_types_custom block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/elastigroup_gke#instance_types_custom ElastigroupGke#instance_types_custom}
	InstanceTypesCustom interface{} `field:"optional" json:"instanceTypesCustom" yaml:"instanceTypesCustom"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/elastigroup_gke#instance_types_ondemand ElastigroupGke#instance_types_ondemand}.
	InstanceTypesOndemand *string `field:"optional" json:"instanceTypesOndemand" yaml:"instanceTypesOndemand"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/elastigroup_gke#instance_types_preemptible ElastigroupGke#instance_types_preemptible}.
	InstanceTypesPreemptible *[]*string `field:"optional" json:"instanceTypesPreemptible" yaml:"instanceTypesPreemptible"`
	// integration_docker_swarm block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/elastigroup_gke#integration_docker_swarm ElastigroupGke#integration_docker_swarm}
	IntegrationDockerSwarm *ElastigroupGkeIntegrationDockerSwarm `field:"optional" json:"integrationDockerSwarm" yaml:"integrationDockerSwarm"`
	// integration_gke block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/elastigroup_gke#integration_gke ElastigroupGke#integration_gke}
	IntegrationGke *ElastigroupGkeIntegrationGke `field:"optional" json:"integrationGke" yaml:"integrationGke"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/elastigroup_gke#ip_forwarding ElastigroupGke#ip_forwarding}.
	IpForwarding interface{} `field:"optional" json:"ipForwarding" yaml:"ipForwarding"`
	// labels block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/elastigroup_gke#labels ElastigroupGke#labels}
	Labels interface{} `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/elastigroup_gke#max_size ElastigroupGke#max_size}.
	MaxSize *float64 `field:"optional" json:"maxSize" yaml:"maxSize"`
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/elastigroup_gke#metadata ElastigroupGke#metadata}
	Metadata interface{} `field:"optional" json:"metadata" yaml:"metadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/elastigroup_gke#min_cpu_platform ElastigroupGke#min_cpu_platform}.
	MinCpuPlatform *string `field:"optional" json:"minCpuPlatform" yaml:"minCpuPlatform"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/elastigroup_gke#min_size ElastigroupGke#min_size}.
	MinSize *float64 `field:"optional" json:"minSize" yaml:"minSize"`
	// network_interface block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/elastigroup_gke#network_interface ElastigroupGke#network_interface}
	NetworkInterface interface{} `field:"optional" json:"networkInterface" yaml:"networkInterface"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/elastigroup_gke#node_image ElastigroupGke#node_image}.
	NodeImage *string `field:"optional" json:"nodeImage" yaml:"nodeImage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/elastigroup_gke#ondemand_count ElastigroupGke#ondemand_count}.
	OndemandCount *float64 `field:"optional" json:"ondemandCount" yaml:"ondemandCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/elastigroup_gke#optimization_windows ElastigroupGke#optimization_windows}.
	OptimizationWindows *[]*string `field:"optional" json:"optimizationWindows" yaml:"optimizationWindows"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/elastigroup_gke#preemptible_percentage ElastigroupGke#preemptible_percentage}.
	PreemptiblePercentage *float64 `field:"optional" json:"preemptiblePercentage" yaml:"preemptiblePercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/elastigroup_gke#provisioning_model ElastigroupGke#provisioning_model}.
	ProvisioningModel *string `field:"optional" json:"provisioningModel" yaml:"provisioningModel"`
	// revert_to_preemptible block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/elastigroup_gke#revert_to_preemptible ElastigroupGke#revert_to_preemptible}
	RevertToPreemptible interface{} `field:"optional" json:"revertToPreemptible" yaml:"revertToPreemptible"`
	// scaling_down_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/elastigroup_gke#scaling_down_policy ElastigroupGke#scaling_down_policy}
	ScalingDownPolicy interface{} `field:"optional" json:"scalingDownPolicy" yaml:"scalingDownPolicy"`
	// scaling_up_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/elastigroup_gke#scaling_up_policy ElastigroupGke#scaling_up_policy}
	ScalingUpPolicy interface{} `field:"optional" json:"scalingUpPolicy" yaml:"scalingUpPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/elastigroup_gke#service_account ElastigroupGke#service_account}.
	ServiceAccount *string `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// shielded_instance_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/elastigroup_gke#shielded_instance_config ElastigroupGke#shielded_instance_config}
	ShieldedInstanceConfig *ElastigroupGkeShieldedInstanceConfig `field:"optional" json:"shieldedInstanceConfig" yaml:"shieldedInstanceConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/elastigroup_gke#should_utilize_commitments ElastigroupGke#should_utilize_commitments}.
	ShouldUtilizeCommitments interface{} `field:"optional" json:"shouldUtilizeCommitments" yaml:"shouldUtilizeCommitments"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/elastigroup_gke#shutdown_script ElastigroupGke#shutdown_script}.
	ShutdownScript *string `field:"optional" json:"shutdownScript" yaml:"shutdownScript"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/elastigroup_gke#startup_script ElastigroupGke#startup_script}.
	StartupScript *string `field:"optional" json:"startupScript" yaml:"startupScript"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.223.0/docs/resources/elastigroup_gke#tags ElastigroupGke#tags}.
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

