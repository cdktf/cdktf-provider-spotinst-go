// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupgcp

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ElastigroupGcpConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.159.0/docs/resources/elastigroup_gcp#desired_capacity ElastigroupGcp#desired_capacity}.
	DesiredCapacity *float64 `field:"required" json:"desiredCapacity" yaml:"desiredCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.159.0/docs/resources/elastigroup_gcp#name ElastigroupGcp#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.159.0/docs/resources/elastigroup_gcp#auto_healing ElastigroupGcp#auto_healing}.
	AutoHealing interface{} `field:"optional" json:"autoHealing" yaml:"autoHealing"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.159.0/docs/resources/elastigroup_gcp#availability_zones ElastigroupGcp#availability_zones}.
	AvailabilityZones *[]*string `field:"optional" json:"availabilityZones" yaml:"availabilityZones"`
	// backend_services block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.159.0/docs/resources/elastigroup_gcp#backend_services ElastigroupGcp#backend_services}
	BackendServices interface{} `field:"optional" json:"backendServices" yaml:"backendServices"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.159.0/docs/resources/elastigroup_gcp#description ElastigroupGcp#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.159.0/docs/resources/elastigroup_gcp#disk ElastigroupGcp#disk}
	Disk interface{} `field:"optional" json:"disk" yaml:"disk"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.159.0/docs/resources/elastigroup_gcp#draining_timeout ElastigroupGcp#draining_timeout}.
	DrainingTimeout *float64 `field:"optional" json:"drainingTimeout" yaml:"drainingTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.159.0/docs/resources/elastigroup_gcp#fallback_to_ondemand ElastigroupGcp#fallback_to_ondemand}.
	FallbackToOndemand interface{} `field:"optional" json:"fallbackToOndemand" yaml:"fallbackToOndemand"`
	// gpu block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.159.0/docs/resources/elastigroup_gcp#gpu ElastigroupGcp#gpu}
	Gpu interface{} `field:"optional" json:"gpu" yaml:"gpu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.159.0/docs/resources/elastigroup_gcp#health_check_grace_period ElastigroupGcp#health_check_grace_period}.
	HealthCheckGracePeriod *float64 `field:"optional" json:"healthCheckGracePeriod" yaml:"healthCheckGracePeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.159.0/docs/resources/elastigroup_gcp#health_check_type ElastigroupGcp#health_check_type}.
	HealthCheckType *string `field:"optional" json:"healthCheckType" yaml:"healthCheckType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.159.0/docs/resources/elastigroup_gcp#id ElastigroupGcp#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.159.0/docs/resources/elastigroup_gcp#instance_name_prefix ElastigroupGcp#instance_name_prefix}.
	InstanceNamePrefix *string `field:"optional" json:"instanceNamePrefix" yaml:"instanceNamePrefix"`
	// instance_types_custom block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.159.0/docs/resources/elastigroup_gcp#instance_types_custom ElastigroupGcp#instance_types_custom}
	InstanceTypesCustom interface{} `field:"optional" json:"instanceTypesCustom" yaml:"instanceTypesCustom"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.159.0/docs/resources/elastigroup_gcp#instance_types_ondemand ElastigroupGcp#instance_types_ondemand}.
	InstanceTypesOndemand *string `field:"optional" json:"instanceTypesOndemand" yaml:"instanceTypesOndemand"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.159.0/docs/resources/elastigroup_gcp#instance_types_preemptible ElastigroupGcp#instance_types_preemptible}.
	InstanceTypesPreemptible *[]*string `field:"optional" json:"instanceTypesPreemptible" yaml:"instanceTypesPreemptible"`
	// integration_docker_swarm block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.159.0/docs/resources/elastigroup_gcp#integration_docker_swarm ElastigroupGcp#integration_docker_swarm}
	IntegrationDockerSwarm *ElastigroupGcpIntegrationDockerSwarm `field:"optional" json:"integrationDockerSwarm" yaml:"integrationDockerSwarm"`
	// integration_gke block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.159.0/docs/resources/elastigroup_gcp#integration_gke ElastigroupGcp#integration_gke}
	IntegrationGke *ElastigroupGcpIntegrationGke `field:"optional" json:"integrationGke" yaml:"integrationGke"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.159.0/docs/resources/elastigroup_gcp#ip_forwarding ElastigroupGcp#ip_forwarding}.
	IpForwarding interface{} `field:"optional" json:"ipForwarding" yaml:"ipForwarding"`
	// labels block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.159.0/docs/resources/elastigroup_gcp#labels ElastigroupGcp#labels}
	Labels interface{} `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.159.0/docs/resources/elastigroup_gcp#max_size ElastigroupGcp#max_size}.
	MaxSize *float64 `field:"optional" json:"maxSize" yaml:"maxSize"`
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.159.0/docs/resources/elastigroup_gcp#metadata ElastigroupGcp#metadata}
	Metadata interface{} `field:"optional" json:"metadata" yaml:"metadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.159.0/docs/resources/elastigroup_gcp#min_size ElastigroupGcp#min_size}.
	MinSize *float64 `field:"optional" json:"minSize" yaml:"minSize"`
	// network_interface block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.159.0/docs/resources/elastigroup_gcp#network_interface ElastigroupGcp#network_interface}
	NetworkInterface interface{} `field:"optional" json:"networkInterface" yaml:"networkInterface"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.159.0/docs/resources/elastigroup_gcp#ondemand_count ElastigroupGcp#ondemand_count}.
	OndemandCount *float64 `field:"optional" json:"ondemandCount" yaml:"ondemandCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.159.0/docs/resources/elastigroup_gcp#preemptible_percentage ElastigroupGcp#preemptible_percentage}.
	PreemptiblePercentage *float64 `field:"optional" json:"preemptiblePercentage" yaml:"preemptiblePercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.159.0/docs/resources/elastigroup_gcp#provisioning_model ElastigroupGcp#provisioning_model}.
	ProvisioningModel *string `field:"optional" json:"provisioningModel" yaml:"provisioningModel"`
	// scaling_down_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.159.0/docs/resources/elastigroup_gcp#scaling_down_policy ElastigroupGcp#scaling_down_policy}
	ScalingDownPolicy interface{} `field:"optional" json:"scalingDownPolicy" yaml:"scalingDownPolicy"`
	// scaling_up_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.159.0/docs/resources/elastigroup_gcp#scaling_up_policy ElastigroupGcp#scaling_up_policy}
	ScalingUpPolicy interface{} `field:"optional" json:"scalingUpPolicy" yaml:"scalingUpPolicy"`
	// scheduled_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.159.0/docs/resources/elastigroup_gcp#scheduled_task ElastigroupGcp#scheduled_task}
	ScheduledTask interface{} `field:"optional" json:"scheduledTask" yaml:"scheduledTask"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.159.0/docs/resources/elastigroup_gcp#service_account ElastigroupGcp#service_account}.
	ServiceAccount *string `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.159.0/docs/resources/elastigroup_gcp#shutdown_script ElastigroupGcp#shutdown_script}.
	ShutdownScript *string `field:"optional" json:"shutdownScript" yaml:"shutdownScript"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.159.0/docs/resources/elastigroup_gcp#startup_script ElastigroupGcp#startup_script}.
	StartupScript *string `field:"optional" json:"startupScript" yaml:"startupScript"`
	// subnets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.159.0/docs/resources/elastigroup_gcp#subnets ElastigroupGcp#subnets}
	Subnets interface{} `field:"optional" json:"subnets" yaml:"subnets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.159.0/docs/resources/elastigroup_gcp#tags ElastigroupGcp#tags}.
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.159.0/docs/resources/elastigroup_gcp#unhealthy_duration ElastigroupGcp#unhealthy_duration}.
	UnhealthyDuration *float64 `field:"optional" json:"unhealthyDuration" yaml:"unhealthyDuration"`
}

