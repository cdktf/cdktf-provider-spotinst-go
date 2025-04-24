// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupgcp


type ElastigroupGcpRevertToPreemptible struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.216.2/docs/resources/elastigroup_gcp#perform_at ElastigroupGcp#perform_at}.
	PerformAt *string `field:"required" json:"performAt" yaml:"performAt"`
}

