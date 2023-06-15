package managedinstanceaws

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ManagedInstanceAwsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.2/docs/resources/managed_instance_aws#image_id ManagedInstanceAws#image_id}.
	ImageId *string `field:"required" json:"imageId" yaml:"imageId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.2/docs/resources/managed_instance_aws#instance_types ManagedInstanceAws#instance_types}.
	InstanceTypes *[]*string `field:"required" json:"instanceTypes" yaml:"instanceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.2/docs/resources/managed_instance_aws#name ManagedInstanceAws#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.2/docs/resources/managed_instance_aws#persist_block_devices ManagedInstanceAws#persist_block_devices}.
	PersistBlockDevices interface{} `field:"required" json:"persistBlockDevices" yaml:"persistBlockDevices"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.2/docs/resources/managed_instance_aws#product ManagedInstanceAws#product}.
	Product *string `field:"required" json:"product" yaml:"product"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.2/docs/resources/managed_instance_aws#subnet_ids ManagedInstanceAws#subnet_ids}.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.2/docs/resources/managed_instance_aws#vpc_id ManagedInstanceAws#vpc_id}.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.2/docs/resources/managed_instance_aws#auto_healing ManagedInstanceAws#auto_healing}.
	AutoHealing interface{} `field:"optional" json:"autoHealing" yaml:"autoHealing"`
	// block_device_mappings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.2/docs/resources/managed_instance_aws#block_device_mappings ManagedInstanceAws#block_device_mappings}
	BlockDeviceMappings interface{} `field:"optional" json:"blockDeviceMappings" yaml:"blockDeviceMappings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.2/docs/resources/managed_instance_aws#block_devices_mode ManagedInstanceAws#block_devices_mode}.
	BlockDevicesMode *string `field:"optional" json:"blockDevicesMode" yaml:"blockDevicesMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.2/docs/resources/managed_instance_aws#cpu_credits ManagedInstanceAws#cpu_credits}.
	CpuCredits *string `field:"optional" json:"cpuCredits" yaml:"cpuCredits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.2/docs/resources/managed_instance_aws#description ManagedInstanceAws#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.2/docs/resources/managed_instance_aws#draining_timeout ManagedInstanceAws#draining_timeout}.
	DrainingTimeout *float64 `field:"optional" json:"drainingTimeout" yaml:"drainingTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.2/docs/resources/managed_instance_aws#ebs_optimized ManagedInstanceAws#ebs_optimized}.
	EbsOptimized interface{} `field:"optional" json:"ebsOptimized" yaml:"ebsOptimized"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.2/docs/resources/managed_instance_aws#elastic_ip ManagedInstanceAws#elastic_ip}.
	ElasticIp *string `field:"optional" json:"elasticIp" yaml:"elasticIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.2/docs/resources/managed_instance_aws#enable_monitoring ManagedInstanceAws#enable_monitoring}.
	EnableMonitoring interface{} `field:"optional" json:"enableMonitoring" yaml:"enableMonitoring"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.2/docs/resources/managed_instance_aws#fall_back_to_od ManagedInstanceAws#fall_back_to_od}.
	FallBackToOd interface{} `field:"optional" json:"fallBackToOd" yaml:"fallBackToOd"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.2/docs/resources/managed_instance_aws#grace_period ManagedInstanceAws#grace_period}.
	GracePeriod *float64 `field:"optional" json:"gracePeriod" yaml:"gracePeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.2/docs/resources/managed_instance_aws#health_check_type ManagedInstanceAws#health_check_type}.
	HealthCheckType *string `field:"optional" json:"healthCheckType" yaml:"healthCheckType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.2/docs/resources/managed_instance_aws#iam_instance_profile ManagedInstanceAws#iam_instance_profile}.
	IamInstanceProfile *string `field:"optional" json:"iamInstanceProfile" yaml:"iamInstanceProfile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.2/docs/resources/managed_instance_aws#id ManagedInstanceAws#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// integration_route53 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.2/docs/resources/managed_instance_aws#integration_route53 ManagedInstanceAws#integration_route53}
	IntegrationRoute53 *ManagedInstanceAwsIntegrationRoute53 `field:"optional" json:"integrationRoute53" yaml:"integrationRoute53"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.2/docs/resources/managed_instance_aws#key_pair ManagedInstanceAws#key_pair}.
	KeyPair *string `field:"optional" json:"keyPair" yaml:"keyPair"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.2/docs/resources/managed_instance_aws#life_cycle ManagedInstanceAws#life_cycle}.
	LifeCycle *string `field:"optional" json:"lifeCycle" yaml:"lifeCycle"`
	// load_balancers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.2/docs/resources/managed_instance_aws#load_balancers ManagedInstanceAws#load_balancers}
	LoadBalancers interface{} `field:"optional" json:"loadBalancers" yaml:"loadBalancers"`
	// managed_instance_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.2/docs/resources/managed_instance_aws#managed_instance_action ManagedInstanceAws#managed_instance_action}
	ManagedInstanceAction *ManagedInstanceAwsManagedInstanceAction `field:"optional" json:"managedInstanceAction" yaml:"managedInstanceAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.2/docs/resources/managed_instance_aws#minimum_instance_lifetime ManagedInstanceAws#minimum_instance_lifetime}.
	MinimumInstanceLifetime *float64 `field:"optional" json:"minimumInstanceLifetime" yaml:"minimumInstanceLifetime"`
	// network_interface block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.2/docs/resources/managed_instance_aws#network_interface ManagedInstanceAws#network_interface}
	NetworkInterface interface{} `field:"optional" json:"networkInterface" yaml:"networkInterface"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.2/docs/resources/managed_instance_aws#optimization_windows ManagedInstanceAws#optimization_windows}.
	OptimizationWindows *[]*string `field:"optional" json:"optimizationWindows" yaml:"optimizationWindows"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.2/docs/resources/managed_instance_aws#orientation ManagedInstanceAws#orientation}.
	Orientation *string `field:"optional" json:"orientation" yaml:"orientation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.2/docs/resources/managed_instance_aws#persist_private_ip ManagedInstanceAws#persist_private_ip}.
	PersistPrivateIp interface{} `field:"optional" json:"persistPrivateIp" yaml:"persistPrivateIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.2/docs/resources/managed_instance_aws#persist_root_device ManagedInstanceAws#persist_root_device}.
	PersistRootDevice interface{} `field:"optional" json:"persistRootDevice" yaml:"persistRootDevice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.2/docs/resources/managed_instance_aws#placement_tenancy ManagedInstanceAws#placement_tenancy}.
	PlacementTenancy *string `field:"optional" json:"placementTenancy" yaml:"placementTenancy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.2/docs/resources/managed_instance_aws#preferred_type ManagedInstanceAws#preferred_type}.
	PreferredType *string `field:"optional" json:"preferredType" yaml:"preferredType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.2/docs/resources/managed_instance_aws#private_ip ManagedInstanceAws#private_ip}.
	PrivateIp *string `field:"optional" json:"privateIp" yaml:"privateIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.2/docs/resources/managed_instance_aws#region ManagedInstanceAws#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// resource_tag_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.2/docs/resources/managed_instance_aws#resource_tag_specification ManagedInstanceAws#resource_tag_specification}
	ResourceTagSpecification interface{} `field:"optional" json:"resourceTagSpecification" yaml:"resourceTagSpecification"`
	// revert_to_spot block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.2/docs/resources/managed_instance_aws#revert_to_spot ManagedInstanceAws#revert_to_spot}
	RevertToSpot *ManagedInstanceAwsRevertToSpot `field:"optional" json:"revertToSpot" yaml:"revertToSpot"`
	// scheduled_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.2/docs/resources/managed_instance_aws#scheduled_task ManagedInstanceAws#scheduled_task}
	ScheduledTask interface{} `field:"optional" json:"scheduledTask" yaml:"scheduledTask"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.2/docs/resources/managed_instance_aws#security_group_ids ManagedInstanceAws#security_group_ids}.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.2/docs/resources/managed_instance_aws#shutdown_script ManagedInstanceAws#shutdown_script}.
	ShutdownScript *string `field:"optional" json:"shutdownScript" yaml:"shutdownScript"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.2/docs/resources/managed_instance_aws#tags ManagedInstanceAws#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.2/docs/resources/managed_instance_aws#unhealthy_duration ManagedInstanceAws#unhealthy_duration}.
	UnhealthyDuration *float64 `field:"optional" json:"unhealthyDuration" yaml:"unhealthyDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.2/docs/resources/managed_instance_aws#user_data ManagedInstanceAws#user_data}.
	UserData *string `field:"optional" json:"userData" yaml:"userData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.2/docs/resources/managed_instance_aws#utilize_reserved_instances ManagedInstanceAws#utilize_reserved_instances}.
	UtilizeReservedInstances interface{} `field:"optional" json:"utilizeReservedInstances" yaml:"utilizeReservedInstances"`
}

