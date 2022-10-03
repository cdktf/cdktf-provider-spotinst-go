package elastigroupaws


type ElastigroupAwsIntegrationBeanstalk struct {
	// deployment_preferences block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_aws#deployment_preferences ElastigroupAws#deployment_preferences}
	DeploymentPreferences *ElastigroupAwsIntegrationBeanstalkDeploymentPreferences `field:"optional" json:"deploymentPreferences" yaml:"deploymentPreferences"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_aws#environment_id ElastigroupAws#environment_id}.
	EnvironmentId *string `field:"optional" json:"environmentId" yaml:"environmentId"`
	// managed_actions block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_aws#managed_actions ElastigroupAws#managed_actions}
	ManagedActions *ElastigroupAwsIntegrationBeanstalkManagedActions `field:"optional" json:"managedActions" yaml:"managedActions"`
}

