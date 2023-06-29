package statefulnodeazure


type StatefulNodeAzureBootDiagnostics struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.124.0/docs/resources/stateful_node_azure#is_enabled StatefulNodeAzure#is_enabled}.
	IsEnabled interface{} `field:"optional" json:"isEnabled" yaml:"isEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.124.0/docs/resources/stateful_node_azure#storage_url StatefulNodeAzure#storage_url}.
	StorageUrl *string `field:"optional" json:"storageUrl" yaml:"storageUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.124.0/docs/resources/stateful_node_azure#type StatefulNodeAzure#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

