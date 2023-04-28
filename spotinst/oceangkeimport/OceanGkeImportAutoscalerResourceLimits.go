package oceangkeimport


type OceanGkeImportAutoscalerResourceLimits struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.115.0/docs/resources/ocean_gke_import#max_memory_gib OceanGkeImport#max_memory_gib}.
	MaxMemoryGib *float64 `field:"optional" json:"maxMemoryGib" yaml:"maxMemoryGib"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.115.0/docs/resources/ocean_gke_import#max_vcpu OceanGkeImport#max_vcpu}.
	MaxVcpu *float64 `field:"optional" json:"maxVcpu" yaml:"maxVcpu"`
}

