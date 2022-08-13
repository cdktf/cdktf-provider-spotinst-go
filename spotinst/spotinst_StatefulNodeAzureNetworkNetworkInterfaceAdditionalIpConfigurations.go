// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type StatefulNodeAzureNetworkNetworkInterfaceAdditionalIpConfigurations struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/stateful_node_azure#name StatefulNodeAzure#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/stateful_node_azure#private_ip_address_version StatefulNodeAzure#private_ip_address_version}.
	PrivateIpAddressVersion *string `field:"required" json:"privateIpAddressVersion" yaml:"privateIpAddressVersion"`
}

