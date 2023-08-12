package oceanecs


type OceanEcsBlockDeviceMappingsEbsDynamicVolumeSize struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.133.0/docs/resources/ocean_ecs#base_size OceanEcs#base_size}.
	BaseSize *float64 `field:"required" json:"baseSize" yaml:"baseSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.133.0/docs/resources/ocean_ecs#resource OceanEcs#resource}.
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.133.0/docs/resources/ocean_ecs#size_per_resource_unit OceanEcs#size_per_resource_unit}.
	SizePerResourceUnit *float64 `field:"required" json:"sizePerResourceUnit" yaml:"sizePerResourceUnit"`
}

