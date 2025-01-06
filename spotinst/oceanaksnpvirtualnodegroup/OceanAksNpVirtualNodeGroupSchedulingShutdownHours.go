// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaksnpvirtualnodegroup


type OceanAksNpVirtualNodeGroupSchedulingShutdownHours struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.205.0/docs/resources/ocean_aks_np_virtual_node_group#is_enabled OceanAksNpVirtualNodeGroup#is_enabled}.
	IsEnabled interface{} `field:"optional" json:"isEnabled" yaml:"isEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.205.0/docs/resources/ocean_aks_np_virtual_node_group#time_windows OceanAksNpVirtualNodeGroup#time_windows}.
	TimeWindows *[]*string `field:"optional" json:"timeWindows" yaml:"timeWindows"`
}

