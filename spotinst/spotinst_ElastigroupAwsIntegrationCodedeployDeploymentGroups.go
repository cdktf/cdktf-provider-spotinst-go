// Prebuilt spotinst Provider for Terraform CDK (cdktf)
package spotinst


type ElastigroupAwsIntegrationCodedeployDeploymentGroups struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_aws#application_name ElastigroupAws#application_name}.
	ApplicationName *string `field:"required" json:"applicationName" yaml:"applicationName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/spotinst/r/elastigroup_aws#deployment_group_name ElastigroupAws#deployment_group_name}.
	DeploymentGroupName *string `field:"required" json:"deploymentGroupName" yaml:"deploymentGroupName"`
}

