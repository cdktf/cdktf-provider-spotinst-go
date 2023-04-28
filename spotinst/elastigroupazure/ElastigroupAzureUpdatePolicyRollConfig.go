package elastigroupazure


type ElastigroupAzureUpdatePolicyRollConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.115.0/docs/resources/elastigroup_azure#batch_size_percentage ElastigroupAzure#batch_size_percentage}.
	BatchSizePercentage *float64 `field:"required" json:"batchSizePercentage" yaml:"batchSizePercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.115.0/docs/resources/elastigroup_azure#grace_period ElastigroupAzure#grace_period}.
	GracePeriod *float64 `field:"optional" json:"gracePeriod" yaml:"gracePeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.115.0/docs/resources/elastigroup_azure#health_check_type ElastigroupAzure#health_check_type}.
	HealthCheckType *string `field:"optional" json:"healthCheckType" yaml:"healthCheckType"`
}

