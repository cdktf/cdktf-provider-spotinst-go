package oceanawslaunchspec


type OceanAwsLaunchSpecElasticIpPoolTagSelector struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aws_launch_spec#tag_key OceanAwsLaunchSpec#tag_key}.
	TagKey *string `field:"required" json:"tagKey" yaml:"tagKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aws_launch_spec#tag_value OceanAwsLaunchSpec#tag_value}.
	TagValue *string `field:"optional" json:"tagValue" yaml:"tagValue"`
}

