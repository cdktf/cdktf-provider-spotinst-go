// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupaws


type ElastigroupAwsItfLoadBalancerListenerRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.153.0/docs/resources/elastigroup_aws#rule_arn ElastigroupAws#rule_arn}.
	RuleArn *string `field:"required" json:"ruleArn" yaml:"ruleArn"`
	// static_target_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.153.0/docs/resources/elastigroup_aws#static_target_group ElastigroupAws#static_target_group}
	StaticTargetGroup *ElastigroupAwsItfLoadBalancerListenerRuleStaticTargetGroup `field:"optional" json:"staticTargetGroup" yaml:"staticTargetGroup"`
}

