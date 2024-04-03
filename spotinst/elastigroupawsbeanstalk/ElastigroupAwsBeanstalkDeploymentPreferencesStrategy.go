// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupawsbeanstalk


type ElastigroupAwsBeanstalkDeploymentPreferencesStrategy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.168.1/docs/resources/elastigroup_aws_beanstalk#action ElastigroupAwsBeanstalk#action}.
	Action *string `field:"optional" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.168.1/docs/resources/elastigroup_aws_beanstalk#should_drain_instances ElastigroupAwsBeanstalk#should_drain_instances}.
	ShouldDrainInstances interface{} `field:"optional" json:"shouldDrainInstances" yaml:"shouldDrainInstances"`
}

