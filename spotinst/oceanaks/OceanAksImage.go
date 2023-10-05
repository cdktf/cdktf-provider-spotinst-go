// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaks


type OceanAksImage struct {
	// marketplace block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.143.0/docs/resources/ocean_aks#marketplace OceanAks#marketplace}
	Marketplace interface{} `field:"optional" json:"marketplace" yaml:"marketplace"`
}

