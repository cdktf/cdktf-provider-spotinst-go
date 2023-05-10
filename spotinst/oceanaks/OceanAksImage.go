package oceanaks


type OceanAksImage struct {
	// marketplace block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.117.0/docs/resources/ocean_aks#marketplace OceanAks#marketplace}
	Marketplace interface{} `field:"optional" json:"marketplace" yaml:"marketplace"`
}

