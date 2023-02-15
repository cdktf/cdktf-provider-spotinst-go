package oceanaws


type OceanAwsClusterOrientation struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aws#availability_vs_cost OceanAws#availability_vs_cost}.
	AvailabilityVsCost *string `field:"optional" json:"availabilityVsCost" yaml:"availabilityVsCost"`
}

