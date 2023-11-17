// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupaws


type ElastigroupAwsNetworkInterface struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.150.0/docs/resources/elastigroup_aws#device_index ElastigroupAws#device_index}.
	DeviceIndex *string `field:"required" json:"deviceIndex" yaml:"deviceIndex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.150.0/docs/resources/elastigroup_aws#associate_ipv6_address ElastigroupAws#associate_ipv6_address}.
	AssociateIpv6Address interface{} `field:"optional" json:"associateIpv6Address" yaml:"associateIpv6Address"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.150.0/docs/resources/elastigroup_aws#associate_public_ip_address ElastigroupAws#associate_public_ip_address}.
	AssociatePublicIpAddress interface{} `field:"optional" json:"associatePublicIpAddress" yaml:"associatePublicIpAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.150.0/docs/resources/elastigroup_aws#delete_on_termination ElastigroupAws#delete_on_termination}.
	DeleteOnTermination interface{} `field:"optional" json:"deleteOnTermination" yaml:"deleteOnTermination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.150.0/docs/resources/elastigroup_aws#description ElastigroupAws#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.150.0/docs/resources/elastigroup_aws#network_interface_id ElastigroupAws#network_interface_id}.
	NetworkInterfaceId *string `field:"optional" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.150.0/docs/resources/elastigroup_aws#private_ip_address ElastigroupAws#private_ip_address}.
	PrivateIpAddress *string `field:"optional" json:"privateIpAddress" yaml:"privateIpAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.150.0/docs/resources/elastigroup_aws#secondary_private_ip_address_count ElastigroupAws#secondary_private_ip_address_count}.
	SecondaryPrivateIpAddressCount *string `field:"optional" json:"secondaryPrivateIpAddressCount" yaml:"secondaryPrivateIpAddressCount"`
}

