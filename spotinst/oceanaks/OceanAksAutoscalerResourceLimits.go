package oceanaks


type OceanAksAutoscalerResourceLimits struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.115.0/docs/resources/ocean_aks#max_memory_gib OceanAks#max_memory_gib}.
	MaxMemoryGib *float64 `field:"optional" json:"maxMemoryGib" yaml:"maxMemoryGib"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.115.0/docs/resources/ocean_aks#max_vcpu OceanAks#max_vcpu}.
	MaxVcpu *float64 `field:"optional" json:"maxVcpu" yaml:"maxVcpu"`
}

