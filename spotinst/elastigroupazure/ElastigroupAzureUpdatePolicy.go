package elastigroupazure


type ElastigroupAzureUpdatePolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_azure#should_roll ElastigroupAzure#should_roll}.
	ShouldRoll interface{} `field:"required" json:"shouldRoll" yaml:"shouldRoll"`
	// roll_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_azure#roll_config ElastigroupAzure#roll_config}
	RollConfig *ElastigroupAzureUpdatePolicyRollConfig `field:"optional" json:"rollConfig" yaml:"rollConfig"`
}

