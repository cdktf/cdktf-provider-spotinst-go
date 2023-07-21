package oceanaws

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OceanAwsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.127.0/docs/resources/ocean_aws#security_groups OceanAws#security_groups}.
	SecurityGroups *[]*string `field:"required" json:"securityGroups" yaml:"securityGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.127.0/docs/resources/ocean_aws#subnet_ids OceanAws#subnet_ids}.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.127.0/docs/resources/ocean_aws#associate_ipv6_address OceanAws#associate_ipv6_address}.
	AssociateIpv6Address interface{} `field:"optional" json:"associateIpv6Address" yaml:"associateIpv6Address"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.127.0/docs/resources/ocean_aws#associate_public_ip_address OceanAws#associate_public_ip_address}.
	AssociatePublicIpAddress interface{} `field:"optional" json:"associatePublicIpAddress" yaml:"associatePublicIpAddress"`
	// autoscaler block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.127.0/docs/resources/ocean_aws#autoscaler OceanAws#autoscaler}
	Autoscaler *OceanAwsAutoscaler `field:"optional" json:"autoscaler" yaml:"autoscaler"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.127.0/docs/resources/ocean_aws#blacklist OceanAws#blacklist}.
	Blacklist *[]*string `field:"optional" json:"blacklist" yaml:"blacklist"`
	// block_device_mappings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.127.0/docs/resources/ocean_aws#block_device_mappings OceanAws#block_device_mappings}
	BlockDeviceMappings interface{} `field:"optional" json:"blockDeviceMappings" yaml:"blockDeviceMappings"`
	// cluster_orientation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.127.0/docs/resources/ocean_aws#cluster_orientation OceanAws#cluster_orientation}
	ClusterOrientation interface{} `field:"optional" json:"clusterOrientation" yaml:"clusterOrientation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.127.0/docs/resources/ocean_aws#controller_id OceanAws#controller_id}.
	ControllerId *string `field:"optional" json:"controllerId" yaml:"controllerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.127.0/docs/resources/ocean_aws#desired_capacity OceanAws#desired_capacity}.
	DesiredCapacity *float64 `field:"optional" json:"desiredCapacity" yaml:"desiredCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.127.0/docs/resources/ocean_aws#draining_timeout OceanAws#draining_timeout}.
	DrainingTimeout *float64 `field:"optional" json:"drainingTimeout" yaml:"drainingTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.127.0/docs/resources/ocean_aws#ebs_optimized OceanAws#ebs_optimized}.
	EbsOptimized interface{} `field:"optional" json:"ebsOptimized" yaml:"ebsOptimized"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.127.0/docs/resources/ocean_aws#fallback_to_ondemand OceanAws#fallback_to_ondemand}.
	FallbackToOndemand interface{} `field:"optional" json:"fallbackToOndemand" yaml:"fallbackToOndemand"`
	// filters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.127.0/docs/resources/ocean_aws#filters OceanAws#filters}
	Filters *OceanAwsFilters `field:"optional" json:"filters" yaml:"filters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.127.0/docs/resources/ocean_aws#grace_period OceanAws#grace_period}.
	GracePeriod *float64 `field:"optional" json:"gracePeriod" yaml:"gracePeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.127.0/docs/resources/ocean_aws#iam_instance_profile OceanAws#iam_instance_profile}.
	IamInstanceProfile *string `field:"optional" json:"iamInstanceProfile" yaml:"iamInstanceProfile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.127.0/docs/resources/ocean_aws#id OceanAws#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.127.0/docs/resources/ocean_aws#image_id OceanAws#image_id}.
	ImageId *string `field:"optional" json:"imageId" yaml:"imageId"`
	// instance_metadata_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.127.0/docs/resources/ocean_aws#instance_metadata_options OceanAws#instance_metadata_options}
	InstanceMetadataOptions *OceanAwsInstanceMetadataOptions `field:"optional" json:"instanceMetadataOptions" yaml:"instanceMetadataOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.127.0/docs/resources/ocean_aws#key_name OceanAws#key_name}.
	KeyName *string `field:"optional" json:"keyName" yaml:"keyName"`
	// load_balancers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.127.0/docs/resources/ocean_aws#load_balancers OceanAws#load_balancers}
	LoadBalancers interface{} `field:"optional" json:"loadBalancers" yaml:"loadBalancers"`
	// logging block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.127.0/docs/resources/ocean_aws#logging OceanAws#logging}
	Logging *OceanAwsLogging `field:"optional" json:"logging" yaml:"logging"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.127.0/docs/resources/ocean_aws#max_size OceanAws#max_size}.
	MaxSize *float64 `field:"optional" json:"maxSize" yaml:"maxSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.127.0/docs/resources/ocean_aws#min_size OceanAws#min_size}.
	MinSize *float64 `field:"optional" json:"minSize" yaml:"minSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.127.0/docs/resources/ocean_aws#monitoring OceanAws#monitoring}.
	Monitoring interface{} `field:"optional" json:"monitoring" yaml:"monitoring"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.127.0/docs/resources/ocean_aws#name OceanAws#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.127.0/docs/resources/ocean_aws#region OceanAws#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// resource_tag_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.127.0/docs/resources/ocean_aws#resource_tag_specification OceanAws#resource_tag_specification}
	ResourceTagSpecification interface{} `field:"optional" json:"resourceTagSpecification" yaml:"resourceTagSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.127.0/docs/resources/ocean_aws#root_volume_size OceanAws#root_volume_size}.
	RootVolumeSize *float64 `field:"optional" json:"rootVolumeSize" yaml:"rootVolumeSize"`
	// scheduled_task block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.127.0/docs/resources/ocean_aws#scheduled_task OceanAws#scheduled_task}
	ScheduledTask interface{} `field:"optional" json:"scheduledTask" yaml:"scheduledTask"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.127.0/docs/resources/ocean_aws#spot_percentage OceanAws#spot_percentage}.
	SpotPercentage *float64 `field:"optional" json:"spotPercentage" yaml:"spotPercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.127.0/docs/resources/ocean_aws#spread_nodes_by OceanAws#spread_nodes_by}.
	SpreadNodesBy *string `field:"optional" json:"spreadNodesBy" yaml:"spreadNodesBy"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.127.0/docs/resources/ocean_aws#tags OceanAws#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// update_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.127.0/docs/resources/ocean_aws#update_policy OceanAws#update_policy}
	UpdatePolicy *OceanAwsUpdatePolicy `field:"optional" json:"updatePolicy" yaml:"updatePolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.127.0/docs/resources/ocean_aws#use_as_template_only OceanAws#use_as_template_only}.
	UseAsTemplateOnly interface{} `field:"optional" json:"useAsTemplateOnly" yaml:"useAsTemplateOnly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.127.0/docs/resources/ocean_aws#user_data OceanAws#user_data}.
	UserData *string `field:"optional" json:"userData" yaml:"userData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.127.0/docs/resources/ocean_aws#utilize_commitments OceanAws#utilize_commitments}.
	UtilizeCommitments interface{} `field:"optional" json:"utilizeCommitments" yaml:"utilizeCommitments"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.127.0/docs/resources/ocean_aws#utilize_reserved_instances OceanAws#utilize_reserved_instances}.
	UtilizeReservedInstances interface{} `field:"optional" json:"utilizeReservedInstances" yaml:"utilizeReservedInstances"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.127.0/docs/resources/ocean_aws#whitelist OceanAws#whitelist}.
	Whitelist *[]*string `field:"optional" json:"whitelist" yaml:"whitelist"`
}

