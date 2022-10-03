package oceanawslaunchspec


type OceanAwsLaunchSpecUpdatePolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aws_launch_spec#should_roll OceanAwsLaunchSpec#should_roll}.
	ShouldRoll interface{} `field:"required" json:"shouldRoll" yaml:"shouldRoll"`
	// roll_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aws_launch_spec#roll_config OceanAwsLaunchSpec#roll_config}
	RollConfig *OceanAwsLaunchSpecUpdatePolicyRollConfig `field:"optional" json:"rollConfig" yaml:"rollConfig"`
}

