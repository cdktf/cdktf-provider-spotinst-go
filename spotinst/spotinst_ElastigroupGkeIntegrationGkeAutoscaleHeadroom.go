// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type ElastigroupGkeIntegrationGkeAutoscaleHeadroom struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_gke#cpu_per_unit ElastigroupGke#cpu_per_unit}.
	CpuPerUnit *float64 `field:"optional" json:"cpuPerUnit" yaml:"cpuPerUnit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_gke#memory_per_unit ElastigroupGke#memory_per_unit}.
	MemoryPerUnit *float64 `field:"optional" json:"memoryPerUnit" yaml:"memoryPerUnit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_gke#num_of_units ElastigroupGke#num_of_units}.
	NumOfUnits *float64 `field:"optional" json:"numOfUnits" yaml:"numOfUnits"`
}

