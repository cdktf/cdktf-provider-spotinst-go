// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupaws


type ElastigroupAwsIntegrationNomadAutoscaleConstraints struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.153.0/docs/resources/elastigroup_aws#key ElastigroupAws#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.153.0/docs/resources/elastigroup_aws#value ElastigroupAws#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

