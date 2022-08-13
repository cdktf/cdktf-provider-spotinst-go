// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type OceanEcsAutoscalerResourceLimits struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_ecs#max_memory_gib OceanEcs#max_memory_gib}.
	MaxMemoryGib *float64 `field:"optional" json:"maxMemoryGib" yaml:"maxMemoryGib"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_ecs#max_vcpu OceanEcs#max_vcpu}.
	MaxVcpu *float64 `field:"optional" json:"maxVcpu" yaml:"maxVcpu"`
}

