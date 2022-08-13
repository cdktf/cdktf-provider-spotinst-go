// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type OceanEcsAutoscalerHeadroom struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_ecs#cpu_per_unit OceanEcs#cpu_per_unit}.
	CpuPerUnit *float64 `field:"optional" json:"cpuPerUnit" yaml:"cpuPerUnit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_ecs#memory_per_unit OceanEcs#memory_per_unit}.
	MemoryPerUnit *float64 `field:"optional" json:"memoryPerUnit" yaml:"memoryPerUnit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_ecs#num_of_units OceanEcs#num_of_units}.
	NumOfUnits *float64 `field:"optional" json:"numOfUnits" yaml:"numOfUnits"`
}

