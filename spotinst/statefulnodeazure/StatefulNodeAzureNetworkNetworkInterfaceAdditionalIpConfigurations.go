package statefulnodeazure


type StatefulNodeAzureNetworkNetworkInterfaceAdditionalIpConfigurations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.121.0/docs/resources/stateful_node_azure#name StatefulNodeAzure#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.121.0/docs/resources/stateful_node_azure#private_ip_address_version StatefulNodeAzure#private_ip_address_version}.
	PrivateIpAddressVersion *string `field:"required" json:"privateIpAddressVersion" yaml:"privateIpAddressVersion"`
}

