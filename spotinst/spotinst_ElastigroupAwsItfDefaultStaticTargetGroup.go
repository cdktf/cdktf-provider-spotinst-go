// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type ElastigroupAwsItfDefaultStaticTargetGroup struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_aws#arn ElastigroupAws#arn}.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_aws#percentage ElastigroupAws#percentage}.
	Percentage *float64 `field:"required" json:"percentage" yaml:"percentage"`
}

