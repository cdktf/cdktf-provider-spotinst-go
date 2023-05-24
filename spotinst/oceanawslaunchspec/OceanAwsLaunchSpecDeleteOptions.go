package oceanawslaunchspec


type OceanAwsLaunchSpecDeleteOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.119.0/docs/resources/ocean_aws_launch_spec#force_delete OceanAwsLaunchSpec#force_delete}.
	ForceDelete interface{} `field:"required" json:"forceDelete" yaml:"forceDelete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.119.0/docs/resources/ocean_aws_launch_spec#delete_nodes OceanAwsLaunchSpec#delete_nodes}.
	DeleteNodes interface{} `field:"optional" json:"deleteNodes" yaml:"deleteNodes"`
}

