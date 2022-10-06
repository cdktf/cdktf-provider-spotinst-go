package statefulnodeazure


type StatefulNodeAzureImageCustomImage struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/stateful_node_azure#custom_image_resource_group_name StatefulNodeAzure#custom_image_resource_group_name}.
	CustomImageResourceGroupName *string `field:"required" json:"customImageResourceGroupName" yaml:"customImageResourceGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/stateful_node_azure#name StatefulNodeAzure#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}
