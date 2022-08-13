// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type ElastigroupAwsItfLoadBalancer struct {
	// listener_rule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_aws#listener_rule ElastigroupAws#listener_rule}
	ListenerRule interface{} `field:"required" json:"listenerRule" yaml:"listenerRule"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_aws#load_balancer_arn ElastigroupAws#load_balancer_arn}.
	LoadBalancerArn *string `field:"required" json:"loadBalancerArn" yaml:"loadBalancerArn"`
}

