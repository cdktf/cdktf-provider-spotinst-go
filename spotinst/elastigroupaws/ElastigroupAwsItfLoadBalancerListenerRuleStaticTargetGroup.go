package elastigroupaws


type ElastigroupAwsItfLoadBalancerListenerRuleStaticTargetGroup struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.127.0/docs/resources/elastigroup_aws#arn ElastigroupAws#arn}.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.127.0/docs/resources/elastigroup_aws#percentage ElastigroupAws#percentage}.
	Percentage *float64 `field:"required" json:"percentage" yaml:"percentage"`
}

