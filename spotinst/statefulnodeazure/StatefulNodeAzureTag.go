package statefulnodeazure


type StatefulNodeAzureTag struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.113.0/docs/resources/stateful_node_azure#tag_key StatefulNodeAzure#tag_key}.
	TagKey *string `field:"required" json:"tagKey" yaml:"tagKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.113.0/docs/resources/stateful_node_azure#tag_value StatefulNodeAzure#tag_value}.
	TagValue *string `field:"optional" json:"tagValue" yaml:"tagValue"`
}

