package oceanaws


type OceanAwsBlockDeviceMappingsEbsDynamicVolumeSize struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.133.0/docs/resources/ocean_aws#base_size OceanAws#base_size}.
	BaseSize *float64 `field:"required" json:"baseSize" yaml:"baseSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.133.0/docs/resources/ocean_aws#resource OceanAws#resource}.
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.133.0/docs/resources/ocean_aws#size_per_resource_unit OceanAws#size_per_resource_unit}.
	SizePerResourceUnit *float64 `field:"required" json:"sizePerResourceUnit" yaml:"sizePerResourceUnit"`
}

