// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type ElastigroupAzureStrategy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_azure#draining_timeout ElastigroupAzure#draining_timeout}.
	DrainingTimeout *float64 `field:"optional" json:"drainingTimeout" yaml:"drainingTimeout"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_azure#low_priority_percentage ElastigroupAzure#low_priority_percentage}.
	LowPriorityPercentage *float64 `field:"optional" json:"lowPriorityPercentage" yaml:"lowPriorityPercentage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_azure#od_count ElastigroupAzure#od_count}.
	OdCount *float64 `field:"optional" json:"odCount" yaml:"odCount"`
}

