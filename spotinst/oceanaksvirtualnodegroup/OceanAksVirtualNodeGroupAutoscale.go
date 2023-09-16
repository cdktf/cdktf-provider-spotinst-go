// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaksvirtualnodegroup


type OceanAksVirtualNodeGroupAutoscale struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.139.0/docs/resources/ocean_aks_virtual_node_group#auto_headroom_percentage OceanAksVirtualNodeGroup#auto_headroom_percentage}.
	AutoHeadroomPercentage *float64 `field:"optional" json:"autoHeadroomPercentage" yaml:"autoHeadroomPercentage"`
	// autoscale_headroom block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.139.0/docs/resources/ocean_aks_virtual_node_group#autoscale_headroom OceanAksVirtualNodeGroup#autoscale_headroom}
	AutoscaleHeadroom interface{} `field:"optional" json:"autoscaleHeadroom" yaml:"autoscaleHeadroom"`
}

