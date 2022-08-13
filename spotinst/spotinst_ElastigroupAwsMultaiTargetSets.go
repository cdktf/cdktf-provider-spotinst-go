// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type ElastigroupAwsMultaiTargetSets struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_aws#balancer_id ElastigroupAws#balancer_id}.
	BalancerId *string `field:"required" json:"balancerId" yaml:"balancerId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_aws#target_set_id ElastigroupAws#target_set_id}.
	TargetSetId *string `field:"required" json:"targetSetId" yaml:"targetSetId"`
}

