// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type ElastigroupAwsBeanstalkManagedActionsPlatformUpdate struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_aws_beanstalk#perform_at ElastigroupAwsBeanstalk#perform_at}.
	PerformAt *string `field:"optional" json:"performAt" yaml:"performAt"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_aws_beanstalk#time_window ElastigroupAwsBeanstalk#time_window}.
	TimeWindow *string `field:"optional" json:"timeWindow" yaml:"timeWindow"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_aws_beanstalk#update_level ElastigroupAwsBeanstalk#update_level}.
	UpdateLevel *string `field:"optional" json:"updateLevel" yaml:"updateLevel"`
}
