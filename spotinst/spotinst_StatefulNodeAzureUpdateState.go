// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type StatefulNodeAzureUpdateState struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/stateful_node_azure#state StatefulNodeAzure#state}.
	State *string `field:"required" json:"state" yaml:"state"`
}

