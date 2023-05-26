package managedinstanceaws


type ManagedInstanceAwsLoadBalancers struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.119.1/docs/resources/managed_instance_aws#type ManagedInstanceAws#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.119.1/docs/resources/managed_instance_aws#arn ManagedInstanceAws#arn}.
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.119.1/docs/resources/managed_instance_aws#auto_weight ManagedInstanceAws#auto_weight}.
	AutoWeight interface{} `field:"optional" json:"autoWeight" yaml:"autoWeight"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.119.1/docs/resources/managed_instance_aws#az_awareness ManagedInstanceAws#az_awareness}.
	AzAwareness interface{} `field:"optional" json:"azAwareness" yaml:"azAwareness"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.119.1/docs/resources/managed_instance_aws#balancer_id ManagedInstanceAws#balancer_id}.
	BalancerId *string `field:"optional" json:"balancerId" yaml:"balancerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.119.1/docs/resources/managed_instance_aws#name ManagedInstanceAws#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.119.1/docs/resources/managed_instance_aws#target_set_id ManagedInstanceAws#target_set_id}.
	TargetSetId *string `field:"optional" json:"targetSetId" yaml:"targetSetId"`
}

