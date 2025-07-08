// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaksnpvirtualnodegroup


type OceanAksNpVirtualNodeGroupScheduling struct {
	// shutdown_hours block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.222.1/docs/resources/ocean_aks_np_virtual_node_group#shutdown_hours OceanAksNpVirtualNodeGroup#shutdown_hours}
	ShutdownHours *OceanAksNpVirtualNodeGroupSchedulingShutdownHours `field:"optional" json:"shutdownHours" yaml:"shutdownHours"`
}

