// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type OceanAksOsDisk struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aks#size_gb OceanAks#size_gb}.
	SizeGb *float64 `field:"required" json:"sizeGb" yaml:"sizeGb"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aks#type OceanAks#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

