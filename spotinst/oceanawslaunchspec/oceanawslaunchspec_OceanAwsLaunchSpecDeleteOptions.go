package oceanawslaunchspec


type OceanAwsLaunchSpecDeleteOptions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aws_launch_spec#force_delete OceanAwsLaunchSpec#force_delete}.
	ForceDelete interface{} `field:"required" json:"forceDelete" yaml:"forceDelete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aws_launch_spec#delete_nodes OceanAwsLaunchSpec#delete_nodes}.
	DeleteNodes interface{} `field:"optional" json:"deleteNodes" yaml:"deleteNodes"`
}

