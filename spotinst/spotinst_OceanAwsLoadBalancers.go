// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type OceanAwsLoadBalancers struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aws#arn OceanAws#arn}.
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aws#name OceanAws#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/ocean_aws#type OceanAws#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

