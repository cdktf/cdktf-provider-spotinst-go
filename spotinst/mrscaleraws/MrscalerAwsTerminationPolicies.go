// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mrscaleraws


type MrscalerAwsTerminationPolicies struct {
	// statements block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.224.0/docs/resources/mrscaler_aws#statements MrscalerAws#statements}
	Statements interface{} `field:"required" json:"statements" yaml:"statements"`
}

