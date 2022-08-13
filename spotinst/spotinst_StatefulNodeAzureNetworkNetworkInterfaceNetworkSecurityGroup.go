// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type StatefulNodeAzureNetworkNetworkInterfaceNetworkSecurityGroup struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/stateful_node_azure#name StatefulNodeAzure#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/stateful_node_azure#network_resource_group_name StatefulNodeAzure#network_resource_group_name}.
	NetworkResourceGroupName *string `field:"optional" json:"networkResourceGroupName" yaml:"networkResourceGroupName"`
}

