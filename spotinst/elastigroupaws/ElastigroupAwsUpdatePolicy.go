package elastigroupaws


type ElastigroupAwsUpdatePolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.131.0/docs/resources/elastigroup_aws#should_resume_stateful ElastigroupAws#should_resume_stateful}.
	ShouldResumeStateful interface{} `field:"required" json:"shouldResumeStateful" yaml:"shouldResumeStateful"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.131.0/docs/resources/elastigroup_aws#should_roll ElastigroupAws#should_roll}.
	ShouldRoll interface{} `field:"required" json:"shouldRoll" yaml:"shouldRoll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.131.0/docs/resources/elastigroup_aws#auto_apply_tags ElastigroupAws#auto_apply_tags}.
	AutoApplyTags interface{} `field:"optional" json:"autoApplyTags" yaml:"autoApplyTags"`
	// roll_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.131.0/docs/resources/elastigroup_aws#roll_config ElastigroupAws#roll_config}
	RollConfig *ElastigroupAwsUpdatePolicyRollConfig `field:"optional" json:"rollConfig" yaml:"rollConfig"`
}

