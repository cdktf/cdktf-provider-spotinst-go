package elastigroupaws


type ElastigroupAwsIntegrationBeanstalkDeploymentPreferences struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.125.0/docs/resources/elastigroup_aws#automatic_roll ElastigroupAws#automatic_roll}.
	AutomaticRoll interface{} `field:"optional" json:"automaticRoll" yaml:"automaticRoll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.125.0/docs/resources/elastigroup_aws#batch_size_percentage ElastigroupAws#batch_size_percentage}.
	BatchSizePercentage *float64 `field:"optional" json:"batchSizePercentage" yaml:"batchSizePercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.125.0/docs/resources/elastigroup_aws#grace_period ElastigroupAws#grace_period}.
	GracePeriod *float64 `field:"optional" json:"gracePeriod" yaml:"gracePeriod"`
	// strategy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.125.0/docs/resources/elastigroup_aws#strategy ElastigroupAws#strategy}
	Strategy *ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesStrategy `field:"optional" json:"strategy" yaml:"strategy"`
}

