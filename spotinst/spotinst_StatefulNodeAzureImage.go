// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type StatefulNodeAzureImage struct {
	// custom_image block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/stateful_node_azure#custom_image StatefulNodeAzure#custom_image}
	CustomImage interface{} `field:"optional" json:"customImage" yaml:"customImage"`
	// gallery block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/stateful_node_azure#gallery StatefulNodeAzure#gallery}
	Gallery interface{} `field:"optional" json:"gallery" yaml:"gallery"`
	// marketplace_image block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/stateful_node_azure#marketplace_image StatefulNodeAzure#marketplace_image}
	MarketplaceImage interface{} `field:"optional" json:"marketplaceImage" yaml:"marketplaceImage"`
}
