// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mrscaleraws


type MrscalerAwsInstanceWeights struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.224.1/docs/resources/mrscaler_aws#instance_type MrscalerAws#instance_type}.
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.224.1/docs/resources/mrscaler_aws#weighted_capacity MrscalerAws#weighted_capacity}.
	WeightedCapacity *float64 `field:"required" json:"weightedCapacity" yaml:"weightedCapacity"`
}

