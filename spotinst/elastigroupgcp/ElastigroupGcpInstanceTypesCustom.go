package elastigroupgcp


type ElastigroupGcpInstanceTypesCustom struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.0/docs/resources/elastigroup_gcp#memory_gib ElastigroupGcp#memory_gib}.
	MemoryGib *float64 `field:"required" json:"memoryGib" yaml:"memoryGib"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.122.0/docs/resources/elastigroup_gcp#vcpu ElastigroupGcp#vcpu}.
	Vcpu *float64 `field:"required" json:"vcpu" yaml:"vcpu"`
}

