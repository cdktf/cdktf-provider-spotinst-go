// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupgcp


type ElastigroupGcpScalingUpPolicyDimensions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.224.1/docs/resources/elastigroup_gcp#name ElastigroupGcp#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.224.1/docs/resources/elastigroup_gcp#value ElastigroupGcp#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

