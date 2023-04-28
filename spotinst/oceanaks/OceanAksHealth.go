package oceanaks


type OceanAksHealth struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.115.0/docs/resources/ocean_aks#grace_period OceanAks#grace_period}.
	GracePeriod *float64 `field:"optional" json:"gracePeriod" yaml:"gracePeriod"`
}

