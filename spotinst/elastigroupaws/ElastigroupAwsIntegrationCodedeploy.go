package elastigroupaws


type ElastigroupAwsIntegrationCodedeploy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_aws#cleanup_on_failure ElastigroupAws#cleanup_on_failure}.
	CleanupOnFailure interface{} `field:"required" json:"cleanupOnFailure" yaml:"cleanupOnFailure"`
	// deployment_groups block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_aws#deployment_groups ElastigroupAws#deployment_groups}
	DeploymentGroups interface{} `field:"required" json:"deploymentGroups" yaml:"deploymentGroups"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_aws#terminate_instance_on_failure ElastigroupAws#terminate_instance_on_failure}.
	TerminateInstanceOnFailure interface{} `field:"required" json:"terminateInstanceOnFailure" yaml:"terminateInstanceOnFailure"`
}
