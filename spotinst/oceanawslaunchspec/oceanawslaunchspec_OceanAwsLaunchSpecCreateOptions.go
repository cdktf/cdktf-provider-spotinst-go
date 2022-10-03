package oceanawslaunchspec


type OceanAwsLaunchSpecCreateOptions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aws_launch_spec#initial_nodes OceanAwsLaunchSpec#initial_nodes}.
	InitialNodes *float64 `field:"required" json:"initialNodes" yaml:"initialNodes"`
}

