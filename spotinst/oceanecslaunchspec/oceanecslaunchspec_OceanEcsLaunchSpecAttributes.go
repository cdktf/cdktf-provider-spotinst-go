package oceanecslaunchspec


type OceanEcsLaunchSpecAttributes struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_ecs_launch_spec#key OceanEcsLaunchSpec#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_ecs_launch_spec#value OceanEcsLaunchSpec#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

