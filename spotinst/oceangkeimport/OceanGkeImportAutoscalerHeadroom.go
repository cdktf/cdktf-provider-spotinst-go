package oceangkeimport


type OceanGkeImportAutoscalerHeadroom struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.126.0/docs/resources/ocean_gke_import#cpu_per_unit OceanGkeImport#cpu_per_unit}.
	CpuPerUnit *float64 `field:"optional" json:"cpuPerUnit" yaml:"cpuPerUnit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.126.0/docs/resources/ocean_gke_import#gpu_per_unit OceanGkeImport#gpu_per_unit}.
	GpuPerUnit *float64 `field:"optional" json:"gpuPerUnit" yaml:"gpuPerUnit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.126.0/docs/resources/ocean_gke_import#memory_per_unit OceanGkeImport#memory_per_unit}.
	MemoryPerUnit *float64 `field:"optional" json:"memoryPerUnit" yaml:"memoryPerUnit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.126.0/docs/resources/ocean_gke_import#num_of_units OceanGkeImport#num_of_units}.
	NumOfUnits *float64 `field:"optional" json:"numOfUnits" yaml:"numOfUnits"`
}

