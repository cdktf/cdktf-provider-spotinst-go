package elastigroupaws


type ElastigroupAwsItfLoadBalancerListenerRule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_aws#rule_arn ElastigroupAws#rule_arn}.
	RuleArn *string `field:"required" json:"ruleArn" yaml:"ruleArn"`
	// static_target_group block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_aws#static_target_group ElastigroupAws#static_target_group}
	StaticTargetGroup *ElastigroupAwsItfLoadBalancerListenerRuleStaticTargetGroup `field:"optional" json:"staticTargetGroup" yaml:"staticTargetGroup"`
}

