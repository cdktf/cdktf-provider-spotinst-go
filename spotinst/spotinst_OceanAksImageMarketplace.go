// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type OceanAksImageMarketplace struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aks#offer OceanAks#offer}.
	Offer *string `field:"optional" json:"offer" yaml:"offer"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aks#publisher OceanAks#publisher}.
	Publisher *string `field:"optional" json:"publisher" yaml:"publisher"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aks#sku OceanAks#sku}.
	Sku *string `field:"optional" json:"sku" yaml:"sku"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aks#version OceanAks#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

