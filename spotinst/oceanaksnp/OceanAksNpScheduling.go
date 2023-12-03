// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oceanaksnp


type OceanAksNpScheduling struct {
	// shutdown_hours block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/spotinst/spotinst/1.151.0/docs/resources/ocean_aks_np#shutdown_hours OceanAksNp#shutdown_hours}
	ShutdownHours *OceanAksNpSchedulingShutdownHours `field:"optional" json:"shutdownHours" yaml:"shutdownHours"`
}

