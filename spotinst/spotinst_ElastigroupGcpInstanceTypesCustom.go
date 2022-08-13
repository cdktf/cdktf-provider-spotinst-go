// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type ElastigroupGcpInstanceTypesCustom struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_gcp#memory_gib ElastigroupGcp#memory_gib}.
	MemoryGib *float64 `field:"required" json:"memoryGib" yaml:"memoryGib"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_gcp#vcpu ElastigroupGcp#vcpu}.
	Vcpu *float64 `field:"required" json:"vcpu" yaml:"vcpu"`
}

