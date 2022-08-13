// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type OceanAwsLaunchSpecStrategy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aws_launch_spec#spot_percentage OceanAwsLaunchSpec#spot_percentage}.
	SpotPercentage *float64 `field:"optional" json:"spotPercentage" yaml:"spotPercentage"`
}

