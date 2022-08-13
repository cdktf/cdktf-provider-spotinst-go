// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type StatefulNodeAzureSecretSourceVault struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/stateful_node_azure#name StatefulNodeAzure#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/stateful_node_azure#resource_group_name StatefulNodeAzure#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
}

