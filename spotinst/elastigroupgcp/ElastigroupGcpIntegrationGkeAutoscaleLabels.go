// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupgcp


type ElastigroupGcpIntegrationGkeAutoscaleLabels struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.172.3/docs/resources/elastigroup_gcp#key ElastigroupGcp#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.172.3/docs/resources/elastigroup_gcp#value ElastigroupGcp#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

