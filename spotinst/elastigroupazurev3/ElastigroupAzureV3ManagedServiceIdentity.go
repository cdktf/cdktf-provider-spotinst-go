package elastigroupazurev3


type ElastigroupAzureV3ManagedServiceIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.126.0/docs/resources/elastigroup_azure_v3#name ElastigroupAzureV3#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.126.0/docs/resources/elastigroup_azure_v3#resource_group_name ElastigroupAzureV3#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
}

