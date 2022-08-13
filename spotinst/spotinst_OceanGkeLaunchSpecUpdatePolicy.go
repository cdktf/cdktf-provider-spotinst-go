// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type OceanGkeLaunchSpecUpdatePolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_gke_launch_spec#should_roll OceanGkeLaunchSpec#should_roll}.
	ShouldRoll interface{} `field:"required" json:"shouldRoll" yaml:"shouldRoll"`
	// roll_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_gke_launch_spec#roll_config OceanGkeLaunchSpec#roll_config}
	RollConfig *OceanGkeLaunchSpecUpdatePolicyRollConfig `field:"optional" json:"rollConfig" yaml:"rollConfig"`
}

