// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type ElastigroupGkeInstanceTypesCustom struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_gke#memory_gib ElastigroupGke#memory_gib}.
	MemoryGib *float64 `field:"required" json:"memoryGib" yaml:"memoryGib"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_gke#vcpu ElastigroupGke#vcpu}.
	Vcpu *float64 `field:"required" json:"vcpu" yaml:"vcpu"`
}

