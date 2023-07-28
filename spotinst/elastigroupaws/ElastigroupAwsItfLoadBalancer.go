package elastigroupaws


type ElastigroupAwsItfLoadBalancer struct {
	// listener_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.130.0/docs/resources/elastigroup_aws#listener_rule ElastigroupAws#listener_rule}
	ListenerRule interface{} `field:"required" json:"listenerRule" yaml:"listenerRule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.130.0/docs/resources/elastigroup_aws#load_balancer_arn ElastigroupAws#load_balancer_arn}.
	LoadBalancerArn *string `field:"required" json:"loadBalancerArn" yaml:"loadBalancerArn"`
}

