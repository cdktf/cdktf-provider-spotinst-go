// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mrscaleraws


type MrscalerAwsProvisioningTimeout struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.142.0/docs/resources/mrscaler_aws#timeout MrscalerAws#timeout}.
	Timeout *float64 `field:"required" json:"timeout" yaml:"timeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.142.0/docs/resources/mrscaler_aws#timeout_action MrscalerAws#timeout_action}.
	TimeoutAction *string `field:"required" json:"timeoutAction" yaml:"timeoutAction"`
}

