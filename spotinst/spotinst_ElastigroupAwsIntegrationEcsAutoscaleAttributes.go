// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type ElastigroupAwsIntegrationEcsAutoscaleAttributes struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_aws#key ElastigroupAws#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_aws#value ElastigroupAws#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

