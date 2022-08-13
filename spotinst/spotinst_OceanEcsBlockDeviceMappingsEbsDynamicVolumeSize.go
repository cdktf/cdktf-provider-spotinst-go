// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type OceanEcsBlockDeviceMappingsEbsDynamicVolumeSize struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_ecs#base_size OceanEcs#base_size}.
	BaseSize *float64 `field:"required" json:"baseSize" yaml:"baseSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_ecs#resource OceanEcs#resource}.
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_ecs#size_per_resource_unit OceanEcs#size_per_resource_unit}.
	SizePerResourceUnit *float64 `field:"required" json:"sizePerResourceUnit" yaml:"sizePerResourceUnit"`
}

