package elastigroupazure


type ElastigroupAzureImageMarketplace struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_azure#offer ElastigroupAzure#offer}.
	Offer *string `field:"required" json:"offer" yaml:"offer"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_azure#publisher ElastigroupAzure#publisher}.
	Publisher *string `field:"required" json:"publisher" yaml:"publisher"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_azure#sku ElastigroupAzure#sku}.
	Sku *string `field:"required" json:"sku" yaml:"sku"`
}
