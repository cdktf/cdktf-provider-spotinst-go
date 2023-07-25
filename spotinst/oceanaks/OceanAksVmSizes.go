package oceanaks


type OceanAksVmSizes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.128.0/docs/resources/ocean_aks#whitelist OceanAks#whitelist}.
	Whitelist *[]*string `field:"optional" json:"whitelist" yaml:"whitelist"`
}

