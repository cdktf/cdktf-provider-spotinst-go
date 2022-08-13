// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type OceanAwsUpdatePolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aws#should_roll OceanAws#should_roll}.
	ShouldRoll interface{} `field:"required" json:"shouldRoll" yaml:"shouldRoll"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aws#auto_apply_tags OceanAws#auto_apply_tags}.
	AutoApplyTags interface{} `field:"optional" json:"autoApplyTags" yaml:"autoApplyTags"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aws#conditioned_roll OceanAws#conditioned_roll}.
	ConditionedRoll interface{} `field:"optional" json:"conditionedRoll" yaml:"conditionedRoll"`
	// roll_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aws#roll_config OceanAws#roll_config}
	RollConfig *OceanAwsUpdatePolicyRollConfig `field:"optional" json:"rollConfig" yaml:"rollConfig"`
}

