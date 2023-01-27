package oceanecs


type OceanEcsClusterOrientation struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_ecs#availability_vs_cost OceanEcs#availability_vs_cost}.
	AvailabilityVsCost *string `field:"optional" json:"availabilityVsCost" yaml:"availabilityVsCost"`
}

