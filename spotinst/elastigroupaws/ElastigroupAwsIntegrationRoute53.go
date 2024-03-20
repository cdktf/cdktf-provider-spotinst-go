// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastigroupaws


type ElastigroupAwsIntegrationRoute53 struct {
	// domains block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.165.0/docs/resources/elastigroup_aws#domains ElastigroupAws#domains}
	Domains interface{} `field:"required" json:"domains" yaml:"domains"`
}

