// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type StatefulNodeAzureStrategyRevertToSpot struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/stateful_node_azure#perform_at StatefulNodeAzure#perform_at}.
	PerformAt *string `field:"required" json:"performAt" yaml:"performAt"`
}

