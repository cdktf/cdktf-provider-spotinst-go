package elastigroupazure


type ElastigroupAzureUpdatePolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.125.0/docs/resources/elastigroup_azure#should_roll ElastigroupAzure#should_roll}.
	ShouldRoll interface{} `field:"required" json:"shouldRoll" yaml:"shouldRoll"`
	// roll_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.125.0/docs/resources/elastigroup_azure#roll_config ElastigroupAzure#roll_config}
	RollConfig *ElastigroupAzureUpdatePolicyRollConfig `field:"optional" json:"rollConfig" yaml:"rollConfig"`
}

