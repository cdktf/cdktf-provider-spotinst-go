package oceanaks


type OceanAksOsDisk struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.112.0/docs/resources/ocean_aks#size_gb OceanAks#size_gb}.
	SizeGb *float64 `field:"required" json:"sizeGb" yaml:"sizeGb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.112.0/docs/resources/ocean_aks#type OceanAks#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

