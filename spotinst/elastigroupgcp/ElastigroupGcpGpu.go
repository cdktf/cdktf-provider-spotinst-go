// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupgcp


type ElastigroupGcpGpu struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.228.0/docs/resources/elastigroup_gcp#count ElastigroupGcp#count}.
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.228.0/docs/resources/elastigroup_gcp#type ElastigroupGcp#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

