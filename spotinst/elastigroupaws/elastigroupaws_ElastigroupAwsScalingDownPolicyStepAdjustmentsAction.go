package elastigroupaws


type ElastigroupAwsScalingDownPolicyStepAdjustmentsAction struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_aws#type ElastigroupAws#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_aws#adjustment ElastigroupAws#adjustment}.
	Adjustment *string `field:"optional" json:"adjustment" yaml:"adjustment"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_aws#maximum ElastigroupAws#maximum}.
	Maximum *string `field:"optional" json:"maximum" yaml:"maximum"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_aws#max_target_capacity ElastigroupAws#max_target_capacity}.
	MaxTargetCapacity *string `field:"optional" json:"maxTargetCapacity" yaml:"maxTargetCapacity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_aws#minimum ElastigroupAws#minimum}.
	Minimum *string `field:"optional" json:"minimum" yaml:"minimum"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_aws#min_target_capacity ElastigroupAws#min_target_capacity}.
	MinTargetCapacity *string `field:"optional" json:"minTargetCapacity" yaml:"minTargetCapacity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_aws#target ElastigroupAws#target}.
	Target *string `field:"optional" json:"target" yaml:"target"`
}
