// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type StatefulNodeAzureSignal struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/stateful_node_azure#timeout StatefulNodeAzure#timeout}.
	Timeout *float64 `field:"required" json:"timeout" yaml:"timeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/stateful_node_azure#type StatefulNodeAzure#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

