package elastigroupaws


type ElastigroupAwsIntegrationBeanstalkDeploymentPreferencesStrategy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_aws#action ElastigroupAws#action}.
	Action *string `field:"optional" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_aws#should_drain_instances ElastigroupAws#should_drain_instances}.
	ShouldDrainInstances interface{} `field:"optional" json:"shouldDrainInstances" yaml:"shouldDrainInstances"`
}

