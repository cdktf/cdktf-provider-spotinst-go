// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type OceanAwsLaunchSpecSchedulingShutdownHours struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aws_launch_spec#time_windows OceanAwsLaunchSpec#time_windows}.
	TimeWindows *[]*string `field:"required" json:"timeWindows" yaml:"timeWindows"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aws_launch_spec#is_enabled OceanAwsLaunchSpec#is_enabled}.
	IsEnabled interface{} `field:"optional" json:"isEnabled" yaml:"isEnabled"`
}

