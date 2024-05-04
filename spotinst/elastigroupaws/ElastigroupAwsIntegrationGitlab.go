// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupaws


type ElastigroupAwsIntegrationGitlab struct {
	// runner block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.171.3/docs/resources/elastigroup_aws#runner ElastigroupAws#runner}
	Runner *ElastigroupAwsIntegrationGitlabRunner `field:"optional" json:"runner" yaml:"runner"`
}

