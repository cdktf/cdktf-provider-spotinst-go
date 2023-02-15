package oceanawslaunchspec


type OceanAwsLaunchSpecTaints struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aws_launch_spec#effect OceanAwsLaunchSpec#effect}.
	Effect *string `field:"required" json:"effect" yaml:"effect"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aws_launch_spec#key OceanAwsLaunchSpec#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aws_launch_spec#value OceanAwsLaunchSpec#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

