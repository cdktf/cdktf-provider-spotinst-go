package managedinstanceaws


type ManagedInstanceAwsIntegrationRoute53DomainsRecordSets struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.0/docs/resources/managed_instance_aws#name ManagedInstanceAws#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.0/docs/resources/managed_instance_aws#use_public_dns ManagedInstanceAws#use_public_dns}.
	UsePublicDns interface{} `field:"optional" json:"usePublicDns" yaml:"usePublicDns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.0/docs/resources/managed_instance_aws#use_public_ip ManagedInstanceAws#use_public_ip}.
	UsePublicIp interface{} `field:"optional" json:"usePublicIp" yaml:"usePublicIp"`
}

