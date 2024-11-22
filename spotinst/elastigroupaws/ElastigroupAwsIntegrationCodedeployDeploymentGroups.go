// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupaws


type ElastigroupAwsIntegrationCodedeployDeploymentGroups struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.0/docs/resources/elastigroup_aws#application_name ElastigroupAws#application_name}.
	ApplicationName *string `field:"required" json:"applicationName" yaml:"applicationName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.199.0/docs/resources/elastigroup_aws#deployment_group_name ElastigroupAws#deployment_group_name}.
	DeploymentGroupName *string `field:"required" json:"deploymentGroupName" yaml:"deploymentGroupName"`
}

