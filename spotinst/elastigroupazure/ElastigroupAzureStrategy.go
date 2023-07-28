package elastigroupazure


type ElastigroupAzureStrategy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.130.0/docs/resources/elastigroup_azure#draining_timeout ElastigroupAzure#draining_timeout}.
	DrainingTimeout *float64 `field:"optional" json:"drainingTimeout" yaml:"drainingTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.130.0/docs/resources/elastigroup_azure#low_priority_percentage ElastigroupAzure#low_priority_percentage}.
	LowPriorityPercentage *float64 `field:"optional" json:"lowPriorityPercentage" yaml:"lowPriorityPercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.130.0/docs/resources/elastigroup_azure#od_count ElastigroupAzure#od_count}.
	OdCount *float64 `field:"optional" json:"odCount" yaml:"odCount"`
}

