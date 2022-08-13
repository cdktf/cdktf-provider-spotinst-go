// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type ElastigroupAwsStatefulInstanceAction struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_aws#stateful_instance_id ElastigroupAws#stateful_instance_id}.
	StatefulInstanceId *string `field:"required" json:"statefulInstanceId" yaml:"statefulInstanceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_aws#type ElastigroupAws#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

