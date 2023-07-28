package statefulnodeazure


type StatefulNodeAzureUpdateState struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.130.0/docs/resources/stateful_node_azure#state StatefulNodeAzure#state}.
	State *string `field:"required" json:"state" yaml:"state"`
}

