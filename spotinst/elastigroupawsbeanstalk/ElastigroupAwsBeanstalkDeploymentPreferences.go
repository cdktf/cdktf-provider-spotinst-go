// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupawsbeanstalk


type ElastigroupAwsBeanstalkDeploymentPreferences struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.225.1/docs/resources/elastigroup_aws_beanstalk#automatic_roll ElastigroupAwsBeanstalk#automatic_roll}.
	AutomaticRoll interface{} `field:"optional" json:"automaticRoll" yaml:"automaticRoll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.225.1/docs/resources/elastigroup_aws_beanstalk#batch_size_percentage ElastigroupAwsBeanstalk#batch_size_percentage}.
	BatchSizePercentage *float64 `field:"optional" json:"batchSizePercentage" yaml:"batchSizePercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.225.1/docs/resources/elastigroup_aws_beanstalk#grace_period ElastigroupAwsBeanstalk#grace_period}.
	GracePeriod *float64 `field:"optional" json:"gracePeriod" yaml:"gracePeriod"`
	// strategy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.225.1/docs/resources/elastigroup_aws_beanstalk#strategy ElastigroupAwsBeanstalk#strategy}
	Strategy interface{} `field:"optional" json:"strategy" yaml:"strategy"`
}

