package statefulnodeazure


type StatefulNodeAzureStrategyRevertToSpot struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.116.0/docs/resources/stateful_node_azure#perform_at StatefulNodeAzure#perform_at}.
	PerformAt *string `field:"required" json:"performAt" yaml:"performAt"`
}

