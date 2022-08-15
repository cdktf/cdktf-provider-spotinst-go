// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type ManagedInstanceAwsNetworkInterface struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/managed_instance_aws#device_index ManagedInstanceAws#device_index}.
	DeviceIndex *string `field:"required" json:"deviceIndex" yaml:"deviceIndex"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/managed_instance_aws#associate_ipv6_address ManagedInstanceAws#associate_ipv6_address}.
	AssociateIpv6Address interface{} `field:"optional" json:"associateIpv6Address" yaml:"associateIpv6Address"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/managed_instance_aws#associate_public_ip_address ManagedInstanceAws#associate_public_ip_address}.
	AssociatePublicIpAddress interface{} `field:"optional" json:"associatePublicIpAddress" yaml:"associatePublicIpAddress"`
}
