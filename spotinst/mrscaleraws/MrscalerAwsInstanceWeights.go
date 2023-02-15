package mrscaleraws


type MrscalerAwsInstanceWeights struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/mrscaler_aws#instance_type MrscalerAws#instance_type}.
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/mrscaler_aws#weighted_capacity MrscalerAws#weighted_capacity}.
	WeightedCapacity *float64 `field:"required" json:"weightedCapacity" yaml:"weightedCapacity"`
}

