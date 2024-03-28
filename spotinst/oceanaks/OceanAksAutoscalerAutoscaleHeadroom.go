// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaks


type OceanAksAutoscalerAutoscaleHeadroom struct {
	// automatic block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.165.1/docs/resources/ocean_aks#automatic OceanAks#automatic}
	Automatic *OceanAksAutoscalerAutoscaleHeadroomAutomatic `field:"optional" json:"automatic" yaml:"automatic"`
}

