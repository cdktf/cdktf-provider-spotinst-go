// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package multaibalancer


type MultaiBalancerTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.139.0/docs/resources/multai_balancer#key MultaiBalancer#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.139.0/docs/resources/multai_balancer#value MultaiBalancer#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

