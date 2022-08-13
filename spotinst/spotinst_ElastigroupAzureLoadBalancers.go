// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type ElastigroupAzureLoadBalancers struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_azure#type ElastigroupAzure#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_azure#auto_weight ElastigroupAzure#auto_weight}.
	AutoWeight interface{} `field:"optional" json:"autoWeight" yaml:"autoWeight"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_azure#balancer_id ElastigroupAzure#balancer_id}.
	BalancerId *string `field:"optional" json:"balancerId" yaml:"balancerId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_azure#target_set_id ElastigroupAzure#target_set_id}.
	TargetSetId *string `field:"optional" json:"targetSetId" yaml:"targetSetId"`
}

