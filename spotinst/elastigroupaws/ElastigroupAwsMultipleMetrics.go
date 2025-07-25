// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupaws


type ElastigroupAwsMultipleMetrics struct {
	// expressions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.224.1/docs/resources/elastigroup_aws#expressions ElastigroupAws#expressions}
	Expressions interface{} `field:"optional" json:"expressions" yaml:"expressions"`
	// metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.224.1/docs/resources/elastigroup_aws#metrics ElastigroupAws#metrics}
	Metrics interface{} `field:"optional" json:"metrics" yaml:"metrics"`
}

