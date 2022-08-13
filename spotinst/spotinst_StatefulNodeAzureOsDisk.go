// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type StatefulNodeAzureOsDisk struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/stateful_node_azure#type StatefulNodeAzure#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/stateful_node_azure#size_gb StatefulNodeAzure#size_gb}.
	SizeGb *float64 `field:"optional" json:"sizeGb" yaml:"sizeGb"`
}

