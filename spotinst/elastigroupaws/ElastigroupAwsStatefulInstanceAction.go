// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupaws


type ElastigroupAwsStatefulInstanceAction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.157.0/docs/resources/elastigroup_aws#stateful_instance_id ElastigroupAws#stateful_instance_id}.
	StatefulInstanceId *string `field:"required" json:"statefulInstanceId" yaml:"statefulInstanceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.157.0/docs/resources/elastigroup_aws#type ElastigroupAws#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

